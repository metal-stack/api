import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Meta, UpdateMeta } from "../../api/v2/common_pb";
import type { ImageConstraint, SizeImageConstraint, SizeImageConstraintQuery } from "../../api/v2/size_imageconstraint_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/size_imageconstraint.proto.
 */
export declare const file_metalstack_admin_v2_size_imageconstraint: GenFile;
/**
 * SizeImageConstraintServiceCreateRequest is the request payload for a size image constraint create request
 *
 * @generated from message metalstack.admin.v2.SizeImageConstraintServiceCreateRequest
 */
export type SizeImageConstraintServiceCreateRequest = Message<"metalstack.admin.v2.SizeImageConstraintServiceCreateRequest"> & {
    /**
     * Id of the size
     *
     * @generated from field: string size = 1;
     */
    size: string;
    /**
     * ImageConstraints to apply to this size
     *
     * @generated from field: repeated metalstack.api.v2.ImageConstraint image_constraints = 2;
     */
    imageConstraints: ImageConstraint[];
    /**
     * Meta for this size image constraint
     *
     * @generated from field: metalstack.api.v2.Meta meta = 3;
     */
    meta?: Meta;
    /**
     * Name of this size image constraint
     *
     * @generated from field: optional string name = 4;
     */
    name?: string;
    /**
     * Description of this size image constraint
     *
     * @generated from field: optional string description = 5;
     */
    description?: string;
};
/**
 * Describes the message metalstack.admin.v2.SizeImageConstraintServiceCreateRequest.
 * Use `create(SizeImageConstraintServiceCreateRequestSchema)` to create a new message.
 */
export declare const SizeImageConstraintServiceCreateRequestSchema: GenMessage<SizeImageConstraintServiceCreateRequest>;
/**
 * SizeImageConstraintServiceCreateResponse is the response payload for a size image constraint create request
 *
 * @generated from message metalstack.admin.v2.SizeImageConstraintServiceCreateResponse
 */
export type SizeImageConstraintServiceCreateResponse = Message<"metalstack.admin.v2.SizeImageConstraintServiceCreateResponse"> & {
    /**
     * SizeImageConstraint created
     *
     * @generated from field: metalstack.api.v2.SizeImageConstraint size_image_constraint = 1;
     */
    sizeImageConstraint?: SizeImageConstraint;
};
/**
 * Describes the message metalstack.admin.v2.SizeImageConstraintServiceCreateResponse.
 * Use `create(SizeImageConstraintServiceCreateResponseSchema)` to create a new message.
 */
export declare const SizeImageConstraintServiceCreateResponseSchema: GenMessage<SizeImageConstraintServiceCreateResponse>;
/**
 * SizeImageConstraintServiceUpdateRequest is the request payload for a size image constraint update request
 *
 * @generated from message metalstack.admin.v2.SizeImageConstraintServiceUpdateRequest
 */
export type SizeImageConstraintServiceUpdateRequest = Message<"metalstack.admin.v2.SizeImageConstraintServiceUpdateRequest"> & {
    /**
     * Id of the size
     *
     * @generated from field: string size = 1;
     */
    size: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request
     *
     * @generated from field: metalstack.api.v2.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta;
    /**
     * ImageConstraints to apply to this size
     *
     * @generated from field: repeated metalstack.api.v2.ImageConstraint image_constraints = 3;
     */
    imageConstraints: ImageConstraint[];
    /**
     * Name of this size image constraint
     *
     * @generated from field: optional string name = 4;
     */
    name?: string;
    /**
     * Description of this size image constraint
     *
     * @generated from field: optional string description = 5;
     */
    description?: string;
};
/**
 * Describes the message metalstack.admin.v2.SizeImageConstraintServiceUpdateRequest.
 * Use `create(SizeImageConstraintServiceUpdateRequestSchema)` to create a new message.
 */
