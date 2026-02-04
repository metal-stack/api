import type { GenExtension, GenFile } from "@bufbuild/protobuf/codegenv2";
import type { RepeatedRules, StringRules } from "../../../buf/validate/validate_pb";
/**
 * Describes the file metalstack/api/v2/predefined_rules.proto.
 */
export declare const file_metalstack_api_v2_predefined_rules: GenFile;
/**
 * Macaddress returns true if the given string is a valid macadress
 *
 * @generated from extension: optional bool macaddress = 80048951;
 */
export declare const macaddress: GenExtension<StringRules, boolean>;
/**
 * IsName returns true if name field satisfies our requirements
 *
 * @generated from extension: optional bool is_name = 80048952;
 */
export declare const is_name: GenExtension<StringRules, boolean>;
/**
 * IsDescription returns true if description field satisfies our requirements
 *
 * @generated from extension: optional bool is_description = 80048953;
 */
export declare const is_description: GenExtension<StringRules, boolean>;
/**
 * IsPartition returns true if partition field satisfies our requirements
 *
 * @generated from extension: optional bool is_partition = 80048954;
 */
export declare const is_partition: GenExtension<StringRules, boolean>;
/**
 * IsPrefix validates if the given string is a valid prefix
 *
 * @generated from extension: optional bool is_prefix = 80048955;
 */
export declare const is_prefix: GenExtension<StringRules, boolean>;
/**
 * IsUri validates if the given string is a valid URI
 *
 * @generated from extension: optional bool is_uri = 80048956;
 */
export declare const is_uri: GenExtension<StringRules, boolean>;
/**
 * Prefixes validates if a slice of prefixes in string form are valid
 *
 * @generated from extension: optional bool prefixes = 80048956;
 */
export declare const prefixes: GenExtension<RepeatedRules, boolean>;
/**
 * Ips validates if a slice of ips in string form are valid
 *
 * @generated from extension: optional bool ips = 80048957;
 */
export declare const ips: GenExtension<RepeatedRules, boolean>;
