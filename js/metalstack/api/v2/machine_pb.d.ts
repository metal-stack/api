import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Labels, Meta, UpdateLabels, UpdateMeta } from "./common_pb";
import type { FilesystemLayout } from "./filesystem_pb";
import type { Image } from "./image_pb";
import type { NATType, NetworkType } from "./network_pb";
import type { DNSServer, NTPServer, Partition } from "./partition_pb";
import type { Size } from "./size_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/machine.proto.
 */
export declare const file_metalstack_api_v2_machine: GenFile;
/**
 * MachineServiceGetRequest is the request payload for a machine get request
 *
 * @generated from message metalstack.api.v2.MachineServiceGetRequest
 */
export type MachineServiceGetRequest = Message<"metalstack.api.v2.MachineServiceGetRequest"> & {
    /**
     * UUID of the machine to get
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Project of the machine
     *
     * @generated from field: string project = 2;
     */
    project: string;
};
/**
 * Describes the message metalstack.api.v2.MachineServiceGetRequest.
 * Use `create(MachineServiceGetRequestSchema)` to create a new message.
 */
export declare const MachineServiceGetRequestSchema: GenMessage<MachineServiceGetRequest>;
/**
 * MachineServiceGetResponse is the request payload for a machine get response
 *
 * @generated from message metalstack.api.v2.MachineServiceGetResponse
 */
export type MachineServiceGetResponse = Message<"metalstack.api.v2.MachineServiceGetResponse"> & {
    /**
     * Machine is the machine requested
     *
     * @generated from field: metalstack.api.v2.Machine machine = 1;
     */
    machine?: Machine;
};
/**
 * Describes the message metalstack.api.v2.MachineServiceGetResponse.
 * Use `create(MachineServiceGetResponseSchema)` to create a new message.
 */
export declare const MachineServiceGetResponseSchema: GenMessage<MachineServiceGetResponse>;
/**
 * MachineServiceCreateRequest is the request payload for a machine create request
 *
 * @generated from message metalstack.api.v2.MachineServiceCreateRequest
 */
export type MachineServiceCreateRequest = Message<"metalstack.api.v2.MachineServiceCreateRequest"> & {
    /**
     * Project of the machine
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * UUID if this field is set, this specific machine will be allocated if it is not in available state and not currently allocated.
     * this field overrules size and partition
     *
     * @generated from field: optional string uuid = 2;
     */
    uuid?: string;
    /**
     * Name of this machine
     *
     * @generated from field: string name = 3;
     */
    name: string;
    /**
     * Description for this machine
     *
     * @generated from field: optional string description = 4;
     */
    description?: string;
    /**
     * Hostname the hostname for the allocated machine (defaults to metal)
     *
     * @generated from field: optional string hostname = 5;
     */
    hostname?: string;
    /**
     * Partition the partition id to assign this machine to
     *
     * @generated from field: string partition = 6;
     */
    partition: string;
    /**
     * Size of the machine to create
     *
     * @generated from field: string size = 7;
     */
    size: string;
    /**
     * Image which should be installed on this machine
     *
     * @generated from field: string image = 8;
     */
    image: string;
    /**
     * FilesystemLayout which should be applied for the operating system installation
     * Is defaulted by a lookup at the available fsls for this size and image.
     * Can be specified to test new fsls during development of fsls
     *
     * @generated from field: optional string filesystem_layout = 9;
     */
    filesystemLayout?: string;
    /**
     * SSHPublicKeys defines the ssh public key to be installed on the machine to access it via ssh
     *
     * @generated from field: repeated string ssh_public_keys = 10;
     */
    sshPublicKeys: string[];
    /**
     * Userdata contains instructions required to bootstrap the machine
     * AWS limits the max userdata size to 16k, lets allow twice as much
     *
     * @generated from field: optional string userdata = 11;
     */
    userdata?: string;
    /**
     * Labels to be attached to this machine allocation
     *
     * @generated from field: metalstack.api.v2.Labels labels = 12;
     */
    labels?: Labels;
    /**
     * Networks the networks that this machine will be placed in.
     *
     * @generated from field: repeated metalstack.api.v2.MachineAllocationNetwork networks = 13;
     */
    networks: MachineAllocationNetwork[];
    /**
     * IPs to to attach to this machine additionally
     *
     * @generated from field: repeated string ips = 14;
     */
    ips: string[];
    /**
     * PlacementTags by default machines are spread across the racks inside a partition for every project.
     * if placement tags are provided, the machine candidate has an additional anti-affinity to other machines having the same tags
     *
     * @generated from field: repeated string placement_tags = 15;
     */
    placementTags: string[];
    /**
     * DNSServer the dns servers used for the machine
     *
     * @generated from field: repeated metalstack.api.v2.DNSServer dns_server = 16;
     */
    dnsServer: DNSServer[];
    /**
     * NTPServer the ntp servers used for the machine
     *
     * @generated from field: repeated metalstack.api.v2.NTPServer ntp_server = 17;
     */
    ntpServer: NTPServer[];
    /**
     * AllocationType of this machine
     *
     * @generated from field: metalstack.api.v2.MachineAllocationType allocation_type = 18;
     */
    allocationType: MachineAllocationType;
    /**
     * FirewallSpec provides firewall specific parameters if allocationType is firewall
     *
     * @generated from field: metalstack.api.v2.FirewallSpec firewall_spec = 19;
     */
    firewallSpec?: FirewallSpec;
};
/**
 * Describes the message metalstack.api.v2.MachineServiceCreateRequest.
 * Use `create(MachineServiceCreateRequestSchema)` to create a new message.
 */
export declare const MachineServiceCreateRequestSchema: GenMessage<MachineServiceCreateRequest>;
/**
 * FirewallSpec contains firewall specific firewall creation parameters
 *
 * @generated from message metalstack.api.v2.FirewallSpec
 */
export type FirewallSpec = Message<"metalstack.api.v2.FirewallSpec"> & {
    /**
     * FirewallRules to be applied if this is a firewall
     *
     * @generated from field: metalstack.api.v2.FirewallRules firewall_rules = 19;
     */
    firewallRules?: FirewallRules;
};
/**
 * Describes the message metalstack.api.v2.FirewallSpec.
 * Use `create(FirewallSpecSchema)` to create a new message.
 */
export declare const FirewallSpecSchema: GenMessage<FirewallSpec>;
/**
 * MachineServiceCreateResponse is the request payload for a machine create response
 *
 * @generated from message metalstack.api.v2.MachineServiceCreateResponse
 */
export type MachineServiceCreateResponse = Message<"metalstack.api.v2.MachineServiceCreateResponse"> & {
    /**
     * Machine which was created
     *
     * @generated from field: metalstack.api.v2.Machine machine = 1;
     */
    machine?: Machine;
};
/**
 * Describes the message metalstack.api.v2.MachineServiceCreateResponse.
 * Use `create(MachineServiceCreateResponseSchema)` to create a new message.
 */
