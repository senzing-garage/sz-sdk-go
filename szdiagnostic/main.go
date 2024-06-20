package szdiagnostic

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Log message prefix.
const Prefix = "szdiagnostic."

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for szdiagnostic implementations.
var IDMessages = map[int]string{
	1:    "Enter " + Prefix + "CheckDatastorePerformance(%d).",
	2:    "Exit  " + Prefix + "CheckDatastorePerformance(%d) returned (%s, %v).",
	3:    "Enter " + Prefix + "ClearLastException().",
	4:    "Exit  " + Prefix + "ClearLastException() returned (%v).",
	5:    "Enter " + Prefix + "Destroy().",
	6:    "Exit  " + Prefix + "Destroy() returned (%v).",
	7:    "Enter " + Prefix + "GetDatastoreInfo().",
	8:    "Exit  " + Prefix + "GetDatastoreInfo() returned (%s, %v).",
	9:    "Enter " + Prefix + "GetFeature(%d).",
	10:   "Exit  " + Prefix + "GetFeature(%d) returned (%s, %v).",
	11:   "Enter " + Prefix + "GetLastException().",
	12:   "Exit  " + Prefix + "GetLastException() returned (%s, %v).",
	13:   "Enter " + Prefix + "GetLastExceptionCode().",
	14:   "Exit  " + Prefix + "GetLastExceptionCode() returned (%d, %v).",
	15:   "Enter " + Prefix + "Initialize(%s, %s, %d).",
	16:   "Exit  " + Prefix + "Initialize(%s, %s, %d) returned (%v).",
	17:   "Enter " + Prefix + "PurgeRepository().",
	18:   "Exit  " + Prefix + "PurgeRepository() returned (%v).",
	19:   "Enter " + Prefix + "Reinitialize(%d).",
	20:   "Exit  " + Prefix + "Reinitialize(%d) returned (%v).",
	703:  "Enter " + Prefix + "RegisterObserver(%s).",
	704:  "Exit  " + Prefix + "RegisterObserver(%s) returned (%v).",
	705:  "Enter " + Prefix + "SetLogLevel(%v).",
	706:  "Exit  " + Prefix + "SetLogLevel(%v) returned (%v).",
	707:  "Enter " + Prefix + "UnregisterObserver(%s).",
	708:  "Exit  " + Prefix + "UnregisterObserver(%s) returned (%v).",
	4001: Prefix + "G2Diagnostic_checkDatastorePerformance(%d) failed. Return code: %d",
	4002: Prefix + "G2Diagnostic_destroy() failed.  Return code: %d",
	4003: Prefix + "G2Diagnostic_getDatastoreInfo() failed. Return code: %d",
	4004: Prefix + "G2Diagnostic_getFeature(%d) failed. Return code: %d",
	4005: Prefix + "G2Diagnostic_init(%s, %s, %d) failed. Return code: %d",
	4006: Prefix + "G2Diagnostic_initWithConfigID(%s, %s, %d, %d) failed. Return code: %d",
	4007: Prefix + "G2Diagnostic_purgeRepository() failed. Return Code: %d",
	4008: Prefix + "G2Diagnostic_reinit(%d) failed. Return Code: %d",
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
	5912: "During test setup, call to szconfigmanager.Initialize() failed.",
	5913: "During test setup, call to szconfigmanager.AddConfig() failed.",
	5914: "During test setup, call to szconfigmanager.SetDefaultConfigID() failed.",
	5915: "During test setup, call to szconfigmanager.Destroy() failed.",
	5916: "During test setup, call to szengine.Initialize() failed.",
	5917: "During test setup, call to szengine.AddRecord() failed.",
	5918: "During test setup, call to szengine.Destroy() failed.",
	5920: "During test setup, call to setupSenzingConfig() failed.",
	5921: "During test setup, call to setupPurgeRepository() failed.",
	5922: "During test setup, call to setupAddRecords() failed.",
	5931: "During test setup, call to szengine.Initialize() failed.",
	5932: "During test setup, call to szdiagnostic.PurgeRepository() failed.",
	5933: "During test setup, call to szengine.Destroy() failed.",
	8001: Prefix + "CheckDatastorePerformance",
	8002: Prefix + "Destroy",
	8003: Prefix + "GetDatastoreInfo",
	8004: Prefix + "GetFeature",
	8005: Prefix + "Initialize",
	8007: Prefix + "PurgeRepository",
	8008: Prefix + "Reinitialize",
	8702: Prefix + "RegisterObserver",
	8703: Prefix + "SetLogLevel",
	8704: Prefix + "UnregisterObserver",
}

// Status strings for specific szdiagnostic messages.
var IDStatuses = map[int]string{}
