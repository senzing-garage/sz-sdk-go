package senzing

import (
	"context"
	"encoding/json"

	"github.com/senzing-garage/sz-sdk-json-type-definition/go/typedef"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// --- Config -----------------------------------------------------------------

func UnmarshalSzConfigAddDataSourceResponse(ctx context.Context, jsonString string) (*typedef.SzConfigAddDataSourceResponse, error) {
	result := &typedef.SzConfigAddDataSourceResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzConfigExportConfigResponse(ctx context.Context, jsonString string) (*typedef.SzConfigExportConfigResponse, error) {
	result := &typedef.SzConfigExportConfigResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzConfigGetDataSourcesResponse(ctx context.Context, jsonString string) (*typedef.SzConfigGetDataSourcesResponse, error) {
	result := &typedef.SzConfigGetDataSourcesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- ConfigManager ----------------------------------------------------------

func UnmarshalSzConfigManagerGetConfigResponse(ctx context.Context, jsonString string) (*typedef.SzConfigManagerGetConfigResponse, error) {
	result := &typedef.SzConfigManagerGetConfigResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzConfigManagerGetConfigListResponse(ctx context.Context, jsonString string) (*typedef.SzConfigManagerGetConfigListResponse, error) {
	result := &typedef.SzConfigManagerGetConfigListResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Diagnostic -------------------------------------------------------------

func UnmarshalSzDiagnosticCheckDatastorePerformanceResponse(ctx context.Context, jsonString string) (*typedef.SzDiagnosticCheckDatastorePerformanceResponse, error) {
	result := &typedef.SzDiagnosticCheckDatastorePerformanceResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzDiagnosticGetDatastoreInfoResponse(ctx context.Context, jsonString string) (*typedef.SzDiagnosticGetDatastoreInfoResponse, error) {
	result := &typedef.SzDiagnosticGetDatastoreInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzDiagnosticGetFeatureResponse(ctx context.Context, jsonString string) (*typedef.SzDiagnosticGetFeatureResponse, error) {
	result := &typedef.SzDiagnosticGetFeatureResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Engine -----------------------------------------------------------------

func UnmarshalSzEngineAddRecordResponse(ctx context.Context, jsonString string) (*typedef.SzEngineAddRecordResponse, error) {
	result := &typedef.SzEngineAddRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineDeleteRecordResponse(ctx context.Context, jsonString string) (*typedef.SzEngineDeleteRecordResponse, error) {
	result := &typedef.SzEngineDeleteRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineFetchNextResponse(ctx context.Context, jsonString string) (*typedef.SzEngineFetchNextResponse, error) {
	result := &typedef.SzEngineFetchNextResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineFindInterestingEntitiesByEntityIdResponse(ctx context.Context, jsonString string) (*typedef.SzEngineFindInterestingEntitiesByEntityIDResponse, error) {
	result := &typedef.SzEngineFindInterestingEntitiesByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineFindInterestingEntitiesByRecordIdResponse(ctx context.Context, jsonString string) (*typedef.SzEngineFindInterestingEntitiesByRecordIDResponse, error) {
	result := &typedef.SzEngineFindInterestingEntitiesByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineFindNetworkByEntityIdResponse(ctx context.Context, jsonString string) (*typedef.SzEngineFindNetworkByEntityIDResponse, error) {
	result := &typedef.SzEngineFindNetworkByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineFindNetworkByRecordIdResponse(ctx context.Context, jsonString string) (*typedef.SzEngineFindNetworkByRecordIDResponse, error) {
	result := &typedef.SzEngineFindNetworkByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineFindPathByEntityIdResponse(ctx context.Context, jsonString string) (*typedef.SzEngineFindPathByEntityIDResponse, error) {
	result := &typedef.SzEngineFindPathByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineFindPathByRecordIdResponse(ctx context.Context, jsonString string) (*typedef.SzEngineFindPathByRecordIDResponse, error) {
	result := &typedef.SzEngineFindPathByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineGetEntityByEntityIdResponse(ctx context.Context, jsonString string) (*typedef.SzEngineGetEntityByEntityIDResponse, error) {
	result := &typedef.SzEngineGetEntityByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineGetEntityByRecordIdResponse(ctx context.Context, jsonString string) (*typedef.SzEngineGetEntityByRecordIDResponse, error) {
	result := &typedef.SzEngineGetEntityByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineGetRecordResponse(ctx context.Context, jsonString string) (*typedef.SzEngineGetRecordResponse, error) {
	result := &typedef.SzEngineGetRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineGetRedoRecordResponse(ctx context.Context, jsonString string) (*typedef.SzEngineGetRedoRecordResponse, error) {
	result := &typedef.SzEngineGetRedoRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineGetStatsResponse(ctx context.Context, jsonString string) (*typedef.SzEngineGetStatsResponse, error) {
	result := &typedef.SzEngineGetStatsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineGetVirtualEntityByRecordIdResponse(ctx context.Context, jsonString string) (*typedef.SzEngineGetVirtualEntityByRecordIDResponse, error) {
	result := &typedef.SzEngineGetVirtualEntityByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineHowEntityByEntityIdResponse(ctx context.Context, jsonString string) (*typedef.SzEngineHowEntityByEntityIDResponse, error) {
	result := &typedef.SzEngineHowEntityByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineProcessRedoRecordResponse(ctx context.Context, jsonString string) (*typedef.SzEngineProcessRedoRecordResponse, error) {
	result := &typedef.SzEngineProcessRedoRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineReevaluateEntityResponse(ctx context.Context, jsonString string) (*typedef.SzEngineReevaluateEntityResponse, error) {
	result := &typedef.SzEngineReevaluateEntityResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineReevaluateRecordResponse(ctx context.Context, jsonString string) (*typedef.SzEngineReevaluateRecordResponse, error) {
	result := &typedef.SzEngineReevaluateRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineSearchByAttributesResponse(ctx context.Context, jsonString string) (*typedef.SzEngineSearchByAttributesResponse, error) {
	result := &typedef.SzEngineSearchByAttributesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineStreamExportJsonEntityReportResponse(ctx context.Context, jsonString string) (*typedef.SzEngineStreamExportJSONEntityReportResponse, error) {
	result := &typedef.SzEngineStreamExportJSONEntityReportResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineWhyEntitiesResponse(ctx context.Context, jsonString string) (*typedef.SzEngineWhyEntitiesResponse, error) {
	result := &typedef.SzEngineWhyEntitiesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineWhyRecordInEntityResponse(ctx context.Context, jsonString string) (*typedef.SzEngineWhyRecordInEntityResponse, error) {
	result := &typedef.SzEngineWhyRecordInEntityResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzEngineWhyRecordsResponse(ctx context.Context, jsonString string) (*typedef.SzEngineWhyRecordsResponse, error) {
	result := &typedef.SzEngineWhyRecordsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Product ----------------------------------------------------------------

func UnmarshalSzProductGetLicenseResponse(ctx context.Context, jsonString string) (*typedef.SzProductGetLicenseResponse, error) {
	result := &typedef.SzProductGetLicenseResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzProductGetVersionResponse(ctx context.Context, jsonString string) (*typedef.SzProductGetVersionResponse, error) {
	result := &typedef.SzProductGetVersionResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}
