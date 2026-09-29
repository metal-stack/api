import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv2";
import type { TenantMember as TenantMember$1 } from "../../api/v2/tenant_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/cli/v2/tenant.proto.
 */
export declare const file_metalstack_cli_v2_tenant: GenFile;
/**
 * TenantMember defines a user that participates in a tenant.
 *
 * @generated from message metalstack.cli.v2.TenantMember
 */
export type TenantMember = Message<"metalstack.cli.v2.TenantMember"> & {
    /**
     * tenant is the login tenant to which the member belongs.
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * TenantMember is the tenant member.
     *
     * @generated from field: metalstack.api.v2.TenantMember tenant_member = 2;
     */
    tenantMember?: TenantMember$1 | undefined;
};
/**
 * Describes the message metalstack.cli.v2.TenantMember.
 * Use `create(TenantMemberSchema)` to create a new message.
 */
export declare const TenantMemberSchema: GenMessage<TenantMember>;