export declare const MachineServiceCreateResponseSchema: GenMessage<MachineServiceCreateResponse>;
/**
 * MachineServiceUpdateRequest is the request payload for a machine update request
 *
 * @generated from message metalstack.api.v2.MachineServiceUpdateRequest
 */
export type MachineServiceUpdateRequest = Message<"metalstack.api.v2.MachineServiceUpdateRequest"> & {
    /**
     * UUID of the machine to modify
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request
     *
     * @generated from field: metalstack.api.v2.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta;
    /**
     * Project of the machine
     *
     * @generated from field: string project = 3;
     */
    project: string;
    /**
     * Description of this machine allocation
     *
     * @generated from field: optional string description = 4;
     */
    description?: string;
    /**
     * Labels to update on this machine allocation
     *
     * @generated from field: optional metalstack.api.v2.UpdateLabels labels = 5;
     */
    labels?: UpdateLabels;
    /**
     * SSHPublicKeys which should be update of this machine allocation
     *
     * @generated from field: repeated string ssh_public_keys = 6;
     */
    sshPublicKeys: string[];
};
/**
 * Describes the message metalstack.api.v2.MachineServiceUpdateRequest.
 * Use `create(MachineServiceUpdateRequestSchema)` to create a new message.
 */
export declare const MachineServiceUpdateRequestSchema: GenMessage<MachineServiceUpdateRequest>;
/**
 * MachineServiceUpdateResponse is the request payload for a machine update response
 *
 * @generated from message metalstack.api.v2.MachineServiceUpdateResponse
 */
export type MachineServiceUpdateResponse = Message<"metalstack.api.v2.MachineServiceUpdateResponse"> & {
    /**
     * Machine which was updated
     *
     * @generated from field: metalstack.api.v2.Machine machine = 1;
     */
    machine?: Machine;
};
/**
 * Describes the message metalstack.api.v2.MachineServiceUpdateResponse.
 * Use `create(MachineServiceUpdateResponseSchema)` to create a new message.
 */
export declare const MachineServiceUpdateResponseSchema: GenMessage<MachineServiceUpdateResponse>;
/**
 * MachineServiceListRequest is the request payload for a machine list request
 *
 * @generated from message metalstack.api.v2.MachineServiceListRequest
 */
export type MachineServiceListRequest = Message<"metalstack.api.v2.MachineServiceListRequest"> & {
    /**
     * Project of the machines to list
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Query to list one ore more machines
     *
     * @generated from field: metalstack.api.v2.MachineQuery query = 2;
     */
    query?: MachineQuery;
};
/**
 * Describes the message metalstack.api.v2.MachineServiceListRequest.
 * Use `create(MachineServiceListRequestSchema)` to create a new message.
 */
export declare const MachineServiceListRequestSchema: GenMessage<MachineServiceListRequest>;
/**
 * MachineServiceListResponse is the request payload for a machine list response
 *
 * @generated from message metalstack.api.v2.MachineServiceListResponse
 */
export type MachineServiceListResponse = Message<"metalstack.api.v2.MachineServiceListResponse"> & {
    /**
     * Machines are the machines requested by a list request
     *
     * @generated from field: repeated metalstack.api.v2.Machine machines = 1;
     */
    machines: Machine[];
};
/**
 * Describes the message metalstack.api.v2.MachineServiceListResponse.
 * Use `create(MachineServiceListResponseSchema)` to create a new message.
 */
export declare const MachineServiceListResponseSchema: GenMessage<MachineServiceListResponse>;
/**
 * MachineServiceDeleteRequest is the request payload for a machine delete request
 *
 * @generated from message metalstack.api.v2.MachineServiceDeleteRequest
 */
export type MachineServiceDeleteRequest = Message<"metalstack.api.v2.MachineServiceDeleteRequest"> & {
    /**
     * UUID of the machine to delete
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Project of the machine
     *
     * @generated from field: string project = 2;
     */
    project: string;
};
/**
 * Describes the message metalstack.api.v2.MachineServiceDeleteRequest.
 * Use `create(MachineServiceDeleteRequestSchema)` to create a new message.
 */
export declare const MachineServiceDeleteRequestSchema: GenMessage<MachineServiceDeleteRequest>;
/**
 * MachineServiceDeleteResponse is the request payload for a machine delete response
 *
 * @generated from message metalstack.api.v2.MachineServiceDeleteResponse
 */
export type MachineServiceDeleteResponse = Message<"metalstack.api.v2.MachineServiceDeleteResponse"> & {
    /**
     * Machine which was deleteds
     *
     * @generated from field: metalstack.api.v2.Machine machine = 1;
     */
    machine?: Machine;
};
/**
 * Describes the message metalstack.api.v2.MachineServiceDeleteResponse.
 * Use `create(MachineServiceDeleteResponseSchema)` to create a new message.
 */
export declare const MachineServiceDeleteResponseSchema: GenMessage<MachineServiceDeleteResponse>;
/**
 * Machine represents a physical bare metal machine.
 *
 * @generated from message metalstack.api.v2.Machine
 */
export type Machine = Message<"metalstack.api.v2.Machine"> & {
    /**
     * UUID of this machine
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Meta for this machine
     *
     * @generated from field: metalstack.api.v2.Meta meta = 2;
     */
    meta?: Meta;
    /**
     * Partition where this machine resides
     *
     * @generated from field: metalstack.api.v2.Partition partition = 3;
     */
    partition?: Partition;
    /**
     * Rack where this machine is located
     *
     * @generated from field: string rack = 4;
     */
    rack: string;
    /**
     * Size of this machine
     *
     * @generated from field: metalstack.api.v2.Size size = 5;
     */
    size?: Size;
    /**
     * Hardware specs of this machine
     *
     * @generated from field: metalstack.api.v2.MachineHardware hardware = 6;
     */
    hardware?: MachineHardware;
    /**
     * Bios details for this machine
     *
     * @generated from field: metalstack.api.v2.MachineBios bios = 7;
     */
    bios?: MachineBios;
    /**
     * Allocation details
     *
     * @generated from field: metalstack.api.v2.MachineAllocation allocation = 8;
     */
    allocation?: MachineAllocation;
    /**
     * Status contains several status details related to this machine
     *
     * @generated from field: metalstack.api.v2.MachineStatus status = 9;
     */
    status?: MachineStatus;
    /**
     * MachineRecentProvisioningEvents contains the recent provisioning events
     *
     * @generated from field: metalstack.api.v2.MachineRecentProvisioningEvents recent_provisioning_events = 10;
     */
    recentProvisioningEvents?: MachineRecentProvisioningEvents;
};
/**
 * Describes the message metalstack.api.v2.Machine.
 * Use `create(MachineSchema)` to create a new message.
 */
export declare const MachineSchema: GenMessage<Machine>;
/**
 * MachineStatus contains several status details related to this machine
 *
 * @generated from message metalstack.api.v2.MachineStatus
 */
