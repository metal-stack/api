package request

import (
	"context"
	"fmt"
	"slices"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/metal-stack/api/go/permissions"
)

type (
	// tokenPermission represents the unflattened permissions from a token.
	// It maps the method to the allowed subjects per method.
	// The subject will be set to "*" in case of admin roles.
	// This works because a certain method can either be tenant or project scoped.
	// Therefore a single subject is enough to decide
	tokenPermissions map[string]map[string]bool
)

func (a *authorizer) getTokenPermissions(ctx context.Context, token *apiv2.Token) (tokenPermissions, error) {
	tp := tokenPermissions{}
	if token == nil {
		return nil, nil
	}

	var (
		projectRoles       map[string]apiv2.ProjectRole
		tenantRoles        map[string]apiv2.TenantRole
		adminRole          *apiv2.AdminRole
		servicePermissions = permissions.GetServicePermissions()
	)

	pat, err := a.projectsAndTenantsGetter(ctx, token.User)
	if err != nil {
		return nil, err
	}

	projectRoles = pat.ProjectRoles
	tenantRoles = pat.TenantRoles
	a.log.Debug("decide", "adminsubjects", a.adminSubjects, "user", token.User)
	if slices.Contains(a.adminSubjects, token.User) {
		// we do not store admin roles in the masterdata-api, but we can use this from the static configuration passed to the api-server
		adminRole = token.AdminRole
	}

	if token.TokenType == apiv2.TokenType_TOKEN_TYPE_USER {
		// as we do not store roles in the user token, we set the roles from the information in the masterdata-db
		token.ProjectRoles = projectRoles
		token.TenantRoles = tenantRoles
		token.AdminRole = adminRole
		// user tokens should never have permissions cause they are not stored in the masterdata-db
		token.Permissions = nil
	}

	// Admin Roles
	if token.AdminRole != nil {
		var allMethods []string
		for method := range servicePermissions.Methods {
			allMethods = append(allMethods, method)
		}
		slices.Sort(allMethods)

		switch *token.AdminRole {
		case apiv2.AdminRole_ADMIN_ROLE_EDITOR:
			for _, method := range allMethods {
				tp[method] = map[string]bool{"*": true}
			}
			// Return here because all methods are allowed with all permissions
			return tp, nil

		case apiv2.AdminRole_ADMIN_ROLE_VIEWER:
			adminEditorMethods := servicePermissions.Roles.Admin[apiv2.AdminRole_ADMIN_ROLE_EDITOR.Enum().String()]
			adminViewerMethods := servicePermissions.Roles.Admin[apiv2.AdminRole_ADMIN_ROLE_VIEWER.Enum().String()]
			nonViewerMethods := slices.DeleteFunc(adminEditorMethods, func(method string) bool {
				return slices.Contains(adminViewerMethods, method)
			})
			remainingMethods := slices.DeleteFunc(allMethods, func(method string) bool {
				return slices.Contains(nonViewerMethods, method)
			})
			for _, method := range remainingMethods {
				tp[method] = map[string]bool{"*": true}
			}
			// Do not return here because it might be that some permissions are granted later

		default:
			return nil, fmt.Errorf("given admin role:%s is not valid", *token.AdminRole)
		}
	}

	// Permission
	for _, permission := range token.Permissions {
		subject := permission.Subject
		for _, method := range permission.Methods {
			if _, ok := tp[method]; !ok {
				tp[method] = map[string]bool{}
			}
			tp[method][subject] = true
		}
	}

	// Tenant Roles
	for subject, role := range token.TenantRoles {
		tenantMethods := servicePermissions.Roles.Tenant[role.Enum().String()]
		for _, method := range tenantMethods {
			if _, ok := tp[method]; !ok {
				tp[method] = map[string]bool{}
			}
			tp[method][subject] = true
		}
	}

	// Project Roles
	for subject, role := range token.ProjectRoles {
		projectMethods := servicePermissions.Roles.Project[role.Enum().String()]
		for _, method := range projectMethods {
			if _, ok := tp[method]; !ok {
				tp[method] = map[string]bool{}
			}
			tp[method][subject] = true
		}
	}

	// Public and Self Methods only on user tokens
	if token.TokenType == apiv2.TokenType_TOKEN_TYPE_USER {
		for method := range servicePermissions.Visibility.Public {
			if _, ok := tp[method]; !ok {
				tp[method] = map[string]bool{}
			}
			tp[method]["*"] = true
		}

		for method := range servicePermissions.Visibility.Self {
			if _, ok := tp[method]; !ok {
				tp[method] = map[string]bool{}
			}
			// Subjects of self service must also be validated inside the service implementation
			tp[method]["*"] = true
		}
	}

	// TODO infra

	return tp, nil
}
