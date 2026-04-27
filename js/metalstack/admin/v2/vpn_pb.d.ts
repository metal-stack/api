import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Duration, Timestamp } from "@bufbuild/protobuf/wkt";
import type { VPNNode } from "../../api/v2/vpn_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/vpn.proto.
 */
export declare const file_metalstack_admin_v2_vpn: GenFile;
/**
 * VPNServiceAuthKeyRequest is the request payload for generating a VPN authentication key.
 *
 * @generated from message metalstack.admin.v2.VPNServiceAuthKeyRequest
 */
export type VPNServiceAuthKeyRequest = Message<"metalstack.admin.v2.VPNServiceAuthKeyRequest"> & {
    /**
     * Project for which a VPN authentication key should be generated
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Ephemeral defines if the authentication key should be ephemeral
     *
     * @generated from field: bool ephemeral = 2;
     */
    ephemeral: boolean;
    /**
     * Expires defines the duration after which the authentication key expires
     *
     * @generated from field: google.protobuf.Duration expires = 3;
     */
    expires?: Duration | undefined;
    /**
     * Reason must be provided why access to the VPN is requested
     * Reason is only forwarded to an audit sink
     *
     * @generated from field: string reason = 4;
     */
    reason: string;
};
/**
 * Describes the message metalstack.admin.v2.VPNServiceAuthKeyRequest.
 * Use `create(VPNServiceAuthKeyRequestSchema)` to create a new message.
 */
export declare const VPNServiceAuthKeyRequestSchema: GenMessage<VPNServiceAuthKeyRequest>;
/**
 * VPNServiceAuthKeyResponse is the response payload for generating a VPN authentication key.
 *
 * @generated from message metalstack.admin.v2.VPNServiceAuthKeyResponse
 */
export type VPNServiceAuthKeyResponse = Message<"metalstack.admin.v2.VPNServiceAuthKeyResponse"> & {
    /**
     * Address is the address of the VPN control plane
     *
     * @generated from field: string address = 1;
     */
    address: string;
    /**
     * AuthKey is the key to connect to the VPN at the given address
     * This key can only be seen once
     *
     * @generated from field: string auth_key = 2;
     */
    authKey: string;
    /**
     * Ephemeral defines if the authentication key should be ephemeral
     *
     * @generated from field: bool ephemeral = 3;
     */
    ephemeral: boolean;
    /**
     * ExpiresAt this key cannot be used after this timestamp
     *
     * @generated from field: google.protobuf.Timestamp expires_at = 4;
     */
    expiresAt?: Timestamp | undefined;
    /**
     * CreatedAt this key was created at this timestamp
     *
     * @generated from field: google.protobuf.Timestamp created_at = 5;
     */
    createdAt?: Timestamp | undefined;
};
/**
 * Describes the message metalstack.admin.v2.VPNServiceAuthKeyResponse.
 * Use `create(VPNServiceAuthKeyResponseSchema)` to create a new message.
 */
export declare const VPNServiceAuthKeyResponseSchema: GenMessage<VPNServiceAuthKeyResponse>;
/**
 * VPNServiceListNodesRequest is the request payload for listing VPN nodes.
 *
 * @generated from message metalstack.admin.v2.VPNServiceListNodesRequest
 */
export type VPNServiceListNodesRequest = Message<"metalstack.admin.v2.VPNServiceListNodesRequest"> & {
    /**
     * Project filters nodes by this project
     *
     * @generated from field: optional string project = 1;
     */
    project?: string | undefined;
};
/**
 * Describes the message metalstack.admin.v2.VPNServiceListNodesRequest.
 * Use `create(VPNServiceListNodesRequestSchema)` to create a new message.
 */
export declare const VPNServiceListNodesRequestSchema: GenMessage<VPNServiceListNodesRequest>;
/**
 * VPNServiceListNodesResponse is the response payload for listing VPN nodes.
 *
 * @generated from message metalstack.admin.v2.VPNServiceListNodesResponse
 */
export type VPNServiceListNodesResponse = Message<"metalstack.admin.v2.VPNServiceListNodesResponse"> & {
    /**
     * Nodes connected to the VPN
     *
     * @generated from field: repeated metalstack.api.v2.VPNNode nodes = 1;
     */
    nodes: VPNNode[];
};
/**
 * Describes the message metalstack.admin.v2.VPNServiceListNodesResponse.
 * Use `create(VPNServiceListNodesResponseSchema)` to create a new message.
 */
export declare const VPNServiceListNodesResponseSchema: GenMessage<VPNServiceListNodesResponse>;
/**
 * VPNService provides VPN management operations.
 *
 * @generated from service metalstack.admin.v2.VPNService
 */
export declare const VPNService: GenService<{
    /**
     * GenerateAuthKey generates an authentication key for a project to join a machine to the project VPN.
     *
     * @generated from rpc metalstack.admin.v2.VPNService.AuthKey
     */
    authKey: {
        methodKind: "unary";
        input: typeof VPNServiceAuthKeyRequestSchema;
        output: typeof VPNServiceAuthKeyResponseSchema;
    };
    /**
     * Returns the list of machines connected to the VPN.
     *
     * @generated from rpc metalstack.admin.v2.VPNService.ListNodes
     */
    listNodes: {
        methodKind: "unary";
        input: typeof VPNServiceListNodesRequestSchema;
        output: typeof VPNServiceListNodesResponseSchema;
    };
}>;
