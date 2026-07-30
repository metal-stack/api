# Domain Model Review

A review of the metal-stack API domain model, identifying weaknesses and potential improvements.

## Overview

The API follows a clean three-layer architecture (public / admin / infra) with well-defined auth scopes, consistent CRUD patterns, strong validation via buf/CEL, and optimistic locking. The machine lifecycle (PXE → register → wait → install → phone home) is thoroughly modeled. Below are the weaknesses found.

---

## 1. `User` vs `Tenant` Identity Ambiguity

**File:** `proto/metalstack/api/v2/user.proto:21-36`, `proto/metalstack/api/v2/project.proto:187`

`User` has `login`, `tenants`, `projects`; `Tenant` also has `login`. The `ProjectServiceCreateRequest` field `login` has a TODO comment: _"is login really a good name?"_ The `Tenant.created_by` field is typed as `string` with no validation, linking back to a user concept that is never formally defined as a message type in a request/response — users are only materialized as `User` in `UserServiceGetResponse`.

This naming confusion between user/tenant login persists across the model. A human "user" and an organizational "tenant" share the same identifier namespace and validation rules (`is_tenant_login`).

## 2. `MachineAllocation` creates a large embedded optional chunk

**File:** `proto/metalstack/api/v2/machine.proto:250-348`

`Machine` embeds an `allocation` field of type `MachineAllocation` which contains: `uuid`, `name`, `description`, `meta`, `created_by`, `project`, `image`, `filesystem`, `networks`, `hostname`, `ssh_public_keys`, `userdata`, `allocation_type`, `firewall_rules`, `dns_servers`, `ntp_servers`, `vpn`. Almost all of these are only meaningful when the machine _is_ allocated.

When a machine is free (unallocated), half the `Machine` message is structurally empty. This is a large embedded optional value that is always serialized, and consumers must null-check `allocation` to know the machine's lifecycle phase.

## 3. `MachineState` enum has gaps

**File:** `proto/metalstack/api/v2/machine.proto:509-518`

States: `UNSPECIFIED`, `TAINTED`, `LOCKED`, `AVAILABLE`. There is no `RESERVED`, `ALLOCATED`, or `IN_PROGRESS` state. Whether a machine is allocated is implied by the presence of the `allocation` sub-object rather than being a first-class state value. This means:

- Filtering by lifecycle phase requires checking a nullable sub-object relationship, not a simple enum
- The state enum describes _operator interventions_ (tainted/locked) but not the _normal progression_ of a machine's lifecycle

## 4. `MachineLiveliness` has no tracking context

**File:** `proto/metalstack/api/v2/machine.proto:696-705`

