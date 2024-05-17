# sz-sdk-go

If you are beginning your journey with [Senzing],
please start with [Senzing Quick Start guides].

You are in the [Senzing Garage]
where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## :warning: WARNING: sz-sdk-go is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

The Senzing sz-sdk-go packages provide interface definitions for implementations of the Senzing Go Software Development Kit.
`sz-sdk-go` contains no "running code", but serves to unify implementations of the Senzing Go SDK.

[![Go Reference Badge]][Package reference]
[![Go Report Card Badge]][Go Report Card]
[![License Badge]][License]
[![golangci-lint.yaml Badge]][golangci-lint.yaml]
[![go-test-linux.yaml Badge]][go-test-linux.yaml]
[![go-test-darwin.yaml Badge]][go-test-darwin.yaml]
[![go-test-windows.yaml Badge]][go-test-windows.yaml]

## Overview

When programming, the recommendation is to "program to an interface".
The Senzing Go SDK API interfaces are documented at:

- [Szconfig]
- [Szconfigmanager]
- [Szdiagnostic]
- [Szengine]
- [Szproduct]

The Senzing sz-sdk-go packages provide interface definitions for the following implementations.
For hints on usage, review the top-most `main.go` and the `XXX_test.go` files in package directories.

- [Senzing/sz-sdk-go-base] - an SDK for calling the Senzing C library locally.
- [Senzing/sz-sdk-go-grpc] - an SDK for calling the Senzing C library via
  [gRPC](https://grpc.io/) network requests.
- [Senzing/sz-sdk-go-mock] - [mock objects]
  for developing and testing without the Senzing C library.
- [Senzing/go-sdk-abstract-factory] - An
  [abstract factory pattern]
  for switching among implementations.

Documentation and examples for the implementations are at:

- [sz-sdk-go-base]
- [sz-sdk-go-grpc]
- [sz-sdk-go-mock]
- [go-sdk-abstract-factory]

## References

1. [Development](docs/development.md)
1. [Errors](docs/errors.md)
1. [Examples](docs/examples.md)
1. [Package reference]

[abstract factory pattern]: https://en.wikipedia.org/wiki/Abstract_factory_pattern
[Go Reference Badge]: https://pkg.go.dev/badge/github.com/senzing-garage/sz-sdk-go.svg
[Go Report Card]: https://goreportcard.com/report/github.com/senzing-garage/sz-sdk-go
[Go Report Card Badge]: https://goreportcard.com/badge/github.com/senzing-garage/sz-sdk-go
[golangci-lint.yaml]: https://github.com/senzing-garage/sz-sdk-go/actions/workflows/golangci-lint.yaml
[go-test-linux.yaml]: https://github.com/senzing-garage/sz-sdk-go/actions/workflows/go-test-linux.yaml
[go-test-darwin.yaml]: https://github.com/senzing-garage/sz-sdk-go/actions/workflows/go-test-darwin.yaml
[go-test-windows.yaml]: https://github.com/senzing-garage/sz-sdk-go/actions/workflows/go-test-windows.yaml
[golangci-lint.yaml Badge]: https://github.com/senzing-garage/sz-sdk-go/actions/workflows/golangci-lint.yaml/badge.svg
[go-test-linux.yaml Badge]: https://github.com/senzing-garage/sz-sdk-go/actions/workflows/go-test-linux.yaml/badge.svg
[go-test-darwin.yaml Badge]: https://github.com/senzing-garage/sz-sdk-go/actions/workflows/go-test-darwin.yaml/badge.svg
[go-test-windows.yaml Badge]: https://github.com/senzing-garage/sz-sdk-go/actions/workflows/go-test-windows.yaml/badge.svg
[go-sdk-abstract-factory]: https://pkg.go.dev/github.com/senzing-garage/go-sdk-abstract-factory
[License]: https://github.com/senzing-garage/sz-sdk-go/blob/main/LICENSE
[License Badge]: https://img.shields.io/badge/License-Apache2-brightgreen.svg
[mock objects]: https://en.wikipedia.org/wiki/Mock_object
[Package reference]: https://pkg.go.dev/github.com/senzing-garage/sz-sdk-go
[Senzing]: https://senzing.com/
[Senzing Garage]: https://github.com/senzing-garage
[Senzing Quick Start guides]: https://docs.senzing.com/quickstart/
[Senzing/sz-sdk-go-base]: https://github.com/senzing-garage/sz-sdk-go-base
[Senzing/sz-sdk-go-grpc]: https://github.com/senzing-garage/sz-sdk-go-grpc
[Senzing/sz-sdk-go-mock]: https://github.com/senzing-garage/sz-sdk-go-mock
[Senzing/go-sdk-abstract-factory]: https://github.com/senzing-garage/go-sdk-abstract-factory
[Szconfig]: https://pkg.go.dev/github.com/senzing-garage/sz-sdk-go/szconfig#Szconfig
[Szconfigmanager]: https://pkg.go.dev/github.com/senzing-garage/sz-sdk-go/szconfigmanager#Szconfigmanager
[Szdiagnostic]: https://pkg.go.dev/github.com/senzing-garage/sz-sdk-go/szdiagnostic#Szdiagnostic
[Szengine]: https://pkg.go.dev/github.com/senzing-garage/sz-sdk-go/szengine#Szengine
[Szproduct]: https://pkg.go.dev/github.com/senzing-garage/sz-sdk-go/szproduct#Szproduct
[sz-sdk-go-base]: https://pkg.go.dev/github.com/senzing-garage/sz-sdk-go-base
[sz-sdk-go-grpc]: https://pkg.go.dev/github.com/senzing-garage/sz-sdk-go-grpc
[sz-sdk-go-mock]: https://pkg.go.dev/github.com/senzing-garage/sz-sdk-go-mock