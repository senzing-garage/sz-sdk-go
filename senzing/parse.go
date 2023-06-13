package senzing

import (
	"context"
	"encoding/json"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

/*
ParseConfigAddDataSourceResponse...
*/
func ParseConfigAddDataSourceResponse(ctx context.Context, jsonString string) (*ConfigAddDataSourceResponse, error) {
	result := &ConfigAddDataSourceResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

/*
ParseConfigListDataSourcesResponse...
*/
func ParseConfigListDataSourcesResponse(ctx context.Context, jsonString string) (*ConfigListDataSourcesResponse, error) {
	result := &ConfigListDataSourcesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

/*
ParseConfigSaveResponse...
*/
func ParseConfigSaveResponse(ctx context.Context, jsonString string) (*ConfigSaveResponse, error) {
	result := &ConfigSaveResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

/*
ParseProductVersionResponse...
*/
func ParseProductLicenseResponse(ctx context.Context, jsonString string) (*ProductLicenseResponse, error) {
	result := &ProductLicenseResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

/*
ParseProductVersionResponse...
*/
func ParseProductVersionResponse(ctx context.Context, jsonString string) (*ProductVersionResponse, error) {
	result := &ProductVersionResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}
