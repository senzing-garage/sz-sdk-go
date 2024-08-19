package szproduct

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Log message prefix.
const Prefix = "szproduct."

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for szproduct implementations.
var IDMessages = map[int]string{
	1:    "Enter " + Prefix + "ClearLastException().",
	2:    "Exit  " + Prefix + "ClearLastException() returned (%v).",
	3:    "Enter " + Prefix + "Destroy().",
	4:    "Exit  " + Prefix + "Destroy() returned (%v).",
	5:    "Enter " + Prefix + "GetLastException().",
	6:    "Exit  " + Prefix + "GetLastException() returned (%s, %v).",
	7:    "Enter " + Prefix + "GetLastExceptionCode().",
	8:    "Exit  " + Prefix + "GetLastExceptionCode() returned (%d, %v).",
	9:    "Enter " + Prefix + "GetLicense().",
	10:   "Exit  " + Prefix + "GetLicense() returned (%s, %v).",
	11:   "Enter " + Prefix + "GetVersion().",
	12:   "Exit  " + Prefix + "GetVersion() returned (%s, %v).",
	13:   "Enter " + Prefix + "Initialize(%s, %s, %d).",
	14:   "Exit  " + Prefix + "Initialize(%s, %s, %d) returned (%v).",
	703:  "Enter " + Prefix + "RegisterObserver(%s).",
	704:  "Exit  " + Prefix + "RegisterObserver(%s) returned (%v).",
	705:  "Enter " + Prefix + "SetLogLevel(%s).",
	706:  "Exit  " + Prefix + "SetLogLevel(%s) returned (%v).",
	707:  "Enter " + Prefix + "UnregisterObserver(%s).",
	708:  "Exit  " + Prefix + "UnregisterObserver(%s) returned (%v).",
	4001: Prefix + "SzProduct_destroy() failed. Return code: %d",
	4002: Prefix + "SzProduct_init(%s, %s, %d) failed. Return code: %d",
	8001: Prefix + "Destroy",
	8002: Prefix + "Initialize",
	8003: Prefix + "GetLicense",
	8004: Prefix + "GetVersion",
	8702: Prefix + "RegisterObserver",
	8703: Prefix + "SetLogLevel",
	8704: Prefix + "UnregisterObserver",
}

// Status strings for specific szproduct messages.
var IDStatuses = map[int]string{}
