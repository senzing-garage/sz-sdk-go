# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
[markdownlint](https://dlaa.me/markdownlint/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

-

## [0.2.6] - 2023-01-11

### Changed in 0.2.6

- Support `LD_LIBRARY_PATH` environment variable in `Makefile`
- Improve C string buffer resizing
- Improve `#cgo LDFLAGS:` parameters
- Minor bug fixes

## [0.2.5] - 2023-01-04

### Changed in 0.2.5

- Refactor code to reduce cyclomatic complexities and ineffectual assignments

## [0.2.4] - 2022-12-12

### Changed in 0.2.4

- Make getLastException(), getLastExceptionCode(), and clearLastException() private
- Added code to support *goroutines*

## [0.2.3] - 2022-12-08

### Changed in 0.2.3

- Improved godoc documentation
- Improved tests to initialize database prior to each run
- Add ability for GitHub Action testing

## [0.2.2] - 2022-11-18

### Changed in 0.2.2

- Update to github.com/senzing/go-logging v1.1.1

## [0.2.1] - 2022-11-09

### Added to 0.2.1

- Improved error detection

## [0.2.0] - 2022-11-08

### Added to 0.2.0

- Entry/Exit tracing when log level is set to TRACE

## [0.1.0] - 2022-11-02

### Added to 0.1.0

- Initial functionality
