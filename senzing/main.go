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

type CfgCfbom struct {
	CfcallId  int64 `json:"CFCALL_ID"`
	ExecOrder int64 `json:"EXEC_ORDER"`
	FelemId   int64 `json:"FELEM_ID"`
	FtypeId   int64 `json:"FTYPE_ID"`
}

type CfgCfcall struct {
	CfcallId  int64 `json:"CFCALL_ID"`
	CfuncId   int64 `json:"CFUNC_ID"`
	ExecOrder int64 `json:"EXEC_ORDER"`
	FtypeId   int64 `json:"FTYPE_ID"`
}

type CfgCfrtn struct {
	CfrtnId        int64  `json:"CFRTN_ID"`
	CfuncId        int64  `json:"CFUNC_ID"`
	CfuncRtnval    string `json:"CFUNC_RTNVAL"`
	CloseScore     int64  `json:"CLOSE_SCORE"`
	ExecOrder      int64  `json:"EXEC_ORDER"`
	FtypeId        int64  `json:"FTYPE_ID"`
	LikelyScore    int64  `json:"LIKELY_SCORE"`
	PlausibleScore int64  `json:"PLAUSIBLE_SCORE"`
	SameScore      int64  `json:"SAME_SCORE"`
	UnlikelyScore  int64  `json:"UN_LIKELY_SCORE"`
}

type CfgCfunc struct {
	AnonSupport   string `json:"ANON_SUPPORT"`
	CfuncCode     string `json:"CFUNC_CODE"`
	CfuncDesc     string `json:"CFUNC_DESC"`
	CfuncId       int64  `json:"CFUNC_ID"`
	ConnectStr    string `json:"CONNECT_STR"`
	FuncLib       string `json:"FUNC_LIB"`
	FuncVer       string `json:"FUNC_VER"`
	JavaClassName string `json:"JAVA_CLASS_NAME"`
	Language      string `json:"LANGUAGE"`
}

type CfgDfbom struct {
	DfcallId  int64 `json:"DFCALL_ID"`
	ExecOrder int64 `json:"EXEC_ORDER"`
	FelemId   int64 `json:"FELEM_ID"`
	FtypeId   int64 `json:"FTYPE_ID"`
}

type CfgDfcall struct {
	DfcallId  int64 `json:"DFCALL_ID"`
	DfuncId   int64 `json:"DFUNC_ID"`
	ExecOrder int64 `json:"EXEC_ORDER"`
	FtypeId   int64 `json:"FTYPE_ID"`
}

type CfgDfunc struct {
	AnonSupport   string `json:"ANON_SUPPORT"`
	ConnectStr    string `json:"CONNECT_STR"`
	DfuncCode     string `json:"DFUNC_CODE"`
	DfuncDesc     string `json:"DFUNC_DESC"`
	DfuncId       int64  `json:"DFUNC_ID"`
	FuncLib       string `json:"FUNC_LIB"`
	FuncVer       string `json:"FUNC_VER"`
	JavaClassName string `json:"JAVA_CLASS_NAME"`
	Language      string `json:"LANGUAGE"`
}

type CfgDsrc struct {
	Conversational string `json:"CONVERSATIONAL"`
	DsrcCode       string `json:"DSRC_CODE"`
	DsrcDesc       string `json:"DSRC_DESC"`
	DsrcId         int64  `json:"DSRC_ID"`
	DsrcRely       int64  `json:"DSRC_RELY"`
	RetentionLevel string `json:"RETENTION_LEVEL"`
}

type CfgDsrcInterest struct{}

type CfgEbom struct {
	EtypeId   int64  `json:"ETYPE_ID"`
	ExecOrder int64  `json:"EXEC_ORDER"`
	FtypeId   int64  `json:"FTYPE_ID"`
	UtypeCode string `json:"UTYPE_CODE"`
}

type CfgEclass struct {
	EclassCode string `json:"ECLASS_CODE"`
	EclassDesc string `json:"ECLASS_DESC"`
	EclassId   int64  `json:"ECLASS_ID"`
	Resolve    string `json:"RESOLVE"`
}

