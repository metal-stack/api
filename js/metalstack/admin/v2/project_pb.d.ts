import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Labels } from "../../api/v2/common_pb";
import type { Project } from "../../api/v2/project_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/project.proto.
 */
export declare const file_metalstack_admin_v2_project: GenFile;
/**
 * ProjectServiceListRequest is the request payload for the project list request
 *
 * @generated from message metalstack.admin.v2.ProjectServiceListRequest
 */
export type ProjectServiceListRequest = Message<"metalstack.admin.v2.ProjectServiceListRequest"> & {
    /**
     * Tenant lists only projects of this tenant
     *
     * @generated from field: optional string tenant = 1;
     */
    tenant?: string;
    /**
     * Labels lists only projects containing the given labels
     *
     * @generated from field: optional metalstack.api.v2.Labels labels = 2;
     */
    labels?: Labels;
};
/**
 * Describes the message metalstack.admin.v2.ProjectServiceListRequest.
 * Use `create(ProjectServiceListRequestSchema)` to create a new message.
 */
export declare const ProjectServiceListRequestSchema: GenMessage<ProjectServiceListRequest>;
/**
 * ProjectServiceListResponse is the response payload for the project list request
 *
 * @generated from message metalstack.admin.v2.ProjectServiceListResponse
 */
export type ProjectServiceListResponse = Message<"metalstack.admin.v2.ProjectServiceListResponse"> & {
    /**
     * Projects is a list of all projects
     *
     * @generated from field: repeated metalstack.api.v2.Project projects = 1;
     */
    projects: Project[];
};
/**
 * Describes the message metalstack.admin.v2.ProjectServiceListResponse.
 * Use `create(ProjectServiceListResponseSchema)` to create a new message.
 */
export declare const ProjectServiceListResponseSchema: GenMessage<ProjectServiceListResponse>;
/**
 * ProjectService serves project related functions
 *
 * @generated from service metalstack.admin.v2.ProjectService
 */
export declare const ProjectService: GenService<{
    /**
     * List projects based on various filter criteria
     *
     * @generated from rpc metalstack.admin.v2.ProjectService.List
     */
    list: {
        methodKind: "unary";
        input: typeof ProjectServiceListRequestSchema;
        output: typeof ProjectServiceListResponseSchema;
    };
}>;
