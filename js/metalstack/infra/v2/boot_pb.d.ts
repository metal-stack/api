import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { MachineAllocation, MachineBios, MachineBMC, MachineFRU, MachineHardware } from "../../api/v2/machine_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/infra/v2/boot.proto.
 */
export declare const file_metalstack_infra_v2_boot: GenFile;
/**
 * BootServiceDhcpRequest is called once a machine issues a dhcp request
 *
 * @generated from message metalstack.infra.v2.BootServiceDhcpRequest
 */
export type BootServiceDhcpRequest = Message<"metalstack.infra.v2.BootServiceDhcpRequest"> & {
    /**
     * UUID of the machine
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Partition where this machine is located
     *
     * @generated from field: string partition = 2;
     */
    partition: string;
};
/**
 * Describes the message metalstack.infra.v2.BootServiceDhcpRequest.
 * Use `create(BootServiceDhcpRequestSchema)` to create a new message.
 */
export declare const BootServiceDhcpRequestSchema: GenMessage<BootServiceDhcpRequest>;
/**
 * BootServiceDhcpResponse contains the response to a dhcp request
 *
 * @generated from message metalstack.infra.v2.BootServiceDhcpResponse
 */
export type BootServiceDhcpResponse = Message<"metalstack.infra.v2.BootServiceDhcpResponse"> & {};
/**
 * Describes the message metalstack.infra.v2.BootServiceDhcpResponse.
 * Use `create(BootServiceDhcpResponseSchema)` to create a new message.
 */
export declare const BootServiceDhcpResponseSchema: GenMessage<BootServiceDhcpResponse>;
/**
 * BootServiceBootRequest is called to get specified parameters to boot a machine with the given mac
 *
 * @generated from message metalstack.infra.v2.BootServiceBootRequest
 */
export type BootServiceBootRequest = Message<"metalstack.infra.v2.BootServiceBootRequest"> & {
    /**
     * Mac address of the machine
     *
     * @generated from field: string mac = 1;
     */
    mac: string;
    /**
     * Partition where this machine is located
     *
     * @generated from field: string partition = 2;
     */
    partition: string;
};
/**
 * Describes the message metalstack.infra.v2.BootServiceBootRequest.
 * Use `create(BootServiceBootRequestSchema)` to create a new message.
 */
export declare const BootServiceBootRequestSchema: GenMessage<BootServiceBootRequest>;
/**
 * BootServiceBootResponse contains additional infos which are required to boot a machine
 *
 * @generated from message metalstack.infra.v2.BootServiceBootResponse
 */
export type BootServiceBootResponse = Message<"metalstack.infra.v2.BootServiceBootResponse"> & {
    /**
     * Kernel is the url to the linux kernel to boot
     *
     * @generated from field: string kernel = 1;
     */
    kernel: string;
    /**
     * Initial ram disk is the url to the initial ram disk to boot
     *
     * @generated from field: repeated string init_ram_disks = 2;
     */
    initRamDisks: string[];
    /**
     * CMDLine contains kernel command line parameters to boot
     *
     * @generated from field: optional string cmdline = 3;
     */
    cmdline?: string;
};
/**
 * Describes the message metalstack.infra.v2.BootServiceBootResponse.
 * Use `create(BootServiceBootResponseSchema)` to create a new message.
 */
export declare const BootServiceBootResponseSchema: GenMessage<BootServiceBootResponse>;
/**
 * BootServiceRegisterRequest is called from metal-hammer to register a machine with as much hardware details as possible
 *
 * @generated from message metalstack.infra.v2.BootServiceRegisterRequest
 */
export type BootServiceRegisterRequest = Message<"metalstack.infra.v2.BootServiceRegisterRequest"> & {
    /**
     * UUID of this machine
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Hardware details of this machine
     *
     * @generated from field: metalstack.api.v2.MachineHardware hardware = 2;
     */
    hardware?: MachineHardware;
    /**
     * Bios details of this machine
     *
     * @generated from field: metalstack.api.v2.MachineBios bios = 3;
     */
    bios?: MachineBios;
    /**
     * BMC details of this machine
     *
     * @generated from field: metalstack.api.v2.MachineBMC bmc = 4;
     */
    bmc?: MachineBMC;
    /**
     * FRU details of this machine
     *
     * @generated from field: metalstack.api.v2.MachineFRU fru = 5;
     */
    fru?: MachineFRU;
    /**
     * Tags of this machine
     *
     * @generated from field: repeated string tags = 6;
     */
    tags: string[];
    /**
     * MetalHammer version this machine was bootet into
     *
     * @generated from field: string metal_hammer_version = 7;
     */
    metalHammerVersion: string;
    /**
     * Partition where this machine is located
     *
     * @generated from field: string partition = 8;
     */
    partition: string;
};
/**
 * Describes the message metalstack.infra.v2.BootServiceRegisterRequest.
 * Use `create(BootServiceRegisterRequestSchema)` to create a new message.
 */
