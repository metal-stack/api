# Proto API Findings

Prompt: please identify inconsistencies and logical problems and write them into findings.md

This document captures inconsistencies and logical problems identified across the proto API definitions.

---

## 1. Comment Mismatches with Field / Message Purpose

### C1. partition.proto - "Meta for this ip" on Partition message (line 27)✅

```
// Meta for this ip
Meta meta = 2;
```

The field belongs to the `Partition` message but the comment says "ip". Should be "Meta for this partition".

### C2. partition.proto - "Ip the partition" on PartitionServiceGetResponse (line 87)✅

```
// Ip the partition
Machine machine = 1;
```
Comment says "Ip" instead of "Partition".

### C3. partition.proto - "Ips the partitions" on PartitionServiceListResponse (line 92)✅
```
// Ips the partitions
repeated Machine machines = 1;
```
Comment says "Ips" instead of "Partitions".

### C4. image.proto - Multiple wrong references ("imageLayout" and "ip")✅
- Line 70: `// Id of this imageLayout` should say "this image"
- Line 73: `// Meta for this ip` should say "this image"
- Lines 76-78: `// Name of this imageLayout` and `// Description of this imageLayout` should say "this image"

### C5. project.proto - Delete and Update request comments say "to get"✅
- Line 191 (`ProjectServiceDeleteRequest`, `uuid` field): `// Project is the uuid of the project to get` should say "to delete"
- Line 203 (`ProjectServiceUpdateRequest`, `uuid` field): same comment, should say "to update"

### C6. version.proto comments (lines 22-27) - Validation misuse
All four fields (`version`, `revision`, `git_sha1`, `build_date`) use `is_description` which is semantically wrong. These are structured data (semantic version, revision, git SHA1, build date), not free-form description text.

### C7. machine.proto - `firewall_spec` non-optional with conditional comment
- Line 147-148: `FirewallSpec firewall_spec = 18;` is non-optional, but the comment says `"if allocationType is firewall"`. Since it's not optional, an empty `FirewallSpec` is always sent regardless of `allocation_type`, making the conditional comment misleading.

### C8. response messages say "request payload"✅
All machine response messages incorrectly say "request payload" in their comments:
- `MachineServiceGetResponse` (line 73): "request payload" → "response payload"
- `MachineServiceCreateResponse` (line 158): same
- `MachineServiceUpdateResponse` (line 198): same
- `MachineServiceListResponse` (line 212): same
- `MachineServiceDeleteResponse` (line 226): same
- Same pattern in admin/v2/machine.proto lines 57 and 72

### C9. machine.proto - "deleteds" typo (line 228)✅
```
// Machine which was deleteds
```

### C10. common.proto - "Labels consists labels" (line 125)
```
// Labels consists labels
```

### C11. common.proto - "infra role" vs "machine role" (line 104)✅
```
// MachineRole are used to define which infra role a microservice must provide to call this method
```
Should say "which **machine** role" since the extension is for `MachineRole`.

### C12. project.proto - Wrong type in comment (line 232)✅
```
// ProjectServiceInviteRequest is the response payload to a invite member request
```
Comment says "ProjectServiceInviteRequest" but this is the `InviteResponse` message.

### C13. project.proto - Wrong type in comment (line 306)✅
```
// ProjectServiceInvitesListResponse is the response payload to a accept invite request
```
Comment says "InvitesListResponse" but this is the `InviteAcceptResponse` message.

### C14. tenant.proto - Wrong type in comment (line 292)✅
```
// TenantServiceInvitesListResponse is the response payload to a accept invite request
```
Same issue as C13.

### C15. ip.proto - Awkward comment (lines 121-122)✅
```
// Ip the ip to update
```
Should read "ID of the IP to update" or "IP address to update".

### C16. machine.proto - "overrules" nonstandard word (line 85)✅
```
// this field overrules size and partition.
```
Should be "overrides" or "takes precedence over".

### C17. admin/v2/machine.proto - Grammar errors (line 105)✅
```
// Query to list one ore more bmcs of more machines
```
"ore" → "or", "many machines" likely intended instead of "more machines".

### C18. machine.proto - "store" should be "stored" (lines 112, 178, 332)✅
```
// At most 50 keys can be store.
```
Should be "stored".

