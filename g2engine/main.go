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
	SetLogLevel(ctx context.Context, logLevel logger.Level) error
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
	2001: "Call to G2_addRecord(%s, %s, %s, %s) failed. Return code: %d",
	2002: "Call to G2_addRecordWithInfo(%s, %s, %s, %s, %d) failed. Return code: %d",
	2003: "Call to G2_addRecordWithInfoWithReturnedRecordID(%s, %s, %s, %d) failed. Return code: %d",
	2004: "Call to G2_addRecordWithReturnedRecordID(%s, %s, %s) failed. Return code: %d",
	2005: "Call to G2_checkRecord(%s, %s) failed. Return code: %d",
	2006: "Call to G2_closeExport(%v) failed. Return code: %d",
	2007: "Call to G2_deleteRecord(%s, %s, %s) failed. Return code: %d",
	2008: "Call to G2_deleteRecordWithInfo(%s, %s, %s, %d) failed. Return code: %d",
	2009: "Call to G2_destroy() failed. Return code: %d",
	2010: "Call to G2_exportConfigAndConfigID() failed. Return code: %d",
	2011: "Call to G2_exportConfig() failed. Return code: %d",
	2012: "Call to G2_exportCSVEntityReport(%s, %d) failed. Return code: %d",
	2013: "Call to G2_exportJSONEntityReport(%d) failed. Return code: %d",
	2014: "Call to G2_fetchNext(%v) failed. Return code: %d",
	2015: "Call to G2_findInterestingEntitiesByEntityID(%d, %d) failed. Return code: %d",
	2016: "Call to G2_findInterestingEntitiesByRecordID(%s, %s, %d) failed. Return code: %d",
	2017: "Call to G2_findNetworkByEntityID(%s, %d, %d, %d) failed. Return code: %d",
	2018: "Call to G2_findNetworkByEntityID_V2(%s, %d, %d, %d, %d) failed. Return code: %d",
	2019: "Call to G2_findNetworkByRecordID(%s, %d, %d, %d) failed. Return code: %d",
	2020: "Call to G2_findNetworkByRecordID_V2(%s, %d, %d, %d, %d) failed. Return code: %d",
	2021: "Call to G2_findPathByEntityID(%d, %d, %d) failed. Return code: %d",
	2022: "Call to G2_findPathByEntityID_V2(%d, %d, %d, %d) failed. Return code: %d",
	2023: "Call to G2_findPathByRecordID(%s, %s, %s, %s, %d) failed. Return code: %d",
	2024: "Call to G2_findPathByRecordID_V2(%s, %s, %s, %s, %d, %d) failed. Return code: %d",
	2025: "Call to G2_findPathExcludingByEntityID(%d, %d, %d, %s) failed. Return code: %d",
	2026: "Call to G2_findPathExcludingByEntityID_V2(%d, %d, %d, %s, %d) failed. Return code: %d",
	2027: "Call to G2_findPathExcludingByRecordID(%s, %s, %s, %s %d, %s) failed. Return code: %d",
	2028: "Call to G2_findPathExcludingByRecordID_V2(%s, %s, %s, %s %d, %s, %d) failed. Return code: %d",
	2029: "Call to G2_findPathIncludingSourceByEntityID(%d, %d, %d, %s, %s) failed. Return code: %d",
	2030: "Call to G2_findPathIncludingSourceByEntityID_V2(%d, %d, %d, %s, %s, %d) failed. Return code: %d",
	2031: "Call to G2_findPathIncludingSourceByRecordID(%s, %s, %s, %s %d, %s, %s) failed. Return code: %d",
	2032: "Call to G2_findPathIncludingSourceByRecordID_V2(%s, %s, %s, %s %d, %s, %s, %d) failed. Return code: %d",
	2033: "Call to G2_getActiveConfigID() failed. Return code: %d",
	2034: "Call to G2_getEntityByEntityID(%d) failed. Return code: %d",
	2035: "Call to G2_getEntityByEntityID_V2(%d, %d) failed. Return code: %d",
	2036: "Call to G2_getEntityByRecordID(%s, %s) failed. Return code: %d",
	2037: "Call to G2_getEntityByRecordID_V2(%s, %s, %d) failed. Return code: %d",
	2038: "Call to G2_getRecord(%s, %s) failed. Return code: %d",
	2039: "Call to G2_getRecord_V2(%s, %s, %d) failed. Return code: %d",
	2040: "Call to G2_getRedoRecord() failed. Return code: %d",
	2041: "Call to G2_getRepositoryLastModifiedTime() failed. Return code: %d",
	2042: "Call to G2_getVirtualEntityByRecordID(%s) failed. Return code: %d",
	2043: "Call to G2_getVirtualEntityByRecordID_V2(%s, %d) failed. Return code: %d",
	2044: "Call to G2_howEntityByEntityID(%d) failed. Return code: %d",
	2045: "Call to G2_howEntityByEntityID_V2(%d, %d) failed. Return code: %d",
	2046: "Call to G2_init(%s, %s, %d) failed. Return code: %d",
	2047: "Call to G2_initWithConfigID(%s, %s, %d, %d) failed. Return code: %d",
	2048: "Call to G2_primeEngine() failed. Return code: %d",
	2049: "Call to G2_process(%s) failed. Return code: %d",
	2050: "Call to G2_processRedoRecord() failed. Return code: %d",
	2051: "Call to G2_processRedoRecordWithInfo(%d) failed. Return code: %d",
	2052: "Call to G2_processWithInfo(%s, %d) failed. Return code: %d",
	2053: "Call to G2_processWithResponse(%s) failed. Return code: %d",
	2054: "Call to G2_processWithResponseResize(%s) failed. Return code: %d",
	2055: "Call to G2_purgeRepository() failed. Return code: %d",
	2056: "Call to G2_reevaluateEntity(%d, %d) failed. Return code: %d",
	2057: "Call to G2_reevaluateEntityWithInfo(%d, %d) failed. Return code: %d",
	2058: "Call to G2_reevaluateRecord(%s, %s, %d) failed. Return code: %d",
	2059: "Call to G2_reevaluateRecordWithInfo(%s, %s, %d) failed. Return code: %d",
	2060: "Call to G2_reinit(%d) failed. Return code: %d",
	2061: "Call to G2_replaceRecord(%s, %s, %s, %s) failed. Return code: %d",
	2062: "Call to G2_replaceRecordWithInfo(%s, %s, %s, %s, %d) failed. Return code: %d",
	2063: "Call to G2_searchByAttributes(%s) failed. Return code: %d",
	2064: "Call to G2_searchByAttributes_V2(%s, %d) failed. Return code: %d",
	2065: "Call to G2_stats() failed. Return code: %d",
	2066: "Call to G2_whyEntities(%d, %d) failed. Return code: %d",
	2067: "Call to G2_whyEntities_V2(%d, %d, %d) failed. Return code: %d",
	2068: "Call to G2_whyEntityByEntityID(%d) failed. Return code: %d",
	2069: "Call to G2_whyEntityByEntityID_V2(%d, %d) failed. Return code: %d",
	2070: "Call to G2_whyEntityByRecordID(%s, %s) failed. Return code: %d",
	2071: "Call to G2_whyEntityByRecordID_V2(%s, %s, %d) failed. Return code: %d",
	2072: "Call to G2_whyRecords(%s, %s, %s, %s) failed. Return code: %d",
	2073: "Call to G2_whyRecords_V2(%s, %s, %s, %s, %d) failed. Return code: %d",
	2999: "Cannot retrieve last error message.",
	4001: "Enter AddRecord(%s, %s, %s, %s).",
	4002: "Exit  AddRecord(%s, %s, %s, %s) returned (%v).",
	4003: "Enter AddRecordWithInfo(%s, %s, %s, %s, %d).",
	4004: "Exit  AddRecordWithInfo(%s, %s, %s, %s, %d) returned (%s, %v).",
	4005: "Enter AddRecordWithInfoWithReturnedRecordID(%s, %s, %s, %d).",
	4006: "Exit  AddRecordWithInfoWithReturnedRecordID(%s, %s, %s, %d) returned (%s, %s, %v).",
	4007: "Enter AddRecordWithReturnedRecordID(%s, %s, %s).",
	4008: "Exit  AddRecordWithReturnedRecordID(%s, %s, %s) returned (%s, %v).",
	4009: "Enter CheckRecord(%s, %s).",
	4010: "Exit  CheckRecord(%s, %s) returned (%s, %v).",
	4011: "Enter ClearLastException().",
	4012: "Exit  ClearLastException() returned (%v).",
	4013: "Enter CloseExport(%v).",
	4014: "Exit  CloseExport(%v) returned (%v).",
	4015: "Enter CountRedoRecords().",
	4016: "Exit  CountRedoRecords() returned (%d, %v).",
	4017: "Enter DeleteRecord(%s, %s, %s).",
	4018: "Exit  DeleteRecord(%s, %s, %s) returned (%v).",
	4019: "Enter DeleteRecordWithInfo(%s, %s, %s, %d).",
	4020: "Exit  DeleteRecordWithInfo(%s, %s, %s, %d) returned (%s, %v).",
	4021: "Enter Destroy().",
	4022: "Exit  Destroy() returned (%v).",
	4023: "Enter ExportConfigAndConfigID().",
	4024: "Exit  ExportConfigAndConfigID() returned (%s, %d, %v).",
	4025: "Enter ExportConfig().",
	4026: "Exit  ExportConfig() returned (%s, %v).",
	4027: "Enter ExportCSVEntityReport(%s, %d).",
	4028: "Exit  ExportCSVEntityReport(%s, %d) returned (%v, %v).",
	4029: "Enter ExportJSONEntityReport(%d).",
	4030: "Exit  ExportJSONEntityReport(%d) returned (%v, %v).",
	4031: "Enter FetchNext(%v).",
	4032: "Exit  FetchNext(%v) returned (%s, %v).",
	4033: "Enter FindInterestingEntitiesByEntityID(%d, %d).",
	4034: "Exit  FindInterestingEntitiesByEntityID(%d, %d) returned (%s, %v).",
	4035: "Enter FindInterestingEntitiesByRecordID(%s, %s, %d).",
	4036: "Exit  FindInterestingEntitiesByRecordID(%s, %s, %d) returned (%s, %v).",
	4037: "Enter FindNetworkByEntityID(%s, %d, %d, %d).",
	4038: "Exit  FindNetworkByEntityID(%s, %d, %d, %d) returned (%s, %v).",
	4039: "Enter FindNetworkByEntityID_V2(%s, %d, %d, %d, %d).",
	4040: "Exit  FindNetworkByEntityID_V2(%s, %d, %d, %d, %d) returned (%s, %v).",
	4041: "Enter FindNetworkByRecordID(%s, %d, %d, %d).",
	4042: "Exit  FindNetworkByRecordID(%s, %d, %d, %d) returned (%s, %v).",
	4043: "Enter FindNetworkByRecordID_V2(%s, %d, %d, %d, %d).",
	4044: "Exit  FindNetworkByRecordID_V2(%s, %d, %d, %d, %d) returned (%s, %v).",
	4045: "Enter FindPathByEntityID(%d, %d, %d).",
	4046: "Exit  FindPathByEntityID(%d, %d, %d) returned (%s, %v).",
	4047: "Enter FindPathByEntityID_V2(%d, %d, %d, %d).",
	4048: "Exit  FindPathByEntityID_V2(%d, %d, %d, %d) returned (%s, %v).",
	4049: "Enter FindPathByRecordID(%s, %s, %s, %s, %d).",
	4050: "Exit  FindPathByRecordID(%s, %s, %s, %s, %d) returned (%s, %v).",
	4051: "Enter FindPathByRecordID_V2(%s, %s, %s, %s, %d, %d).",
	4052: "Exit  FindPathByRecordID_V2(%s, %s, %s, %s, %d, %d) returned (%s, %v).",
	4053: "Enter FindPathExcludingByEntityID(%d, %d, %d, %s).",
	4054: "Exit  FindPathExcludingByEntityID(%d, %d, %d, %s) returned (%s, %v).",
	4055: "Enter FindPathExcludingByEntityID_V2(%d, %d, %d, %s, %d).",
	4056: "Exit  FindPathExcludingByEntityID_V2(%d, %d, %d, %s, %d) returned (%s, %v).",
	4057: "Enter FindPathExcludingByRecordID(%s, %s, %s, %s, %d, %s).",
	4058: "Exit  FindPathExcludingByRecordID(%s, %s, %s, %s, %d, %s) returned (%s, %v).",
	4059: "Enter FindPathExcludingByRecordID_V2(%s, %s, %s, %s, %d, %s, %d).",
	4060: "Exit  FindPathExcludingByRecordID_V2(%s, %s, %s, %s, %d, %s, %d) returned (%v).",
	4061: "Enter FindPathIncludingSourceByEntityID(%d, %d, %d, %s, %s).",
	4062: "Exit  FindPathIncludingSourceByEntityID(%d, %d, %d, %s, %s) returned (%s, %v).",
	4063: "Enter FindPathIncludingSourceByEntityID_V2(%d, %d, %d, %s, %s, %d).",
	4064: "Exit  FindPathIncludingSourceByEntityID_V2(%d, %d, %d, %s, %s, %d) returned (%s, %v).",
	4065: "Enter FindPathIncludingSourceByRecordID(%s, %s, %s, %s, %d, %s, %s).",
	4066: "Exit  FindPathIncludingSourceByRecordID(%s, %s, %s, %s, %d, %s, %s) returned (%s, %v).",
	4067: "Enter FindPathIncludingSourceByRecordID_V2(%s, %s, %s, %s, %d, %s, %s, %d).",
	4068: "Exit  FindPathIncludingSourceByRecordID_V2(%s, %s, %s, %s, %d, %s, %s, %d) returned (%s, %v).",
	4069: "Enter GetActiveConfigID().",
	4070: "Exit  GetActiveConfigID() returned (%d, %v).",
	4071: "Enter GetEntityByEntityID(%d).",
	4072: "Exit  GetEntityByEntityID(%d) returned (%s, %v).",
	4073: "Enter GetEntityByEntityID_V2(%d, %d).",
	4074: "Exit  GetEntityByEntityID_V2(%d, %d) returned (%s, %v).",
	4075: "Enter GetEntityByRecordID(%s, %s).",
	4076: "Exit  GetEntityByRecordID(%s, %s) returned (%s, %v).",
	4077: "Enter GetEntityByRecordID_V2(%s, %s, %d).",
	4078: "Exit  GetEntityByRecordID_V2(%s, %s, %d) returned (%s, %v).",
	4079: "Enter GetLastException().",
	4080: "Exit  GetLastException() returned (%s, %v).",
	4081: "Enter GetLastExceptionCode().",
	4082: "Exit  GetLastExceptionCode() returned (%d, %v).",
	4083: "Enter GetRecord(%s, %s).",
	4084: "Exit  GetRecord(%s, %s) returned (%s, %v).",
	4085: "Enter GetRecord_V2(%s, %s, %d).",
	4086: "Exit  GetRecord_V2(%s, %s, %d) returned (%s, %v).",
	4087: "Enter GetRedoRecord().",
	4088: "Exit  GetRedoRecord() returned (%s, %v).",
	4089: "Enter GetRepositoryLastModifiedTime().",
	4090: "Exit  GetRepositoryLastModifiedTime() returned (%d, %v).",
	4091: "Enter GetVirtualEntityByRecordID(%s).",
	4092: "Exit  GetVirtualEntityByRecordID(%s) returned (%s, %v).",
	4093: "Enter GetVirtualEntityByRecordID_V2(%s, %d).",
	4094: "Exit  GetVirtualEntityByRecordID_V2(%s, %d) returned (%s, %v).",
	4095: "Enter HowEntityByEntityID(%d).",
	4096: "Exit  HowEntityByEntityID(%d) returned (%s, %v).",
	4097: "Enter HowEntityByEntityID_V2(%d, %d).",
	4098: "Exit  HowEntityByEntityID_V2(%d, %d) returned (%s, %v).",
	4099: "Enter Init(%s, %s, %d).",
	4100: "Exit  Init(%s, %s, %d) returned (%v).",
	4101: "Enter InitWithConfigID(%s, %s, %d, %d).",
	4102: "Exit  InitWithConfigID(%s, %s, %d, %d) returned (%v).",
	4103: "Enter PrimeEngine().",
	4104: "Exit  PrimeEngine() returned (%v).",
	4105: "Enter Process(%s).",
	4106: "Exit  Process(%s) returned (%v).",
	4107: "Enter ProcessRedoRecord().",
	4108: "Exit  ProcessRedoRecord() returned (%s, %v).",
	4109: "Enter ProcessRedoRecordWithInfo(%d).",
	4110: "Exit  ProcessRedoRecordWithInfo(%d) returned (%s, %s, %v).",
	4111: "Enter ProcessWithInfo(%s, %d).",
	4112: "Exit  ProcessWithInfo(%s, %d) returned (%s, %v).",
	4113: "Enter ProcessWithResponse(%s).",
	4114: "Exit  ProcessWithResponse(%s) returned (%s, %v).",
	4115: "Enter ProcessWithResponseResize(%s).",
	4116: "Exit  ProcessWithResponseResize(%s) returned (%s, %v).",
	4117: "Enter PurgeRepository().",
	4118: "Exit  PurgeRepository() returned (%v).",
	4119: "Enter ReevaluateEntity(%d, %d).",
	4120: "Exit  ReevaluateEntity(%d, %d) returned (%v).",
	4121: "Enter ReevaluateEntityWithInfo(%d, %d).",
	4122: "Exit  ReevaluateEntityWithInfo(%d, %d) returned (%s, %v).",
	4123: "Enter ReevaluateRecord(%s, %s, %d).",
	4124: "Exit  ReevaluateRecord(%s, %s, %d) returned (%v).",
	4125: "Enter ReevaluateRecordWithInfo(%s, %s, %d).",
	4126: "Exit  ReevaluateRecordWithInfo(%s, %s, %d) returned (%s, %v).",
	4127: "Enter Reinit(%d).",
	4128: "Exit  Reinit(%d) returned (%v).",
	4129: "Enter ReplaceRecord(%s, %s, %s, %s).",
	4130: "Exit  ReplaceRecord(%s, %s, %s, %s) returned (%v).",
	4131: "Enter ReplaceRecordWithInfo(%s, %s, %s, %s, %d).",
	4132: "Exit  ReplaceRecordWithInfo(%s, %s, %s, %s, %d) returned (%s, %v).",
	4133: "Enter SearchByAttributes(%s).",
	4134: "Exit  SearchByAttributes(%s) returned (%s, %v).",
	4135: "Enter SearchByAttributes_V2(%s, %d).",
	4136: "Exit  SearchByAttributes_V2(%s, %d) returned (%s, %v).",
	4137: "Enter SetLogLevel(%v).",
	4138: "Exit  SetLogLevel(%v) returned (%v).",
	4139: "Enter Stats().",
	4140: "Exit  Stats() returned (%s, %v).",
	4141: "Enter WhyEntities(%d, %d).",
	4142: "Exit  WhyEntities(%d, %d) returned (%s, %v).",
	4143: "Enter WhyEntities_V2(%d, %d, %d).",
	4144: "Exit  WhyEntities_V2(%d, %d, %d) returned (%s, %v).",
	4145: "Enter WhyEntityByEntityID(%d).",
	4146: "Exit  WhyEntityByEntityID(%d) returned (%s, %v).",
	4147: "Enter WhyEntityByEntityID_V2(%d, %d).",
	4148: "Exit  WhyEntityByEntityID_V2(%d, %d) returned (%s, %v).",
	4149: "Enter WhyEntityByRecordID(%s, %s).",
	4150: "Exit  WhyEntityByRecordID(%s, %s) returned (%s, %v).",
	4151: "Enter WhyEntityByRecordID_V2(%s, %s, %d).",
	4152: "Exit  WhyEntityByRecordID_V2(%s, %s, %d) returned (%s, %v).",
	4153: "Enter WhyRecords(%s, %s, %s, %s).",
	4154: "Exit  WhyRecords(%s, %s, %s, %s) returned (%s, %v).",
	4155: "Enter WhyRecords_V2(%s, %s, %s, %s, %d).",
	4156: "Exit  WhyRecords_V2(%s, %s, %s, %s, %d) returned (%s, %v).",
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
