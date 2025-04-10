package main

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"

	sprig "github.com/go-task/slim-sprig/v3"
	v1 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/metal-stack/api/go/permissions"
	"github.com/metal-stack/api/go/tests/protoparser"

	_ "embed"
)

const (
	// serverReflectionInfo1alpha1 is always allowed to access to get a list of exposed services for example with grpcurl
	serverReflectionInfov1alpha1 = "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"
	// serverReflectionInfo is always allowed to access to get a list of exposed services for example with grpcurl
	serverReflectionInfo = "/grpc.reflection.v1.ServerReflection/ServerReflectionInfo"
)

var (
	//go:embed go_servicepermissions.tpl
	servicePermissionsTpl string
	//go:embed go_mock_client.tpl
	mockClientTpl string
	//go:embed go_client.tpl
	clientTpl string
)

type api struct {
	Name     string
	Services []string
	Path     string
}

func main() {
	fmt.Println("generating service permissions")

	perms, err := servicePermissions("../proto")
	if err != nil {
		panic(err)
	}

	err = writeTemplate("../go/permissions/servicepermissions.go", servicePermissionsTpl, perms)
	if err != nil {
		panic(err)
	}

	fmt.Println("generating clients")

	svcs, err := svcs("../proto")
	if err != nil {
		panic(err)
	}

	err = writeTemplate("../go/client/client.go", clientTpl, svcs)
	if err != nil {
		panic(err)
	}

	err = writeTemplate("../go/tests/mock_clients.go", mockClientTpl, svcs)
	if err != nil {
		panic(err)
	}
}