export type MachineStatus = Message<"metalstack.api.v2.MachineStatus"> & {
    /**
     * Condition describes the availability
     *
     * @generated from field: metalstack.api.v2.MachineCondition condition = 1;
     */
    condition?: MachineCondition;
    /**
     * LEDState indicates the state of the indicator LED on this machine
     *
     * @generated from field: metalstack.api.v2.MachineChassisIdentifyLEDState led_state = 2;
     */
    ledState?: MachineChassisIdentifyLEDState;
    /**
     * Liveliness of this machine
     *
     * @generated from field: metalstack.api.v2.MachineLiveliness liveliness = 3;
     */
    liveliness: MachineLiveliness;
    /**
     * MetalHammerVersion the version of metal hammer which put the machine in waiting state
     *
     * @generated from field: string metal_hammer_version = 4;
     */
    metalHammerVersion: string;
};
/**
 * Describes the message metalstack.api.v2.MachineStatus.
 * Use `create(MachineStatusSchema)` to create a new message.
 */
export declare const MachineStatusSchema: GenMessage<MachineStatus>;
/**
 * MachineCondition describes the availability of this machine
 *
 * @generated from message metalstack.api.v2.MachineCondition
 */
export type MachineCondition = Message<"metalstack.api.v2.MachineCondition"> & {
    /**
     * State the state of this machine. empty means available for all
     *
     * @generated from field: metalstack.api.v2.MachineState state = 1;
     */
    state: MachineState;
    /**
     * Description a description why this machine is in the given state
     *
     * @generated from field: string description = 2;
     */
    description: string;
    /**
     * Issuer the user that changed the state
     *
     * @generated from field: string issuer = 3;
     */
    issuer: string;
};
/**
 * Describes the message metalstack.api.v2.MachineCondition.
 * Use `create(MachineConditionSchema)` to create a new message.
 */
export declare const MachineConditionSchema: GenMessage<MachineCondition>;
/**
 * MachineAllocation contains properties if this machine is allocated
 *
 * @generated from message metalstack.api.v2.MachineAllocation
 */
export type MachineAllocation = Message<"metalstack.api.v2.MachineAllocation"> & {
    /**
     * UUID of this machine allocation
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Meta for this machine allocation
     *
     * @generated from field: metalstack.api.v2.Meta meta = 2;
     */
    meta?: Meta;
    /**
     * Name of this allocation
     *
     * @generated from field: string name = 3;
     */
    name: string;
    /**
     * Description of this allocation
     *
     * @generated from field: string description = 4;
     */
    description: string;
    /**
     * Created By indicates who created this machine allocation
     *
     * @generated from field: string created_by = 5;
     */
    createdBy: string;
    /**
     * Project of the allocation
     *
     * @generated from field: string project = 6;
     */
    project: string;
    /**
     * Image to be used to install on this machine
     *
     * @generated from field: metalstack.api.v2.Image image = 7;
     */
    image?: Image;
    /**
     * FilesystemLayout to create on the disks
     *
     * @generated from field: metalstack.api.v2.FilesystemLayout filesystem_layout = 8;
     */
    filesystemLayout?: FilesystemLayout;
    /**
     * Networks this machine should be attached to
     *
     * @generated from field: repeated metalstack.api.v2.MachineNetwork networks = 9;
     */
    networks: MachineNetwork[];
    /**
     * Hostname of the allocated machine
     *
     * @generated from field: string hostname = 10;
     */
    hostname: string;
    /**
     * SSHPublicKeys which should be installed on this machine
     *
     * @generated from field: repeated string ssh_public_keys = 11;
     */
    sshPublicKeys: string[];
    /**
     * Userdata contains instructions required to bootstrap the machine.
     * AWS limits the max userdata size to 16k, lets allow twice as much
     *
     * @generated from field: string userdata = 12;
     */
    userdata: string;
    /**
     * AllocationType of this machine
     *
     * @generated from field: metalstack.api.v2.MachineAllocationType allocation_type = 13;
     */
    allocationType: MachineAllocationType;
    /**
     * FirewallRules to be applied if this is a firewall
     *
     * @generated from field: metalstack.api.v2.FirewallRules firewall_rules = 14;
     */
    firewallRules?: FirewallRules;
    /**
     * DNSServers for this machine
     *
     * @generated from field: repeated metalstack.api.v2.DNSServer dns_server = 15;
     */
    dnsServer: DNSServer[];
    /**
     * NTPServers for this machine
     *
     * @generated from field: repeated metalstack.api.v2.NTPServer ntp_server = 16;
     */
    ntpServer: NTPServer[];
    /**
     * VPN connection configuration
     *
     * @generated from field: metalstack.api.v2.MachineVPN vpn = 17;
     */
    vpn?: MachineVPN;
};
/**
 * Describes the message metalstack.api.v2.MachineAllocation.
 * Use `create(MachineAllocationSchema)` to create a new message.
 */
export declare const MachineAllocationSchema: GenMessage<MachineAllocation>;
/**
 * MachineAllocationNetwork defines which network should be attached to a machine and if ips should be autoacquired
 *
 * @generated from message metalstack.api.v2.MachineAllocationNetwork
 */
export type MachineAllocationNetwork = Message<"metalstack.api.v2.MachineAllocationNetwork"> & {
    /**
     * Network the id of the network that this machine will be placed in
     *
     * @generated from field: string network = 1;
     */
    network: string;
    /**
     * NoAutoAcquireIp will prevent automatic ip acquirement per network if set to true.
     * By default one ip address is acquired per network for the machine
     *
     * @generated from field: optional bool no_auto_acquire_ip = 2;
     */
    noAutoAcquireIp?: boolean;
};
/**
 * Describes the message metalstack.api.v2.MachineAllocationNetwork.
 * Use `create(MachineAllocationNetworkSchema)` to create a new message.
 */
export declare const MachineAllocationNetworkSchema: GenMessage<MachineAllocationNetwork>;
/**
 * FirewallRules can be defined during firewall allocation
 *
 * @generated from message metalstack.api.v2.FirewallRules
 */
export type FirewallRules = Message<"metalstack.api.v2.FirewallRules"> & {
    /**
     * Egress list of egress rules to be deployed during firewall allocation
     *
     * @generated from field: repeated metalstack.api.v2.FirewallEgressRule egress = 1;
     */
    egress: FirewallEgressRule[];
    /**
     * Ingress list of ingress rules to be deployed during firewall allocation
     *
     * @generated from field: repeated metalstack.api.v2.FirewallIngressRule ingress = 2;
     */
    ingress: FirewallIngressRule[];
};
/**
 * Describes the message metalstack.api.v2.FirewallRules.
 * Use `create(FirewallRulesSchema)` to create a new message.
 */
export declare const FirewallRulesSchema: GenMessage<FirewallRules>;
/**
 * FirewallEgressRule defines rules for outgoing traffic
 *
 * @generated from message metalstack.api.v2.FirewallEgressRule
 */
