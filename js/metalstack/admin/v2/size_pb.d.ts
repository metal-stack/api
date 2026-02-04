import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { UpdateLabels, UpdateMeta } from "../../api/v2/common_pb";
import type { Size, SizeConstraint } from "../../api/v2/size_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/size.proto.
 */
export declare const file_metalstack_admin_v2_size: GenFile;
/**
 * SizeServiceCreateRequest is the request payload for a size create request
 *
 * @generated from message metalstack.admin.v2.SizeServiceCreateRequest
 */
export type SizeServiceCreateRequest = Message<"metalstack.admin.v2.SizeServiceCreateRequest"> & {
    /**
     * Size is the size to create
     *
     * @generated from field: metalstack.api.v2.Size size = 1;
     */
    size?: Size;
};
/**
 * Describes the message metalstack.admin.v2.SizeServiceCreateRequest.
 * Use `create(SizeServiceCreateRequestSchema)` to create a new message.
 */
export declare const SizeServiceCreateRequestSchema: GenMessage<SizeServiceCreateRequest>;
/**
 * SizeServiceGetResponse is the response payload for a size create request
 *
 * @generated from message metalstack.admin.v2.SizeServiceCreateResponse
 */
export type SizeServiceCreateResponse = Message<"metalstack.admin.v2.SizeServiceCreateResponse"> & {
    /**
     * Size the size
     *
     * @generated from field: metalstack.api.v2.Size size = 1;
     */
    size?: Size;
};
/**
 * Describes the message metalstack.admin.v2.SizeServiceCreateResponse.
 * Use `create(SizeServiceCreateResponseSchema)` to create a new message.
 */
export declare const SizeServiceCreateResponseSchema: GenMessage<SizeServiceCreateResponse>;
/**
 * SizeServiceUpdateRequest is the request payload for a size update request
 *
 * @generated from message metalstack.admin.v2.SizeServiceUpdateRequest
 */
export type SizeServiceUpdateRequest = Message<"metalstack.admin.v2.SizeServiceUpdateRequest"> & {
    /**
     * Id of this size
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
     * Name of this size
     *
     * @generated from field: optional string name = 3;
     */
    name?: string;
    /**
     * Description of this size
     *
     * @generated from field: optional string description = 4;
     */
    description?: string;
    /**
     * Constraints which must match that a specific machine is considered of this size
     *
     * @generated from field: repeated metalstack.api.v2.SizeConstraint constraints = 5;
     */
    constraints: SizeConstraint[];
    /**
     * Labels to update on this size
     *
     * @generated from field: optional metalstack.api.v2.UpdateLabels labels = 6;
     */
    labels?: UpdateLabels;
};
/**
 * Describes the message metalstack.admin.v2.SizeServiceUpdateRequest.
 * Use `create(SizeServiceUpdateRequestSchema)` to create a new message.
 */
export declare const SizeServiceUpdateRequestSchema: GenMessage<SizeServiceUpdateRequest>;
/**
 * SizeServiceUpdateResponse is the response payload for a size update request
 *
 * @generated from message metalstack.admin.v2.SizeServiceUpdateResponse
 */
export type SizeServiceUpdateResponse = Message<"metalstack.admin.v2.SizeServiceUpdateResponse"> & {
    /**
     * Size the size
     *
     * @generated from field: metalstack.api.v2.Size size = 1;
     */
    size?: Size;
};
/**
 * Describes the message metalstack.admin.v2.SizeServiceUpdateResponse.
 * Use `create(SizeServiceUpdateResponseSchema)` to create a new message.
 */
export declare const SizeServiceUpdateResponseSchema: GenMessage<SizeServiceUpdateResponse>;
/**
 * SizeServiceDeleteRequest is the request payload for a size delete request
 *
 * @generated from message metalstack.admin.v2.SizeServiceDeleteRequest
 */
export type SizeServiceDeleteRequest = Message<"metalstack.admin.v2.SizeServiceDeleteRequest"> & {
    /**
     * ID of the size to delete
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message metalstack.admin.v2.SizeServiceDeleteRequest.
 * Use `create(SizeServiceDeleteRequestSchema)` to create a new message.
 */
export declare const SizeServiceDeleteRequestSchema: GenMessage<SizeServiceDeleteRequest>;
/**
 * SizeServiceDeleteResponse is the response payload for a size delete request
 *
 * @generated from message metalstack.admin.v2.SizeServiceDeleteResponse
 */
export type SizeServiceDeleteResponse = Message<"metalstack.admin.v2.SizeServiceDeleteResponse"> & {
    /**
     * Size the size
     *
     * @generated from field: metalstack.api.v2.Size size = 1;
     */
    size?: Size;
};
/**
 * Describes the message metalstack.admin.v2.SizeServiceDeleteResponse.
 * Use `create(SizeServiceDeleteResponseSchema)` to create a new message.
 */
export declare const SizeServiceDeleteResponseSchema: GenMessage<SizeServiceDeleteResponse>;
/**
 * SizeService serves size related functions
 *
 * @generated from service metalstack.admin.v2.SizeService
 */
export declare const SizeService: GenService<{
    /**
     * Create a size
     *
     * @generated from rpc metalstack.admin.v2.SizeService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof SizeServiceCreateRequestSchema;
        output: typeof SizeServiceCreateResponseSchema;
    };
    /**
     * Update a size
     *
     * @generated from rpc metalstack.admin.v2.SizeService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof SizeServiceUpdateRequestSchema;
        output: typeof SizeServiceUpdateResponseSchema;
    };
    /**
     * Delete a size
     *
     * @generated from rpc metalstack.admin.v2.SizeService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof SizeServiceDeleteRequestSchema;
        output: typeof SizeServiceDeleteResponseSchema;
    };
}>;
