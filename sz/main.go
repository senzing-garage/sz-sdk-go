package sz

import (
	"context"
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

// The SzAbstractFactory interface is the interface for all Senzing factories in the Abstract Factory pattern
type SzAbstractFactory interface {
	CreateConfig(ctx context.Context) (SzConfig, error)
	CreateConfigManager(ctx context.Context) (SzConfigManager, error)
	CreateDiagnostic(ctx context.Context) (SzDiagnostic, error)
	CreateEngine(ctx context.Context) (SzEngine, error)
	CreateProduct(ctx context.Context) (SzProduct, error)
}

// The SzConfig interface is a Golang representation of Senzing's libg2config.h
type SzConfig interface {
	AddDataSource(ctx context.Context, configHandle uintptr, dataSourceCode string) (string, error)
	CloseConfig(ctx context.Context, configHandle uintptr) error
	CreateConfig(ctx context.Context) (uintptr, error)
	DeleteDataSource(ctx context.Context, configHandle uintptr, dataSourceCode string) error
	Destroy(ctx context.Context) error
	ExportConfig(ctx context.Context, configHandle uintptr) (string, error)
	GetDataSources(ctx context.Context, configHandle uintptr) (string, error)
	ImportConfig(ctx context.Context, configDefinition string) (uintptr, error)
	Initialize(ctx context.Context, instanceName string, settings string, verboseLogging int64) error
}

// The SzConfigManager interface is a Golang representation of Senzing's libg2configmgr.h
type SzConfigManager interface {
	AddConfig(ctx context.Context, configDefinition string, configComments string) (int64, error)
	Destroy(ctx context.Context) error
	GetConfig(ctx context.Context, configId int64) (string, error)
	GetConfigList(ctx context.Context) (string, error)
	GetDefaultConfigId(ctx context.Context) (int64, error)
	Initialize(ctx context.Context, instanceName string, settings string, verboseLogging int64) error
	ReplaceDefaultConfigId(ctx context.Context, currentDefaultConfigId int64, newDefaultConfigId int64) error
	SetDefaultConfigId(ctx context.Context, configId int64) error
}

// The SzDiagnostic interface is a Golang representation of Senzing's libg2diagnostic.h
type SzDiagnostic interface {
	CheckDatabasePerformance(ctx context.Context, secondsToRun int) (string, error)
	Destroy(ctx context.Context) error
	Initialize(ctx context.Context, instanceName string, settings string, verboseLogging int64, configId int64) error
	PurgeRepository(ctx context.Context) error
	Reinitialize(ctx context.Context, configId int64) error
}

// The SzEngine interface is a Golang representation of Senzing's libg2.h
type SzEngine interface {
	AddRecord(ctx context.Context, dataSourceCode string, recordId string, recordDefinition string, flags int64) (string, error)
	CloseExport(ctx context.Context, exportHandle uintptr) error
	CountRedoRecords(ctx context.Context) (int64, error)
	DeleteRecord(ctx context.Context, dataSourceCode string, recordId string, flags int64) (string, error)
	Destroy(ctx context.Context) error
	ExportCsvEntityReport(ctx context.Context, csvColumnList string, flags int64) (uintptr, error)
	ExportCsvEntityReportIterator(ctx context.Context, csvColumnList string, flags int64) chan StringFragment
	ExportJsonEntityReport(ctx context.Context, flags int64) (uintptr, error)
	ExportJsonEntityReportIterator(ctx context.Context, flags int64) chan StringFragment
	FetchNext(ctx context.Context, exportHandle uintptr) (string, error)
	FindNetworkByEntityId(ctx context.Context, entityList string, maxDegrees int64, buildOutDegree int64, maxEntities int64, flags int64) (string, error)
	FindNetworkByRecordId(ctx context.Context, recordList string, maxDegrees int64, buildOutDegree int64, maxEntities int64, flags int64) (string, error)
	FindPathByEntityId(ctx context.Context, startEntityId int64, endEntityId int64, maxDegrees int64, exclusions string, requiredDataSources string, flags int64) (string, error)
	FindPathByRecordId(ctx context.Context, startDataSourceCode string, startRecordId string, endDataSourceCode string, endRecordId string, maxDegrees int64, exclusions string, requiredDataSources string, flags int64) (string, error)
	GetActiveConfigId(ctx context.Context) (int64, error)
	GetEntityByEntityId(ctx context.Context, entityId int64, flags int64) (string, error)
	GetEntityByRecordId(ctx context.Context, dataSourceCode string, recordId string, flags int64) (string, error)
	GetRecord(ctx context.Context, dataSourceCode string, recordId string, flags int64) (string, error)
	GetRedoRecord(ctx context.Context) (string, error)
	GetRepositoryLastModifiedTime(ctx context.Context) (int64, error)
	GetStats(ctx context.Context) (string, error)
	GetVirtualEntityByRecordId(ctx context.Context, recordList string, flags int64) (string, error)
	HowEntityByEntityId(ctx context.Context, entityId int64, flags int64) (string, error)
	Initialize(ctx context.Context, instanceName string, settings string, verboseLogging int64, configId int64) error
	PrimeEngine(ctx context.Context) error
	ProcessRedoRecord(ctx context.Context, redoRecord string, flags int64) (string, error)
	ReevaluateEntity(ctx context.Context, entityId int64, flags int64) (string, error)
	ReevaluateRecord(ctx context.Context, dataSourceCode string, recordId string, flags int64) (string, error)
	Reinitialize(ctx context.Context, configId int64) error
	ReplaceRecord(ctx context.Context, dataSourceCode string, recordId string, recordDefinition string, flags int64) (string, error)
	SearchByAttributes(ctx context.Context, attributes string, searchProfile string, flags int64) (string, error)
	WhyEntities(ctx context.Context, entityId1 int64, entityId2 int64, flags int64) (string, error)
	WhyRecordInEntity(ctx context.Context, dataSourceCode string, recordId string, flags int64) (string, error)
	WhyRecords(ctx context.Context, dataSourceCode1 string, recordId1 string, dataSourceCode2 string, recordId2 string, flags int64) (string, error)
}

// The SzProduct interface is a Golang representation of Senzing's libg2product.h
type SzProduct interface {
	Destroy(ctx context.Context) error
	GetLicense(ctx context.Context) (string, error)
	GetVersion(ctx context.Context) (string, error)
	Initialize(ctx context.Context, instanceName string, settings string, verboseLogging int64) error
}
