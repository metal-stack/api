import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/audit.proto.
 */
export declare const file_metalstack_api_v2_audit: GenFile;
/**
 * AuditTrace is an audit trace.
 *
 * @generated from message metalstack.api.v2.AuditTrace
 */
export type AuditTrace = Message<"metalstack.api.v2.AuditTrace"> & {
    /**
     * Uuid of the audit trace
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Timestamp is the timestamp when the request arrived at the API
     *
     * @generated from field: google.protobuf.Timestamp timestamp = 2;
     */
    timestamp?: Timestamp | undefined;
    /**
     * User is the login user who called the API method
     *
     * @generated from field: string user = 3;
     */
    user: string;
    /**
     * Tenant is the tenant targeted by the API call
     *
     * @generated from field: string tenant = 4;
     */
    tenant: string;
    /**
     * Project is the project targeted by the API call
     *
     * @generated from field: optional string project = 5;
     */
    project?: string | undefined;
    /**
     * Method is the API method that was called
     *
     * @generated from field: string method = 6;
     */
    method: string;
    /**
     * Body is the payload of the API call. In the request phase this contains the payload sent by the client, in the response phase it contains the payload returned by the API server
     *
     * @generated from field: optional string body = 7;
     */
    body?: string | undefined;
    /**
     * SourceIP contains the source IP address of the API call
     *
     * @generated from field: string source_ip = 8;
     */
    sourceIp: string;
    /**
     * ResultCode is a status code describing the result of the API call. It is set for traces in the response phase and contains official gRPC status codes
     *
     * @generated from field: optional int32 result_code = 9;
     */
    resultCode?: number | undefined;
    /**
     * Phase represents the phase of the audit trace
     *
     * @generated from field: metalstack.api.v2.AuditPhase phase = 10;
     */
    phase: AuditPhase;
};
/**
 * Describes the message metalstack.api.v2.AuditTrace.
 * Use `create(AuditTraceSchema)` to create a new message.
 */
export declare const AuditTraceSchema: GenMessage<AuditTrace>;
/**
 * AuditQuery is the query for audit traces
 *
 * @generated from message metalstack.api.v2.AuditQuery
 */
export type AuditQuery = Message<"metalstack.api.v2.AuditQuery"> & {
    /**
     * Uuid of the audit trace
     *
     * @generated from field: optional string uuid = 2;
     */
    uuid?: string | undefined;
    /**
     * From describes the start of the time window in which to list audit traces. Defaults to the last eight hours
     *
     * @generated from field: optional google.protobuf.Timestamp from = 3;
     */
    from?: Timestamp | undefined;
    /**
     * To describes the end of the time window in which to list audit traces. Defaults to the time the request was issued
     *
     * @generated from field: optional google.protobuf.Timestamp to = 4;
     */
    to?: Timestamp | undefined;
    /**
     * User is the user who called the api method
     *
     * @generated from field: optional string user = 5;
     */
    user?: string | undefined;
    /**
     * Project is the project targeted by the api call
     *
     * @generated from field: optional string project = 6;
     */
    project?: string | undefined;
    /**
     * Method is the api method that was called
     *
     * @generated from field: optional string method = 7;
     */
    method?: string | undefined;
    /**
     * Source IP contains the ip address of the caller
     *
     * @generated from field: optional string source_ip = 8;
     */
    sourceIp?: string | undefined;
    /**
     * Result Code is a string describing the result of the api call
     *
     * @generated from field: optional int32 result_code = 9;
     */
    resultCode?: number | undefined;
    /**
     * Body is a string providing text-search of the body field
     *
     * @generated from field: optional string body = 10;
     */
    body?: string | undefined;
    /**
     * Limit is a number limiting the length of the response (min: 1, max: 1000, defaults to 200)
     *
     * @generated from field: optional int32 limit = 11;
     */
    limit?: number | undefined;
    /**
     * Phase specifies the audit phase
     *
     * @generated from field: optional metalstack.api.v2.AuditPhase phase = 12;
     */
    phase?: AuditPhase | undefined;
};
/**
 * Describes the message metalstack.api.v2.AuditQuery.
 * Use `create(AuditQuerySchema)` to create a new message.
 */
export declare const AuditQuerySchema: GenMessage<AuditQuery>;
/**
 * AuditServiceListRequest is the request payload for a audit list request
 *
 * @generated from message metalstack.api.v2.AuditServiceListRequest
 */
