import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Labels, Meta, ProjectRole, UpdateLabels, UpdateMeta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/project.proto.
 */
export declare const file_metalstack_api_v2_project: GenFile;
/**
 * Project defines a group of resources belonging to a tenant
 * a tenant can have multiple projects
 *
 * @generated from message metalstack.api.v2.Project
 */
export type Project = Message<"metalstack.api.v2.Project"> & {
    /**
     * Uuid of this project
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Meta for this project
     *
     * @generated from field: metalstack.api.v2.Meta meta = 2;
     */
    meta?: Meta;
    /**
     * Name of this project must be unique per tenant
     *
     * @generated from field: string name = 3;
     */
    name: string;
    /**
     * Description of this project
     *
     * @generated from field: string description = 4;
     */
    description: string;
    /**
     * Tenant this project belongs to
     *
     * @generated from field: string tenant = 5;
     */
    tenant: string;
    /**
     * AvatarUrl of the Project
     *
     * @generated from field: optional string avatar_url = 6;
     */
    avatarUrl?: string;
};
/**
 * Describes the message metalstack.api.v2.Project.
 * Use `create(ProjectSchema)` to create a new message.
 */
export declare const ProjectSchema: GenMessage<Project>;
/**
 * ProjectMember defines a user that participates at a project
 *
 * @generated from message metalstack.api.v2.ProjectMember
 */
export type ProjectMember = Message<"metalstack.api.v2.ProjectMember"> & {
    /**
     * Id is the user id of the member
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * Role is the role of the member
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
     * CreatedAt the date when the member was added to the project
     *
     * @generated from field: google.protobuf.Timestamp created_at = 10;
     */
    createdAt?: Timestamp;
};
/**
 * Describes the message metalstack.api.v2.ProjectMember.
 * Use `create(ProjectMemberSchema)` to create a new message.
 */
export declare const ProjectMemberSchema: GenMessage<ProjectMember>;
/**
 * ProjectInvite defines invite to project
 *
 * @generated from message metalstack.api.v2.ProjectInvite
 */
export type ProjectInvite = Message<"metalstack.api.v2.ProjectInvite"> & {
    /**
     * Secret is the secret part of the invite, typically part of the url
     *
     * @generated from field: string secret = 1;
     */
    secret: string;
    /**
     * Project is the project id for which this invite was created
     *
     * @generated from field: string project = 2;
     */
    project: string;
    /**
     * Role is the role in this project the user will get after accepting the invitation
     *
     * @generated from field: metalstack.api.v2.ProjectRole role = 3;
     */
    role: ProjectRole;
    /**
     * Joined is false as long as a user has not accepted the invite
     *
     * @generated from field: bool joined = 4;
     */
    joined: boolean;
    /**
     * ProjectName is the project name for which this invite was created
     *
     * @generated from field: string project_name = 5;
     */
    projectName: string;
    /**
     * Tenant is the login of tenant who invites to join this project
     *
     * @generated from field: string tenant = 6;
     */
    tenant: string;
    /**
     * TenantName is the name of tenant who invites to join this project
     *
     * @generated from field: string tenant_name = 7;
     */
    tenantName: string;
    /**
     * ExpiresAt the date when this invite expires
     *
     * @generated from field: google.protobuf.Timestamp expires_at = 10;
     */
    expiresAt?: Timestamp;
    /**
     * JoinedAt the date when the member accepted this invite
     *
     * @generated from field: google.protobuf.Timestamp joined_at = 11;
     */
    joinedAt?: Timestamp;
};
/**
 * Describes the message metalstack.api.v2.ProjectInvite.
 * Use `create(ProjectInviteSchema)` to create a new message.
 */
export declare const ProjectInviteSchema: GenMessage<ProjectInvite>;
/**
 * ProjectServiceListRequest is the request payload to list all projects
 *
 * @generated from message metalstack.api.v2.ProjectServiceListRequest
 */
