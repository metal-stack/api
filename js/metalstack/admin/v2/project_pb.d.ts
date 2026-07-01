import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Project, ProjectQuery } from "../../api/v2/project_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/project.proto.
 */
export declare const file_metalstack_admin_v2_project: GenFile;
/**
 * ProjectServiceListRequest is the request payload for listing projects.
 *
 * @generated from message metalstack.admin.v2.ProjectServiceListRequest
 */
export type ProjectServiceListRequest = Message<"metalstack.admin.v2.ProjectServiceListRequest"> & {
    /**
     * Query for projects
     *
     * @generated from field: optional metalstack.api.v2.ProjectQuery query = 1;
     */
    query?: ProjectQuery | undefined;
};
/**
 * Describes the message metalstack.admin.v2.ProjectServiceListRequest.
 * Use `create(ProjectServiceListRequestSchema)` to create a new message.
 */
export declare const ProjectServiceListRequestSchema: GenMessage<ProjectServiceListRequest>;
/**
 * ProjectServiceListResponse is the response payload for listing projects.
 *
 * @generated from message metalstack.admin.v2.ProjectServiceListResponse
 */
export type ProjectServiceListResponse = Message<"metalstack.admin.v2.ProjectServiceListResponse"> & {
    /**
     * Projects contains the list of projects
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
 * ProjectService provides project management operations.
 *
 * @generated from service metalstack.admin.v2.ProjectService
 */
export declare const ProjectService: GenService<{
    /**
     * Returns the list of projects matching the filter criteria.
     *
     * @generated from rpc metalstack.admin.v2.ProjectService.List
     */
    list: {
        methodKind: "unary";
        input: typeof ProjectServiceListRequestSchema;
        output: typeof ProjectServiceListResponseSchema;
    };
}>;
