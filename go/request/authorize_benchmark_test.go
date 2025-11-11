package request

import (
	"context"
	"log/slog"
	"testing"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
)

func Benchmark_allow(b *testing.B) {
	a := &authorizer{
		log:           slog.Default(),
		adminSubjects: []string{},
	}
	a.projectsAndTenantsGetter = func(ctx context.Context, userId string) (*ProjectsAndTenants, error) {
		return &ProjectsAndTenants{}, nil
	}

	for b.Loop() {
		token := &apiv2.Token{
			TokenType: apiv2.TokenType_TOKEN_TYPE_API,
			Permissions: []*apiv2.MethodPermission{
				{Subject: "project-a", Methods: []string{"/metalstack.api.v2.IPService/Get"}},
			},
		}
		message := "/metalstack.api.v2.IPService/Get"
		subject := "project-a"
		a.allowed(b.Context(), token, message, subject)
	}
}
