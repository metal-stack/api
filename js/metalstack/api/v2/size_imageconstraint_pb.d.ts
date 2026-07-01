import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Meta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/size_imageconstraint.proto.
 */
export declare const file_metalstack_api_v2_size_imageconstraint: GenFile;
/**
 * SizeImageConstraintServiceTryRequest is the request payload for a size image constraint try request
 *
 * @generated from message metalstack.api.v2.SizeImageConstraintServiceTryRequest
 */
export type SizeImageConstraintServiceTryRequest = Message<"metalstack.api.v2.SizeImageConstraintServiceTryRequest"> & {
    /**
     * Size to try
     *
     * @generated from field: string size = 1;
     */
    size: string;
    /**
     * Image to try
     *
     * @generated from field: string image = 2;
     */
    image: string;
};
/**
 * Describes the message metalstack.api.v2.SizeImageConstraintServiceTryRequest.
 * Use `create(SizeImageConstraintServiceTryRequestSchema)` to create a new message.
 */
export declare const SizeImageConstraintServiceTryRequestSchema: GenMessage<SizeImageConstraintServiceTryRequest>;
/**
 * SizeImageConstraintServiceTryResponse is the response payload for a size image constraint try request
 *
 * @generated from message metalstack.api.v2.SizeImageConstraintServiceTryResponse
 */
export type SizeImageConstraintServiceTryResponse = Message<"metalstack.api.v2.SizeImageConstraintServiceTryResponse"> & {};
/**
 * Describes the message metalstack.api.v2.SizeImageConstraintServiceTryResponse.
 * Use `create(SizeImageConstraintServiceTryResponseSchema)` to create a new message.
 */
export declare const SizeImageConstraintServiceTryResponseSchema: GenMessage<SizeImageConstraintServiceTryResponse>;
/**
 * SizeImageConstraint expresses optional restrictions for specific size to image combinations
 * this might be required if the support for a specific hardware in a given size is only supported
 * with a newer version of the image.
 *
 * If the size in question is not found, no restrictions apply.
 * If the image in question is not found, no restrictions apply as well.
 * If the image in question is found, but does not match the given expression, machine creation must be forbidden.
 *
 * @generated from message metalstack.api.v2.SizeImageConstraint
 */
export type SizeImageConstraint = Message<"metalstack.api.v2.SizeImageConstraint"> & {
    /**
     * Size where this constraint should apply
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
    meta?: Meta | undefined;
    /**
     * Name of this size image constraint
     *
     * @generated from field: optional string name = 4;
     */
    name?: string | undefined;
    /**
     * Description of this size image constraint
     *
     * @generated from field: optional string description = 5;
     */
    description?: string | undefined;
};
/**
 * Describes the message metalstack.api.v2.SizeImageConstraint.
 * Use `create(SizeImageConstraintSchema)` to create a new message.
 */
export declare const SizeImageConstraintSchema: GenMessage<SizeImageConstraint>;
/**
 * ImageConstraint defines a constraint for a image
 * examples:
 * images:
 *    ubuntu: ">= 20.04.20211011"
 *    debian: ">= 10.0.20210101"
 *
 * @generated from message metalstack.api.v2.ImageConstraint
 */
export type ImageConstraint = Message<"metalstack.api.v2.ImageConstraint"> & {
    /**
     * Image of the constraint
     *
     * @generated from field: string image = 1;
     */
    image: string;
    /**
     * SemverMatch which defines in semver match format which image version should apply
     *
     * @generated from field: string semver_match = 2;
     */
    semverMatch: string;
};
/**
 * Describes the message metalstack.api.v2.ImageConstraint.
 * Use `create(ImageConstraintSchema)` to create a new message.
 */
export declare const ImageConstraintSchema: GenMessage<ImageConstraint>;
/**
 * SizeImageConstraintQuery is used to search size image constraints
 *
 * @generated from message metalstack.api.v2.SizeImageConstraintQuery
 */
export type SizeImageConstraintQuery = Message<"metalstack.api.v2.SizeImageConstraintQuery"> & {
    /**
     * Size of the size image constraint
     *
     * @generated from field: optional string size = 1;
     */
    size?: string | undefined;
    /**
     * Name of the size image constraint to query
     *
     * @generated from field: optional string name = 2;
     */
    name?: string | undefined;
    /**
     * Description of the size image constraint to query
     *
     * @generated from field: optional string description = 3;
     */
    description?: string | undefined;
};
/**
 * Describes the message metalstack.api.v2.SizeImageConstraintQuery.
 * Use `create(SizeImageConstraintQuerySchema)` to create a new message.
 */
export declare const SizeImageConstraintQuerySchema: GenMessage<SizeImageConstraintQuery>;
/**
 * SizeImageConstraintService provides size and image constraint validation operations.
 *
 * @generated from service metalstack.api.v2.SizeImageConstraintService
 */
export declare const SizeImageConstraintService: GenService<{
    /**
     * Try validates if a given combination of size and image is possible.
     *
     * @generated from rpc metalstack.api.v2.SizeImageConstraintService.Try
     */
    try: {
        methodKind: "unary";
        input: typeof SizeImageConstraintServiceTryRequestSchema;
        output: typeof SizeImageConstraintServiceTryResponseSchema;
    };
}>;
