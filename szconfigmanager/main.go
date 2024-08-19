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
	9:    "Enter " + Prefix + "GetConfigs().",
	10:   "Exit  " + Prefix + "GetConfigs() returned (%s, %v).",
	11:   "Enter " + Prefix + "GetDefaultConfigID().",
	12:   "Exit  " + Prefix + "GetDefaultConfigID() returned (%d, %v).",
	13:   "Enter " + Prefix + "GetLastException().",
	14:   "Exit  " + Prefix + "GetLastException() returned (%s, %v).",
	15:   "Enter " + Prefix + "GetLastExceptionCode().",
	16:   "Exit  " + Prefix + "GetLastExceptionCode() returned (%d, %v).",
	17:   "Enter " + Prefix + "Initialize(%s, %s, %d).",
	18:   "Exit  " + Prefix + "Initialize(%s, %s, %d) returned (%v).",
	19:   "Enter " + Prefix + "ReplaceDefaultConfigID(%d, %d).",
	20:   "Exit  " + Prefix + "ReplaceDefaultConfigID(%d, %d) returned (%v).",
	21:   "Enter " + Prefix + "SetDefaultConfigID(%d).",
	22:   "Exit  " + Prefix + "SetDefaultConfigID(%d) returned (%v).",
	703:  "Enter " + Prefix + "RegisterObserver(%s).",
	704:  "Exit  " + Prefix + "RegisterObserver(%s) returned (%v).",
	705:  "Enter " + Prefix + "SetLogLevel(%s).",
	706:  "Exit  " + Prefix + "SetLogLevel(%s) returned (%v).",
	707:  "Enter " + Prefix + "UnregisterObserver(%s).",
	708:  "Exit  " + Prefix + "UnregisterObserver(%s) returned (%v).",
	4001: Prefix + "SzConfigMgr_addConfig(%s, %s) failed. Return code: %d",
	4002: Prefix + "SzConfigMgr_destroy() failed. Return code: %d",
	4003: Prefix + "SzConfigMgr_getConfig(%d) failed. Return code: %d",
	4004: Prefix + "SzConfigMgr_getConfigList() failed. Return code: %d",
	4005: Prefix + "SzConfigMgr_getDefaultConfigID() failed. Return code: %d",
	4006: Prefix + "SzConfigMgr_init(%s, %s, %d) failed. Return code: %d",
	4007: Prefix + "SzConfigMgr_replaceDefaultConfigID(%d, %d) failed. Return code: %d",
	4008: Prefix + "SzConfigMgr_setDefaultConfigID(%d) failed. Return code: %d",
	8001: Prefix + "AddConfig",
	8002: Prefix + "Destroy",
	8003: Prefix + "GetConfig",
	8004: Prefix + "GetConfigs",
	8005: Prefix + "GetDefaultConfigID",
	8006: Prefix + "Initialize",
	8007: Prefix + "ReplaceDefaultConfigID",
	8008: Prefix + "SetDefaultConfigID",
	8702: Prefix + "RegisterObserver",
	8703: Prefix + "SetLogLevel",
	8704: Prefix + "UnregisterObserver",
}

// Status strings for specific szconfigmanager messages.
var IDStatuses = map[int]string{}