type CfgEfbom struct {
	EfcallId  int64  `json:"EFCALL_ID"`
	ExecOrder int64  `json:"EXEC_ORDER"`
	FelemId   int64  `json:"FELEM_ID"`
	FelemReq  string `json:"FELEM_REQ"`
	FtypeId   int64  `json:"FTYPE_ID"`
}

type CfgEfcall struct {
	EfcallId     int64  `json:"EFCALL_ID"`
	EfeatFtypeId int64  `json:"EFEAT_FTYPE_ID"`
	EfuncId      int64  `json:"EFUNC_ID"`
	ExecOrder    int64  `json:"EXEC_ORDER"`
	FelemId      int64  `json:"FELEM_ID"`
	FtypeId      int64  `json:"FTYPE_ID"`
	IsVirtual    string `json:"IS_VIRTUAL"`
}

type CfgEfunc struct {
	ConnectStr    string `json:"CONNECT_STR"`
	EfuncCode     string `json:"EFUNC_CODE"`
	EfuncDesc     string `json:"EFUNC_DESC"`
	EfuncId       int64  `json:"EFUNC_ID"`
	FuncLib       string `json:"FUNC_LIB"`
	FuncVer       string `json:"FUNC_VER"`
	JavaClassName string `json:"JAVA_CLASS_NAME"`
	Language      string `json:"LANGUAGE"`
}

type CfgErfrag struct {
	ErfragCode    string `json:"ERFRAG_CODE"`
	ErfragDepends string `json:"ERFRAG_DEPENDS"`
	ErfragDesc    string `json:"ERFRAG_DESC"`
	ErfragId      int64  `json:"ERFRAG_ID"`
	ErfragSource  string `json:"ERFRAG_SOURCE"`
}

type CfgErrule struct {
	DisqErfragCode string `json:"DISQ_ERFRAG_CODE"`
	ErruleCode     string `json:"ERRULE_CODE"`
	ErruleDesc     string `json:"ERRULE_DESC"`
	ErruleId       int64  `json:"ERRULE_ID"`
	ErruleTier     int64  `json:"ERRULE_TIER"`
	QualErfragCode string `json:"QUAL_ERFRAG_CODE"`
	RefScore       int64  `json:"REF_SCORE"`
	Relate         string `json:"RELATE"`
	Resolve        string `json:"RESOLVE"`
	RtypeId        int64  `json:"RTYPE_ID"`
}

type CfgEtype struct {
	EclassId  int64  `json:"ECLASS_ID"`
	EtypeCode string `json:"ETYPE_CODE"`
	EtypeDesc string `json:"ETYPE_DESC"`
	EtypeId   int64  `json:"ETYPE_ID"`
}

type CfgFbom struct {
	Derived      string `json:"DERIVED"`
	DisplayDelim string `json:"DISPLAY_DELIM"`
	DisplayLevel int64  `json:"DISPLAY_LEVEL"`
	ExecOrder    int64  `json:"EXEC_ORDER"`
	FelemId      int64  `json:"FELEM_ID"`
	FtypeId      int64  `json:"FTYPE_ID"`
}

type CfgFbovr struct {
	EclassId  int64  `json:"ECLASS_ID"`
	FtypeExcl string `json:"FTYPE_EXCL"`
	FtypeFreq string `json:"FTYPE_FREQ"`
	FtypeId   int64  `json:"FTYPE_ID"`
	FtypeStab string `json:"FTYPE_STAB"`
	UtypeCode string `json:"UTYPE_CODE"`
}

type CfgFclass struct {
	FclassCode string `json:"FCLASS_CODE"`
	FclassDesc string `json:"FCLASS_DESC"`
	FclassId   int64  `json:"FCLASS_ID"`
}
type CfgFelem struct {
	DataType  string `json:"DATA_TYPE"`
	FelemCode string `json:"FELEM_CODE"`
	FelemDesc string `json:"FELEM_DESC"`
	FelemId   int64  `json:"FELEM_ID"`
	Tokenize  string `json:"TOKENIZE"`
}

