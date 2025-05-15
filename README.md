# metal-stack.io api

[![Release](https://github.com/metal-stack/api/actions/workflows/main.yml/badge.svg)](https://github.com/metal-stack/api/actions/workflows/main.yml) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/metal-stack/api) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/metal-stack/api) ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/metal-stack/api) [![Go Report Card](https://goreportcard.com/badge/github.com/metal-stack/api)](https://goreportcard.com/report/github.com/metal-stack/api)

This is the POC V2 API of [metal-stack.io](https://metal-stack.io) to implement MEP-4.


TODO:

- validate all fields in all messages, also review them
- get rid of exposed nsq, replace with grpc streaming
- try to implement a generic reconcile which spans multiple backends

Meeting Minutes:

- We stick with networks without a project and treat them as global resource
- Write down a table with use cases which are actually possible and check if the new implementation still allows them
- If project is "" its a global resource, otherwise its project scoped. This requires more checks regarding visibility.

List of possible use cases:

- Admin creates a underlay network for a partition
  - creation of firewall and switch underlay ip allocation, currently only the firewall ip is allocated
    the switch VNI and Loopback ips are actually manually allocated by the ansible deployment
- Admin creates a super network for a partition (should be possible without partition for zonal support)
  - only project scoped child network allocation is possible, no ip address allocation
  - In Addition to defaultchildprefixlength a min childprefixlength per AF should be possible
- Admin creates a external network, either internet or other campus networks where communication must be possible
  - Only project scoped ip address allocation is possible, no child networks
- Admin creates a external super network which is used for project scoped child networks and inherit vrfShared
  - Users can allocate child which are then project scoped with their own project.
- Admin creates a shared network which is project scoped in a partition which can be used for storage for example.
  - Users can allocate a IP address in this network which are scoped with their project.

- User allocates a child network from a super network to be used as node network, this is scoped to his project
- User allocates ip addresses from his child networks, which are then scoped with the same project as the network itself.
- User allocates ip addresses from global networks, which are then project scoped with his project
- User allocates ip addresses from shared networks, which are then project scoped with his project

Future Use Cases:

tenant-super-namespaced-network             10.0.0.0/16         ""
├─╴ 10.0.0.0/16                                                 project-a
   ├──╴10.0.0.0/24                                              project-a
   ├──╴10.0.1.0/24                                              project-a
├─╴ 10.0.0.0/16                                                 project-b
   ├──╴10.0.0.0/24                                              project-b
   ├──╴10.0.1.0/24                                              project-b

- Admin can create super networks for project namespaced child networks. These networks will always start at the same prefix as the parent:
  Example:
    Parent: 10.0.0.0/16
      Child 1 in project 1 : 10.0.0.0/24
      Child 2 in project 1 : 10.0.1.0/24
      Child 3 in project 1 : 10.0.2.0/24
      Child 4 in project 1 : 10.0.3.0/24
      Child 5 in project 1 : 10.0.4.0/24

      Child 1 in project 2 : 10.0.0.0/24
      Child 2 in project 2 : 10.0.1.0/24
      Child 3 in project 2 : 10.0.2.0/24
      Child 4 in project 2 : 10.0.3.0/24
      Child 5 in project 2 : 10.0.4.0/24

# Meeting Minutes 2nd Call

- Check if is possible to create a Internet access network for single customer usage with specified destination prefixes


```proto
enum NetworkType {
  // NETWORK_TYPE_UNSPECIFIED indicates a unknown network type
  NETWORK_TYPE_UNSPECIFIED = 0;
  // NETWORK_TYPE_EXTERNAL indicates network where multiple projects can allocate ips, it offers connectivity to external networks
  // In most cases this is the internet network or a network which offers connectivity to legacy datacenter networks.
  // If it is not project scoped everyone can allocate Ips in this network, otherwise only from the same project ip allocation is possible.
  NETWORK_TYPE_EXTERNAL = 1 [(enum_string_value) = "external"];
  // NETWORK_TYPE_UNDERLAY indicates a underlay network
  // The underlay network connects all switches and the firewalls to build a EVPN dataplane
  // It is not project scoped. Is part of the dataplane and reserved for administrative purposes.
  NETWORK_TYPE_UNDERLAY = 2 [(enum_string_value) = "underlay"];

  // NETWORK_TYPE_SUPER indicates a super network which is only used to create child networks
  // If the vrf id is given, child networks will inherit this vrf.
  // If the vrf id is nil in this network, child vrf is taken from the pool.
  // If the partition is given, child networks inherit the partition.
  // If the partition is nil, child networks also do not have a partition (i.e. requires vrf is distributed across all partitions).
  // For child creation destination prefixes will be inherited
  // If this is project scoped, child project must match, otherwise can be freely specified.
  NETWORK_TYPE_SUPER = 3 [(enum_string_value) = "super"];
  // NETWORK_TYPE_SUPER_NAMESPACED indicates a super network which is only used to create child networks.
  // All rules from NETWORK_TYPE_SUPER apply for them as well.
  // In addition, a network namespace will be created for every project. Child networks per project will have disjunct prefixes.
  // Prefix allocation will start again with the same base cidr for every project / namespace.
  // This will allow the creation of much more child networks from a given super network size.
  NETWORK_TYPE_SUPER_NAMESPACED = 4 [(enum_string_value) = "super-namespaced"];

  // NETWORK_TYPE_CHILD indicates a child network of a project.
  // This is the only network type that can be created by a user.
  // Connectivity to external networks is not possible without going through an additional firewall in this network which creates connectivity to other networks.
  // Such a network will be created either from a super, or super namespaced.
  NETWORK_TYPE_CHILD = 5 [(enum_string_value) = "child"];
  // NETWORK_TYPE_CHILD_SHARED indicates a child network of a project which allows the allocation of ips from different projects.
  // Connectivity to external networks is not possible, as for normal child networks.
  // These networks are usually used to provide connectivity to shared services which are created in child networks, e.g. storage.
  // With this approach the number of hops can be reduced to the bare minimum in order to increase availability and performance.
  NETWORK_TYPE_CHILD_SHARED = 6 [(enum_string_value) = "child-shared"];
}
```

Please validate these enums with the given use-cases above.
