# g2-sdk-go

## :warning: WARNING: g2-sdk-go is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

The Senzing g2-sdk-go packages provide interface definitions for implementations of the Senzing Go Software Development Kit.
`g2-sdk-go` contains no "running code", but serves to unify implementations of the Senzing Go SDK.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/g2-sdk-go.svg)](https://pkg.go.dev/github.com/senzing/g2-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/g2-sdk-go)](https://goreportcard.com/report/github.com/senzing/g2-sdk-go)
[![go-test.yaml](https://github.com/Senzing/g2-sdk-go/actions/workflows/go-test.yaml/badge.svg)](https://github.com/Senzing/g2-sdk-go/actions/workflows/go-test.yaml)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/Senzing/g2-sdk-go/blob/main/LICENSE)

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

- [Senzing/g2-sdk-go-base](https://github.com/Senzing/g2-sdk-go-base) - an SDK for calling the Senzing C library locally.
- [Senzing/g2-sdk-go-grpc](https://github.com/Senzing/g2-sdk-go-grpc) - an SDK for calling the Senzing C library via
  [gRPC](https://grpc.io/) network requests.
- [Senzing/g2-sdk-go-mock](https://github.com/Senzing/g2-sdk-go-mock) - [mock objects](https://en.wikipedia.org/wiki/Mock_object)
  for developing and testing without the Senzing C library.
- [Senzing/go-sdk-abstract-factory](https://github.com/Senzing/go-sdk-abstract-factory) - An
  [abstract factory pattern](https://en.wikipedia.org/wiki/Abstract_factory_pattern)
  for switching among implementations.

Documentation and examples for the implementations are at:

- [g2-sdk-go-base](https://pkg.go.dev/github.com/senzing/g2-sdk-go-base)
- [g2-sdk-go-grpc](https://pkg.go.dev/github.com/senzing/g2-sdk-go-grpc)
- [g2-sdk-go-mock](https://pkg.go.dev/github.com/senzing/g2-sdk-go-mock)
- [go-sdk-abstract-factory](https://pkg.go.dev/github.com/senzing/go-sdk-abstract-factory)

## References

1. [API documentation](https://pkg.go.dev/github.com/senzing/g2-sdk-go)
1. [Development](docs/development.md)
1. [Errors](docs/errors.md)
1. [Examples](docs/examples.md)
