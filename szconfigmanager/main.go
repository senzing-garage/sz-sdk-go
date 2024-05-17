package szconfigmanager

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Log message prefix.
const Prefix = "szconfigmanager."

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for szconfigmanager implementations.
var IDMessages = map[int]string{
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
	11:   "Enter " + Prefix + "GetDefaultConfigId().",
	12:   "Exit  " + Prefix + "GetDefaultConfigId() returned (%d, %v).",
	13:   "Enter " + Prefix + "GetLastException().",
	14:   "Exit  " + Prefix + "GetLastException() returned (%s, %v).",
	15:   "Enter " + Prefix + "GetLastExceptionCode().",
	16:   "Exit  " + Prefix + "GetLastExceptionCode() returned (%d, %v).",
	17:   "Enter " + Prefix + "Initialize(%s, %s, %d).",
	18:   "Exit  " + Prefix + "Initialize(%s, %s, %d) returned (%v).",
	19:   "Enter " + Prefix + "ReplaceDefaultConfigId(%d, %d).",
	20:   "Exit  " + Prefix + "ReplaceDefaultConfigId(%d, %d) returned (%v).",
	21:   "Enter " + Prefix + "SetDefaultConfigId(%d).",
	22:   "Exit  " + Prefix + "SetDefaultConfigId(%d) returned (%v).",
	701:  "Enter " + Prefix + "GetSdkId().",
	702:  "Exit  " + Prefix + "GetSdkId() returned (%s).",
	703:  "Enter " + Prefix + "RegisterObserver(%s).",
	704:  "Exit  " + Prefix + "RegisterObserver(%s) returned (%v).",
	705:  "Enter " + Prefix + "SetLogLevel(%v).",
	706:  "Exit  " + Prefix + "SetLogLevel(%v) returned (%v).",
	707:  "Enter " + Prefix + "UnregisterObserver(%s).",
	708:  "Exit  " + Prefix + "UnregisterObserver(%s) returned (%v).",
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
	5914: "During test setup, call to szconfigmanager.SetDefaultConfigId() failed.",
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
	8001: Prefix + "AddConfig",
	8002: Prefix + "Destroy",
	8003: Prefix + "GetConfig",
	8004: Prefix + "GetConfigList",
	8005: Prefix + "GetDefaultConfigId",
	8006: Prefix + "Initialize",
	8007: Prefix + "ReplaceDefaultConfigId",
	8008: Prefix + "SetDefaultConfigId",
	8701: Prefix + "GetSdkId",
	8702: Prefix + "RegisterObserver",
	8703: Prefix + "SetLogLevel",
	8704: Prefix + "UnregisterObserver",
}

// Status strings for specific szconfigmanager messages.
var IDStatuses = map[int]string{}
