import { Interceptor } from "@connectrpc/connect";
import type { Client as ConnectClient } from "@connectrpc/connect";
import { AuditService as Adminv2AuditService } from "./metalstack/admin/v2/audit_pb";
import { ComponentService as Adminv2ComponentService } from "./metalstack/admin/v2/component_pb";
import { FilesystemService as Adminv2FilesystemService } from "./metalstack/admin/v2/filesystem_pb";
import { ImageService as Adminv2ImageService } from "./metalstack/admin/v2/image_pb";
import { IPService as Adminv2IPService } from "./metalstack/admin/v2/ip_pb";
import { MachineService as Adminv2MachineService } from "./metalstack/admin/v2/machine_pb";
import { NetworkService as Adminv2NetworkService } from "./metalstack/admin/v2/network_pb";
import { PartitionService as Adminv2PartitionService } from "./metalstack/admin/v2/partition_pb";
import { ProjectService as Adminv2ProjectService } from "./metalstack/admin/v2/project_pb";
import { SizeService as Adminv2SizeService } from "./metalstack/admin/v2/size_pb";
import { SizeImageConstraintService as Adminv2SizeImageConstraintService } from "./metalstack/admin/v2/size_imageconstraint_pb";
import { SizeReservationService as Adminv2SizeReservationService } from "./metalstack/admin/v2/size_reservation_pb";
import { SwitchService as Adminv2SwitchService } from "./metalstack/admin/v2/switch_pb";
import { TaskService as Adminv2TaskService } from "./metalstack/admin/v2/task_pb";
import { TenantService as Adminv2TenantService } from "./metalstack/admin/v2/tenant_pb";
import { TokenService as Adminv2TokenService } from "./metalstack/admin/v2/token_pb";
import { VPNService as Adminv2VPNService } from "./metalstack/admin/v2/vpn_pb";
import { AuditService as Apiv2AuditService } from "./metalstack/api/v2/audit_pb";
import { FilesystemService as Apiv2FilesystemService } from "./metalstack/api/v2/filesystem_pb";
import { HealthService as Apiv2HealthService } from "./metalstack/api/v2/health_pb";
import { ImageService as Apiv2ImageService } from "./metalstack/api/v2/image_pb";
import { IPService as Apiv2IPService } from "./metalstack/api/v2/ip_pb";
import { MachineService as Apiv2MachineService } from "./metalstack/api/v2/machine_pb";
import { MethodService as Apiv2MethodService } from "./metalstack/api/v2/method_pb";
import { NetworkService as Apiv2NetworkService } from "./metalstack/api/v2/network_pb";
import { PartitionService as Apiv2PartitionService } from "./metalstack/api/v2/partition_pb";
import { ProjectService as Apiv2ProjectService } from "./metalstack/api/v2/project_pb";
import { SizeService as Apiv2SizeService } from "./metalstack/api/v2/size_pb";
import { SizeImageConstraintService as Apiv2SizeImageConstraintService } from "./metalstack/api/v2/size_imageconstraint_pb";
import { SizeReservationService as Apiv2SizeReservationService } from "./metalstack/api/v2/size_reservation_pb";
import { TenantService as Apiv2TenantService } from "./metalstack/api/v2/tenant_pb";
import { TokenService as Apiv2TokenService } from "./metalstack/api/v2/token_pb";
import { UserService as Apiv2UserService } from "./metalstack/api/v2/user_pb";
import { VersionService as Apiv2VersionService } from "./metalstack/api/v2/version_pb";
import { BMCService as Infrav2BMCService } from "./metalstack/infra/v2/bmc_pb";
import { BootService as Infrav2BootService } from "./metalstack/infra/v2/boot_pb";
import { ComponentService as Infrav2ComponentService } from "./metalstack/infra/v2/component_pb";
import { EventService as Infrav2EventService } from "./metalstack/infra/v2/event_pb";
import { SwitchService as Infrav2SwitchService } from "./metalstack/infra/v2/switch_pb";
export interface ClientConfig {
    baseUrl: string;
    token?: string;
    interceptors?: Interceptor[];
}
export interface Client {
    adminv2(): Adminv2;
    apiv2(): Apiv2;
    infrav2(): Infrav2;
}
export interface Adminv2 {
    audit(): ConnectClient<typeof Adminv2AuditService>;
    component(): ConnectClient<typeof Adminv2ComponentService>;
    filesystem(): ConnectClient<typeof Adminv2FilesystemService>;
    image(): ConnectClient<typeof Adminv2ImageService>;
    ip(): ConnectClient<typeof Adminv2IPService>;
    machine(): ConnectClient<typeof Adminv2MachineService>;
    network(): ConnectClient<typeof Adminv2NetworkService>;
    partition(): ConnectClient<typeof Adminv2PartitionService>;
    project(): ConnectClient<typeof Adminv2ProjectService>;
    size(): ConnectClient<typeof Adminv2SizeService>;
    sizeImageConstraint(): ConnectClient<typeof Adminv2SizeImageConstraintService>;
    sizeReservation(): ConnectClient<typeof Adminv2SizeReservationService>;
    switch(): ConnectClient<typeof Adminv2SwitchService>;
    task(): ConnectClient<typeof Adminv2TaskService>;
    tenant(): ConnectClient<typeof Adminv2TenantService>;
    token(): ConnectClient<typeof Adminv2TokenService>;
    vpn(): ConnectClient<typeof Adminv2VPNService>;
}
export interface Apiv2 {
    audit(): ConnectClient<typeof Apiv2AuditService>;
    filesystem(): ConnectClient<typeof Apiv2FilesystemService>;
    health(): ConnectClient<typeof Apiv2HealthService>;
    image(): ConnectClient<typeof Apiv2ImageService>;
    ip(): ConnectClient<typeof Apiv2IPService>;
    machine(): ConnectClient<typeof Apiv2MachineService>;
    method(): ConnectClient<typeof Apiv2MethodService>;
    network(): ConnectClient<typeof Apiv2NetworkService>;
    partition(): ConnectClient<typeof Apiv2PartitionService>;
    project(): ConnectClient<typeof Apiv2ProjectService>;
    size(): ConnectClient<typeof Apiv2SizeService>;
    sizeImageConstraint(): ConnectClient<typeof Apiv2SizeImageConstraintService>;
    sizeReservation(): ConnectClient<typeof Apiv2SizeReservationService>;
    tenant(): ConnectClient<typeof Apiv2TenantService>;
    token(): ConnectClient<typeof Apiv2TokenService>;
    user(): ConnectClient<typeof Apiv2UserService>;
    version(): ConnectClient<typeof Apiv2VersionService>;
}
export interface Infrav2 {
    bmc(): ConnectClient<typeof Infrav2BMCService>;
    boot(): ConnectClient<typeof Infrav2BootService>;
    component(): ConnectClient<typeof Infrav2ComponentService>;
    event(): ConnectClient<typeof Infrav2EventService>;
    switch(): ConnectClient<typeof Infrav2SwitchService>;
}
export declare function newClient(config: ClientConfig): Client;
