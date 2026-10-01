import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Labels, ProjectRole } from "../../api/v2/common_pb";
import type { Project, ProjectQuery } from "../../api/v2/project_pb";
import type { ProjectMember } from "../../api/v2/project_member_pb";
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
     * Query for projects.
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
     * Projects contains the list of projects.
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
 * ProjectServiceCreateRequest is the request payload to Create a project.
 *
 * @generated from message metalstack.admin.v2.ProjectServiceCreateRequest
 */
export type ProjectServiceCreateRequest = Message<"metalstack.admin.v2.ProjectServiceCreateRequest"> & {
    /**
     * Login is the tenant of this project.
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Name of this project, unique per tenant.
     *
     * @generated from field: string name = 2;
     */
    name: string;
    /**
     * Description of this project.
     *
     * @generated from field: string description = 3;
     */
    description: string;
    /**
     * Avatar URL of the project.
     *
     * @generated from field: optional string avatar_url = 4;
     */
    avatarUrl?: string | undefined;
    /**
     * Labels on the project.
     *
     * @generated from field: metalstack.api.v2.Labels labels = 5;
     */
    labels?: Labels | undefined;
    /**
     * Project is the uuid of the project.
     *
     * @generated from field: optional string project = 6;
     */
    project?: string | undefined;
};
/**
 * Describes the message metalstack.admin.v2.ProjectServiceCreateRequest.
 * Use `create(ProjectServiceCreateRequestSchema)` to create a new message.
 */
export declare const ProjectServiceCreateRequestSchema: GenMessage<ProjectServiceCreateRequest>;
/**
 * ProjectServiceCreateResponse is the response payload for creating a project.
 *
 * @generated from message metalstack.admin.v2.ProjectServiceCreateResponse
 */
export type ProjectServiceCreateResponse = Message<"metalstack.admin.v2.ProjectServiceCreateResponse"> & {
    /**
     * Project contains the created project.
     *
     * @generated from field: metalstack.api.v2.Project project = 1;
     */
    project?: Project | undefined;
};
/**
 * Describes the message metalstack.admin.v2.ProjectServiceCreateResponse.
 * Use `create(ProjectServiceCreateResponseSchema)` to create a new message.
 */
export declare const ProjectServiceCreateResponseSchema: GenMessage<ProjectServiceCreateResponse>;
/**
 * ProjectServiceAddMemberRequest is the request payload for adding a member to a tenant.
 *
 * @generated from message metalstack.admin.v2.ProjectServiceAddMemberRequest
 */
export type ProjectServiceAddMemberRequest = Message<"metalstack.admin.v2.ProjectServiceAddMemberRequest"> & {
    /**
     * Project is the uuid of the project.
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Login of the member to add.
     *
     * @generated from field: string member = 2;
     */
    member: string;
    /**
     * Role to assign to the new member.
     *
     * @generated from field: metalstack.api.v2.ProjectRole role = 3;
     */
    role: ProjectRole;
};
/**
 * Describes the message metalstack.admin.v2.ProjectServiceAddMemberRequest.
 * Use `create(ProjectServiceAddMemberRequestSchema)` to create a new message.
 */
export declare const ProjectServiceAddMemberRequestSchema: GenMessage<ProjectServiceAddMemberRequest>;
/**
 * ProjectServiceAddMemberResponse is the response payload for the add member request.
 *
 * @generated from message metalstack.admin.v2.ProjectServiceAddMemberResponse
 */
export type ProjectServiceAddMemberResponse = Message<"metalstack.admin.v2.ProjectServiceAddMemberResponse"> & {
    /**
     * Member is the added project member.
     *
     * @generated from field: metalstack.api.v2.ProjectMember member = 1;
     */
    member?: ProjectMember | undefined;
};
/**
 * Describes the message metalstack.admin.v2.ProjectServiceAddMemberResponse.
 * Use `create(ProjectServiceAddMemberResponseSchema)` to create a new message.
 */
export declare const ProjectServiceAddMemberResponseSchema: GenMessage<ProjectServiceAddMemberResponse>;
/**
 * ProjectServiceRemoveMemberRequest is the request payload for removing a member from a project.
 *
 * @generated from message metalstack.admin.v2.ProjectServiceRemoveMemberRequest
 */
export type ProjectServiceRemoveMemberRequest = Message<"metalstack.admin.v2.ProjectServiceRemoveMemberRequest"> & {
    /**
     * Project is the uuid of the project.
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Login of the member to remove.
     *
     * @generated from field: string member = 2;
     */
    member: string;
};
/**
 * Describes the message metalstack.admin.v2.ProjectServiceRemoveMemberRequest.
 * Use `create(ProjectServiceRemoveMemberRequestSchema)` to create a new message.
 */
export declare const ProjectServiceRemoveMemberRequestSchema: GenMessage<ProjectServiceRemoveMemberRequest>;
/**
 * ProjectServiceRemoveMemberResponse is the response payload for the remove member request.
 *
 * @generated from message metalstack.admin.v2.ProjectServiceRemoveMemberResponse
 */
export type ProjectServiceRemoveMemberResponse = Message<"metalstack.admin.v2.ProjectServiceRemoveMemberResponse"> & {
    /**
     * Member is the removed project member.
     *
     * @generated from field: metalstack.api.v2.ProjectMember member = 1;
     */
    member?: ProjectMember | undefined;
};
/**
 * Describes the message metalstack.admin.v2.ProjectServiceRemoveMemberResponse.
 * Use `create(ProjectServiceRemoveMemberResponseSchema)` to create a new message.
 */
export declare const ProjectServiceRemoveMemberResponseSchema: GenMessage<ProjectServiceRemoveMemberResponse>;
/**
 * ProjectService provides project management operations.
 *
 * @generated from service metalstack.admin.v2.ProjectService
 */
export declare const ProjectService: GenService<{
    /**
     * Creates a new project.
     *
     * @generated from rpc metalstack.admin.v2.ProjectService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof ProjectServiceCreateRequestSchema;
        output: typeof ProjectServiceCreateResponseSchema;
    };
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
    /**
     * Add a member to a project.
     *
     * @generated from rpc metalstack.admin.v2.ProjectService.AddMember
     */
    addMember: {
        methodKind: "unary";
        input: typeof ProjectServiceAddMemberRequestSchema;
        output: typeof ProjectServiceAddMemberResponseSchema;
    };
    /**
     * RemoveMember removes a member from a project.
     *
     * @generated from rpc metalstack.admin.v2.ProjectService.RemoveMember
     */
    removeMember: {
        methodKind: "unary";
        input: typeof ProjectServiceRemoveMemberRequestSchema;
        output: typeof ProjectServiceRemoveMemberResponseSchema;
    };
}>;
