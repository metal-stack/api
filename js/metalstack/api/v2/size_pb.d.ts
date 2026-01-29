import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Labels, Meta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/size.proto.
 */
export declare const file_metalstack_api_v2_size: GenFile;
/**
 * SizeServiceGetRequest is the request payload for a size get request
 *
 * @generated from message metalstack.api.v2.SizeServiceGetRequest
 */
export type SizeServiceGetRequest = Message<"metalstack.api.v2.SizeServiceGetRequest"> & {
    /**
     * ID of the size to get
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message metalstack.api.v2.SizeServiceGetRequest.
 * Use `create(SizeServiceGetRequestSchema)` to create a new message.
 */
export declare const SizeServiceGetRequestSchema: GenMessage<SizeServiceGetRequest>;
/**
 * SizeServiceListRequest is the request payload for a size list request
 *
 * @generated from message metalstack.api.v2.SizeServiceListRequest
 */
export type SizeServiceListRequest = Message<"metalstack.api.v2.SizeServiceListRequest"> & {
    /**
     * Query for sizes
     *
     * @generated from field: metalstack.api.v2.SizeQuery query = 1;
     */
    query?: SizeQuery;
};
/**
 * Describes the message metalstack.api.v2.SizeServiceListRequest.
 * Use `create(SizeServiceListRequestSchema)` to create a new message.
 */
export declare const SizeServiceListRequestSchema: GenMessage<SizeServiceListRequest>;
/**
 * SizeServiceGetResponse is the response payload for a size get request
 *
 * @generated from message metalstack.api.v2.SizeServiceGetResponse
 */
export type SizeServiceGetResponse = Message<"metalstack.api.v2.SizeServiceGetResponse"> & {
    /**
     * Size the size
     *
     * @generated from field: metalstack.api.v2.Size size = 1;
     */
    size?: Size;
};
/**
 * Describes the message metalstack.api.v2.SizeServiceGetResponse.
 * Use `create(SizeServiceGetResponseSchema)` to create a new message.
 */
export declare const SizeServiceGetResponseSchema: GenMessage<SizeServiceGetResponse>;
/**
 * SizeServiceListResponse is the response payload for a size list request
 *
 * @generated from message metalstack.api.v2.SizeServiceListResponse
 */
export type SizeServiceListResponse = Message<"metalstack.api.v2.SizeServiceListResponse"> & {
    /**
     * Sizes the sizes
     *
     * @generated from field: repeated metalstack.api.v2.Size sizes = 1;
     */
    sizes: Size[];
};
/**
 * Describes the message metalstack.api.v2.SizeServiceListResponse.
 * Use `create(SizeServiceListResponseSchema)` to create a new message.
 */
export declare const SizeServiceListResponseSchema: GenMessage<SizeServiceListResponse>;
/**
 * Size
 *
 * @generated from message metalstack.api.v2.Size
 */
export type Size = Message<"metalstack.api.v2.Size"> & {
    /**
     * Id of this size
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * Meta for this size
     *
     * @generated from field: metalstack.api.v2.Meta meta = 2;
     */
    meta?: Meta;
    /**
     * Name of this size
     *
     * @generated from field: optional string name = 4;
     */
    name?: string;
    /**
     * Description of this size
     *
     * @generated from field: optional string description = 5;
     */
    description?: string;
    /**
     * Constraints which must match that a specific machine is considered of this size
     *
     * @generated from field: repeated metalstack.api.v2.SizeConstraint constraints = 6;
     */
    constraints: SizeConstraint[];
};
/**
 * Describes the message metalstack.api.v2.Size.
 * Use `create(SizeSchema)` to create a new message.
 */
export declare const SizeSchema: GenMessage<Size>;
/**
 * SizeConstraint defines the boundaries for certain type of machine property which must match to identify this machine as this size.
 *
 * @generated from message metalstack.api.v2.SizeConstraint
 */
export type SizeConstraint = Message<"metalstack.api.v2.SizeConstraint"> & {
    /**
     * Type a machine matches to a size in order to make them easier to categorize
     *
     * @generated from field: metalstack.api.v2.SizeConstraintType type = 1;
     */
    type: SizeConstraintType;
    /**
     * Min the minimum value of the constraint
     *
     * @generated from field: uint64 min = 2;
     */
    min: bigint;
    /**
     * Max the maximum value of the constraint
     *
     * @generated from field: uint64 max = 3;
     */
    max: bigint;
    /**
     * Identifier glob pattern which matches to the given type, for example gpu pci id
     *
     * @generated from field: optional string identifier = 4;
     */
    identifier?: string;
};
/**
 * Describes the message metalstack.api.v2.SizeConstraint.
 * Use `create(SizeConstraintSchema)` to create a new message.
 */
export declare const SizeConstraintSchema: GenMessage<SizeConstraint>;
/**
 * SizeQuery is used to search sizes
 *
 * @generated from message metalstack.api.v2.SizeQuery
 */
export type SizeQuery = Message<"metalstack.api.v2.SizeQuery"> & {
    /**
     * ID of the size to get
     *
     * @generated from field: optional string id = 1;
     */
    id?: string;
    /**
     * Name of the size to query
     *
     * @generated from field: optional string name = 2;
     */
    name?: string;
    /**
     * Description of the size to query
     *
     * @generated from field: optional string description = 3;
     */
    description?: string;
    /**
     * Labels lists only sizes containing the given labels
     *
     * @generated from field: optional metalstack.api.v2.Labels labels = 4;
     */
    labels?: Labels;
};
/**
 * Describes the message metalstack.api.v2.SizeQuery.
 * Use `create(SizeQuerySchema)` to create a new message.
 */
export declare const SizeQuerySchema: GenMessage<SizeQuery>;
/**
 * SizeConstraintType defines the property for which a constraint is defined
 *
 * @generated from enum metalstack.api.v2.SizeConstraintType
 */
export declare enum SizeConstraintType {
    /**
     * SIZE_CONSTRAINT_TYPE_UNSPECIFIED type is not specified
     *
     * @generated from enum value: SIZE_CONSTRAINT_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * SIZE_CONSTRAINT_TYPE_CORES defines the number of cores as property
     *
     * @generated from enum value: SIZE_CONSTRAINT_TYPE_CORES = 1;
     */
    CORES = 1,
    /**
     * SIZE_CONSTRAINT_TYPE_MEMORY defines the amount of memory as property
     *
     * @generated from enum value: SIZE_CONSTRAINT_TYPE_MEMORY = 2;
     */
    MEMORY = 2,
    /**
     * SIZE_CONSTRAINT_TYPE_STORAGE defines the amount of storage as property
     *
     * @generated from enum value: SIZE_CONSTRAINT_TYPE_STORAGE = 3;
     */
    STORAGE = 3,
    /**
     * SIZE_CONSTRAINT_TYPE_GPU defines the number of gpus as property
     *
     * @generated from enum value: SIZE_CONSTRAINT_TYPE_GPU = 4;
     */
    GPU = 4
}
/**
 * Describes the enum metalstack.api.v2.SizeConstraintType.
 */
export declare const SizeConstraintTypeSchema: GenEnum<SizeConstraintType>;
/**
 * SizeService serves size related functions
 *
 * @generated from service metalstack.api.v2.SizeService
 */
export declare const SizeService: GenService<{
    /**
     * Get a size
     *
     * @generated from rpc metalstack.api.v2.SizeService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof SizeServiceGetRequestSchema;
        output: typeof SizeServiceGetResponseSchema;
    };
    /**
     * List all sizes
     *
     * @generated from rpc metalstack.api.v2.SizeService.List
     */
    list: {
        methodKind: "unary";
        input: typeof SizeServiceListRequestSchema;
        output: typeof SizeServiceListResponseSchema;
    };
}>;
