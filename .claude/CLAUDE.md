# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

The Senzing sz-sdk-go packages provide interface definitions for implementations of the Senzing Go Software Development Kit. This repository contains no "running code", but serves to unify implementations of the Senzing Go SDK including sz-sdk-go-core, sz-sdk-go-grpc, and sz-sdk-go-mock.

## Common Commands

### Test

```bash
make test               # Run tests with formatted output
go test ./...           # Run tests directly
go test -v ./... -run TestName  # Run a specific test
```

### Lint

```bash
make lint               # Run all linters (golangci-lint, govulncheck, cspell)
make golangci-lint      # Run golangci-lint only
make fix                # Auto-fix linting issues
```

### Coverage

```bash
make coverage           # Generate coverage report and open in browser
make check-coverage     # Check coverage thresholds (80% file/package/total)
```

### Other

```bash
make clean              # Clean build artifacts and caches
make dependencies       # Update Go dependencies
make dependencies-for-development  # Install development tools
```

## Architecture

### Project Structure

- `szconfig/` - SzConfig interface for managing Senzing configurations
- `szconfigmanager/` - SzConfigManager interface for configuration management
- `szdiagnostic/` - SzDiagnostic interface for diagnostic operations
- `szengine/` - SzEngine interface for entity resolution operations
- `szproduct/` - SzProduct interface for product information
- `senzing/` - Common Senzing types and response structures
- `szerror/` - Error types and handling for the SDK

### Key Patterns

**Interface Pattern**: Each package defines interfaces in `main.go` that implementations (sz-sdk-go-core, sz-sdk-go-grpc, sz-sdk-go-mock) must implement.

**Response Types**: The `senzing/` package contains response type definitions used across implementations.

**Error Handling**: Use `szerror/` package types for consistent error handling across implementations.

### Makefile System

The Makefile uses OS detection with platform-specific includes:

- `makefiles/osdetect.mk` - Detects OS type and architecture
- `makefiles/{darwin,linux,windows}.mk` - OS-specific target implementations

## Linting Configuration

Golangci-lint config: `.github/linters/.golangci.yaml`

- Line length: 120 characters
- Coverage thresholds: 80% (configurable in `.github/coverage/testcoverage.yaml`)
- Uses extensive linter set including exhaustruct, wrapcheck, err113