type CfgFtype struct {
	Anonymize      string `json:"ANONYMIZE"`
	Derivation     string `json:"DERIVATION"`
	Derived        string `json:"DERIVED"`
	FclassId       int64  `json:"FCLASS_ID"`
	FtypeCode      string `json:"FTYPE_CODE"`
	FtypeDesc      string `json:"FTYPE_DESC"`
	FtypeExcl      string `json:"FTYPE_EXCL"`
	FtypeFreq      string `json:"FTYPE_FREQ"`
	FtypeId        int64  `json:"FTYPE_ID"`
	FtypeStab      string `json:"FTYPE_STAB"`
	PersistHistory string `json:"PERSIST_HISTORY"`
	RtypeID        int64  `json:"RTYPE_ID"`
	ShowInMatchKey string `json:"SHOW_IN_MATCH_KEY"`
	UsedForCand    string `json:"USED_FOR_CAND"`
	Version        int64  `json:"VERSION"`
}

type CfgGenericThreshold struct {
	Behavior     string `json:"BEHAVIOR"`
	CandidateCap int64  `json:"CANDIDATE_CAP"`
	FtypeId      int64  `json:"FTYPE_ID"`
	GplanId      int64  `json:"GPLAN_ID"`
	ScoringCap   int64  `json:"SCORING_CAP"`
	SendToRedo   string `json:"SEND_TO_REDO"`
}

type CfgGplan struct {
	GplanCode string `json:"GPLAN_CODE"`
	GplanDesc string `json:"GPLAN_DESC"`
	GplanId   int64  `json:"GPLAN_ID"`
}

type CfgLens struct {
	LensCode string `json:"LENS_CODE"`
	LensDesc string `json:"LENS_DESC"`
	LensId   int64  `json:"LENS_ID"`
}

type CfgLensrl struct{}

type CfgRclass struct {
	IsDisclosed string `json:"IS_DISCLOSED"`
	RclassCode  string `json:"RCLASS_CODE"`
	RclassDesc  string `json:"RCLASS_DESC"`
	RclassId    int64  `json:"RCLASS_ID"`
}

type CfgRtype struct {
	BreakRes    string `json:"BREAK_RES"`
	RclassId    int64  `json:"RCLASS_ID"`
	RelStrength int64  `json:"REL_STRENGTH"`
	RtypeCode   string `json:"RTYPE_CODE"`
	RtypeDesc   string `json:"RTYPE_DESC"`
	RtypeId     int64  `json:"RTYPE_ID"`
}

type CfgSfcall struct {
	ExecOrder int64 `json:"EXEC_ORDER"`
	FelemId   int64 `json:"FELEM_ID"`
	FtypeId   int64 `json:"FTYPE_ID"`
	SfcallId  int64 `json:"SFCALL_ID"`
	SfuncId   int64 `json:"SFUNC_ID"`
}

type CfgSfunc struct {
	ConnectStr    string `json:"CONNECT_STR"`
	FuncLib       string `json:"FUNC_LIB"`
	FuncVer       string `json:"FUNC_VER"`
	JavaClassName string `json:"JAVA_CLASS_NAME"`
	Language      string `json:"LANGUAGE"`
	SfuncCode     string `json:"SFUNC_CODE"`
	SfuncDesc     string `json:"SFUNC_DESC"`
	SfuncId       int64  `json:"SFUNC_ID"`
}

type Config struct {
	ConfigComments string `json:"CONFIG_COMMENTS"`
	ConfigId       int64  `json:"CONFIG_ID"`
	SysCreateDt    string `json:"SYS_CREATE_DT"`
}

type ConfigBaseVersion struct {
	BuildDate            string               `json:"BUILD_DATE"`
	BuildNumber          string               `json:"BUILD_NUMBER"`
	BuildVersion         string               `json:"BUILD_VERSION"`
	CompatibilityVersion CompatibilityVersion `json:"COMPATIBILITY_VERSION"`
	Version              string               `json:"VERSION"`
}

type Feature struct {
	FeatDesc       string         `json:"FEAT_DESC"`
	FeatDescValues []FeatureValue `json:"FEAT_DESC_VALUES"`
	LibFeatId      int64          `json:"LIB_FEAT_ID"`
	UsageType      string         `json:"USAGE_TYPE"`
}

