# metal-stack.io api

[![Release](https://github.com/metal-stack/api/actions/workflows/main.yml/badge.svg)](https://github.com/metal-stack/api/actions/workflows/main.yml) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/metal-stack/api) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/metal-stack/api) ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/metal-stack/api) [![Go Report Card](https://goreportcard.com/badge/github.com/metal-stack/api)](https://goreportcard.com/report/github.com/metal-stack/api)

This is the POC V2 API of [metal-stack.io](https://metal-stack.io) to implement MEP-4.


TODO:

- validate all fields in all messages, also review them
- get rid of exposed nsq, replace with grpc streaming
- try to implement a generic reconcile which spans multiple backends
