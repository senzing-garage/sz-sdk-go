package senzing

import (
	"context"
)

// ----------------------------------------------------------------------------
// Types - struct
// ----------------------------------------------------------------------------

// Type StringFragment struct is used as a return value when iterating over log strings.
type StringFragment struct {
	Error error
	Value string
}

// ----------------------------------------------------------------------------
// Types - interface
// ----------------------------------------------------------------------------

/*
Type SzAbstractFactory interface is the interface for Senzing factories following the [Abstract Factory pattern].

[Abstract Factory pattern]: https://en.wikipedia.org/wiki/Abstract_factory_pattern
*/
type SzAbstractFactory interface {
	CreateConfigManager(ctx context.Context) (SzConfigManager, error)
	CreateDiagnostic(ctx context.Context) (SzDiagnostic, error)
	CreateEngine(ctx context.Context) (SzEngine, error)
	CreateProduct(ctx context.Context) (SzProduct, error)
	Destroy(ctx context.Context) error
	Reinitialize(ctx context.Context, configID int64) error
}

// Type SzConfig interface is a Golang representation of Senzing's libSzconfig.h.
type SzConfig interface {
	AddDataSource(ctx context.Context, dataSourceCode string) (string, error)
	DeleteDataSource(ctx context.Context, dataSourceCode string) (string, error)
	Export(ctx context.Context) (string, error)
	GetDataSources(ctx context.Context) (string, error)
}

// Type SzConfigManager interface is a Golang representation of Senzing's libSzconfigmgr.h.
type SzConfigManager interface {
	CreateConfigFromConfigID(ctx context.Context, configID int64) (SzConfig, error)
	CreateConfigFromString(ctx context.Context, configDefinition string) (SzConfig, error)
	CreateConfigFromTemplate(ctx context.Context) (SzConfig, error)
	GetConfigs(ctx context.Context) (string, error)
	GetDefaultConfigID(ctx context.Context) (int64, error)
	RegisterConfig(ctx context.Context, configDefinition string, configComment string) (int64, error)
	ReplaceDefaultConfigID(ctx context.Context, currentDefaultConfigID int64, newDefaultConfigID int64) error
	SetDefaultConfig(ctx context.Context, configDefinition string, configComment string) (int64, error)
	SetDefaultConfigID(ctx context.Context, configID int64) error
}

// Type SzDiagnostic interface is a Golang representation of Senzing's libSzdiagnostic.h.
type SzDiagnostic interface {
	CheckDatastorePerformance(ctx context.Context, secondsToRun int) (string, error)
	GetDatastoreInfo(ctx context.Context) (string, error)
	GetFeature(ctx context.Context, featureID int64) (string, error)
	PurgeRepository(ctx context.Context) error
}

// Type SzEngine interface is a Golang representation of Senzing's libSz.h.
type SzEngine interface {
	AddRecord(
		ctx context.Context,
		dataSourceCode string,
		recordID string,
		recordDefinition string,
		flags int64) (string, error)
	CloseExport(ctx context.Context, exportHandle uintptr) error
	CountRedoRecords(ctx context.Context) (int64, error)
	DeleteRecord(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	ExportCsvEntityReport(ctx context.Context, csvColumnList string, flags int64) (uintptr, error)
	ExportCsvEntityReportIterator(ctx context.Context, csvColumnList string, flags int64) chan StringFragment
	ExportJSONEntityReport(ctx context.Context, flags int64) (uintptr, error)
	ExportJSONEntityReportIterator(ctx context.Context, flags int64) chan StringFragment
	FetchNext(ctx context.Context, exportHandle uintptr) (string, error)
	FindInterestingEntitiesByEntityID(
		ctx context.Context,
		entityID int64,
		flags int64) (string, error)
	FindInterestingEntitiesByRecordID(
		ctx context.Context,
		dataSourceCode string,
		recordID string,
		flags int64) (string, error)
	FindNetworkByEntityID(
		ctx context.Context,
		entityIDs string,
		maxDegrees int64,
		buildOutDegrees int64,
		buildOutMaxEntities int64,
		flags int64) (string, error)
	FindNetworkByRecordID(
		ctx context.Context,
		recordKeys string,
		maxDegrees int64,
		buildOutDegrees int64,
		buildOutMaxEntities int64,
		flags int64) (string, error)
	FindPathByEntityID(
		ctx context.Context,
		startEntityID int64,
		endEntityID int64,
		maxDegrees int64,
		avoidEntityIDs string,
		requiredDataSources string,
		flags int64) (string, error)
	FindPathByRecordID(
		ctx context.Context,
		startDataSourceCode string,
		startRecordID string,
		endDataSourceCode string,
		endRecordID string,
		maxDegrees int64,
		avoidRecordKeys string,
		requiredDataSources string,
		flags int64) (string, error)
	GetActiveConfigID(ctx context.Context) (int64, error)
	GetEntityByEntityID(ctx context.Context, entityID int64, flags int64) (string, error)
	GetEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	GetRecord(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	GetRedoRecord(ctx context.Context) (string, error)
	GetStats(ctx context.Context) (string, error)
	GetVirtualEntityByRecordID(ctx context.Context, recordKeys string, flags int64) (string, error)
	HowEntityByEntityID(ctx context.Context, entityID int64, flags int64) (string, error)
	PreprocessRecord(ctx context.Context, recordDefinition string, flags int64) (string, error)
	PrimeEngine(ctx context.Context) error
	ProcessRedoRecord(ctx context.Context, redoRecord string, flags int64) (string, error)
	ReevaluateEntity(ctx context.Context, entityID int64, flags int64) (string, error)
	ReevaluateRecord(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	SearchByAttributes(ctx context.Context, attributes string, searchProfile string, flags int64) (string, error)
	WhyEntities(ctx context.Context, entityID1 int64, entityID2 int64, flags int64) (string, error)
	WhyRecordInEntity(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	WhyRecords(
		ctx context.Context,
		dataSourceCode1 string,
		recordID1 string,
		dataSourceCode2 string,
		recordID2 string,
		flags int64) (string, error)
	WhySearch(ctx context.Context, attributes string, entityID int64, searchProfile string, flags int64) (string, error)
}

// Type SzProduct interface is a Golang representation of Senzing's libSzproduct.h.
type SzProduct interface {
	GetLicense(ctx context.Context) (string, error)
	GetVersion(ctx context.Context) (string, error)
}
