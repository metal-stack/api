from contextlib import contextmanager
from typing import Optional

from metalstack.api.v2 import ip_pb2, ip_connecpy, network_pb2, network_connecpy


from connecpy.context import ClientContext
from connecpy.exceptions import ConnecpyServerException


class ClientWrapper:
    def __init__(self, client, token):
        self._client = client
        self._token = token

    def __getattr__(self, name):
        attr = getattr(self._client, name)
        if callable(attr):
            def wrapper(*args, **kwargs):
                headers = kwargs.get("headers", {})
                headers.setdefault("Authorization", "Bearer " + self._token)
                kwargs["headers"] = headers
                return attr(*args, **kwargs)
            return wrapper
        return attr

class ServiceDriver:
    def __init__(self, baseurl: str, token: str, timeout: int = 10):
        self.baseurl = baseurl
        self.token = token
        self.timeout = timeout

    @contextmanager
    def ip(self):
        with ip_connecpy.IPServiceClient(self.baseurl, timeout=self.timeout) as client:
           yield ClientWrapper(client, self.token)

    @contextmanager
    def network(self):
        with network_connecpy.NetworkServiceClient(self.baseurl, timeout=self.timeout) as client:
            yield ClientWrapper(client, self.token)

    def list_ips(self, project: str):
        try:
            with self.ip() as client:
                response = client.List(
                    ctx=ClientContext(),
                    request=ip_pb2.IPServiceListRequest(project=project),
                )
                return response.ips
        except ConnecpyServerException as e:
            print("Error listing IPs:", e.code, e.message)
            return []


# Example usage
def main():
    driver = ServiceDriver(baseurl="https://example.com", token="your_token", timeout=10)
    ips = driver.list_ips(project="project_id")
    for ip in ips:
        print(ip.ip, ip.name, ip.project, ip.network)
