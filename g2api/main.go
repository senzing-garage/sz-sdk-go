package g2api

import (
	"context"

	"github.com/senzing-garage/go-observing/observer"
)

// ----------------------------------------------------------------------------
// Types - struct
// ----------------------------------------------------------------------------

// StringFragment is used as a return value when iterating over log strings.
type StringFragment struct {
	Error error
	Value string
}

// ----------------------------------------------------------------------------
// Types - interface
// ----------------------------------------------------------------------------

// The G2config interface is a Golang representation of Senzing's libg2config.h
type G2config interface {
	AddDataSource(ctx context.Context, configHandle uintptr, inputJson string) (string, error)
	Close(ctx context.Context, configHandle uintptr) error
	Create(ctx context.Context) (uintptr, error)
	DeleteDataSource(ctx context.Context, configHandle uintptr, inputJson string) error
	Destroy(ctx context.Context) error
	GetObserverOrigin(ctx context.Context) string
	GetSdkId(ctx context.Context) string
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int64) error
	ListDataSources(ctx context.Context, configHandle uintptr) (string, error)
	Load(ctx context.Context, jsonConfig string) (uintptr, error)
	RegisterObserver(ctx context.Context, observer observer.Observer) error
	Save(ctx context.Context, configHandle uintptr) (string, error)
	SetLogLevel(ctx context.Context, logLevelName string) error
	SetObserverOrigin(ctx context.Context, origin string)
	UnregisterObserver(ctx context.Context, observer observer.Observer) error
}

// The G2configmgr interface is a Golang representation of Senzing's libg2configmgr.h
type G2configmgr interface {
	AddConfig(ctx context.Context, configStr string, configComments string) (int64, error)
	Destroy(ctx context.Context) error
	GetConfig(ctx context.Context, configID int64) (string, error)
	GetConfigList(ctx context.Context) (string, error)
	GetDefaultConfigID(ctx context.Context) (int64, error)
	GetObserverOrigin(ctx context.Context) string
	GetSdkId(ctx context.Context) string
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int64) error
	RegisterObserver(ctx context.Context, observer observer.Observer) error
	ReplaceDefaultConfigID(ctx context.Context, oldConfigID int64, newConfigID int64) error
	SetDefaultConfigID(ctx context.Context, configID int64) error
	SetLogLevel(ctx context.Context, logLevelName string) error
	SetObserverOrigin(ctx context.Context, origin string)
	UnregisterObserver(ctx context.Context, observer observer.Observer) error
}

// The G2diagnostic interface is a Golang representation of Senzing's libg2diagnostic.h
type G2diagnostic interface {
	CheckDBPerf(ctx context.Context, secondsToRun int) (string, error)
	Destroy(ctx context.Context) error
	GetObserverOrigin(ctx context.Context) string
	GetSdkId(ctx context.Context) string
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int64) error
	InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int64) error
	RegisterObserver(ctx context.Context, observer observer.Observer) error
	Reinit(ctx context.Context, initConfigID int64) error
	SetLogLevel(ctx context.Context, logLevelName string) error
	SetObserverOrigin(ctx context.Context, origin string)
	UnregisterObserver(ctx context.Context, observer observer.Observer) error
}

