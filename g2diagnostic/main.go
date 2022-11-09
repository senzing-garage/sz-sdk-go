/*
Package g2diagnostic is a Go wrapper over Senzing's G2Diagnostic C binding.

To use G2diagnostic, the LD_LIBRARY_PATH environment variable must include
a path to Senzing's libraries.  Example:

	export LD_LIBRARY_PATH=/opt/senzing/g2/lib
*/
package g2diagnostic

import (
	"context"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2diagnostic interface {
	CheckDBPerf(ctx context.Context, secondsToRun int) (string, error)
	ClearLastException(ctx context.Context) error
	CloseEntityListBySize(ctx context.Context, entityListBySizeHandle uintptr) error
	Destroy(ctx context.Context) error
	FetchNextEntityBySize(ctx context.Context, entityListBySizeHandle uintptr) (string, error)
	FindEntitiesByFeatureIDs(ctx context.Context, features string) (string, error)
	GetAvailableMemory(ctx context.Context) (int64, error)
	GetDataSourceCounts(ctx context.Context) (string, error)
	GetDBInfo(ctx context.Context) (string, error)
	GetEntityDetails(ctx context.Context, entityID int64, includeInternalFeatures int) (string, error)
	GetEntityListBySize(ctx context.Context, entitySize int) (uintptr, error)
	GetEntityResume(ctx context.Context, entityID int64) (string, error)
	GetEntitySizeBreakdown(ctx context.Context, minimumEntitySize int, includeInternalFeatures int) (string, error)
	GetFeature(ctx context.Context, libFeatID int64) (string, error)
	GetGenericFeatures(ctx context.Context, featureType string, maximumEstimatedCount int) (string, error)
	GetLastException(ctx context.Context) (string, error)
	GetLastExceptionCode(ctx context.Context) (int, error)
	GetLogicalCores(ctx context.Context) (int, error)
	GetMappingStatistics(ctx context.Context, includeInternalFeatures int) (string, error)
	GetPhysicalCores(ctx context.Context) (int, error)
	GetRelationshipDetails(ctx context.Context, relationshipID int64, includeInternalFeatures int) (string, error)
	GetResolutionStatistics(ctx context.Context) (string, error)
	GetTotalSystemMemory(ctx context.Context) (int64, error)
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error
	InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error
	Reinit(ctx context.Context, initConfigID int64) error
	SetLogLevel(ctx context.Context, logLevel logger.Level) error
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdTemplate = "senzing-6003%04d"

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var IdMessages = map[int]string{
	2001: "Call to G2Diagnostic_checkDBPerf(%d) failed. Return code: %d",
	2002: "Call to G2Diagnostic_closeEntityListBySize() failed. Return code: %d",
	2003: "Call to G2Diagnostic_destroy() failed.  Return code: %d",
	2004: "Call to G2Diagnostic_fetchNextEntityBySize() failed.  Return code: %d",
	2005: "Call to G2Diagnostic_findEntitiesByFeatureIDs(%s) failed. Return code: %d",
	2006: "Call to G2Diagnostic_getDataSourceCounts() failed. Return code: %d",
	2007: "Call to G2Diagnostic_getDBInfo() failed. Return code: %d",
	2008: "Call to G2Diagnostic_getEntityDetails(%d, %d) failed. Return code: %d",
	2009: "Call to G2Diagnostic_getEntityListBySize(%d) failed. Return code: %d",
	2010: "Call to G2Diagnostic_getEntityResume(%d) failed. Return code: %d",
	2011: "Call to G2Diagnostic_getEntitySizeBreakdown(%d, %d) failed. Return code: %d",
	2012: "Call to G2Diagnostic_getFeature(%d) failed. Return code: %d",
	2013: "Call to G2Diagnostic_getGenericFeatures(%s, %d) failed. Return code: %d",
	2014: "Call to G2Diagnostic_getLastException() failed. Return code: %d",
	2015: "Call to G2Diagnostic_getMappingStatistics(%d) failed. Return code: %d",
	2016: "Call to G2Diagnostic_getRelationshipDetails(%d, %d) failed. Return code: %d",
	2017: "Call to G2Diagnostic_getResolutionStatistics() failed. Return code: %d",
	2018: "Call to G2Diagnostic_init(%s, %s, %d) failed. Return code: %d",
	2019: "Call to G2Diagnostic_initWithConfigID(%s, %s, %d, %d) failed. Return code: %d",
	2020: "Call to G2Diagnostic_reinit(%d) failed. Return Code: %d",
	4001: "Enter CheckDBPerf(%d).",
	4002: "Exit  CheckDBPerf(%d) returned (%s, %v).",
	4003: "Enter ClearLastException().",
	4004: "Exit  ClearLastException() returned (%v).",
	4005: "Enter CloseEntityListBySize().",
	4006: "Exit  CloseEntityListBySize() returned (%v).",
	4007: "Enter Destroy().",
	4008: "Exit  Destroy() returned (%v).",
	4009: "Enter FetchNextEntityBySize().",
	4010: "Exit  FetchNextEntityBySize() returned (%s, %v).",
	4011: "Enter FindEntitiesByFeatureIDs(%s).",
	4012: "Exit  FindEntitiesByFeatureIDs(%s) returned (%s, %v).",
	4013: "Enter GetAvailableMemory().",
	4014: "Exit  GetAvailableMemory() returned (%d, %v).",
	4015: "Enter GetDataSourceCounts().",
	4016: "Exit  GetDataSourceCounts() returned (%s, %v).",
	4017: "Enter GetDBInfo().",
	4018: "Exit  GetDBInfo()  returned (%s, %v).",
	4019: "Enter GetEntityDetails(%d, %d).",
	4020: "Exit  GetEntityDetails(%d, %d) returned (%s, %v).",
	4021: "Enter GetEntityListBySize(%d).",
	4022: "Exit  GetEntityListBySize(%d) returned (%v, %v).",
	4023: "Enter GetEntityResume(%d).",
	4024: "Exit  GetEntityResume(%d) returned (%s, %v).",
	4025: "Enter GetEntitySizeBreakdown(%d, %d).",
	4026: "Exit  GetEntitySizeBreakdown(%d, %d) returned (%s, %v).",
	4027: "Enter GetFeature(%d).",
	4028: "Exit  GetFeature(%d) returned (%s, %v).",
	4029: "Enter GetGenericFeatures(%s, %d).",
	4030: "Exit  GetGenericFeatures(%s, %d) returned (%s, %v).",
	4031: "Enter GetLastException().",
	4032: "Exit  GetLastException() returned (%s, %v).",
	4033: "Enter GetLastExceptionCode().",
	4034: "Exit  GetLastExceptionCode() returned (%d, %v).",
	4035: "Enter GetLogicalCores().",
	4036: "Exit  GetLogicalCores() returned (%d, %v).",
	4037: "Enter GetMappingStatistics(%d).",
	4038: "Exit  GetMappingStatistics(%d) returned (%s, %v).",
	4039: "Enter GetPhysicalCores().",
	4040: "Exit  GetPhysicalCores() returned (%d, %v).",
	4041: "Enter GetRelationshipDetails(%d, %d).",
	4042: "Exit  GetRelationshipDetails(%d, %d) returned (%s, %v).",
	4043: "Enter GetResolutionStatistics().",
	4044: "Exit  GetResolutionStatistics() returned (%s, %v).",
	4045: "Enter GetTotalSystemMemory().",
	4046: "Exit  GetTotalSystemMemory() returned (%d, %v).",
	4047: "Enter Init(%s, %s, %d).",
	4048: "Exit  Init(%s, %s, %d) returned (%v).",
	4049: "Enter InitWithConfigID(%s, %s, %d, %d).",
	4050: "Exit  InitWithConfigID(%s, %s, %d, %d) returned (%v).",
	4051: "Enter Reinit(%d).",
	4052: "Exit  Reinit(%d) returned (%v).",
	4053: "Enter SetLogLevel(%v).",
	4054: "Exit  SetLogLevel(%v) returned (%v).",
}

var IdRanges = map[int]string{
	0000: logger.LevelInfoName,
	1000: logger.LevelWarnName,
	2000: logger.LevelErrorName,
	3000: logger.LevelDebugName,
	4000: logger.LevelTraceName,
	5000: logger.LevelFatalName,
	6000: logger.LevelPanicName,
}

var IdRangesLogLevel = map[int]logger.Level{
	0000: logger.LevelInfo,
	1000: logger.LevelWarn,
	2000: logger.LevelError,
	3000: logger.LevelDebug,
	4000: logger.LevelTrace,
	5000: logger.LevelFatal,
	6000: logger.LevelPanic,
}

var IdStatuses = map[int]string{
	2001: logger.LevelErrorName,
	2002: logger.LevelErrorName,
	2003: logger.LevelErrorName,
	2004: logger.LevelErrorName,
	2005: logger.LevelErrorName,
	2006: logger.LevelErrorName,
	2007: logger.LevelErrorName,
	2008: logger.LevelErrorName,
	2009: logger.LevelErrorName,
	2010: logger.LevelErrorName,
	2011: logger.LevelErrorName,
	2012: logger.LevelErrorName,
	2013: logger.LevelErrorName,
	2014: logger.LevelErrorName,
	2015: logger.LevelErrorName,
	2016: logger.LevelErrorName,
	2017: logger.LevelErrorName,
	2018: logger.LevelErrorName,
	2019: logger.LevelErrorName,
	2999: logger.LevelErrorName,
}
