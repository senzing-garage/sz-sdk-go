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

type DataSource struct {
	DsrcCode string `json:"DSRC_CODE"`
	DsrcId   int64  `json:"DSRC_ID"`
}

type CfgAttr struct {
	Advanced     string `json:"ADVANCED"`
	AttrClass    string `json:"ATTR_CLASS"`
	AttrCode     string `json:"ATTR_CODE"`
	AttrId       int64  `json:"ATTR_ID"`
	DefaultValue string `json:"DEFAULT_VALUE"`
	FelemCode    string `json:"FELEM_CODE"`
	FelemReq     string `json:"FELEM_REQ"`
	FtypeCode    string `json:"FTYPE_CODE"`
	Internal     string `json:"INTERNAL"`
}

type CfgCfbom struct{}
type CfgCfcall struct{}
type CfgCfrtn struct{}
type CfgCfunc struct{}
type CfgDfbom struct{}
type CfgDfcall struct{}
type CfgDfunc struct{}
type CfgDsrc struct{}
type CfgDsrcInterest struct{}
type CfgEbom struct{}
type CfgEclass struct{}
type CfgEfbom struct{}
type CfgEfcall struct{}
type CfgEfunc struct{}
type CfgErfrag struct{}
type CfgErrule struct{}
type CfgEtype struct{}
type CfgFbom struct{}
type CfgFbovr struct{}
type CfgFclass struct{}
type CfgFelem struct{}
type CfgFtype struct{}
type CfgGenericThreshold struct{}
type CfgGplan struct{}
type CfgLens struct{}
type CfgLensrl struct{}
type CfgRclass struct{}
type CfgRtype struct{}
type CfgSfcall struct{}
type CfgSfunc struct{}
type SysOom struct{}
type ConfigBaseVersion struct{}

type G2Config struct {
	CfgAttr             []CfgAttr             `json:"CFG_ATTR"`
	CfgCfbom            []CfgCfbom            `json:"CFG_CFBOM"`
	CfgCfcall           []CfgCfcall           `json:"CFG_CFCALL"`
	CfgCfrtn            []CfgCfrtn            `json:"CFG_CFRTN"`
	CfgCfunc            []CfgCfunc            `json:"CFG_CFUNC"`
	CfgDfbom            []CfgDfbom            `json:"CFG_DFBOM"`
	CfgDfcall           []CfgDfcall           `json:"CFG_DFCALL"`
	CfgDfunc            []CfgDfunc            `json:"CFG_DFUNC"`
	CfgDsrc             []CfgDsrc             `json:"CFG_DSRC"`
	CfgDsrcInterest     []CfgDsrcInterest     `json:"CFG_DSRC_INTEREST"`
	CfgEbom             []CfgEbom             `json:"CFG_EBOM"`
	CfgEclass           []CfgEclass           `json:"CFG_ECLASS"`
	CfgEfbom            []CfgEfbom            `json:"CFG_EFBOM"`
	CfgEfcall           []CfgEfcall           `json:"CFG_EFCALL"`
	CfgEfunc            []CfgEfunc            `json:"CFG_EFUNC"`
	CfgErfrag           []CfgErfrag           `json:"CFG_ERFRAG"`
	CfgErrule           []CfgErrule           `json:"CFG_ERRULE"`
	CfgEtype            []CfgEtype            `json:"CFG_ETYPE"`
	CfgFbom             []CfgFbom             `json:"CFG_FBOM"`
	CfgFbovr            []CfgFbovr            `json:"CFG_FBOVR"`
	CfgFclass           []CfgFclass           `json:"CFG_FCLASS"`
	CfgFelem            []CfgFelem            `json:"CFG_FELEM"`
	CfgFtype            []CfgFtype            `json:"CFG_FTYPE"`
	CfgGenericThreshold []CfgGenericThreshold `json:"CFG_GENERIC_THRESHOLD"`
	CfgGplan            []CfgGplan            `json:"CFG_GPLAN"`
	CfgLens             []CfgLens             `json:"CFG_LENS"`
	CfgLensrl           []CfgLensrl           `json:"CFG_LENSRL"`
	CfgRclass           []CfgRclass           `json:"CFG_RCLASS"`
	CfgRtype            []CfgRtype            `json:"CFG_RTYPE"`
	CfgSfcall           []CfgSfcall           `json:"CFG_SFCALL"`
	CfgSfunc            []CfgSfunc            `json:"CFG_SFUNC"`
	SysOom              []SysOom              `json:"SYS_OOM"`
	ConfigBaseVersion   []ConfigBaseVersion   `json:"CONFIG_BASE_VERSION"`
}

type SchemaVersion struct {
	EngineSchemaVersion          string `json:"ENGINE_SCHEMA_VERSION"`
	MinimumRequiredSchemaVersion string `json:"MINIMUM_REQUIRED_SCHEMA_VERSION"`
	MaximumRequiredSchemaVersion string `json:"MAXIMUM_REQUIRED_SCHEMA_VERSION"`
}

// ----------------------------------------------------------------------------
// Types - structs - Config
// ----------------------------------------------------------------------------

// The ConfigAddDataSourceResponse...
type ConfigAddDataSourceResponse struct {
	DsrcId int64 `json:"DSRC_ID"`
}

type ConfigListDataSourcesResponse struct {
	DataSources []DataSource `json:"DATA_SOURCES"`
}

type ConfigSaveResponse struct {
	G2Config G2Config `json:"G2_CONFIG"`
}

// ----------------------------------------------------------------------------
// Types - structs - Product
// ----------------------------------------------------------------------------

// The ProductLicenseResponse...
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