export declare const SizeImageConstraintServiceUpdateRequestSchema: GenMessage<SizeImageConstraintServiceUpdateRequest>;
/**
 * SizeImageConstraintServiceUpdateResponse is the response payload for a size image constraint update request
 *
 * @generated from message metalstack.admin.v2.SizeImageConstraintServiceUpdateResponse
 */
export type SizeImageConstraintServiceUpdateResponse = Message<"metalstack.admin.v2.SizeImageConstraintServiceUpdateResponse"> & {
    /**
     * SizeImageConstraint updated
     *
     * @generated from field: metalstack.api.v2.SizeImageConstraint size_image_constraint = 1;
     */
    sizeImageConstraint?: SizeImageConstraint;
};
/**
 * Describes the message metalstack.admin.v2.SizeImageConstraintServiceUpdateResponse.
 * Use `create(SizeImageConstraintServiceUpdateResponseSchema)` to create a new message.
 */
export declare const SizeImageConstraintServiceUpdateResponseSchema: GenMessage<SizeImageConstraintServiceUpdateResponse>;
/**
 * SizeImageConstraintServiceDeleteRequest is the request payload for a size image constraint delete request
 *
 * @generated from message metalstack.admin.v2.SizeImageConstraintServiceDeleteRequest
 */
export type SizeImageConstraintServiceDeleteRequest = Message<"metalstack.admin.v2.SizeImageConstraintServiceDeleteRequest"> & {
    /**
     * Id of the size
     *
     * @generated from field: string size = 1;
     */
    size: string;
};
/**
 * Describes the message metalstack.admin.v2.SizeImageConstraintServiceDeleteRequest.
 * Use `create(SizeImageConstraintServiceDeleteRequestSchema)` to create a new message.
 */
export declare const SizeImageConstraintServiceDeleteRequestSchema: GenMessage<SizeImageConstraintServiceDeleteRequest>;
/**
 * SizeImageConstraintServiceDeleteResponse is the response payload for a size image constraint delete request
 *
 * @generated from message metalstack.admin.v2.SizeImageConstraintServiceDeleteResponse
 */
export type SizeImageConstraintServiceDeleteResponse = Message<"metalstack.admin.v2.SizeImageConstraintServiceDeleteResponse"> & {
    /**
     * SizeImageConstraint deleted
     *
     * @generated from field: metalstack.api.v2.SizeImageConstraint size_image_constraint = 1;
     */
    sizeImageConstraint?: SizeImageConstraint;
};
/**
 * Describes the message metalstack.admin.v2.SizeImageConstraintServiceDeleteResponse.
 * Use `create(SizeImageConstraintServiceDeleteResponseSchema)` to create a new message.
 */
export declare const SizeImageConstraintServiceDeleteResponseSchema: GenMessage<SizeImageConstraintServiceDeleteResponse>;
/**
 * SizeImageConstraintServiceGetRequest is the request payload for a size image constraint get request
 *
 * @generated from message metalstack.admin.v2.SizeImageConstraintServiceGetRequest
 */
export type SizeImageConstraintServiceGetRequest = Message<"metalstack.admin.v2.SizeImageConstraintServiceGetRequest"> & {
    /**
     * Id of the size to get
     *
     * @generated from field: string size = 1;
     */
    size: string;
};
/**
 * Describes the message metalstack.admin.v2.SizeImageConstraintServiceGetRequest.
 * Use `create(SizeImageConstraintServiceGetRequestSchema)` to create a new message.
 */
export declare const SizeImageConstraintServiceGetRequestSchema: GenMessage<SizeImageConstraintServiceGetRequest>;
/**
 * SizeImageConstraintServiceGetResponse is the response payload for a size image constraint get request
 *
 * @generated from message metalstack.admin.v2.SizeImageConstraintServiceGetResponse
 */
