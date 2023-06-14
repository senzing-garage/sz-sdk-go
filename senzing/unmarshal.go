package senzing

import (
	"context"
	"encoding/json"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// --- Config -----------------------------------------------------------------

/*
UnmarshalConfigAddDataSourceResponse...
*/
func UnmarshalConfigAddDataSourceResponse(ctx context.Context, jsonString string) (*ConfigAddDataSourceResponse, error) {
	result := &ConfigAddDataSourceResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

/*
UnmarshalConfigListDataSourcesResponse...
*/
func UnmarshalConfigListDataSourcesResponse(ctx context.Context, jsonString string) (*ConfigListDataSourcesResponse, error) {
	result := &ConfigListDataSourcesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

/*
UnmarshalConfigSaveResponse...
*/
func UnmarshalConfigSaveResponse(ctx context.Context, jsonString string) (*ConfigSaveResponse, error) {
	result := &ConfigSaveResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Configmgr --------------------------------------------------------------

/*
UnmarshalConfigmgrGetConfigResponse...
*/
func UnmarshalConfigmgrGetConfigResponse(ctx context.Context, jsonString string) (*ConfigmgrGetConfigResponse, error) {
	result := &ConfigmgrGetConfigResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

/*
UnmarshalConfigmgrGetConfigListResponse...
*/
func UnmarshalConfigmgrGetConfigListResponse(ctx context.Context, jsonString string) (*ConfigmgrGetConfigListResponse, error) {
	result := &ConfigmgrGetConfigListResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Diagnostic -------------------------------------------------------------

func UnmarshalDiagnosticCheckDBPerfResponse(ctx context.Context, jsonString string) (*DiagnosticCheckDBPerfResponse, error) {
	return nil, nil
}

func UnmarshalDiagnosticFetchNextEntityBySizeResponse(ctx context.Context, jsonString string) (*DiagnosticFetchNextEntityBySizeResponse, error) {
	return nil, nil
}

func UnmarshalDiagnosticFindEntitiesByFeatureIDsResponse(ctx context.Context, jsonString string) (*DiagnosticFindEntitiesByFeatureIDsResponse, error) {
	return nil, nil
}
func UnmarshalDiagnosticGetDataSourceCountsResponse(ctx context.Context, jsonString string) (*DiagnosticGetDataSourceCountsResponse, error) {
	return nil, nil
}

func UnmarshalDiagnosticGetDBInfoResponse(ctx context.Context, jsonString string) (*DiagnosticGetDBInfoResponse, error) {
	return nil, nil
}
func UnmarshalDiagnosticGetEntityDetailsResponse(ctx context.Context, jsonString string) (*DiagnosticGetEntityDetailsResponse, error) {
	return nil, nil
}

func UnmarshalDiagnosticGetEntityListBySizeResponse(ctx context.Context, jsonString string) (*DiagnosticGetEntityListBySizeResponse, error) {
	return nil, nil
}
func UnmarshalDiagnosticGetEntityResumeResponse(ctx context.Context, jsonString string) (*DiagnosticGetEntityResumeResponse, error) {
	return nil, nil
}

func UnmarshalDiagnosticGetEntitySizeBreakdownResponse(ctx context.Context, jsonString string) (*DiagnosticGetEntitySizeBreakdownResponse, error) {
	return nil, nil
}
func UnmarshalDiagnosticGetFeatureResponse(ctx context.Context, jsonString string) (*DiagnosticGetFeatureResponse, error) {
	return nil, nil
}

func UnmarshalDiagnosticGetGenericFeaturesResponse(ctx context.Context, jsonString string) (*DiagnosticGetGenericFeaturesResponse, error) {
	return nil, nil
}
func UnmarshalDiagnosticGetMappingStatisticsResponse(ctx context.Context, jsonString string) (*DiagnosticGetMappingStatisticsResponse, error) {
	return nil, nil
}

func UnmarshalDiagnosticGetRelationshipDetailsResponse(ctx context.Context, jsonString string) (*DiagnosticGetRelationshipDetailsResponse, error) {
	return nil, nil
}
func UnmarshalDiagnosticGetResolutionStatisticsResponse(ctx context.Context, jsonString string) (*DiagnosticGetResolutionStatisticsResponse, error) {
	return nil, nil
}

func UnmarshalDiagnosticStreamEntityListBySizeResponse(ctx context.Context, jsonString string) (*DiagnosticStreamEntityListBySizeResponse, error) {
	return nil, nil
}

// --- Engine -----------------------------------------------------------------

func UnmarshalEngineAddRecordWithInfoResponse(ctx context.Context, jsonString string) (*EngineAddRecordWithInfoResponse, error) {
	return nil, nil
}

func UnmarshalEngineAddRecordWithInfoWithReturnedRecordIDResponse(ctx context.Context, jsonString string) (*EngineAddRecordWithInfoWithReturnedRecordIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineAddRecordWithReturnedRecordIDResponse(ctx context.Context, jsonString string) (*EngineAddRecordWithReturnedRecordIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineCheckRecordResponse(ctx context.Context, jsonString string) (*EngineCheckRecordResponse, error) {
	return nil, nil
}

func UnmarshalEngineDeleteRecordWithInfoResponse(ctx context.Context, jsonString string) (*EngineDeleteRecordWithInfoResponse, error) {
	return nil, nil
}

func UnmarshalEngineExportConfigAndConfigIDResponse(ctx context.Context, jsonString string) (*EngineExportConfigAndConfigIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineExportConfigResponse(ctx context.Context, jsonString string) (*EngineExportConfigResponse, error) {
	return nil, nil
}

func UnmarshalEngineFetchNextResponse(ctx context.Context, jsonString string) (*EngineFetchNextResponse, error) {
	return nil, nil
}

func UnmarshalEngineFindInterestingEntitiesByEntityIDResponse(ctx context.Context, jsonString string) (*EngineFindInterestingEntitiesByEntityIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineFindInterestingEntitiesByRecordIDResponse(ctx context.Context, jsonString string) (*EngineFindInterestingEntitiesByRecordIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineFindNetworkByEntityID_V2Response(ctx context.Context, jsonString string) (*EngineFindNetworkByEntityID_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineFindNetworkByEntityIDResponse(ctx context.Context, jsonString string) (*EngineFindNetworkByEntityIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineFindNetworkByRecordID_V2Response(ctx context.Context, jsonString string) (*EngineFindNetworkByRecordID_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineFindNetworkByRecordIDResponse(ctx context.Context, jsonString string) (*EngineFindNetworkByRecordIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineFindPathByEntityID_V2Response(ctx context.Context, jsonString string) (*EngineFindPathByEntityID_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineFindPathByEntityIDResponse(ctx context.Context, jsonString string) (*EngineFindPathByEntityIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineFindPathByRecordID_V2Response(ctx context.Context, jsonString string) (*EngineFindPathByRecordID_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineFindPathByRecordIDResponse(ctx context.Context, jsonString string) (*EngineFindPathByRecordIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineFindPathExcludingByEntityID_V2Response(ctx context.Context, jsonString string) (*EngineFindPathExcludingByEntityID_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineFindPathExcludingByEntityIDResponse(ctx context.Context, jsonString string) (*EngineFindPathExcludingByEntityIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineFindPathExcludingByRecordID_V2Response(ctx context.Context, jsonString string) (*EngineFindPathExcludingByRecordID_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineFindPathExcludingByRecordIDResponse(ctx context.Context, jsonString string) (*EngineFindPathExcludingByRecordIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineFindPathIncludingSourceByEntityID_V2Response(ctx context.Context, jsonString string) (*EngineFindPathIncludingSourceByEntityID_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineFindPathIncludingSourceByEntityIDResponse(ctx context.Context, jsonString string) (*EngineFindPathIncludingSourceByEntityIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineFindPathIncludingSourceByRecordID_V2Response(ctx context.Context, jsonString string) (*EngineFindPathIncludingSourceByRecordID_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineFindPathIncludingSourceByRecordIDResponse(ctx context.Context, jsonString string) (*EngineFindPathIncludingSourceByRecordIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineGetEntityByEntityID_V2Response(ctx context.Context, jsonString string) (*EngineGetEntityByEntityID_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineGetEntityByEntityIDResponse(ctx context.Context, jsonString string) (*EngineGetEntityByEntityIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineGetEntityByRecordID_V2Response(ctx context.Context, jsonString string) (*EngineGetEntityByRecordID_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineGetEntityByRecordIDResponse(ctx context.Context, jsonString string) (*EngineGetEntityByRecordIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineGetRecord_V2Response(ctx context.Context, jsonString string) (*EngineGetRecord_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineGetRecordResponse(ctx context.Context, jsonString string) (*EngineGetRecordResponse, error) {
	return nil, nil
}

func UnmarshalEngineGetRedoRecordResponse(ctx context.Context, jsonString string) (*EngineGetRedoRecordResponse, error) {
	return nil, nil
}

func UnmarshalEngineGetVirtualEntityByRecordID_V2Response(ctx context.Context, jsonString string) (*EngineGetVirtualEntityByRecordID_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineGetVirtualEntityByRecordIDResponse(ctx context.Context, jsonString string) (*EngineGetVirtualEntityByRecordIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineHowEntityByEntityID_V2Response(ctx context.Context, jsonString string) (*EngineHowEntityByEntityID_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineHowEntityByEntityIDResponse(ctx context.Context, jsonString string) (*EngineHowEntityByEntityIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineProcessRedoRecordResponse(ctx context.Context, jsonString string) (*EngineProcessRedoRecordResponse, error) {
	return nil, nil
}

func UnmarshalEngineProcessRedoRecordWithInfoResponse(ctx context.Context, jsonString string) (*EngineProcessRedoRecordWithInfoResponse, error) {
	return nil, nil
}

func UnmarshalEngineProcessWithInfoResponse(ctx context.Context, jsonString string) (*EngineProcessWithInfoResponse, error) {
	return nil, nil
}

func UnmarshalEngineProcessWithResponseResizeResponse(ctx context.Context, jsonString string) (*EngineProcessWithResponseResizeResponse, error) {
	return nil, nil
}

func UnmarshalEngineProcessWithResponseResponse(ctx context.Context, jsonString string) (*EngineProcessWithResponseResponse, error) {
	return nil, nil
}

func UnmarshalEngineReevaluateEntityWithInfoResponse(ctx context.Context, jsonString string) (*EngineReevaluateEntityWithInfoResponse, error) {
	return nil, nil
}
func UnmarshalEngineReevaluateRecordWithInfoResponse(ctx context.Context, jsonString string) (*EngineReevaluateRecordWithInfoResponse, error) {
	return nil, nil
}

func UnmarshalEngineReplaceRecordWithInfoResponse(ctx context.Context, jsonString string) (*EngineReplaceRecordWithInfoResponse, error) {
	return nil, nil
}

func UnmarshalEngineSearchByAttributes_V2Response(ctx context.Context, jsonString string) (*EngineSearchByAttributes_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineSearchByAttributesResponse(ctx context.Context, jsonString string) (*EngineSearchByAttributesResponse, error) {
	return nil, nil
}

func UnmarshalEngineStatsResponse(ctx context.Context, jsonString string) (*EngineStatsResponse, error) {
	return nil, nil
}

func UnmarshalEngineStreamExportJSONEntityReportResponse(ctx context.Context, jsonString string) (*EngineStreamExportJSONEntityReportResponse, error) {
	return nil, nil
}

func UnmarshalEngineWhyEntities_V2Response(ctx context.Context, jsonString string) (*EngineWhyEntities_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineWhyEntitiesResponse(ctx context.Context, jsonString string) (*EngineWhyEntitiesResponse, error) {
	return nil, nil
}

func UnmarshalEngineWhyEntityByEntityID_V2Response(ctx context.Context, jsonString string) (*EngineWhyEntityByEntityID_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineWhyEntityByEntityIDResponse(ctx context.Context, jsonString string) (*EngineWhyEntityByEntityIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineWhyEntityByRecordID_V2Response(ctx context.Context, jsonString string) (*EngineWhyEntityByRecordID_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineWhyEntityByRecordIDResponse(ctx context.Context, jsonString string) (*EngineWhyEntityByRecordIDResponse, error) {
	return nil, nil
}

func UnmarshalEngineWhyRecords_V2Response(ctx context.Context, jsonString string) (*EngineWhyRecords_V2Response, error) {
	return nil, nil
}

func UnmarshalEngineWhyRecordsResponse(ctx context.Context, jsonString string) (*EngineWhyRecordsResponse, error) {
	return nil, nil
}

// --- Product ----------------------------------------------------------------

/*
UnmarshalProductVersionResponse...
*/
func UnmarshalProductLicenseResponse(ctx context.Context, jsonString string) (*ProductLicenseResponse, error) {
	result := &ProductLicenseResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalProductValidateLicenseFileResponse(ctx context.Context, jsonString string) (*ProductValidateLicenseFileResponse, error) {
	return nil, nil
}

func UnmarshalProductValidateLicenseStringBase64Response(ctx context.Context, jsonString string) (*ProductValidateLicenseStringBase64Response, error) {
	return nil, nil
}

/*
UnmarshalProductVersionResponse...
*/
func UnmarshalProductVersionResponse(ctx context.Context, jsonString string) (*ProductVersionResponse, error) {
	result := &ProductVersionResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}
