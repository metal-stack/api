import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/infra/v2/bmc.proto.
 */
export declare const file_metalstack_infra_v2_bmc: GenFile;
/**
 * UpdateBMCInfoRequest
 *
 * @generated from message metalstack.infra.v2.UpdateBMCInfoRequest
 */
export type UpdateBMCInfoRequest = Message<"metalstack.infra.v2.UpdateBMCInfoRequest"> & {};
/**
 * Describes the message metalstack.infra.v2.UpdateBMCInfoRequest.
 * Use `create(UpdateBMCInfoRequestSchema)` to create a new message.
 */
export declare const UpdateBMCInfoRequestSchema: GenMessage<UpdateBMCInfoRequest>;
/**
 * UpdateBMCInfoResponse
 *
 * @generated from message metalstack.infra.v2.UpdateBMCInfoResponse
 */
export type UpdateBMCInfoResponse = Message<"metalstack.infra.v2.UpdateBMCInfoResponse"> & {};
/**
 * Describes the message metalstack.infra.v2.UpdateBMCInfoResponse.
 * Use `create(UpdateBMCInfoResponseSchema)` to create a new message.
 */
export declare const UpdateBMCInfoResponseSchema: GenMessage<UpdateBMCInfoResponse>;
/**
 * BMCService serves bmc related functions
 *
 * @generated from service metalstack.infra.v2.BMCService
 */
export declare const BMCService: GenService<{
    /**
     * UpdateBMCInfo
     *
     * @generated from rpc metalstack.infra.v2.BMCService.UpdateBMCInfo
     */
    updateBMCInfo: {
        methodKind: "unary";
        input: typeof UpdateBMCInfoRequestSchema;
        output: typeof UpdateBMCInfoResponseSchema;
    };
}>;