export type FirewallEgressRule = Message<"metalstack.api.v2.FirewallEgressRule"> & {
    /**
     * Protocol the protocol for the rule, defaults to tcp
     *
     * @generated from field: metalstack.api.v2.IPProtocol protocol = 1;
     */
    protocol: IPProtocol;
    /**
     * Ports the ports affected by this rule
     *
     * @generated from field: repeated uint32 ports = 2;
     */
    ports: number[];
    /**
     * To the destination cidrs affected by this rule
     *
     * @generated from field: repeated string to = 3;
     */
    to: string[];
    /**
     * Comment for this rule
     *
     * @generated from field: string comment = 4;
     */
    comment: string;
};
/**
 * Describes the message metalstack.api.v2.FirewallEgressRule.
 * Use `create(FirewallEgressRuleSchema)` to create a new message.
 */
export declare const FirewallEgressRuleSchema: GenMessage<FirewallEgressRule>;
/**
 * FirewallIngressRule defines rules for incoming traffic
 *
 * @generated from message metalstack.api.v2.FirewallIngressRule
 */
export type FirewallIngressRule = Message<"metalstack.api.v2.FirewallIngressRule"> & {
    /**
     * Protocol the protocol for the rule, defaults to tcp
     *
     * @generated from field: metalstack.api.v2.IPProtocol protocol = 1;
     */
    protocol: IPProtocol;
    /**
     * Ports the ports affected by this rule
     *
     * @generated from field: repeated uint32 ports = 2;
     */
    ports: number[];
    /**
     * To the destination cidrs affected by this rule
     *
     * @generated from field: repeated string to = 3;
     */
    to: string[];
    /**
     * From the source cidrs affected by this rule
     *
     * @generated from field: repeated string from = 4;
     */
    from: string[];
    /**
     * Comment for this rule
     *
     * @generated from field: string comment = 5;
     */
    comment: string;
};
/**
 * Describes the message metalstack.api.v2.FirewallIngressRule.
 * Use `create(FirewallIngressRuleSchema)` to create a new message.
 */
export declare const FirewallIngressRuleSchema: GenMessage<FirewallIngressRule>;
/**
 * MachineNetwork contains details which network should be created on a allocated machine
 *
 * @generated from message metalstack.api.v2.MachineNetwork
 */
export type MachineNetwork = Message<"metalstack.api.v2.MachineNetwork"> & {
    /**
     * Network the networkID of the allocated machine in this vrf
     *
     * @generated from field: string network = 1;
     */
    network: string;
    /**
     * Prefixes the prefixes of this network
     *
     * @generated from field: repeated string prefixes = 2;
     */
    prefixes: string[];
    /**
     * DestinationPrefixes prefixes that are reachable within this network
     *
     * @generated from field: repeated string destination_prefixes = 3;
     */
    destinationPrefixes: string[];
    /**
     * IPs the ip addresses of the allocated machine in this vrf
     *
     * @generated from field: repeated string ips = 4;
     */
    ips: string[];
    /**
     * NetworkType the type of network of this vrf
     *
     * @generated from field: metalstack.api.v2.NetworkType network_type = 5;
     */
    networkType: NetworkType;
    /**
     * NatType what type of nat if any should be used
     *
     * @generated from field: metalstack.api.v2.NATType nat_type = 6;
     */
    natType: NATType;
    /**
     * VRF the vrf id
     *
     * @generated from field: uint64 vrf = 7;
     */
    vrf: bigint;
    /**
     * ASN the autonomous system number for this network
     *
     * @generated from field: uint32 asn = 8;
     */
    asn: number;
};
/**
 * Describes the message metalstack.api.v2.MachineNetwork.
 * Use `create(MachineNetworkSchema)` to create a new message.
 */
export declare const MachineNetworkSchema: GenMessage<MachineNetwork>;
/**
 * MachineHardware contains hardware details
 *
 * @generated from message metalstack.api.v2.MachineHardware
 */
export type MachineHardware = Message<"metalstack.api.v2.MachineHardware"> & {
    /**
     * Memory the total memory of the machine in bytes
     *
     * @generated from field: uint64 memory = 1;
     */
    memory: bigint;
    /**
     * Disks the list of block devices of this machine
     *
     * @generated from field: repeated metalstack.api.v2.MachineBlockDevice disks = 3;
     */
    disks: MachineBlockDevice[];
    /**
     * CPUs the cpu details
     *
     * @generated from field: repeated metalstack.api.v2.MetalCPU cpus = 4;
     */
    cpus: MetalCPU[];
    /**
     * GPUs the gpu details
     *
     * @generated from field: repeated metalstack.api.v2.MetalGPU gpus = 5;
     */
    gpus: MetalGPU[];
    /**
     * Nics the list of network interfaces of this machine
     *
     * @generated from field: repeated metalstack.api.v2.MachineNic nics = 6;
     */
    nics: MachineNic[];
};
/**
 * Describes the message metalstack.api.v2.MachineHardware.
 * Use `create(MachineHardwareSchema)` to create a new message.
 */
export declare const MachineHardwareSchema: GenMessage<MachineHardware>;
/**
 * MetalCPU contains details of a cpu in this machine
 *
 * @generated from message metalstack.api.v2.MetalCPU
 */
export type MetalCPU = Message<"metalstack.api.v2.MetalCPU"> & {
    /**
     * Vendor of this cpu
     *
     * @generated from field: string vendor = 1;
     */
    vendor: string;
    /**
     * Model of this cpu
     *
     * @generated from field: string model = 2;
     */
    model: string;
    /**
     * Cores of this cpu
     *
     * @generated from field: uint32 cores = 3;
     */
    cores: number;
    /**
     * Threads of this cpu
     *
     * @generated from field: uint32 threads = 4;
     */
    threads: number;
};
/**
 * Describes the message metalstack.api.v2.MetalCPU.
 * Use `create(MetalCPUSchema)` to create a new message.
 */
export declare const MetalCPUSchema: GenMessage<MetalCPU>;
/**
 * MetalGPU contains details of a gpu in this machine
 *
 * @generated from message metalstack.api.v2.MetalGPU
 */
export type MetalGPU = Message<"metalstack.api.v2.MetalGPU"> & {
    /**
     * Vendor of this gpu
     *
     * @generated from field: string vendor = 1;
     */
    vendor: string;
    /**
     * Model of this gpu
     *
     * @generated from field: string model = 2;
     */
    model: string;
};
/**
 * Describes the message metalstack.api.v2.MetalGPU.
 * Use `create(MetalGPUSchema)` to create a new message.
 */
export declare const MetalGPUSchema: GenMessage<MetalGPU>;
/**
 * MachineNic contains details of a network interface of this machine
 *
 * @generated from message metalstack.api.v2.MachineNic
 */