export type AuditServiceListRequest = Message<"metalstack.api.v2.AuditServiceListRequest"> & {
    /**
     * Login for this tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Query for audit traces
     *
     * @generated from field: metalstack.api.v2.AuditQuery query = 2;
     */
    query?: AuditQuery | undefined;
};
/**
 * Describes the message metalstack.api.v2.AuditServiceListRequest.
 * Use `create(AuditServiceListRequestSchema)` to create a new message.
 */
export declare const AuditServiceListRequestSchema: GenMessage<AuditServiceListRequest>;
/**
 * AuditServiceListResponse is the response payload of a audit list request
 *
 * @generated from message metalstack.api.v2.AuditServiceListResponse
 */
export type AuditServiceListResponse = Message<"metalstack.api.v2.AuditServiceListResponse"> & {
    /**
     * Traces is a list of audit traces
     *
     * @generated from field: repeated metalstack.api.v2.AuditTrace traces = 1;
     */
    traces: AuditTrace[];
};
/**
 * Describes the message metalstack.api.v2.AuditServiceListResponse.
 * Use `create(AuditServiceListResponseSchema)` to create a new message.
 */
export declare const AuditServiceListResponseSchema: GenMessage<AuditServiceListResponse>;
/**
 * AuditServiceGetRequest is the request payload of a audit get request
 *
 * @generated from message metalstack.api.v2.AuditServiceGetRequest
 */
export type AuditServiceGetRequest = Message<"metalstack.api.v2.AuditServiceGetRequest"> & {
    /**
     * Login for this tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Uuid of the audit trace
     *
     * @generated from field: string uuid = 2;
     */
    uuid: string;
    /**
     * Phase specifies the audit phase. Defaults to request
     *
     * @generated from field: optional metalstack.api.v2.AuditPhase phase = 3;
     */
    phase?: AuditPhase | undefined;
};
/**
 * Describes the message metalstack.api.v2.AuditServiceGetRequest.
 * Use `create(AuditServiceGetRequestSchema)` to create a new message.
 */
export declare const AuditServiceGetRequestSchema: GenMessage<AuditServiceGetRequest>;
/**
 * AuditServiceGetResponse is the response payload of a audit get request
 *
 * @generated from message metalstack.api.v2.AuditServiceGetResponse
 */
export type AuditServiceGetResponse = Message<"metalstack.api.v2.AuditServiceGetResponse"> & {
    /**
     * Trace is the audit trace
     *
     * @generated from field: metalstack.api.v2.AuditTrace trace = 1;
     */
    trace?: AuditTrace | undefined;
};
/**
 * Describes the message metalstack.api.v2.AuditServiceGetResponse.
 * Use `create(AuditServiceGetResponseSchema)` to create a new message.
 */
export declare const AuditServiceGetResponseSchema: GenMessage<AuditServiceGetResponse>;
/**
 * AuditPhase specifies the phase of an audit trace.
 *
 * @generated from enum metalstack.api.v2.AuditPhase
 */
export declare enum AuditPhase {
    /**
     * AUDIT_PHASE_UNSPECIFIED is not specified
     *
     * @generated from enum value: AUDIT_PHASE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * AUDIT_PHASE_REQUEST defines an audit trace in the request phase
     *
     * @generated from enum value: AUDIT_PHASE_REQUEST = 1;
     */
    REQUEST = 1,
    /**
     * AUDIT_PHASE_RESPONSE defines an audit trace in the response phase
     *
     * @generated from enum value: AUDIT_PHASE_RESPONSE = 2;
     */
    RESPONSE = 2
}
/**
 * Describes the enum metalstack.api.v2.AuditPhase.
 */
export declare const AuditPhaseSchema: GenEnum<AuditPhase>;
/**
 * AuditService provides audit logging operations.
 *
 * @generated from service metalstack.api.v2.AuditService
 */
export declare const AuditService: GenService<{
    /**
     * Returns the audit trace with the specified UUID.
     *
     * @generated from rpc metalstack.api.v2.AuditService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof AuditServiceGetRequestSchema;
        output: typeof AuditServiceGetResponseSchema;
    };
    /**
     * Returns the list of audit traces.
     *
     * @generated from rpc metalstack.api.v2.AuditService.List
     */
    list: {
        methodKind: "unary";
        input: typeof AuditServiceListRequestSchema;
        output: typeof AuditServiceListResponseSchema;
    };
}>;
