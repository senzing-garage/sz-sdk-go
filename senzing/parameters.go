package senzing

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

/*
Empty values for parameter.
*/
const (
	SzInitializeWithDefaultConfiguration int64  = 0  // Use the default configuration
	SzNoAttributes                       string = "" // No attributes requested
	SzNoAvoidance                        string = "" // No avoidances requested
	SzNoLogging                          int64  = 0  // No logging requested
	SzNoRequiredDatasources              string = "" // No datasources requested
	SzNoSearchProfile                    string = "" // No search profile requested
	SzVerboseLogging                     int64  = 1  // Verbose logging requested
	SzWithoutInfo                        int64  = 0  // No returned information requested
)
