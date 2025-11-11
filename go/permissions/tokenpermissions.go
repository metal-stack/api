package permissions

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"slices"

	"connectrpc.com/connect"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
)

type (
	ProjectsAndTenants struct {
		Projects      []*apiv2.Project
		Tenants       []*apiv2.Tenant
		DefaultTenant *apiv2.Tenant
		ProjectRoles  map[string]apiv2.ProjectRole
		TenantRoles   map[string]apiv2.TenantRole
	}

	projectsAndTenantsGetter func(ctx context.Context, userId string) (*ProjectsAndTenants, error)

	// tokenPermission represents the unflattened permissions from a token.
	// It maps the method to the allowed subjects per method.
	// The subject will be set to "*" in case of admin roles.
	// This works because a certain method can either be tenant or project scoped.
	// Therefore a single subject is enough to decide
	tokenPermissions map[string]map[string]bool

	authorizer struct {
		log                      *slog.Logger
		adminSubjects            []string
		projectsAndTenantsGetter func(ctx context.Context, userId string) (*ProjectsAndTenants, error)
	}

	// Authorizer provides methods to authorize requests with a given token
	Authorizer interface {
		// Allowed checks if with the given token, access to method with this subject is allowed.
		// If the access is not allowed, a PermissionDenied Error is returned with a proper error message.
		// req is only fully populated after a intercepter call.
		Allowed(ctx context.Context, token *apiv2.Token, req connect.AnyRequest) error
	}
)

func NewAuthorizer(log *slog.Logger, adminSubjects []string, patg projectsAndTenantsGetter) Authorizer {
	return &authorizer{
		log:                      log,
		adminSubjects:            adminSubjects,
		projectsAndTenantsGetter: patg,
	}
}

func (a *authorizer) Allowed(ctx context.Context, token *apiv2.Token, req connect.AnyRequest) error {
	var (
		method  = req.Spec().Procedure
		subject string
	)
	if req == nil {
		return connect.NewError(connect.CodeInternal, fmt.Errorf("request is nil"))
	}

	a.log.Info("allowed", "req", req.Spec())

	if IsProjectScope(req) {
		project, ok := GetProjectFromRequest(req)
		if ok {
			subject = project
		}
	}

	if IsTenantScope(req) {
		tenant, ok := GetTenantFromRequest(req)
		if ok {
			subject = tenant
		}
	}

	return a.allowed(ctx, token, method, subject)
}

// Allowed implements Authorizer.
func (a *authorizer) allowed(ctx context.Context, token *apiv2.Token, method string, subject string) error {
	permissions, err := a.getTokenPermissions(ctx, token)
	if err != nil {
		return connect.NewError(connect.CodeInternal, err)
	}
	if permissions == nil {
		return connect.NewError(connect.CodePermissionDenied, errors.New("no permissions found in token"))
	}

	subjects, ok := permissions[method]
	if !ok {
		return connect.NewError(connect.CodePermissionDenied, fmt.Errorf("access to:%q is not allowed because it is not part of the token permissions", method))
	}

	if _, allSubjectsAllowed := subjects["*"]; allSubjectsAllowed {
		// This token contains permissions to access this method regardless of subject
		return nil
	}

	if _, subjectAllowed := subjects[subject]; !subjectAllowed {
		return connect.NewError(connect.CodePermissionDenied, fmt.Errorf("access to:%q with subject:%q is not allowed because it is not part of the token permissions", method, subject))
	}

	return nil
}

func (a *authorizer) getTokenPermissions(ctx context.Context, token *apiv2.Token) (tokenPermissions, error) {
	tp := tokenPermissions{}
	if token == nil {
		return nil, nil
	}

	var (
		projectRoles       map[string]apiv2.ProjectRole
		tenantRoles        map[string]apiv2.TenantRole
		adminRole          *apiv2.AdminRole
		servicePermissions = GetServicePermissions()
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
