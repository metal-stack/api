import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Labels, UpdateLabels, UpdateMeta } from "../../api/v2/common_pb";
import type { ChildPrefixLength, NATType, Network, NetworkAddressFamily, NetworkQuery, NetworkType } from "../../api/v2/network_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/network.proto.
 */
export declare const file_metalstack_admin_v2_network: GenFile;
/**
 * NetworkServiceGetRequest
 *
 * @generated from message metalstack.admin.v2.NetworkServiceGetRequest
 */
export type NetworkServiceGetRequest = Message<"metalstack.admin.v2.NetworkServiceGetRequest"> & {
    /**
     * ID of the network to get
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceGetRequest.
 * Use `create(NetworkServiceGetRequestSchema)` to create a new message.
 */
export declare const NetworkServiceGetRequestSchema: GenMessage<NetworkServiceGetRequest>;
/**
 * NetworkServiceGetResponse
 *
 * @generated from message metalstack.admin.v2.NetworkServiceGetResponse
 */
export type NetworkServiceGetResponse = Message<"metalstack.admin.v2.NetworkServiceGetResponse"> & {
    /**
     * Network which was requested to get
     *
     * @generated from field: metalstack.api.v2.Network network = 1;
     */
    network?: Network;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceGetResponse.
 * Use `create(NetworkServiceGetResponseSchema)` to create a new message.
 */
export declare const NetworkServiceGetResponseSchema: GenMessage<NetworkServiceGetResponse>;
/**
 * NetworkServiceCreateRequest
 *
 * @generated from message metalstack.admin.v2.NetworkServiceCreateRequest
 */
export type NetworkServiceCreateRequest = Message<"metalstack.admin.v2.NetworkServiceCreateRequest"> & {
    /**
     * Id of this network
     *
     * @generated from field: optional string id = 1;
     */
    id?: string;
    /**
     * Name of this network
     *
     * @generated from field: optional string name = 2;
     */
    name?: string;
    /**
     * Description of this network
     *
     * @generated from field: optional string description = 3;
     */
    description?: string;
    /**
     * Partition where this network will be created
     *
     * @generated from field: optional string partition = 4;
     */
    partition?: string;
    /**
     * Project where this network belongs to
     *
     * @generated from field: optional string project = 5;
     */
    project?: string;
    /**
     * Type of the network to create
     *
     * @generated from field: metalstack.api.v2.NetworkType type = 6;
     */
    type: NetworkType;
    /**
     * Labels on this network
     *
     * @generated from field: optional metalstack.api.v2.Labels labels = 7;
     */
    labels?: Labels;
    /**
     * Prefixes in this network
     *
     * @generated from field: repeated string prefixes = 8;
     */
    prefixes: string[];
    /**
     * Destination Prefixes in this network
     *
     * @generated from field: repeated string destination_prefixes = 9;
     */
    destinationPrefixes: string[];
    /**
     * Default Child Prefix length defines the bitlength of a child network created per addressfamily, if not specified during the allocate request
     *
     * @generated from field: metalstack.api.v2.ChildPrefixLength default_child_prefix_length = 10;
     */
    defaultChildPrefixLength?: ChildPrefixLength;
    /**
     * Min Child Prefix length asserts that during child network creation the requested bit length is greater or equal the min child prefix length
     *
     * @generated from field: metalstack.api.v2.ChildPrefixLength min_child_prefix_length = 11;
     */
    minChildPrefixLength?: ChildPrefixLength;
    /**
     * NATType of this network
     *
     * @generated from field: optional metalstack.api.v2.NATType nat_type = 12;
     */
    natType?: NATType;
    /**
     * VRF of this network has this VNI.
     *
     * @generated from field: optional uint32 vrf = 13;
     */
    vrf?: number;
    /**
     * Parent Network points to the id of the parent network if any
     *
     * @generated from field: optional string parent_network = 14;
     */
    parentNetwork?: string;
    /**
     * AdditionalAnnouncableCidrs will be added to the allow list on the switch which prefixes might be announced
     *
     * @generated from field: repeated string additional_announcable_cidrs = 15;
     */
    additionalAnnouncableCidrs: string[];
    /**
     * Length per addressfamily
     *
     * @generated from field: optional metalstack.api.v2.ChildPrefixLength length = 16;
     */
    length?: ChildPrefixLength;
    /**
     * AddressFamily to create, defaults to the same as the parent
     *
     * @generated from field: optional metalstack.api.v2.NetworkAddressFamily address_family = 17;
     */
    addressFamily?: NetworkAddressFamily;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceCreateRequest.
 * Use `create(NetworkServiceCreateRequestSchema)` to create a new message.
 */
export declare const NetworkServiceCreateRequestSchema: GenMessage<NetworkServiceCreateRequest>;
/**
 * NetworkServiceUpdateRequest is the request payload for a network update request
 *
 * @generated from message metalstack.admin.v2.NetworkServiceUpdateRequest
 */
export type NetworkServiceUpdateRequest = Message<"metalstack.admin.v2.NetworkServiceUpdateRequest"> & {
    /**
     * Id of this network
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request
     *
     * @generated from field: metalstack.api.v2.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta;
    /**
     * Name of this network
     *
     * @generated from field: optional string name = 3;
     */
    name?: string;
    /**
     * Description of this network
     *
     * @generated from field: optional string description = 4;
     */
    description?: string;
    /**
     * Labels to update on this network
     *
     * @generated from field: optional metalstack.api.v2.UpdateLabels labels = 5;
     */
    labels?: UpdateLabels;
    /**
     * Prefixes in this network
     *
     * @generated from field: repeated string prefixes = 6;
     */
    prefixes: string[];
    /**
     * Destination Prefixes in this network
     *
     * @generated from field: repeated string destination_prefixes = 7;
     */
    destinationPrefixes: string[];
    /**
     * Default Child Prefix length defines the bit length of a child network created per addressfamily, of not specified during the allocate request
     *
     * @generated from field: optional metalstack.api.v2.ChildPrefixLength default_child_prefix_length = 10;
     */
    defaultChildPrefixLength?: ChildPrefixLength;
    /**
     * Min Child Prefix length asserts that during child network creation the requested bit length is greater or equal the min child prefix length
     *
     * @generated from field: optional metalstack.api.v2.ChildPrefixLength min_child_prefix_length = 11;
     */
    minChildPrefixLength?: ChildPrefixLength;
    /**
     * NATType of this network
     *
     * @generated from field: optional metalstack.api.v2.NATType nat_type = 13;
     */
    natType?: NATType;
    /**
     * AdditionalAnnouncableCidrs will be added to the allow list on the switch which prefixes might be announced
     *
     * @generated from field: repeated string additional_announcable_cidrs = 16;
     */
    additionalAnnouncableCidrs: string[];
    /**
     * Force update, actually only prevents accidental removal of additional_announcable_cidrs which will destroy your dataplane in fact.
     *
     * @generated from field: bool force = 20;
     */
    force: boolean;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceUpdateRequest.
 * Use `create(NetworkServiceUpdateRequestSchema)` to create a new message.
 */
export declare const NetworkServiceUpdateRequestSchema: GenMessage<NetworkServiceUpdateRequest>;
/**
 * NetworkServiceDeleteRequest is the request payload for a network delete request
 *
 * @generated from message metalstack.admin.v2.NetworkServiceDeleteRequest
 */
export type NetworkServiceDeleteRequest = Message<"metalstack.admin.v2.NetworkServiceDeleteRequest"> & {
    /**
     * ID of the network to get
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceDeleteRequest.
 * Use `create(NetworkServiceDeleteRequestSchema)` to create a new message.
 */
export declare const NetworkServiceDeleteRequestSchema: GenMessage<NetworkServiceDeleteRequest>;
/**
 * NetworkServiceListRequest
 *
 * @generated from message metalstack.admin.v2.NetworkServiceListRequest
 */
export type NetworkServiceListRequest = Message<"metalstack.admin.v2.NetworkServiceListRequest"> & {
    /**
     * Query which specifies which networks to return
     *
     * @generated from field: metalstack.api.v2.NetworkQuery query = 2;
     */
    query?: NetworkQuery;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceListRequest.
 * Use `create(NetworkServiceListRequestSchema)` to create a new message.
 */
export declare const NetworkServiceListRequestSchema: GenMessage<NetworkServiceListRequest>;
/**
 * NetworkServiceCreateResponse is the response payload for a network create request
 *
 * @generated from message metalstack.admin.v2.NetworkServiceCreateResponse
 */
export type NetworkServiceCreateResponse = Message<"metalstack.admin.v2.NetworkServiceCreateResponse"> & {
    /**
     * Network the network
     *
     * @generated from field: metalstack.api.v2.Network network = 1;
     */
    network?: Network;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceCreateResponse.
 * Use `create(NetworkServiceCreateResponseSchema)` to create a new message.
 */
export declare const NetworkServiceCreateResponseSchema: GenMessage<NetworkServiceCreateResponse>;
/**
 * NetworkServiceUpdateResponse is the response payload for a network update request
 *
 * @generated from message metalstack.admin.v2.NetworkServiceUpdateResponse
 */
export type NetworkServiceUpdateResponse = Message<"metalstack.admin.v2.NetworkServiceUpdateResponse"> & {
    /**
     * Network the network
     *
     * @generated from field: metalstack.api.v2.Network network = 1;
     */
    network?: Network;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceUpdateResponse.
 * Use `create(NetworkServiceUpdateResponseSchema)` to create a new message.
 */
export declare const NetworkServiceUpdateResponseSchema: GenMessage<NetworkServiceUpdateResponse>;
/**
 * NetworkServiceCapacityResponse is the response payload for a network delete request
 *
 * @generated from message metalstack.admin.v2.NetworkServiceDeleteResponse
 */
export type NetworkServiceDeleteResponse = Message<"metalstack.admin.v2.NetworkServiceDeleteResponse"> & {
    /**
     * Network the network
     *
     * @generated from field: metalstack.api.v2.Network network = 1;
     */
    network?: Network;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceDeleteResponse.
 * Use `create(NetworkServiceDeleteResponseSchema)` to create a new message.
 */
export declare const NetworkServiceDeleteResponseSchema: GenMessage<NetworkServiceDeleteResponse>;
/**
 * NetworkServiceListResponse
 *
 * @generated from message metalstack.admin.v2.NetworkServiceListResponse
 */
export type NetworkServiceListResponse = Message<"metalstack.admin.v2.NetworkServiceListResponse"> & {
    /**
     * Networks are the requested networks
     *
     * @generated from field: repeated metalstack.api.v2.Network networks = 1;
     */
    networks: Network[];
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceListResponse.
 * Use `create(NetworkServiceListResponseSchema)` to create a new message.
 */
export declare const NetworkServiceListResponseSchema: GenMessage<NetworkServiceListResponse>;
/**
 * NetworkService serves network address related functions
 *
 * @generated from service metalstack.admin.v2.NetworkService
 */
export declare const NetworkService: GenService<{
    /**
     * Get a network
     *
     * @generated from rpc metalstack.admin.v2.NetworkService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof NetworkServiceGetRequestSchema;
        output: typeof NetworkServiceGetResponseSchema;
    };
    /**
     * Create a network
     *
     * @generated from rpc metalstack.admin.v2.NetworkService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof NetworkServiceCreateRequestSchema;
        output: typeof NetworkServiceCreateResponseSchema;
    };
    /**
     * Update a network
     *
     * @generated from rpc metalstack.admin.v2.NetworkService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof NetworkServiceUpdateRequestSchema;
        output: typeof NetworkServiceUpdateResponseSchema;
    };
    /**
     * Delete a network
     *
     * @generated from rpc metalstack.admin.v2.NetworkService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof NetworkServiceDeleteRequestSchema;
        output: typeof NetworkServiceDeleteResponseSchema;
    };
    /**
     * List all networks
     *
     * @generated from rpc metalstack.admin.v2.NetworkService.List
     */
    list: {
        methodKind: "unary";
        input: typeof NetworkServiceListRequestSchema;
        output: typeof NetworkServiceListResponseSchema;
    };
}>;
