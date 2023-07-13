package g2configmgr

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Log message prefix.
const Prefix = "g2configmgr."

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for g2configmgr implementations.
var IdMessages = map[int]string{
	1:    "Enter " + Prefix + "AddConfig(%s, %s).",
	2:    "Exit  " + Prefix + "AddConfig(%s, %s) returned (%d, %v).",
	3:    "Enter " + Prefix + "ClearLastException().",
	4:    "Exit  " + Prefix + "ClearLastException() returned (%v).",
	5:    "Enter " + Prefix + "Destroy().",
	6:    "Exit  " + Prefix + "Destroy() returned (%v).",
	7:    "Enter " + Prefix + "GetConfig(%d).",
	8:    "Exit  " + Prefix + "GetConfig(%d) returned (%s, %v).",
	9:    "Enter " + Prefix + "GetConfigList().",
	10:   "Exit  " + Prefix + "GetConfigList() returned (%s, %v).",
	11:   "Enter " + Prefix + "GetDefaultConfigID().",
	12:   "Exit  " + Prefix + "GetDefaultConfigID() returned (%d, %v).",
	13:   "Enter " + Prefix + "GetLastException().",
	14:   "Exit  " + Prefix + "GetLastException() returned (%s, %v).",
	15:   "Enter " + Prefix + "GetLastExceptionCode().",
	16:   "Exit  " + Prefix + "GetLastExceptionCode() returned (%d, %v).",
	17:   "Enter " + Prefix + "Init(%s, %s, %d).",
	18:   "Exit  " + Prefix + "Init(%s, %s, %d) returned (%v).",
	19:   "Enter " + Prefix + "ReplaceDefaultConfigID(%d, %d).",
	20:   "Exit  " + Prefix + "ReplaceDefaultConfigID(%d, %d) returned (%v).",
	21:   "Enter " + Prefix + "SetDefaultConfigID(%d).",
	22:   "Exit  " + Prefix + "SetDefaultConfigID(%d) returned (%v).",
	23:   "Enter " + Prefix + "SetLogLevel(%v).",
	24:   "Exit  " + Prefix + "SetLogLevel(%v) returned (%v).",
	25:   "Enter " + Prefix + "RegisterObserver(%s).",
	26:   "Exit  " + Prefix + "RegisterObserver(%s) returned (%v).",
	27:   "Enter " + Prefix + "UnregisterObserver(%s).",
	28:   "Exit  " + Prefix + "UnregisterObserver(%s) returned (%v).",
	29:   "Enter " + Prefix + "GetSdkId().",
	30:   "Exit  " + Prefix + "GetSdkId() returned (%s).",
	4001: Prefix + "G2ConfigMgr_addConfig(%s, %s) failed. Return code: %d",
	4002: Prefix + "G2ConfigMgr_destroy() failed. Return code: %d",
	4003: Prefix + "G2ConfigMgr_getConfig(%d) failed. Return code: %d",
	4004: Prefix + "G2ConfigMgr_getConfigList() failed. Return code: %d",
	4005: Prefix + "G2ConfigMgr_getDefaultConfigID() failed. Return code: %d",
	4006: Prefix + "G2ConfigMgr_getLastException() failed. Return code: %d",
	4007: Prefix + "G2ConfigMgr_init(%s, %s, %d) failed. Return code: %d",
	4008: Prefix + "G2ConfigMgr_replaceDefaultConfigID(%d, %d) failed. Return code: %d",
	4009: Prefix + "G2ConfigMgr_setDefaultConfigID(%d) failed. Return code: %d",
	5901: "During test setup, call to messagelogger.NewSenzingApiLogger() failed.",
	5902: "During test setup, call to g2eg2engineconfigurationjson.BuildSimpleSystemConfigurationJsonViaMap() failed.",
	5903: "During test setup, call to g2engine.Init() failed.",
	5904: "During test setup, call to g2engine.PurgeRepository() failed.",
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
	5932: "During test setup, call to g2engine.PurgeRepository() failed.",
	5933: "During test setup, call to g2engine.Destroy() failed.",
	8001: Prefix + "AddConfig",
	8002: Prefix + "Destroy",
	8003: Prefix + "GetConfig",
	8004: Prefix + "GetConfigList",
	8005: Prefix + "GetDefaultConfigID",
	8006: Prefix + "Init",
	8007: Prefix + "ReplaceDefaultConfigID",
	8008: Prefix + "SetDefaultConfigID",
	8009: Prefix + "GetSdkId",
	8010: Prefix + "RegisterObserver",
	8011: Prefix + "SetLogLevel",
	8012: Prefix + "UnregisterObserver",
}

// Status strings for specific g2configmgr messages.
var IdStatuses = map[int]string{}
