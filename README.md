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

## Overview

When programming, the recommendation is to "program to an interface".
The Senzing Go SDK API interfaces are documented at:

- [G2Config](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2config#G2config)
- [G2Configmgr](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2configmgr#G2configmgr)
- [G2Diagnostic](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2diagnostic#G2diagnostic)
- [G2Engine](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2engine#G2engine)
- [G2Product](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2product#G2product)

The Senzing g2-sdk-go packages provide interface definitions for the following implementations:

- [Senzing/g2-sdk-go-base](https://github.com/Senzing/g2-sdk-go-base) - an SDK for calling the Senzing C library locally.
- [Senzing/g2-sdk-go-grpc](https://github.com/Senzing/g2-sdk-go-grpc) - an SDK for calling the Senzing C library via gRPC network requests.
- [Senzing/g2-sdk-go-mock](https://github.com/Senzing/g2-sdk-go-mock) - [mock objects](https://en.wikipedia.org/wiki/Mock_object) for developing without the Senzing C library.

Documentation and examples for the implementations are at:

- [Senzing/g2-sdk-go-base](https://pkg.go.dev/github.com/senzing/g2-sdk-go-base)
- [Senzing/g2-sdk-go-grpc](https://pkg.go.dev/github.com/senzing/g2-sdk-go-grpc)
- [Senzing/g2-sdk-go-mock](https://pkg.go.dev/github.com/senzing/g2-sdk-go-mock)