type Features struct {
	Address    []Feature `json:"ADDRESS"`
	Dob        []Feature `json:"DOB"`
	Email      []Feature `json:"EMAIL"`
	Name       []Feature `json:"NAME"`
	Phone      []Feature `json:"PHONE"`
	RecordType []Feature `json:"RECORD_TYPE"`
}

type FeatureScore struct {
	CandidateFeat          string `json:"CANDIDATE_FEAT"`
	CandidateFeatId        int64  `json:"CANDIDATE_FEAT_ID"`
	CandidateFeatUsageType string `json:"CANDIDATE_FEAT_USAGE_TYPE"`
	FullScore              int64  `json:"FULL_SCORE"`
	InboundFeat            string `json:"INBOUND_FEAT"`
	InboundFeatId          int64  `json:"INBOUND_FEAT_ID"`
	InboundFeatUsageType   string `json:"INBOUND_FEAT_USAGE_TYPE"`
	ScoreBehavior          string `json:"SCORE_BEHAVIOR"`
	ScoreBucket            string `json:"SCORE_BUCKET"`
}

type FeatureScores struct {
	Address    []FeatureScore `json:"ADDRESS"`
	Dob        []FeatureScore `json:"DOB"`
	Email      []FeatureScore `json:"EMAIL"`
	Name       []FeatureScore `json:"NAME"`
	Phone      []FeatureScore `json:"PHONE"`
	RecordType []FeatureScore `json:"RECORD_TYPE"`
}

type FeatureValue struct {
	FeatDesc  string `json:"FEAT_DESC"`
	LibFeatId int64  `json:"LIB_FEAT_ID"`
}

type FinalState struct {
	NeedReevaluation int64           `json:"NEED_REEVALUATION"`
	VirtualEntities  []VirtualEntity `json:"VIRTUAL_ENTITIES"`
}

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
	ConfigBaseVersion   ConfigBaseVersion     `json:"CONFIG_BASE_VERSION"`
	SysOom              []SysOom              `json:"SYS_OOM"`
}

type HowResults struct {
	FinalState      FinalState       `json:"FINAL_STATE"`
	ResolutionSteps []ResolutionStep `json:"RESOLUTION_STEPS"`
}

type MatchInfo struct {
	ErruleCode    string        `json:"ERRULE_CODE"`
	FeatureScores FeatureScores `json:"FEATURE_SCORES"`
	MatchKey      string        `json:"MATCH_KEY"`
}

type MemberRecord struct {
	InternalId int64    `json:"INTERNAL_ID"`
	Records    []Record `json:"RECORDS"`
}

type Record struct {
	DataSource     string `json:"DATA_SOURCE"`
	EntityDesc     string `json:"ENTITY_DESC"`
	EntityKey      string `json:"ENTITY_KEY"`
	EntityType     string `json:"ENTITY_TYPE"`
	ErrultCode     string `json:"ERRULE_CODE"`
	InternalId     int64  `json:"INTERNAL_ID"`
	LastSeenDt     string `json:"LAST_SEEN_DT"`
	MatchKey       string `json:"MATCH_KEY"`
	MatchLevel     int64  `json:"MATCH_LEVEL"`
	MatchLevelCode string `json:"MATCH_LEVEL_CODE"`
	RecordId       string `json:"RECORD_ID"`
}

type RecordSummary struct {
	DataSource  string `json:"DATA_SOURCE"`
	FirstSeenDt string `json:"FIRST_SEEN_DT"`
	LastSeenDt  string `json:"LAST_SEEN_DT"`
	RecordCount int64  `json:"RECORD_COUNT"`
}

type RelatedEntity struct{}

type ResolutionStep struct {
	InboundVirtualEntityId string        `json:"INBOUND_VIRTUAL_ENTITY_ID"`
	MatchInfo              MatchInfo     `json:"MATCH_INFO"`
	ResultVirtualEntityId  string        `json:"RESULT_VIRTUAL_ENTITY_ID"`
	Step                   int64         `json:"STEP"`
	VirtualEntity1         VirtualEntity `json:"VIRTUAL_ENTITY_1"`
	VirtualEntity2         VirtualEntity `json:"VIRTUAL_ENTITY_2"`
}

