# -*- coding: utf-8 -*-
# Generated by https://github.com/i2y/connecpy/protoc-gen-connecpy.  DO NOT EDIT!
# source: metalstack/admin/v2/size.proto

from typing import Optional, Protocol, Union

import httpx

from connecpy.async_client import AsyncConnecpyClient
from connecpy.base import Endpoint
from connecpy.server import ConnecpyServer
from connecpy.client import ConnecpyClient
from connecpy.context import ClientContext, ServiceContext
import metalstack.admin.v2.size_pb2 as metalstack_dot_admin_dot_v2_dot_size__pb2


class SizeService(Protocol):
    async def Create(self, req: metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateResponse: ...
    async def Update(self, req: metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateResponse: ...
    async def Delete(self, req: metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteResponse: ...


class SizeServiceServer(ConnecpyServer):
    def __init__(self, *, service: SizeService, server_path_prefix=""):
        super().__init__()
        self._prefix = f"{server_path_prefix}/metalstack.admin.v2.SizeService"
        self._endpoints = {
            "Create": Endpoint[metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateRequest, metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateResponse](
                service_name="SizeService",
                name="Create",
                function=getattr(service, "Create"),
                input=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateRequest,
                output=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateResponse,
                allowed_methods=("POST",),
            ),
            "Update": Endpoint[metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateRequest, metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateResponse](
                service_name="SizeService",
                name="Update",
                function=getattr(service, "Update"),
                input=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateRequest,
                output=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateResponse,
                allowed_methods=("POST",),
            ),
            "Delete": Endpoint[metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteRequest, metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteResponse](
                service_name="SizeService",
                name="Delete",
                function=getattr(service, "Delete"),
                input=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteRequest,
                output=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteResponse,
                allowed_methods=("POST",),
            ),
        }

    def serviceName(self):
        return "metalstack.admin.v2.SizeService"


class SizeServiceSync(Protocol):
    def Create(self, req: metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateResponse: ...
    def Update(self, req: metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateResponse: ...
    def Delete(self, req: metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteResponse: ...


class SizeServiceServerSync(ConnecpyServer):
    def __init__(self, *, service: SizeServiceSync, server_path_prefix=""):
        super().__init__()
        self._prefix = f"{server_path_prefix}/metalstack.admin.v2.SizeService"
        self._endpoints = {
            "Create": Endpoint[metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateRequest, metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateResponse](
                service_name="SizeService",
                name="Create",
                function=getattr(service, "Create"),
                input=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateRequest,
                output=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateResponse,
                allowed_methods=("POST",),
            ),
            "Update": Endpoint[metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateRequest, metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateResponse](
                service_name="SizeService",
                name="Update",
                function=getattr(service, "Update"),
                input=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateRequest,
                output=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateResponse,
                allowed_methods=("POST",),
            ),
            "Delete": Endpoint[metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteRequest, metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteResponse](
                service_name="SizeService",
                name="Delete",
                function=getattr(service, "Delete"),
                input=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteRequest,
                output=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteResponse,
                allowed_methods=("POST",),
            ),
        }

    def serviceName(self):
        return "metalstack.admin.v2.SizeService"


class SizeServiceClient(ConnecpyClient):
    def Create(
        self,
        request: metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateResponse:
        method = "POST"
        return self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.SizeService/Create",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateResponse,
            method=method,
            **kwargs,
        )

    def Update(
        self,
        request: metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateResponse:
        method = "POST"
        return self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.SizeService/Update",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateResponse,
            method=method,
            **kwargs,
        )

    def Delete(
        self,
        request: metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteResponse:
        method = "POST"
        return self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.SizeService/Delete",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteResponse,
            method=method,
            **kwargs,
        )


class AsyncSizeServiceClient(AsyncConnecpyClient):
    async def Create(
        self,
        request: metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        session: Union[httpx.AsyncClient, None] = None,
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateResponse:
        method = "POST"
        return await self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.SizeService/Create",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceCreateResponse,
            method=method,
            session=session,
            **kwargs,
        )

    async def Update(
        self,
        request: metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        session: Union[httpx.AsyncClient, None] = None,
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateResponse:
        method = "POST"
        return await self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.SizeService/Update",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceUpdateResponse,
            method=method,
            session=session,
            **kwargs,
        )

    async def Delete(
        self,
        request: metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        session: Union[httpx.AsyncClient, None] = None,
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteResponse:
        method = "POST"
        return await self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.SizeService/Delete",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_size__pb2.SizeServiceDeleteResponse,
            method=method,
            session=session,
            **kwargs,
        )
