package senzing

import (
	"context"
	"encoding/json"

	"github.com/senzing-garage/g2-sdk-json-type-definition/go/typedef"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// --- Config -----------------------------------------------------------------

func UnmarshalG2configAddDataSourceResponse(ctx context.Context, jsonString string) (*typedef.G2configAddDataSourceResponse, error) {
	result := &typedef.G2configAddDataSourceResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2configListDataSourcesResponse(ctx context.Context, jsonString string) (*typedef.G2configListDataSourcesResponse, error) {
	result := &typedef.G2configListDataSourcesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2configSaveResponse(ctx context.Context, jsonString string) (*typedef.G2configSaveResponse, error) {
	result := &typedef.G2configSaveResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Configmgr --------------------------------------------------------------

func UnmarshalG2configmgrGetConfigResponse(ctx context.Context, jsonString string) (*typedef.G2configmgrGetConfigResponse, error) {
	result := &typedef.G2configmgrGetConfigResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2configmgrGetConfigListResponse(ctx context.Context, jsonString string) (*typedef.G2configmgrGetConfigListResponse, error) {
	result := &typedef.G2configmgrGetConfigListResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Diagnostic -------------------------------------------------------------

func UnmarshalG2diagnosticCheckDBPerfResponse(ctx context.Context, jsonString string) (*typedef.G2diagnosticCheckDbperfResponse, error) {
	result := &typedef.G2diagnosticCheckDbperfResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2diagnosticFetchNextEntityBySizeResponse(ctx context.Context, jsonString string) (*typedef.G2diagnosticFetchNextEntityBySizeResponse, error) {
	result := &typedef.G2diagnosticFetchNextEntityBySizeResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2diagnosticGetDataSourceCountsResponse(ctx context.Context, jsonString string) (*typedef.G2diagnosticGetDataSourceCountsResponse, error) {
	result := &typedef.G2diagnosticGetDataSourceCountsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2diagnosticGetDBInfoResponse(ctx context.Context, jsonString string) (*typedef.G2diagnosticGetDbinfoResponse, error) {
	result := &typedef.G2diagnosticGetDbinfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2diagnosticGetEntityDetailsResponse(ctx context.Context, jsonString string) (*typedef.G2diagnosticGetEntityDetailsResponse, error) {
	result := &typedef.G2diagnosticGetEntityDetailsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalG2diagnosticGetEntityListBySizeResponse(ctx context.Context, jsonString string) (*typedef.G2diagnosticGetEntityListBySizeResponse, error) {
// 	result := &typedef.G2diagnosticGetEntityListBySizeResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalG2diagnosticGetEntityResumeResponse(ctx context.Context, jsonString string) (*typedef.G2diagnosticGetEntityResumeResponse, error) {
	result := &typedef.G2diagnosticGetEntityResumeResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2diagnosticGetEntitySizeBreakdownResponse(ctx context.Context, jsonString string) (*typedef.G2diagnosticGetEntitySizeBreakdownResponse, error) {
	result := &typedef.G2diagnosticGetEntitySizeBreakdownResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2diagnosticGetFeatureResponse(ctx context.Context, jsonString string) (*typedef.G2diagnosticGetFeatureResponse, error) {
	result := &typedef.G2diagnosticGetFeatureResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2diagnosticGetGenericFeaturesResponse(ctx context.Context, jsonString string) (*typedef.G2diagnosticGetGenericFeaturesResponse, error) {
	result := &typedef.G2diagnosticGetGenericFeaturesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2diagnosticGetMappingStatisticsResponse(ctx context.Context, jsonString string) (*typedef.G2diagnosticGetMappingStatisticsResponse, error) {
	result := &typedef.G2diagnosticGetMappingStatisticsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2diagnosticGetRelationshipDetailsResponse(ctx context.Context, jsonString string) (*typedef.G2diagnosticGetRelationshipDetailsResponse, error) {
	result := &typedef.G2diagnosticGetRelationshipDetailsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2diagnosticGetResolutionStatisticsResponse(ctx context.Context, jsonString string) (*typedef.G2diagnosticGetResolutionStatisticsResponse, error) {
	result := &typedef.G2diagnosticGetResolutionStatisticsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalG2diagnosticStreamEntityListBySizeResponse(ctx context.Context, jsonString string) (*typedef.G2diagnosticStreamEntityListBySizeResponse, error) {
// 	result := &typedef.G2diagnosticStreamEntityListBySizeResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

// --- G2engine -----------------------------------------------------------------

func UnmarshalG2engineAddRecordWithInfoResponse(ctx context.Context, jsonString string) (*typedef.G2engineAddRecordWithInfoResponse, error) {
	result := &typedef.G2engineAddRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineAddRecordWithInfoWithReturnedRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineAddRecordWithInfoWithReturnedRecordIDResponse, error) {
	result := &typedef.G2engineAddRecordWithInfoWithReturnedRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalG2engineAddRecordWithReturnedRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineAddRecordWithReturnedRecordIDResponse, error) {
// 	result := &typedef.G2engineAddRecordWithReturnedRecordIDResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalG2engineCheckRecordResponse(ctx context.Context, jsonString string) (*typedef.G2engineCheckRecordResponse, error) {
	result := &typedef.G2engineCheckRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineDeleteRecordWithInfoResponse(ctx context.Context, jsonString string) (*typedef.G2engineDeleteRecordWithInfoResponse, error) {
	result := &typedef.G2engineDeleteRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineExportConfigAndConfigIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineExportConfigAndConfigIDResponse, error) {
	result := &typedef.G2engineExportConfigAndConfigIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineExportConfigResponse(ctx context.Context, jsonString string) (*typedef.G2engineExportConfigResponse, error) {
	result := &typedef.G2engineExportConfigResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalG2engineFetchNextResponse(ctx context.Context, jsonString string) (*typedef.G2engineFetchNextResponse, error) {
// 	result := &typedef.G2engineFetchNextResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalG2engineFindInterestingEntitiesByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindInterestingEntitiesByEntityIDResponse, error) {
	result := &typedef.G2engineFindInterestingEntitiesByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindInterestingEntitiesByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindInterestingEntitiesByRecordIDResponse, error) {
	result := &typedef.G2engineFindInterestingEntitiesByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindNetworkByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindNetworkByEntityIdv2response, error) {
	result := &typedef.G2engineFindNetworkByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindNetworkByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindNetworkByEntityIDResponse, error) {
	result := &typedef.G2engineFindNetworkByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindNetworkByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindNetworkByRecordIdv2response, error) {
	result := &typedef.G2engineFindNetworkByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindNetworkByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindNetworkByRecordIDResponse, error) {
	result := &typedef.G2engineFindNetworkByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindPathByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindPathByEntityIdv2response, error) {
	result := &typedef.G2engineFindPathByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindPathByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindPathByEntityIDResponse, error) {
	result := &typedef.G2engineFindPathByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindPathByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindPathByRecordIdv2response, error) {
	result := &typedef.G2engineFindPathByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindPathByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindPathByRecordIDResponse, error) {
	result := &typedef.G2engineFindPathByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindPathExcludingByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindPathExcludingByEntityIdv2response, error) {
	result := &typedef.G2engineFindPathExcludingByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindPathExcludingByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindPathExcludingByEntityIDResponse, error) {
	result := &typedef.G2engineFindPathExcludingByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindPathExcludingByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindPathExcludingByRecordIdv2response, error) {
	result := &typedef.G2engineFindPathExcludingByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindPathExcludingByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindPathExcludingByRecordIDResponse, error) {
	result := &typedef.G2engineFindPathExcludingByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindPathIncludingSourceByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindPathIncludingSourceByEntityIdv2response, error) {
	result := &typedef.G2engineFindPathIncludingSourceByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindPathIncludingSourceByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindPathIncludingSourceByEntityIDResponse, error) {
	result := &typedef.G2engineFindPathIncludingSourceByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindPathIncludingSourceByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindPathIncludingSourceByRecordIdv2response, error) {
	result := &typedef.G2engineFindPathIncludingSourceByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineFindPathIncludingSourceByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindPathIncludingSourceByRecordIDResponse, error) {
	result := &typedef.G2engineFindPathIncludingSourceByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineGetEntityByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineGetEntityByEntityIdv2response, error) {
	result := &typedef.G2engineGetEntityByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineGetEntityByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineGetEntityByEntityIDResponse, error) {
	result := &typedef.G2engineGetEntityByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineGetEntityByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineGetEntityByRecordIdv2response, error) {
	result := &typedef.G2engineGetEntityByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineGetEntityByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineGetEntityByRecordIDResponse, error) {
	result := &typedef.G2engineGetEntityByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineGetRecordV2Response(ctx context.Context, jsonString string) (*typedef.G2engineGetRecordV2response, error) {
	result := &typedef.G2engineGetRecordV2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineGetRecordResponse(ctx context.Context, jsonString string) (*typedef.G2engineGetRecordResponse, error) {
	result := &typedef.G2engineGetRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalG2engineGetRedoRecordResponse(ctx context.Context, jsonString string) (*typedef.G2engineGetRedoRecordResponse, error) {
// 	result := &typedef.G2engineGetRedoRecordResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalG2engineGetVirtualEntityByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineGetVirtualEntityByRecordIdv2response, error) {
	result := &typedef.G2engineGetVirtualEntityByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineGetVirtualEntityByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineGetVirtualEntityByRecordIDResponse, error) {
	result := &typedef.G2engineGetVirtualEntityByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineHowEntityByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineHowEntityByEntityIdv2response, error) {
	result := &typedef.G2engineHowEntityByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineHowEntityByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineHowEntityByEntityIDResponse, error) {
	result := &typedef.G2engineHowEntityByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalG2engineProcessRedoRecordResponse(ctx context.Context, jsonString string) (*typedef.G2engineProcessRedoRecordResponse, error) {
// 	result := &typedef.G2engineProcessRedoRecordResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalG2engineProcessRedoRecordWithInfoResponse(ctx context.Context, jsonString string) (*typedef.G2engineProcessRedoRecordWithInfoResponse, error) {
	result := &typedef.G2engineProcessRedoRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineProcessWithInfoResponse(ctx context.Context, jsonString string) (*typedef.G2engineProcessWithInfoResponse, error) {
	result := &typedef.G2engineProcessWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineProcessWithResponseResizeResponse(ctx context.Context, jsonString string) (*typedef.G2engineProcessWithResponseResizeResponse, error) {
	result := &typedef.G2engineProcessWithResponseResizeResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineProcessWithResponseResponse(ctx context.Context, jsonString string) (*typedef.G2engineProcessWithResponseResponse, error) {
	result := &typedef.G2engineProcessWithResponseResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineReevaluateEntityWithInfoResponse(ctx context.Context, jsonString string) (*typedef.G2engineReevaluateEntityWithInfoResponse, error) {
	result := &typedef.G2engineReevaluateEntityWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineReevaluateRecordWithInfoResponse(ctx context.Context, jsonString string) (*typedef.G2engineReevaluateRecordWithInfoResponse, error) {
	result := &typedef.G2engineReevaluateRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineReplaceRecordWithInfoResponse(ctx context.Context, jsonString string) (*typedef.G2engineReplaceRecordWithInfoResponse, error) {
	result := &typedef.G2engineReplaceRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineSearchByAttributesV2Response(ctx context.Context, jsonString string) (*typedef.G2engineSearchByAttributesV2response, error) {
	result := &typedef.G2engineSearchByAttributesV2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineSearchByAttributesV3Response(ctx context.Context, jsonString string) (*typedef.G2engineSearchByAttributesV3response, error) {
	result := &typedef.G2engineSearchByAttributesV3response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineSearchByAttributesResponse(ctx context.Context, jsonString string) (*typedef.G2engineSearchByAttributesResponse, error) {
	result := &typedef.G2engineSearchByAttributesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineStatsResponse(ctx context.Context, jsonString string) (*typedef.G2engineStatsResponse, error) {
	result := &typedef.G2engineStatsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalG2engineStreamExportJSONEntityReportResponse(ctx context.Context, jsonString string) (*typedef.G2engineStreamExportJSONEntityReportResponse, error) {
// 	result := &typedef.G2engineStreamExportJSONEntityReportResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalG2engineWhyEntitiesV2Response(ctx context.Context, jsonString string) (*typedef.G2engineWhyEntitiesV2response, error) {
	result := &typedef.G2engineWhyEntitiesV2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineWhyEntitiesResponse(ctx context.Context, jsonString string) (*typedef.G2engineWhyEntitiesResponse, error) {
	result := &typedef.G2engineWhyEntitiesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineWhyEntityByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineWhyEntityByEntityIdv2response, error) {
	result := &typedef.G2engineWhyEntityByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineWhyEntityByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineWhyEntityByEntityIDResponse, error) {
	result := &typedef.G2engineWhyEntityByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineWhyEntityByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineWhyEntityByRecordIdv2response, error) {
	result := &typedef.G2engineWhyEntityByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineWhyEntityByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineWhyEntityByRecordIDResponse, error) {
	result := &typedef.G2engineWhyEntityByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineWhyRecordsV2Response(ctx context.Context, jsonString string) (*typedef.G2engineWhyRecordsV2response, error) {
	result := &typedef.G2engineWhyRecordsV2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2engineWhyRecordsResponse(ctx context.Context, jsonString string) (*typedef.G2engineWhyRecordsResponse, error) {
	result := &typedef.G2engineWhyRecordsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Product ----------------------------------------------------------------

func UnmarshalG2productLicenseResponse(ctx context.Context, jsonString string) (*typedef.G2productLicenseResponse, error) {
	result := &typedef.G2productLicenseResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalG2productVersionResponse(ctx context.Context, jsonString string) (*typedef.G2productVersionResponse, error) {
	result := &typedef.G2productVersionResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}