export type MachineNic = Message<"metalstack.api.v2.MachineNic"> & {
    /**
     * Mac the macaddress of this interface
     *
     * @generated from field: string mac = 1;
     */
    mac: string;
    /**
     * Name of this interface
     *
     * @generated from field: string name = 2;
     */
    name: string;
    /**
     * Identifier the unique identifier of this network interface
     *
     * @generated from field: string identifier = 3;
     */
    identifier: string;
    /**
     * Vendor of this network card
     *
     * @generated from field: string vendor = 4;
     */
    vendor: string;
    /**
     * Model of this network card
     *
     * @generated from field: string model = 5;
     */
    model: string;
    /**
     * Speed in bits/second of this network card
     *
     * @generated from field: uint64 speed = 6;
     */
    speed: bigint;
    /**
     * Neighbors the neighbors visible to this network interface
     *
     * @generated from field: repeated metalstack.api.v2.MachineNic neighbors = 7;
     */
    neighbors: MachineNic[];
    /**
     * Hostname the nic belongs to
     *
     * @generated from field: string hostname = 8;
     */
    hostname: string;
};
/**
 * Describes the message metalstack.api.v2.MachineNic.
 * Use `create(MachineNicSchema)` to create a new message.
 */
export declare const MachineNicSchema: GenMessage<MachineNic>;
/**
 * MachineBlockDevice contains details of a block device of this machine
 *
 * @generated from message metalstack.api.v2.MachineBlockDevice
 */
export type MachineBlockDevice = Message<"metalstack.api.v2.MachineBlockDevice"> & {
    /**
     * Name of this block device
     *
     * @generated from field: string name = 1;
     */
    name: string;
    /**
     * Size of this block device in bytes
     *
     * @generated from field: uint64 size = 2;
     */
    size: bigint;
};
/**
 * Describes the message metalstack.api.v2.MachineBlockDevice.
 * Use `create(MachineBlockDeviceSchema)` to create a new message.
 */
export declare const MachineBlockDeviceSchema: GenMessage<MachineBlockDevice>;
/**
 * MachineChassisIdentifyLEDState describes the identifier led state
 *
 * @generated from message metalstack.api.v2.MachineChassisIdentifyLEDState
 */
export type MachineChassisIdentifyLEDState = Message<"metalstack.api.v2.MachineChassisIdentifyLEDState"> & {
    /**
     * Value the state of this chassis identify LED. empty means LED-OFF
     *
     * @generated from field: string value = 1;
     */
    value: string;
    /**
     * Description a description why this chassis identify LED is in the given state
     *
     * @generated from field: string description = 2;
     */
    description: string;
};
/**
 * Describes the message metalstack.api.v2.MachineChassisIdentifyLEDState.
 * Use `create(MachineChassisIdentifyLEDStateSchema)` to create a new message.
 */
export declare const MachineChassisIdentifyLEDStateSchema: GenMessage<MachineChassisIdentifyLEDState>;
/**
 * MachineBios contains BIOS details of this machine
 *
 * @generated from message metalstack.api.v2.MachineBios
 */
export type MachineBios = Message<"metalstack.api.v2.MachineBios"> & {
    /**
     * Version the bios version
     *
     * @generated from field: string version = 1;
     */
    version: string;
    /**
     * Vendor the bios vendor
     *
     * @generated from field: string vendor = 2;
     */
    vendor: string;
    /**
     * Date the bios date as string because every vendor has different ideas howto describe the date
     *
     * @generated from field: string date = 3;
     */
    date: string;
};
/**
 * Describes the message metalstack.api.v2.MachineBios.
 * Use `create(MachineBiosSchema)` to create a new message.
 */
export declare const MachineBiosSchema: GenMessage<MachineBios>;
/**
 * MachineRecentProvisioningEvents the recent provisioning events for this machine
 *
 * @generated from message metalstack.api.v2.MachineRecentProvisioningEvents
 */
export type MachineRecentProvisioningEvents = Message<"metalstack.api.v2.MachineRecentProvisioningEvents"> & {
    /**
     * Events the log of recent machine provisioning events
     *
     * @generated from field: repeated metalstack.api.v2.MachineProvisioningEvent events = 1;
     */
    events: MachineProvisioningEvent[];
    /**
     * LastEventTime the time where the last event was received
     *
     * @generated from field: google.protobuf.Timestamp last_event_time = 2;
     */
    lastEventTime?: Timestamp;
    /**
     * LastErrorEvent the last erroneous event received
     *
     * @generated from field: metalstack.api.v2.MachineProvisioningEvent last_error_event = 3;
     */
    lastErrorEvent?: MachineProvisioningEvent;
    /**
     * State can be either CrashLoop, FailedReclaim or something else
     *
     * @generated from field: metalstack.api.v2.MachineProvisioningEventState state = 4;
     */
    state: MachineProvisioningEventState;
};
/**
 * Describes the message metalstack.api.v2.MachineRecentProvisioningEvents.
 * Use `create(MachineRecentProvisioningEventsSchema)` to create a new message.
 */
export declare const MachineRecentProvisioningEventsSchema: GenMessage<MachineRecentProvisioningEvents>;
/**
 * MachineProvisioningEvent is a event which has occurred during provisioning
 *
 * @generated from message metalstack.api.v2.MachineProvisioningEvent
 */
export type MachineProvisioningEvent = Message<"metalstack.api.v2.MachineProvisioningEvent"> & {
    /**
     * Time the time that this event was received
     *
     * @generated from field: google.protobuf.Timestamp time = 1;
     */
    time?: Timestamp;
    /**
     * Event the event emitted by the machine
     *
     * @generated from field: metalstack.api.v2.MachineProvisioningEventType event = 2;
     */
    event: MachineProvisioningEventType;
    /**
     * Message an additional message to add to the event
     *
     * @generated from field: string message = 3;
     */
    message: string;
};
/**
 * Describes the message metalstack.api.v2.MachineProvisioningEvent.
 * Use `create(MachineProvisioningEventSchema)` to create a new message.
 */
export declare const MachineProvisioningEventSchema: GenMessage<MachineProvisioningEvent>;
/**
 * MachineVPN contains configuration data for the VPN connection
 *
 * @generated from message metalstack.api.v2.MachineVPN
 */
export type MachineVPN = Message<"metalstack.api.v2.MachineVPN"> & {
    /**
     * Address of VPN control plane
     *
     * @generated from field: string control_plane_address = 1;
     */
    controlPlaneAddress: string;
    /**
     * Auth key used to connect to VPN
     *
     * @generated from field: string auth_key = 2;
     */
    authKey: string;
    /**
     * Connected indicate if this machine is connected to the VPN
     *
     * TODO add machine ips
     *
     * @generated from field: bool connected = 3;
     */
    connected: boolean;
};
/**
 * Describes the message metalstack.api.v2.MachineVPN.
 * Use `create(MachineVPNSchema)` to create a new message.
 */
export declare const MachineVPNSchema: GenMessage<MachineVPN>;
/**
 * MachineQuery contains fields which can be specified to list specific machines.
 *
 * @generated from message metalstack.api.v2.MachineQuery
 */
