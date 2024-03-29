package szconfig

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Log message prefix.
const Prefix = "szconfig."

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for szconfig implementations.
var IdMessages = map[int]string{
	1:    "Enter " + Prefix + "AddDataSource(%v, %s).",
	2:    "Exit  " + Prefix + "AddDataSource(%v, %s) returned (%s, %v).",
	3:    "Enter " + Prefix + "ClearLastException().",
	4:    "Exit  " + Prefix + "ClearLastException() returned (%v).",
	5:    "Enter " + Prefix + "Close(%v).",
	6:    "Exit  " + Prefix + "Close(%v) returned (%v).",
	7:    "Enter " + Prefix + "Create().",
	8:    "Exit  " + Prefix + "Create() returned (%v, %v).",
	9:    "Enter " + Prefix + "DeleteDataSource(%v, %s).",
	10:   "Exit  " + Prefix + "DeleteDataSource(%v, %s) returned (%v).",
	11:   "Enter " + Prefix + "Destroy().",
	12:   "Exit  " + Prefix + "Destroy() returned (%v).",
	13:   "Enter " + Prefix + "GetLastException().",
	14:   "Exit  " + Prefix + "GetLastException() returned (%s, %v).",
	15:   "Enter " + Prefix + "GetLastExceptionCode().",
	16:   "Exit  " + Prefix + "GetLastExceptionCode() returned (%d, %v).",
	17:   "Enter " + Prefix + "Initialize(%s, %s, %d).",
	18:   "Exit  " + Prefix + "Initialize(%s, %s, %d) returned (%v).",
	19:   "Enter " + Prefix + "GetDataSources(%v).",
	20:   "Exit  " + Prefix + "GetDataSources(%v) returned (%s, %v).",
	21:   "Enter " + Prefix + "Load(%s).",
	22:   "Exit  " + Prefix + "Load(%s) returned (%v, %v).",
	23:   "Enter " + Prefix + "GetJsonString(%v).",
	24:   "Exit  " + Prefix + "GetJsonString(%v) returned (%s, %v).",
	25:   "Enter " + Prefix + "SetLogLevel(%v).",
	26:   "Exit  " + Prefix + "SetLogLevel(%v) returned (%v).",
	27:   "Enter " + Prefix + "RegisterObserver(%s).",
	28:   "Exit  " + Prefix + "RegisterObserver(%s) returned (%v).",
	29:   "Enter " + Prefix + "UnregisterObserver(%s).",
	30:   "Exit  " + Prefix + "UnregisterObserver(%s) returned (%v).",
	31:   "Enter " + Prefix + "GetSdkId().",
	32:   "Exit  " + Prefix + "GetSdkId() returned (%s).",
	4001: Prefix + "G2Config_addDataSource(%v, %s) failed. Return code: %d",
	4002: Prefix + "G2Config_close(%v) failed. Return code: %d",
	4003: Prefix + "G2Config_create() failed. Return code: %d",
	4004: Prefix + "G2Config_deleteDataSource(%v, %s) failed. Return code: %d",
	4005: Prefix + "G2Config_getLastException() failed. Return code: %d",
	4006: Prefix + "G2Config_destroy() failed. Return code: %d",
	4007: Prefix + "G2Config_init(%s, %s, %d) failed. Return code: %d",
	4008: Prefix + "G2Config_listDataSources() failed. Return code: %d",
	4009: Prefix + "G2Config_load(%s) failed. Return code: %d",
	4010: Prefix + "G2Config_save(%v) failed. Return code: %d",
	5901: "During test setup, call to messagelogger.NewSenzingApiLogger() failed.",
	5902: "During test setup, call to szengineconfigurationjson.BuildSimpleSystemConfigurationJsonViaMap() failed.",
	5903: "During test setup, call to szengine.Initialize() failed.",
	5904: "During test setup, call to szdiagnostic.PurgeRepository() failed.",
	5905: "During test setup, call to szengine.Destroy() failed.",
	5906: "During test setup, call to szconfig.Initialize() failed.",
	5907: "During test setup, call to szconfig.Create() failed.",
	5908: "During test setup, call to szconfig.AddDataSource() failed.",
	5909: "During test setup, call to szconfig.GetJsonString() failed.",
	5910: "During test setup, call to szconfig.Close() failed.",
	5911: "During test setup, call to szconfig.Destroy() failed.",
	5912: "During test setup, call to szconfigmgr.Initialize() failed.",
	5913: "During test setup, call to szconfigmgr.AddConfig() failed.",
	5914: "During test setup, call to szconfigmgr.SetDefaultConfigId() failed.",
	5915: "During test setup, call to szconfigmgr.Destroy() failed.",
	5916: "During test setup, call to szengine.Initialize() failed.",
	5917: "During test setup, call to szengine.AddRecord() failed.",
	5918: "During test setup, call to szengine.Destroy() failed.",
	5920: "During test setup, call to setupSenzingConfig() failed.",
	5921: "During test setup, call to setupPurgeRepository() failed.",
	5922: "During test setup, call to setupAddRecords() failed.",
	5931: "During test setup, call to szengine.Initialize() failed.",
	5932: "During test setup, call to szdiagnostic.PurgeRepository() failed.",
	5933: "During test setup, call to szengine.Destroy() failed.",
	8001: Prefix + "AddDataSource",
	8002: Prefix + "Close",
	8003: Prefix + "Create",
	8004: Prefix + "DeleteDataSource",
	8005: Prefix + "Destroy",
	8006: Prefix + "Initialize",
	8007: Prefix + "GetDataSources",
	8008: Prefix + "Load",
	8009: Prefix + "GetJsonString",
	8010: Prefix + "GetSdkId",
	8011: Prefix + "RegisterObserver",
	8012: Prefix + "SetLogLevel",
	8013: Prefix + "UnregisterObserver",
}

// Status strings for specific szconfig messages.
var IdStatuses = map[int]string{}