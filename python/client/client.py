# Code generated generate_clients.go. DO NOT EDIT.

from contextlib import contextmanager
from connecpy.context import ClientContext
from connecpy.exceptions import ConnecpyServerException

from metalstack.admin.v2 import filesystem_pb2 as admin_filesystem_pb2
from metalstack.admin.v2 import filesystem_connecpy as admin_filesystem_connecpy
from metalstack.admin.v2 import image_pb2 as admin_image_pb2
from metalstack.admin.v2 import image_connecpy as admin_image_connecpy
from metalstack.admin.v2 import ip_pb2 as admin_ip_pb2
from metalstack.admin.v2 import ip_connecpy as admin_ip_connecpy
from metalstack.admin.v2 import network_pb2 as admin_network_pb2
from metalstack.admin.v2 import network_connecpy as admin_network_connecpy
from metalstack.admin.v2 import partition_pb2 as admin_partition_pb2
from metalstack.admin.v2 import partition_connecpy as admin_partition_connecpy
from metalstack.admin.v2 import size_pb2 as admin_size_pb2
from metalstack.admin.v2 import size_connecpy as admin_size_connecpy
from metalstack.admin.v2 import tenant_pb2 as admin_tenant_pb2
from metalstack.admin.v2 import tenant_connecpy as admin_tenant_connecpy
from metalstack.admin.v2 import token_pb2 as admin_token_pb2
from metalstack.admin.v2 import token_connecpy as admin_token_connecpy

from metalstack.api.v2 import filesystem_pb2 as api_filesystem_pb2
from metalstack.api.v2 import filesystem_connecpy as api_filesystem_connecpy
from metalstack.api.v2 import health_pb2 as api_health_pb2
from metalstack.api.v2 import health_connecpy as api_health_connecpy
from metalstack.api.v2 import image_pb2 as api_image_pb2
from metalstack.api.v2 import image_connecpy as api_image_connecpy
from metalstack.api.v2 import ip_pb2 as api_ip_pb2
from metalstack.api.v2 import ip_connecpy as api_ip_connecpy
from metalstack.api.v2 import method_pb2 as api_method_pb2
from metalstack.api.v2 import method_connecpy as api_method_connecpy
from metalstack.api.v2 import network_pb2 as api_network_pb2
from metalstack.api.v2 import network_connecpy as api_network_connecpy
from metalstack.api.v2 import partition_pb2 as api_partition_pb2
from metalstack.api.v2 import partition_connecpy as api_partition_connecpy
from metalstack.api.v2 import project_pb2 as api_project_pb2
from metalstack.api.v2 import project_connecpy as api_project_connecpy
from metalstack.api.v2 import size_pb2 as api_size_pb2
from metalstack.api.v2 import size_connecpy as api_size_connecpy
from metalstack.api.v2 import tenant_pb2 as api_tenant_pb2
from metalstack.api.v2 import tenant_connecpy as api_tenant_connecpy
from metalstack.api.v2 import token_pb2 as api_token_pb2
from metalstack.api.v2 import token_connecpy as api_token_connecpy
from metalstack.api.v2 import user_pb2 as api_user_pb2
from metalstack.api.v2 import user_connecpy as api_user_connecpy
from metalstack.api.v2 import version_pb2 as api_version_pb2
from metalstack.api.v2 import version_connecpy as api_version_connecpy

from metalstack.infra.v2 import bmc_pb2 as infra_bmc_pb2
from metalstack.infra.v2 import bmc_connecpy as infra_bmc_connecpy



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



