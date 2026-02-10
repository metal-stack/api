import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { UpdateLabels, UpdateMeta } from "../../api/v2/common_pb";
import type { SizeReservation, SizeReservationQuery } from "../../api/v2/size_reservation_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/size_reservation.proto.
 */
export declare const file_metalstack_admin_v2_size_reservation: GenFile;
/**
 * SizeReservationServiceCreateRequest is the request payload for a size reservation create request
 *
 * @generated from message metalstack.admin.v2.SizeReservationServiceCreateRequest
 */
export type SizeReservationServiceCreateRequest = Message<"metalstack.admin.v2.SizeReservationServiceCreateRequest"> & {
    /**
     * SizeReservation is the size reservation to create
     *
     * @generated from field: metalstack.api.v2.SizeReservation size_reservation = 1;
     */
    sizeReservation?: SizeReservation;
};
/**
 * Describes the message metalstack.admin.v2.SizeReservationServiceCreateRequest.
 * Use `create(SizeReservationServiceCreateRequestSchema)` to create a new message.
 */
export declare const SizeReservationServiceCreateRequestSchema: GenMessage<SizeReservationServiceCreateRequest>;
/**
 * SizeReservationServiceCreateResponse is the response payload for a size reservation create request
 *
 * @generated from message metalstack.admin.v2.SizeReservationServiceCreateResponse
 */
export type SizeReservationServiceCreateResponse = Message<"metalstack.admin.v2.SizeReservationServiceCreateResponse"> & {
    /**
     * SizeReservation the size reservation
     *
     * @generated from field: metalstack.api.v2.SizeReservation size_reservation = 1;
     */
    sizeReservation?: SizeReservation;
};
/**
 * Describes the message metalstack.admin.v2.SizeReservationServiceCreateResponse.
 * Use `create(SizeReservationServiceCreateResponseSchema)` to create a new message.
 */
export declare const SizeReservationServiceCreateResponseSchema: GenMessage<SizeReservationServiceCreateResponse>;
/**
 * SizeReservationServiceUpdateRequest is the request payload for a size reservation update request
 *
 * @generated from message metalstack.admin.v2.SizeReservationServiceUpdateRequest
 */
export type SizeReservationServiceUpdateRequest = Message<"metalstack.admin.v2.SizeReservationServiceUpdateRequest"> & {
    /**
     * Id of this size reservation
     *
     * @generated from field: string id = 1;
     */
    id: string;
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
     * UpdateMeta contains the timestamp and strategy to be used in this update request
     *
     * @generated from field: metalstack.api.v2.UpdateMeta update_meta = 4;
     */
    updateMeta?: UpdateMeta;
    /**
     * Partition ids of this size reservation
     *
     * @generated from field: repeated string partitions = 5;
     */
    partitions: string[];
    /**
     * Amount of reservations of this size reservation
     *
     * @generated from field: optional int32 amount = 6;
     */
    amount?: number;
    /**
     * Labels to update of this size reservation
     *
     * @generated from field: optional metalstack.api.v2.UpdateLabels labels = 7;
     */
    labels?: UpdateLabels;
};
/**
 * Describes the message metalstack.admin.v2.SizeReservationServiceUpdateRequest.
 * Use `create(SizeReservationServiceUpdateRequestSchema)` to create a new message.
 */
export declare const SizeReservationServiceUpdateRequestSchema: GenMessage<SizeReservationServiceUpdateRequest>;
/**
 * SizeReservationServiceUpdateResponse is the response payload for a size reservation update request
 *
 * @generated from message metalstack.admin.v2.SizeReservationServiceUpdateResponse
 */
export type SizeReservationServiceUpdateResponse = Message<"metalstack.admin.v2.SizeReservationServiceUpdateResponse"> & {
    /**
     * SizeReservation the size reservation
     *
     * @generated from field: metalstack.api.v2.SizeReservation size_reservation = 1;
     */
    sizeReservation?: SizeReservation;
};
/**
 * Describes the message metalstack.admin.v2.SizeReservationServiceUpdateResponse.
 * Use `create(SizeReservationServiceUpdateResponseSchema)` to create a new message.
 */
