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

// --- Engine -----------------------------------------------------------------

// --- Product ----------------------------------------------------------------

/*
UnmarshalProductVersionResponse...
*/
func UnmarshalProductLicenseResponse(ctx context.Context, jsonString string) (*ProductLicenseResponse, error) {
	result := &ProductLicenseResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

/*
UnmarshalProductVersionResponse...
*/
func UnmarshalProductVersionResponse(ctx context.Context, jsonString string) (*ProductVersionResponse, error) {
	result := &ProductVersionResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}
