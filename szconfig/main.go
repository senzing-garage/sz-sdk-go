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
var IDMessages = map[int]string{
	1:    "Enter " + Prefix + "AddDataSource(%s).",
	2:    "Exit  " + Prefix + "AddDataSource(%s) returned (%s, %v).",
	3:    "Enter " + Prefix + "ClearLastException().",
	4:    "Exit  " + Prefix + "ClearLastException() returned (%v).",
	5:    "Enter " + Prefix + "CloseConfig(%v).",
	6:    "Exit  " + Prefix + "CloseConfig(%v) returned (%v).",
	7:    "Enter " + Prefix + "CreateConfig().",
	8:    "Exit  " + Prefix + "CreateConfig() returned (%v, %v).",
	9:    "Enter " + Prefix + "DeleteDataSource(%s).",
	10:   "Exit  " + Prefix + "DeleteDataSource(%s) returned (%v).",
	11:   "Enter " + Prefix + "Destroy().",
	12:   "Exit  " + Prefix + "Destroy() returned (%v).",
	13:   "Enter " + Prefix + "Export().",
	14:   "Exit  " + Prefix + "Export() returned (%s, %v).",
	15:   "Enter " + Prefix + "GetDataSources().",
	16:   "Exit  " + Prefix + "GetDataSources() returned (%s, %v).",
	17:   "Enter " + Prefix + "GetLastException().",
	18:   "Exit  " + Prefix + "GetLastException() returned (%s, %v).",
	19:   "Enter " + Prefix + "GetLastExceptionCode().",
	20:   "Exit  " + Prefix + "GetLastExceptionCode() returned (%d, %v).",
	21:   "Enter " + Prefix + "ImportConfig(%s).",
	22:   "Exit  " + Prefix + "ImportConfig(%s) returned (%v, %v).",
	23:   "Enter " + Prefix + "Initialize(%s, %s, %d).",
	24:   "Exit  " + Prefix + "Initialize(%s, %s, %d) returned (%v).",
	703:  "Enter " + Prefix + "RegisterObserver(%s).",
	704:  "Exit  " + Prefix + "RegisterObserver(%s) returned (%v).",
	705:  "Enter " + Prefix + "SetLogLevel(%s).",
	706:  "Exit  " + Prefix + "SetLogLevel(%s) returned (%v).",
	707:  "Enter " + Prefix + "UnregisterObserver(%s).",
	708:  "Exit  " + Prefix + "UnregisterObserver(%s) returned (%v).",
	4001: Prefix + "SzConfig_addDataSource(%v, %s) failed. Return code: %d",
	4002: Prefix + "SzConfig_close(%v) failed. Return code: %d",
	4003: Prefix + "SzConfig_create() failed. Return code: %d",
	4004: Prefix + "SzConfig_deleteDataSource(%v, %s) failed. Return code: %d",
	4005: Prefix + "SzConfig_destroy() failed. Return code: %d",
	4006: Prefix + "SzConfig_getLastException() failed. Return code: %d",
	4007: Prefix + "SzConfig_init(%s, %s, %d) failed. Return code: %d",
	4008: Prefix + "SzConfig_listDataSources(%v) failed. Return code: %d",
	4009: Prefix + "SzConfig_load(%s) failed. Return code: %d",
	4010: Prefix + "SzConfig_save(%v) failed. Return code: %d",
	8001: Prefix + "AddDataSource",
	8002: Prefix + "CloseConfig",
	8003: Prefix + "CreateConfig",
	8004: Prefix + "DeleteDataSource",
	8005: Prefix + "Destroy",
	8006: Prefix + "Export",
	8007: Prefix + "Initialize",
	8008: Prefix + "GetDataSources",
	8009: Prefix + "ImportConfig",
	8702: Prefix + "RegisterObserver",
	8703: Prefix + "SetLogLevel",
	8704: Prefix + "UnregisterObserver",
}

// Status strings for specific szconfig messages.
var IDStatuses = map[int]string{}