export type ProjectServiceListRequest = Message<"metalstack.api.v2.ProjectServiceListRequest"> & {
    /**
     * Id lists only projects with this id
     *
     * @generated from field: optional string id = 1;
     */
    id?: string;
    /**
     * Name lists only projects with this name
     *
     * @generated from field: optional string name = 2;
     */
    name?: string;
    /**
     * Tenant lists only projects of this tenant
     *
     * @generated from field: optional string tenant = 3;
     */
    tenant?: string;
    /**
     * Labels lists only projects containing the given labels
     *
     * @generated from field: optional metalstack.api.v2.Labels labels = 6;
     */
    labels?: Labels;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceListRequest.
 * Use `create(ProjectServiceListRequestSchema)` to create a new message.
 */
export declare const ProjectServiceListRequestSchema: GenMessage<ProjectServiceListRequest>;
/**
 * ProjectServiceListResponse is the response payload to list all projects
 *
 * @generated from message metalstack.api.v2.ProjectServiceListResponse
 */
export type ProjectServiceListResponse = Message<"metalstack.api.v2.ProjectServiceListResponse"> & {
    /**
     * Projects is a list of all your projects
     *
     * @generated from field: repeated metalstack.api.v2.Project projects = 1;
     */
    projects: Project[];
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceListResponse.
 * Use `create(ProjectServiceListResponseSchema)` to create a new message.
 */
export declare const ProjectServiceListResponseSchema: GenMessage<ProjectServiceListResponse>;
/**
 * ProjectServiceGetRequest is the request payload to get a project
 *
 * @generated from message metalstack.api.v2.ProjectServiceGetRequest
 */
export type ProjectServiceGetRequest = Message<"metalstack.api.v2.ProjectServiceGetRequest"> & {
    /**
     * Project is the uuid of the project to get
     *
     * @generated from field: string project = 1;
     */
    project: string;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceGetRequest.
 * Use `create(ProjectServiceGetRequestSchema)` to create a new message.
 */
export declare const ProjectServiceGetRequestSchema: GenMessage<ProjectServiceGetRequest>;
/**
 * ProjectServiceGetResponse is the response payload to get a projects
 *
 * @generated from message metalstack.api.v2.ProjectServiceGetResponse
 */
export type ProjectServiceGetResponse = Message<"metalstack.api.v2.ProjectServiceGetResponse"> & {
    /**
     * Project is the project
     *
     * @generated from field: metalstack.api.v2.Project project = 1;
     */
    project?: Project;
    /**
     * ProjectMembers in this project, projects guests will only see direct project members and not implicit memberships from tenant permissions
     *
     * @generated from field: repeated metalstack.api.v2.ProjectMember project_members = 2;
     */
    projectMembers: ProjectMember[];
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceGetResponse.
 * Use `create(ProjectServiceGetResponseSchema)` to create a new message.
 */
export declare const ProjectServiceGetResponseSchema: GenMessage<ProjectServiceGetResponse>;
/**
 * ProjectServiceCreateRequest is the request payload to Create a project
 *
 * @generated from message metalstack.api.v2.ProjectServiceCreateRequest
 */
export type ProjectServiceCreateRequest = Message<"metalstack.api.v2.ProjectServiceCreateRequest"> & {
    /**
     * Login is the tenant of this project
     * TODO: is login really a good name?
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Name of this project, unique per tenant
     *
     * @generated from field: string name = 2;
     */
    name: string;
    /**
     * Description of this project
     *
     * @generated from field: string description = 3;
     */
    description: string;
    /**
     * Avatar URL of the project
     *
     * @generated from field: optional string avatar_url = 4;
     */
    avatarUrl?: string;
    /**
     * Labels on the project
     *
     * @generated from field: metalstack.api.v2.Labels labels = 5;
     */
    labels?: Labels;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceCreateRequest.
 * Use `create(ProjectServiceCreateRequestSchema)` to create a new message.
 */
export declare const ProjectServiceCreateRequestSchema: GenMessage<ProjectServiceCreateRequest>;
/**
 * ProjectServiceCreateResponse is the response payload of creation of a project
 *
 * @generated from message metalstack.api.v2.ProjectServiceCreateResponse
 */
export type ProjectServiceCreateResponse = Message<"metalstack.api.v2.ProjectServiceCreateResponse"> & {
    /**
     * Project is the project
     *
     * @generated from field: metalstack.api.v2.Project project = 1;
     */
    project?: Project;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceCreateResponse.
 * Use `create(ProjectServiceCreateResponseSchema)` to create a new message.
 */
export declare const ProjectServiceCreateResponseSchema: GenMessage<ProjectServiceCreateResponse>;
/**
 * ProjectServiceDeleteRequest is the request payload to delete a project
 *
 * @generated from message metalstack.api.v2.ProjectServiceDeleteRequest
 */
export type ProjectServiceDeleteRequest = Message<"metalstack.api.v2.ProjectServiceDeleteRequest"> & {
    /**
     * Project is the uuid of the project to get
     *
     * @generated from field: string project = 1;
     */
    project: string;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceDeleteRequest.
 * Use `create(ProjectServiceDeleteRequestSchema)` to create a new message.
 */
export declare const ProjectServiceDeleteRequestSchema: GenMessage<ProjectServiceDeleteRequest>;
/**
 * ProjectServiceDeleteResponse is the response payload to delete a project
 *
 * @generated from message metalstack.api.v2.ProjectServiceDeleteResponse
 */
export type ProjectServiceDeleteResponse = Message<"metalstack.api.v2.ProjectServiceDeleteResponse"> & {
    /**
     * Project is the project
     *
     * @generated from field: metalstack.api.v2.Project project = 1;
     */
    project?: Project;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceDeleteResponse.
 * Use `create(ProjectServiceDeleteResponseSchema)` to create a new message.
 */
export declare const ProjectServiceDeleteResponseSchema: GenMessage<ProjectServiceDeleteResponse>;
/**
 * ProjectServiceUpdateRequest is the request payload to update a project
 *
 * @generated from message metalstack.api.v2.ProjectServiceUpdateRequest
 */
export type ProjectServiceUpdateRequest = Message<"metalstack.api.v2.ProjectServiceUpdateRequest"> & {
    /**
     * Project is the uuid of the project to get
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request
     *
     * @generated from field: metalstack.api.v2.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta;
    /**
     * Name of this project unique per tenant
     *
     * @generated from field: optional string name = 3;
     */
    name?: string;
    /**
     * Description of this project
     *
     * @generated from field: optional string description = 4;
     */
    description?: string;
    /**
     * Avatar URL of the project
     *
     * @generated from field: optional string avatar_url = 5;
     */
    avatarUrl?: string;
    /**
     * Labels on this project
     *
     * @generated from field: optional metalstack.api.v2.UpdateLabels labels = 6;
     */
    labels?: UpdateLabels;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceUpdateRequest.
 * Use `create(ProjectServiceUpdateRequestSchema)` to create a new message.
 */
export declare const ProjectServiceUpdateRequestSchema: GenMessage<ProjectServiceUpdateRequest>;
/**
 * ProjectServiceUpdateResponse is the response payload to update a project
 *
 * @generated from message metalstack.api.v2.ProjectServiceUpdateResponse
 */
export type ProjectServiceUpdateResponse = Message<"metalstack.api.v2.ProjectServiceUpdateResponse"> & {
    /**
     * Project is the project
     *
     * @generated from field: metalstack.api.v2.Project project = 1;
     */
    project?: Project;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceUpdateResponse.
 * Use `create(ProjectServiceUpdateResponseSchema)` to create a new message.
 */
export declare const ProjectServiceUpdateResponseSchema: GenMessage<ProjectServiceUpdateResponse>;
/**
 * ProjectServiceInviteRequest is used to invite a member to a project
 *
 * @generated from message metalstack.api.v2.ProjectServiceInviteRequest
 */
export type ProjectServiceInviteRequest = Message<"metalstack.api.v2.ProjectServiceInviteRequest"> & {
    /**
     * Project is the uuid of the project
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Role of this user in this project
     *
     * @generated from field: metalstack.api.v2.ProjectRole role = 3;
     */
    role: ProjectRole;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceInviteRequest.
 * Use `create(ProjectServiceInviteRequestSchema)` to create a new message.
 */
export declare const ProjectServiceInviteRequestSchema: GenMessage<ProjectServiceInviteRequest>;
/**
 * ProjectServiceInviteRequest is the response payload to a invite member request
 *
 * @generated from message metalstack.api.v2.ProjectServiceInviteResponse
 */
export type ProjectServiceInviteResponse = Message<"metalstack.api.v2.ProjectServiceInviteResponse"> & {
    /**
     * Inviter contains a secret which can be sent to a potential user
     * can be appended to the invitation endpoint at our cloud console like
     * console.metalstack.cloud/invite/<secret>
     *
     * @generated from field: metalstack.api.v2.ProjectInvite invite = 1;
     */
    invite?: ProjectInvite;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceInviteResponse.
 * Use `create(ProjectServiceInviteResponseSchema)` to create a new message.
 */
export declare const ProjectServiceInviteResponseSchema: GenMessage<ProjectServiceInviteResponse>;
/**
 * ProjectServiceInvitesListRequest is the request payload to a list invites request
 *
 * @generated from message metalstack.api.v2.ProjectServiceInvitesListRequest
 */
export type ProjectServiceInvitesListRequest = Message<"metalstack.api.v2.ProjectServiceInvitesListRequest"> & {
    /**
     * Project is the uuid of the project
     *
     * @generated from field: string project = 1;
     */
    project: string;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceInvitesListRequest.
 * Use `create(ProjectServiceInvitesListRequestSchema)` to create a new message.
 */
export declare const ProjectServiceInvitesListRequestSchema: GenMessage<ProjectServiceInvitesListRequest>;
/**
 * ProjectServiceInvitesListResponse is the response payload to a list invites request
 *
 * @generated from message metalstack.api.v2.ProjectServiceInvitesListResponse
 */
export type ProjectServiceInvitesListResponse = Message<"metalstack.api.v2.ProjectServiceInvitesListResponse"> & {
    /**
     * Invites not already accepted the invitation to this project
     *
     * @generated from field: repeated metalstack.api.v2.ProjectInvite invites = 1;
     */
    invites: ProjectInvite[];
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceInvitesListResponse.
 * Use `create(ProjectServiceInvitesListResponseSchema)` to create a new message.
 */
export declare const ProjectServiceInvitesListResponseSchema: GenMessage<ProjectServiceInvitesListResponse>;
/**
 * ProjectServiceInviteGetRequest is the request payload to get a invite
 *
 * @generated from message metalstack.api.v2.ProjectServiceInviteGetRequest
 */
export type ProjectServiceInviteGetRequest = Message<"metalstack.api.v2.ProjectServiceInviteGetRequest"> & {
    /**
     * Secret of the invite to list
     *
     * @generated from field: string secret = 1;
     */
    secret: string;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceInviteGetRequest.
 * Use `create(ProjectServiceInviteGetRequestSchema)` to create a new message.
 */
export declare const ProjectServiceInviteGetRequestSchema: GenMessage<ProjectServiceInviteGetRequest>;
/**
 * ProjectServiceInviteGetResponse is the response payload to a get invite request
 *
 * @generated from message metalstack.api.v2.ProjectServiceInviteGetResponse
 */
export type ProjectServiceInviteGetResponse = Message<"metalstack.api.v2.ProjectServiceInviteGetResponse"> & {
    /**
     * Invite is the invite
     *
     * @generated from field: metalstack.api.v2.ProjectInvite invite = 1;
     */
    invite?: ProjectInvite;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceInviteGetResponse.
 * Use `create(ProjectServiceInviteGetResponseSchema)` to create a new message.
 */
export declare const ProjectServiceInviteGetResponseSchema: GenMessage<ProjectServiceInviteGetResponse>;
/**
 * ProjectServiceLeaveRequest is used to leave a project
 *
 * @generated from message metalstack.api.v2.ProjectServiceLeaveRequest
 */
export type ProjectServiceLeaveRequest = Message<"metalstack.api.v2.ProjectServiceLeaveRequest"> & {
    /**
     * Project is the uuid of the project
     *
     * @generated from field: string project = 1;
     */
    project: string;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceLeaveRequest.
 * Use `create(ProjectServiceLeaveRequestSchema)` to create a new message.
 */
export declare const ProjectServiceLeaveRequestSchema: GenMessage<ProjectServiceLeaveRequest>;
/**
 * ProjectServiceLeaveResponse is the response payload to a leave project request
 *
 * @generated from message metalstack.api.v2.ProjectServiceLeaveResponse
 */
export type ProjectServiceLeaveResponse = Message<"metalstack.api.v2.ProjectServiceLeaveResponse"> & {};
/**
 * Describes the message metalstack.api.v2.ProjectServiceLeaveResponse.
 * Use `create(ProjectServiceLeaveResponseSchema)` to create a new message.
 */
export declare const ProjectServiceLeaveResponseSchema: GenMessage<ProjectServiceLeaveResponse>;
/**
 * ProjectServiceRemoveMemberRequest is used to remove a member from a project
 *
 * @generated from message metalstack.api.v2.ProjectServiceRemoveMemberRequest
 */
export type ProjectServiceRemoveMemberRequest = Message<"metalstack.api.v2.ProjectServiceRemoveMemberRequest"> & {
    /**
     * Project is the uuid of the project
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Member is the id of the member to remove from this project
     *
     * @generated from field: string member = 2;
     */
    member: string;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceRemoveMemberRequest.
 * Use `create(ProjectServiceRemoveMemberRequestSchema)` to create a new message.
 */
export declare const ProjectServiceRemoveMemberRequestSchema: GenMessage<ProjectServiceRemoveMemberRequest>;
/**
 * ProjectServiceRemoveMemberResponse is the response payload to a remove member request
 *
 * @generated from message metalstack.api.v2.ProjectServiceRemoveMemberResponse
 */
export type ProjectServiceRemoveMemberResponse = Message<"metalstack.api.v2.ProjectServiceRemoveMemberResponse"> & {};
/**
 * Describes the message metalstack.api.v2.ProjectServiceRemoveMemberResponse.
 * Use `create(ProjectServiceRemoveMemberResponseSchema)` to create a new message.
 */
export declare const ProjectServiceRemoveMemberResponseSchema: GenMessage<ProjectServiceRemoveMemberResponse>;
/**
 * ProjectServiceUpdateMemberRequest is used to update a member of a project
 *
 * @generated from message metalstack.api.v2.ProjectServiceUpdateMemberRequest
 */
export type ProjectServiceUpdateMemberRequest = Message<"metalstack.api.v2.ProjectServiceUpdateMemberRequest"> & {
    /**
     * Project is the uuid of the project
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Member is the id of the member to remove from this project
     *
     * @generated from field: string member = 2;
     */
    member: string;
    /**
     * Role is the role in this project the user will get after the update
     *
     * @generated from field: metalstack.api.v2.ProjectRole role = 3;
     */
    role: ProjectRole;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceUpdateMemberRequest.
 * Use `create(ProjectServiceUpdateMemberRequestSchema)` to create a new message.
 */
export declare const ProjectServiceUpdateMemberRequestSchema: GenMessage<ProjectServiceUpdateMemberRequest>;
/**
 * ProjectServiceUpdateMemberResponse is the response payload to a update member request
 *
 * @generated from message metalstack.api.v2.ProjectServiceUpdateMemberResponse
 */
export type ProjectServiceUpdateMemberResponse = Message<"metalstack.api.v2.ProjectServiceUpdateMemberResponse"> & {
    /**
     * ProjectMember is the updated project member
     *
     * @generated from field: metalstack.api.v2.ProjectMember project_member = 5;
     */
    projectMember?: ProjectMember;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceUpdateMemberResponse.
 * Use `create(ProjectServiceUpdateMemberResponseSchema)` to create a new message.
 */
export declare const ProjectServiceUpdateMemberResponseSchema: GenMessage<ProjectServiceUpdateMemberResponse>;
/**
 * ProjectServiceInviteAcceptRequest is the request payload to a accept invite request
 *
 * @generated from message metalstack.api.v2.ProjectServiceInviteAcceptRequest
 */
export type ProjectServiceInviteAcceptRequest = Message<"metalstack.api.v2.ProjectServiceInviteAcceptRequest"> & {
    /**
     * Secret is the invitation secret part of the invitation url
     *
     * @generated from field: string secret = 1;
     */
    secret: string;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceInviteAcceptRequest.
 * Use `create(ProjectServiceInviteAcceptRequestSchema)` to create a new message.
 */
export declare const ProjectServiceInviteAcceptRequestSchema: GenMessage<ProjectServiceInviteAcceptRequest>;
/**
 * ProjectServiceInvitesListResponse is the response payload to a accept invite request
 *
 * @generated from message metalstack.api.v2.ProjectServiceInviteAcceptResponse
 */
export type ProjectServiceInviteAcceptResponse = Message<"metalstack.api.v2.ProjectServiceInviteAcceptResponse"> & {
    /**
     * Project ID of the project joined
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * ProjectName if the project joined
     *
     * @generated from field: string project_name = 2;
     */
    projectName: string;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceInviteAcceptResponse.
 * Use `create(ProjectServiceInviteAcceptResponseSchema)` to create a new message.
 */
export declare const ProjectServiceInviteAcceptResponseSchema: GenMessage<ProjectServiceInviteAcceptResponse>;
/**
 * ProjectServiceInviteDeleteRequest is the request payload to a delete invite
 *
 * @generated from message metalstack.api.v2.ProjectServiceInviteDeleteRequest
 */
export type ProjectServiceInviteDeleteRequest = Message<"metalstack.api.v2.ProjectServiceInviteDeleteRequest"> & {
    /**
     * Project is the uuid of the project
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Secret of the invite to delete
     *
     * @generated from field: string secret = 2;
     */
    secret: string;
};
/**
 * Describes the message metalstack.api.v2.ProjectServiceInviteDeleteRequest.
 * Use `create(ProjectServiceInviteDeleteRequestSchema)` to create a new message.
 */
export declare const ProjectServiceInviteDeleteRequestSchema: GenMessage<ProjectServiceInviteDeleteRequest>;
/**
 * ProjectServiceInviteDeleteResponse is the response payload of a delete invite request
 *
 * @generated from message metalstack.api.v2.ProjectServiceInviteDeleteResponse
 */
export type ProjectServiceInviteDeleteResponse = Message<"metalstack.api.v2.ProjectServiceInviteDeleteResponse"> & {};
/**
 * Describes the message metalstack.api.v2.ProjectServiceInviteDeleteResponse.
 * Use `create(ProjectServiceInviteDeleteResponseSchema)` to create a new message.
 */
export declare const ProjectServiceInviteDeleteResponseSchema: GenMessage<ProjectServiceInviteDeleteResponse>;
/**
 * ProjectService serves project related functions
 *
 * @generated from service metalstack.api.v2.ProjectService
 */
export declare const ProjectService: GenService<{
    /**
     * List all accessible projects
     *
     * @generated from rpc metalstack.api.v2.ProjectService.List
     */
    list: {
        methodKind: "unary";
        input: typeof ProjectServiceListRequestSchema;
        output: typeof ProjectServiceListResponseSchema;
    };
    /**
     * Get a project
     *
     * @generated from rpc metalstack.api.v2.ProjectService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof ProjectServiceGetRequestSchema;
        output: typeof ProjectServiceGetResponseSchema;
    };
    /**
     * Create a project
     *
     * @generated from rpc metalstack.api.v2.ProjectService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof ProjectServiceCreateRequestSchema;
        output: typeof ProjectServiceCreateResponseSchema;
    };
    /**
     * Delete a project
     *
     * @generated from rpc metalstack.api.v2.ProjectService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof ProjectServiceDeleteRequestSchema;
        output: typeof ProjectServiceDeleteResponseSchema;
    };
    /**
     * Update a project
     *
     * @generated from rpc metalstack.api.v2.ProjectService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof ProjectServiceUpdateRequestSchema;
        output: typeof ProjectServiceUpdateResponseSchema;
    };
    /**
     * Leave project
     *
     * @generated from rpc metalstack.api.v2.ProjectService.Leave
     */
    leave: {
        methodKind: "unary";
        input: typeof ProjectServiceLeaveRequestSchema;
        output: typeof ProjectServiceLeaveResponseSchema;
    };
    /**
     * RemoveMember remove a user from a project
     *
     * @generated from rpc metalstack.api.v2.ProjectService.RemoveMember
     */
    removeMember: {
        methodKind: "unary";
        input: typeof ProjectServiceRemoveMemberRequestSchema;
        output: typeof ProjectServiceRemoveMemberResponseSchema;
    };
    /**
     * UpdateMember update a user for a project
     *
     * @generated from rpc metalstack.api.v2.ProjectService.UpdateMember
     */
    updateMember: {
        methodKind: "unary";
        input: typeof ProjectServiceUpdateMemberRequestSchema;
        output: typeof ProjectServiceUpdateMemberResponseSchema;
    };
    /**
     * Invite a user to a project
     *
     * @generated from rpc metalstack.api.v2.ProjectService.Invite
     */
    invite: {
        methodKind: "unary";
        input: typeof ProjectServiceInviteRequestSchema;
        output: typeof ProjectServiceInviteResponseSchema;
    };
    /**
     * InviteAccept is called from a user to accept a invitation
     *
     * @generated from rpc metalstack.api.v2.ProjectService.InviteAccept
     */
    inviteAccept: {
        methodKind: "unary";
        input: typeof ProjectServiceInviteAcceptRequestSchema;
        output: typeof ProjectServiceInviteAcceptResponseSchema;
    };
    /**
     * InviteDelete deletes a pending invitation
     *
     * @generated from rpc metalstack.api.v2.ProjectService.InviteDelete
     */
    inviteDelete: {
        methodKind: "unary";
        input: typeof ProjectServiceInviteDeleteRequestSchema;
        output: typeof ProjectServiceInviteDeleteResponseSchema;
    };
    /**
     * InvitesList list all invites to a project
     *
     * @generated from rpc metalstack.api.v2.ProjectService.InvitesList
     */
    invitesList: {
        methodKind: "unary";
        input: typeof ProjectServiceInvitesListRequestSchema;
        output: typeof ProjectServiceInvitesListResponseSchema;
    };
    /**
     * InviteGet get an invite
     *
     * @generated from rpc metalstack.api.v2.ProjectService.InviteGet
     */
    inviteGet: {
        methodKind: "unary";
        input: typeof ProjectServiceInviteGetRequestSchema;
        output: typeof ProjectServiceInviteGetResponseSchema;
    };
}>;
