import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { IP, IPQuery } from "../../api/v2/ip_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/ip.proto.
 */
export declare const file_metalstack_admin_v2_ip: GenFile;
/**
 * IPServiceListRequest is the request payload for listing IP addresses.
 *
 * @generated from message metalstack.admin.v2.IPServiceListRequest
 */
export type IPServiceListRequest = Message<"metalstack.admin.v2.IPServiceListRequest"> & {
    /**
     * Query to search for IP addresses
     *
     * @generated from field: metalstack.api.v2.IPQuery query = 1;
     */
    query?: IPQuery;
};
/**
 * Describes the message metalstack.admin.v2.IPServiceListRequest.
 * Use `create(IPServiceListRequestSchema)` to create a new message.
 */
export declare const IPServiceListRequestSchema: GenMessage<IPServiceListRequest>;
/**
 * IPServiceListResponse is the response payload for listing IP addresses.
 *
 * @generated from message metalstack.admin.v2.IPServiceListResponse
 */
export type IPServiceListResponse = Message<"metalstack.admin.v2.IPServiceListResponse"> & {
    /**
     * IPs contains the list of IP addresses
     *
     * @generated from field: repeated metalstack.api.v2.IP ips = 1;
     */
    ips: IP[];
};
/**
 * Describes the message metalstack.admin.v2.IPServiceListResponse.
 * Use `create(IPServiceListResponseSchema)` to create a new message.
 */
export declare const IPServiceListResponseSchema: GenMessage<IPServiceListResponse>;
/**
 * IPService provides IP address management operations.
 *
 * @generated from service metalstack.admin.v2.IPService
 */
export declare const IPService: GenService<{
    /**
     * Returns the list of all IP addresses.
     *
     * @generated from rpc metalstack.admin.v2.IPService.List
     */
    list: {
        methodKind: "unary";
        input: typeof IPServiceListRequestSchema;
        output: typeof IPServiceListResponseSchema;
    };
}>;
