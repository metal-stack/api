# -*- coding: utf-8 -*-
# Generated by https://github.com/i2y/connecpy/protoc-gen-connecpy.  DO NOT EDIT!
# source: metalstack/api/v2/version.proto

from typing import Optional, Protocol, Union

import httpx

from connecpy.async_client import AsyncConnecpyClient
from connecpy.base import Endpoint
from connecpy.server import ConnecpyServer
from connecpy.client import ConnecpyClient
from connecpy.context import ClientContext, ServiceContext
import metalstack.api.v2.version_pb2 as metalstack_dot_api_dot_v2_dot_version__pb2


class VersionService(Protocol):
    async def Get(self, req: metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetRequest, ctx: ServiceContext) -> metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetResponse: ...


class VersionServiceServer(ConnecpyServer):
    def __init__(self, *, service: VersionService, server_path_prefix=""):
        super().__init__()
        self._prefix = f"{server_path_prefix}/metalstack.api.v2.VersionService"
        self._endpoints = {
            "Get": Endpoint[metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetRequest, metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetResponse](
                service_name="VersionService",
                name="Get",
                function=getattr(service, "Get"),
                input=metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetRequest,
                output=metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetResponse,
                allowed_methods=("POST",),
            ),
        }

    def serviceName(self):
        return "metalstack.api.v2.VersionService"


class VersionServiceSync(Protocol):
    def Get(self, req: metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetRequest, ctx: ServiceContext) -> metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetResponse: ...


class VersionServiceServerSync(ConnecpyServer):
    def __init__(self, *, service: VersionServiceSync, server_path_prefix=""):
        super().__init__()
        self._prefix = f"{server_path_prefix}/metalstack.api.v2.VersionService"
        self._endpoints = {
            "Get": Endpoint[metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetRequest, metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetResponse](
                service_name="VersionService",
                name="Get",
                function=getattr(service, "Get"),
                input=metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetRequest,
                output=metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetResponse,
                allowed_methods=("POST",),
            ),
        }

    def serviceName(self):
        return "metalstack.api.v2.VersionService"


class VersionServiceClient(ConnecpyClient):
    def Get(
        self,
        request: metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        **kwargs,
    ) -> metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetResponse:
        method = "POST"
        return self._make_request(
            url=f"{server_path_prefix}/metalstack.api.v2.VersionService/Get",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetResponse,
            method=method,
            **kwargs,
        )


class AsyncVersionServiceClient(AsyncConnecpyClient):
    async def Get(
        self,
        request: metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        session: Union[httpx.AsyncClient, None] = None,
        **kwargs,
    ) -> metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetResponse:
        method = "POST"
        return await self._make_request(
            url=f"{server_path_prefix}/metalstack.api.v2.VersionService/Get",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_api_dot_v2_dot_version__pb2.VersionServiceGetResponse,
            method=method,
            session=session,
            **kwargs,
        )
