import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Component, ComponentQuery } from "../../api/v2/component_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/component.proto.
 */
export declare const file_metalstack_admin_v2_component: GenFile;
/**
 * ComponentServiceListRequest
 *
 * @generated from message metalstack.admin.v2.ComponentServiceListRequest
 */
export type ComponentServiceListRequest = Message<"metalstack.admin.v2.ComponentServiceListRequest"> & {
    /**
     * Query components.
     *
     * @generated from field: metalstack.api.v2.ComponentQuery query = 1;
     */
    query?: ComponentQuery;
};
/**
 * Describes the message metalstack.admin.v2.ComponentServiceListRequest.
 * Use `create(ComponentServiceListRequestSchema)` to create a new message.
 */
export declare const ComponentServiceListRequestSchema: GenMessage<ComponentServiceListRequest>;
/**
 * ComponentServiceListResponse
 *
 * @generated from message metalstack.admin.v2.ComponentServiceListResponse
 */
export type ComponentServiceListResponse = Message<"metalstack.admin.v2.ComponentServiceListResponse"> & {
    /**
     * Components
     *
     * @generated from field: repeated metalstack.api.v2.Component components = 1;
     */
    components: Component[];
};
/**
 * Describes the message metalstack.admin.v2.ComponentServiceListResponse.
 * Use `create(ComponentServiceListResponseSchema)` to create a new message.
 */
export declare const ComponentServiceListResponseSchema: GenMessage<ComponentServiceListResponse>;
/**
 * ComponentServiceGetRequest
 *
 * @generated from message metalstack.admin.v2.ComponentServiceGetRequest
 */
export type ComponentServiceGetRequest = Message<"metalstack.admin.v2.ComponentServiceGetRequest"> & {
    /**
     * UUID identifies the component to get
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message metalstack.admin.v2.ComponentServiceGetRequest.
 * Use `create(ComponentServiceGetRequestSchema)` to create a new message.
 */
export declare const ComponentServiceGetRequestSchema: GenMessage<ComponentServiceGetRequest>;
/**
 * ComponentServiceGetResponse
 *
 * @generated from message metalstack.admin.v2.ComponentServiceGetResponse
 */
export type ComponentServiceGetResponse = Message<"metalstack.admin.v2.ComponentServiceGetResponse"> & {
    /**
     * Component
     *
     * @generated from field: metalstack.api.v2.Component component = 1;
     */
    component?: Component;
};
/**
 * Describes the message metalstack.admin.v2.ComponentServiceGetResponse.
 * Use `create(ComponentServiceGetResponseSchema)` to create a new message.
 */
export declare const ComponentServiceGetResponseSchema: GenMessage<ComponentServiceGetResponse>;
/**
 * ComponentServiceGetRequest
 *
 * @generated from message metalstack.admin.v2.ComponentServiceDeleteRequest
 */
export type ComponentServiceDeleteRequest = Message<"metalstack.admin.v2.ComponentServiceDeleteRequest"> & {
    /**
     * UUID identifies the component to delete
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message metalstack.admin.v2.ComponentServiceDeleteRequest.
 * Use `create(ComponentServiceDeleteRequestSchema)` to create a new message.
 */
export declare const ComponentServiceDeleteRequestSchema: GenMessage<ComponentServiceDeleteRequest>;
/**
 * ComponentServiceGetResponse
 *
 * @generated from message metalstack.admin.v2.ComponentServiceDeleteResponse
 */
export type ComponentServiceDeleteResponse = Message<"metalstack.admin.v2.ComponentServiceDeleteResponse"> & {
    /**
     * Component
     *
     * @generated from field: metalstack.api.v2.Component component = 1;
     */
    component?: Component;
};
/**
 * Describes the message metalstack.admin.v2.ComponentServiceDeleteResponse.
 * Use `create(ComponentServiceDeleteResponseSchema)` to create a new message.
 */
export declare const ComponentServiceDeleteResponseSchema: GenMessage<ComponentServiceDeleteResponse>;
/**
 * ComponentService serves microservice related functions
 *
 * @generated from service metalstack.admin.v2.ComponentService
 */
export declare const ComponentService: GenService<{
    /**
     * Get a single component
     *
     * @generated from rpc metalstack.admin.v2.ComponentService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof ComponentServiceGetRequestSchema;
        output: typeof ComponentServiceGetResponseSchema;
    };
    /**
     * Delete a component
     *
     * @generated from rpc metalstack.admin.v2.ComponentService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof ComponentServiceDeleteRequestSchema;
        output: typeof ComponentServiceDeleteResponseSchema;
    };
    /**
     * List all components with their status
     *
     * @generated from rpc metalstack.admin.v2.ComponentService.List
     */
    list: {
        methodKind: "unary";
        input: typeof ComponentServiceListRequestSchema;
        output: typeof ComponentServiceListResponseSchema;
    };
}>;
