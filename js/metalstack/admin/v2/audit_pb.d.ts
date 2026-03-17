import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { AuditPhase, AuditQuery, AuditTrace } from "../../api/v2/audit_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/audit.proto.
 */
export declare const file_metalstack_admin_v2_audit: GenFile;
/**
 * AuditServiceListRequest is the request payload to list audits
 *
 * @generated from message metalstack.admin.v2.AuditServiceListRequest
 */
export type AuditServiceListRequest = Message<"metalstack.admin.v2.AuditServiceListRequest"> & {
    /**
     * Query for audit traces
     *
     * @generated from field: metalstack.api.v2.AuditQuery query = 1;
     */
    query?: AuditQuery;
};
/**
 * Describes the message metalstack.admin.v2.AuditServiceListRequest.
 * Use `create(AuditServiceListRequestSchema)` to create a new message.
 */
export declare const AuditServiceListRequestSchema: GenMessage<AuditServiceListRequest>;
/**
 * AuditServiceListResponse is the response payload of a audit list request
 *
 * @generated from message metalstack.admin.v2.AuditServiceListResponse
 */
export type AuditServiceListResponse = Message<"metalstack.admin.v2.AuditServiceListResponse"> & {
    /**
     * Traces is a list of audit traces
     *
     * @generated from field: repeated metalstack.api.v2.AuditTrace traces = 1;
     */
    traces: AuditTrace[];
};
/**
 * Describes the message metalstack.admin.v2.AuditServiceListResponse.
 * Use `create(AuditServiceListResponseSchema)` to create a new message.
 */
export declare const AuditServiceListResponseSchema: GenMessage<AuditServiceListResponse>;
/**
 * AuditServiceGetRequest is the request payload of a audit get request
 *
 * @generated from message metalstack.admin.v2.AuditServiceGetRequest
 */
export type AuditServiceGetRequest = Message<"metalstack.admin.v2.AuditServiceGetRequest"> & {
    /**
     * Uuid of the audit trace
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Phase specifies the audit phase. Defaults to request
     *
     * @generated from field: optional metalstack.api.v2.AuditPhase phase = 2;
     */
    phase?: AuditPhase;
};
/**
 * Describes the message metalstack.admin.v2.AuditServiceGetRequest.
 * Use `create(AuditServiceGetRequestSchema)` to create a new message.
 */
export declare const AuditServiceGetRequestSchema: GenMessage<AuditServiceGetRequest>;
/**
 * AuditServiceGetResponse is the response payload of a audit get request
 *
 * @generated from message metalstack.admin.v2.AuditServiceGetResponse
 */
export type AuditServiceGetResponse = Message<"metalstack.admin.v2.AuditServiceGetResponse"> & {
    /**
     * Trace is the audit trace
     *
     * @generated from field: metalstack.api.v2.AuditTrace trace = 1;
     */
    trace?: AuditTrace;
};
/**
 * Describes the message metalstack.admin.v2.AuditServiceGetResponse.
 * Use `create(AuditServiceGetResponseSchema)` to create a new message.
 */
export declare const AuditServiceGetResponseSchema: GenMessage<AuditServiceGetResponse>;
/**
 * AuditService serves audit related functions
 *
 * @generated from service metalstack.admin.v2.AuditService
 */
export declare const AuditService: GenService<{
    /**
     * Get an audit trace
     *
     * @generated from rpc metalstack.admin.v2.AuditService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof AuditServiceGetRequestSchema;
        output: typeof AuditServiceGetResponseSchema;
    };
    /**
     * List all audit traces
     *
     * @generated from rpc metalstack.admin.v2.AuditService.List
     */
    list: {
        methodKind: "unary";
        input: typeof AuditServiceListRequestSchema;
        output: typeof AuditServiceListResponseSchema;
    };
}>;