func servicePermissions(root string) (*permissions.ServicePermissions, error) {
	var (
		walk = func(root string) ([]string, error) {
			var files []string
			err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				if strings.HasSuffix(info.Name(), ".proto") {
					files = append(files, path)
				}
				return nil
			})
			return files, err
		}
		roles = permissions.Roles{
			Admin:   permissions.Admin{},
			Infra:   permissions.Infra{},
			Tenant:  permissions.Tenant{},
			Project: permissions.Project{},
		}
		methods    = permissions.Methods{}
		visibility = permissions.Visibility{
			Public: map[string]bool{
				// Allow service reflection to list available methods
				serverReflectionInfov1alpha1: true,
				serverReflectionInfo:         true,
			},
			Self:    map[string]bool{},
			Admin:   map[string]bool{},
			Infra:   map[string]bool{},
			Tenant:  map[string]bool{},
			Project: map[string]bool{},
		}
		auditable = permissions.Auditable{}
		services  = []string{}
	)

	files, err := walk(root)
	if err != nil {
		return nil, err
	}

	for _, filename := range files {
		filename := filename
		fd, err := protoparser.Parse(filename)
		if err != nil {
			return nil, err
		}
		for _, serviceDesc := range fd.GetService() {
			serviceDesc := serviceDesc
			services = append(services, fmt.Sprintf("%s.%s", *fd.Package, *serviceDesc.Name))
			for _, method := range serviceDesc.GetMethod() {
				method := method
				methodName := fmt.Sprintf("/%s.%s/%s", *fd.Package, *serviceDesc.Name, *method.Name)
				methodOpts := method.Options.GetUninterpretedOption()
				for _, methodOpt := range methodOpts {
					methodOpt := methodOpt
					for _, namePart := range methodOpt.Name {
						namePart := namePart
						if !*namePart.IsExtension {
							continue
						}
						auditable[methodName] = true
						// Tenant
						switch *methodOpt.IdentifierValue {
						case v1.TenantRole_TENANT_ROLE_OWNER.String():
							roles.Tenant[v1.TenantRole_TENANT_ROLE_OWNER.String()] = append(roles.Tenant[v1.TenantRole_TENANT_ROLE_OWNER.String()], methodName)
							visibility.Tenant[methodName] = true
						case v1.TenantRole_TENANT_ROLE_EDITOR.String():
							roles.Tenant[v1.TenantRole_TENANT_ROLE_EDITOR.String()] = append(roles.Tenant[v1.TenantRole_TENANT_ROLE_EDITOR.String()], methodName)
							visibility.Tenant[methodName] = true
						case v1.TenantRole_TENANT_ROLE_VIEWER.String():
							roles.Tenant[v1.TenantRole_TENANT_ROLE_VIEWER.String()] = append(roles.Tenant[v1.TenantRole_TENANT_ROLE_VIEWER.String()], methodName)
							visibility.Tenant[methodName] = true
						case v1.TenantRole_TENANT_ROLE_GUEST.String():
							roles.Tenant[v1.TenantRole_TENANT_ROLE_GUEST.String()] = append(roles.Tenant[v1.TenantRole_TENANT_ROLE_GUEST.String()], methodName)
							visibility.Tenant[methodName] = true
						case v1.TenantRole_TENANT_ROLE_UNSPECIFIED.String():
							// noop
						// Project
						case v1.ProjectRole_PROJECT_ROLE_OWNER.String():
							roles.Project[v1.ProjectRole_PROJECT_ROLE_OWNER.String()] = append(roles.Project[v1.ProjectRole_PROJECT_ROLE_OWNER.String()], methodName)
							visibility.Project[methodName] = true
						case v1.ProjectRole_PROJECT_ROLE_EDITOR.String():
							visibility.Project[methodName] = true
							roles.Project[v1.ProjectRole_PROJECT_ROLE_EDITOR.String()] = append(roles.Project[v1.ProjectRole_PROJECT_ROLE_EDITOR.String()], methodName)
						case v1.ProjectRole_PROJECT_ROLE_VIEWER.String():
							visibility.Project[methodName] = true
							roles.Project[v1.ProjectRole_PROJECT_ROLE_VIEWER.String()] = append(roles.Project[v1.ProjectRole_PROJECT_ROLE_VIEWER.String()], methodName)
						case v1.ProjectRole_PROJECT_ROLE_UNSPECIFIED.String():
							// noop
						// Admin
						case v1.AdminRole_ADMIN_ROLE_EDITOR.String():
							roles.Admin[v1.AdminRole_ADMIN_ROLE_EDITOR.String()] = append(roles.Admin[v1.AdminRole_ADMIN_ROLE_EDITOR.String()], methodName)
							visibility.Admin[methodName] = true
						case v1.AdminRole_ADMIN_ROLE_VIEWER.String():
							roles.Admin[v1.AdminRole_ADMIN_ROLE_VIEWER.String()] = append(roles.Admin[v1.AdminRole_ADMIN_ROLE_VIEWER.String()], methodName)
							visibility.Admin[methodName] = true
						case v1.AdminRole_ADMIN_ROLE_UNSPECIFIED.String():
							// noop
						// Infra
						case v1.InfraRole_INFRA_ROLE_EDITOR.String():
							roles.Infra[v1.InfraRole_INFRA_ROLE_EDITOR.String()] = append(roles.Admin[v1.InfraRole_INFRA_ROLE_EDITOR.String()], methodName)
							visibility.Infra[methodName] = true
						case v1.InfraRole_INFRA_ROLE_VIEWER.String():
							roles.Infra[v1.InfraRole_INFRA_ROLE_VIEWER.String()] = append(roles.Admin[v1.InfraRole_INFRA_ROLE_VIEWER.String()], methodName)
							visibility.Infra[methodName] = true
						case v1.InfraRole_INFRA_ROLE_UNSPECIFIED.String():
							// noop
						// Visibility
						case v1.Visibility_VISIBILITY_PUBLIC.String():
							visibility.Public[methodName] = true
						case v1.Visibility_VISIBILITY_SELF.String():
							visibility.Self[methodName] = true
						case v1.Visibility_VISIBILITY_UNSPECIFIED.String():
							// noop
						// Auditable
						case v1.Auditing_AUDITING_EXCLUDED.String():
							auditable[methodName] = false
						case v1.Auditing_AUDITING_INCLUDED.String():
							auditable[methodName] = true
						case v1.Auditing_AUDITING_UNSPECIFIED.String():
							auditable[methodName] = true
						// noop
						default:
							return nil, fmt.Errorf("unknown method identifier value detected:%s", *methodOpt.IdentifierValue)

						}
						methods[methodName] = true
					}
				}
			}
		}
	}
	slices.Sort(services)
	sp := &permissions.ServicePermissions{
		Roles:      roles,
		Methods:    methods,
		Visibility: visibility,
		Auditable:  auditable,
		Services:   services,
	}

	return sp, nil
}

func svcs(root string) (map[string]api, error) {
	var (
		result = map[string]api{}
		walk   = func(root string) ([]string, error) {
			var files []string
			err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				if strings.HasSuffix(info.Name(), ".proto") {
					files = append(files, path)
				}
				return nil
			})
			return files, err
		}
	)

	files, err := walk(root)
	if err != nil {
		return nil, err
	}

	for _, filename := range files {
		fd, err := protoparser.Parse(filename)
		if err != nil {
			return nil, err
		}
		_, name, _ := strings.Cut(*fd.Package, "metalstack.")
		name = strings.ReplaceAll(name, ".", "")

		a, ok := result[name]
		if !ok {
			a = api{
				Name: name,
				Path: path.Dir(strings.TrimPrefix(filename, root)),
			}
		}
		for _, serviceDesc := range fd.GetService() {
			a.Services = append(a.Services, *serviceDesc.Name)
		}
		result[name] = a
	}

	return result, nil
}

func writeTemplate(dest, text string, data any) error {
	t, err := template.New("").Funcs(sprig.FuncMap()).Parse(text)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return err
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	fmt.Println("wrote " + dest)

	return os.WriteFile(dest, p, 0755) // nolint:gosec
}
