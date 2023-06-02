package g2struct

// ----------------------------------------------------------------------------
// Types - sub-structs
//     Names do not start with "G2*"
// ----------------------------------------------------------------------------

type CompatibilityVersion struct {
	ConfigVersion string `json:"CONFIG_VERSION"`
}

type SchemaVersion struct {
	EngineSchemaVersion          string `json:"ENGINE_SCHEMA_VERSION"`
	MinimumRequiredSchemaVersion string `json:"MINIMUM_REQUIRED_SCHEMA_VERSION"`
	MaximumRequiredSchemaVersion string `json:"MAXIMUM_REQUIRED_SCHEMA_VERSION"`
}

// ----------------------------------------------------------------------------
// Types - structs
//     All names start with G2*
// ----------------------------------------------------------------------------

// The G2config interface is a Golang representation of Senzing's libg2config.h
type G2ProductVersionResponse struct {
	ProductName          string               `json:"PRODUCT_NAME"`
	Version              string               `json:"VERSION"`
	BuildVersion         string               `json:"BUILD_VERSION"`
	BuildDate            string               `json:"BUILD_DATE"`
	BuildNumber          string               `json:"BUILD_NUMBER"`
	CompatibilityVersion CompatibilityVersion `json:"COMPATIBILITY_VERSION"`
	SchemaVersion        SchemaVersion        `json:"SCHEMA_VERSION"`
}
