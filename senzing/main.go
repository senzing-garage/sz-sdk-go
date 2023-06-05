package senzing

// ----------------------------------------------------------------------------
// Types - interface
// ----------------------------------------------------------------------------

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
// Types - structs - Product
//     All names start with G2*
// ----------------------------------------------------------------------------

// The ProductVersionResponse...
type ProductLicenseResponse struct {
	Customer     string `json:"customer"`
	Contract     string `json:"contract"`
	IssueDate    string `json:"issueDate"`
	LicenseType  string `json:"licenseType"`
	LicenseLevel string `json:"licenseLevel"`
	Billing      string `json:"billing"`
	ExpireDate   string `json:"expireDate"`
	RecordLimit  int64  `json:"recordLimit"`
}

// The ProductVersionResponse...
type ProductVersionResponse struct {
	ProductName          string               `json:"PRODUCT_NAME"`
	Version              string               `json:"VERSION"`
	BuildVersion         string               `json:"BUILD_VERSION"`
	BuildDate            string               `json:"BUILD_DATE"`
	BuildNumber          string               `json:"BUILD_NUMBER"`
	CompatibilityVersion CompatibilityVersion `json:"COMPATIBILITY_VERSION"`
	SchemaVersion        SchemaVersion        `json:"SCHEMA_VERSION"`
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Log message prefix.
const Prefix = "senzing."

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for g2product implementations.
var IdMessages = map[int]string{
	1:    "Enter " + Prefix + "ClearLastException().",
	2:    "Exit  " + Prefix + "ClearLastException() returned (%v).",
	5901: "During test setup, call to messagelogger.NewSenzingApiLogger() failed.",
	8001: Prefix + "Destroy",
}

// Status strings for specific g2product messages.
var IdStatuses = map[int]string{}
