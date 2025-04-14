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
	15:   "Enter " + Prefix + "Initialize(%s, %s, %d, %d).",
	16:   "Exit  " + Prefix + "Initialize(%s, %s, %d, %d) returned (%v).",
	17:   "Enter " + Prefix + "PurgeRepository().",
	18:   "Exit  " + Prefix + "PurgeRepository() returned (%v).",
	19:   "Enter " + Prefix + "Reinitialize(%d).",
	20:   "Exit  " + Prefix + "Reinitialize(%d) returned (%v).",
	703:  "Enter " + Prefix + "RegisterObserver(%s).",
	704:  "Exit  " + Prefix + "RegisterObserver(%s) returned (%v).",
	705:  "Enter " + Prefix + "SetLogLevel(%s).",
	706:  "Exit  " + Prefix + "SetLogLevel(%s) returned (%v).",
	707:  "Enter " + Prefix + "UnregisterObserver(%s).",
	708:  "Exit  " + Prefix + "UnregisterObserver(%s) returned (%v).",
	4001: Prefix + "SzDiagnostic_checkDatastorePerformance(%d) failed. Return code: %d",
	4002: Prefix + "SzDiagnostic_destroy() failed.  Return code: %d",
	4003: Prefix + "SzDiagnostic_getDatastoreInfo() failed. Return code: %d",
	4004: Prefix + "SzDiagnostic_getFeature(%d) failed. Return code: %d",
	4005: Prefix + "SzDiagnostic_init(%s, %s, %d) failed. Return code: %d",
	4006: Prefix + "SzDiagnostic_initWithConfigID(%s, %s, %d, %d) failed. Return code: %d",
	4007: Prefix + "SzDiagnostic_purgeRepository() failed. Return Code: %d",
	4008: Prefix + "SzDiagnostic_reinit(%d) failed. Return Code: %d",
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
