package senzing

import (
	"context"
	"encoding/json"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

/*
ParseProductVersionResponse...

Input
  - ctx: A context to control lifecycle.
*/
func ParseProductVersionResponse(ctx context.Context, jsonString string) (*ProductVersionResponse, error) {
	result := &ProductVersionResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}