type ResolvedEntity struct {
	EntityId      int64           `json:"ENTITY_ID"`
	EntityName    string          `json:"ENTITY_NAME"`
	Features      Features        `json:"FEATURES"`
	LastSeenDt    string          `json:"LAST_SEEN_DT"`
	Records       []Record        `json:"RECORDS"`
	RecordSummary []RecordSummary `json:"RECORD_SUMMARY"`
}

type ResolvedEntityFromSearch struct {
	MatchInfo MatchInfoFromSearch `json:"MATCH_INFO"`
	Entity    EntityFromSearch    `json:"ENTITY"`
}

type SchemaVersion struct {
	EngineSchemaVersion          string `json:"ENGINE_SCHEMA_VERSION"`
	MaximumRequiredSchemaVersion string `json:"MAXIMUM_REQUIRED_SCHEMA_VERSION"`
	MinimumRequiredSchemaVersion string `json:"MINIMUM_REQUIRED_SCHEMA_VERSION"`
}

type SysOom struct {
	FelemId    int64  `json:"FELEM_ID"`
	FtypeId    int64  `json:"FTYPE_ID"`
	LensId     int64  `json:"LENS_ID"`
	LibFeatId  int64  `json:"LIB_FEAT_ID"`
	LibFelemId int64  `json:"LIB_FELEM_ID"`
	NextThresh int64  `json:"NEXT_THRESH"`
	OomLevel   string `json:"OOM_LEVEL"`
	OomType    string `json:"OOM_TYPE"`
	Thresh1Cnt int64  `json:"THRESH1_CNT"`
	Thresh1Oom int64  `json:"THRESH1_OOM"`
}

type VirtualEntity struct {
	MemberRecords   []MemberRecord `json:"MEMBER_RECORDS"`
	VirtualEntityId string         `json:"VIRTUAL_ENTITY_ID"`
}

// ----------------------------------------------------------------------------
// Types - structs - Config
// ----------------------------------------------------------------------------

// The ConfigAddDataSourceResponse...
type ConfigAddDataSourceResponse struct {
	DsrcId int64 `json:"DSRC_ID"`
}

// The ConfigListDataSourcesResponse...
type ConfigListDataSourcesResponse struct {
	DataSources []DataSource `json:"DATA_SOURCES"`
}

// The ConfigSaveResponse...
type ConfigSaveResponse struct {
	G2Config G2Config `json:"G2_CONFIG"`
}

// ----------------------------------------------------------------------------
// Types - structs - Configmgr
// ----------------------------------------------------------------------------

// The ConfigmgrGetConfigResponse...
type ConfigmgrGetConfigResponse struct {
	G2Config G2Config `json:"G2_CONFIG"`
}

// The ConfigmgrGetConfigListResponse...
type ConfigmgrGetConfigListResponse struct {
	Configs []Config `json:"CONFIGS"`
}

// ----------------------------------------------------------------------------
// Types - structs - Diagnostic
// ----------------------------------------------------------------------------

type DiagnosticCheckDBPerfResponse struct{}
type DiagnosticFetchNextEntityBySizeResponse struct{}
type DiagnosticFindEntitiesByFeatureIDsResponse struct{}
type DiagnosticGetDataSourceCountsResponse struct{}
type DiagnosticGetDBInfoResponse struct{}
type DiagnosticGetEntityDetailsResponse struct{}
type DiagnosticGetEntityListBySizeResponse struct{}
type DiagnosticGetEntityResumeResponse struct{}
type DiagnosticGetEntitySizeBreakdownResponse struct{}
type DiagnosticGetFeatureResponse struct{}
type DiagnosticGetGenericFeaturesResponse struct{}
type DiagnosticGetMappingStatisticsResponse struct{}
type DiagnosticGetRelationshipDetailsResponse struct{}
type DiagnosticGetResolutionStatisticsResponse struct{}
type DiagnosticStreamEntityListBySizeResponse struct{}

