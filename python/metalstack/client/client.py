# Code generated generate_clients.go. DO NOT EDIT.

from contextlib import contextmanager
from connecpy.context import ClientContext

import metalstack.admin.v2.filesystem_connecpy as admin_filesystem_connecpy
import metalstack.admin.v2.image_connecpy as admin_image_connecpy
import metalstack.admin.v2.ip_connecpy as admin_ip_connecpy
import metalstack.admin.v2.network_connecpy as admin_network_connecpy
import metalstack.admin.v2.partition_connecpy as admin_partition_connecpy
import metalstack.admin.v2.size_connecpy as admin_size_connecpy
import metalstack.admin.v2.tenant_connecpy as admin_tenant_connecpy
import metalstack.admin.v2.token_connecpy as admin_token_connecpy

import metalstack.api.v2.filesystem_connecpy as api_filesystem_connecpy
import metalstack.api.v2.health_connecpy as api_health_connecpy
import metalstack.api.v2.image_connecpy as api_image_connecpy
import metalstack.api.v2.ip_connecpy as api_ip_connecpy
import metalstack.api.v2.method_connecpy as api_method_connecpy
import metalstack.api.v2.network_connecpy as api_network_connecpy
import metalstack.api.v2.partition_connecpy as api_partition_connecpy
import metalstack.api.v2.project_connecpy as api_project_connecpy
import metalstack.api.v2.size_connecpy as api_size_connecpy
import metalstack.api.v2.tenant_connecpy as api_tenant_connecpy
import metalstack.api.v2.token_connecpy as api_token_connecpy
import metalstack.api.v2.user_connecpy as api_user_connecpy
import metalstack.api.v2.version_connecpy as api_version_connecpy

import metalstack.infra.v2.bmc_connecpy as infra_bmc_connecpy



class Driver:
    def __init__(self, baseurl: str, token: str, timeout: int = 10):
        self.baseurl = baseurl
        self.token = token
        self.timeout = timeout

    def admin(self):
        return AdminDriver(baseurl=self.baseurl, token=self.token, timeout=self.timeout)

    def api(self):
        return ApiDriver(baseurl=self.baseurl, token=self.token, timeout=self.timeout)

    def infra(self):
        return InfraDriver(baseurl=self.baseurl, token=self.token, timeout=self.timeout)


class AdminDriver:
    def __init__(self, baseurl: str, token: str, timeout: int = 10):
        self.baseurl = baseurl
        self.token = token
        self.timeout = timeout


    @contextmanager
    def filesystem(self):
        with admin_filesystem_connecpy.FilesystemServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def image(self):
        with admin_image_connecpy.ImageServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def ip(self):
        with admin_ip_connecpy.IPServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def network(self):
        with admin_network_connecpy.NetworkServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def partition(self):
        with admin_partition_connecpy.PartitionServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def size(self):
        with admin_size_connecpy.SizeServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def tenant(self):
        with admin_tenant_connecpy.TenantServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def token(self):
        with admin_token_connecpy.TokenServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

class ApiDriver:
    def __init__(self, baseurl: str, token: str, timeout: int = 10):
        self.baseurl = baseurl
        self.token = token
        self.timeout = timeout


    @contextmanager
    def filesystem(self):
        with api_filesystem_connecpy.FilesystemServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def health(self):
        with api_health_connecpy.HealthServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def image(self):
        with api_image_connecpy.ImageServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def ip(self):
        with api_ip_connecpy.IPServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def method(self):
        with api_method_connecpy.MethodServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def network(self):
        with api_network_connecpy.NetworkServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def partition(self):
        with api_partition_connecpy.PartitionServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def project(self):
        with api_project_connecpy.ProjectServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def size(self):
        with api_size_connecpy.SizeServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def tenant(self):
        with api_tenant_connecpy.TenantServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def token(self):
        with api_token_connecpy.TokenServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def user(self):
        with api_user_connecpy.UserServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def version(self):
        with api_version_connecpy.VersionServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

class InfraDriver:
    def __init__(self, baseurl: str, token: str, timeout: int = 10):
        self.baseurl = baseurl
        self.token = token
        self.timeout = timeout


    @contextmanager
    def bmc(self):
        with infra_bmc_connecpy.BMCServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)



class ClientWrapper:
    def __init__(self, client, token):
        self._client = client
        self._token = token

    def __getattr__(self, name):
        attr = getattr(self._client, name)
        if callable(attr):
            def wrapper(*args, **kwargs):
                kwargs.setdefault("ctx", ClientContext())
                headers = kwargs.get("headers", {})
                headers.setdefault("Authorization", "Bearer " + self._token)
                kwargs["headers"] = headers
                return attr(*args, **kwargs)
            return wrapper
        return attr
