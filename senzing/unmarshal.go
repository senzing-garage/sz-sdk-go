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

func UnmarshalSzconfigAddDataSourceResponse(ctx context.Context, jsonString string) (*typedef.G2configAddDataSourceResponse, error) {
	result := &typedef.G2configAddDataSourceResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzconfigListDataSourcesResponse(ctx context.Context, jsonString string) (*typedef.G2configListDataSourcesResponse, error) {
	result := &typedef.G2configListDataSourcesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzconfigSaveResponse(ctx context.Context, jsonString string) (*typedef.G2configSaveResponse, error) {
	result := &typedef.G2configSaveResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Configmgr --------------------------------------------------------------

func UnmarshalSzconfigmgrGetConfigResponse(ctx context.Context, jsonString string) (*typedef.G2configmgrGetConfigResponse, error) {
	result := &typedef.G2configmgrGetConfigResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzconfigmgrGetConfigListResponse(ctx context.Context, jsonString string) (*typedef.G2configmgrGetConfigListResponse, error) {
	result := &typedef.G2configmgrGetConfigListResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Diagnostic -------------------------------------------------------------

func UnmarshalSzdiagnosticCheckDBPerfResponse(ctx context.Context, jsonString string) (*typedef.G2diagnosticCheckDbperfResponse, error) {
	result := &typedef.G2diagnosticCheckDbperfResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- G2engine -----------------------------------------------------------------

func UnmarshalSzengineAddRecordWithInfoResponse(ctx context.Context, jsonString string) (*typedef.G2engineAddRecordWithInfoResponse, error) {
	result := &typedef.G2engineAddRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineDeleteRecordWithInfoResponse(ctx context.Context, jsonString string) (*typedef.G2engineDeleteRecordWithInfoResponse, error) {
	result := &typedef.G2engineDeleteRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineExportConfigAndConfigIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineExportConfigAndConfigIDResponse, error) {
	result := &typedef.G2engineExportConfigAndConfigIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineExportConfigResponse(ctx context.Context, jsonString string) (*typedef.G2engineExportConfigResponse, error) {
	result := &typedef.G2engineExportConfigResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalSzengineFetchNextResponse(ctx context.Context, jsonString string) (*typedef.G2engineFetchNextResponse, error) {
// 	result := &typedef.G2engineFetchNextResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalSzengineFindInterestingEntitiesByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindInterestingEntitiesByEntityIDResponse, error) {
	result := &typedef.G2engineFindInterestingEntitiesByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindInterestingEntitiesByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindInterestingEntitiesByRecordIDResponse, error) {
	result := &typedef.G2engineFindInterestingEntitiesByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindNetworkByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindNetworkByEntityIdv2response, error) {
	result := &typedef.G2engineFindNetworkByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindNetworkByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindNetworkByEntityIDResponse, error) {
	result := &typedef.G2engineFindNetworkByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindNetworkByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindNetworkByRecordIdv2response, error) {
	result := &typedef.G2engineFindNetworkByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindNetworkByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindNetworkByRecordIDResponse, error) {
	result := &typedef.G2engineFindNetworkByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindPathByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindPathByEntityIdv2response, error) {
	result := &typedef.G2engineFindPathByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindPathByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindPathByEntityIDResponse, error) {
	result := &typedef.G2engineFindPathByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindPathByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindPathByRecordIdv2response, error) {
	result := &typedef.G2engineFindPathByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindPathByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindPathByRecordIDResponse, error) {
	result := &typedef.G2engineFindPathByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindPathExcludingByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindPathExcludingByEntityIdv2response, error) {
	result := &typedef.G2engineFindPathExcludingByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindPathExcludingByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindPathExcludingByEntityIDResponse, error) {
	result := &typedef.G2engineFindPathExcludingByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindPathExcludingByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindPathExcludingByRecordIdv2response, error) {
	result := &typedef.G2engineFindPathExcludingByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindPathExcludingByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindPathExcludingByRecordIDResponse, error) {
	result := &typedef.G2engineFindPathExcludingByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindPathIncludingSourceByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindPathIncludingSourceByEntityIdv2response, error) {
	result := &typedef.G2engineFindPathIncludingSourceByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindPathIncludingSourceByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindPathIncludingSourceByEntityIDResponse, error) {
	result := &typedef.G2engineFindPathIncludingSourceByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindPathIncludingSourceByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineFindPathIncludingSourceByRecordIdv2response, error) {
	result := &typedef.G2engineFindPathIncludingSourceByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineFindPathIncludingSourceByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineFindPathIncludingSourceByRecordIDResponse, error) {
	result := &typedef.G2engineFindPathIncludingSourceByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineGetEntityByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineGetEntityByEntityIdv2response, error) {
	result := &typedef.G2engineGetEntityByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineGetEntityByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineGetEntityByEntityIDResponse, error) {
	result := &typedef.G2engineGetEntityByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineGetEntityByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineGetEntityByRecordIdv2response, error) {
	result := &typedef.G2engineGetEntityByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineGetEntityByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineGetEntityByRecordIDResponse, error) {
	result := &typedef.G2engineGetEntityByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineGetRecordV2Response(ctx context.Context, jsonString string) (*typedef.G2engineGetRecordV2response, error) {
	result := &typedef.G2engineGetRecordV2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineGetRecordResponse(ctx context.Context, jsonString string) (*typedef.G2engineGetRecordResponse, error) {
	result := &typedef.G2engineGetRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalSzengineGetRedoRecordResponse(ctx context.Context, jsonString string) (*typedef.G2engineGetRedoRecordResponse, error) {
// 	result := &typedef.G2engineGetRedoRecordResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalSzengineGetVirtualEntityByRecordIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineGetVirtualEntityByRecordIdv2response, error) {
	result := &typedef.G2engineGetVirtualEntityByRecordIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineGetVirtualEntityByRecordIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineGetVirtualEntityByRecordIDResponse, error) {
	result := &typedef.G2engineGetVirtualEntityByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineHowEntityByEntityIDV2Response(ctx context.Context, jsonString string) (*typedef.G2engineHowEntityByEntityIdv2response, error) {
	result := &typedef.G2engineHowEntityByEntityIdv2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineHowEntityByEntityIDResponse(ctx context.Context, jsonString string) (*typedef.G2engineHowEntityByEntityIDResponse, error) {
	result := &typedef.G2engineHowEntityByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalSzengineProcessRedoRecordResponse(ctx context.Context, jsonString string) (*typedef.G2engineProcessRedoRecordResponse, error) {
// 	result := &typedef.G2engineProcessRedoRecordResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalSzengineProcessRedoRecordWithInfoResponse(ctx context.Context, jsonString string) (*typedef.G2engineProcessRedoRecordWithInfoResponse, error) {
	result := &typedef.G2engineProcessRedoRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineReevaluateEntityWithInfoResponse(ctx context.Context, jsonString string) (*typedef.G2engineReevaluateEntityWithInfoResponse, error) {
	result := &typedef.G2engineReevaluateEntityWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineReevaluateRecordWithInfoResponse(ctx context.Context, jsonString string) (*typedef.G2engineReevaluateRecordWithInfoResponse, error) {
	result := &typedef.G2engineReevaluateRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineReplaceRecordWithInfoResponse(ctx context.Context, jsonString string) (*typedef.G2engineReplaceRecordWithInfoResponse, error) {
	result := &typedef.G2engineReplaceRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineSearchByAttributesV2Response(ctx context.Context, jsonString string) (*typedef.G2engineSearchByAttributesV2response, error) {
	result := &typedef.G2engineSearchByAttributesV2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineSearchByAttributesV3Response(ctx context.Context, jsonString string) (*typedef.G2engineSearchByAttributesV3response, error) {
	result := &typedef.G2engineSearchByAttributesV3response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineSearchByAttributesResponse(ctx context.Context, jsonString string) (*typedef.G2engineSearchByAttributesResponse, error) {
	result := &typedef.G2engineSearchByAttributesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineStatsResponse(ctx context.Context, jsonString string) (*typedef.G2engineStatsResponse, error) {
	result := &typedef.G2engineStatsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// func UnmarshalSzengineStreamExportJSONEntityReportResponse(ctx context.Context, jsonString string) (*typedef.G2engineStreamExportJSONEntityReportResponse, error) {
// 	result := &typedef.G2engineStreamExportJSONEntityReportResponse{}
// 	err := json.Unmarshal([]byte(jsonString), result)
// 	return result, err
// }

func UnmarshalSzengineWhyEntitiesV2Response(ctx context.Context, jsonString string) (*typedef.G2engineWhyEntitiesV2response, error) {
	result := &typedef.G2engineWhyEntitiesV2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineWhyEntitiesResponse(ctx context.Context, jsonString string) (*typedef.G2engineWhyEntitiesResponse, error) {
	result := &typedef.G2engineWhyEntitiesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineWhyRecordsV2Response(ctx context.Context, jsonString string) (*typedef.G2engineWhyRecordsV2response, error) {
	result := &typedef.G2engineWhyRecordsV2response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzengineWhyRecordsResponse(ctx context.Context, jsonString string) (*typedef.G2engineWhyRecordsResponse, error) {
	result := &typedef.G2engineWhyRecordsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Product ----------------------------------------------------------------

func UnmarshalSzproductLicenseResponse(ctx context.Context, jsonString string) (*typedef.G2productLicenseResponse, error) {
	result := &typedef.G2productLicenseResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalSzproductVersionResponse(ctx context.Context, jsonString string) (*typedef.G2productVersionResponse, error) {
	result := &typedef.G2productVersionResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}
