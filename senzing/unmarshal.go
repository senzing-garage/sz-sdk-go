package senzing

import (
	"context"
	"encoding/json"

	"github.com/senzing/g2-sdk-json-type-definition/go/typedef"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// --- Config -----------------------------------------------------------------

func UnmarshalConfigAddDataSourceResponse(ctx context.Context, jsonString string) (*typedef.ConfigAddDataSourceResponse, error) {
	result := &typedef.ConfigAddDataSourceResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalConfigListDataSourcesResponse(ctx context.Context, jsonString string) (*typedef.ConfigListDataSourcesResponse, error) {
	result := &typedef.ConfigListDataSourcesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalConfigSaveResponse(ctx context.Context, jsonString string) (*typedef.ConfigSaveResponse, error) {
	result := &typedef.ConfigSaveResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Configmgr --------------------------------------------------------------

func UnmarshalConfigmgrGetConfigResponse(ctx context.Context, jsonString string) (*typedef.ConfigmgrGetConfigResponse, error) {
	result := &typedef.ConfigmgrGetConfigResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalConfigmgrGetConfigListResponse(ctx context.Context, jsonString string) (*typedef.ConfigmgrGetConfigListResponse, error) {
	result := &typedef.ConfigmgrGetConfigListResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Diagnostic -------------------------------------------------------------

func UnmarshalDiagnosticCheckDBPerfResponse(ctx context.Context, jsonString string) (*typedef.DiagnosticCheckDbperfResponse, error) {
	result := &typedef.DiagnosticCheckDbperfResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticFetchNextEntityBySizeResponse(ctx context.Context, jsonString string) (*typedef.DiagnosticFetchNextEntityBySizeResponse, error) {
	result := &typedef.DiagnosticFetchNextEntityBySizeResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticFindEntitiesByFeatureIDsResponse(ctx context.Context, jsonString string) (*typedef.DiagnosticFindEntitiesByFeatureIdsResponse, error) {
	result := &typedef.DiagnosticFindEntitiesByFeatureIdsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetDataSourceCountsResponse(ctx context.Context, jsonString string) (*typedef.DiagnosticGetDataSourceCountsResponse, error) {
	result := &typedef.DiagnosticGetDataSourceCountsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetDBInfoResponse(ctx context.Context, jsonString string) (*typedef.DiagnosticGetDbinfoResponse, error) {
	result := &typedef.DiagnosticGetDbinfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetEntityDetailsResponse(ctx context.Context, jsonString string) (*typedef.DiagnosticGetEntityDetailsResponse, error) {
	result := &typedef.DiagnosticGetEntityDetailsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalDiagnosticGetEntityListBySizeResponse(ctx context.Context, jsonString string) (*typedef.DiagnosticGetEntityListBySizeResponse, error) {
// 	result := &typedef.DiagnosticGetEntityListBySizeResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalDiagnosticGetEntityResumeResponse(ctx context.Context, jsonString string) (*typedef.DiagnosticGetEntityResumeResponse, error) {
	result := &typedef.DiagnosticGetEntityResumeResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetEntitySizeBreakdownResponse(ctx context.Context, jsonString string) (*typedef.DiagnosticGetEntitySizeBreakdownResponse, error) {
	result := &typedef.DiagnosticGetEntitySizeBreakdownResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetFeatureResponse(ctx context.Context, jsonString string) (*typedef.DiagnosticGetFeatureResponse, error) {
	result := &typedef.DiagnosticGetFeatureResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetGenericFeaturesResponse(ctx context.Context, jsonString string) (*typedef.DiagnosticGetGenericFeaturesResponse, error) {
	result := &typedef.DiagnosticGetGenericFeaturesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetMappingStatisticsResponse(ctx context.Context, jsonString string) (*typedef.DiagnosticGetMappingStatisticsResponse, error) {
	result := &typedef.DiagnosticGetMappingStatisticsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetRelationshipDetailsResponse(ctx context.Context, jsonString string) (*typedef.DiagnosticGetRelationshipDetailsResponse, error) {
	result := &typedef.DiagnosticGetRelationshipDetailsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetResolutionStatisticsResponse(ctx context.Context, jsonString string) (*typedef.DiagnosticGetResolutionStatisticsResponse, error) {
	result := &typedef.DiagnosticGetResolutionStatisticsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalDiagnosticStreamEntityListBySizeResponse(ctx context.Context, jsonString string) (*typedef.DiagnosticStreamEntityListBySizeResponse, error) {
// 	result := &typedef.DiagnosticStreamEntityListBySizeResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

// --- Engine -----------------------------------------------------------------

func UnmarshalEngineAddRecordWithInfoResponse(ctx context.Context, jsonString string) (*typedef.EngineAddRecordWithInfoResponse, error) {
	result := &typedef.EngineAddRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineAddRecordWithInfoWithReturnedRecordIDResponse(ctx context.Context, jsonString string) (*typedef.EngineAddRecordWithInfoWithReturnedRecordIdresponse, error) {
	result := &typedef.EngineAddRecordWithInfoWithReturnedRecordIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalEngineAddRecordWithReturnedRecordIDResponse(ctx context.Context, jsonString string) (*typedef.EngineAddRecordWithReturnedRecordIDResponse, error) {
// 	result := &typedef.EngineAddRecordWithReturnedRecordIDResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalEngineCheckRecordResponse(ctx context.Context, jsonString string) (*typedef.EngineCheckRecordResponse, error) {
	result := &typedef.EngineCheckRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineDeleteRecordWithInfoResponse(ctx context.Context, jsonString string) (*typedef.EngineDeleteRecordWithInfoResponse, error) {
	result := &typedef.EngineDeleteRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineExportConfigAndConfigIDResponse(ctx context.Context, jsonString string) (*typedef.EngineExportConfigAndConfigIdresponse, error) {
	result := &typedef.EngineExportConfigAndConfigIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineExportConfigResponse(ctx context.Context, jsonString string) (*typedef.EngineExportConfigResponse, error) {
	result := &typedef.EngineExportConfigResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalEngineFetchNextResponse(ctx context.Context, jsonString string) (*typedef.EngineFetchNextResponse, error) {
// 	result := &typedef.EngineFetchNextResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalEngineFindInterestingEntitiesByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.EngineFindInterestingEntitiesByEntityIdresponse, error) {
	result := &typedef.EngineFindInterestingEntitiesByEntityIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindInterestingEntitiesByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.EngineFindInterestingEntitiesByRecordIdresponse, error) {
	result := &typedef.EngineFindInterestingEntitiesByRecordIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindNetworkByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.EngineFindNetworkByEntityIdv2response, error) {
	result := &typedef.EngineFindNetworkByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindNetworkByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.EngineFindNetworkByEntityIdresponse, error) {
	result := &typedef.EngineFindNetworkByEntityIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindNetworkByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.EngineFindNetworkByRecordIdv2response, error) {
	result := &typedef.EngineFindNetworkByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindNetworkByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.EngineFindNetworkByRecordIdresponse, error) {
	result := &typedef.EngineFindNetworkByRecordIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.EngineFindPathByEntityIdv2response, error) {
	result := &typedef.EngineFindPathByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.EngineFindPathByEntityIdresponse, error) {
	result := &typedef.EngineFindPathByEntityIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.EngineFindPathByRecordIdv2response, error) {
	result := &typedef.EngineFindPathByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.EngineFindPathByRecordIdresponse, error) {
	result := &typedef.EngineFindPathByRecordIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathExcludingByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.EngineFindPathExcludingByEntityIdv2response, error) {
	result := &typedef.EngineFindPathExcludingByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathExcludingByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.EngineFindPathExcludingByEntityIdresponse, error) {
	result := &typedef.EngineFindPathExcludingByEntityIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathExcludingByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.EngineFindPathExcludingByRecordIdv2response, error) {
	result := &typedef.EngineFindPathExcludingByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathExcludingByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.EngineFindPathExcludingByRecordIdresponse, error) {
	result := &typedef.EngineFindPathExcludingByRecordIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathIncludingSourceByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.EngineFindPathIncludingSourceByEntityIdv2response, error) {
	result := &typedef.EngineFindPathIncludingSourceByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathIncludingSourceByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.EngineFindPathIncludingSourceByEntityIdresponse, error) {
	result := &typedef.EngineFindPathIncludingSourceByEntityIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathIncludingSourceByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.EngineFindPathIncludingSourceByRecordIdv2response, error) {
	result := &typedef.EngineFindPathIncludingSourceByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathIncludingSourceByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.EngineFindPathIncludingSourceByRecordIdresponse, error) {
	result := &typedef.EngineFindPathIncludingSourceByRecordIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetEntityByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.EngineGetEntityByEntityIdv2response, error) {
	result := &typedef.EngineGetEntityByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetEntityByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.EngineGetEntityByEntityIdresponse, error) {
	result := &typedef.EngineGetEntityByEntityIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetEntityByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.EngineGetEntityByRecordIdv2response, error) {
	result := &typedef.EngineGetEntityByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetEntityByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.EngineGetEntityByRecordIdresponse, error) {
	result := &typedef.EngineGetEntityByRecordIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetRecordV2Response(ctx context.Context, jsonString string) (*typedef.EngineGetRecordV2response, error) {
	result := &typedef.EngineGetRecordV2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetRecordResponse(ctx context.Context, jsonString string) (*typedef.EngineGetRecordResponse, error) {
	result := &typedef.EngineGetRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalEngineGetRedoRecordResponse(ctx context.Context, jsonString string) (*typedef.EngineGetRedoRecordResponse, error) {
// 	result := &typedef.EngineGetRedoRecordResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalEngineGetVirtualEntityByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.EngineGetVirtualEntityByRecordIdv2response, error) {
	result := &typedef.EngineGetVirtualEntityByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetVirtualEntityByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.EngineGetVirtualEntityByRecordIdresponse, error) {
	result := &typedef.EngineGetVirtualEntityByRecordIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineHowEntityByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.EngineHowEntityByEntityIdv2response, error) {
	result := &typedef.EngineHowEntityByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineHowEntityByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.EngineHowEntityByEntityIdresponse, error) {
	result := &typedef.EngineHowEntityByEntityIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalEngineProcessRedoRecordResponse(ctx context.Context, jsonString string) (*typedef.EngineProcessRedoRecordResponse, error) {
// 	result := &typedef.EngineProcessRedoRecordResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalEngineProcessRedoRecordWithInfoResponse(ctx context.Context, jsonString string) (*typedef.EngineProcessRedoRecordWithInfoResponse, error) {
	result := &typedef.EngineProcessRedoRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineProcessWithInfoResponse(ctx context.Context, jsonString string) (*typedef.EngineProcessWithInfoResponse, error) {
	result := &typedef.EngineProcessWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineProcessWithResponseResizeResponse(ctx context.Context, jsonString string) (*typedef.EngineProcessWithResponseResizeResponse, error) {
	result := &typedef.EngineProcessWithResponseResizeResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineProcessWithResponseResponse(ctx context.Context, jsonString string) (*typedef.EngineProcessWithResponseResponse, error) {
	result := &typedef.EngineProcessWithResponseResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineReevaluateEntityWithInfoResponse(ctx context.Context, jsonString string) (*typedef.EngineReevaluateEntityWithInfoResponse, error) {
	result := &typedef.EngineReevaluateEntityWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineReevaluateRecordWithInfoResponse(ctx context.Context, jsonString string) (*typedef.EngineReevaluateRecordWithInfoResponse, error) {
	result := &typedef.EngineReevaluateRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineReplaceRecordWithInfoResponse(ctx context.Context, jsonString string) (*typedef.EngineReplaceRecordWithInfoResponse, error) {
	result := &typedef.EngineReplaceRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineSearchByAttributesV2Response(ctx context.Context, jsonString string) (*typedef.EngineSearchByAttributesV2response, error) {
	result := &typedef.EngineSearchByAttributesV2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineSearchByAttributesV3Response(ctx context.Context, jsonString string) (*typedef.EngineSearchByAttributesV3response, error) {
	result := &typedef.EngineSearchByAttributesV3response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineSearchByAttributesResponse(ctx context.Context, jsonString string) (*typedef.EngineSearchByAttributesResponse, error) {
	result := &typedef.EngineSearchByAttributesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineStatsResponse(ctx context.Context, jsonString string) (*typedef.EngineStatsResponse, error) {
	result := &typedef.EngineStatsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalEngineStreamExportJSONEntityReportResponse(ctx context.Context, jsonString string) (*typedef.EngineStreamExportJSONEntityReportResponse, error) {
// 	result := &typedef.EngineStreamExportJSONEntityReportResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalEngineWhyEntitiesV2Response(ctx context.Context, jsonString string) (*typedef.EngineWhyEntitiesV2response, error) {
	result := &typedef.EngineWhyEntitiesV2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineWhyEntitiesResponse(ctx context.Context, jsonString string) (*typedef.EngineWhyEntitiesResponse, error) {
	result := &typedef.EngineWhyEntitiesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineWhyEntityByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.EngineWhyEntityByEntityIdv2response, error) {
	result := &typedef.EngineWhyEntityByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineWhyEntityByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.EngineWhyEntityByEntityIdresponse, error) {
	result := &typedef.EngineWhyEntityByEntityIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineWhyEntityByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.EngineWhyEntityByRecordIdv2response, error) {
	result := &typedef.EngineWhyEntityByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineWhyEntityByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.EngineWhyEntityByRecordIdresponse, error) {
	result := &typedef.EngineWhyEntityByRecordIdresponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineWhyRecordsV2Response(ctx context.Context, jsonString string) (*typedef.EngineWhyRecordsV2response, error) {
	result := &typedef.EngineWhyRecordsV2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineWhyRecordsResponse(ctx context.Context, jsonString string) (*typedef.EngineWhyRecordsResponse, error) {
	result := &typedef.EngineWhyRecordsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Product ----------------------------------------------------------------

func UnmarshalProductLicenseResponse(ctx context.Context, jsonString string) (*typedef.ProductLicenseResponse, error) {
	result := &typedef.ProductLicenseResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalProductVersionResponse(ctx context.Context, jsonString string) (*typedef.ProductVersionResponse, error) {
	result := &typedef.ProductVersionResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}
