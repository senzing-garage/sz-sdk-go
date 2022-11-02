/*
Package g2engine is a Go wrapper over Senzing's G2Engine C binding.

To use G2engine, the LD_LIBRARY_PATH environment variable must include
a path to Senzing's libraries.  Example:

	export LD_LIBRARY_PATH=/opt/senzing/g2/lib
*/
package g2engine

import (
	"context"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type FlagMask int64

type G2engine interface {
	AddRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error
	AddRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error)
	AddRecordWithInfoWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string, flags int64) (string, string, error)
	AddRecordWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string) (string, error)
	CheckRecord(ctx context.Context, record string, recordQueryList string) (string, error)
	ClearLastException(ctx context.Context) error
	CloseExport(ctx context.Context, responseHandle uintptr) error
	CountRedoRecords(ctx context.Context) (int64, error)
	DeleteRecord(ctx context.Context, dataSourceCode string, recordID string, loadID string) error
	DeleteRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, loadID string, flags int64) (string, error)
	Destroy(ctx context.Context) error
	ExportConfig(ctx context.Context) (string, error)
	ExportConfigAndConfigID(ctx context.Context) (string, int64, error)
	ExportCSVEntityReport(ctx context.Context, csvColumnList string, flags int64) (uintptr, error)
	ExportJSONEntityReport(ctx context.Context, flags int64) (uintptr, error)
	FetchNext(ctx context.Context, responseHandle uintptr) (string, error)
	FindInterestingEntitiesByEntityID(ctx context.Context, entityID int64, flags int64) (string, error)
	FindInterestingEntitiesByRecordID(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	FindNetworkByEntityID(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error)
	FindNetworkByEntityID_V2(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error)
	FindNetworkByRecordID(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error)
	FindNetworkByRecordID_V2(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error)
	FindPathByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int) (string, error)
	FindPathByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, flags int64) (string, error)
	FindPathByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int) (string, error)
	FindPathByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, flags int64) (string, error)
	FindPathExcludingByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string) (string, error)
	FindPathExcludingByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, flags int64) (string, error)
	FindPathExcludingByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string) (string, error)
	FindPathExcludingByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, flags int64) (string, error)
	FindPathIncludingSourceByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, requiredDsrcs string) (string, error)
	FindPathIncludingSourceByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, requiredDsrcs string, flags int64) (string, error)
	FindPathIncludingSourceByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string) (string, error)
	FindPathIncludingSourceByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string, flags int64) (string, error)
	GetActiveConfigID(ctx context.Context) (int64, error)
	GetEntityByEntityID(ctx context.Context, entityID int64) (string, error)
	GetEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error)
	GetEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error)
	GetEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	GetLastException(ctx context.Context) (string, error)
	GetLastExceptionCode(ctx context.Context) (int, error)
	GetRecord(ctx context.Context, dataSourceCode string, recordID string) (string, error)
	GetRecord_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	GetRedoRecord(ctx context.Context) (string, error)
	GetRepositoryLastModifiedTime(ctx context.Context) (int64, error)
	GetVirtualEntityByRecordID(ctx context.Context, recordList string) (string, error)
	GetVirtualEntityByRecordID_V2(ctx context.Context, recordList string, flags int64) (string, error)
	HowEntityByEntityID(ctx context.Context, entityID int64) (string, error)
	HowEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error)
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error
	InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error
	PrimeEngine(ctx context.Context) error
	Process(ctx context.Context, record string) error
	ProcessRedoRecord(ctx context.Context) (string, error)
	ProcessRedoRecordWithInfo(ctx context.Context, flags int64) (string, string, error)
	ProcessWithInfo(ctx context.Context, record string, flags int64) (string, error)
	ProcessWithResponse(ctx context.Context, record string) (string, error)
	ProcessWithResponseResize(ctx context.Context, record string) (string, error)
	PurgeRepository(ctx context.Context) error
	ReevaluateEntity(ctx context.Context, entityID int64, flags int64) error
	ReevaluateEntityWithInfo(ctx context.Context, entityID int64, flags int64) (string, error)
	ReevaluateRecord(ctx context.Context, dataSourceCode string, recordID string, flags int64) error
	ReevaluateRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	Reinit(ctx context.Context, initConfigID int64) error
	ReplaceRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error
	ReplaceRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error)
	SearchByAttributes(ctx context.Context, jsonData string) (string, error)
	SearchByAttributes_V2(ctx context.Context, jsonData string, flags int64) (string, error)
	Stats(ctx context.Context) (string, error)
	WhyEntities(ctx context.Context, entityID1 int64, entityID2 int64) (string, error)
	WhyEntities_V2(ctx context.Context, entityID1 int64, entityID2 int64, flags int64) (string, error)
	WhyEntityByEntityID(ctx context.Context, entityID int64) (string, error)
	WhyEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error)
	WhyEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error)
	WhyEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	WhyRecords(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string) (string, error)
	WhyRecords_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, flags int64) (string, error)
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdTemplate = "senzing-6004%04d"