export type MachineQuery = Message<"metalstack.api.v2.MachineQuery"> & {
    /**
     * UUID of the machine to get
     *
     * @generated from field: optional string uuid = 1;
     */
    uuid?: string;
    /**
     * Name of the machine to get
     *
     * @generated from field: optional string name = 2;
     */
    name?: string;
    /**
     * Partition of the machine to get
     *
     * @generated from field: optional string partition = 3;
     */
    partition?: string;
    /**
     * Size of the machine to get
     *
     * @generated from field: optional string size = 4;
     */
    size?: string;
    /**
     * Rack of the machine to get
     *
     * @generated from field: optional string rack = 5;
     */
    rack?: string;
    /**
     * Labels for which this machine should get filtered
     *
     * @generated from field: optional metalstack.api.v2.Labels labels = 6;
     */
    labels?: Labels;
    /**
     * Allocation specific machine queries
     *
     * @generated from field: optional metalstack.api.v2.MachineAllocationQuery allocation = 7;
     */
    allocation?: MachineAllocationQuery;
    /**
     * Network specific machine queries
     *
     * @generated from field: optional metalstack.api.v2.MachineNetworkQuery network = 8;
     */
    network?: MachineNetworkQuery;
    /**
     * Nic specific machine queries
     *
     * @generated from field: optional metalstack.api.v2.MachineNicQuery nic = 9;
     */
    nic?: MachineNicQuery;
    /**
     * Disk specific machine queries
     *
     * @generated from field: optional metalstack.api.v2.MachineDiskQuery disk = 10;
     */
    disk?: MachineDiskQuery;
    /**
     * IPMI specific machine queries
     *
     * @generated from field: optional metalstack.api.v2.MachineIPMIQuery ipmi = 11;
     */
    ipmi?: MachineIPMIQuery;
    /**
     * FRU specific machine queries
     *
     * @generated from field: optional metalstack.api.v2.MachineFRUQuery fru = 12;
     */
    fru?: MachineFRUQuery;
    /**
     * Hardware specific machine query
     *
     * @generated from field: optional metalstack.api.v2.MachineHardwareQuery hardware = 13;
     */
    hardware?: MachineHardwareQuery;
    /**
     * State this machine has
     *
     * @generated from field: optional metalstack.api.v2.MachineState state = 14;
     */
    state?: MachineState;
};
/**
 * Describes the message metalstack.api.v2.MachineQuery.
 * Use `create(MachineQuerySchema)` to create a new message.
 */
export declare const MachineQuerySchema: GenMessage<MachineQuery>;
/**
 * MachineAllocationQuery allocation specific query parameters
 *
 * @generated from message metalstack.api.v2.MachineAllocationQuery
 */
export type MachineAllocationQuery = Message<"metalstack.api.v2.MachineAllocationQuery"> & {
    /**
     * UUID of the allocation of the machine to get
     *
     * @generated from field: optional string uuid = 1;
     */
    uuid?: string;
    /**
     * Name of the machine to get
     *
     * @generated from field: optional string name = 2;
     */
    name?: string;
    /**
     * Project of the machine to get
     *
     * @generated from field: optional string project = 3;
     */
    project?: string;
    /**
     * Image of the machine to get
     *
     * @generated from field: optional string image = 4;
     */
    image?: string;
    /**
     * FilesystemLayout of the machine to get
     *
     * @generated from field: optional string filesystem_layout = 5;
     */
    filesystemLayout?: string;
    /**
     * Hostname of the machine to get
     *
     * @generated from field: optional string hostname = 6;
     */
    hostname?: string;
    /**
     * AllocationType of this machine
     *
     * @generated from field: optional metalstack.api.v2.MachineAllocationType allocation_type = 7;
     */
    allocationType?: MachineAllocationType;
    /**
     * Labels for which this machine allocation should get filtered
     *
     * @generated from field: optional metalstack.api.v2.Labels labels = 8;
     */
    labels?: Labels;
};
/**
 * Describes the message metalstack.api.v2.MachineAllocationQuery.
 * Use `create(MachineAllocationQuerySchema)` to create a new message.
 */
export declare const MachineAllocationQuerySchema: GenMessage<MachineAllocationQuery>;
/**
 * MachineNetworkQuery network specific machine queries
 *
 * @generated from message metalstack.api.v2.MachineNetworkQuery
 */
export type MachineNetworkQuery = Message<"metalstack.api.v2.MachineNetworkQuery"> & {
    /**
     * Networks this machine is connected to
     *
     * @generated from field: repeated string networks = 1;
     */
    networks: string[];
    /**
     * Prefixes this machine is connected to
     *
     * @generated from field: repeated string prefixes = 2;
     */
    prefixes: string[];
    /**
     * DestinationPrefixes this machine is connected to
     *
     * @generated from field: repeated string destination_prefixes = 3;
     */
    destinationPrefixes: string[];
    /**
     * IPs this machine has
     *
     * @generated from field: repeated string ips = 4;
     */
    ips: string[];
    /**
     * VRFs this machine is connected to
     *
     * @generated from field: repeated uint64 vrfs = 5;
     */
    vrfs: bigint[];
    /**
     * ASNs this machine is connected to
     *
     * @generated from field: repeated uint32 asns = 6;
     */
    asns: number[];
};
/**
 * Describes the message metalstack.api.v2.MachineNetworkQuery.
 * Use `create(MachineNetworkQuerySchema)` to create a new message.
 */
export declare const MachineNetworkQuerySchema: GenMessage<MachineNetworkQuery>;
/**
 * MachineNicQuery nic specific machine queries
 *
 * @generated from message metalstack.api.v2.MachineNicQuery
 */
export type MachineNicQuery = Message<"metalstack.api.v2.MachineNicQuery"> & {
    /**
     * Macs this machine nic has
     *
     * @generated from field: repeated string macs = 1;
     */
    macs: string[];
    /**
     * Names this machine nic has
     *
     * @generated from field: repeated string names = 2;
     */
    names: string[];
    /**
     * NeighborMacs this machine nic has
     *
     * @generated from field: repeated string neighbor_macs = 3;
     */
    neighborMacs: string[];
    /**
     * NeighborNames this machine nic has
     *
     * @generated from field: repeated string neighbor_names = 4;
     */
    neighborNames: string[];
};
/**
 * Describes the message metalstack.api.v2.MachineNicQuery.
 * Use `create(MachineNicQuerySchema)` to create a new message.
 */
export declare const MachineNicQuerySchema: GenMessage<MachineNicQuery>;
/**
 * MachineDiskQuery disk specific machine queries
 *
 * @generated from message metalstack.api.v2.MachineDiskQuery
 */
export type MachineDiskQuery = Message<"metalstack.api.v2.MachineDiskQuery"> & {
    /**
     * Names of disks in this machine
     *
     * @generated from field: repeated string names = 1;
     */
    names: string[];
    /**
     * Sizes of disks in this machine
     *
     * @generated from field: repeated uint64 sizes = 2;
     */
    sizes: bigint[];
};
/**
 * Describes the message metalstack.api.v2.MachineDiskQuery.
 * Use `create(MachineDiskQuerySchema)` to create a new message.
 */
