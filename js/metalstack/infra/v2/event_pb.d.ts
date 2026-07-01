import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { MachineProvisioningEvent } from "../../api/v2/machine_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/infra/v2/event.proto.
 */
export declare const file_metalstack_infra_v2_event: GenFile;
/**
 * EventServiceSendRequest is the request payload for sending provisioning events.
 *
 * @generated from message metalstack.infra.v2.EventServiceSendRequest
 */
export type EventServiceSendRequest = Message<"metalstack.infra.v2.EventServiceSendRequest"> & {
    /**
     * Events grouped by machine IDs
     *
     * @generated from field: map<string, metalstack.api.v2.MachineProvisioningEvent> events = 1;
     */
    events: {
        [key: string]: MachineProvisioningEvent;
    };
};
/**
 * Describes the message metalstack.infra.v2.EventServiceSendRequest.
 * Use `create(EventServiceSendRequestSchema)` to create a new message.
 */
export declare const EventServiceSendRequestSchema: GenMessage<EventServiceSendRequest>;
/**
 * EventServiceSendResponse is the response payload for sending provisioning events.
 *
 * @generated from message metalstack.infra.v2.EventServiceSendResponse
 */
export type EventServiceSendResponse = Message<"metalstack.infra.v2.EventServiceSendResponse"> & {
    /**
     * Events counts the number of events successfully stored in the database
     *
     * @generated from field: uint64 events = 1;
     */
    events: bigint;
    /**
     * Failed contains IDs of all machines whose events could not be stored in the database
     *
     * @generated from field: repeated string failed = 2;
     */
    failed: string[];
};
/**
 * Describes the message metalstack.infra.v2.EventServiceSendResponse.
 * Use `create(EventServiceSendResponseSchema)` to create a new message.
 */
export declare const EventServiceSendResponseSchema: GenMessage<EventServiceSendResponse>;
/**
 * EventService provides machine provisioning event logging operations.
 *
 * @generated from service metalstack.infra.v2.EventService
 */
export declare const EventService: GenService<{
    /**
     * Sends a series of machine provisioning events.
     *
     * @generated from rpc metalstack.infra.v2.EventService.Send
     */
    send: {
        methodKind: "unary";
        input: typeof EventServiceSendRequestSchema;
        output: typeof EventServiceSendResponseSchema;
    };
}>;