const (
	G2_EXPORT_INCLUDE_RESOLVED         FlagMask = 0         // 0 we should include entities with "resolved" relationships
	G2_EXPORT_INCLUDE_POSSIBLY_SAME    FlagMask = 1 << iota // 1 we should include entities with "possibly same" relationships
	G2_EXPORT_INCLUDE_POSSIBLY_RELATED                      // 2 we should include entities with "possibly related" relationships
	G2_EXPORT_INCLUDE_NAME_ONLY                             // 3 we should include entities with "name only" relationships
	G2_EXPORT_INCLUDE_DISCLOSED                             // 4 we should include entities with "disclosed" relationships
	G2_EXPORT_INCLUDE_SINGLETONS                            // 5 we should include singleton entities

	/* flags for outputting entity relation data  */
	G2_ENTITY_INCLUDE_POSSIBLY_SAME_RELATIONS    // 6 include "possibly same" relationships on entities
	G2_ENTITY_INCLUDE_POSSIBLY_RELATED_RELATIONS // 7 include "possibly related" relationships on entities
	G2_ENTITY_INCLUDE_NAME_ONLY_RELATIONS        // 8 include "name only" relationships on entities
	G2_ENTITY_INCLUDE_DISCLOSED_RELATIONS        // 9 include "disclosed" relationships on entities

	/* flags for outputting entity feature data  */
	G2_ENTITY_INCLUDE_ALL_FEATURES            // 10 include all features for entities
	G2_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES // 11 include only representative features on entities

	/* flags for getting extra information about an entity  */
	G2_ENTITY_INCLUDE_ENTITY_NAME            // 12 include the name of the entity
	G2_ENTITY_INCLUDE_RECORD_SUMMARY         // 13 include the record summary of the entity
	G2_ENTITY_INCLUDE_RECORD_DATA            // 14 include the basic record data for the entity
	G2_ENTITY_INCLUDE_RECORD_MATCHING_INFO   // 15 include the record matching info for the entity
	G2_ENTITY_INCLUDE_RECORD_JSON_DATA       // 16 include the record json data for the entity
	G2_ENTITY_INCLUDE_RECORD_FORMATTED_DATA  // 17 include the record formattted data for the entity
	G2_ENTITY_INCLUDE_RECORD_FEATURE_IDS     // 18 include the features identifiers for the records
	G2_ENTITY_INCLUDE_RELATED_ENTITY_NAME    // 19 include the name of the related entities
	G2_ENTITY_INCLUDE_RELATED_MATCHING_INFO  // 20 include the record matching info of the related entities
	G2_ENTITY_INCLUDE_RELATED_RECORD_SUMMARY // 21 include the record summary of the related entities
	G2_ENTITY_INCLUDE_RELATED_RECORD_DATA    // 22 include the basic record of the related entities

	/* flags for extra feature data  */
	G2_ENTITY_OPTION_INCLUDE_INTERNAL_FEATURES // 23 include internal features
	G2_ENTITY_OPTION_INCLUDE_FEATURE_STATS     // 24 include statistics on features

	/* flags for finding entity path data  */
	G2_FIND_PATH_PREFER_EXCLUDE // 25 excluded entities are still allowed, but not preferred

	/* flags for including search result feature scores  */
	G2_INCLUDE_FEATURE_SCORES // 26 include feature scores
	G2_SEARCH_INCLUDE_STATS   //  27 include statistics from search results
)