// ----------------------------------------------------------------------------
// Types - structs - Engine
// ----------------------------------------------------------------------------

type EngineAddRecordWithInfoResponse struct{}
type EngineAddRecordWithInfoWithReturnedRecordIDResponse struct{}
type EngineAddRecordWithReturnedRecordIDResponse struct{}
type EngineCheckRecordResponse struct{}
type EngineDeleteRecordWithInfoResponse struct{}
type EngineExportConfigAndConfigIDResponse struct{}
type EngineExportConfigResponse struct{}
type EngineFetchNextResponse struct{}
type EngineFindInterestingEntitiesByEntityIDResponse struct{}
type EngineFindInterestingEntitiesByRecordIDResponse struct{}
type EngineFindNetworkByEntityIDV2Response struct{}
type EngineFindNetworkByEntityIDResponse struct{}
type EngineFindNetworkByRecordIDV2Response struct{}
type EngineFindNetworkByRecordIDResponse struct{}
type EngineFindPathByEntityIDV2Response struct{}
type EngineFindPathByEntityIDResponse struct{}
type EngineFindPathByRecordIDV2Response struct{}
type EngineFindPathByRecordIDResponse struct{}
type EngineFindPathExcludingByEntityIDV2Response struct{}
type EngineFindPathExcludingByEntityIDResponse struct{}
type EngineFindPathExcludingByRecordIDV2Response struct{}
type EngineFindPathExcludingByRecordIDResponse struct{}
type EngineFindPathIncludingSourceByEntityIDV2Response struct{}
type EngineFindPathIncludingSourceByEntityIDResponse struct{}
type EngineFindPathIncludingSourceByRecordIDV2Response struct{}
type EngineFindPathIncludingSourceByRecordIDResponse struct{}
type EngineGetEntityByEntityIDV2Response struct{}

type EngineGetEntityByEntityIDResponse struct {
	ResolvedEntity  ResolvedEntity  `json:"RESOLVED_ENTITY"`
	RelatedEntities []RelatedEntity `json:"RELATED_ENTITIES"`
}

type EngineGetEntityByRecordIDV2Response struct{}
type EngineGetEntityByRecordIDResponse struct{}
type EngineGetRecordV2Response struct{}
type EngineGetRecordResponse struct{}
type EngineGetRedoRecordResponse struct{}
type EngineGetVirtualEntityByRecordIDV2Response struct{}
type EngineGetVirtualEntityByRecordIDResponse struct{}
type EngineHowEntityByEntityIDV2Response struct{}

type EngineHowEntityByEntityIDResponse struct {
	HowResults HowResults `json:"HOW_RESULTS"`
}

type EngineProcessRedoRecordResponse struct{}
type EngineProcessRedoRecordWithInfoResponse struct{}
type EngineProcessWithInfoResponse struct{}
type EngineProcessWithResponseResizeResponse struct{}
type EngineProcessWithResponseResponse struct{}
type EngineReevaluateEntityWithInfoResponse struct{}
type EngineReevaluateRecordWithInfoResponse struct{}
type EngineReplaceRecordWithInfoResponse struct{}
type EngineSearchByAttributesV2Response struct{}

type EngineSearchByAttributesResponse struct {
	ResolvedEntities []ResolvedEntityFromSearch `json:"RESOLVED_ENTITIES"`
}

type EngineStatsResponse struct{}
type EngineStreamExportJSONEntityReportResponse struct{}
type EngineWhyEntitiesV2Response struct{}
type EngineWhyEntitiesResponse struct{}
type EngineWhyEntityByEntityIDV2Response struct{}
type EngineWhyEntityByEntityIDResponse struct{}
type EngineWhyEntityByRecordIDV2Response struct{}
type EngineWhyEntityByRecordIDResponse struct{}
type EngineWhyRecordsV2Response struct{}
type EngineWhyRecordsResponse struct{}

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

type ProductValidateLicenseFileResponse struct{}
type ProductValidateLicenseStringBase64Response struct{}

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