The enum values are `ALIVE`, `DEAD`, `UNKNOWN` — with no timestamp, threshold metadata, or staleness window. Compare with `MachineBMCReport.updated_at` which explicitly tracks when BMC data was last refreshed. The liveliness sits outside the BMC domain semantically (it's on `MachineStatus`, not `MachineBMCReport`), but BMC is what determines it — two loosely coupled sources of truth for the same physical signal.

## 5. `Tenant.login` doubles as identity key and mutable display identifier

**File:** `proto/metalstack/api/v2/tenant.proto:95-110`

`login` is the primary key — used in `TenantServiceGetRequest.login`, `TenantServiceDeleteRequest.login`, etc. — but it is a mutable human-readable string. Changing the key of an entity mid-lifecycle is risky. Most entities in this model use UUIDs as stable identifiers (`project.uuid`, `machine.uuid`, `ip.uuid`). Tenant alone uses a human-readable string as its identity anchor.

## 6. Project invitations have no user targeting

**File:** `proto/metalstack/api/v2/project.proto:238-253`

`ProjectServiceInviteRequest` takes a `project` and `role` but no `user` or `member`. Any person who obtains the invite secret can accept it. The only guard is the `joined` boolean on the `ProjectInvite` message. There is no binding between a specific user and an invitation, so secret sharing is the only mechanism controlling invite use.

## 7. `MachineNic.neighbors` is unbounded recursion

**File:** `proto/metalstack/api/v2/machine.proto:495`

`MachineNic.neighbors` is `repeated MachineNic` — a graph adjacency list stored as full NIC objects. A NIC can contain NICs which contain NICs, with no depth limit. In practice the data is shallow, but the schema allows arbitrary nesting.

## 8. `TaskInfo` payload is opaque bytes

**File:** `proto/metalstack/admin/v2/task.proto:88-137`

`payload` and `result` are both `bytes` with no typing, no `oneof`, and no `google.protobuf.Any`. This pushes schema enforcement entirely to the consumer. Compare with the rest of the model which is strongly typed throughout via protobuf messages.

## 9. Machine Issues and Provisioning Events — parallel classification systems

**File:** `proto/metalstack/api/v2/machine.proto:635-1003`

`MachineIssueType` and `MachineProvisioningEventType` overlap semantically:

- `MACHINE_ISSUE_TYPE_CRASH_LOOP` ↔ `MACHINE_PROVISIONING_EVENT_STATE_CRASHLOOP`
- `MACHINE_ISSUE_TYPE_FAILED_MACHINE_RECLAIM` ↔ `MACHINE_PROVISIONING_EVENT_STATE_FAILED_RECLAIM`

Two classification systems for the same underlying machine signals (provisioning state), with no formal mapping between them. Issues also mix infrastructure concerns (`ASN_UNIQUENESS`, `BMC_INFO_OUTDATED`) with lifecycle concerns (`CRASH_LOOP`, `FAILED_MACHINE_RECLAIM`).

## 10. `NetworkType` conflates multiple orthogonal concerns

**File:** `proto/metalstack/api/v2/network.proto:281-321`

The enum collapses three separate axes into one value:

| Axis                 | Values                      |
|----------------------|-----------------------------|
| Topological role     | `UNDERLAY`, `EXTERNAL`      |
| Subnetting purpose   | `SUPER`, `SUPER_NAMESPACED` |
| Project relationship | `CHILD`, `CHILD_SHARED`     |

A network could logically be both "super" (children can be created from it) and "external" (providing internet connectivity), but the enum forces a single value.

## 11. `IP` has dual identity: UUID and IP address

**File:** `proto/metalstack/api/v2/ip.proto:67-86`

`IP` has a `uuid` field, but `IPServiceGetRequest` and `IPServiceDeleteRequest` use the IP address string (`ip` field) as the lookup key. The `IPServiceUpdateRequest` also uses the `ip` address string. The `IPQuery` supports both `uuid` and `ip` as filters. Two identifiers for one entity, neither marked as authoritative, and the "name" field adds a third naming axis.

## 12. No billing model

**File:** `go/permissions/permissions.go:20`

The permissions package defines a `Chargeable` type (`map[string]bool`) but it is never populated. The protos have no cost, price, or SKU fields on any resource (machines, IPs, networks, reservations). Billing is handled outside this API, but the domain model has no hooks for it — no way to know what consuming a resource costs or who pays for it.

## 13. Infra/Machine roles have no service identity

**File:** `proto/metalstack/api/v2/common.proto:49-66, 59-66`

`InfraRole` and `MachineRole` define _what level of access_ a service has (editor/viewer), but there is no concept of _which service_ is calling. Audit traces (`AuditTrace`) capture `user` for end-user calls, but for infra calls the caller identity is reduced to a role level, not a service name. This makes infra audit trails harder to trace across the system.

## 14. Labels use a different update model than everything else

**File:** `proto/metalstack/api/v2/common.proto:124-186`

All entities carry `Labels` (a `map[string]string`), but the update path uses `UpdateLabels` with replace/patch strategies (`Labels` vs `LabelsPatch`). Every other entity field uses `optional` presence for partial updates. Labels follow a different mental model (explicit merge strategies), adding client complexity for what is otherwise just key-value metadata.
