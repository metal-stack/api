import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { IP, IPQuery } from "../../api/v2/ip_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/ip.proto.
 */
export declare const file_metalstack_admin_v2_ip: GenFile;
/**
 * IPServiceListRequest
 *
 * @generated from message metalstack.admin.v2.IPServiceListRequest
 */
export type IPServiceListRequest = Message<"metalstack.admin.v2.IPServiceListRequest"> & {
    /**
     * Query to search for one or more ips
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
 * IPServiceListResponse
 *
 * @generated from message metalstack.admin.v2.IPServiceListResponse
 */
export type IPServiceListResponse = Message<"metalstack.admin.v2.IPServiceListResponse"> & {
    /**
     * IPs are the list of ips
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
 * IPService serves ip address related functions
 *
 * @generated from service metalstack.admin.v2.IPService
 */
export declare const IPService: GenService<{
    /**
     * List all ips
     *
     * @generated from rpc metalstack.admin.v2.IPService.List
     */
    list: {
        methodKind: "unary";
        input: typeof IPServiceListRequestSchema;
        output: typeof IPServiceListResponseSchema;
    };
}>;