export declare const BootServiceRegisterRequestSchema: GenMessage<BootServiceRegisterRequest>;
/**
 * BootServiceRegisterResponse response to a BootServiceRegisterResponse request
 *
 * @generated from message metalstack.infra.v2.BootServiceRegisterResponse
 */
export type BootServiceRegisterResponse = Message<"metalstack.infra.v2.BootServiceRegisterResponse"> & {
    /**
     * UUID of this machine
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Size is the calculated size from given hardware details
     *
     * @generated from field: string size = 2;
     */
    size: string;
    /**
     * Partition of this machine
     *
     * @generated from field: string partition = 3;
     */
    partition: string;
};
/**
 * Describes the message metalstack.infra.v2.BootServiceRegisterResponse.
 * Use `create(BootServiceRegisterResponseSchema)` to create a new message.
 */
export declare const BootServiceRegisterResponseSchema: GenMessage<BootServiceRegisterResponse>;
/**
 * BootServiceWaitRequest is called when a machine was registered and is waiting for allocation
 *
 * @generated from message metalstack.infra.v2.BootServiceWaitRequest
 */
export type BootServiceWaitRequest = Message<"metalstack.infra.v2.BootServiceWaitRequest"> & {
    /**
     * UUID of this machine
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message metalstack.infra.v2.BootServiceWaitRequest.
 * Use `create(BootServiceWaitRequestSchema)` to create a new message.
 */
export declare const BootServiceWaitRequestSchema: GenMessage<BootServiceWaitRequest>;
/**
 * BootServiceWaitResponse response to a wait request
 *
 * @generated from message metalstack.infra.v2.BootServiceWaitResponse
 */
export type BootServiceWaitResponse = Message<"metalstack.infra.v2.BootServiceWaitResponse"> & {
    /**
     * Allocation contains the machine.allocation to actually install the machine
     *
     * @generated from field: metalstack.api.v2.MachineAllocation allocation = 1;
     */
    allocation?: MachineAllocation;
};
/**
 * Describes the message metalstack.infra.v2.BootServiceWaitResponse.
 * Use `create(BootServiceWaitResponseSchema)` to create a new message.
 */
export declare const BootServiceWaitResponseSchema: GenMessage<BootServiceWaitResponse>;
/**
 * BootServiceInstallationSucceededRequest is sent from metal-hammer to the api to report the installation succeeded
 *
 * @generated from message metalstack.infra.v2.BootServiceInstallationSucceededRequest
 */
export type BootServiceInstallationSucceededRequest = Message<"metalstack.infra.v2.BootServiceInstallationSucceededRequest"> & {
    /**
     * UUID of the machine to boot
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * ConsolePassword
     *
     * @generated from field: string console_password = 2;
     */
    consolePassword: string;
};
/**
 * Describes the message metalstack.infra.v2.BootServiceInstallationSucceededRequest.
 * Use `create(BootServiceInstallationSucceededRequestSchema)` to create a new message.
 */
export declare const BootServiceInstallationSucceededRequestSchema: GenMessage<BootServiceInstallationSucceededRequest>;
/**
 * BootServiceInstallationSucceededResponse is the response to a BootServiceInstallationSucceededRequest
 *
 * @generated from message metalstack.infra.v2.BootServiceInstallationSucceededResponse
 */
export type BootServiceInstallationSucceededResponse = Message<"metalstack.infra.v2.BootServiceInstallationSucceededResponse"> & {};
/**
 * Describes the message metalstack.infra.v2.BootServiceInstallationSucceededResponse.
 * Use `create(BootServiceInstallationSucceededResponseSchema)` to create a new message.
 */
export declare const BootServiceInstallationSucceededResponseSchema: GenMessage<BootServiceInstallationSucceededResponse>;
/**
 * BootServiceSuperUserPasswordRequest this call returns the password for the machine superuser
 *
 * @generated from message metalstack.infra.v2.BootServiceSuperUserPasswordRequest
 */
export type BootServiceSuperUserPasswordRequest = Message<"metalstack.infra.v2.BootServiceSuperUserPasswordRequest"> & {
    /**
     * UUID of this machine
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message metalstack.infra.v2.BootServiceSuperUserPasswordRequest.
 * Use `create(BootServiceSuperUserPasswordRequestSchema)` to create a new message.
 */
export declare const BootServiceSuperUserPasswordRequestSchema: GenMessage<BootServiceSuperUserPasswordRequest>;
/**
 * BootServiceSuperUserPasswordResponse the super user password is returned
 *
 * @generated from message metalstack.infra.v2.BootServiceSuperUserPasswordResponse
 */
export type BootServiceSuperUserPasswordResponse = Message<"metalstack.infra.v2.BootServiceSuperUserPasswordResponse"> & {
    /**
     * FeatureDisabled on set the superuserpassword in the bmc if this feature is not disabled.
     *
     * @generated from field: bool feature_disabled = 1;
     */
    featureDisabled: boolean;
    /**
     * SuperUserPassword is the password of the superuser on the ipmi device
     *
     * @generated from field: string super_user_password = 2;
     */
    superUserPassword: string;
};
/**
 * Describes the message metalstack.infra.v2.BootServiceSuperUserPasswordResponse.
 * Use `create(BootServiceSuperUserPasswordResponseSchema)` to create a new message.
 */
export declare const BootServiceSuperUserPasswordResponseSchema: GenMessage<BootServiceSuperUserPasswordResponse>;
/**
 * BootService is used for all boot related requests, either pixiecore or metal-hammer
 *
 * Pixiecore
 *
 * @generated from service metalstack.infra.v2.BootService
 */
export declare const BootService: GenService<{
    /**
     * Dhcp is the first dhcp request (option 97). A ProvisioningEventPXEBooting is fired
     *
     * @generated from rpc metalstack.infra.v2.BootService.Dhcp
     */
    dhcp: {
        methodKind: "unary";
        input: typeof BootServiceDhcpRequestSchema;
        output: typeof BootServiceDhcpResponseSchema;
    };
    /**
     * Boot is called from pixie once the machine got the first dhcp response and ipxie asks for subsequent kernel and initrd
     *
     * @generated from rpc metalstack.infra.v2.BootService.Boot
     */
    boot: {
        methodKind: "unary";
        input: typeof BootServiceBootRequestSchema;
        output: typeof BootServiceBootResponseSchema;
    };
    /**
     * SuperUserPassword metal-hammer takes the configured root password for the bmc from metal-apiserver and configure the bmc accordingly
     *
     * @generated from rpc metalstack.infra.v2.BootService.SuperUserPassword
     */
    superUserPassword: {
        methodKind: "unary";
        input: typeof BootServiceSuperUserPasswordRequestSchema;
        output: typeof BootServiceSuperUserPasswordResponseSchema;
    };
    /**
     * Register is called from metal-hammer after hardware inventory is finished, tells metal-apiserver all gory details about that machine
     *
     * @generated from rpc metalstack.infra.v2.BootService.Register
     */
    register: {
        methodKind: "unary";
        input: typeof BootServiceRegisterRequestSchema;
        output: typeof BootServiceRegisterResponseSchema;
    };
    /**
     * Wait is a hanging call that waits until the machine gets allocated by a user
     *
     * @generated from rpc metalstack.infra.v2.BootService.Wait
     */
    wait: {
        methodKind: "server_streaming";
        input: typeof BootServiceWaitRequestSchema;
        output: typeof BootServiceWaitResponseSchema;
    };
    /**
     * InstallationSucceeded tells metal-apiserver installation was either successful
     *
     * @generated from rpc metalstack.infra.v2.BootService.InstallationSucceeded
     */
    installationSucceeded: {
        methodKind: "unary";
        input: typeof BootServiceInstallationSucceededRequestSchema;
        output: typeof BootServiceInstallationSucceededResponseSchema;
    };
}>;
