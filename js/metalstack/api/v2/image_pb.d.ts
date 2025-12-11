import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Labels, Meta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/image.proto.
 */
export declare const file_metalstack_api_v2_image: GenFile;
/**
 * ImageServiceGetRequest is the request payload for a image get request
 *
 * @generated from message metalstack.api.v2.ImageServiceGetRequest
 */
export type ImageServiceGetRequest = Message<"metalstack.api.v2.ImageServiceGetRequest"> & {
    /**
     * ID of the image to get
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message metalstack.api.v2.ImageServiceGetRequest.
 * Use `create(ImageServiceGetRequestSchema)` to create a new message.
 */
export declare const ImageServiceGetRequestSchema: GenMessage<ImageServiceGetRequest>;
/**
 * ImageServiceListRequest is the request payload for a image list request
 *
 * @generated from message metalstack.api.v2.ImageServiceListRequest
 */
export type ImageServiceListRequest = Message<"metalstack.api.v2.ImageServiceListRequest"> & {
    /**
     * Query for images
     *
     * @generated from field: metalstack.api.v2.ImageQuery query = 1;
     */
    query?: ImageQuery;
};
/**
 * Describes the message metalstack.api.v2.ImageServiceListRequest.
 * Use `create(ImageServiceListRequestSchema)` to create a new message.
 */
export declare const ImageServiceListRequestSchema: GenMessage<ImageServiceListRequest>;
/**
 * ImageServiceLatestRequest is the request payload for a image latest request
 *
 * @generated from message metalstack.api.v2.ImageServiceLatestRequest
 */
export type ImageServiceLatestRequest = Message<"metalstack.api.v2.ImageServiceLatestRequest"> & {
    /**
     * OS for which the latest image should be fetched
     * should contain os and major.minor then latest patch version of this os is returned
     *
     * @generated from field: string os = 1;
     */
    os: string;
};
/**
 * Describes the message metalstack.api.v2.ImageServiceLatestRequest.
 * Use `create(ImageServiceLatestRequestSchema)` to create a new message.
 */
export declare const ImageServiceLatestRequestSchema: GenMessage<ImageServiceLatestRequest>;
/**
 * ImageServiceGetResponse is the response payload for a image get request
 *
 * @generated from message metalstack.api.v2.ImageServiceGetResponse
 */
export type ImageServiceGetResponse = Message<"metalstack.api.v2.ImageServiceGetResponse"> & {
    /**
     * Image the image
     *
     * @generated from field: metalstack.api.v2.Image image = 1;
     */
    image?: Image;
};
/**
 * Describes the message metalstack.api.v2.ImageServiceGetResponse.
 * Use `create(ImageServiceGetResponseSchema)` to create a new message.
 */
export declare const ImageServiceGetResponseSchema: GenMessage<ImageServiceGetResponse>;
/**
 * ImageServiceListResponse is the response payload for a image list request
 *
 * @generated from message metalstack.api.v2.ImageServiceListResponse
 */
export type ImageServiceListResponse = Message<"metalstack.api.v2.ImageServiceListResponse"> & {
    /**
     * Images the images
     *
     * @generated from field: repeated metalstack.api.v2.Image images = 1;
     */
    images: Image[];
};
/**
 * Describes the message metalstack.api.v2.ImageServiceListResponse.
 * Use `create(ImageServiceListResponseSchema)` to create a new message.
 */
export declare const ImageServiceListResponseSchema: GenMessage<ImageServiceListResponse>;
/**
 * ImageServiceLatestResponse is the response payload for a image latest request
 *
 * @generated from message metalstack.api.v2.ImageServiceLatestResponse
 */
export type ImageServiceLatestResponse = Message<"metalstack.api.v2.ImageServiceLatestResponse"> & {
    /**
     * Image which is the latest for one os
     *
     * @generated from field: metalstack.api.v2.Image image = 1;
     */
    image?: Image;
};
/**
 * Describes the message metalstack.api.v2.ImageServiceLatestResponse.
 * Use `create(ImageServiceLatestResponseSchema)` to create a new message.
 */
export declare const ImageServiceLatestResponseSchema: GenMessage<ImageServiceLatestResponse>;
/**
 * Image
 *
 * @generated from message metalstack.api.v2.Image
 */
export type Image = Message<"metalstack.api.v2.Image"> & {
    /**
     * Id of this imageLayout
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * Meta for this ip
     *
     * @generated from field: metalstack.api.v2.Meta meta = 2;
     */
    meta?: Meta;
    /**
     * URL where this image is located
     *
     * @generated from field: string url = 3;
     */
    url: string;
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
 * Describes the message metalstack.api.v2.Image.
 * Use `create(ImageSchema)` to create a new message.
 */
export declare const ImageSchema: GenMessage<Image>;
/**
 * ImageUsage reports which machines/firewalls actually use this image
 *
 * @generated from message metalstack.api.v2.ImageUsage
 */
export type ImageUsage = Message<"metalstack.api.v2.ImageUsage"> & {
    /**
     * Image with usage
     *
     * @generated from field: metalstack.api.v2.Image image = 1;
     */
    image?: Image;
    /**
     * UsedBy the following machines/firewalls
     *
     * @generated from field: repeated string used_by = 2;
     */
    usedBy: string[];
};
/**
 * Describes the message metalstack.api.v2.ImageUsage.
 * Use `create(ImageUsageSchema)` to create a new message.
 */
export declare const ImageUsageSchema: GenMessage<ImageUsage>;
/**
 * ImageQuery is used to search images
 *
 * @generated from message metalstack.api.v2.ImageQuery
 */
export type ImageQuery = Message<"metalstack.api.v2.ImageQuery"> & {
    /**
     * ID of the image to get
     *
     * @generated from field: optional string id = 1;
     */
    id?: string;
    /**
     * OS of the image
     *
     * @generated from field: optional string os = 2;
     */
    os?: string;
    /**
     * Version of the Image
     *
     * @generated from field: optional string version = 3;
     */
    version?: string;
    /**
     * Name of the image to query
     *
     * @generated from field: optional string name = 4;
     */
    name?: string;
    /**
     * Description of the image to query
     *
     * @generated from field: optional string description = 5;
     */
    description?: string;
    /**
     * Url of the image to query
     *
     * @generated from field: optional string url = 6;
     */
    url?: string;
    /**
     * Feature of the image to query
     *
     * @generated from field: optional metalstack.api.v2.ImageFeature feature = 7;
     */
    feature?: ImageFeature;
    /**
     * Classification of the image to query
     *
     * @generated from field: optional metalstack.api.v2.ImageClassification classification = 8;
     */
    classification?: ImageClassification;
    /**
     * Labels lists only images containing the given labels
     *
     * @generated from field: optional metalstack.api.v2.Labels labels = 9;
     */
    labels?: Labels;
};
/**
 * Describes the message metalstack.api.v2.ImageQuery.
 * Use `create(ImageQuerySchema)` to create a new message.
 */
export declare const ImageQuerySchema: GenMessage<ImageQuery>;
/**
 * ImageFeature
 *
 * @generated from enum metalstack.api.v2.ImageFeature
 */
export declare enum ImageFeature {
    /**
     * IMAGE_FEATURE_UNSPECIFIED is not specified
     *
     * @generated from enum value: IMAGE_FEATURE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * IMAGE_FEATURE_MACHINE indicates this image is usable for a machine
     *
     * @generated from enum value: IMAGE_FEATURE_MACHINE = 1;
     */
    MACHINE = 1,
    /**
     * IMAGE_FEATURE_FIREWALL indicates this image is usable for a firewall
     *
     * @generated from enum value: IMAGE_FEATURE_FIREWALL = 2;
     */
    FIREWALL = 2
}
/**
 * Describes the enum metalstack.api.v2.ImageFeature.
 */
export declare const ImageFeatureSchema: GenEnum<ImageFeature>;
/**
 * Image
 *
 * @generated from enum metalstack.api.v2.ImageClassification
 */
export declare enum ImageClassification {
    /**
     * IMAGE_CLASSIFICATION_UNSPECIFIED is not specified
     *
     * @generated from enum value: IMAGE_CLASSIFICATION_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * IMAGE_CLASSIFICATION_PREVIEW indicates that this image is in preview
     *
     * @generated from enum value: IMAGE_CLASSIFICATION_PREVIEW = 1;
     */
    PREVIEW = 1,
    /**
     * IMAGE_CLASSIFICATION_SUPPORTED indicates that this image is supported
     *
     * @generated from enum value: IMAGE_CLASSIFICATION_SUPPORTED = 2;
     */
    SUPPORTED = 2,
    /**
     * IMAGE_CLASSIFICATION_DEPRECATED indicates that this image is deprecated
     *
     * @generated from enum value: IMAGE_CLASSIFICATION_DEPRECATED = 3;
     */
    DEPRECATED = 3
}
/**
 * Describes the enum metalstack.api.v2.ImageClassification.
 */
export declare const ImageClassificationSchema: GenEnum<ImageClassification>;
/**
 * ImageService serves image related functions
 *
 * @generated from service metalstack.api.v2.ImageService
 */
export declare const ImageService: GenService<{
    /**
     * Get a image
     *
     * @generated from rpc metalstack.api.v2.ImageService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof ImageServiceGetRequestSchema;
        output: typeof ImageServiceGetResponseSchema;
    };
    /**
     * List all images
     *
     * @generated from rpc metalstack.api.v2.ImageService.List
     */
    list: {
        methodKind: "unary";
        input: typeof ImageServiceListRequestSchema;
        output: typeof ImageServiceListResponseSchema;
    };
    /**
     * Latest image for a specific os
     *
     * @generated from rpc metalstack.api.v2.ImageService.Latest
     */
    latest: {
        methodKind: "unary";
        input: typeof ImageServiceLatestRequestSchema;
        output: typeof ImageServiceLatestResponseSchema;
    };
}>;
