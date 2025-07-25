# -*- coding: utf-8 -*-
# Generated by https://github.com/i2y/connecpy/protoc-gen-connecpy.  DO NOT EDIT!
# source: metalstack/admin/v2/partition.proto

from typing import Optional, Protocol, Union

import httpx

from connecpy.async_client import AsyncConnecpyClient
from connecpy.base import Endpoint
from connecpy.server import ConnecpyServer
from connecpy.client import ConnecpyClient
from connecpy.context import ClientContext, ServiceContext
import metalstack.admin.v2.partition_pb2 as metalstack_dot_admin_dot_v2_dot_partition__pb2


class PartitionService(Protocol):
    async def Create(self, req: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateResponse: ...
    async def Update(self, req: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateResponse: ...
    async def Delete(self, req: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteResponse: ...
    async def Capacity(self, req: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityResponse: ...


class PartitionServiceServer(ConnecpyServer):
    def __init__(self, *, service: PartitionService, server_path_prefix=""):
        super().__init__()
        self._prefix = f"{server_path_prefix}/metalstack.admin.v2.PartitionService"
        self._endpoints = {
            "Create": Endpoint[metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateRequest, metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateResponse](
                service_name="PartitionService",
                name="Create",
                function=getattr(service, "Create"),
                input=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateRequest,
                output=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateResponse,
                allowed_methods=("POST",),
            ),
            "Update": Endpoint[metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateRequest, metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateResponse](
                service_name="PartitionService",
                name="Update",
                function=getattr(service, "Update"),
                input=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateRequest,
                output=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateResponse,
                allowed_methods=("POST",),
            ),
            "Delete": Endpoint[metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteRequest, metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteResponse](
                service_name="PartitionService",
                name="Delete",
                function=getattr(service, "Delete"),
                input=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteRequest,
                output=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteResponse,
                allowed_methods=("POST",),
            ),
            "Capacity": Endpoint[metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityRequest, metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityResponse](
                service_name="PartitionService",
                name="Capacity",
                function=getattr(service, "Capacity"),
                input=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityRequest,
                output=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityResponse,
                allowed_methods=("POST",),
            ),
        }

    def serviceName(self):
        return "metalstack.admin.v2.PartitionService"


class PartitionServiceSync(Protocol):
    def Create(self, req: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateResponse: ...
    def Update(self, req: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateResponse: ...
    def Delete(self, req: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteResponse: ...
    def Capacity(self, req: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityRequest, ctx: ServiceContext) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityResponse: ...


class PartitionServiceServerSync(ConnecpyServer):
    def __init__(self, *, service: PartitionServiceSync, server_path_prefix=""):
        super().__init__()
        self._prefix = f"{server_path_prefix}/metalstack.admin.v2.PartitionService"
        self._endpoints = {
            "Create": Endpoint[metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateRequest, metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateResponse](
                service_name="PartitionService",
                name="Create",
                function=getattr(service, "Create"),
                input=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateRequest,
                output=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateResponse,
                allowed_methods=("POST",),
            ),
            "Update": Endpoint[metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateRequest, metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateResponse](
                service_name="PartitionService",
                name="Update",
                function=getattr(service, "Update"),
                input=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateRequest,
                output=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateResponse,
                allowed_methods=("POST",),
            ),
            "Delete": Endpoint[metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteRequest, metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteResponse](
                service_name="PartitionService",
                name="Delete",
                function=getattr(service, "Delete"),
                input=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteRequest,
                output=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteResponse,
                allowed_methods=("POST",),
            ),
            "Capacity": Endpoint[metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityRequest, metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityResponse](
                service_name="PartitionService",
                name="Capacity",
                function=getattr(service, "Capacity"),
                input=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityRequest,
                output=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityResponse,
                allowed_methods=("POST",),
            ),
        }

    def serviceName(self):
        return "metalstack.admin.v2.PartitionService"


class PartitionServiceClient(ConnecpyClient):
    def Create(
        self,
        request: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateResponse:
        method = "POST"
        return self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.PartitionService/Create",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateResponse,
            method=method,
            **kwargs,
        )

    def Update(
        self,
        request: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateResponse:
        method = "POST"
        return self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.PartitionService/Update",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateResponse,
            method=method,
            **kwargs,
        )

    def Delete(
        self,
        request: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteResponse:
        method = "POST"
        return self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.PartitionService/Delete",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteResponse,
            method=method,
            **kwargs,
        )

    def Capacity(
        self,
        request: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityResponse:
        method = "POST"
        return self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.PartitionService/Capacity",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityResponse,
            method=method,
            **kwargs,
        )


class AsyncPartitionServiceClient(AsyncConnecpyClient):
    async def Create(
        self,
        request: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        session: Union[httpx.AsyncClient, None] = None,
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateResponse:
        method = "POST"
        return await self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.PartitionService/Create",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCreateResponse,
            method=method,
            session=session,
            **kwargs,
        )

    async def Update(
        self,
        request: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        session: Union[httpx.AsyncClient, None] = None,
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateResponse:
        method = "POST"
        return await self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.PartitionService/Update",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceUpdateResponse,
            method=method,
            session=session,
            **kwargs,
        )

    async def Delete(
        self,
        request: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        session: Union[httpx.AsyncClient, None] = None,
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteResponse:
        method = "POST"
        return await self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.PartitionService/Delete",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceDeleteResponse,
            method=method,
            session=session,
            **kwargs,
        )

    async def Capacity(
        self,
        request: metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityRequest,
        *,
        ctx: Optional[ClientContext] = None,
        server_path_prefix: str = "",
        session: Union[httpx.AsyncClient, None] = None,
        **kwargs,
    ) -> metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityResponse:
        method = "POST"
        return await self._make_request(
            url=f"{server_path_prefix}/metalstack.admin.v2.PartitionService/Capacity",
            ctx=ctx,
            request=request,
            response_class=metalstack_dot_admin_dot_v2_dot_partition__pb2.PartitionServiceCapacityResponse,
            method=method,
            session=session,
            **kwargs,
        )
