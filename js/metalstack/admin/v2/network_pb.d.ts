import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Labels, UpdateLabels, UpdateMeta } from "../../api/v2/common_pb";
import type { ChildPrefixLength, ExternalNetworkMemberQuery, ExternalNetworkMembers, NATType, Network, NetworkAddressFamily, NetworkQuery, NetworkType } from "../../api/v2/network_pb";
import type { Switch } from "../../api/v2/switch_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/network.proto.
 */
export declare const file_metalstack_admin_v2_network: GenFile;
/**
 * NetworkServiceGetRequest is the request payload for getting a network.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceGetRequest
 */
export type NetworkServiceGetRequest = Message<"metalstack.admin.v2.NetworkServiceGetRequest"> & {
    /**
     * ID of the network to get.
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
 * NetworkServiceGetResponse is the response payload for getting a network.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceGetResponse
 */
export type NetworkServiceGetResponse = Message<"metalstack.admin.v2.NetworkServiceGetResponse"> & {
    /**
     * Network contains the requested network.
     *
     * @generated from field: metalstack.api.v2.Network network = 1;
     */
    network?: Network | undefined;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceGetResponse.
 * Use `create(NetworkServiceGetResponseSchema)` to create a new message.
 */
export declare const NetworkServiceGetResponseSchema: GenMessage<NetworkServiceGetResponse>;
/**
 * NetworkServiceCreateRequest is the request payload for creating a network.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceCreateRequest
 */
export type NetworkServiceCreateRequest = Message<"metalstack.admin.v2.NetworkServiceCreateRequest"> & {
    /**
     * Id of this network.
     *
     * @generated from field: optional string id = 1;
     */
    id?: string | undefined;
    /**
     * Name of this network.
     *
     * @generated from field: optional string name = 2;
     */
    name?: string | undefined;
    /**
     * Description of this network.
     *
     * @generated from field: optional string description = 3;
     */
    description?: string | undefined;
    /**
     * Partition where this network will be created.
     *
     * @generated from field: optional string partition = 4;
     */
    partition?: string | undefined;
    /**
     * Project where this network belongs to.
     *
     * @generated from field: optional string project = 5;
     */
    project?: string | undefined;
    /**
     * Type of the network to create.
     *
     * @generated from field: metalstack.api.v2.NetworkType type = 6;
     */
    type: NetworkType;
    /**
     * Labels on this network.
     *
     * @generated from field: optional metalstack.api.v2.Labels labels = 7;
     */
    labels?: Labels | undefined;
    /**
     * Prefixes in this network.
     *
     * @generated from field: repeated string prefixes = 8;
     */
    prefixes: string[];
    /**
     * Destination Prefixes in this network.
     *
     * @generated from field: repeated string destination_prefixes = 9;
     */
    destinationPrefixes: string[];
    /**
     * Default Child Prefix length defines the bitlength of a child network created per addressfamily, if not specified during the allocate request.
     *
     * @generated from field: metalstack.api.v2.ChildPrefixLength default_child_prefix_length = 10;
     */
    defaultChildPrefixLength?: ChildPrefixLength | undefined;
    /**
     * Min Child Prefix length asserts that during child network creation the requested bit length is greater or equal the min child prefix length.
     *
     * @generated from field: metalstack.api.v2.ChildPrefixLength min_child_prefix_length = 11;
     */
    minChildPrefixLength?: ChildPrefixLength | undefined;
    /**
     * NATType of this network.
     *
     * @generated from field: optional metalstack.api.v2.NATType nat_type = 12;
     */
    natType?: NATType | undefined;
    /**
     * VRF of this network has this VNI.
     *
     * @generated from field: optional uint32 vrf = 13;
     */
    vrf?: number | undefined;
    /**
     * Parent Network points to the id of the parent network if any.
     *
     * @generated from field: optional string parent_network = 14;
     */
    parentNetwork?: string | undefined;
    /**
     * AdditionalAnnouncableCidrs will be added to the allow list on the switch which prefixes might be announced.
     *
     * @generated from field: repeated string additional_announcable_cidrs = 15;
     */
    additionalAnnouncableCidrs: string[];
    /**
     * Length per addressfamily.
     *
     * @generated from field: optional metalstack.api.v2.ChildPrefixLength length = 16;
     */
    length?: ChildPrefixLength | undefined;
    /**
     * AddressFamily to create, defaults to the same as the parent.
     *
     * @generated from field: optional metalstack.api.v2.NetworkAddressFamily address_family = 17;
     */
    addressFamily?: NetworkAddressFamily | undefined;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceCreateRequest.
 * Use `create(NetworkServiceCreateRequestSchema)` to create a new message.
 */
export declare const NetworkServiceCreateRequestSchema: GenMessage<NetworkServiceCreateRequest>;
/**
 * NetworkServiceUpdateRequest is the request payload for updating a network.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceUpdateRequest
 */
export type NetworkServiceUpdateRequest = Message<"metalstack.admin.v2.NetworkServiceUpdateRequest"> & {
    /**
     * Id of this network.
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request.
     *
     * @generated from field: metalstack.api.v2.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta | undefined;
    /**
     * Name of this network.
     *
     * @generated from field: optional string name = 3;
     */
    name?: string | undefined;
    /**
     * Description of this network.
     *
     * @generated from field: optional string description = 4;
     */
    description?: string | undefined;
    /**
     * Labels to update on this network.
     *
     * @generated from field: optional metalstack.api.v2.UpdateLabels labels = 5;
     */
    labels?: UpdateLabels | undefined;
    /**
     * Prefixes in this network.
     *
     * @generated from field: repeated string prefixes = 6;
     */
    prefixes: string[];
    /**
     * Destination Prefixes in this network.
     *
     * @generated from field: repeated string destination_prefixes = 7;
     */
    destinationPrefixes: string[];
    /**
     * Default Child Prefix length defines the bit length of a child network created per addressfamily, if not specified during the allocate request.
     *
     * @generated from field: optional metalstack.api.v2.ChildPrefixLength default_child_prefix_length = 8;
     */
    defaultChildPrefixLength?: ChildPrefixLength | undefined;
    /**
     * Min Child Prefix length asserts that during child network creation the requested bit length is greater or equal the min child prefix length.
     *
     * @generated from field: optional metalstack.api.v2.ChildPrefixLength min_child_prefix_length = 9;
     */
    minChildPrefixLength?: ChildPrefixLength | undefined;
    /**
     * NATType of this network.
     *
     * @generated from field: optional metalstack.api.v2.NATType nat_type = 10;
     */
    natType?: NATType | undefined;
    /**
     * AdditionalAnnouncableCidrs will be added to the allow list on the switch which prefixes might be announced.
     *
     * @generated from field: repeated string additional_announcable_cidrs = 11;
     */
    additionalAnnouncableCidrs: string[];
    /**
     * Force update, prevents accidental removal of additional_announcable_cidrs which will destroy your dataplane.
     *
     * @generated from field: bool force = 12;
     */
    force: boolean;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceUpdateRequest.
 * Use `create(NetworkServiceUpdateRequestSchema)` to create a new message.
 */
export declare const NetworkServiceUpdateRequestSchema: GenMessage<NetworkServiceUpdateRequest>;
/**
 * NetworkServiceDeleteRequest is the request payload for deleting a network.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceDeleteRequest
 */
export type NetworkServiceDeleteRequest = Message<"metalstack.admin.v2.NetworkServiceDeleteRequest"> & {
    /**
     * ID of the network to get.
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
 * NetworkServiceListRequest is the request payload for listing networks.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceListRequest
 */
export type NetworkServiceListRequest = Message<"metalstack.admin.v2.NetworkServiceListRequest"> & {
    /**
     * Query specifies which networks to return.
     *
     * @generated from field: metalstack.api.v2.NetworkQuery query = 1;
     */
    query?: NetworkQuery | undefined;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceListRequest.
 * Use `create(NetworkServiceListRequestSchema)` to create a new message.
 */
export declare const NetworkServiceListRequestSchema: GenMessage<NetworkServiceListRequest>;
/**
 * NetworkServiceListExternalMembersRequest is the request payload for listing external members of a network.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceListExternalMembersRequest
 */
export type NetworkServiceListExternalMembersRequest = Message<"metalstack.admin.v2.NetworkServiceListExternalMembersRequest"> & {
    /**
     * Network to list the members for.
     *
     * @generated from field: string network = 1;
     */
    network: string;
    /**
     * Query specifies additional filters for the list request.
     *
     * @generated from field: metalstack.api.v2.ExternalNetworkMemberQuery query = 2;
     */
    query?: ExternalNetworkMemberQuery | undefined;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceListExternalMembersRequest.
 * Use `create(NetworkServiceListExternalMembersRequestSchema)` to create a new message.
 */
export declare const NetworkServiceListExternalMembersRequestSchema: GenMessage<NetworkServiceListExternalMembersRequest>;
/**
 * NetworkServiceAddExternalMemberRequest is the request payload for adding an external member to a network.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceAddExternalMemberRequest
 */
export type NetworkServiceAddExternalMemberRequest = Message<"metalstack.admin.v2.NetworkServiceAddExternalMemberRequest"> & {
    /**
     * Network to add the member to.
     *
     * @generated from field: string network = 1;
     */
    network: string;
    /**
     * SwitchPorts contains the ports of a switch that should be added to the network.
     * Make sure to add both switches of a rack and the same ports of both in the same request.
     *
     * @generated from field: repeated metalstack.api.v2.ExternalNetworkMembers switch_ports = 2;
     */
    switchPorts: ExternalNetworkMembers[];
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceAddExternalMemberRequest.
 * Use `create(NetworkServiceAddExternalMemberRequestSchema)` to create a new message.
 */
export declare const NetworkServiceAddExternalMemberRequestSchema: GenMessage<NetworkServiceAddExternalMemberRequest>;
/**
 * NetworkServiceRemoveExternalMemberRequest is the request payload for removing an external member from a network.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceRemoveExternalMemberRequest
 */
export type NetworkServiceRemoveExternalMemberRequest = Message<"metalstack.admin.v2.NetworkServiceRemoveExternalMemberRequest"> & {
    /**
     * Network to remove the member from.
     *
     * @generated from field: string network = 1;
     */
    network: string;
    /**
     * SwitchPorts contains the ports of a switch that should be removed from the network.
     *
     * @generated from field: repeated metalstack.api.v2.ExternalNetworkMembers switch_ports = 2;
     */
    switchPorts: ExternalNetworkMembers[];
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceRemoveExternalMemberRequest.
 * Use `create(NetworkServiceRemoveExternalMemberRequestSchema)` to create a new message.
 */
export declare const NetworkServiceRemoveExternalMemberRequestSchema: GenMessage<NetworkServiceRemoveExternalMemberRequest>;
/**
 * NetworkServiceCreateResponse is the response payload for creating a network.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceCreateResponse
 */
export type NetworkServiceCreateResponse = Message<"metalstack.admin.v2.NetworkServiceCreateResponse"> & {
    /**
     * Network contains the created network.
     *
     * @generated from field: metalstack.api.v2.Network network = 1;
     */
    network?: Network | undefined;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceCreateResponse.
 * Use `create(NetworkServiceCreateResponseSchema)` to create a new message.
 */
export declare const NetworkServiceCreateResponseSchema: GenMessage<NetworkServiceCreateResponse>;
/**
 * NetworkServiceUpdateResponse is the response payload for updating a network.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceUpdateResponse
 */
export type NetworkServiceUpdateResponse = Message<"metalstack.admin.v2.NetworkServiceUpdateResponse"> & {
    /**
     * Network contains the updated network.
     *
     * @generated from field: metalstack.api.v2.Network network = 1;
     */
    network?: Network | undefined;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceUpdateResponse.
 * Use `create(NetworkServiceUpdateResponseSchema)` to create a new message.
 */
export declare const NetworkServiceUpdateResponseSchema: GenMessage<NetworkServiceUpdateResponse>;
/**
 * NetworkServiceDeleteResponse is the response payload for deleting a network.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceDeleteResponse
 */
export type NetworkServiceDeleteResponse = Message<"metalstack.admin.v2.NetworkServiceDeleteResponse"> & {
    /**
     * Network contains the deleted network.
     *
     * @generated from field: metalstack.api.v2.Network network = 1;
     */
    network?: Network | undefined;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceDeleteResponse.
 * Use `create(NetworkServiceDeleteResponseSchema)` to create a new message.
 */
export declare const NetworkServiceDeleteResponseSchema: GenMessage<NetworkServiceDeleteResponse>;
/**
 * NetworkServiceListResponse is the response payload for listing networks.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceListResponse
 */
export type NetworkServiceListResponse = Message<"metalstack.admin.v2.NetworkServiceListResponse"> & {
    /**
     * Networks contains the list of networks.
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
 * NetworkServiceListExternalMembersResponse is the response payload for listing external members of a network.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceListExternalMembersResponse
 */
export type NetworkServiceListExternalMembersResponse = Message<"metalstack.admin.v2.NetworkServiceListExternalMembersResponse"> & {
    /**
     * Members are the queried external members of the network.
     *
     * @generated from field: repeated metalstack.api.v2.ExternalNetworkMembers members = 1;
     */
    members: ExternalNetworkMembers[];
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceListExternalMembersResponse.
 * Use `create(NetworkServiceListExternalMembersResponseSchema)` to create a new message.
 */
export declare const NetworkServiceListExternalMembersResponseSchema: GenMessage<NetworkServiceListExternalMembersResponse>;
/**
 * NetworkServiceAddExternalMemberResponse is the response payload for adding an external member to a network.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceAddExternalMemberResponse
 */
export type NetworkServiceAddExternalMemberResponse = Message<"metalstack.admin.v2.NetworkServiceAddExternalMemberResponse"> & {
    /**
     * Switch contains the updated switch.
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch | undefined;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceAddExternalMemberResponse.
 * Use `create(NetworkServiceAddExternalMemberResponseSchema)` to create a new message.
 */
export declare const NetworkServiceAddExternalMemberResponseSchema: GenMessage<NetworkServiceAddExternalMemberResponse>;
/**
 * NetworkServiceRemoveExternalMemberResponse is the response payload for removing an external member from a network.
 *
 * @generated from message metalstack.admin.v2.NetworkServiceRemoveExternalMemberResponse
 */
export type NetworkServiceRemoveExternalMemberResponse = Message<"metalstack.admin.v2.NetworkServiceRemoveExternalMemberResponse"> & {
    /**
     * Switch contains the updated switch.
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch | undefined;
};
/**
 * Describes the message metalstack.admin.v2.NetworkServiceRemoveExternalMemberResponse.
 * Use `create(NetworkServiceRemoveExternalMemberResponseSchema)` to create a new message.
 */
export declare const NetworkServiceRemoveExternalMemberResponseSchema: GenMessage<NetworkServiceRemoveExternalMemberResponse>;
/**
 * NetworkService provides network management operations.
 *
 * @generated from service metalstack.admin.v2.NetworkService
 */
export declare const NetworkService: GenService<{
    /**
     * Returns the network with the specified ID.
     *
     * @generated from rpc metalstack.admin.v2.NetworkService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof NetworkServiceGetRequestSchema;
        output: typeof NetworkServiceGetResponseSchema;
    };
    /**
     * Creates a new network.
     *
     * @generated from rpc metalstack.admin.v2.NetworkService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof NetworkServiceCreateRequestSchema;
        output: typeof NetworkServiceCreateResponseSchema;
    };
    /**
     * Updates a network.
     *
     * @generated from rpc metalstack.admin.v2.NetworkService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof NetworkServiceUpdateRequestSchema;
        output: typeof NetworkServiceUpdateResponseSchema;
    };
    /**
     * Deletes a network.
     *
     * @generated from rpc metalstack.admin.v2.NetworkService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof NetworkServiceDeleteRequestSchema;
        output: typeof NetworkServiceDeleteResponseSchema;
    };
    /**
     * Returns the list of all networks.
     *
     * @generated from rpc metalstack.admin.v2.NetworkService.List
     */
    list: {
        methodKind: "unary";
        input: typeof NetworkServiceListRequestSchema;
        output: typeof NetworkServiceListResponseSchema;
    };
    /**
     * Lists external members of a network.
     *
     * @generated from rpc metalstack.admin.v2.NetworkService.ListExternalMembers
     */
    listExternalMembers: {
        methodKind: "unary";
        input: typeof NetworkServiceListExternalMembersRequestSchema;
        output: typeof NetworkServiceListExternalMembersResponseSchema;
    };
    /**
     * Adds an external member to a network.
     *
     * @generated from rpc metalstack.admin.v2.NetworkService.AddExternalMember
     */
    addExternalMember: {
        methodKind: "unary";
        input: typeof NetworkServiceAddExternalMemberRequestSchema;
        output: typeof NetworkServiceAddExternalMemberResponseSchema;
    };
    /**
     * Removes an external member from a network.
     *
     * @generated from rpc metalstack.admin.v2.NetworkService.RemoveExternalMember
     */
    removeExternalMember: {
        methodKind: "unary";
        input: typeof NetworkServiceRemoveExternalMemberRequestSchema;
        output: typeof NetworkServiceRemoveExternalMemberResponseSchema;
    };
}>;