// The G2engine interface is a Golang representation of Senzing's libg2.h
type G2engine interface {
	AddRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error
	AddRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error)
	CloseExport(ctx context.Context, responseHandle uintptr) error
	CountRedoRecords(ctx context.Context) (int64, error)
	DeleteRecord(ctx context.Context, dataSourceCode string, recordID string, loadID string) error
	DeleteRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, loadID string, flags int64) (string, error)
	Destroy(ctx context.Context) error
	ExportConfig(ctx context.Context) (string, error)
	ExportConfigAndConfigID(ctx context.Context) (string, int64, error)
	ExportCSVEntityReport(ctx context.Context, csvColumnList string, flags int64) (uintptr, error)
	ExportCSVEntityReportIterator(ctx context.Context, csvColumnList string, flags int64) chan StringFragment
	ExportJSONEntityReport(ctx context.Context, flags int64) (uintptr, error)
	ExportJSONEntityReportIterator(ctx context.Context, flags int64) chan StringFragment
	FetchNext(ctx context.Context, responseHandle uintptr) (string, error)
	FindInterestingEntitiesByEntityID(ctx context.Context, entityID int64, flags int64) (string, error)
	FindInterestingEntitiesByRecordID(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	FindNetworkByEntityID_V2(ctx context.Context, entityList string, maxDegree int64, buildOutDegree int64, maxEntities int64, flags int64) (string, error)
	FindNetworkByEntityID(ctx context.Context, entityList string, maxDegree int64, buildOutDegree int64, maxEntities int64) (string, error)
	FindNetworkByRecordID_V2(ctx context.Context, recordList string, maxDegree int64, buildOutDegree int64, maxEntities int64, flags int64) (string, error)
	FindNetworkByRecordID(ctx context.Context, recordList string, maxDegree int64, buildOutDegree int64, maxEntities int64) (string, error)
	FindPathByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int64, flags int64) (string, error)
	FindPathByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int64) (string, error)
	FindPathByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int64, flags int64) (string, error)
	FindPathByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int64) (string, error)
	FindPathExcludingByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int64, excludedEntities string, flags int64) (string, error)
	FindPathExcludingByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int64, excludedEntities string) (string, error)
	FindPathExcludingByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int64, excludedRecords string, flags int64) (string, error)
	FindPathExcludingByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int64, excludedRecords string) (string, error)
	FindPathIncludingSourceByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int64, excludedEntities string, requiredDsrcs string, flags int64) (string, error)
	FindPathIncludingSourceByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int64, excludedEntities string, requiredDsrcs string) (string, error)
	FindPathIncludingSourceByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int64, excludedRecords string, requiredDsrcs string, flags int64) (string, error)
	FindPathIncludingSourceByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int64, excludedRecords string, requiredDsrcs string) (string, error)
	GetActiveConfigID(ctx context.Context) (int64, error)
	GetEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error)
	GetEntityByEntityID(ctx context.Context, entityID int64) (string, error)
	GetEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	GetEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error)
	GetObserverOrigin(ctx context.Context) string
	GetRecord_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	GetRecord(ctx context.Context, dataSourceCode string, recordID string) (string, error)
	GetRedoRecord(ctx context.Context) (string, error)
	GetRepositoryLastModifiedTime(ctx context.Context) (int64, error)
	GetSdkId(ctx context.Context) string
	GetVirtualEntityByRecordID_V2(ctx context.Context, recordList string, flags int64) (string, error)
	GetVirtualEntityByRecordID(ctx context.Context, recordList string) (string, error)
	HowEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error)
	HowEntityByEntityID(ctx context.Context, entityID int64) (string, error)
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int64) error
	InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int64) error
	PrimeEngine(ctx context.Context) error
	// ProcessRedoRecord(ctx context.Context) (string, error)
	// ProcessRedoRecordWithInfo(ctx context.Context, flags int64) (string, string, error)
	ReevaluateEntity(ctx context.Context, entityID int64, flags int64) error
	ReevaluateEntityWithInfo(ctx context.Context, entityID int64, flags int64) (string, error)
	ReevaluateRecord(ctx context.Context, dataSourceCode string, recordID string, flags int64) error
	ReevaluateRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	RegisterObserver(ctx context.Context, observer observer.Observer) error
	Reinit(ctx context.Context, initConfigID int64) error
	ReplaceRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error
	ReplaceRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error)
	SearchByAttributes_V2(ctx context.Context, jsonData string, flags int64) (string, error)
	SearchByAttributes(ctx context.Context, jsonData string) (string, error)
	SetLogLevel(ctx context.Context, logLevelName string) error
	SetObserverOrigin(ctx context.Context, origin string)
	Stats(ctx context.Context) (string, error)
	UnregisterObserver(ctx context.Context, observer observer.Observer) error
	WhyEntities_V2(ctx context.Context, entityID1 int64, entityID2 int64, flags int64) (string, error)
	WhyEntities(ctx context.Context, entityID1 int64, entityID2 int64) (string, error)
	WhyRecords_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, flags int64) (string, error)
	WhyRecords(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string) (string, error)
}

// The G2product interface is a Golang representation of Senzing's libg2product.h
type G2product interface {
	Destroy(ctx context.Context) error
	GetObserverOrigin(ctx context.Context) string
	GetSdkId(ctx context.Context) string
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int64) error
	License(ctx context.Context) (string, error)
	RegisterObserver(ctx context.Context, observer observer.Observer) error
	SetLogLevel(ctx context.Context, logLevelName string) error
	SetObserverOrigin(ctx context.Context, origin string)
	UnregisterObserver(ctx context.Context, observer observer.Observer) error
	Version(ctx context.Context) (string, error)
}
