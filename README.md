# g2-sdk-go

If you are beginning your journey with
[Senzing](https://senzing.com/),
please start with
[Senzing Quick Start guides](https://docs.senzing.com/quickstart/).

You are in the
[Senzing Garage](https://github.com/senzing-garage)
where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## :warning: WARNING: g2-sdk-go is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

The Senzing g2-sdk-go packages provide interface definitions for implementations of the Senzing Go Software Development Kit.
`g2-sdk-go` contains no "running code", but serves to unify implementations of the Senzing Go SDK.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/g2-sdk-go.svg)](https://pkg.go.dev/github.com/senzing/g2-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/g2-sdk-go)](https://goreportcard.com/report/github.com/senzing/g2-sdk-go)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/senzing-garage/g2-sdk-go/blob/main/LICENSE)

[![gosec.yaml](https://github.com/senzing-garage/g2-sdk-go/actions/workflows/gosec.yaml/badge.svg)](https://github.com/senzing-garage/g2-sdk-go/actions/workflows/gosec.yaml)
[![go-test-linux.yaml](https://github.com/senzing-garage/g2-sdk-go/actions/workflows/go-test-linux.yaml/badge.svg)](https://github.com/senzing-garage/g2-sdk-go/actions/workflows/go-test-linux.yaml)
[![go-test-darwin.yaml](https://github.com/senzing-garage/g2-sdk-go/actions/workflows/go-test-darwin.yaml/badge.svg)](https://github.com/senzing-garage/g2-sdk-go/actions/workflows/go-test-darwin.yaml)
[![go-test-windows.yaml](https://github.com/senzing-garage/g2-sdk-go/actions/workflows/go-test-windows.yaml/badge.svg)](https://github.com/senzing-garage/g2-sdk-go/actions/workflows/go-test-windows.yaml)

## Overview

When programming, the recommendation is to "program to an interface".
The Senzing Go SDK API interfaces are documented at:

- [G2Config](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2config#G2config)
- [G2Configmgr](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2configmgr#G2configmgr)
- [G2Diagnostic](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2diagnostic#G2diagnostic)
- [G2Engine](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2engine#G2engine)
- [G2Product](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2product#G2product)

The Senzing g2-sdk-go packages provide interface definitions for the following implementations.
For hints on usage, review the top-most `main.go` and the `XXX_test.go` files in package directories.

- [Senzing/g2-sdk-go-base](https://github.com/senzing-garage/g2-sdk-go-base) - an SDK for calling the Senzing C library locally.
- [Senzing/g2-sdk-go-grpc](https://github.com/senzing-garage/g2-sdk-go-grpc) - an SDK for calling the Senzing C library via
  [gRPC](https://grpc.io/) network requests.
- [Senzing/g2-sdk-go-mock](https://github.com/senzing-garage/g2-sdk-go-mock) - [mock objects](https://en.wikipedia.org/wiki/Mock_object)
  for developing and testing without the Senzing C library.
- [Senzing/go-sdk-abstract-factory](https://github.com/senzing-garage/go-sdk-abstract-factory) - An
  [abstract factory pattern](https://en.wikipedia.org/wiki/Abstract_factory_pattern)
  for switching among implementations.

Documentation and examples for the implementations are at:

- [g2-sdk-go-base](https://pkg.go.dev/github.com/senzing/g2-sdk-go-base)
- [g2-sdk-go-grpc](https://pkg.go.dev/github.com/senzing/g2-sdk-go-grpc)
- [g2-sdk-go-mock](https://pkg.go.dev/github.com/senzing/g2-sdk-go-mock)
- [go-sdk-abstract-factory](https://pkg.go.dev/github.com/senzing/go-sdk-abstract-factory)

## References

1. [Development](docs/development.md)
1. [Errors](docs/errors.md)
1. [Examples](docs/examples.md)
1. [Package reference](https://pkg.go.dev/github.com/senzing/g2-sdk-go)
