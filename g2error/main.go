package g2error

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2ErrorTypeIds int

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const (
	G2 G2ErrorTypeIds = iota
	G2BadUserInput
	G2Configuration
	G2Database
	G2DatabaseConnectionLost
	G2IncompleteRecord
	G2MalformedJson
	G2MessageBuffer
	G2MissingConfiguration
	G2MissingDataSource
	G2Module
	G2ModuleEmptyMessage
	G2ModuleGeneric
	G2ModuleInvalidXML
	G2ModuleLicense
	G2ModuleNotInitialized
	G2ModuleResolveMissingResEnt
	G2NotFound
	G2RepositoryPurged
	G2Retryable
	G2RetryTimeoutExceeded
	G2UnacceptableJsonKeyValue
	G2Unhandled
	G2Unrecoverable
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for g2engine implementations.
var G2ErrorTypes = map[int]G2ErrorTypeIds{
	1:     G2ModuleInvalidXML,
	2:     G2Unhandled,
	7:     G2ModuleEmptyMessage,
	10:    G2RetryTimeoutExceeded,
	23:    G2UnacceptableJsonKeyValue,
	24:    G2UnacceptableJsonKeyValue,
	25:    G2UnacceptableJsonKeyValue,
	26:    G2UnacceptableJsonKeyValue,
	27:    G2NotFound,
	32:    G2UnacceptableJsonKeyValue,
	33:    G2NotFound,
	34:    G2Configuration,
	35:    G2Configuration,
	36:    G2Configuration,
	37:    G2NotFound,
	47:    G2ModuleGeneric,
	48:    G2ModuleNotInitialized,
	49:    G2ModuleNotInitialized,
	50:    G2ModuleNotInitialized,
	51:    G2UnacceptableJsonKeyValue,
	53:    G2ModuleNotInitialized,
	54:    G2RepositoryPurged,
	61:    G2Configuration,
	62:    G2Configuration,
	63:    G2ModuleNotInitialized,
	64:    G2Configuration,
	999:   G2ModuleLicense,
	1001:  G2Database,
	1007:  G2DatabaseConnectionLost,
	2089:  G2NotFound,
	2134:  G2ModuleResolveMissingResEnt,
	2208:  G2Configuration,
	7221:  G2Configuration,
	7426:  G2BadUserInput,
	7344:  G2NotFound,
	9000:  G2ModuleLicense,
	30020: G2UnacceptableJsonKeyValue,
	30110: G2MessageBuffer,
	30111: G2MessageBuffer,
	30112: G2MessageBuffer,
	30121: G2MalformedJson,
	30122: G2MalformedJson,
	30123: G2MalformedJson,

	// The 999nn series is for testing.

	99900: G2,
	99901: G2BadUserInput,
}
