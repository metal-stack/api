import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Labels, Meta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/size_reservation.proto.
 */
export declare const file_metalstack_api_v2_size_reservation: GenFile;
/**
 * SizeReservationServiceGetRequest is the request payload for a size get request
 *
 * @generated from message metalstack.api.v2.SizeReservationServiceGetRequest
 */
export type SizeReservationServiceGetRequest = Message<"metalstack.api.v2.SizeReservationServiceGetRequest"> & {
    /**
     * ID of the size reservation to get
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * Project of the size reservation
     *
     * @generated from field: string project = 2;
     */
    project: string;
};
/**
 * Describes the message metalstack.api.v2.SizeReservationServiceGetRequest.
 * Use `create(SizeReservationServiceGetRequestSchema)` to create a new message.
 */
export declare const SizeReservationServiceGetRequestSchema: GenMessage<SizeReservationServiceGetRequest>;
/**
 * SizeReservationServiceListRequest is the request payload for a size list request
 *
 * @generated from message metalstack.api.v2.SizeReservationServiceListRequest
 */
export type SizeReservationServiceListRequest = Message<"metalstack.api.v2.SizeReservationServiceListRequest"> & {
    /**
     * Project of the size reservation
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Query for size reservations
     *
     * @generated from field: metalstack.api.v2.SizeReservationQuery query = 2;
     */
    query?: SizeReservationQuery;
};
/**
 * Describes the message metalstack.api.v2.SizeReservationServiceListRequest.
 * Use `create(SizeReservationServiceListRequestSchema)` to create a new message.
 */
export declare const SizeReservationServiceListRequestSchema: GenMessage<SizeReservationServiceListRequest>;
/**
 * SizeReservationServiceGetResponse is the response payload for a size reservation get request
 *
 * @generated from message metalstack.api.v2.SizeReservationServiceGetResponse
 */
export type SizeReservationServiceGetResponse = Message<"metalstack.api.v2.SizeReservationServiceGetResponse"> & {
    /**
     *  Size reservation
     *
     * @generated from field: metalstack.api.v2.SizeReservation size_reservation = 1;
     */
    sizeReservation?: SizeReservation;
};
/**
 * Describes the message metalstack.api.v2.SizeReservationServiceGetResponse.
 * Use `create(SizeReservationServiceGetResponseSchema)` to create a new message.
 */
export declare const SizeReservationServiceGetResponseSchema: GenMessage<SizeReservationServiceGetResponse>;
/**
 * SizeReservationServiceListResponse is the response payload for a size reservation list request
 *
 * @generated from message metalstack.api.v2.SizeReservationServiceListResponse
 */
export type SizeReservationServiceListResponse = Message<"metalstack.api.v2.SizeReservationServiceListResponse"> & {
    /**
     * Size reservations
     *
     * @generated from field: repeated metalstack.api.v2.SizeReservation size_reservations = 1;
     */
    sizeReservations: SizeReservation[];
};
/**
 * Describes the message metalstack.api.v2.SizeReservationServiceListResponse.
 * Use `create(SizeReservationServiceListResponseSchema)` to create a new message.
 */
export declare const SizeReservationServiceListResponseSchema: GenMessage<SizeReservationServiceListResponse>;
/**
 * SizeReservation
 *
 * @generated from message metalstack.api.v2.SizeReservation
 */
export type SizeReservation = Message<"metalstack.api.v2.SizeReservation"> & {
    /**
     * Id of this size reservation, is generated on creation
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * Meta for this size reservation
     *
     * @generated from field: metalstack.api.v2.Meta meta = 2;
     */
    meta?: Meta;
    /**
     * Name of this size reservation
     *
     * @generated from field: string name = 3;
     */
    name: string;
    /**
     * Description of this size reservation
     *
     * @generated from field: string description = 4;
     */
    description: string;
    /**
     * Project of the size reservation
     *
     * @generated from field: string project = 5;
     */
    project: string;
    /**
     * Size id of this size reservation
     *
     * @generated from field: string size = 6;
     */
    size: string;
    /**
     * Partition ids of this size reservation
     *
     * @generated from field: repeated string partitions = 7;
     */
    partitions: string[];
    /**
     * Amount of reservations of this size reservation
     *
     * @generated from field: int32 amount = 8;
     */
    amount: number;
};
/**
 * Describes the message metalstack.api.v2.SizeReservation.
 * Use `create(SizeReservationSchema)` to create a new message.
 */
export declare const SizeReservationSchema: GenMessage<SizeReservation>;
/**
 * SizeReservationQuery is used to search size reservations
 *
 * @generated from message metalstack.api.v2.SizeReservationQuery
 */
export type SizeReservationQuery = Message<"metalstack.api.v2.SizeReservationQuery"> & {
    /**
     * ID of the size reservation to get
     *
     * @generated from field: optional string id = 1;
     */
    id?: string;
    /**
     * Name of this size reservation
     *
     * @generated from field: optional string name = 2;
     */
    name?: string;
    /**
     * Description of this size reservation
     *
     * @generated from field: optional string description = 3;
     */
    description?: string;
    /**
     * Size id of this size reservation
     *
     * @generated from field: optional string size = 4;
     */
    size?: string;
    /**
     * Project of the size reservation
     *
     * @generated from field: optional string project = 5;
     */
    project?: string;
    /**
     * Partition of the size reservation
     *
     * @generated from field: optional string partition = 6;
     */
    partition?: string;
    /**
     * Labels lists only size reservations containing the given labels
     *
     * @generated from field: optional metalstack.api.v2.Labels labels = 7;
     */
    labels?: Labels;
};
/**
 * Describes the message metalstack.api.v2.SizeReservationQuery.
 * Use `create(SizeReservationQuerySchema)` to create a new message.
 */
export declare const SizeReservationQuerySchema: GenMessage<SizeReservationQuery>;
/**
 * SizeReservationService serves size reservation related functions
 *
 * @generated from service metalstack.api.v2.SizeReservationService
 */
export declare const SizeReservationService: GenService<{
    /**
     * Get a size reservation
     *
     * @generated from rpc metalstack.api.v2.SizeReservationService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof SizeReservationServiceGetRequestSchema;
        output: typeof SizeReservationServiceGetResponseSchema;
    };
    /**
     * List size reservations
     *
     * @generated from rpc metalstack.api.v2.SizeReservationService.List
     */
    list: {
        methodKind: "unary";
        input: typeof SizeReservationServiceListRequestSchema;
        output: typeof SizeReservationServiceListResponseSchema;
    };
}>;
