import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv2";
import type { ProjectMember as ProjectMember$1 } from "../../api/v2/project_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/cli/v2/project.proto.
 */
export declare const file_metalstack_cli_v2_project: GenFile;
/**
 * ProjectMember defines a user that participates in a project.
 *
 * @generated from message metalstack.cli.v2.ProjectMember
 */
export type ProjectMember = Message<"metalstack.cli.v2.ProjectMember"> & {
    /**
     * project is the project id to which the member belongs.
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * ProjectMember is the project member.
     *
     * @generated from field: metalstack.api.v2.ProjectMember project_member = 2;
     */
    projectMember?: ProjectMember$1 | undefined;
};
/**
 * Describes the message metalstack.cli.v2.ProjectMember.
 * Use `create(ProjectMemberSchema)` to create a new message.
 */
export declare const ProjectMemberSchema: GenMessage<ProjectMember>;
