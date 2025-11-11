package request

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/metal-stack/api/go/permissions"
)

type (

	// ProjectAndTenants contains the content of a user token stored in the database if any.
	ProjectsAndTenants struct {
		Projects      []*apiv2.Project
		Tenants       []*apiv2.Tenant
		DefaultTenant *apiv2.Tenant
		ProjectRoles  map[string]apiv2.ProjectRole
		TenantRoles   map[string]apiv2.TenantRole
	}

	// helper to be able to test user token content fetched from the database
	projectsAndTenantsGetter func(ctx context.Context, userId string) (*ProjectsAndTenants, error)

	authorizer struct {
		log                      *slog.Logger
		adminSubjects            []string
		projectsAndTenantsGetter func(ctx context.Context, userId string) (*ProjectsAndTenants, error)
	}

	// Authorizer provides methods to authorize requests with a given token
	Authorizer interface {
		// Allowed checks if with the given token the request is allowed.
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

	if permissions.IsProjectScope(req) {
		project, ok := permissions.GetProjectFromRequest(req)
		if ok {
			subject = project
		}
	}

	if permissions.IsTenantScope(req) {
		tenant, ok := permissions.GetTenantFromRequest(req)
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
