package g2error

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2ErrorTypeIds int

// ----------------------------------------------------------------------------
// "Category" errors
// One of these errors must be the last in each G2ErrorTypes value list.
// ----------------------------------------------------------------------------

type G2BadInputError struct {
	error
	G2ErrorTypeIds []G2ErrorTypeIds
}
type G2BaseError struct {
	error
	G2ErrorTypeIds []G2ErrorTypeIds
}
type G2ConfigurationError struct {
	error
	G2ErrorTypeIds []G2ErrorTypeIds
}
type G2RetryableError struct {
	error
	G2ErrorTypeIds []G2ErrorTypeIds
}
type G2UnrecoverableError struct {
	error
	G2ErrorTypeIds []G2ErrorTypeIds
}

// ----------------------------------------------------------------------------
// Detail errors
// ----------------------------------------------------------------------------

type G2DatabaseConnectionLostError struct{ error }
type G2DatabaseError struct{ error }
type G2LicenseError struct{ error }
type G2NotFoundError struct{ error }
type G2NotInitializedError struct{ error }
type G2RetryTimeoutExceededError struct{ error }
type G2UnhandledError struct{ error }
type G2UnknownDatasourceError struct{ error }

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const (
	G2Base G2ErrorTypeIds = iota
	G2BadInput
	G2Configuration
	G2Database
	G2DatabaseConnectionLost
	G2License
	G2NotFound
	G2NotInitialized
	G2Retryable
	G2RetryTimeoutExceeded
	G2Unhandled
	G2UnknownDatasource
	G2Unrecoverable
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for g2engine implementations.
// Note: The lists of G2ErrorTypeIds are from innermost error to outer most error.
// Example:  #10 is G2RetryableError{G2RetryTimeoutExceededError{errors.New(message)}}
var G2ErrorTypes = map[int][]G2ErrorTypeIds{
	2:     {G2BadInput},                            // EAS_ERR_INVALID_XML                                                                    "Invalid XML"
	5:     {G2Base},                                // EAS_ERR_EXCEEDED_MAX_RETRIES                                                           "Exceeded the Maximum Number of Retries Allowed"
	7:     {G2BadInput},                            // EAS_ERR_EMPTY_XML_MESSAGE                                                              "Empty XML Message"
	10:    {G2RetryTimeoutExceeded, G2Retryable},   // EAS_ERR_RETRY_TIMEOUT                                                                  "Retry timeout exceeded"
	14:    {G2Configuration},                       // EAS_ERR_INVALID_DATASTORE_CONFIGURATION_TYPE                                           "Invalid Datastore Configuration Type"
	19:    {G2Configuration},                       // EAS_ERR_NO_CONFIGURATION_FOUND                                                         "Configuration not found"
	20:    {G2Configuration},                       // EAS_ERR_CONFIG_CANNOT_BE_NULL_DATABASE                                                 "Configuration cannot be loaded from database connection"
	21:    {G2Configuration},                       // EAS_ERR_CONFIG_CANNOT_BE_NULL_CONFIG_FILE                                              "Configuration cannot be loaded from config file"
	22:    {G2BadInput},                            // EAS_ERR_INVALID_DOCTYPE                                                                "Invalid DocType {0}"
	23:    {G2BadInput},                            // EAS_ERR_CONFLICTING_DATA_SOURCE_VALUES                                                 "Conflicting DATA_SOURCE values '{0}' and '{1}'"
	24:    {G2BadInput},                            // EAS_ERR_CONFLICTING_RECORD_ID_VALUES                                                   "Conflicting RECORD_ID values '{0}' and '{1}'"
	25:    {G2BadInput},                            // EAS_ERR_CONFLICTING_LOAD_ID_VALUES                                                     "Conflicting LOAD_ID values '{0}' and '{1}'"
	26:    {G2BadInput},                            // EAS_ERR_RESERVED_WORD_USED_IN_DOCUMENT                                                 "Inbound data contains a reserved keyword '{0}'"
	27:    {G2UnknownDatasource, G2BadInput},       // EAS_ERR_UNKNOWN_DSRC_CODE_VALUE                                                        "Unknown DATA_SOURCE value '{0}'"
	28:    {G2Configuration},                       // EAS_ERR_INVALID_JSON_CONFIG_DOCUMENT                                                   "Invalid JSON config document"
	29:    {G2Base},                                // EAS_ERR_INVALID_HANDLE                                                                 "Invalid Handle"
	30:    {G2Configuration},                       // EAS_ERR_INVALID_MATCH_LEVEL                                                            "Invalid match level '{0}'"
	33:    {G2NotFound, G2BadInput},                // EAS_ERR_UNKNOWN_DSRC_RECORD_ID                                                         "Unknown record: dsrc[{0}], record[{1}]"
	34:    {G2Configuration},                       // EAS_ERR_AMBIGUOUS_ENTITY_FTYPE_MISSING                                                 "AMBIGUOUS_ENTITY Feature Type is not configured"
	35:    {G2Configuration},                       // EAS_ERR_AMBIGUOUS_TIER_FELEM_MISSING                                                   "AMBIGUOUS_TIER Feature Element is not configured"
	36:    {G2Configuration},                       // EAS_ERR_AMBIGUOUS_FTYPE_ID_FELEM_MISSING                                               "AMBIGUOUS_FTYPE_ID Feature Element is not configured"
	37:    {G2NotFound, G2BadInput},                // EAS_ERR_UNKNOWN_RESOLVED_ENTITY_VALUE                                                  "Unknown resolved entity value '{0}'"
	38:    {G2Base},                                // EAS_ERR_RECORD_HAS_NO_RESOLVED_ENTITY                                                  "Data source record has no resolved entity: dsrc[{0}], recordID[{1}]"
	39:    {G2Base},                                // EAS_ERR_NO_OBSERVED_ENTITY_FOR_DSRC_ENTITY_KEY                                         "No observed entity for entity key: dsrc[{0}], record_id[{1}], key[{2}]"
	40:    {G2Configuration},                       // EAS_ERR_CONFIG_COMPATIBILITY_MISMATCH                                                  "The engine configuration compatibility version [{0}] does not match the version of the provided config[{1}]."
	41:    {G2Base},                                // EAS_ERR_DOCUMENT_PREPROCESSING_FAILED                                                  "Document preprocessing failed"
	42:    {G2Base},                                // EAS_ERR_DOCUMENT_LOAD_PROCESSING_FAILED                                                "Document load processing failed"
	43:    {G2Base},                                // EAS_ERR_DOCUMENT_ER_PROCESSING_FAILED                                                  "Document ER processing failed"
	44:    {G2Base},                                // EAS_ERR_CHECK_ENTITY_PROCESSING_FAILED                                                 "Check entity processing failed"
	45:    {G2Base},                                // EAS_ERR_UMF_PROC_PROCESSING_FAILED                                                     "UMF procedure processing failed"
	46:    {G2Base},                                // EAS_ERR_DOCUMENT_HASHING_PROCESSING_FAILED                                             "Document hashing-processing failed"
	47:    {G2Base},                                // EAS_ERR_SESSION_IS_INVALID                                                             "Session is invalid"
	48:    {G2NotInitialized, G2Unrecoverable},     // EAS_ERR_G2_NOT_INITIALIZED                                                             "G2 is not initialized"
	49:    {G2NotInitialized, G2Unrecoverable},     // EAS_ERR_G2AUDIT_NOT_INITIALIZED                                                        "G2Audit is not initialized"
	50:    {G2NotInitialized, G2Unrecoverable},     // EAS_ERR_G2HASHER_NOT_INITIALIZED                                                       "G2Hasher is not initialized"
	51:    {G2BadInput},                            // EAS_ERR_BOTH_RECORD_ID_AND_ENT_SRC_KEY_SPECIFIED                                       "Cannot use both Record ID and Entity Source Key in record"
	52:    {G2Base},                                // EAS_ERR_UNKNOWN_RELATIONSHIP_ID_VALUE                                                  "Unknown relationship ID value '{0}'"
	53:    {G2BadInput},                            // EAS_ERR_G2DIAGNOSTIC_NOT_INITIALIZED                                                   "G2Diagnostic is not initialized"
	54:    {G2Database, G2Unrecoverable},           // EAS_ERR_G2_DATA_REPOSITORY_WAS_PURGED                                                  "Data repository was purged"
	55:    {G2Base},                                // EAS_ERR_NO_RESOLVED_ENTITY_FOR_DSRC_ENTITY_KEY                                         "No resolved entity for entity key: dsrc[{0}], record_id[{1}], key[{2}]"
	56:    {G2Base},                                // EAS_ERR_NO_RECORDS_EXIST_FOR_RESOLVED_ENTITY                                           "No data source records exist for entity ID: entityID[{0}]"
	57:    {G2Base},                                // EAS_ERR_UNKNOWN_FEATURE_ID_VALUE                                                       "Unknown feature ID value '{0}'"
	58:    {G2Base},                                // EAS_ERR_G2_INITIALIZATION_FAILURE                                                      "G2 initialization process has failed"
	60:    {G2Configuration},                       // EAS_ERR_CONFIG_DATABASE_MISMATCH                                                       "The engine configuration does not match the records loaded into the repository:  errors[{0}]."
	61:    {G2Configuration},                       // EAS_ERR_AMBIGUOUS_SUPPRESSED_LIBFEAT_FELEM_MISSING                                     "AMBIGUOUS_SUPRESSED_LIBFEAT Feature Element is not configured"
	62:    {G2Configuration},                       // EAS_ERR_AMBIGUOUS_TYPE_FELEM_MISSING                                                   "AMBIGUOUS_TYPE Feature Element is not configured"
	63:    {G2NotInitialized, G2Unrecoverable},     // EAS_ERR_G2CONFIGMGR_NOT_INITIALIZED                                                    "G2ConfigMgr is not initialized"
	64:    {G2Configuration},                       // EAS_ERR_CONFUSED_ENTITY_FTYPE_MISSING                                                  "CONFUSED_ENTITY Feature Type is not configured"
	65:    {G2BadInput},                            // EAS_ERR_UNKNOWN_ENTITY_TYPE_ID                                                         "Unknown entity type ID '{0}'"
	66:    {G2BadInput},                            // EAS_ERR_UNKNOWN_GENERIC_PLAN_VALUE                                                     "Unknown generic plan value '{0}'"
	67:    {G2Configuration},                       // EAS_ERR_INVALID_GENERIC_PLAN_VALUE                                                     "Invalid Generic Plan ID [{0}] configured for the '{1}' retention level.'"
	68:    {G2Base},                                // EAS_ERR_UNKNOWN_ER_RESULT                                                              "Unknown ER-result."
	69:    {G2Base},                                // EAS_ERR_NO_CANDIDATES                                                                  "No candidates."
	76:    {G2Base},                                // EAS_ERR_INBOUND_FEATURE_VERSION_NEWER_THAN_CONFIG                                      "Inbound Feature Version [{0}] is newer than configured version [{1}] for feature type[{2}]."
	77:    {G2Base},                                // EAS_ERR_ERROR_WHEN_PRIMING_GNR                                                         "Error when priming GNR resources '{0}'"
	78:    {G2Base},                                // EAS_ERR_ERROR_WHEN_ENCRYPTING                                                          "Error when encrypting '{0}'"
	79:    {G2Base},                                // EAS_ERR_ERROR_WHEN_DECRYPTING                                                          "Error when decryting '{0}'"
	80:    {G2Base},                                // EAS_ERR_ERROR_WHEN_VALIDATING_ENCRYPTION_SIGNATURE_COMPATIBILITY                       "Error when validating encryption signature compatibility '{0}'"
	81:    {G2Base},                                // EAS_ERR_ERROR_WHEN_CHECKING_DISTINCT_FEATURE_GENERALIZATION                            "Error when checking distinct feature generalization '{0}'"
	82:    {G2Base},                                // EAS_ERR_ERROR_WHEN_RUNNING_DQM                                                         "Error when running DQM '{0}'"
	83:    {G2Base},                                // EAS_ERR_ERROR_WHEN_CREATING_EFEATS                                                     "Error when creating EFEATS '{0}'"
	84:    {G2Base},                                // EAS_ERR_ERROR_WHEN_SIMPLE_SCORING                                                      "Error when simple scoring '{0}'"
	85:    {G2Base},                                // EAS_ERR_ERROR_WHEN_SCORING_PAIR                                                        "Error when scoring a pair '{0}'"
	86:    {G2Base},                                // EAS_ERR_ERROR_WHEN_SCORING_SET                                                         "Error when scoring a set '{0}'"
	87:    {G2Unhandled, G2Unrecoverable},          // EAS_ERR_SRD_EXCEPTION                                                                  "SRD Exception '{0}'"
	88:    {G2BadInput},                            // EAS_ERR_UNKNOWN_SEARCH_PROFILE_VALUE                                                   "Unknown search profile value '{0}'"
	89:    {G2Configuration},                       // EAS_ERR_MISCONFIGURED_SEARCH_PROFILE_VALUE                                             "Misconfigured search profile value '{0}'"
	90:    {G2Configuration},                       // EAS_ERR_CANNOT_ADD_LIBRARY_FEATURES_TO_DATASTORE                                       "Cannot add library features to datastore"
	999:   {G2License, G2Unrecoverable},            // EAS_ERR_LICENSE_HAS_EXPIRED                                                            "License has expired"
	1000:  {G2Database, G2Unrecoverable},           // EAS_ERR_UNHANDLED_DATABASE_ERROR                                                       "Unhandled Database Error '{0}'"
	1001:  {G2Database, G2Unrecoverable},           // EAS_ERR_CRITICAL_DATABASE_ERROR                                                        "Critical Database Error '{0}'"
	1002:  {G2Database, G2Unrecoverable},           // EAS_ERR_DATABASE_MEMORY_ERROR                                                          "Database Memory Error '{0}'"
	1003:  {G2Database, G2Unrecoverable},           // EAS_ERR_TABLE_SPACE_OR_LOG_VIOLATION                                                   "Table Space or Log Violation '{0}'"
	1004:  {G2Database, G2Unrecoverable},           // EAS_ERR_RESOURCE_CONTENTION                                                            "Resource Contention '{0}'"
	1005:  {G2Database, G2Unrecoverable},           // EAS_ERR_USER_DEFINED_PROC_ERROR                                                        "User Defined Procedure or Function Error '{0}'"
	1006:  {G2DatabaseConnectionLost, G2Retryable}, // EAS_ERR_DATABASE_CONNECTION_FAILURE                                                    "Database Connection Failure '{0}'"
	1007:  {G2DatabaseConnectionLost, G2Retryable}, // EAS_ERR_DATABASE_CONNECTION_LOST                                                       "Database Connection Lost '{0}'"
	1008:  {G2Database, G2Unrecoverable},           // EAS_ERR_DEADLOCK_ERROR                                                                 "Deadlock Error '{0}'"
	1009:  {G2Database, G2Unrecoverable},           // EAS_ERR_INSUFFICIENT_PERMISSIONS                                                       "Insufficient Permissions '{0}'"
	1010:  {G2Database, G2Unrecoverable},           // EAS_ERR_TRANSACTION_ERROR                                                              "Transaction Error '{0}'"
	1011:  {G2Database, G2Unrecoverable},           // EAS_ERR_UNIQUE_CONSTRAINT_VIOLATION                                                    "Unique Constraint Violation '{0}'"
	1012:  {G2Database, G2Unrecoverable},           // EAS_ERR_CONSTRAINT_VIOLATION                                                           "Constraint Violation '{0}'"
	1013:  {G2Database, G2Unrecoverable},           // EAS_ERR_SYNTAX_ERROR                                                                   "Syntax Error '{0}'"
	1014:  {G2Database, G2Unrecoverable},           // EAS_ERR_CURSOR_ERROR                                                                   "Cursor Error '{0}'"
	1015:  {G2Database, G2Unrecoverable},           // EAS_ERR_DATATYPE_ERROR                                                                 "Data Type Error '{0}'"
	1016:  {G2Database, G2Unrecoverable},           // EAS_ERR_TRANSACTION_ABORTED_ERROR                                                      "Transaction Aborted '{0}'"
	1017:  {G2Database, G2Unrecoverable},           // EAS_ERR_DATABASE_OPERATOR_NOT_SET                                                      "Database operator not set '{0}'"
	1018:  {G2Database, G2Unrecoverable},           // EAS_ERR_DATABASE_EXCEPTION_GENERATOR_NOT_SET                                           "Database exception generator not set '{0}'"
	1019:  {G2Configuration},                       // EAS_ERR_DATABASE_SCHEMA_TABLES_NOT_FOUND                                               "Datastore schema tables not found. [{0}]"
	2001:  {G2Configuration},                       // EAS_ERR_FEATURE_HAS_NO_FTYPE_CODE                                                      "Cannot process feature with no FTYPE_CODE[{0}]"
	2002:  {G2Base},                                // EAS_ERR_REQUESTED_CONFIG_FOR_INVALID_FTYPE_CODE                                        "Requested config for invalid FTYPE_CODE[{0}]"
	2003:  {G2Base},                                // EAS_ERR_NO_FELEM_CODE                                                                  "Cannot process OBS_FELEM with no FELEM_CODE[{0}]"
	2005:  {G2Base},                                // EAS_ERR_INVALID_FELEM_CODE                                                             "FELEM_CODE[{0}] is not configured for FTYPE_CODE[{1}]"
	2006:  {G2Base},                                // EAS_ERR_MISSING_ENT_SRC_KEY                                                            "OBS_ENT is missing ENT_SRC_KEY"
	2007:  {G2Base},                                // EAS_ERR_MISSING_OBS_SRC_KEY                                                            "OBS is missing OBS_SRC_KEY"
	2012:  {G2Configuration},                       // EAS_ERR_ERRULE_CONFIGURED_FOR_RESOLVE_AND_RELATE                                       "ER Rule [{0}] is configured for both resolve and relate."
	2015:  {G2Configuration},                       // EAS_ERR_INVALID_FTYPE_CODE                                                             "Invalid FTYPE_CODE[{0}]"
	2027:  {G2Base},                                // EAS_ERR_PLUGIN_INIT                                                                    "Plugin initialization error {0}"
	2029:  {G2Configuration},                       // EAS_ERR_REQUESTED_CONFIG_FOR_INVALID_PLUGIN                                            "Configuration not found for plugin type: {0}"
	2034:  {G2Configuration},                       // EAS_ERR_INVALID_CFRTN_VAL                                                              "CFRTN_ID[{0}]/FTYPE[{1}] is expecting CFRTN_VAL[{2}] which is not offered by CFUNC_ID[{3}][{4}]. Available scores are [{5}]"
	2036:  {G2Configuration},                       // EAS_ERR_FTYPE_HAS_NO_BOM                                                               "FType configured with no Feature Elements (Bill of Materials)  FTYPE_ID[{0}] FTYPE_CODE[{1}]"
	2037:  {G2Configuration},                       // EAS_ERR_FUNC_CALL_HAS_NO_BOM                                                           "Function call ({3}) configured with no Bill of Materials  {4}[{0}] FTYPE_ID[{1}] FTYPE_CODE[{2}]"
	2038:  {G2Configuration},                       // EAS_ERR_DISTINCT_FEATURE_HAS_NO_BOM                                                    "Distinct feature call configured with no Bill of Materials  DFCALL_ID[{0}]"
	2041:  {G2Configuration},                       // EAS_ERR_EFCALL_HAS_NO_BOM                                                              "EFeature creation call configured with no Bill of Materials  EFCALL_ID[{0}]"
	2045:  {G2Configuration},                       // EAS_ERR_CFRTN_REFERS_BAD_CFUNC_ID                                                      "CFG_CFRTN references CFUNC_ID[{0}] which is not configured"
	2047:  {G2Configuration},                       // EAS_ERR_MISSING_DSRC_CODE                                                              "OBS_SRC_KEY[{0}] is missing DSRC_CODE tag which is required"
	2048:  {G2Configuration},                       // EAS_ERR_FEAT_FREQ_INVALID                                                              "FEATURE CODE[{0}] FEATURE FREQUENCY[{1}] is an invalid frequency"
	2049:  {G2Configuration},                       // EAS_ERR_FUNC_INVALID                                                                   "{2} [{0}] is invalid for {3}[{1}]"
	2050:  {G2Configuration},                       // EAS_ERR_QUAL_FRAG_NOT_FOUND                                                            "Rule[{0}] Qualifier Fragment[{1}]: Fragment not found"
	2051:  {G2Configuration},                       // EAS_ERR_DISQUAL_FRAG_NOT_FOUND                                                         "Rule[{0}] Disqualifier Fragment[{1}]: Fragment not found"
	2057:  {G2BadInput},                            // EAS_ERR_BAD_DSRC_ACTION                                                                "OBS_SRC_KEY[{0}] has DSRC_ACTION[{1}] which is invalid.  Valid values are [A]dd, [C]hange, [D]elete, [P]rune or E[X]tensive Evaluation"
	2061:  {G2Configuration},                       // EAS_ERR_DUPLICATE_LOOKUP_IDENTIFIER                                                    "Duplicate [{0}] with identifier value [{1}].  Only unique values are allowed."
	2062:  {G2Configuration},                       // EAS_ERR_INVALID_LOOKUP_IDENTIFIER                                                      "Requested lookup of [{0}] using unknown value [{1}].  Value not found."
	2065:  {G2Configuration},                       // EAS_ERR_FTYPE_HAS_MULTIPLE_DEFINITIONS                                                 "FType configured with multiple definitions. FTYPE_CODE[{0}] used in FTYPE_ID[{1}] and FTYPE_ID[{2}]"
	2066:  {G2Configuration},                       // EAS_ERR_FELEM_HAS_MULTIPLE_DEFINITIONS                                                 "FElem configured with multiple definitions. FELEM_CODE[{0}] used in FELEM_ID[{1}] and FELEM_ID[{2}]"
	2067:  {G2Configuration},                       // EAS_ERR_ERFRAG_HAS_MULTIPLE_DEFINITIONS                                                "ER Fragment code configured with multiple definitions. ERFRAG_CODE[{0}] used in ERFRAG_ID[{1}] and ERFRAG_ID[{2}]"
	2069:  {G2Configuration},                       // EAS_ERR_BOM_CONFIG_INVALID_FOR_SIMPLE_PLUGIN                                           "Configured plugin for CFCALL_ID[{0}] requires exactly one value in BOM"
	2070:  {G2Configuration},                       // EAS_ERR_EFCALL_HAS_INVALID_FUNCTION                                                    "EFeature creation call configured with invalid function ID EFCALL_ID[{0}] EFUNC_ID[{1}]"
	2071:  {G2Configuration},                       // EAS_ERR_EFBOM_HAS_INVALID_EFCALL                                                       "EFeature BOM configured with invalid EFCALL_ID[{0}]"
	2073:  {G2Base},                                // EAS_ERR_LOADING_LIBRARY                                                                "Library loading error {0}"
	2074:  {G2Base},                                // EAS_ERR_SCORING_MANAGER_PLUGIN                                                         "Scoring manager: id {0} and {1} do not match"
	2075:  {G2Configuration},                       // EAS_ERR_TABLE_CONFIGURED_WITH_INVALID_FTYPE_CODE                                       "Table {0} configured with an invalid type FTYPE_CODE[{1}]"
	2076:  {G2Configuration},                       // EAS_ERR_TABLE_CONFIGURED_WITH_INVALID_FELEM_CODE                                       "Table {0} configured with an invalid type FELEM_CODE[{1}]"
	2077:  {G2Configuration},                       // EAS_ERR_TABLE_CONFIGURED_WITH_INVALID_ETYPE_ID                                         "Table {0} configured with an invalid type ETYPE_ID[{1}]"
	2079:  {G2Configuration},                       // EAS_ERR_EFBOM_CONFIGURED_WITH_INVALID_FTYPE_ID                                         "CFG_EFBOM configured with an invalid type FTYPE_ID[{0}]"
	2080:  {G2Configuration},                       // EAS_ERR_EFBOM_CONFIGURED_WITH_INVALID_FELEM_ID                                         "CFG_EFBOM configured with an invalid type FELEM_ID[{0}]"
	2081:  {G2Configuration},                       // EAS_ERR_FUNC_CALL_CONFIGURED_WITH_INVALID_FTYPE_ID                                     "{1} configured with an invalid type FTYPE_ID[{0}]"
	2082:  {G2Configuration},                       // EAS_ERR_FUNC_CALL_CONFIGURED_WITH_INVALID_FUNC_ID                                      "{1} configured with an invalid type {2}[{0}]"
	2083:  {G2Configuration},                       // EAS_ERR_FUNC_BOM_CONFIGURED_WITH_INVALID_FTYPE_ID                                      "{1} configured with an invalid type FTYPE_ID[{0}]"
	2084:  {G2Configuration},                       // EAS_ERR_FUNC_BOM_CONFIGURED_WITH_INVALID_FELEM_ID                                      "{1} configured with an invalid type FELEM_ID[{0}]"
	2087:  {G2Configuration},                       // EAS_ERR_TABLE_CONFIGURED_WITH_DUPLICATE_CONFIG_FOR_EFU_TYPE                            "Table [{0}] configured with duplicate config values for ETYPE_ID[{1}]/FTYPE_ID[{2}]/UTYPE_CODE[{3}]"
	2088:  {G2Configuration},                       // EAS_ERR_TABLE_CONFIGURED_WITH_INVALID_RCLASS_ID                                        "Table {0} configured with an invalid RCLASS_ID[{1}]"
	2089:  {G2Configuration},                       // EAS_ERR_UNKNOWN_FCLASS_ID                                                              "UNKNOWN FCLASS ID[{0}]"
	2090:  {G2Configuration},                       // EAS_ERR_SFCALL_HAS_INVALID_FUNCTION                                                    "Feature standardization call configured with invalid function ID SFCALL_ID[{0}] SFUNC_ID[{1}]"
	2091:  {G2Configuration},                       // EAS_ERR_TABLE_CONFIGURED_WITH_BOTH_FTYPE_ID_AND_FELEM_ID                               "{0} configured with both an FTYPE_ID[{1}] and FELEM_ID[{2}]"
	2092:  {G2Configuration},                       // EAS_ERR_TABLE_CONFIGURED_WITH_NEITHER_FTYPE_ID_NOR_FELEM_ID                            "{0} configured with neither an FTYPE_ID nor an FELEM_ID"
	2093:  {G2Configuration},                       // EAS_ERR_TABLE_CONFIGURED_WITH_DUPLICATE_EXEC_ORDER_FOR_IDENTIFIER_LIST                 "Table [{0}] configured with duplicate execution order value [{3}] for identifiers[{1}] with values [{2}]"
	2094:  {G2Configuration},                       // EAS_ERR_DUPLICATE_VALUE_FOR_FIELD_IN_TABLE                                             "Duplicate value [{2}] of field [{1}] in config [{0}]"
	2095:  {G2Configuration},                       // EAS_ERR_TABLE_CONFIGURED_WITH_INVALID_FTYPE_CODE_FELEM_CODE_PAIR                       "Table {0} configured with an invalid FTYPE_CODE[{1}]/FELEM_CODE[{2}] pair"
	2099:  {G2Configuration},                       // EAS_ERR_COUNTER_CONFIG_INVALID_THRESHOLD                                               "Next Threshold for a counter should be no less than 10, but has NEXT_THRESH{0}"
	2101:  {G2Configuration},                       // EAS_ERR_XPATH_OP_UNSUPPORTED                                                           "XPath operation unsupported [{0}]"
	2102:  {G2Configuration},                       // EAS_ERR_XPATH_AXIS_UNSUPPORTED                                                         "XPath axis unsupported [{0}]"
	2103:  {G2Configuration},                       // EAS_ERR_XPATH_TEST_UNSUPPORTED                                                         "XPath test unsupported [{0}]"
	2104:  {G2Configuration},                       // EAS_ERR_XPATH_TYPE_UNSUPPORTED                                                         "XPath type unsupported [{0}]"
	2105:  {G2Configuration},                       // EAS_ERR_XPATH_NODE_PREFIX_UNSUPPORTED                                                  "XPath node prefix unsupported [{0}]"
	2106:  {G2Configuration},                       // EAS_ERR_XPATH_NODE_NAME_UNSUPPORTED                                                    "XPath node name unsupported position[{0}], name[{1}]"
	2107:  {G2Configuration},                       // EAS_ERR_XPATH_BEHAVIOR_TYPE_UNSUPPORTED                                                "XPath behavior type unsupported [{0}]"
	2108:  {G2Configuration},                       // EAS_ERR_XPATH_BUCKET_UNSUPPORTED                                                       "XPath bucket type unsupported [{0}]"
	2109:  {G2Configuration},                       // EAS_ERR_XPATH_VALUE_TYPE_UNSUPPORTED                                                   "XPath value type unsupported [{0}]"
	2110:  {G2Configuration},                       // EAS_ERR_XPATH_PLUS_TYPE_UNSUPPORTED                                                    "XPath plus operand type unsupported [{0}]"
	2111:  {G2Configuration},                       // EAS_ERR_XPATH_FRAGMENT_NOT_EVALUATED                                                   "XPath fragment not evaluated[{0}]"
	2112:  {G2Configuration},                       // EAS_ERR_XPATH_FRAGMENT_NOT_CONFIGURED                                                  "XPath fragment not configured[{0}]"
	2113:  {G2Configuration},                       // EAS_ERR_XPATH_FUNCTION_UNSUPPORTED                                                     "XPath function unsupported [{0}]"
	2114:  {G2Configuration},                       // EAS_ERR_INVALID_FTYPE_SCORESET                                                         "Cannot set score for invalid FTYPE_ID [{0}]"
	2116:  {G2Base},                                // EAS_ERR_UNITIALIZED_AMBIGUOUS_CACHE                                                    "Uninitialized Ambiguous Test Cache"
	2117:  {G2Configuration},                       // EAS_ERR_SCORING_CALL_HAS_NO_BOM                                                        "Scoring call configured with no Bill of Materials  CFCALL_ID[{0}]."
	2118:  {G2Configuration},                       // EAS_ERR_BOM_CONFIG_INVALID_FOR_SCORING_PLUGIN                                          "Configured plugin for CFCALL_ID[{0}] has invalid BOM."
	2120:  {G2Configuration},                       // EAS_ERR_TABLE_CONFIGURED_WITH_INVALID_FTYPE_ID                                         "Table {0} configured with an invalid type FTYPE_ID[{1}]"
	2121:  {G2Configuration},                       // EAS_ERR_TABLE_CONFIGURED_WITH_INVALID_FELEM_ID                                         "Table {0} configured with an invalid type FELEM_ID[{1}]"
	2123:  {G2Configuration},                       // EAS_ERR_CFUNC_CONFIGURED_WITH_NO_CFRTN                                                 "CFG_CFUNC [{0}] feature type [{1}] configured without any corresponding return values in CFG_CFRTN"
	2131:  {G2Configuration},                       // EAS_ERR_OBS_ENT_NOT_FOUND                                                              "Requested resolution of OBS_ENT_ID that is not loaded OBS_ENT_ID[{0}]"
	2135:  {G2Configuration},                       // EAS_ERR_UMF_MAPPING_CONFIG_ERROR                                                       "Error in UMF Mapping Config[{0}]"
	2136:  {G2Configuration},                       // EAS_ERR_UMF_MAPPING_MISSING_REQUIRED_FIELD                                             "Error in UMF Mapping, missing required field[{0}]"
	2137:  {G2Configuration},                       // EAS_ERR_UMF_MAPPING_MALFORMED_INPUT                                                    "Error in UMF Mapping, input message is malformed[{0}]"
	2138:  {G2Configuration},                       // EAS_ERR_INVALID_CFRTN_INDEX                                                            "CFRTN_ID[{0}] is out of range. Valid range is 0-7"
	2139:  {G2Configuration},                       // EAS_ERR_DSRC_INTEREST_CONFIGURED_WITH_INVALID_DSRCID                                   "Data Source Interest configured with invalid Data Source ID DSRC_ID[{0}]"
	2205:  {G2Configuration},                       // EAS_ERR_ENTITY_TYPE_CODE_ALREADY_EXISTS                                                "Entity type code [{0}] already exists."
	2206:  {G2Configuration},                       // EAS_ERR_ENTITY_TYPE_ID_ALREADY_EXISTS                                                  "Entity type ID [{0}] already exists."
	2207:  {G2Configuration},                       // EAS_ERR_DATA_SOURCE_CODE_DOES_NOT_EXIST                                                "Data source code [{0}] does not exist."
	2208:  {G2Configuration},                       // EAS_ERR_DATA_SOURCE_CODE_ALREADY_EXISTS                                                "Data source code [{0}] already exists."
	2209:  {G2Configuration},                       // EAS_ERR_DATA_SOURCE_ID_ALREADY_EXISTS                                                  "Data source ID [{0}] already exists."
	2210:  {G2Configuration},                       // EAS_ERR_FELEM_CODE_DOES_NOT_EXIST                                                      "Feature element code [{0}] does not exist."
	2211:  {G2Configuration},                       // EAS_ERR_FELEM_CODE_ALREADY_EXISTS                                                      "Feature element code [{0}] already exists."
	2212:  {G2Configuration},                       // EAS_ERR_FELEM_ID_ALREADY_EXISTS                                                        "Feature element ID [{0}] already exists."
	2213:  {G2Configuration},                       // EAS_ERR_INVALID_FELEM_DATA_TYPE                                                        "Invalid feature element datatype [{0}] found.  Datatype must be in [{1}]."
	2214:  {G2Configuration},                       // EAS_ERR_FELEM_IS_CONFIGURED_FOR_USE_IN_FEATURES                                        "Feature element [{0}] is configured for use in feature(s) [{1}]."
	2215:  {G2Configuration},                       // EAS_ERR_FTYPE_CODE_DOES_NOT_EXIST                                                      "Feature type code [{0}] does not exist."
	2216:  {G2Configuration},                       // EAS_ERR_FTYPE_CODE_ALREADY_EXISTS                                                      "Feature type code [{0}] already exists."
	2217:  {G2Configuration},                       // EAS_ERR_FTYPE_ID_ALREADY_EXISTS                                                        "Feature type ID [{0}] already exists."
	2218:  {G2Configuration},                       // EAS_ERR_FEATURE_FREQUENCY_IS_INVALID                                                   "Feature type frequency [{0}] is invalid."
	2219:  {G2Configuration},                       // EAS_ERR_FEATURE_ELEMENT_LIST_IS_EMPTY                                                  "Feature element list is empty."
	2220:  {G2Configuration},                       // EAS_ERR_STANDARDIZATION_FUNCTION_DOES_NOT_EXIST                                        "Standardization function [{0}] does not exist."
	2221:  {G2Configuration},                       // EAS_ERR_FUNCTION_USES_BOTH_FTYPE_AND_FELEM_TRIGGER                                     "Function call requested uses both triggering feature type [{0}] and triggering feature element code [{1}].  Cannot use both triggering feature type and triggering feature element code."
	2222:  {G2Configuration},                       // EAS_ERR_EXPRESSION_FUNCTION_DOES_NOT_EXIST                                             "Expression function [{0}] does not exist."
	2223:  {G2Configuration},                       // EAS_ERR_EXPRESSION_FUNCTION_FEATURE_ELEMENT_LIST_IS_EMPTY                              "Expression function feature element list is empty."
	2224:  {G2Configuration},                       // EAS_ERR_COMPARISON_FUNCTION_DOES_NOT_EXIST                                             "Comparison function [{0}] does not exist."
	2225:  {G2Configuration},                       // EAS_ERR_COMPARISON_FUNCTION_FEATURE_ELEMENT_LIST_IS_EMPTY                              "Comparison function feature element list is empty."
	2226:  {G2Configuration},                       // EAS_ERR_DISTINCT_FUNCTION_DOES_NOT_EXIST                                               "Distinct feature function [{0}] does not exist."
	2227:  {G2Configuration},                       // EAS_ERR_DISTINCT_FUNCTION_FEATURE_ELEMENT_LIST_IS_EMPTY                                "Distinct feature function feature element list is empty."
	2228:  {G2Configuration},                       // EAS_ERR_FELEM_CODE_MUST_BE_UNIQUE_IN_FELEM_LIST                                        "Feature element code [{0}] must be unique in felem list."
	2230:  {G2Configuration},                       // EAS_ERR_FTYPE_CODE_AND_FELEM_CODE_MUST_BE_UNIQUE_IN_EXPRESSED_FUNCTION_CALL            "Feature type [{0}] and feature element [{1}] must be unique in expressed feature function call."
	2231:  {G2Configuration},                       // EAS_ERR_FTYPE_CODE_AND_FELEM_CODE_IN_EXPRESSED_FUNCTION_CALL_DO_NOT_EXIST_IN_FEATURE   "Feature type [{0}] and feature element [{1}] requested for expressed feature function call, but don't exist in feature [{0}]."
	2232:  {G2Configuration},                       // EAS_ERR_FELEM_CODE_MUST_BE_UNIQUE_IN_COMPARISON_FUNCTION_CALL                          "Feature element [{0}] must be unique in comparison feature function call."
	2233:  {G2Configuration},                       // EAS_ERR_FELEM_CODE_IN_COMPARISON_FUNCTION_CALL_DOES_NOT_EXIST_IN_FEATURE               "Feature element [{0}] requested for comparison feature function call, but doesn't exist in feature [{1}]."
	2234:  {G2Configuration},                       // EAS_ERR_FELEM_CODE_MUST_BE_UNIQUE_IN_DISTINCT_FUNCTION_CALL                            "Feature element [{0}] must be unique in distinct feature function call."
	2235:  {G2Configuration},                       // EAS_ERR_FELEM_CODE_IN_DISTINCT_FUNCTION_CALL_DOES_NOT_EXIST_IN_FEATURE                 "Feature element [{0}] requested for distinct feature function call, but doesn't exist in feature [{1}]."
	2236:  {G2Configuration},                       // EAS_ERR_EXEC_ORDER_IS_NOT_SPECIFIED_FOR_FUNCTION                                       "Exec order not specified for function."
	2237:  {G2Configuration},                       // EAS_ERR_SFCALL_ID_ALREADY_EXISTS                                                       "Standardization function call ID [{0}] already exists."
	2238:  {G2Configuration},                       // EAS_ERR_EFCALL_ID_ALREADY_EXISTS                                                       "Expression function call ID [{0}] already exists."
	2239:  {G2Configuration},                       // EAS_ERR_CFCALL_ID_ALREADY_EXISTS                                                       "Comparison function call ID [{0}] already exists."
	2240:  {G2Configuration},                       // EAS_ERR_DFCALL_ID_ALREADY_EXISTS                                                       "Distinct feature function call ID [{0}] already exists."
	2241:  {G2Configuration},                       // EAS_ERR_FTYPE_CODE_REQUIRED_BY_SEPARATE_EXPRESSED_FUNCTION_CALL                        "Feature type [{0}] required for separate expressed feature function call [{1}]."
	2242:  {G2Configuration},                       // EAS_ERR_SFCALL_ID_DOES_NOT_EXIST                                                       "Standardization function call ID [{0}] does not exist."
	2243:  {G2Configuration},                       // EAS_ERR_EFCALL_ID_DOES_NOT_EXIST                                                       "Expression function call ID [{0}] does not exist."
	2244:  {G2Configuration},                       // EAS_ERR_CFCALL_ID_DOES_NOT_EXIST                                                       "Comparison function call ID [{0}] does not exist."
	2245:  {G2Configuration},                       // EAS_ERR_DFCALL_ID_DOES_NOT_EXIST                                                       "Distinct feature function call ID [{0}] does not exist."
	2246:  {G2Configuration},                       // EAS_ERR_BOM_EXEC_ORDER_ALREADY_EXISTS                                                  "BOM exec order value [{0}] already exists."
	2247:  {G2Configuration},                       // EAS_ERR_COMPARISON_FUNCTION_CALL_DOES_NOT_EXIST_FOR_FEATURE                            "Comparison function call does not exist for feature [{0}]."
	2248:  {G2Configuration},                       // EAS_ERR_DISTINCT_FUNCTION_CALL_DOES_NOT_EXIST_FOR_FEATURE                              "Distinct feature function call does not exist for feature [{0}]."
	2249:  {G2Configuration},                       // EAS_ERR_CONFLICTING_SPECIFIERS_FOR_FUNCTION_CALL                                       "Conflicting specifiers: Function call ID [{0}] does not match function call ID [{1}] from feature type."
	2250:  {G2Configuration},                       // EAS_ERR_ATTR_CODE_DOES_NOT_EXIST                                                       "Attribute code [{0}] does not exist."
	2251:  {G2Configuration},                       // EAS_ERR_ATTR_CODE_ALREADY_EXISTS                                                       "Attribute code [{0}] already exists."
	2252:  {G2Configuration},                       // EAS_ERR_ATTR_ID_ALREADY_EXISTS                                                         "Attribute ID [{0}] already exists."
	2253:  {G2Configuration},                       // EAS_ERR_ATTR_CLASS_CODE_DOES_NOT_EXIST                                                 "Attribute class code [{0}] does not exist."
	2254:  {G2Configuration},                       // EAS_ERR_FUNCTION_USES_NEITHER_FTYPE_NOR_FELEM_TRIGGER                                  "Function call requested uses neither triggering feature type [{0}] nor triggering feature element code [{1}].  At least one trigger must be specified."
	2255:  {G2Configuration},                       // EAS_ERR_FEATURE_CLASS_CODE_DOES_NOT_EXIST                                              "Feature class code [{0}] does not exist."
	2256:  {G2Configuration},                       // EAS_ERR_RELATIONSHIP_TYPE_CODE_DOES_NOT_EXIST                                          "Relationship type code [{0}] does not exist."
	2257:  {G2Configuration},                       // EAS_ERR_FELEM_CODE_NOT_IN_FEATURE                                                      "Feature element code [{0}] not included in feature[{1}]."
	2258:  {G2Configuration},                       // EAS_ERR_ER_FRAGMENT_DOES_NOT_EXIST                                                     "ER fragment code [{0}] does not exist."
	2259:  {G2Configuration},                       // EAS_ERR_ER_RULE_DOES_NOT_EXIST                                                         "ER rule code [{0}] does not exist."
	2260:  {G2Configuration},                       // EAS_ERR_ERFRAG_ID_ALREADY_EXISTS                                                       "ER fragment ID [{0}] already exists."
	2261:  {G2Configuration},                       // EAS_ERR_ERRULE_ID_ALREADY_EXISTS                                                       "ER rule ID [{0}] already exists."
	2262:  {G2Configuration},                       // EAS_ERR_ERFRAG_CODE_ALREADY_EXISTS                                                     "ER fragment code [{0}] already exists."
	2263:  {G2Configuration},                       // EAS_ERR_ERRULE_CODE_ALREADY_EXISTS                                                     "ER rule code [{0}] already exists."
	2264:  {G2Configuration},                       // EAS_ERR_ERFRAG_CODE_DOES_NOT_EXIST                                                     "ER fragment code [{0}] does not exist."
	2266:  {G2Configuration},                       // EAS_ERR_ERFRAG_CODE_MUST_BE_UNIQUE_IN_DEPENDENCY_LIST                                  "ER fragment code [{0}] must be unique in dependency list."
	2267:  {G2Configuration},                       // EAS_ERR_SECTION_NAME_ALREADY_EXISTS                                                    "Section name [{0}] already exists."
	2268:  {G2Configuration},                       // EAS_ERR_SECTION_NAME_DOES_NOT_EXIST                                                    "Section name [{0}] does not exist."
	2269:  {G2Configuration},                       // EAS_ERR_SECTION_FIELD_NAME_ALREADY_EXISTS                                              "Section field name [{0}] already exists."
	2270:  {G2Configuration},                       // EAS_ERR_SFUNC_ID_ALREADY_EXISTS                                                        "Feature standardization function ID [{0}] already exists."
	2271:  {G2Configuration},                       // EAS_ERR_SFUNC_CODE_ALREADY_EXISTS                                                      "Feature standardization function code [{0}] already exists."
	2272:  {G2Configuration},                       // EAS_ERR_EFUNC_ID_ALREADY_EXISTS                                                        "Feature expression function ID [{0}] already exists."
	2273:  {G2Configuration},                       // EAS_ERR_EFUNC_CODE_ALREADY_EXISTS                                                      "Feature expression function code [{0}] already exists."
	2274:  {G2Configuration},                       // EAS_ERR_CFUNC_ID_ALREADY_EXISTS                                                        "Feature comparison function ID [{0}] already exists."
	2275:  {G2Configuration},                       // EAS_ERR_CFUNC_CODE_ALREADY_EXISTS                                                      "Feature comparison function code [{0}] already exists."
	2276:  {G2Configuration},                       // EAS_ERR_DFUNC_ID_ALREADY_EXISTS                                                        "Feature distinct function ID [{0}] already exists."
	2277:  {G2Configuration},                       // EAS_ERR_DFUNC_CODE_ALREADY_EXISTS                                                      "Feature distinct function code [{0}] already exists."
	2278:  {G2Configuration},                       // EAS_ERR_COMPATIBILITY_VERSION_NOT_FOUND_IN_CONFIG                                      "Compatibility version not found in document."
	2279:  {G2Configuration},                       // EAS_ERR_CFRTN_ID_ALREADY_EXISTS                                                        "Feature comparison function return ID [{0}] already exists."
	2280:  {G2Configuration},                       // EAS_ERR_CFUNC_CODE_DOES_NOT_EXIST                                                      "Feature comparison function code [{0}] does not exist."
	2281:  {G2Configuration},                       // EAS_ERR_CFRTN_VALUE_ALREADY_EXISTS                                                     "Feature comparison function return value [{0}] already exists for comparison function [{1}] ftype [{2}]."
	2282:  {G2Configuration},                       // EAS_ERR_CFUNC_EXEC_ORDER_ALREADY_EXISTS                                                "Feature comparison function exec order value [{0}] already exists for comparison function [{1}] ftype [{2}]."
	2283:  {G2Configuration},                       // EAS_ERR_EFUNC_CODE_DOES_NOT_EXIST                                                      "Feature expression function code [{0}] does not exist."
	2285:  {G2Base},                                // EAS_ERR_INVALID_FORMAT_FOR_ENTITIES                                                    "Invalid format for ENTITIES."
	2286:  {G2Base},                                // EAS_ERR_NO_ENTITY_ID_FOUND_FOR_ENTITY                                                  "No entity ID found for entity."
	2287:  {G2Base},                                // EAS_ERR_NO_DATA_SOURCE_FOUND                                                           "No data source found."
	2288:  {G2Base},                                // EAS_ERR_NO_RECORD_ID_FOUND                                                             "No record ID found."
	2289:  {G2Configuration},                       // EAS_ERR_INVALID_FEATURE_CLASS_FOR_FEATURE_TYPE                                         "Invalid feature class [{0}] for feature type [{1}]."
	2290:  {G2Configuration},                       // EAS_ERR_FRAGMENT_IS_CONFIGURED_FOR_USE_IN_RULES                                        "Rule fragment [{0}] is configured for use in rules(s) [{1}]."
	2291:  {G2Configuration},                       // EAS_ERR_FRAGMENT_IS_CONFIGURED_FOR_USE_IN_FRAGMENT                                     "Rule fragment [{0}] is configured for use in fragments(s) [{1}]."
	2292:  {G2Base},                                // EAS_ERR_CANT_RETRIEVE_OBS_FEATURE_DATA_FOR_OBS_ENT                                     "Could not retrieve observed feature data for observed entity [{0}]."
	2293:  {G2Base},                                // EAS_ERR_NO_RECORDS_SPECIFIED                                                           "No records specified."
	2294:  {G2Base},                                // EAS_ERR_DATA_SOURCE_ID_DOES_NOT_EXIST                                                  "Data source ID [{0}] does not exist."
	2295:  {G2Base},                                // EAS_ERR_ETYPE_CODE_DOES_NOT_EXIST                                                      "Entity type code [{0}] does not exist."
	7209:  {G2Configuration},                       // EAS_ERR_DB_BAD_BACKEND_TYPE                                                            "Invalid [SQL] Backend Parameter. Valid values are SQL or HYBRID"
	7211:  {G2Configuration},                       // EAS_ERR_DB_BAD_CLUSTER_SIZE                                                            "Cluster [{0}] is configured with an invalid size. Size must be equal to 1."
	7212:  {G2Configuration},                       // EAS_ERR_DB_BAD_CLUSTER_NODE                                                            "Cluster [{0}] Node [{1}] is not configured."
	7216:  {G2Configuration},                       // EAS_ERR_DB_BAD_CLUSTER_DEFINITION                                                      "Cluster [{0}] is not properly configured"
	7217:  {G2Configuration},                       // EAS_ERR_DB_CONFLICTING_DEFAULT_SHARD_CONFIG                                            "Cannot specify both default backend database and default backend cluster"
	7218:  {G2Configuration},                       // EAS_ERR_DB_CLUSTER_DOES_NOT_EXIST                                                      "Cluster [{0}] does not exist"
	7220:  {G2Configuration},                       // EAS_ERR_NO_CONFIG_REGISTERED_IN_DATASTORE                                              "No engine configuration registered in datastore (see https://senzing.zendesk.com/hc/en-us/articles/360036587313)."
	7221:  {G2Configuration},                       // EAS_ERR_NO_CONFIG_REGISTERED_FOR_DATA_ID                                               "No engine configuration registered with data ID [{0}]."
	7222:  {G2Base},                                // EAS_ERR_FAILED_TO_SET_SYS_VAR_IN_DATASTORE                                             "Could not set system variable value in database for Group[{0}],Code[{1}],Value[{2}]."
	7223:  {G2Configuration},                       // EAS_ERR_INVALID_SCHEMA_VERSION_IN_DATASTORE                                            "Invalid version number for datastore schema [version '{0}']."
	7224:  {G2Configuration},                       // EAS_ERR_INVALID_SCHEMA_VERSION_IN_ENGINE                                               "Invalid version number for engine schema [version '{0}']."
	7226:  {G2Configuration},                       // EAS_ERR_INCOMPATIBLE_DATASTORE_SCHEMA_VERSION                                          "Incompatible datastore schema version: [Engine version '{0}'.  Datastore version '{1}' is installed, but must be between '{2}' and '{3}'.]"
	7227:  {G2Configuration},                       // EAS_ERR_CONFLICTING_SCHEMA_VERSIONS_IN_DATASTORE                                       "Conflicting version numbers for datastore schema [{0}]."
	7228:  {G2Configuration},                       // EAS_ERR_INVALID_SCHEMA_VERSION                                                         "Invalid schema version number [version '{0}']."
	7230:  {G2Configuration},                       // EAS_ERR_ENGINE_CONFIGURATION_FILE_NOT_FOUND                                            "Engine configuration file not found [{0}]."
	7232:  {G2Configuration},                       // EAS_ERR_ENGINE_CONFIGURATION_NOT_FOUND                                                 "No engine configuration found."
	7233:  {G2Configuration},                       // EAS_ERR_DATASTORE_ENCRYPTION_SIGNATURE_IS_INCOMPATIBLE                                 "Datastore encryption signature is not compatible."
	7234:  {G2Configuration},                       // EAS_ERR_FAILED_TO_GET_ENCRYPTION_SIGNATURE                                             "Failed to get encryption signature: '{0}'"
	7235:  {G2Configuration},                       // EAS_ERR_FTYPE_CONFIGURED_AS_REL_BUT_NO_RTYPE                                           "FTYPE_CODE[{0}] IS CONFIGURED AS A RELATIONSHIP FEATURE TYPE BUT RTYPE_ID IS NOT SET."
	7236:  {G2Configuration},                       // EAS_ERR_DUPLICATE_BEHAVIOR_OVERRIDE_KEY_IN_CFG_FBOVR                                   "Duplicate behavior override keys in CFG_FBOVR -- FTYPE_ID[{0}], UTYPE_CODE[{1}] referenced in CFG_FBOVR."
	7237:  {G2Configuration},                       // EAS_ERR_UNKNOWN_FTYPE_IN_TABLE                                                         "Unknown FTYPE_ID[{0}] referenced in {1}."
	7239:  {G2Configuration},                       // EAS_ERR_INVALID_GENERIC_THRESHOLD_CANDIDATE_CAP                                        "Invalid generic threshold {0} cap [{1}] for [GPLAN_ID[{2}], BEHAVIOR[{3}], FTYPE_ID[{4}]]."
	7240:  {G2Configuration},                       // EAS_ERR_INCORRECT_BEHAVIOR_REFERENCED                                                  "Incorrect BEHAVIOR[{0}] referenced in CFG_GENERIC_THRESHOLD for [GPLAN_ID[{1}], FTYPE_ID[{2}]].  FType configured for behavior [{3}]"
	7241:  {G2Configuration},                       // EAS_ERR_UNKNOWN_GPLAN_IN_TABLE                                                         "Unknown FTYPE_ID[{0}] referenced in {1}."
	7242:  {G2Configuration},                       // EAS_ERR_MULTIPLE_GENERIC_THRESHOLD_DEFINITIONS                                         "Multiple Generic Threshold definitions for [GPLAN_ID[{0}], BEHAVIOR[{1}], FTYPE_ID[{2}]]."
	7243:  {G2Configuration},                       // EAS_ERR_ER_FRAGMENT_HAS_UNDEFINED_DEPENDENT_FRAGMENTS                                  "ER Fragment [{0}] configured with undefined dependent fragments. Fragment [{1}] undefined."
	7244:  {G2Configuration},                       // EAS_ERR_ER_RULE_FRAGMENT_LACKS_REQUIRED_FRAGMENT                                       "ER Rule Fragment configuration lacks the required {0} fragment."
	7245:  {G2Configuration},                       // EAS_ERR_CURRENT_CONFIG_REGISTERED_DOES_NOT_MATCH_DATA_ID                               "Current configuration ID does not match specified data ID [{0}]."
	7246:  {G2Configuration},                       // EAS_ERR_INVALID_MAXIMUM_DATASTORE_SCHEMA_VERSION                                       "Invalid maximum datastore version number for engine schema [version '{0}']."
	7247:  {G2Configuration},                       // EAS_ERR_INVALID_MINIMUM_DATASTORE_SCHEMA_VERSION                                       "Invalid minimum datastore version number for engine schema [version '{0}']."
	7303:  {G2BadInput},                            // EAS_ERR_MANDATORY_SEGMENT_WITH_MISSING_REQUIREMENTS                                    "Mandatory segment with missing requirements:"
	7305:  {G2BadInput},                            // EAS_ERR_MISSING_JSON_ROOT_ELEMENT                                                      "No root element name in json TEMPLATE"
	7313:  {G2BadInput},                            // EAS_ERR_REQUIRED_ELEMENT_WITH_EMPTY_FIELD                                              "A non-empty value for [{0}] must be specified."
	7314:  {G2BadInput},                            // EAS_ERR_REQUIRED_ELEMENT_NOT_FOUND                                                     "A value for [{0}] must be specified."
	7317:  {G2Configuration},                       // EAS_ERR_FAILED_TO_OPEN_FILE                                                            "Failed to open file: {0}"
	7344:  {G2Configuration},                       // EAS_ERR_UNKNOWN_MAPPING_DIRECTIVE                                                      "Invalid mapping directive [{0}] for attribute [{1}]."
	7426:  {G2BadInput},                            // EAS_ERR_XLITERATOR_FAILED                                                              "Transliteration failed"
	7511:  {G2Base},                                // EAS_ERR_ABORT_ER_AND_RETRY                                                             "Detected change in candidate entity[{0}].  Restarting ER evaluation."
	8000:  {G2BadInput},                            // EAS_ERR_GNRNP                                                                          "GNR NameParser Failure"
	8410:  {G2Base},                                // EAS_ERR_UNINITIALIZED_AMBIGUOUS_FEATURE                                                "Cannot use uninitialized ambiguous feature."
	8501:  {G2Configuration},                       // EAS_ERR_SALT_DIGEST_ALGORITHM_NOT_AVAILABLE                                            "Failed to get {0} digest algorithm from ICC."
	8502:  {G2Base},                                // EAS_ERR_SALT_DIGEST_CONTEXT_CREATE_FAILED                                              "Failed to create a digest context."
	8503:  {G2Base},                                // EAS_ERR_SALT_DIGEST_CONTEXT_INIT_FAILED                                                "Failed {0} to initialise a digest context."
	8504:  {G2Base},                                // EAS_ERR_SALT_DIGEST_FAILED                                                             "Failed {0} to digest block {1}."
	8505:  {G2Base},                                // EAS_ERR_SALT_DIGEST_FINAL_FAILED                                                       "Failed {0} to complete digest."
	8508:  {G2Base},                                // EAS_ERR_SALT_DIGEST_UNKNOWN_EXCEPTION                                                  "Unrecognized exception thrown generating digest."
	8509:  {G2Base},                                // EAS_ERR_SALT_DIGEST_ALGORITHM_REQUIRED                                                 "Cannot generate a digest without a valid algorithm."
	8514:  {G2Base},                                // EAS_ERR_SALT_RANDOM_FAILED                                                             "Failed {0} to get random content"
	8516:  {G2Configuration},                       // EAS_ERR_SALT_MUST_BE_SIZE                                                              "A salt value must be {0} bytes long but the provided one is {1} bytes."
	8517:  {G2Configuration},                       // EAS_ERR_SALT_DOES_NOT_MATCH_CHECKSUM                                                   "The salt value does not match the recorded checksum."
	8520:  {G2Base},                                // EAS_ERR_SALT_G2SS_INIT_FAILED                                                          "Secure Store initialization failed."
	8521:  {G2Base},                                // EAS_ERR_SALT_G2SS_TOKEN_MUST_BE_INIT                                                   "Hashing with a named salt requires the Secure Store to be initialised."
	8522:  {G2Configuration},                       // EAS_ERR_SALT_G2SS_SOPIN_NOT_VALID                                                      "The Security Officer (SO) PIN is not correct."
	8524:  {G2Base},                                // EAS_ERR_SALT_G2SS_INIT_UNKNOWN_EXCEPTION                                               "Secure Store initialization failed with an unrecognised exception"
	8525:  {G2Configuration},                       // EAS_ERR_SALT_G2SS_REQUIRED_FOR_LOAD                                                    "Secure Store is required to load salt"
	8526:  {G2Configuration},                       // EAS_ERR_SALT_G2SS_REQUIRED_FOR_GENERATE                                                "Secure Store is required to generate salt"
	8527:  {G2Configuration},                       // EAS_ERR_SALT_G2SS_REQUIRED_FOR_IMPORT                                                  "Secure Store is required to import salt"
	8528:  {G2Configuration},                       // EAS_ERR_SALT_G2SS_REQUIRED_FOR_EXPORT                                                  "Secure Store is required to export salt"
	8529:  {G2Configuration},                       // EAS_ERR_SALT_G2SS_REQUIRED_FOR_DELETE                                                  "Secure Store is required to delete salt"
	8530:  {G2Base},                                // EAS_ERR_SALT_CANNOT_OVERWRITE                                                          "You cannot overwrite an existing salt called {0}"
	8536:  {G2Configuration},                       // EAS_ERR_SALT_G2SS_REQUIRED_FOR_LEGACY                                                  "Secure Store is required to add a legacy salt"
	8538:  {G2Configuration},                       // EAS_ERR_SALT_G2SS_REQUIRED_FOR_METHOD                                                  "Secure Store is required to change hashing method"
	8539:  {G2Base},                                // EAS_ERR_SALT_G2SS_ERROR_CHANGING_METHOD                                                "Secure Store error changing hashing method"
	8540:  {G2Configuration},                       // EAS_ERR_SALT_WRONG_SIZE                                                                "The object called {0} is not a salt"
	8541:  {G2Base},                                // EAS_ERR_SALT_BASE64_DECODE_ERROR                                                       "Base64 decoding error in salt {0} at character {1}"
	8542:  {G2Base},                                // EAS_ERR_SALT_UNINITIALISED                                                             "Must load a salt before using it."
	8543:  {G2Configuration},                       // EAS_ERR_SALT_NOT_FOUND                                                                 "There is no salt called {0} in the Secure Store."
	8544:  {G2Configuration},                       // EAS_ERR_SALT_PASSWORD_NOT_STRONG_ENOUGH                                                "The password must be stronger: {0}"
	8545:  {G2Configuration},                       // EAS_ERR_SALT_ADMIN_NAME_REQUIRED                                                       "Specify -name and the name to use for the salt"
	8556:  {G2Configuration},                       // EAS_ERR_SALT_ADMIN_METHOD_NOT_RECOGNISED                                               "Hashing method {0} not supported."
	8557:  {G2Configuration},                       // EAS_ERR_SALT_METHOD_DOES_NOT_MATCH                                                     "The hashing method in the configuration ({1}) does not match the method ({2}) of the salt {0}"
	8593:  {G2Base},                                // EAS_ERR_SALT_HMAC_CONTEXT_INIT_FAILED                                                  "Failed {0} to initialise an HMAC context."
	8594:  {G2Base},                                // EAS_ERR_SALT_HMAC_FAILED                                                               "Failed {0} to HMAC block {1}."
	8595:  {G2Base},                                // EAS_ERR_SALT_HMAC_FINAL_FAILED                                                         "Failed {0} to complete HMAC."
	8598:  {G2Base},                                // EAS_ERR_SALT_HMAC_UNKNOWN_EXCEPTION                                                    "Unrecognized exception thrown generating HMAC."
	8599:  {G2Configuration},                       // EAS_ERR_SALT_UNKNOWN_HASHING_METHOD                                                    "Unrecognized hashing method ({0}) requested."
	8601:  {G2Configuration},                       // EAS_ERR_HASHER_REQUIRES_SECURE_STORE                                                   "Using a named salt requires the Secure Store configured and running"
	8602:  {G2Configuration},                       // EAS_ERR_HASHER_CHECKSUM_DOES_NOT_MATCH                                                 "The hashing checksum configured ({1}) does not match the checksum ({2}) of the salt named {0}"
	8603:  {G2Base},                                // EAS_ERR_HASHER_UNABLE_TO_RECORD_SALT                                                   "Unable to record the configured salt"
	8604:  {G2Configuration},                       // EAS_ERR_HASHER_REQUIRES_FUNCTION                                                       "Using hashing requires a configured hashing function"
	8605:  {G2Configuration},                       // EAS_ERR_HASHER_EPHEMERAL_OR_NAMED_SALT                                                 "Specify either a named salt or an ephemeral one. Can not have both"
	8606:  {G2Configuration},                       // EAS_ERR_HASHER_SALT_REQUIRED                                                           "Hashing requires a salt to be configured."
	8607:  {G2Configuration},                       // EAS_ERR_HASHER_INVALID_ARGS                                                            "Invalid arguments to hashing function. Either a parameter wasn't provided or a buffer was too small: location={0}, dataPtr={1}, dataLength={2}, outputPtr={3}, outputLength={4}, output={5}"
	8608:  {G2Configuration},                       // EAS_ERR_NO_SALT_VALUE_CONFIGURED                                                       "No salt value is configured. A salt value must be configured if you wish to export the token library."
	8701:  {G2Configuration},                       // EAS_ERR_PARAMETER_NOT_READABLE                                                         "The parameter store does not support a read interface"
	8702:  {G2Configuration},                       // EAS_ERR_PARAMETER_NOT_WRITABLE                                                         "The parameter store does not support a write interface"
	9000:  {G2License, G2Unrecoverable},            // EAS_LIMIT_MAX_OBS_ENT                                                                  "LIMIT: Maximum number of records ingested: {0}"
	9107:  {G2Configuration},                       // EAS_ERR_CANT_GET_PARAMETER_FROM_THE_STORE                                              "Cannot get parameter [{0}] from parameter store"
	9110:  {G2Configuration},                       // EAS_ERR_INSUFFICIENT_CONFIG                                                            "Insufficient configuration for the {0} table!"
	9111:  {G2Configuration},                       // EAS_ERR_PARSE_FRAGMENT                                                                 "ERROR parsing FragmentID[{0}] FragmentName[{1}] : [{2}] is an invalid RuleID dependency"
	9112:  {G2Configuration},                       // EAS_ERR_FAILED_TO_OPEN_INI_FILE_FOR_WRITING                                            "Failed to open ini file for writing [{0}]"
	9113:  {G2Configuration},                       // EAS_ERR_FAILED_TO_OPEN_INI_FILE_FOR_READING                                            "Failed to open ini file for reading [{0}]"
	9115:  {G2BadInput},                            // EAS_ERR_INPUT_NOT_STANDARDIZED                                                         "Cannot process Observation that has not been standardized"
	9116:  {G2Configuration},                       // EAS_ERR_CONFIG_TABLE_NOT_FOUND                                                         "CONFIG information for {0} not found!"
	9117:  {G2Configuration},                       // EAS_ERR_CONFIG_TABLE_COLUMN_NOT_FOUND                                                  "CONFIG information for {0} not found in {1}!"
	9118:  {G2Configuration},                       // EAS_ERR_CONFIG_TABLE_COLUMN_INDEX_NOT_FOUND                                            "Invalid column index {0} queried from {1} container!"
	9119:  {G2Configuration},                       // EAS_ERR_CONFIG_TABLE_COLUMN_NAME_NOT_FOUND                                             "Invalid column name {0} queried from {1} container!"
	9120:  {G2Configuration},                       // EAS_ERR_CONFIG_TABLE_MALFORMED                                                         "CONFIG information for {0} is malformed!"
	9210:  {G2Configuration},                       // EAS_ERR_DIGEST_CONTEXT_INIT_FAILED                                                     "Unable to initialize Digest Context."
	9220:  {G2Configuration},                       // EAS_ERR_FTYPE_CANNOT_BE_HASHED                                                         "FType configured to be hashed, but cannot be scored.  FTYPE_ID[{0}] FTYPE_CODE[{1}]"
	9222:  {G2Configuration},                       // EAS_ERR_FTYPE_CONFIGURED_TO_BE_HASHED_MISSING_SALT                                     "A Feature Type is marked for hashing, but a valid salt value was not found.  FTYPE_ID[{0}] FTYPE_CODE[{1}]"
	9224:  {G2Configuration},                       // EAS_ERR_FTYPE_CONFIGURED_TO_BE_HASHED                                                  "FType configured to be hashed, but no hashable data found.  FTYPE_ID[{0}] FTYPE_CODE[{1}]"
	9228:  {G2Configuration},                       // EAS_ERR_UNEXPECTED_SALT_CHECKUM_LIST                                                   "The SALT checksum on the Observation does not match the EXPECTED SALT checksum: EXPECTED=[{0}] Observation=[{1}]"
	9240:  {G2Configuration},                       // EAS_ERR_CIPHER_CONTEXT_INIT_FAILED                                                     "Unable to initialize an ICC Context."
	9241:  {G2Configuration},                       // EAS_ERR_CIPHER_OP_FAILED                                                               "Unable to perform a required ICC operation."
	9250:  {G2Configuration},                       // EAS_ERR_G2SS_INVALID_LIB                                                               "Invalid ({1}) Secure Store plug-in library: {0}"
	9251:  {G2Configuration},                       // EAS_ERR_G2SS_INVALID_URL                                                               "Invalid Secure Store URL: {0}"
	9252:  {G2Configuration},                       // EAS_ERR_G2SS_INVALID_PIN                                                               "Invalid Secure Store credential specification: {0}"
	9253:  {G2Configuration},                       // EAS_ERR_G2SS_TOKEN_INIT_FAILED                                                         "Secure Store token initialization failed: {0}."
	9254:  {G2Configuration},                       // EAS_ERR_G2SS_TOKEN_UNINITIALISED                                                       "Cannot open a Secure Store session when the token is uninitialized."
	9255:  {G2Configuration},                       // EAS_ERR_G2SS_USER_PIN_UNINITIALISED                                                    "Secure Store credential is uninitialized."
	9256:  {G2Configuration},                       // EAS_ERR_G2SS_SESSION_OPEN                                                              "Cannot open a Secure Store session when one is already open."
	9257:  {G2Configuration},                       // EAS_ERR_G2SS_NO_SESSION                                                                "Cannot use Secure Store without a session."
	9258:  {G2Configuration},                       // EAS_ERR_G2SS_SESSION_OPEN_FAILED                                                       "Secure Store session could not be opened: {0}."
	9259:  {G2Configuration},                       // EAS_ERR_G2SS_ADMIN_LOGIN_FAILED                                                        "Secure Store admin login failed: {0}."
	9260:  {G2Configuration},                       // EAS_ERR_G2SS_USER_LOGIN_FAILED                                                         "Secure Store user login failed: {0}."
	9261:  {G2Configuration},                       // EAS_ERR_G2SS_PKCS11_ERROR                                                              "Secure Store function failed: {0}"
	9264:  {G2Configuration},                       // EAS_ERR_G2SS_LOGOUT_FAILED                                                             "Secure Store logout failed: {0}."
	9265:  {G2Configuration},                       // EAS_ERR_G2SS_NEED_RW_SESSION                                                           "Secure Store session must be read/write."
	9266:  {G2Configuration},                       // EAS_ERR_G2SS_UNABLE_TO_VERIFY_KEY                                                      "Secure Store key does not meet requirements."
	9267:  {G2Base},                                // EAS_ERR_G2SS_UNABLE_TO_CREATE_KEY                                                      "Secure Store key creation failed."
	9268:  {G2Base},                                // EAS_ERR_G2SS_UNABLE_TO_CHANGE_PIN                                                      "Secure Store password change failed: {0}."
	9269:  {G2Configuration},                       // EAS_ERR_G2SS_INVALID_OLD_CREDENTIAL                                                    "Secure Store old credential is invalid."
	9270:  {G2Configuration},                       // EAS_ERR_G2SS_INVALID_NEW_CREDENTIAL                                                    "Secure Store new credential is invalid."
	9271:  {G2Base},                                // EAS_ERR_G2SS_OUT_OF_MEMORY                                                             "Secure Store out of memory."
	9272:  {G2Base},                                // EAS_ERR_G2SS_FIND_INIT_FAILED                                                          "Secure Store object locating failed: {0}."
	9273:  {G2Base},                                // EAS_ERR_G2SS_FIND_FAILED                                                               "Secure Store object find failed: {0}."
	9274:  {G2Base},                                // EAS_ERR_G2SS_CRYPTO_SETUP_FAILED                                                       "Secure Store setup of encryption failed: {0}."
	9275:  {G2Base},                                // EAS_ERR_G2SS_ENCRYPT_START_FAILED                                                      "Secure Store unable to start encryption: {0}."
	9276:  {G2Base},                                // EAS_ERR_G2SS_ENCRYPT_SIZE_FAILED                                                       "Secure Store unable to get the size of encrypted data: {0}."
	9277:  {G2Base},                                // EAS_ERR_G2SS_ENCRYPT_FAILED                                                            "Secure Store encryption failed: {0}."
	9278:  {G2Base},                                // EAS_ERR_G2SS_DECRYPT_START_FAILED                                                      "Secure Store unable to start decryption: {0}."
	9279:  {G2Base},                                // EAS_ERR_G2SS_DECRYPT_FAILED                                                            "Secure Store decryption failed: {0}."
	9280:  {G2Base},                                // EAS_ERR_G2SS_OBJECT_SAVE_FAILED                                                        "Secure Store unable to save object: {0}."
	9281:  {G2Base},                                // EAS_ERR_G2SS_OBJECT_DELETE_FAILED                                                      "Secure Store unable to delete object: {0}."
	9282:  {G2Base},                                // EAS_ERR_G2SS_OBJECT_CHANGE_FAILED                                                      "Secure Store unable to modify object: {0}."
	9283:  {G2Base},                                // EAS_ERR_G2SS_UNINITIALISED                                                             "Secure Store has not been initialized"
	9284:  {G2Base},                                // EAS_ERR_G2SS_INVALID_SLOT_ID                                                           "Can not obtain info on specified slot. Possibly invalid slot ID specified in Secure Store URL: {0}"
	9285:  {G2Configuration},                       // EAS_ERR_G2SS_NO_TOKEN_IN_SLOT                                                          "No security token present in slot specified by Secure Store URL: slot ID = {0}"
	9286:  {G2Configuration},                       // EAS_ERR_G2SS_TOKEN_NOT_FOUND                                                           "Can not obtain info for security token. Possibly invalid token label and/or slot ID specified in Secure Store URL: {0}"
	9287:  {G2Base},                                // EAS_ERR_G2SS_TOKEN_IMPL_ERROR                                                          "An internal error occurred in the security token implementation library: Return Code = {0}"
	9288:  {G2Base},                                // EAS_ERR_G2SS_USER_PIN_PROMPT_FAILED                                                    "Was unable to prompt user for security token authentication."
	9289:  {G2Base},                                // EAS_ERR_G2SS_LABEL_CHANGED_SINCE_CONFIG_INIT                                           "Secure Store has been reconfigured since loading."
	9290:  {G2Base},                                // EAS_ERR_G2SS_OBJECT_NOT_FOUND                                                          "Secure Store does not have an object called {0}."
	9292:  {G2Configuration},                       // EAS_ERR_G2SS_NO_PASSWORD                                                               "No password supplied"
	9293:  {G2Configuration},                       // EAS_ERR_G2SS_NO_SEC_STORE_PREFIX                                                       "Secure Store expects a different format (starting with {0}) when a password is supplied"
	9295:  {G2Configuration},                       // EAS_ERR_G2SS_NO_DATA_OBJECTS                                                           "There are no Secure Store objects stored on the token"
	9296:  {G2Configuration},                       // EAS_ERR_G2SS_SEC_STORE_ARCHIVE_BAD                                                     "The exported archive appears to be corrupted around object {0}"
	9297:  {G2Configuration},                       // EAS_ERR_G2SS_FILE_NOT_FOUND                                                            "Secure Store failed to open {0}"
	9298:  {G2Configuration},                       // EAS_ERR_G2SS_FILE_CONTENTS_BAD                                                         "Secure Store contents of {0} not usable."
	9299:  {G2Base},                                // EAS_ERR_G2SS_CLASS_NOT_INIT                                                            "Secure Store internal error."
	9300:  {G2Configuration},                       // EAS_ERR_G2SS_PASSWORD_CHECK_ERROR                                                      "Secure Store internal error ({0}) checking password."
	9301:  {G2Configuration},                       // EAS_ERR_MISSING_SEQUENCE_ENTRY                                                         "Missing Sequence Entry[{0}] in the SYS_SEQUENCE table!"
	9305:  {G2Base},                                // EAS_ERR_SEQUENCE_RETRIES_FAILED                                                        "Retries failed to retrieve Sequence Entry[{0}] in the SYS_SEQUENCE table!  This may mean the CACHE_SIZE is too small."
	9308:  {G2Configuration},                       // EAS_ERR_MISSING_STATUS_ENTRY                                                           "Could not retrieve status entry[{0}] in the SYS_STATUS table!"
	9309:  {G2Configuration},                       // EAS_ERR_SEQUENCE_HAS_BEEN_RESET                                                        "Sequence entry[{0}] has been reset."
	9310:  {G2Configuration},                       // EAS_ERR_INVALID_STATUS_ENTRY_VALUE                                                     "Invalid value for status entry[{0}] in the SYS_STATUS table!"
	9311:  {G2Base},                                // EAS_ERR_COULD_NOT_RECORD_USAGE_TYPE                                                    "Could not record usage type [{0}] in the LIB_UTYPE table!"
	9406:  {G2Base},                                // EAS_ERR_G2SS_SESSION_MUST_NOT_BE_OPEN                                                  "Secure Store cannot fetch a value with sync if a session is already open."
	9408:  {G2Configuration},                       // EAS_ERR_G2SS_PASSWORD_INADEQUATE                                                       "The provided password is not strong enough: {0}"
	9409:  {G2Configuration},                       // EAS_ERR_G2SS_FUNCTION_LIST_NOT_SET                                                     "The security token interface is not yet set"
	9410:  {G2Base},                                // EAS_ERR_G2SS_PKCS_INIT_FAILED                                                          "Initializing token driver failed {0}"
	9411:  {G2Base},                                // EAS_ERR_G2SS_PKCS_FINAL_FAILED                                                         "Finalizing token driver failed {0}"
	9413:  {G2Configuration},                       // EAS_ERR_G2SS_INCORRECT_PASSWORD                                                        "The export file password appears to be incorrect."
	9414:  {G2BadInput},                            // EAS_ERR_STRING_IS_INVALID_UTF8                                                         "Invalid data string. Data must be in UTF-8."
	9500:  {G2Configuration},                       // EAS_ERR_TOKEN_LIBRARY_CHECKSUM_MISMATCH                                                "Cannot load token library. The checksum does not match the configuration of this node. Found: [{0}] Expected: [{1}]"
	9701:  {G2Base},                                // EAS_ERR_CANT_RETRIEVE_INDEX_FROM_MEMORY_ROW                                            "Cannot retrieve index[{0}] from memory row of key[{1}], out of range!"
	9703:  {G2Base},                                // EAS_ERR_MEMTBL_COL_INDEX_TOO_BIG                                                       "Current field in memory row is passed end of row"
	9802:  {G2Configuration},                       // EAS_ERR_INBOUND_OBS_CONFIG_CHECKSUM_MISMATCH                                           "Configuration checksum on inbound observation [{0}] does not match this nodes configuration checksum [{1}]. Cannot process."
	9803:  {G2Configuration},                       // EAS_ERR_CALC_CONFIGCHKSUM_AND_PARAMSTORE_CONFIGCHKSUM_DONT_MATCH                       "The calculated configuration checksum [{0}] does not match the CONFIGURATION_CHECKSUM value in the parameter store [{1}]."
	30011: {G2Base},                                // EAS_ERR_DELETE_WITH_RESOLVE_ONLY                                                       "Cannot delete an entity with type RESOLVE_ONLY"
	30101: {G2Base},                                // EAS_ERR_INVALID_SESSION_HANDLE                                                         "Invalid Session Handle [{0}]"
	30102: {G2Base},                                // EAS_ERR_INVALID_REPORT_HANDLE                                                          "Invalid Report Handle [{0}]"
	30103: {G2Base},                                // EAS_ERR_INVALID_EXPORT_HANDLE                                                          "Invalid Export Handle [{0}]"
	30110: {G2Base},                                // EAS_ERR_RESPONSE_MESSAGE_SIZE_LARGER_THAN_BUFFER_SIZE                                  "Response message size [{0}] is larger than buffer size [{1}]"
	30111: {G2Base},                                // EAS_ERR_RESPONSE_RESIZE_FUNCTION_IS_NOT_PROVIDED                                       "Resize function is not provided"
	30112: {G2Base},                                // EAS_ERR_RESPONSE_RESIZE_FUNCTION_GAVE_INVALID_RESULT                                   "Resize function returned an invalid result"
	30121: {G2BadInput},                            // EAS_ERR_JSON_PARSING_FAILURE                                                           "JSON Parsing Failure [code={0},offset={1}]"
	30122: {G2BadInput},                            // EAS_ERR_JSON_PARSING_FAILURE_MUST_BE_OBJECT_OR_ARRAY                                   "JSON Parsing Failure.  JSON must be object or array."
	30123: {G2BadInput},                            // EAS_ERR_JSON_PARSING_FAILURE_OBJECT_HAS_DUPLICATE_KEYS                                 "Json object has duplicate keys."
	30131: {G2BadInput},                            // EAS_ERR_UNKNOWN_COLUMN_REQUESTED_FOR_CSV_EXPORT                                        "Invalid column [{0}] requested for CSV export."

	// Deleted:

	// 70:    {G2x},                                // EAS_ERR_FAILED_TO_READ_FROM_OUTPUT_FILE                                                "Failed to read from the output file."
	// 71:    {G2x},                                // EAS_ERR_DIRECTORY_DOES_NOT_EXIST                                                       "{0} directory [{1}] does not exist.  Check your command-line arguments."
	// 72:    {G2x},                                // EAS_ERR_NO_CONFIG_UPGRADE_FILES                                                        "No configuration upgrade files found in directory [{0}]."
	// 73:    {G2x},                                // EAS_ERR_G2CONFIGTOOL_NOT_FOUND                                                         "G2ConfigTool.py not found."
	// 74:    {G2x},                                // EAS_ERR_ERROR_EXECUTING_CONFIG_SCRIPT                                                  "Error when executing configuration script. {0} {1}"
	// 75:    {G2x},                                // EAS_ERR_NO_DATABASE_UPGRADE_FILES                                                      "No database upgrade files found in directory [{0}]."
	// 2004:  {G2x},                                // EAS_WRN_FELEM_CODE_BUT_NO_VALUE                                                        "OBS_FELEM has FELEM_CODE but no FELEM_VALUE[{0}]"
	// 2008:  {G2x},                                // EAS_WRN_FEATURE_HAS_NO_FELEM                                                           "Cannot process feature with no elements[{0}]"
	// 2059:  {G2x},                                // EAS_ERR_INHERIT_MISSING_RES_ENT_ID                                                     "OBS_ENT INHERIT TAG IS MISSING REQUIRED RES_ENT_ID TAG"
	// 2060:  {G2x},                                // EAS_ERR_INHERIT_MULTIPLE                                                               "OBS_ENT HAS MULTIPLE INHERIT TAGS.  ONLY ONE IS ALLOWED"
	// 2115:  {G2x},                                // EAS_INF_UNRESOLVE_ABORTED_TIME_EXCEEDED                                                "Unresolve for entities aborted"
	// 9501:  {G2x},                                // EAS_TOKEN_LIBRARY_ALREADY_HASHED                                                       "Cannot hash token library. The Token Library contains previous hashed data"

	// Duplicates:

	// 2091:  {G2x},                                // EAS_INF_OBS_ENT_ID_NOT_PROVIDED                                                        "Requested resolution without providing an OBS_ENT_ID"
	// 7245:  {G2x},                                // EAS_ERR_INVALID_MINIMUM_DATASTORE_SCHEMA_VERSION                                       "Invalid minimum datastore version number for engine schema [version '{0}']."

	// The 999nn series is for testing.

	99900: {G2Base},
	99910: {G2BadInput},
	99911: {G2NotFound, G2BadInput},
	99912: {G2UnknownDatasource, G2BadInput},
	99920: {G2Configuration},
	99930: {G2Retryable},
	99931: {G2DatabaseConnectionLost, G2Retryable},
	99932: {G2RetryTimeoutExceeded, G2Retryable},
	99940: {G2Unrecoverable},
	99941: {G2Database, G2Unrecoverable},
	99942: {G2License, G2Unrecoverable},
	99943: {G2NotInitialized, G2Unrecoverable},
	99944: {G2Unhandled, G2Unrecoverable},
}

// A list of all G2ErrorTypeIds.
var AllG2ErrorTypes = []G2ErrorTypeIds{
	G2BadInput,
	G2Base,
	G2Configuration,
	G2Database,
	G2DatabaseConnectionLost,
	G2License,
	G2NotFound,
	G2NotInitialized,
	G2Retryable,
	G2RetryTimeoutExceeded,
	G2Unhandled,
	G2UnknownDatasource,
	G2Unrecoverable,
}
