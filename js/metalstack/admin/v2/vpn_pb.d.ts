import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Duration, Timestamp } from "@bufbuild/protobuf/wkt";
import type { VPNNode } from "../../api/v2/vpn_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/vpn.proto.
 */
export declare const file_metalstack_admin_v2_vpn: GenFile;
/**
 * VPNServiceAuthKeyRequest is the request payload for a vpn authkey request.
 *
 * @generated from message metalstack.admin.v2.VPNServiceAuthKeyRequest
 */
export type VPNServiceAuthKeyRequest = Message<"metalstack.admin.v2.VPNServiceAuthKeyRequest"> & {
    /**
     * Project for which a vpn authkey should be generated.
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Ephemeral defines if the authkey should be ephemeral.
     *
     * @generated from field: bool ephemeral = 2;
     */
    ephemeral: boolean;
    /**
     * Expires defines the duration after which the authkey expires.
     *
     * @generated from field: google.protobuf.Duration expires = 3;
     */
    expires?: Duration;
    /**
     * Reason must be provided why access to the vpn is requested
     * reason is only forwarded to a audit sink
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
 * VPNServiceAuthKeyResponse is the request payload for a authkey response
 *
 * @generated from message metalstack.admin.v2.VPNServiceAuthKeyResponse
 */
export type VPNServiceAuthKeyResponse = Message<"metalstack.admin.v2.VPNServiceAuthKeyResponse"> & {
    /**
     * Address is the address of the vpn control plane.
     *
     * @generated from field: string address = 1;
     */
    address: string;
    /**
     * AuthKey is the key to connect to the vpn at the given address.
     * This key can only be seen once.
     *
     * @generated from field: string auth_key = 2;
     */
    authKey: string;
    /**
     * Ephemeral defines if the authkey should be ephemeral.
     *
     * @generated from field: bool ephemeral = 3;
     */
    ephemeral: boolean;
    /**
     * ExpiresAt this key cannot be used after this timestamp.
     *
     * @generated from field: google.protobuf.Timestamp expires_at = 4;
     */
    expiresAt?: Timestamp;
    /**
     * CreatedAt this key was created at this timestamp.
     *
     * @generated from field: google.protobuf.Timestamp created_at = 5;
     */
    createdAt?: Timestamp;
};
/**
 * Describes the message metalstack.admin.v2.VPNServiceAuthKeyResponse.
 * Use `create(VPNServiceAuthKeyResponseSchema)` to create a new message.
 */
export declare const VPNServiceAuthKeyResponseSchema: GenMessage<VPNServiceAuthKeyResponse>;
/**
 * VPNServiceListNodesRequest is the request payload for a vpn list nodes request
 *
 * @generated from message metalstack.admin.v2.VPNServiceListNodesRequest
 */
export type VPNServiceListNodesRequest = Message<"metalstack.admin.v2.VPNServiceListNodesRequest"> & {
    /**
     * Project if given only nodes of this user are returned
     *
     * @generated from field: optional string project = 1;
     */
    project?: string;
};
/**
 * Describes the message metalstack.admin.v2.VPNServiceListNodesRequest.
 * Use `create(VPNServiceListNodesRequestSchema)` to create a new message.
 */
export declare const VPNServiceListNodesRequestSchema: GenMessage<VPNServiceListNodesRequest>;
/**
 * VPNServiceListNodesResponse is the response payload for a vpn list nodes request
 *
 * @generated from message metalstack.admin.v2.VPNServiceListNodesResponse
 */
export type VPNServiceListNodesResponse = Message<"metalstack.admin.v2.VPNServiceListNodesResponse"> & {
    /**
     * Nodes connected to the vpn
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
 * VPNService serves vpn related functions
 *
 * @generated from service metalstack.admin.v2.VPNService
 */
export declare const VPNService: GenService<{
    /**
     * AuthKey generates a authkey for a project to join a machine to the project vpn
     *
     * @generated from rpc metalstack.admin.v2.VPNService.AuthKey
     */
    authKey: {
        methodKind: "unary";
        input: typeof VPNServiceAuthKeyRequestSchema;
        output: typeof VPNServiceAuthKeyResponseSchema;
    };
    /**
     * ListNodes returns a list of machines actually connected to the vpn
     *
     * @generated from rpc metalstack.admin.v2.VPNService.ListNodes
     */
    listNodes: {
        methodKind: "unary";
        input: typeof VPNServiceListNodesRequestSchema;
        output: typeof VPNServiceListNodesResponseSchema;
    };
}>;
