# -*- coding: utf-8 -*-
# Generated by https://github.com/i2y/connecpy/protoc-gen-connecpy.  DO NOT EDIT!
# source: metalstack/admin/v2/tenant.proto

from typing import Optional, Protocol, Union

import httpx

from connecpy.async_client import AsyncConnecpyClient
from connecpy.base import Endpoint
from connecpy.server import ConnecpyServer
from connecpy.client import ConnecpyClient
from connecpy.context import ClientContext, ServiceContext
import metalstack.admin.v2.tenant_pb2 as metalstack_dot_admin_dot_v2_dot_tenant__pb2


class TenantService(Protocol):
    async def Create(self, req: metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateResponse: ...
    async def List(self, req: metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListResponse: ...


class TenantServiceServer(ConnecpyServer):
    def __init__(self, *, service: TenantService, server_path_prefix=""):
        super().__init__()
        self._prefix = f"{server_path_prefix}/metalstack.admin.v2.TenantService"
        self._endpoints = {
            "Create": Endpoint[metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateRequest, metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateResponse](
                service_name="TenantService",
                name="Create",
                function=getattr(service, "Create"),
                input=metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateRequest,
                output=metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateResponse,
                allowed_methods=("POST",),
            ),
            "List": Endpoint[metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListRequest, metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListResponse](
                service_name="TenantService",
                name="List",
                function=getattr(service, "List"),
                input=metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListRequest,
                output=metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListResponse,
                allowed_methods=("POST",),
            ),
        }

    def serviceName(self):
        return "metalstack.admin.v2.TenantService"


class TenantServiceSync(Protocol):
    def Create(self, req: metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateResponse: ...
    def List(self, req: metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListResponse: ...


class TenantServiceServerSync(ConnecpyServer):
    def __init__(self, *, service: TenantServiceSync, server_path_prefix=""):
        super().__init__()
        self._prefix = f"{server_path_prefix}/metalstack.admin.v2.TenantService"
        self._endpoints = {
            "Create": Endpoint[metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateRequest, metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateResponse](
                service_name="TenantService",
                name="Create",
                function=getattr(service, "Create"),
                input=metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateRequest,
                output=metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateResponse,
                allowed_methods=("POST",),
            ),
            "List": Endpoint[metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListRequest, metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListResponse](
                service_name="TenantService",
                name="List",
                function=getattr(service, "List"),
                input=metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListRequest,
                output=metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListResponse,
                allowed_methods=("POST",),
            ),
        }

    def serviceName(self):
        return "metalstack.admin.v2.TenantService"


class TenantServiceClient(ConnecpyClient):
    def Create(
        self,
        request: metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateResponse:
        method = "POST"
        return self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.TenantService/Create",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateResponse,
            method=method,
            **kwargs,
        )

    def List(
        self,
        request: metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListResponse:
        method = "POST"
        return self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.TenantService/List",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListResponse,
            method=method,
            **kwargs,
        )


class AsyncTenantServiceClient(AsyncConnecpyClient):
    async def Create(
        self,
        request: metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        session: Union[httpx.AsyncClient, None] = None,
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateResponse:
        method = "POST"
        return await self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.TenantService/Create",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceCreateResponse,
            method=method,
            session=session,
            **kwargs,
        )

    async def List(
        self,
        request: metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        session: Union[httpx.AsyncClient, None] = None,
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListResponse:
        method = "POST"
        return await self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.TenantService/List",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_tenant__pb2.TenantServiceListResponse,
            method=method,
            session=session,
            **kwargs,
        )
