import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { UpdateMeta } from "../../api/v2/common_pb";
import type { Image, ImageClassification, ImageFeature, ImageQuery, ImageUsage } from "../../api/v2/image_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/image.proto.
 */
export declare const file_metalstack_admin_v2_image: GenFile;
/**
 * ImageServiceCreateRequest
 *
 * @generated from message metalstack.admin.v2.ImageServiceCreateRequest
 */
export type ImageServiceCreateRequest = Message<"metalstack.admin.v2.ImageServiceCreateRequest"> & {
    /**
     * Image is the image
     *
     * @generated from field: metalstack.api.v2.Image image = 1;
     */
    image?: Image;
};
/**
 * Describes the message metalstack.admin.v2.ImageServiceCreateRequest.
 * Use `create(ImageServiceCreateRequestSchema)` to create a new message.
 */
export declare const ImageServiceCreateRequestSchema: GenMessage<ImageServiceCreateRequest>;
/**
 * ImageServiceCreateResponse
 *
 * @generated from message metalstack.admin.v2.ImageServiceCreateResponse
 */
export type ImageServiceCreateResponse = Message<"metalstack.admin.v2.ImageServiceCreateResponse"> & {
    /**
     * Image is the image
     *
     * @generated from field: metalstack.api.v2.Image image = 1;
     */
    image?: Image;
};
/**
 * Describes the message metalstack.admin.v2.ImageServiceCreateResponse.
 * Use `create(ImageServiceCreateResponseSchema)` to create a new message.
 */
export declare const ImageServiceCreateResponseSchema: GenMessage<ImageServiceCreateResponse>;
/**
 * ImageServiceUpdateRequest
 *
 * @generated from message metalstack.admin.v2.ImageServiceUpdateRequest
 */
export type ImageServiceUpdateRequest = Message<"metalstack.admin.v2.ImageServiceUpdateRequest"> & {
    /**
     * Id of this image
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
     * URL where this image is located
     *
     * @generated from field: optional string url = 3;
     */
    url?: string;
    /**
     * Name of this imageLayout
     *
     * @generated from field: optional string name = 4;
     */
    name?: string;
    /**
     * Description of this imageLayout
     *
     * @generated from field: optional string description = 5;
     */
    description?: string;
    /**
     * Features of this image
     *
     * @generated from field: repeated metalstack.api.v2.ImageFeature features = 6;
     */
    features: ImageFeature[];
    /**
     * Classification of this image
     *
     * @generated from field: metalstack.api.v2.ImageClassification classification = 7;
     */
    classification: ImageClassification;
    /**
     * ExpiresAt usage is not possible after this date
     *
     * @generated from field: google.protobuf.Timestamp expires_at = 8;
     */
    expiresAt?: Timestamp;
};
/**
 * Describes the message metalstack.admin.v2.ImageServiceUpdateRequest.
 * Use `create(ImageServiceUpdateRequestSchema)` to create a new message.
 */
export declare const ImageServiceUpdateRequestSchema: GenMessage<ImageServiceUpdateRequest>;
/**
 * ImageServiceUpdateResponse
 *
 * @generated from message metalstack.admin.v2.ImageServiceUpdateResponse
 */
export type ImageServiceUpdateResponse = Message<"metalstack.admin.v2.ImageServiceUpdateResponse"> & {
    /**
     * Image is the image
     *
     * @generated from field: metalstack.api.v2.Image image = 1;
     */
    image?: Image;
};
/**
 * Describes the message metalstack.admin.v2.ImageServiceUpdateResponse.
 * Use `create(ImageServiceUpdateResponseSchema)` to create a new message.
 */
export declare const ImageServiceUpdateResponseSchema: GenMessage<ImageServiceUpdateResponse>;
/**
 * ImageServiceDeleteRequest
 *
 * @generated from message metalstack.admin.v2.ImageServiceDeleteRequest
 */
export type ImageServiceDeleteRequest = Message<"metalstack.admin.v2.ImageServiceDeleteRequest"> & {
    /**
     * ID of the image to delete
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message metalstack.admin.v2.ImageServiceDeleteRequest.
 * Use `create(ImageServiceDeleteRequestSchema)` to create a new message.
 */
export declare const ImageServiceDeleteRequestSchema: GenMessage<ImageServiceDeleteRequest>;
/**
 * message ImageServiceDeleteResponse {
 *
 * @generated from message metalstack.admin.v2.ImageServiceDeleteResponse
 */
export type ImageServiceDeleteResponse = Message<"metalstack.admin.v2.ImageServiceDeleteResponse"> & {
    /**
     * ImageLayout the imagelayout
     *
     * @generated from field: metalstack.api.v2.Image image = 1;
     */
    image?: Image;
};
/**
 * Describes the message metalstack.admin.v2.ImageServiceDeleteResponse.
 * Use `create(ImageServiceDeleteResponseSchema)` to create a new message.
 */
export declare const ImageServiceDeleteResponseSchema: GenMessage<ImageServiceDeleteResponse>;
/**
 * ImageServiceUsageRequest
 *
 * @generated from message metalstack.admin.v2.ImageServiceUsageRequest
 */
export type ImageServiceUsageRequest = Message<"metalstack.admin.v2.ImageServiceUsageRequest"> & {
    /**
     * Query for which images the usage should be reported
     *
     * @generated from field: metalstack.api.v2.ImageQuery query = 1;
     */
    query?: ImageQuery;
};
/**
 * Describes the message metalstack.admin.v2.ImageServiceUsageRequest.
 * Use `create(ImageServiceUsageRequestSchema)` to create a new message.
 */
export declare const ImageServiceUsageRequestSchema: GenMessage<ImageServiceUsageRequest>;
/**
 * ImageServiceUsageResponse
 *
 * @generated from message metalstack.admin.v2.ImageServiceUsageResponse
 */
export type ImageServiceUsageResponse = Message<"metalstack.admin.v2.ImageServiceUsageResponse"> & {
    /**
     * Images with usage
     *
     * @generated from field: repeated metalstack.api.v2.ImageUsage image_usage = 1;
     */
    imageUsage: ImageUsage[];
};
/**
 * Describes the message metalstack.admin.v2.ImageServiceUsageResponse.
 * Use `create(ImageServiceUsageResponseSchema)` to create a new message.
 */
export declare const ImageServiceUsageResponseSchema: GenMessage<ImageServiceUsageResponse>;
/**
 * ImageService serves image related functions
 *
 * @generated from service metalstack.admin.v2.ImageService
 */
export declare const ImageService: GenService<{
    /**
     * Create a image
     *
     * @generated from rpc metalstack.admin.v2.ImageService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof ImageServiceCreateRequestSchema;
        output: typeof ImageServiceCreateResponseSchema;
    };
    /**
     * Update a image
     *
     * @generated from rpc metalstack.admin.v2.ImageService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof ImageServiceUpdateRequestSchema;
        output: typeof ImageServiceUpdateResponseSchema;
    };
    /**
     * Delete a image
     *
     * @generated from rpc metalstack.admin.v2.ImageService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof ImageServiceDeleteRequestSchema;
        output: typeof ImageServiceDeleteResponseSchema;
    };
    /**
     * Usage of images
     *
     * @generated from rpc metalstack.admin.v2.ImageService.Usage
     */
    usage: {
        methodKind: "unary";
        input: typeof ImageServiceUsageRequestSchema;
        output: typeof ImageServiceUsageResponseSchema;
    };
}>;