export declare const SizeReservationServiceUpdateResponseSchema: GenMessage<SizeReservationServiceUpdateResponse>;
/**
 * SizeReservationServiceDeleteRequest is the request payload for a size reservation delete request
 *
 * @generated from message metalstack.admin.v2.SizeReservationServiceDeleteRequest
 */
export type SizeReservationServiceDeleteRequest = Message<"metalstack.admin.v2.SizeReservationServiceDeleteRequest"> & {
    /**
     * ID of the size reservation to delete
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message metalstack.admin.v2.SizeReservationServiceDeleteRequest.
 * Use `create(SizeReservationServiceDeleteRequestSchema)` to create a new message.
 */
export declare const SizeReservationServiceDeleteRequestSchema: GenMessage<SizeReservationServiceDeleteRequest>;
/**
 * SizeReservationServiceDeleteResponse is the response payload for a size reservation delete request
 *
 * @generated from message metalstack.admin.v2.SizeReservationServiceDeleteResponse
 */
export type SizeReservationServiceDeleteResponse = Message<"metalstack.admin.v2.SizeReservationServiceDeleteResponse"> & {
    /**
     * SizeReservation the size reservation
     *
     * @generated from field: metalstack.api.v2.SizeReservation size_reservation = 1;
     */
    sizeReservation?: SizeReservation;
};
/**
 * Describes the message metalstack.admin.v2.SizeReservationServiceDeleteResponse.
 * Use `create(SizeReservationServiceDeleteResponseSchema)` to create a new message.
 */
export declare const SizeReservationServiceDeleteResponseSchema: GenMessage<SizeReservationServiceDeleteResponse>;
/**
 * SizeReservationServiceListRequest is the request payload for a size list request
 *
 * @generated from message metalstack.admin.v2.SizeReservationServiceListRequest
 */
export type SizeReservationServiceListRequest = Message<"metalstack.admin.v2.SizeReservationServiceListRequest"> & {
    /**
     * Query for size reservations
     *
     * @generated from field: metalstack.api.v2.SizeReservationQuery query = 1;
     */
    query?: SizeReservationQuery;
};
/**
 * Describes the message metalstack.admin.v2.SizeReservationServiceListRequest.
 * Use `create(SizeReservationServiceListRequestSchema)` to create a new message.
 */
export declare const SizeReservationServiceListRequestSchema: GenMessage<SizeReservationServiceListRequest>;
/**
 * SizeReservationServiceListResponse is the response payload for a size reservation list request
 *
 * @generated from message metalstack.admin.v2.SizeReservationServiceListResponse
 */
export type SizeReservationServiceListResponse = Message<"metalstack.admin.v2.SizeReservationServiceListResponse"> & {
    /**
     * Size reservations
     *
     * @generated from field: repeated metalstack.api.v2.SizeReservation size_reservations = 1;
     */
    sizeReservations: SizeReservation[];
};
/**
 * Describes the message metalstack.admin.v2.SizeReservationServiceListResponse.
 * Use `create(SizeReservationServiceListResponseSchema)` to create a new message.
 */
export declare const SizeReservationServiceListResponseSchema: GenMessage<SizeReservationServiceListResponse>;
/**
 * SizeReservationService serves size reservation related functions
 *
 * @generated from service metalstack.admin.v2.SizeReservationService
 */
export declare const SizeReservationService: GenService<{
    /**
     * Create a size reservation
     *
     * @generated from rpc metalstack.admin.v2.SizeReservationService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof SizeReservationServiceCreateRequestSchema;
        output: typeof SizeReservationServiceCreateResponseSchema;
    };
    /**
     * Update a size reservation
     *
     * @generated from rpc metalstack.admin.v2.SizeReservationService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof SizeReservationServiceUpdateRequestSchema;
        output: typeof SizeReservationServiceUpdateResponseSchema;
    };
    /**
     * Delete a size reservation
     *
     * @generated from rpc metalstack.admin.v2.SizeReservationService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof SizeReservationServiceDeleteRequestSchema;
        output: typeof SizeReservationServiceDeleteResponseSchema;
    };
    /**
     * List size reservations
     *
     * @generated from rpc metalstack.admin.v2.SizeReservationService.List
     */
    list: {
        methodKind: "unary";
        input: typeof SizeReservationServiceListRequestSchema;
        output: typeof SizeReservationServiceListResponseSchema;
    };
}>;
