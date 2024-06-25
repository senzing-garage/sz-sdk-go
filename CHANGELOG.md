# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], [markdownlint],
and this project adheres to [Semantic Versioning].

## [Unreleased]

-

## [0.13.4] - 2024-06-13

### Changed in 0.13.4

- Updated to changed in native Senzing C API

## [0.13.3] - 2024-06-11

### Changed in 0.13.3

## [0.13.2] - 2024-05-21

### Changed in 0.13.2

- Change "Id" to "ID"
- Change "Json" to "JSON"

## [0.13.1] - 2024-05-20

### Changed in 0.13.1

- `GetConfigList()` became `GetConfigs`

## [0.13.0] - 2024-05-17

### Changed in 0.13.0

- `senzing` migrated to `response`
- `sz` migrated to  `senzing`
- `bin/generate_response_response_test.py` modified for `response`
- Renamed variable to Go naming standards (e.g. `ID`, `JSON`, camelCase)
- Changes to pass new battery of GitHub action tests
- Refactored `szerror`

## [0.12.4] - 2024-05-13

### Changed in 0.12.4

- Identifiers for messages

## [0.12.3] - 2024-05-07

### Added in 0.12.3

- `FindInterestingEntitiesByEntityId` and `FindInterestingEntitiesByRecordId`
- `GetFeature`

### Changed in 0.12.3

- Flag variable names

## [0.12.2] - 2024-05-03

### Deleted in 0.12.2

- `GetRepositoryLastModifiedTime`

## [0.12.1] - 2024-04-26

### Changed in 0.12.1

- Changed from `CheckDatabasePerformance` to `CheckDatastorePerformance`
- Added `GetDatastoreInfo`

## [0.12.0] - 2024-04-19

### Changed in 0.12.0

- Updated to Senzing V4 method signatures
- Updated dependencies

## [0.11.1] - 2024-04-02

### Changed in 0.11.1

- Changed package from "szapi" to "szinterface"
- Changed package from "szconfigmgr" to "szconfigmanager"

## [0.11.0] - 2024-03-29

### Changed in 0.11.0

- Changed repository from "g2-sdk-go" to "sz-sdk-go"

## [0.10.1] - 2024-02-29

### Changed in 0.10.1

- Added G2Diagnostic.PurgeRepository()

## [0.10.0] - 2024-02-26

### Changed in 0.10.0

- Updated dependencies
- Updated names used in `g2-sdk-json-type-definition`
- Deleted methods not used in V4

## [0.9.0] - 2024-01-26

### Changed in 0.9.0

- Renamed module to github.com/senzing-garage/g2-sdk-go
- Refactor to [template-go]

## [0.8.0] - 2023-12-29

### Changed in 0.8.0

- Update dependencies
  - github.com/senzing-garage/g2-sdk-json-type-definition v0.2.0
  - github.com/senzing-garage/go-observing v0.3.0

## [0.7.6] - 2023-12-12

### Added in 0.7.6

- Additional support for Iterators

## [0.7.5] - 2023-12-11

### Added in 0.7.5

- `ExportCSVEntityReportIterator` and `ExportJSONEntityReportIterator`

## [0.7.4] - 2023-10-16

### Changed in 0.7.4

- Refactor to [template-go]
- Update dependencies
  - github.com/senzing/go-observing v0.2.8

## [0.7.3] - 2023-10-12

### Deleted in 0.7.3

- `g2product.ValidateLicenseFile`
- `g2product.ValidateLicenseStringBase64`

## [0.7.2] - 2023-10-12

### Changed in 0.7.2

- Changed from `int` to `int64` where required by the SenzingAPI

## [0.7.1] - 2023-10-10

### Changed in 0.7.1

- Updated `g2error` error codes

## [0.7.0] - 2023-09-01

### Changed in 0.7.0

- Support for SenzingAPI 3.8.0

### Removed in 0.7.0

- In `g2api.G2diagnostic.`
  - CloseEntityListBySize
  - FetchNextEntityBySize
  - FindEntitiesByFeatureIDs
  - GetDataSourceCounts
  - GetEntityDetails
  - GetEntityListBySize
  - GetEntityResume
  - GetEntitySizeBreakdown
  - GetFeature
  - GetGenericFeatures
  - GetMappingStatistics
  - GetRelationshipDetails
  - GetResolutionStatistics
- In `g2api.G2diagnostic.`
  - AddRecordWithInfoWithReturnedRecordID
  - AddRecordWithReturnedRecordID
  - CheckRecord
  - ProcessRedoRecord
  - ProcessRedoRecordWithInfo
  - ProcessWithResponse
  - ProcessWithResponseResize

## [0.6.9] - 2023-09-01

### Changed in 0.6.9

- Last version before SenzingAPI 3.8.0

## [0.6.8] - 2023-08-04

### Changed in 0.6.8

- Refactor to `template-go`

## [0.6.7] - 2023-07-05

### Added in 0.6.7

- Completed `senzing` package for unmarshalling

## [0.6.6] - 2023-06-16

### Added in 0.6.6

- Beginnings of `senzing` package

### Changed in 0.6.6

- Update dependencies
  - github.com/aquilax/truncate v1.0.0

## [0.6.5] - 2023-06-16

### Changed in 0.6.5

- Update dependencies
  - github.com/senzing/go-observing v0.2.6
  - github.com/stretchr/testify v1.8.4

## [0.6.4] - 2023-05-26

### Changed in 0.6.4

- Fixed method signature for `g2config.Load()`

## [0.6.3] - 2023-05-19

### Changed in 0.6.3

- Fixed error conversion
- Added `gosec` security testing
- Update dependencies
  - github.com/senzing/go-observing v0.2.5
  - github.com/stretchr/testify v1.8.3

## [0.6.2] - 2023-05-10

### Changed in 0.6.2

- Added `GetObserverOrigin()` and `SetObserverOrigin()` to g2api package
- Update dependencies
  - github.com/senzing/go-observing v0.2.2

## [0.6.1] - 2023-04-21

### Changed in 0.6.1

- Update dependencies
  - github.com/senzing/go-observing v0.2.1

## [0.6.0] - 2023-04-19

### Changed in 0.6.0

- Changed `SetLogLevel(ctx context.Context, logLevel logger.Level)` to `SetLogLevel(ctx context.Context, logLevelName string)`

## [0.5.1] - 2023-03-27

### Added in 0.5.1

- Added g2error.Convert()

## [0.5.0] - 2023-03-22

### Added in 0.5.0

- Added g2error.

## [0.4.1] - 2023-02-21

### Changed in 0.4.1

- Change GetSdkId() signature.

## [0.4.0] - 2023-02-14

### Changed in 0.4.0

- Major refactor.  Now g2-sdk-go only contains interface information, no implementation
- See g2-sdk-go-base for an implementation

## [0.3.1] - 2023-02-09

### Changed in 0.3.1

- Update GitActions
- Fix inaccurate assignments

## [0.3.0] - 2023-02-03

### Changed in 0.3.0

- Added Observer support

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

[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[markdownlint]: https://dlaa.me/markdownlint/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html
[template-go]: https://github.com/senzing-garage/template-go
