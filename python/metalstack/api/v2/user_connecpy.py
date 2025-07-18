# -*- coding: utf-8 -*-
# Generated by https://github.com/i2y/connecpy/protoc-gen-connecpy.  DO NOT EDIT!
# source: metalstack/api/v2/user.proto

from typing import Optional, Protocol, Union

import httpx

from connecpy.async_client import AsyncConnecpyClient
from connecpy.base import Endpoint
from connecpy.server import ConnecpyServer
from connecpy.client import ConnecpyClient
from connecpy.context import ClientContext, ServiceContext
import metalstack.api.v2.user_pb2 as metalstack_dot_api_dot_v2_dot_user__pb2


class UserService(Protocol):
    async def Get(self, req: metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetRequest, ctx: ServiceContext) -> metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetResponse: ...


class UserServiceServer(ConnecpyServer):
    def __init__(self, *, service: UserService, server_path_prefix=""):
        super().__init__()
        self._prefix = f"{server_path_prefix}/metalstack.api.v2.UserService"
        self._endpoints = {
            "Get": Endpoint[metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetRequest, metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetResponse](
                service_name="UserService",
                name="Get",
                function=getattr(service, "Get"),
                input=metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetRequest,
                output=metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetResponse,
                allowed_methods=("POST",),
            ),
        }

    def serviceName(self):
        return "metalstack.api.v2.UserService"


class UserServiceSync(Protocol):
    def Get(self, req: metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetRequest, ctx: ServiceContext) -> metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetResponse: ...


class UserServiceServerSync(ConnecpyServer):
    def __init__(self, *, service: UserServiceSync, server_path_prefix=""):
        super().__init__()
        self._prefix = f"{server_path_prefix}/metalstack.api.v2.UserService"
        self._endpoints = {
            "Get": Endpoint[metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetRequest, metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetResponse](
                service_name="UserService",
                name="Get",
                function=getattr(service, "Get"),
                input=metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetRequest,
                output=metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetResponse,
                allowed_methods=("POST",),
            ),
        }

    def serviceName(self):
        return "metalstack.api.v2.UserService"


class UserServiceClient(ConnecpyClient):
    def Get(
        self,
        request: metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        **kwargs,
    ) -> metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetResponse:
        method = "POST"
        return self._make_request(
            url=f"{server_path_prefix}/metalstack.api.v2.UserService/Get",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetResponse,
            method=method,
            **kwargs,
        )


class AsyncUserServiceClient(AsyncConnecpyClient):
    async def Get(
        self,
        request: metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        session: Union[httpx.AsyncClient, None] = None,
        **kwargs,
    ) -> metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetResponse:
        method = "POST"
        return await self._make_request(
            url=f"{server_path_prefix}/metalstack.api.v2.UserService/Get",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_api_dot_v2_dot_user__pb2.UserServiceGetResponse,
            method=method,
            session=session,
            **kwargs,
        )