export declare const MachineDiskQuerySchema: GenMessage<MachineDiskQuery>;
/**
 * MachineIPMIQuery machine ipmi specific machine queries
 *
 * @generated from message metalstack.api.v2.MachineIPMIQuery
 */
export type MachineIPMIQuery = Message<"metalstack.api.v2.MachineIPMIQuery"> & {
    /**
     * Address of the ipmi system of this machine
     *
     * @generated from field: optional string address = 1;
     */
    address?: string;
    /**
     * Mac of the ipmi system of this machine
     *
     * @generated from field: optional string mac = 2;
     */
    mac?: string;
    /**
     * User of the ipmi system of this machine
     *
     * @generated from field: optional string user = 3;
     */
    user?: string;
    /**
     * Interface of the ipmi system of this machine
     *
     * @generated from field: optional string interface = 4;
     */
    interface?: string;
};
/**
 * Describes the message metalstack.api.v2.MachineIPMIQuery.
 * Use `create(MachineIPMIQuerySchema)` to create a new message.
 */
export declare const MachineIPMIQuerySchema: GenMessage<MachineIPMIQuery>;
/**
 * MachineFRUQuery machine fru specific machine queries
 *
 * @generated from message metalstack.api.v2.MachineFRUQuery
 */
export type MachineFRUQuery = Message<"metalstack.api.v2.MachineFRUQuery"> & {
    /**
     * ChassisPartNumber of this machine
     *
     * @generated from field: optional string chassis_part_number = 1;
     */
    chassisPartNumber?: string;
    /**
     * ChassisPartSerial of this machine
     *
     * @generated from field: optional string chassis_part_serial = 2;
     */
    chassisPartSerial?: string;
    /**
     * BoardMFG of this machine
     *
     * @generated from field: optional string board_mfg = 3;
     */
    boardMfg?: string;
    /**
     * BoardSerial of this machine
     *
     * @generated from field: optional string board_serial = 4;
     */
    boardSerial?: string;
    /**
     * BoardPartNumber of this machine
     *
     * @generated from field: optional string board_part_number = 5;
     */
    boardPartNumber?: string;
    /**
     * ProductManufacturer of this machine
     *
     * @generated from field: optional string product_manufacturer = 6;
     */
    productManufacturer?: string;
    /**
     * ProductPartNumber of this machine
     *
     * @generated from field: optional string product_part_number = 7;
     */
    productPartNumber?: string;
    /**
     * ProductSerial of this machine
     *
     * @generated from field: optional string product_serial = 8;
     */
    productSerial?: string;
};
/**
 * Describes the message metalstack.api.v2.MachineFRUQuery.
 * Use `create(MachineFRUQuerySchema)` to create a new message.
 */
export declare const MachineFRUQuerySchema: GenMessage<MachineFRUQuery>;
/**
 * MachineHardwareQuery machine hardware specific machine queries
 *
 * @generated from message metalstack.api.v2.MachineHardwareQuery
 */
export type MachineHardwareQuery = Message<"metalstack.api.v2.MachineHardwareQuery"> & {
    /**
     * Memory the total memory of the machine in bytes
     *
     * @generated from field: optional uint64 memory = 1;
     */
    memory?: bigint;
    /**
     * CPUCores the number of cpu cores
     *
     * @generated from field: optional uint32 cpu_cores = 2;
     */
    cpuCores?: number;
};
/**
 * Describes the message metalstack.api.v2.MachineHardwareQuery.
 * Use `create(MachineHardwareQuerySchema)` to create a new message.
 */
export declare const MachineHardwareQuerySchema: GenMessage<MachineHardwareQuery>;
/**
 * IPProtocol defines tcp|udp
 *
 * @generated from enum metalstack.api.v2.IPProtocol
 */
export declare enum IPProtocol {
    /**
     * IP_PROTOCOL_UNSPECIFIED is not specified
     *
     * @generated from enum value: IP_PROTOCOL_UNSPECIFIED = 0;
     */
    IP_PROTOCOL_UNSPECIFIED = 0,
    /**
     * IP_PROTOCOL_TCP is tcp
     *
     * @generated from enum value: IP_PROTOCOL_TCP = 1;
     */
    IP_PROTOCOL_TCP = 1,
    /**
     * IP_PROTOCOL_UDP is udp
     *
     * @generated from enum value: IP_PROTOCOL_UDP = 2;
     */
    IP_PROTOCOL_UDP = 2
}
/**
 * Describes the enum metalstack.api.v2.IPProtocol.
 */
export declare const IPProtocolSchema: GenEnum<IPProtocol>;
/**
 * MachineState defines if the machine was locked or reserved from a operator
 *
 * @generated from enum metalstack.api.v2.MachineState
 */
export declare enum MachineState {
    /**
     * MACHINE_STATE_UNSPECIFIED is not specified
     *
     * @generated from enum value: MACHINE_STATE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * MACHINE_STATE_RESERVED this machine is reserved
     *
     * @generated from enum value: MACHINE_STATE_RESERVED = 1;
     */
    RESERVED = 1,
    /**
     * MACHINE_STATE_LOCKED this machine is locked
     *
     * @generated from enum value: MACHINE_STATE_LOCKED = 2;
     */
    LOCKED = 2,
    /**
     * MACHINE_STATE_LOCKED this machine is available for all
     *
     * @generated from enum value: MACHINE_STATE_AVAILABLE = 3;
     */
    AVAILABLE = 3
}
/**
 * Describes the enum metalstack.api.v2.MachineState.
 */
export declare const MachineStateSchema: GenEnum<MachineState>;
/**
 * MachineProvisioningEventState possible event states
 *
 * @generated from enum metalstack.api.v2.MachineProvisioningEventState
 */
export declare enum MachineProvisioningEventState {
    /**
     * MACHINE_PROVISIONING_EVENT_STATE_UNSPECIFIED is not specified
     *
     * @generated from enum value: MACHINE_PROVISIONING_EVENT_STATE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * MACHINE_PROVISIONING_EVENT_STATE_CRASHLOOP machine is in crash loop
     *
     * @generated from enum value: MACHINE_PROVISIONING_EVENT_STATE_CRASHLOOP = 1;
     */
    CRASHLOOP = 1,
    /**
     * MACHINE_PROVISIONING_EVENT_STATE_FAILED_RECLAIM machine is in failed reclaim
     *
     * @generated from enum value: MACHINE_PROVISIONING_EVENT_STATE_FAILED_RECLAIM = 2;
     */
    FAILED_RECLAIM = 2
}
/**
 * Describes the enum metalstack.api.v2.MachineProvisioningEventState.
 */
export declare const MachineProvisioningEventStateSchema: GenEnum<MachineProvisioningEventState>;
/**
 * MachineProvisioningEventType defines in which phase the machine actually is
 *
 * @generated from enum metalstack.api.v2.MachineProvisioningEventType
 */
