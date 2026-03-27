import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/vpn.proto.
 */
export declare const file_metalstack_api_v2_vpn: GenFile;
/**
 * VPNNode is a machine connected to the vpn
 *
 * @generated from message metalstack.api.v2.VPNNode
 */
export type VPNNode = Message<"metalstack.api.v2.VPNNode"> & {
    /**
     * Id of this node
     *
     * @generated from field: uint64 id = 1;
     */
    id: bigint;
    /**
     * Name of this node
     *
     * @generated from field: string name = 2;
     */
    name: string;
    /**
     * Project of this node, maps to a project
     *
     * @generated from field: string project = 3;
     */
    project: string;
    /**
     * IPAddresses of this node in the vpn
     *
     * @generated from field: repeated string ip_addresses = 4;
     */
    ipAddresses: string[];
    /**
     * LastSeen timestamp when this node reached out the the control plane
     *
     * @generated from field: google.protobuf.Timestamp last_seen = 5;
     */
    lastSeen?: Timestamp;
    /**
     * Online indicates if this node is online
     *
     * @generated from field: bool online = 6;
     */
    online: boolean;
};
/**
 * Describes the message metalstack.api.v2.VPNNode.
 * Use `create(VPNNodeSchema)` to create a new message.
 */
export declare const VPNNodeSchema: GenMessage<VPNNode>;
