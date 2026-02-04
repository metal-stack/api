import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/infra/v2/event.proto.
 */
export declare const file_metalstack_infra_v2_event: GenFile;
/**
 * EventServiceSendRequest.
 *
 * @generated from message metalstack.infra.v2.EventServiceSendRequest
 */
export type EventServiceSendRequest = Message<"metalstack.infra.v2.EventServiceSendRequest"> & {
    /**
     * Events grouped by machine IDs.
     *
     * @generated from field: map<string, metalstack.infra.v2.MachineProvisioningEvent> events = 1;
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
 * EventServiceSendResponse.
 *
 * @generated from message metalstack.infra.v2.EventServiceSendResponse
 */
export type EventServiceSendResponse = Message<"metalstack.infra.v2.EventServiceSendResponse"> & {
    /**
     * Events counts the number of events successfully stored in the database.
     *
     * @generated from field: uint64 events = 1;
     */
    events: bigint;
    /**
     * Failed contains IDs of all machines whose events could not be stored in the database.
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
 * MachineProvisioningEvent contains details about an event.
 *
 * @generated from message metalstack.infra.v2.MachineProvisioningEvent
 */
export type MachineProvisioningEvent = Message<"metalstack.infra.v2.MachineProvisioningEvent"> & {
    /**
     * Time the event occurred at.
     *
     * @generated from field: google.protobuf.Timestamp time = 1;
     */
    time?: Timestamp;
    /**
     * Event that occurred.
     *
     * @generated from field: metalstack.infra.v2.ProvisioningEventType event = 2;
     */
    event: ProvisioningEventType;
    /**
     * Message describing the event in more detail.
     *
     * @generated from field: string message = 3;
     */
    message: string;
};
/**
 * Describes the message metalstack.infra.v2.MachineProvisioningEvent.
 * Use `create(MachineProvisioningEventSchema)` to create a new message.
 */
export declare const MachineProvisioningEventSchema: GenMessage<MachineProvisioningEvent>;
/**
 * ProvisioningEventType is a short description of a machine event.
 *
 * @generated from enum metalstack.infra.v2.ProvisioningEventType
 */
export declare enum ProvisioningEventType {
    /**
     * PROVISIONING_EVENT_TYPE_UNSPECIFIED is unspecified.
     *
     * @generated from enum value: PROVISIONING_EVENT_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * PROVISIONING_EVENT_TYPE_ALIVE means the machine has reported itself to the API not long ago.
     *
     * @generated from enum value: PROVISIONING_EVENT_TYPE_ALIVE = 1;
     */
    ALIVE = 1,
    /**
     * PROVISIONING_EVENT_TYPE_CRASHED means an irregularity in the machine's lifecycle.
     *
     * @generated from enum value: PROVISIONING_EVENT_TYPE_CRASHED = 2;
     */
    CRASHED = 2,
    /**
     * PROVISIONING_EVENT_TYPE_PXE_BOOTING is sent when an unprovisioned machine requests a boot image via PXE.
     *
     * @generated from enum value: PROVISIONING_EVENT_TYPE_PXE_BOOTING = 3;
     */
    PXE_BOOTING = 3,
    /**
     * PROVISIONING_EVENT_TYPE_PLANNED_REBOOT means the machine was scheduled for reboot.
     *
     * @generated from enum value: PROVISIONING_EVENT_TYPE_PLANNED_REBOOT = 4;
     */
    PLANNED_REBOOT = 4,
    /**
     * PROVISIONING_EVENT_TYPE_PREPARING means the metal-hammer has started.
     *
     * @generated from enum value: PROVISIONING_EVENT_TYPE_PREPARING = 5;
     */
    PREPARING = 5,
    /**
     * PROVISIONING_EVENT_TYPE_REGISTERING means the metal-hammer is attempting to register the machine at the API.
     *
     * @generated from enum value: PROVISIONING_EVENT_TYPE_REGISTERING = 6;
     */
    REGISTERING = 6,
    /**
     * PROVISIONING_EVENT_TYPE_WAITING means the machine has successfully reached the state where it is waiting for allocation.
     *
     * @generated from enum value: PROVISIONING_EVENT_TYPE_WAITING = 7;
     */
    WAITING = 7,
    /**
     * PROVISIONING_EVENT_TYPE_INSTALLING means the machine was allocated and the requested OS is being installed.
     *
     * @generated from enum value: PROVISIONING_EVENT_TYPE_INSTALLING = 8;
     */
    INSTALLING = 8,
    /**
     * PROVISIONING_EVENT_TYPE_BOOTING_NEW_KERNEL means the machine has successfully been installed and is now booting into the new OS.
     *
     * @generated from enum value: PROVISIONING_EVENT_TYPE_BOOTING_NEW_KERNEL = 9;
     */
    BOOTING_NEW_KERNEL = 9,
    /**
     * PROVISIONING_EVENT_TYPE_PHONED_HOME is sent periodically by an allocated machine to indicate its liveliness.
     *
     * @generated from enum value: PROVISIONING_EVENT_TYPE_PHONED_HOME = 10;
     */
    PHONED_HOME = 10,
    /**
     * PROVISIONING_EVENT_TYPE_MACHINE_RECLAIM means the machine was freed and is about to return into the pool of waiting machines.
     *
     * @generated from enum value: PROVISIONING_EVENT_TYPE_MACHINE_RECLAIM = 11;
     */
    MACHINE_RECLAIM = 11
}
/**
 * Describes the enum metalstack.infra.v2.ProvisioningEventType.
 */
export declare const ProvisioningEventTypeSchema: GenEnum<ProvisioningEventType>;
/**
 * EventService serves event related functions.
 *
 * @generated from service metalstack.infra.v2.EventService
 */
export declare const EventService: GenService<{
    /**
     * Send a series of machine provisioning events.
     *
     * @generated from rpc metalstack.infra.v2.EventService.Send
     */
    send: {
        methodKind: "unary";
        input: typeof EventServiceSendRequestSchema;
        output: typeof EventServiceSendResponseSchema;
    };
}>;
