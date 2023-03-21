package g2error

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2ErrorTypeIds int

// ----------------------------------------------------------------------------
// "Category" errors
// ----------------------------------------------------------------------------

type G2BadUserInputError struct {
	error
	G2ErrorTypeIds []G2ErrorTypeIds
}
type G2BaseError struct {
	error
	G2ErrorTypeIds []G2ErrorTypeIds
}
type G2RetryableError struct {
	error
	G2ErrorTypeIds []G2ErrorTypeIds
}
type G2UnrecoverableError struct {
	error
	G2ErrorTypeIds []G2ErrorTypeIds
}

// ----------------------------------------------------------------------------
// Detail errors
// ----------------------------------------------------------------------------

type G2ConfigurationError struct{ error }
type G2DatabaseConnectionLostError struct{ error }
type G2DatabaseError struct{ error }
type G2IncompleteRecordError struct{ error }
type G2MalformedJsonError struct{ error }
type G2MessageBufferError struct{ error }
type G2MissingConfigurationError struct{ error }
type G2MissingDataSourceError struct{ error }
type G2ModuleEmptyMessageError struct{ error }
type G2ModuleError struct{ error }
type G2ModuleGenericError struct{ error }
type G2ModuleInvalidXMLError struct{ error }
type G2ModuleLicenseError struct{ error }
type G2ModuleNotInitializedError struct{ error }
type G2ModuleResolveMissingResEntError struct{ error }
type G2NotFoundError struct{ error }
type G2RepositoryPurgedError struct{ error }
type G2RetryTimeoutExceededError struct{ error }
type G2UnacceptableJsonKeyValueError struct{ error }
type G2UnhandledError struct{ error }

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
// Note: The lists of G2ErrorTypeIds are from innermost error to outer most error.
// Example:  #99901 is G2BadUserInputError{G2ModuleInvalidXMLError{errors.New(message)}}
var G2ErrorTypes = map[int][]G2ErrorTypeIds{
	1:     {G2ModuleInvalidXML, G2Unrecoverable},
	2:     {G2Unhandled, G2Unrecoverable},
	7:     {G2ModuleEmptyMessage, G2Unrecoverable},
	10:    {G2RetryTimeoutExceeded, G2Retryable},
	23:    {G2UnacceptableJsonKeyValue, G2BadUserInput},
	24:    {G2UnacceptableJsonKeyValue, G2BadUserInput},
	25:    {G2UnacceptableJsonKeyValue, G2BadUserInput},
	26:    {G2UnacceptableJsonKeyValue, G2BadUserInput},
	27:    {G2NotFound, G2BadUserInput},
	32:    {G2UnacceptableJsonKeyValue, G2BadUserInput},
	33:    {G2NotFound, G2BadUserInput},
	34:    {G2Configuration, G2Retryable},
	35:    {G2Configuration, G2Retryable},
	36:    {G2Configuration, G2Retryable},
	37:    {G2NotFound, G2BadUserInput},
	47:    {G2ModuleGeneric, G2Unrecoverable},
	48:    {G2ModuleNotInitialized, G2Unrecoverable},
	49:    {G2ModuleNotInitialized, G2Unrecoverable},
	50:    {G2ModuleNotInitialized, G2Unrecoverable},
	51:    {G2UnacceptableJsonKeyValue, G2BadUserInput},
	53:    {G2ModuleNotInitialized, G2Unrecoverable},
	54:    {G2RepositoryPurged, G2Retryable},
	61:    {G2Configuration, G2Retryable},
	62:    {G2Configuration, G2Retryable},
	63:    {G2ModuleNotInitialized, G2Unrecoverable},
	64:    {G2Configuration, G2Retryable},
	999:   {G2ModuleLicense, G2Unrecoverable},
	1001:  {G2Database, G2Unrecoverable},
	1007:  {G2DatabaseConnectionLost, G2Retryable},
	2089:  {G2NotFound, G2BadUserInput},
	2134:  {G2ModuleResolveMissingResEnt, G2Unrecoverable},
	2208:  {G2Configuration, G2Retryable},
	7221:  {G2Configuration, G2Retryable},
	7426:  {G2BadUserInput},
	7344:  {G2NotFound, G2BadUserInput},
	9000:  {G2ModuleLicense, G2Unrecoverable},
	30020: {G2UnacceptableJsonKeyValue, G2BadUserInput},
	30110: {G2MessageBuffer, G2Retryable},
	30111: {G2MessageBuffer, G2Retryable},
	30112: {G2MessageBuffer, G2Retryable},
	30121: {G2MalformedJson, G2BadUserInput},
	30122: {G2MalformedJson, G2BadUserInput},
	30123: {G2MalformedJson, G2BadUserInput},

	// The 999nn series is for testing.

	99900: {G2},
	99901: {G2ModuleInvalidXML, G2BadUserInput},
}
