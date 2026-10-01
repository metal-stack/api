import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Meta, ProjectRole, UpdateMeta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/project_member.proto.
 */
export declare const file_metalstack_api_v2_project_member: GenFile;
/**
 * ProjectMember defines a user that participates in a project.
 *
 * @generated from message metalstack.api.v2.ProjectMember
 */
export type ProjectMember = Message<"metalstack.api.v2.ProjectMember"> & {
    /**
     * Member is the user id of the member.
     *
     * @generated from field: string member = 1;
     */
    member: string;
    /**
     * Role is the role of the member.
     *
     * @generated from field: metalstack.api.v2.ProjectRole role = 2;
     */
    role: ProjectRole;
    /**
     * InheritedMembership indicates that this member has implicit permissions on the project through his membership within the tenant.
     * This member does not have direct project membership but gains permissions on this project from the role he has in the tenant.
     * Inherited memberships are not included in member lists for users with guest permission but only for direct tenant members.
     *
     * @generated from field: bool inherited_membership = 3;
     */
    inheritedMembership: boolean;
    /**
     * CreatedAt the date when the member was added to the project.
     *
     * @generated from field: google.protobuf.Timestamp created_at = 4;
     */
    createdAt?: Timestamp | undefined;
    /**
     * Meta for this project member.
     *
     * @generated from field: metalstack.api.v2.Meta meta = 5;
     */
    meta?: Meta | undefined;
};
/**
 * Describes the message metalstack.api.v2.ProjectMember.
 * Use `create(ProjectMemberSchema)` to create a new message.
 */
export declare const ProjectMemberSchema: GenMessage<ProjectMember>;
/**
 * ProjectQuery is used to search projects.
 *
 * @generated from message metalstack.api.v2.ProjectMemberQuery
 */
export type ProjectMemberQuery = Message<"metalstack.api.v2.ProjectMemberQuery"> & {
    /**
     * Member is the user id of the member.
     *
     * @generated from field: optional string member = 1;
     */
    member?: string | undefined;
    /**
     * Role is the role of the member.
     *
     * @generated from field: optional metalstack.api.v2.ProjectRole role = 2;
     */
    role?: ProjectRole | undefined;
};
/**
 * Describes the message metalstack.api.v2.ProjectMemberQuery.
 * Use `create(ProjectMemberQuerySchema)` to create a new message.
 */
export declare const ProjectMemberQuerySchema: GenMessage<ProjectMemberQuery>;
/**
 * ProjectMemberServiceListRequest is the request payload to list all project members.
 *
 * @generated from message metalstack.api.v2.ProjectMemberServiceListRequest
 */
export type ProjectMemberServiceListRequest = Message<"metalstack.api.v2.ProjectMemberServiceListRequest"> & {
    /**
     * Project is the uuid of the project.
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Query for projects.
     *
     * @generated from field: metalstack.api.v2.ProjectMemberQuery query = 2;
     */
    query?: ProjectMemberQuery | undefined;
};
/**
 * Describes the message metalstack.api.v2.ProjectMemberServiceListRequest.
 * Use `create(ProjectMemberServiceListRequestSchema)` to create a new message.
 */
export declare const ProjectMemberServiceListRequestSchema: GenMessage<ProjectMemberServiceListRequest>;
/**
 * ProjectMemberServiceGetRequest is the request payload to get a project.
 *
 * @generated from message metalstack.api.v2.ProjectMemberServiceGetRequest
 */
export type ProjectMemberServiceGetRequest = Message<"metalstack.api.v2.ProjectMemberServiceGetRequest"> & {
    /**
     * Project is the uuid of the project.
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Member is the user id of the member.
     *
     * @generated from field: string member = 2;
     */
    member: string;
};
/**
 * Describes the message metalstack.api.v2.ProjectMemberServiceGetRequest.
 * Use `create(ProjectMemberServiceGetRequestSchema)` to create a new message.
 */
export declare const ProjectMemberServiceGetRequestSchema: GenMessage<ProjectMemberServiceGetRequest>;
/**
 * ProjectMemberServiceCreateRequest is the request payload to Create a project.
 *
 * @generated from message metalstack.api.v2.ProjectMemberServiceCreateRequest
 */
