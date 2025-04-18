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
  - In Addition to defaultchildprefixlength a max childprefixlength per AF should be possible
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