export type SizeImageConstraintServiceGetResponse = Message<"metalstack.admin.v2.SizeImageConstraintServiceGetResponse"> & {
    /**
     * SizeImageConstraint get
     *
     * @generated from field: metalstack.api.v2.SizeImageConstraint size_image_constraint = 1;
     */
    sizeImageConstraint?: SizeImageConstraint;
};
/**
 * Describes the message metalstack.admin.v2.SizeImageConstraintServiceGetResponse.
 * Use `create(SizeImageConstraintServiceGetResponseSchema)` to create a new message.
 */
export declare const SizeImageConstraintServiceGetResponseSchema: GenMessage<SizeImageConstraintServiceGetResponse>;
/**
 * SizeImageConstraintServiceListRequest is the request payload for a size image constraint list request
 *
 * @generated from message metalstack.admin.v2.SizeImageConstraintServiceListRequest
 */
export type SizeImageConstraintServiceListRequest = Message<"metalstack.admin.v2.SizeImageConstraintServiceListRequest"> & {
    /**
     * Query for size image constraints
     *
     * @generated from field: metalstack.api.v2.SizeImageConstraintQuery query = 1;
     */
    query?: SizeImageConstraintQuery;
};
/**
 * Describes the message metalstack.admin.v2.SizeImageConstraintServiceListRequest.
 * Use `create(SizeImageConstraintServiceListRequestSchema)` to create a new message.
 */
export declare const SizeImageConstraintServiceListRequestSchema: GenMessage<SizeImageConstraintServiceListRequest>;
/**
 * SizeImageConstraintServiceListResponse is the response payload for a size image constraint list request
 *
 * @generated from message metalstack.admin.v2.SizeImageConstraintServiceListResponse
 */
export type SizeImageConstraintServiceListResponse = Message<"metalstack.admin.v2.SizeImageConstraintServiceListResponse"> & {
    /**
     * SizeImageConstraints listed
     *
     * @generated from field: repeated metalstack.api.v2.SizeImageConstraint size_image_constraints = 1;
     */
    sizeImageConstraints: SizeImageConstraint[];
};
/**
 * Describes the message metalstack.admin.v2.SizeImageConstraintServiceListResponse.
 * Use `create(SizeImageConstraintServiceListResponseSchema)` to create a new message.
 */
export declare const SizeImageConstraintServiceListResponseSchema: GenMessage<SizeImageConstraintServiceListResponse>;
/**
 * SizeImageConstraintService serves size and image constraint related functions
 *
 * @generated from service metalstack.admin.v2.SizeImageConstraintService
 */
export declare const SizeImageConstraintService: GenService<{
    /**
     * Create a size image constraint
     *
     * @generated from rpc metalstack.admin.v2.SizeImageConstraintService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof SizeImageConstraintServiceCreateRequestSchema;
        output: typeof SizeImageConstraintServiceCreateResponseSchema;
    };
    /**
     * Update a size image constraint
     *
     * @generated from rpc metalstack.admin.v2.SizeImageConstraintService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof SizeImageConstraintServiceUpdateRequestSchema;
        output: typeof SizeImageConstraintServiceUpdateResponseSchema;
    };
    /**
     * Delete a size image constraint
     *
     * @generated from rpc metalstack.admin.v2.SizeImageConstraintService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof SizeImageConstraintServiceDeleteRequestSchema;
        output: typeof SizeImageConstraintServiceDeleteResponseSchema;
    };
    /**
     * Get a size image constraint
     *
     * @generated from rpc metalstack.admin.v2.SizeImageConstraintService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof SizeImageConstraintServiceGetRequestSchema;
        output: typeof SizeImageConstraintServiceGetResponseSchema;
    };
    /**
     * List a size image constraint
     *
     * @generated from rpc metalstack.admin.v2.SizeImageConstraintService.List
     */
    list: {
        methodKind: "unary";
        input: typeof SizeImageConstraintServiceListRequestSchema;
        output: typeof SizeImageConstraintServiceListResponseSchema;
    };
}>;