### C19. predefined_rules.proto - "macadress" typo (line 11)✅
Missing second "s": should be "macaddress".

---

## 2. Validation Rule Misuse

### V1. machine.proto - BMC fields use `is_description` (lines 279-287)
`MachineBMC` has six fields (`user`, `password`, `interface`, `version`, `power_state`, `address`) all validated with `is_description` (a rule meaning `size() <= 256`). This is wrong:
- `user` (BMC username) should use `is_name` or a dedicated username rule
- `password` should use min/max length validation
- `interface` should use a network interface name validation rule
- `version` should use a version-like pattern validation
- `power_state` should have a reasonable length constraint
- `address` at line 289 already uses `host_and_port` which is correct — only 5 of 6 fields are misvalidated

### V2. machine.proto - BIOS fields use `is_description` (lines 579-581)
`MachineBios` fields `version`, `vendor`, and `date` all use `is_description`. These are BIOS identifiers, not descriptions.

### V3. machine.proto - FRU fields all use `is_description` (lines 593-607)
All eight fields in `MachineFRU` (`chassis_part_number`, `chassis_part_serial`, `board_mfg`, `board_mfg_serial`, `board_part_number`, `product_manufacturer`, `product_part_number`, `product_serial`) use `is_description`. These are FRU identifiers.

### V4. machine.proto - version.proto style issues
Same systematic misuse of `is_description` as in V1 above, applied to BIOS and FRU hardware identifiers that are structured data, not free-form descriptions.

### V5. v2/size_imageconstraint.proto - semver_match uses `is_description` (lines 57-58)
```
string semver_match = 2 [(buf.validate.field).string.(metalstack.api.v2.is_description) = true];
```
A semver match string (e.g., `">= 20.04.20211011"`) is structured data, not a description.

### V6. admin/v2/switch.proto - Wrong rules (lines 100, 102)
- `management_user` uses `is_name` (likely wrong for a username, should use a dedicated username rule)
- `console_command` uses `is_description` (should use command-line validation)

---

## 3. Validation Gaps

### G1. vpn.proto - No validation on VPNNode (lines 15-22)✅
```
message VPNNode {
  uint64 id = 1;
  string name = 2;
  string project = 3;
  repeated string ip_addresses = 4;
  google.protobuf.Timestamp last_seen = 5;
  bool online = 6;
}
```
At minimum, `name` should use `is_name` and `ip_addresses` should be validated for proper IP format.

### G2. partition.proto - Boot commandline has no validation (line 57) ✅
```
string commandline = 3;
```
Kernel command lines can be long and should have max length validation.

---

## 4. Functional Inconsistencies

### F1. BMC query cannot search by hostname
- `MachineBMC.address` (machine.proto line 495): uses `host_and_port` (accepts `<ip>:<port>` or `<hostname>:<port>`)
- `MachineBMCQuery.address` (machine.proto line 507): uses `ip` only (pure IP addresses)

This means you can store and use a BMC with a hostname, but you **cannot query** by that hostname. The query only matches IP addresses, making hostname-based BMC lookup impossible.

### F2. Duplicate `MachineProvisioningEvent` definitions ✅
Two completely separate definitions exist with identical structure:
- `metalstack.api.v2.MachineProvisioningEvent` (machine.proto line 690)
- `metalstack.infra.v2.MachineProvisioningEvent` (infra/v2/event.proto line 33)

Both have enums with identical values ("Alive", "Crashed", "PXE Booting", etc.) but different enum name prefixes. The infra/event.proto does not import machine.proto — it duplicates everything. Risk of drift.

---

## 5. Cross-Version Inconsistencies

### X1. Partition `min_items` missing in admin update
- `api/v2/size_reservation.proto` (SizeReservation): partitions have `min_items: 1` — at least one required
- `admin/v2/size_reservation.proto` (SizeReservationServiceUpdateRequest): partitions at field 5 is **missing** `min_items: 1`, allowing empty partitions in updates

---

## 6. Naming Pattern Inconsistencies in filesystem.proto

### N1. `MatchMachine` message confusing naming
The message `MatchMachine` at line 73 has a field called `filesystem_layout` which is a string parameter, but the oneof case name is `machine_and_filesystemlayout`. The naming is inconsistent — the message name and field name don't clearly communicate what this match operation does.
