package g2config

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Log message prefix.
const Prefix = "g2config."

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for g2config implementations.
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
	17:   "Enter " + Prefix + "Init(%s, %s, %d).",
	18:   "Exit  " + Prefix + "Init(%s, %s, %d) returned (%v).",
	19:   "Enter " + Prefix + "ListDataSources(%v).",
	20:   "Exit  " + Prefix + "ListDataSources(%v) returned (%s, %v).",
	21:   "Enter " + Prefix + "Load(%s).",
	22:   "Exit  " + Prefix + "Load(%s) returned (%v, %v).",
	23:   "Enter " + Prefix + "Save(%v).",
	24:   "Exit  " + Prefix + "Save(%v) returned (%s, %v).",
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
	5902: "During test setup, call to g2eg2engineconfigurationjson.BuildSimpleSystemConfigurationJsonViaMap() failed.",
	5903: "During test setup, call to g2engine.Init() failed.",
	5904: "During test setup, call to g2diagnostic.PurgeRepository() failed.",
	5905: "During test setup, call to g2engine.Destroy() failed.",
	5906: "During test setup, call to g2config.Init() failed.",
	5907: "During test setup, call to g2config.Create() failed.",
	5908: "During test setup, call to g2config.AddDataSource() failed.",
	5909: "During test setup, call to g2config.Save() failed.",
	5910: "During test setup, call to g2config.Close() failed.",
	5911: "During test setup, call to g2config.Destroy() failed.",
	5912: "During test setup, call to g2configmgr.Init() failed.",
	5913: "During test setup, call to g2configmgr.AddConfig() failed.",
	5914: "During test setup, call to g2configmgr.SetDefaultConfigID() failed.",
	5915: "During test setup, call to g2configmgr.Destroy() failed.",
	5916: "During test setup, call to g2engine.Init() failed.",
	5917: "During test setup, call to g2engine.AddRecord() failed.",
	5918: "During test setup, call to g2engine.Destroy() failed.",
	5920: "During test setup, call to setupSenzingConfig() failed.",
	5921: "During test setup, call to setupPurgeRepository() failed.",
	5922: "During test setup, call to setupAddRecords() failed.",
	5931: "During test setup, call to g2engine.Init() failed.",
	5932: "During test setup, call to g2diagnostic.PurgeRepository() failed.",
	5933: "During test setup, call to g2engine.Destroy() failed.",
	8001: Prefix + "AddDataSource",
	8002: Prefix + "Close",
	8003: Prefix + "Create",
	8004: Prefix + "DeleteDataSource",
	8005: Prefix + "Destroy",
	8006: Prefix + "Init",
	8007: Prefix + "ListDataSources",
	8008: Prefix + "Load",
	8009: Prefix + "Save",
	8010: Prefix + "GetSdkId",
	8011: Prefix + "RegisterObserver",
	8012: Prefix + "SetLogLevel",
	8013: Prefix + "UnregisterObserver",
}

// Status strings for specific g2config messages.
var IdStatuses = map[int]string{}