export type ProjectMemberServiceCreateRequest = Message<"metalstack.api.v2.ProjectMemberServiceCreateRequest"> & {
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
 * Describes the message metalstack.api.v2.ProjectMemberServiceCreateRequest.
 * Use `create(ProjectMemberServiceCreateRequestSchema)` to create a new message.
 */
export declare const ProjectMemberServiceCreateRequestSchema: GenMessage<ProjectMemberServiceCreateRequest>;
/**
 * ProjectMemberServiceUpdateRequest is the request payload to update a project.
 *
 * @generated from message metalstack.api.v2.ProjectMemberServiceUpdateRequest
 */
export type ProjectMemberServiceUpdateRequest = Message<"metalstack.api.v2.ProjectMemberServiceUpdateRequest"> & {
    /**
     * Project is the uuid of the project.
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Member is the id of the member to remove from this project.
     *
     * @generated from field: string member = 2;
     */
    member: string;
    /**
     * Role is the role in this project the user will get after the update.
     *
     * @generated from field: metalstack.api.v2.ProjectRole role = 3;
     */
    role: ProjectRole;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request.
     *
     * @generated from field: metalstack.api.v2.UpdateMeta update_meta = 4;
     */
    updateMeta?: UpdateMeta | undefined;
};
/**
 * Describes the message metalstack.api.v2.ProjectMemberServiceUpdateRequest.
 * Use `create(ProjectMemberServiceUpdateRequestSchema)` to create a new message.
 */
export declare const ProjectMemberServiceUpdateRequestSchema: GenMessage<ProjectMemberServiceUpdateRequest>;
/**
 * ProjectMemberServiceDeleteRequest is the request payload to delete a project.
 *
 * @generated from message metalstack.api.v2.ProjectMemberServiceDeleteRequest
 */
export type ProjectMemberServiceDeleteRequest = Message<"metalstack.api.v2.ProjectMemberServiceDeleteRequest"> & {
    /**
     * Project is the uuid of the project.
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Member is the id of the member to remove from this project.
     *
     * @generated from field: string member = 2;
     */
    member: string;
};
/**
 * Describes the message metalstack.api.v2.ProjectMemberServiceDeleteRequest.
 * Use `create(ProjectMemberServiceDeleteRequestSchema)` to create a new message.
 */
export declare const ProjectMemberServiceDeleteRequestSchema: GenMessage<ProjectMemberServiceDeleteRequest>;
/**
 * ProjectMemberServiceListResponse is the response payload to list all project members.
 *
 * @generated from message metalstack.api.v2.ProjectMemberServiceListResponse
 */
export type ProjectMemberServiceListResponse = Message<"metalstack.api.v2.ProjectMemberServiceListResponse"> & {
    /**
     * Members is a list of project members.
     *
     * @generated from field: repeated metalstack.api.v2.ProjectMember members = 1;
     */
    members: ProjectMember[];
};
/**
 * Describes the message metalstack.api.v2.ProjectMemberServiceListResponse.
 * Use `create(ProjectMemberServiceListResponseSchema)` to create a new message.
 */
export declare const ProjectMemberServiceListResponseSchema: GenMessage<ProjectMemberServiceListResponse>;
/**
 * ProjectMemberServiceGetResponse is the response payload to get a projects.
 *
 * @generated from message metalstack.api.v2.ProjectMemberServiceGetResponse
 */
export type ProjectMemberServiceGetResponse = Message<"metalstack.api.v2.ProjectMemberServiceGetResponse"> & {
    /**
     * Member in this project, projects guests will only see direct project members and not implicit memberships from tenant permissions.
     *
     * @generated from field: repeated metalstack.api.v2.ProjectMember member = 1;
     */
    member: ProjectMember[];
};
/**
 * Describes the message metalstack.api.v2.ProjectMemberServiceGetResponse.
 * Use `create(ProjectMemberServiceGetResponseSchema)` to create a new message.
 */
export declare const ProjectMemberServiceGetResponseSchema: GenMessage<ProjectMemberServiceGetResponse>;
/**
 * ProjectMemberServiceCreateResponse is the response payload to create a project member.
 *
 * @generated from message metalstack.api.v2.ProjectMemberServiceCreateResponse
 */
export type ProjectMemberServiceCreateResponse = Message<"metalstack.api.v2.ProjectMemberServiceCreateResponse"> & {
    /**
     * Member is the created project member
     *
     * @generated from field: metalstack.api.v2.ProjectMember member = 1;
     */
    member?: ProjectMember | undefined;
};
/**
 * Describes the message metalstack.api.v2.ProjectMemberServiceCreateResponse.
 * Use `create(ProjectMemberServiceCreateResponseSchema)` to create a new message.
 */
export declare const ProjectMemberServiceCreateResponseSchema: GenMessage<ProjectMemberServiceCreateResponse>;
/**
 * ProjectMemberServiceDeleteResponse is the response payload to delete a project member.
 *
 * @generated from message metalstack.api.v2.ProjectMemberServiceDeleteResponse
 */
export type ProjectMemberServiceDeleteResponse = Message<"metalstack.api.v2.ProjectMemberServiceDeleteResponse"> & {
    /**
     * Member is the deleted project member
     *
     * @generated from field: metalstack.api.v2.ProjectMember member = 1;
     */
    member?: ProjectMember | undefined;
};
/**
 * Describes the message metalstack.api.v2.ProjectMemberServiceDeleteResponse.
 * Use `create(ProjectMemberServiceDeleteResponseSchema)` to create a new message.
 */
export declare const ProjectMemberServiceDeleteResponseSchema: GenMessage<ProjectMemberServiceDeleteResponse>;
/**
 * ProjectMemberServiceUpdateResponse is the response payload to update a project member.
 *
 * @generated from message metalstack.api.v2.ProjectMemberServiceUpdateResponse
 */
export type ProjectMemberServiceUpdateResponse = Message<"metalstack.api.v2.ProjectMemberServiceUpdateResponse"> & {
    /**
     * Member is the updated project member
     *
     * @generated from field: metalstack.api.v2.ProjectMember member = 1;
     */
    member?: ProjectMember | undefined;
};
/**
 * Describes the message metalstack.api.v2.ProjectMemberServiceUpdateResponse.
 * Use `create(ProjectMemberServiceUpdateResponseSchema)` to create a new message.
 */
export declare const ProjectMemberServiceUpdateResponseSchema: GenMessage<ProjectMemberServiceUpdateResponse>;
/**
 * ProjectMemberServiceLeaveRequest is used to leave a project. This way a member is not required to ask a project owner to remove him.
 *
 * @generated from message metalstack.api.v2.ProjectMemberServiceLeaveRequest
 */
export type ProjectMemberServiceLeaveRequest = Message<"metalstack.api.v2.ProjectMemberServiceLeaveRequest"> & {
    /**
     * Project of the uuid of the project that the user wants to leave.
     *
     * @generated from field: string project = 1;
     */
    project: string;
};
/**
 * Describes the message metalstack.api.v2.ProjectMemberServiceLeaveRequest.
 * Use `create(ProjectMemberServiceLeaveRequestSchema)` to create a new message.
 */
export declare const ProjectMemberServiceLeaveRequestSchema: GenMessage<ProjectMemberServiceLeaveRequest>;
/**
 * ProjectMemberServiceLeaveResponse is the response payload to a leave project request.
 *
 * @generated from message metalstack.api.v2.ProjectMemberServiceLeaveResponse
 */
export type ProjectMemberServiceLeaveResponse = Message<"metalstack.api.v2.ProjectMemberServiceLeaveResponse"> & {
    /**
     * Member is the member of this project.
     *
     * @generated from field: metalstack.api.v2.ProjectMember member = 1;
     */
    member?: ProjectMember | undefined;
};
/**
 * Describes the message metalstack.api.v2.ProjectMemberServiceLeaveResponse.
 * Use `create(ProjectMemberServiceLeaveResponseSchema)` to create a new message.
 */
export declare const ProjectMemberServiceLeaveResponseSchema: GenMessage<ProjectMemberServiceLeaveResponse>;
/**
 * ProjectMemberService provides project management operations.
 *
 * @generated from service metalstack.api.v2.ProjectMemberService
 */
export declare const ProjectMemberService: GenService<{
    /**
     * Returns the list of project members.
     *
     * @generated from rpc metalstack.api.v2.ProjectMemberService.List
     */
    list: {
        methodKind: "unary";
        input: typeof ProjectMemberServiceListRequestSchema;
        output: typeof ProjectMemberServiceListResponseSchema;
    };
    /**
     * Returns the project member of the project.
     *
     * @generated from rpc metalstack.api.v2.ProjectMemberService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof ProjectMemberServiceGetRequestSchema;
        output: typeof ProjectMemberServiceGetResponseSchema;
    };
    /**
     * Create a project member.
     *
     * @generated from rpc metalstack.api.v2.ProjectMemberService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof ProjectMemberServiceCreateRequestSchema;
        output: typeof ProjectMemberServiceCreateResponseSchema;
    };
    /**
     * Delete a project member.
     *
     * @generated from rpc metalstack.api.v2.ProjectMemberService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof ProjectMemberServiceDeleteRequestSchema;
        output: typeof ProjectMemberServiceDeleteResponseSchema;
    };
    /**
     * Update a project member.
     *
     * @generated from rpc metalstack.api.v2.ProjectMemberService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof ProjectMemberServiceUpdateRequestSchema;
        output: typeof ProjectMemberServiceUpdateResponseSchema;
    };
    /**
     * Leave a project.
     *
     * @generated from rpc metalstack.api.v2.ProjectMemberService.Leave
     */
    leave: {
        methodKind: "unary";
        input: typeof ProjectMemberServiceLeaveRequestSchema;
        output: typeof ProjectMemberServiceLeaveResponseSchema;
    };
}>;
