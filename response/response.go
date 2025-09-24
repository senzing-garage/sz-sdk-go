package response

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/senzing-garage/sz-sdk-json-type-definition/go/typedef"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// --- Config -----------------------------------------------------------------

func SzConfigRegisterDataSource(
	ctx context.Context,
	jsonString string,
) (*typedef.SzConfigRegisterDataSourceResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzConfigRegisterDataSourceResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzConfigRegisterDataSource cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzConfigExport(ctx context.Context, jsonString string) (*typedef.SzConfigExportConfigResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzConfigExportConfigResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzConfigExportConfig cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzConfigGetDataSourceRegistry(
	ctx context.Context,
	jsonString string,
) (*typedef.SzConfigGetDataSourceRegistryResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzConfigGetDataSourceRegistryResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzConfigGetDataSourceRegistry cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

// --- ConfigManager ----------------------------------------------------------

func SzConfigManagerGetConfigRegistry(
	ctx context.Context,
	jsonString string,
) (*typedef.SzConfigManagerGetConfigRegistryResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzConfigManagerGetConfigRegistryResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzConfigManagerGetConfigRegistry cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

// --- Diagnostic -------------------------------------------------------------

func SzDiagnosticCheckRepositoryPerformance(
	ctx context.Context,
	jsonString string,
) (*typedef.SzDiagnosticCheckRepositoryPerformanceResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzDiagnosticCheckRepositoryPerformanceResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzDiagnosticCheckRepositoryPerformance cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzDiagnosticGetRepositoryInfo(
	ctx context.Context,
	jsonString string,
) (*typedef.SzDiagnosticGetRepositoryInfoResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzDiagnosticGetRepositoryInfoResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzDiagnosticGetRepositoryInfo cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzDiagnosticGetFeature(ctx context.Context, jsonString string) (*typedef.SzDiagnosticGetFeatureResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzDiagnosticGetFeatureResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzDiagnosticGetFeature cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

// --- Engine -----------------------------------------------------------------

func SzEngineAddRecord(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineAddRecordResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineAddRecordResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineAddRecord cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineDeleteRecord(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineDeleteRecordResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineDeleteRecordResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineDeleteRecord cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineFetchNext(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineFetchNextResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineFetchNextResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineFetchNext cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineFindInterestingEntitiesByEntityID(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineFindInterestingEntitiesByEntityIDResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineFindInterestingEntitiesByEntityIDResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineFindInterestingEntitiesByEntityID cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineFindInterestingEntitiesByRecordID(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineFindInterestingEntitiesByRecordIDResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineFindInterestingEntitiesByRecordIDResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineFindInterestingEntitiesByRecordID cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineFindNetworkByEntityID(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineFindNetworkByEntityIDResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineFindNetworkByEntityIDResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineFindNetworkByEntityID cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineFindNetworkByRecordID(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineFindNetworkByRecordIDResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineFindNetworkByRecordIDResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineFindNetworkByRecordID cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineFindPathByEntityID(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineFindPathByEntityIDResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineFindPathByEntityIDResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineFindPathByEntityID cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineFindPathByRecordID(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineFindPathByRecordIDResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineFindPathByRecordIDResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineFindPathByRecordID cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineGetEntityByEntityID(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineGetEntityByEntityIDResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineGetEntityByEntityIDResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineGetEntityByEntityID cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineGetEntityByRecordID(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineGetEntityByRecordIDResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineGetEntityByRecordIDResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineGetEntityByRecordID cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineGetRecord(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineGetRecordResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineGetRecordResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineGetRecord cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineGetRedoRecord(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineGetRedoRecordResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineGetRedoRecordResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineGetRedoRecord cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineGetStats(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineGetStatsResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineGetStatsResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineGetStats cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineGetVirtualEntityByRecordID(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineGetVirtualEntityByRecordIDResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineGetVirtualEntityByRecordIDResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineGetVirtualEntityByRecordID cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineHowEntityByEntityID(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineHowEntityByEntityIDResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineHowEntityByEntityIDResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineHowEntityByEntityID cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineProcessRedoRecord(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineProcessRedoRecordResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineProcessRedoRecordResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineProcessRedoRecord cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineReevaluateEntity(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineReevaluateEntityResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineReevaluateEntityResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineReevaluateEntity cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineReevaluateRecord(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineReevaluateRecordResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineReevaluateRecordResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineReevaluateRecord cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineSearchByAttributes(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineSearchByAttributesResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineSearchByAttributesResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineSearchByAttributes cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineStreamExportJSONEntityReport(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineStreamExportJSONEntityReportResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineStreamExportJSONEntityReportResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineStreamExportJSONEntityReport cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineWhyEntities(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineWhyEntitiesResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineWhyEntitiesResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineWhyEntities cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineWhyRecordInEntity(
	ctx context.Context,
	jsonString string,
) (*typedef.SzEngineWhyRecordInEntityResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineWhyRecordInEntityResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineWhyRecordInEntity cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzEngineWhyRecords(ctx context.Context, jsonString string) (*typedef.SzEngineWhyRecordsResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzEngineWhyRecordsResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzEngineWhyRecords cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

// --- Product ----------------------------------------------------------------

func SzProductGetLicense(ctx context.Context, jsonString string) (*typedef.SzProductGetLicenseResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzProductGetLicenseResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzProductGetLicense cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}

func SzProductGetVersion(ctx context.Context, jsonString string) (*typedef.SzProductGetVersionResponse, error) {
	_ = ctx

	var err error

	result := &typedef.SzProductGetVersionResponse{} //exhaustruct:ignore

	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		err = fmt.Errorf("SzProductGetVersion cannot unmarshal %s: %w", jsonString, err)
	}

	return result, err
}