export declare enum MachineProvisioningEventType {
    /**
     * MACHINE_PROVISIONING_EVENT_TYPE_UNSPECIFIED is not specified
     *
     * @generated from enum value: MACHINE_PROVISIONING_EVENT_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * MACHINE_PROVISIONING_EVENT_TYPE_ALIVE machine is alive
     *
     * @generated from enum value: MACHINE_PROVISIONING_EVENT_TYPE_ALIVE = 1;
     */
    ALIVE = 1,
    /**
     * MACHINE_PROVISIONING_EVENT_TYPE_CRASHED machine crashed
     *
     * @generated from enum value: MACHINE_PROVISIONING_EVENT_TYPE_CRASHED = 2;
     */
    CRASHED = 2,
    /**
     * MACHINE_PROVISIONING_EVENT_TYPE_PXE_BOOTING machine is pxe booting into metal-hammer
     *
     * @generated from enum value: MACHINE_PROVISIONING_EVENT_TYPE_PXE_BOOTING = 3;
     */
    PXE_BOOTING = 3,
    /**
     * MACHINE_PROVISIONING_EVENT_TYPE_PLANNED_REBOOT machine got a reboot instruction
     *
     * @generated from enum value: MACHINE_PROVISIONING_EVENT_TYPE_PLANNED_REBOOT = 4;
     */
    PLANNED_REBOOT = 4,
    /**
     * MACHINE_PROVISIONING_EVENT_TYPE_PREPARING metal-hammer is preparing the machine
     *
     * @generated from enum value: MACHINE_PROVISIONING_EVENT_TYPE_PREPARING = 5;
     */
    PREPARING = 5,
    /**
     * MACHINE_PROVISIONING_EVENT_TYPE_REGISTERING metal-hammer registers machine at the apiserver
     *
     * @generated from enum value: MACHINE_PROVISIONING_EVENT_TYPE_REGISTERING = 6;
     */
    REGISTERING = 6,
    /**
     * MACHINE_PROVISIONING_EVENT_TYPE_WAITING machine is waiting for installation
     *
     * @generated from enum value: MACHINE_PROVISIONING_EVENT_TYPE_WAITING = 7;
     */
    WAITING = 7,
    /**
     * MACHINE_PROVISIONING_EVENT_TYPE_INSTALLING metal-hammer is installing the desired os
     *
     * @generated from enum value: MACHINE_PROVISIONING_EVENT_TYPE_INSTALLING = 8;
     */
    INSTALLING = 8,
    /**
     * MACHINE_PROVISIONING_EVENT_TYPE_BOOTING_NEW_KERNEL metal-hammer completed installation and boots into target os
     *
     * @generated from enum value: MACHINE_PROVISIONING_EVENT_TYPE_BOOTING_NEW_KERNEL = 9;
     */
    BOOTING_NEW_KERNEL = 9,
    /**
     * MACHINE_PROVISIONING_EVENT_TYPE_PHONED_HOME machine is installed and phones home
     *
     * @generated from enum value: MACHINE_PROVISIONING_EVENT_TYPE_PHONED_HOME = 10;
     */
    PHONED_HOME = 10,
    /**
     * MACHINE_PROVISIONING_EVENT_TYPE_MACHINE_RECLAIM machine is not allocated, but phones home
     *
     * @generated from enum value: MACHINE_PROVISIONING_EVENT_TYPE_MACHINE_RECLAIM = 11;
     */
    MACHINE_RECLAIM = 11
}
/**
 * Describes the enum metalstack.api.v2.MachineProvisioningEventType.
 */
export declare const MachineProvisioningEventTypeSchema: GenEnum<MachineProvisioningEventType>;
/**
 * MachineLiveliness specifies the liveliness of a machine
 *
 * @generated from enum metalstack.api.v2.MachineLiveliness
 */
export declare enum MachineLiveliness {
    /**
     * MACHINE_LIVELINESS_UNSPECIFIED is not defined
     *
     * @generated from enum value: MACHINE_LIVELINESS_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * MACHINE_LIVELINESS_ALIVE liveliness is alive
     *
     * @generated from enum value: MACHINE_LIVELINESS_ALIVE = 1;
     */
    ALIVE = 1,
    /**
     * MACHINE_LIVELINESS_DEAD liveliness is dead
     *
     * @generated from enum value: MACHINE_LIVELINESS_DEAD = 2;
     */
    DEAD = 2,
    /**
     * MACHINE_LIVELINESS_UNKNOWN liveliness is unknown
     *
     * @generated from enum value: MACHINE_LIVELINESS_UNKNOWN = 3;
     */
    UNKNOWN = 3
}
/**
 * Describes the enum metalstack.api.v2.MachineLiveliness.
 */
export declare const MachineLivelinessSchema: GenEnum<MachineLiveliness>;
/**
 * MachineAllocationType defines if this is a machine or a firewall
 *
 * @generated from enum metalstack.api.v2.MachineAllocationType
 */
export declare enum MachineAllocationType {
    /**
     * MACHINE_ALLOCATION_TYPE_UNSPECIFIED is unspecified
     *
     * @generated from enum value: MACHINE_ALLOCATION_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * MACHINE_ALLOCATION_TYPE_MACHINE is a machine
     *
     * @generated from enum value: MACHINE_ALLOCATION_TYPE_MACHINE = 1;
     */
    MACHINE = 1,
    /**
     * MACHINE_ALLOCATION_TYPE_FIREWALL is a firewall
     *
     * @generated from enum value: MACHINE_ALLOCATION_TYPE_FIREWALL = 2;
     */
    FIREWALL = 2
}
/**
 * Describes the enum metalstack.api.v2.MachineAllocationType.
 */
export declare const MachineAllocationTypeSchema: GenEnum<MachineAllocationType>;
/**
 * MachineService serves machine related functions
 *
 * @generated from service metalstack.api.v2.MachineService
 */
export declare const MachineService: GenService<{
    /**
     * Get a machine
     *
     * @generated from rpc metalstack.api.v2.MachineService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof MachineServiceGetRequestSchema;
        output: typeof MachineServiceGetResponseSchema;
    };
    /**
     * Create a machine
     *
     * @generated from rpc metalstack.api.v2.MachineService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof MachineServiceCreateRequestSchema;
        output: typeof MachineServiceCreateResponseSchema;
    };
    /**
     * Update a machine
     *
     * @generated from rpc metalstack.api.v2.MachineService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof MachineServiceUpdateRequestSchema;
        output: typeof MachineServiceUpdateResponseSchema;
    };
    /**
     * List all machines
     *
     * @generated from rpc metalstack.api.v2.MachineService.List
     */
    list: {
        methodKind: "unary";
        input: typeof MachineServiceListRequestSchema;
        output: typeof MachineServiceListResponseSchema;
    };
    /**
     * Delete a machine
     *
     * @generated from rpc metalstack.api.v2.MachineService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof MachineServiceDeleteRequestSchema;
        output: typeof MachineServiceDeleteResponseSchema;
    };
}>;