const (
	/* flags for exporting entity data  */
	G2_EXPORT_INCLUDE_ALL_ENTITIES      = G2_EXPORT_INCLUDE_RESOLVED | G2_EXPORT_INCLUDE_SINGLETONS
	G2_EXPORT_INCLUDE_ALL_RELATIONSHIPS = G2_EXPORT_INCLUDE_POSSIBLY_SAME | G2_EXPORT_INCLUDE_POSSIBLY_RELATED | G2_EXPORT_INCLUDE_NAME_ONLY | G2_EXPORT_INCLUDE_DISCLOSED

	/* flags for outputting entity relation data  */
	G2_ENTITY_INCLUDE_ALL_RELATIONS  = G2_ENTITY_INCLUDE_POSSIBLY_SAME_RELATIONS | G2_ENTITY_INCLUDE_POSSIBLY_RELATED_RELATIONS | G2_ENTITY_INCLUDE_NAME_ONLY_RELATIONS | G2_ENTITY_INCLUDE_DISCLOSED_RELATIONS
	G2_SEARCH_INCLUDE_FEATURE_SCORES = G2_INCLUDE_FEATURE_SCORES // include feature scores from search results

	/* flags for searching for entities  */
	G2_SEARCH_INCLUDE_RESOLVED         = G2_EXPORT_INCLUDE_RESOLVED
	G2_SEARCH_INCLUDE_POSSIBLY_SAME    = G2_EXPORT_INCLUDE_POSSIBLY_SAME
	G2_SEARCH_INCLUDE_POSSIBLY_RELATED = G2_EXPORT_INCLUDE_POSSIBLY_RELATED
	G2_SEARCH_INCLUDE_NAME_ONLY        = G2_EXPORT_INCLUDE_NAME_ONLY
	G2_SEARCH_INCLUDE_ALL_ENTITIES     = G2_SEARCH_INCLUDE_RESOLVED | G2_SEARCH_INCLUDE_POSSIBLY_SAME | G2_SEARCH_INCLUDE_POSSIBLY_RELATED | G2_SEARCH_INCLUDE_NAME_ONLY

	/* recommended settings */
	G2_RECORD_DEFAULT_FLAGS       = G2_ENTITY_INCLUDE_RECORD_JSON_DATA                                                                                                                                                                                                                                                                                                                   // the recommended default flag values for getting records
	G2_ENTITY_DEFAULT_FLAGS       = G2_ENTITY_INCLUDE_ALL_RELATIONS | G2_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES | G2_ENTITY_INCLUDE_ENTITY_NAME | G2_ENTITY_INCLUDE_RECORD_SUMMARY | G2_ENTITY_INCLUDE_RECORD_DATA | G2_ENTITY_INCLUDE_RECORD_MATCHING_INFO | G2_ENTITY_INCLUDE_RELATED_ENTITY_NAME | G2_ENTITY_INCLUDE_RELATED_RECORD_SUMMARY | G2_ENTITY_INCLUDE_RELATED_MATCHING_INFO // the recommended default flag values for getting entities
	G2_ENTITY_BRIEF_DEFAULT_FLAGS = G2_ENTITY_INCLUDE_RECORD_MATCHING_INFO | G2_ENTITY_INCLUDE_ALL_RELATIONS | G2_ENTITY_INCLUDE_RELATED_MATCHING_INFO                                                                                                                                                                                                                                   // the recommended default flag values for a brief entity result
	G2_EXPORT_DEFAULT_FLAGS       = G2_EXPORT_INCLUDE_ALL_ENTITIES | G2_EXPORT_INCLUDE_ALL_RELATIONSHIPS | G2_ENTITY_INCLUDE_ALL_RELATIONS | G2_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES | G2_ENTITY_INCLUDE_ENTITY_NAME | G2_ENTITY_INCLUDE_RECORD_DATA | G2_ENTITY_INCLUDE_RECORD_MATCHING_INFO | G2_ENTITY_INCLUDE_RELATED_MATCHING_INFO                                                // the recommended default flag values for exporting entities
	G2_FIND_PATH_DEFAULT_FLAGS    = G2_ENTITY_INCLUDE_ALL_RELATIONS | G2_ENTITY_INCLUDE_ENTITY_NAME | G2_ENTITY_INCLUDE_RECORD_SUMMARY | G2_ENTITY_INCLUDE_RELATED_MATCHING_INFO                                                                                                                                                                                                         // the recommended default flag values for finding entity paths
	G2_WHY_ENTITY_DEFAULT_FLAGS   = G2_ENTITY_DEFAULT_FLAGS | G2_ENTITY_INCLUDE_RECORD_FEATURE_IDS | G2_ENTITY_OPTION_INCLUDE_INTERNAL_FEATURES | G2_ENTITY_OPTION_INCLUDE_FEATURE_STATS | G2_INCLUDE_FEATURE_SCORES                                                                                                                                                                     // the recommended default flag values for why-analysis on entities
	G2_HOW_ENTITY_DEFAULT_FLAGS   = G2_ENTITY_DEFAULT_FLAGS | G2_ENTITY_INCLUDE_RECORD_FEATURE_IDS | G2_ENTITY_OPTION_INCLUDE_INTERNAL_FEATURES | G2_ENTITY_OPTION_INCLUDE_FEATURE_STATS | G2_INCLUDE_FEATURE_SCORES                                                                                                                                                                     // the recommended default flag values for how-analysis on entities

	G2_SEARCH_BY_ATTRIBUTES_ALL            = G2_SEARCH_INCLUDE_ALL_ENTITIES | G2_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES | G2_ENTITY_INCLUDE_ENTITY_NAME | G2_ENTITY_INCLUDE_RECORD_SUMMARY | G2_SEARCH_INCLUDE_FEATURE_SCORES                               // the recommended flag values for searching by attributes, returning all matching entities
	G2_SEARCH_BY_ATTRIBUTES_STRONG         = G2_SEARCH_INCLUDE_RESOLVED | G2_SEARCH_INCLUDE_POSSIBLY_SAME | G2_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES | G2_ENTITY_INCLUDE_ENTITY_NAME | G2_ENTITY_INCLUDE_RECORD_SUMMARY | G2_SEARCH_INCLUDE_FEATURE_SCORES // the recommended flag values for searching by attributes, returning only strongly matching entities
	G2_SEARCH_BY_ATTRIBUTES_MINIMAL_ALL    = G2_SEARCH_INCLUDE_ALL_ENTITIES                                                                                                                                                                                 // return minimal data with all matches
	G2_SEARCH_BY_ATTRIBUTES_MINIMAL_STRONG = G2_SEARCH_INCLUDE_RESOLVED | G2_SEARCH_INCLUDE_POSSIBLY_SAME                                                                                                                                                   // return minimal data with only the strongest matches
	G2_SEARCH_BY_ATTRIBUTES_DEFAULT_FLAGS  = G2_SEARCH_BY_ATTRIBUTES_ALL                                                                                                                                                                                    // the recommended default flag values for search-by-attributes
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var IdMessages = map[int]string{
	1:    "Call to G2_addRecord(%s, %s, %s, %s) failed. Return code: %d",
	2:    "Call to G2_addRecordWithInfo(%s, %s, %s, %s, %d) failed. Return code: %d",
	3:    "Call to G2_addRecordWithInfoWithReturnedRecordID(%s, %s, %s, %d) failed. Return code: %d",
	4:    "Call to G2_addRecordWithReturnedRecordID(%s, %s, %s) failed. Return code: %d",
	5:    "Call to G2_checkRecord(%s, %s) failed. Return code: %d",
	6:    "Call to G2_closeExport() failed. Return code: %d",
	7:    "Call to G2_deleteRecord(%s, %s, %s) failed. Return code: %d",
	8:    "Call to G2_deleteRecordWithInfo(%s, %s, %s, %d) failed. Return code: %d",
	9:    "Call to G2_destroy() failed. Return code: %d",
	10:   "Call to G2_exportConfigAndConfigID() failed. Return code: %d",
	11:   "Call to G2_exportConfig() failed. Return code: %d",
	12:   "Call to G2_exportCSVEntityReport(%s, %d) failed. Return code: %d",
	13:   "Call to G2_exportJSONEntityReport(%d) failed. Return code: %d",
	14:   "Call to G2_fetchNext() failed. Return code: %d",
	15:   "Call to G2_findInterestingEntitiesByEntityID(%d, %d) failed. Return code: %d",
	16:   "Call to G2_findInterestingEntitiesByRecordID(%s, %s, %d) failed. Return code: %d",
	17:   "Call to G2_findNetworkByEntityID(%s, %d, %d, %d) failed. Return code: %d",
	18:   "Call to G2_findNetworkByEntityID_V2(%s, %d, %d, %d, %d) failed. Return code: %d",
	19:   "Call to G2_findNetworkByRecordID(%s, %d, %d, %d) failed. Return code: %d",
	20:   "Call to G2_findNetworkByRecordID_V2(%s, %d, %d, %d, %d) failed. Return code: %d",
	21:   "Call to G2_findPathByEntityID(%d, %d, %d) failed. Return code: %d",
	22:   "Call to G2_findPathByEntityID_V2(%d, %d, %d, %d) failed. Return code: %d",
	23:   "Call to G2_findPathByRecordID(%s, %s, %s, %s, %d) failed. Return code: %d",
	24:   "Call to G2_findPathByRecordID_V2(%s, %s, %s, %s, %d, %d) failed. Return code: %d",
	25:   "Call to G2_findPathExcludingByEntityID(%d, %d, %d, %s) failed. Return code: %d",
	26:   "Call to G2_findPathExcludingByEntityID_V2(%d, %d, %d, %s, %d) failed. Return code: %d",
	27:   "Call to G2_findPathExcludingByRecordID(%s, %s, %s, %s %d, %s) failed. Return code: %d",
	28:   "Call to G2_findPathExcludingByRecordID_V2(%s, %s, %s, %s %d, %s, %d) failed. Return code: %d",
	29:   "Call to G2_findPathIncludingSourceByEntityID(%d, %d, %d, %s, %s) failed. Return code: %d",
	30:   "Call to G2_findPathIncludingSourceByEntityID_V2(%d, %d, %d, %s, %s, %d) failed. Return code: %d",
	31:   "Call to G2_findPathIncludingSourceByRecordID(%s, %s, %s, %s %d, %s, %s) failed. Return code: %d",
	32:   "Call to G2_findPathIncludingSourceByRecordID_V2(%s, %s, %s, %s %d, %s, %s, %d) failed. Return code: %d",
	33:   "Call to G2_getActiveConfigID() failed. Return code: %d",
	34:   "Call to G2_getEntityByEntityID(%d) failed. Return code: %d",
	35:   "Call to G2_getEntityByEntityID_V2(%d, %d) failed. Return code: %d",
	36:   "Call to G2_getEntityByRecordID(%s, %s) failed. Return code: %d",
	37:   "Call to GetEntityByRecordID_V2(%s, %s, %d) failed. Return code: %d",
	38:   "Call to G2_getRecord(%s, %s) failed. Return code: %d",
	39:   "Call to G2_getRecord_V2(%s, %s, %d) failed. Return code: %d",
	40:   "Call to G2_getRedoRecord() failed. Return code: %d",
	41:   "Call to G2_getRepositoryLastModifiedTime() failed. Return code: %d",
	42:   "Call to G2_getVirtualEntityByRecordID(%s) failed. Return code: %d",
	43:   "Call to G2_getVirtualEntityByRecordID_V2(%s, %d) failed. Return code: %d",
	44:   "Call to G2_howEntityByEntityID(%d) failed. Return code: %d",
	45:   "Call to G2_howEntityByEntityID_V2(%d, %d) failed. Return code: %d",
	46:   "Call to G2_init(%s, %s, %d) failed. Return code: %d",
	47:   "Call to G2_initWithConfigID(%s, %s, %d, %d) failed. Return code: %d",
	48:   "Call to G2_primeEngine() failed. Return code: %d",
	49:   "Call to G2_process(%s) failed. Return code: %d",
	50:   "Call to G2_processRedoRecord() failed. Return code: %d",
	51:   "Call to G2_processRedoRecordWithInfo(%d) failed. Return code: %d",
	52:   "Call to G2_processWithInfo(%s, %d) failed. Return code: %d",
	53:   "Call to G2_processWithResponse(%s) failed. Return code: %d",
	54:   "Call to G2_processWithResponseResize(%s) failed. Return code: %d",
	55:   "Call to PurgeRepository() failed. Return code: %d",
	56:   "Call to G2_reevaluateEntity(%d, %d) failed. Return code: %d",
	57:   "Call to G2_reevaluateEntityWithInfo(%d, %d) failed. Return code: %d",
	58:   "Call to G2_reevaluateRecord(%s, %s, %d) failed. Return code: %d",
	59:   "Call to G2_reevaluateRecordWithInfo(%s, %s, %d) failed. Return code: %d",
	60:   "Call to G2_reinit(%d) failed. Return code: %d",
	61:   "Call to G2_replaceRecord(%s, %s, %s, %s) failed. Return code: %d",
	62:   "Call to ReplaceRecordWithInfo(%s, %s, %s, %s, %d) failed. Return code: %d",
	63:   "Call to G2_searchByAttributes(%s) failed. Return code: %d",
	64:   "Call to G2_searchByAttributes_V2(%s, %d) failed. Return code: %d",
	65:   "Call to G2_stats() failed. Return code: %d",
	66:   "Call to G2_whyEntities(%d, %d) failed. Return code: %d",
	67:   "Call to G2_whyEntities_V2(%d, %d, %d) failed. Return code: %d",
	68:   "Call to G2_whyEntityByEntityID(%d) failed. Return code: %d",
	69:   "Call to G2_whyEntityByEntityID_V2(%d, %d) failed. Return code: %d",
	70:   "Call to G2_whyEntityByRecordID(%s, %s) failed. Return code: %d",
	71:   "Call to G2_whyEntityByRecordID_V2(%s, %s, %d) failed. Return code: %d",
	72:   "Call to G2_whyRecords(%s, %s, %s, %s) failed. Return code: %d",
	73:   "Call to G2_whyRecords_V2(%s, %s, %s, %s, %d) failed. Return code: %d",
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
	20:   logger.LevelErrorName,
	21:   logger.LevelErrorName,
	22:   logger.LevelErrorName,
	23:   logger.LevelErrorName,
	24:   logger.LevelErrorName,
	25:   logger.LevelErrorName,
	26:   logger.LevelErrorName,
	27:   logger.LevelErrorName,
	28:   logger.LevelErrorName,
	29:   logger.LevelErrorName,
	30:   logger.LevelErrorName,
	31:   logger.LevelErrorName,
	32:   logger.LevelErrorName,
	33:   logger.LevelErrorName,
	34:   logger.LevelErrorName,
	35:   logger.LevelErrorName,
	36:   logger.LevelErrorName,
	37:   logger.LevelErrorName,
	38:   logger.LevelErrorName,
	39:   logger.LevelErrorName,
	40:   logger.LevelErrorName,
	41:   logger.LevelErrorName,
	42:   logger.LevelErrorName,
	43:   logger.LevelErrorName,
	44:   logger.LevelErrorName,
	45:   logger.LevelErrorName,
	46:   logger.LevelErrorName,
	47:   logger.LevelErrorName,
	48:   logger.LevelErrorName,
	49:   logger.LevelErrorName,
	50:   logger.LevelErrorName,
	51:   logger.LevelErrorName,
	52:   logger.LevelErrorName,
	53:   logger.LevelErrorName,
	54:   logger.LevelErrorName,
	55:   logger.LevelErrorName,
	56:   logger.LevelErrorName,
	57:   logger.LevelErrorName,
	58:   logger.LevelErrorName,
	59:   logger.LevelErrorName,
	60:   logger.LevelErrorName,
	61:   logger.LevelErrorName,
	62:   logger.LevelErrorName,
	63:   logger.LevelErrorName,
	64:   logger.LevelErrorName,
	65:   logger.LevelErrorName,
	66:   logger.LevelErrorName,
	67:   logger.LevelErrorName,
	68:   logger.LevelErrorName,
	69:   logger.LevelErrorName,
	70:   logger.LevelErrorName,
	71:   logger.LevelErrorName,
	72:   logger.LevelErrorName,
	73:   logger.LevelErrorName,
	2999: logger.LevelErrorName,
}
