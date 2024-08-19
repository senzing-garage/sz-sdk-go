/*
The Senzing sz-sdk-go packages contain definitions that are common to implementations of the Senzing Go SDK.

# Synopsis

The Senzing [sz-sdk-go] packages contain definitions that are common to implementations of the Senzing Go SDK.
A client of the Senzing Go SDK API would use the senzing package to maintain compatiblity across
various Senzing Go SDK API implementations.

# Overview

The Senzing [sz-sdk-go] implementations enable Go programs to call Senzing library functions.

Implementations:

  - [sz-sdk-go-core] - An implementation that uses the Senzing native C binaries.
  - [sz-sdk-go-grpc] - An implementation that communicates with a [Senzing gRPC server].
  - [sz-sdk-go-mock] - An implementation that creates Senzing [mock objects].

More information at [sz-sdk-go].

[mock objects]: https://en.wikipedia.org/wiki/Mock_object
[Senzing gRPC server]: https://github.com/senzing-garage/serve-grpc
[sz-sdk-go-core]: https://pkg.go.dev/github.com/senzing-garage/sz-sdk-go-core
[sz-sdk-go-grpc]: https://pkg.go.dev/github.com/senzing-garage/sz-sdk-go-grpc
[sz-sdk-go-mock]: https://pkg.go.dev/github.com/senzing-garage/sz-sdk-go-mock
[sz-sdk-go]: https://github.com/senzing-garage/sz-sdk-go
*/
package main
