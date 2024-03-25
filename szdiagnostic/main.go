package szdiagnostic

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Log message prefix.
const Prefix = "g2diagnostic."

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for g2diagnostic implementations.
var IdMessages = map[int]string{
	1:    "Enter " + Prefix + "CheckDatabasePerformance(%d).",
	2:    "Exit  " + Prefix + "CheckDatabasePerformance(%d) returned (%s, %v).",
	3:    "Enter " + Prefix + "ClearLastException().",
	4:    "Exit  " + Prefix + "ClearLastException() returned (%v).",
	7:    "Enter " + Prefix + "Destroy().",
	8:    "Exit  " + Prefix + "Destroy() returned (%v).",
	31:   "Enter " + Prefix + "GetLastException().",
	32:   "Exit  " + Prefix + "GetLastException() returned (%s, %v).",
	33:   "Enter " + Prefix + "GetLastExceptionCode().",
	34:   "Exit  " + Prefix + "GetLastExceptionCode() returned (%d, %v).",
	47:   "Enter " + Prefix + "Initialize(%s, %s, %d).",
	48:   "Exit  " + Prefix + "Initialize(%s, %s, %d) returned (%v).",
	51:   "Enter " + Prefix + "Reinitialize(%d).",
	52:   "Exit  " + Prefix + "Reinitialize(%d) returned (%v).",
	53:   "Enter " + Prefix + "SetLogLevel(%v).",
	54:   "Exit  " + Prefix + "SetLogLevel(%v) returned (%v).",
	55:   "Enter " + Prefix + "RegisterObserver(%s).",
	56:   "Exit  " + Prefix + "RegisterObserver(%s) returned (%v).",
	57:   "Enter " + Prefix + "UnregisterObserver(%s).",
	58:   "Exit  " + Prefix + "UnregisterObserver(%s) returned (%v).",
	59:   "Enter " + Prefix + "GetSdkId().",
	60:   "Exit  " + Prefix + "GetSdkId() returned (%s).",
	61:   "Enter " + Prefix + "PurgeRepository().",
	62:   "Exit  " + Prefix + "PurgeRepository() returned (%v).",
	4001: Prefix + "G2Diagnostic_checkDBPerf(%d) failed. Return code: %d",
	4003: Prefix + "G2Diagnostic_destroy() failed.  Return code: %d",
	4014: Prefix + "G2Diagnostic_getLastException() failed. Return code: %d",
	4018: Prefix + "G2Diagnostic_init(%s, %s, %d) failed. Return code: %d",
	4019: Prefix + "G2Diagnostic_initWithConfigID(%s, %s, %d, %d) failed. Return code: %d",
	4020: Prefix + "G2Diagnostic_reinit(%d) failed. Return Code: %d",
	4021: Prefix + "G2Diagnostic_purgeRepository(%d) failed. Return Code: %d",
	5901: "During test setup, call to messagelogger.NewSenzingApiLogger() failed.",
	5902: "During test setup, call to g2engineconfigurationjson.BuildSimpleSystemConfigurationJsonViaMap() failed.",
	5903: "During test setup, call to g2engine.Initialize() failed.",
	5904: "During test setup, call to g2diagnostic.PurgeRepository() failed.",
	5905: "During test setup, call to g2engine.Destroy() failed.",
	5906: "During test setup, call to g2config.Initialize() failed.",
	5907: "During test setup, call to g2config.Create() failed.",
	5908: "During test setup, call to g2config.AddDataSource() failed.",
	5909: "During test setup, call to g2config.GetJsonString() failed.",
	5910: "During test setup, call to g2config.Close() failed.",
	5911: "During test setup, call to g2config.Destroy() failed.",
	5912: "During test setup, call to g2configmgr.Initialize() failed.",
	5913: "During test setup, call to g2configmgr.AddConfig() failed.",
	5914: "During test setup, call to g2configmgr.SetDefaultConfigId() failed.",
	5915: "During test setup, call to g2configmgr.Destroy() failed.",
	5916: "During test setup, call to g2engine.Initialize() failed.",
	5917: "During test setup, call to g2engine.AddRecord() failed.",
	5918: "During test setup, call to g2engine.Destroy() failed.",
	5920: "During test setup, call to setupSenzingConfig() failed.",
	5921: "During test setup, call to setupPurgeRepository() failed.",
	5922: "During test setup, call to setupAddRecords() failed.",
	5931: "During test setup, call to g2engine.Initialize() failed.",
	5932: "During test setup, call to g2diagnostic.PurgeRepository() failed.",
	5933: "During test setup, call to g2engine.Destroy() failed.",
	8001: Prefix + "CheckDatabasePerformance",
	8003: Prefix + "Destroy",
	8021: Prefix + "Initialize",
	8023: Prefix + "Reinitialize",
	8024: Prefix + "GetSdkId",
	8025: Prefix + "RegisterObserver",
	8026: Prefix + "SetLogLevel",
	8027: Prefix + "UnregisterObserver",
	8028: Prefix + "PurgeRepository",
}

// Status strings for specific g2diagnostic messages.
var IdStatuses = map[int]string{}
