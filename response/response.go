package response

import (
	"context"
	"encoding/json"

	"github.com/senzing-garage/sz-sdk-json-type-definition/go/typedef"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// --- Config -----------------------------------------------------------------

func SzConfigAddDataSource(ctx context.Context, jsonString string) (*typedef.SzConfigAddDataSourceResponse, error) {
	_ = ctx
	result := &typedef.SzConfigAddDataSourceResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzConfigExportConfig(ctx context.Context, jsonString string) (*typedef.SzConfigExportConfigResponse, error) {
	_ = ctx
	result := &typedef.SzConfigExportConfigResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzConfigGetDataSources(ctx context.Context, jsonString string) (*typedef.SzConfigGetDataSourcesResponse, error) {
	_ = ctx
	result := &typedef.SzConfigGetDataSourcesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- ConfigManager ----------------------------------------------------------

func SzConfigManagerGetConfig(ctx context.Context, jsonString string) (*typedef.SzConfigManagerGetConfigResponse, error) {
	_ = ctx
	result := &typedef.SzConfigManagerGetConfigResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzConfigManagerGetConfigList(ctx context.Context, jsonString string) (*typedef.SzConfigManagerGetConfigListResponse, error) {
	_ = ctx
	result := &typedef.SzConfigManagerGetConfigListResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Diagnostic -------------------------------------------------------------

func SzDiagnosticCheckDatastorePerformance(ctx context.Context, jsonString string) (*typedef.SzDiagnosticCheckDatastorePerformanceResponse, error) {
	_ = ctx
	result := &typedef.SzDiagnosticCheckDatastorePerformanceResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzDiagnosticGetDatastoreInfo(ctx context.Context, jsonString string) (*typedef.SzDiagnosticGetDatastoreInfoResponse, error) {
	_ = ctx
	result := &typedef.SzDiagnosticGetDatastoreInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzDiagnosticGetFeature(ctx context.Context, jsonString string) (*typedef.SzDiagnosticGetFeatureResponse, error) {
	_ = ctx
	result := &typedef.SzDiagnosticGetFeatureResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Engine -----------------------------------------------------------------

func SzEngineAddRecord(ctx context.Context, jsonString string) (*typedef.SzEngineAddRecordResponse, error) {
	_ = ctx
	result := &typedef.SzEngineAddRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineDeleteRecord(ctx context.Context, jsonString string) (*typedef.SzEngineDeleteRecordResponse, error) {
	_ = ctx
	result := &typedef.SzEngineDeleteRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineFetchNext(ctx context.Context, jsonString string) (*typedef.SzEngineFetchNextResponse, error) {
	_ = ctx
	result := &typedef.SzEngineFetchNextResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineFindInterestingEntitiesByEntityID(ctx context.Context, jsonString string) (*typedef.SzEngineFindInterestingEntitiesByEntityIDResponse, error) {
	_ = ctx
	result := &typedef.SzEngineFindInterestingEntitiesByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineFindInterestingEntitiesByRecordID(ctx context.Context, jsonString string) (*typedef.SzEngineFindInterestingEntitiesByRecordIDResponse, error) {
	_ = ctx
	result := &typedef.SzEngineFindInterestingEntitiesByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineFindNetworkByEntityID(ctx context.Context, jsonString string) (*typedef.SzEngineFindNetworkByEntityIDResponse, error) {
	_ = ctx
	result := &typedef.SzEngineFindNetworkByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineFindNetworkByRecordID(ctx context.Context, jsonString string) (*typedef.SzEngineFindNetworkByRecordIDResponse, error) {
	_ = ctx
	result := &typedef.SzEngineFindNetworkByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineFindPathByEntityID(ctx context.Context, jsonString string) (*typedef.SzEngineFindPathByEntityIDResponse, error) {
	_ = ctx
	result := &typedef.SzEngineFindPathByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineFindPathByRecordID(ctx context.Context, jsonString string) (*typedef.SzEngineFindPathByRecordIDResponse, error) {
	_ = ctx
	result := &typedef.SzEngineFindPathByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineGetEntityByEntityID(ctx context.Context, jsonString string) (*typedef.SzEngineGetEntityByEntityIDResponse, error) {
	_ = ctx
	result := &typedef.SzEngineGetEntityByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineGetEntityByRecordID(ctx context.Context, jsonString string) (*typedef.SzEngineGetEntityByRecordIDResponse, error) {
	_ = ctx
	result := &typedef.SzEngineGetEntityByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineGetRecord(ctx context.Context, jsonString string) (*typedef.SzEngineGetRecordResponse, error) {
	_ = ctx
	result := &typedef.SzEngineGetRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineGetRedoRecord(ctx context.Context, jsonString string) (*typedef.SzEngineGetRedoRecordResponse, error) {
	_ = ctx
	result := &typedef.SzEngineGetRedoRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineGetStats(ctx context.Context, jsonString string) (*typedef.SzEngineGetStatsResponse, error) {
	_ = ctx
	result := &typedef.SzEngineGetStatsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineGetVirtualEntityByRecordID(ctx context.Context, jsonString string) (*typedef.SzEngineGetVirtualEntityByRecordIDResponse, error) {
	_ = ctx
	result := &typedef.SzEngineGetVirtualEntityByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineHowEntityByEntityID(ctx context.Context, jsonString string) (*typedef.SzEngineHowEntityByEntityIDResponse, error) {
	_ = ctx
	result := &typedef.SzEngineHowEntityByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineProcessRedoRecord(ctx context.Context, jsonString string) (*typedef.SzEngineProcessRedoRecordResponse, error) {
	_ = ctx
	result := &typedef.SzEngineProcessRedoRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineReevaluateEntity(ctx context.Context, jsonString string) (*typedef.SzEngineReevaluateEntityResponse, error) {
	_ = ctx
	result := &typedef.SzEngineReevaluateEntityResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineReevaluateRecord(ctx context.Context, jsonString string) (*typedef.SzEngineReevaluateRecordResponse, error) {
	_ = ctx
	result := &typedef.SzEngineReevaluateRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineSearchByAttributes(ctx context.Context, jsonString string) (*typedef.SzEngineSearchByAttributesResponse, error) {
	_ = ctx
	result := &typedef.SzEngineSearchByAttributesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineStreamExportJSONEntityReport(ctx context.Context, jsonString string) (*typedef.SzEngineStreamExportJSONEntityReportResponse, error) {
	_ = ctx
	result := &typedef.SzEngineStreamExportJSONEntityReportResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineWhyEntities(ctx context.Context, jsonString string) (*typedef.SzEngineWhyEntitiesResponse, error) {
	_ = ctx
	result := &typedef.SzEngineWhyEntitiesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineWhyRecordInEntity(ctx context.Context, jsonString string) (*typedef.SzEngineWhyRecordInEntityResponse, error) {
	_ = ctx
	result := &typedef.SzEngineWhyRecordInEntityResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzEngineWhyRecords(ctx context.Context, jsonString string) (*typedef.SzEngineWhyRecordsResponse, error) {
	_ = ctx
	result := &typedef.SzEngineWhyRecordsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Product ----------------------------------------------------------------

func SzProductGetLicense(ctx context.Context, jsonString string) (*typedef.SzProductGetLicenseResponse, error) {
	_ = ctx
	result := &typedef.SzProductGetLicenseResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func SzProductGetVersion(ctx context.Context, jsonString string) (*typedef.SzProductGetVersionResponse, error) {
	_ = ctx
	result := &typedef.SzProductGetVersionResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}
