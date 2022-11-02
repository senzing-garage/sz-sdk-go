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
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdTemplate = "senzing-6003%04d"

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var IdMessages = map[int]string{
	1:    "Call to G2Diagnostic_checkDBPerf(%d) failed. Return code: %d",
	2:    "Call to G2Diagnostic_closeEntityListBySize() failed. Return code: %d",
	3:    "Call to G2Diagnostic_destroy() failed.  Return code: %d",
	4:    "Call to G2Diagnostic_fetchNextEntityBySize() failed.  Return code: %d",
	5:    "Call to G2Diagnostic_findEntitiesByFeatureIDs(%s) failed. Return code: %d",
	6:    "Call to G2Diagnostic_getDataSourceCounts() failed. Return code: %d",
	7:    "Call to G2Diagnostic_getDBInfo() failed. Return code: %d",
	8:    "Call to G2Diagnostic_getEntityDetails(%d, %d) failed. Return code: %d",
	9:    "Call to G2Diagnostic_getEntityListBySize(%d) failed. Return code: %d",
	10:   "Call to G2Diagnostic_getEntityResume(%d) failed. Return code: %d",
	11:   "Call to G2Diagnostic_getEntitySizeBreakdown(%d, %d) failed. Return code: %d",
	12:   "Call to G2Diagnostic_getFeature(%d) failed. Return code: %d",
	13:   "Call to G2Diagnostic_getGenericFeatures(%s, %d) failed. Return code: %d",
	14:   "Call to G2Diagnostic_getMappingStatistics(%d) failed. Return code: %d",
	15:   "Call to G2Diagnostic_getRelationshipDetails(%d, %d) failed. Return code: %d",
	16:   "Call to G2Diagnostic_getResolutionStatistics() failed. Return code: %d",
	17:   "Call to G2Diagnostic_init(%s, %s, %d) failed. Return code: %d",
	18:   "Call to G2Diagnostic_initWithConfigID(%s, %s, %d, %d) failed. Return code: %d",
	19:   "Call to G2Diagnostic_reinit(%d) failed. Return Code: %d",
	2999: "Cannot retrieve last error message.",
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

var IdStatuses = map[int]string{
	1:    logger.LevelErrorName,
	2:    logger.LevelErrorName,
	3:    logger.LevelErrorName,
	4:    logger.LevelErrorName,
	5:    logger.LevelErrorName,
	6:    logger.LevelErrorName,
	7:    logger.LevelErrorName,
	8:    logger.LevelErrorName,
	9:    logger.LevelErrorName,
	10:   logger.LevelErrorName,
	11:   logger.LevelErrorName,
	12:   logger.LevelErrorName,
	13:   logger.LevelErrorName,
	14:   logger.LevelErrorName,
	15:   logger.LevelErrorName,
	16:   logger.LevelErrorName,
	17:   logger.LevelErrorName,
	18:   logger.LevelErrorName,
	19:   logger.LevelErrorName,
	2999: logger.LevelErrorName,
}
