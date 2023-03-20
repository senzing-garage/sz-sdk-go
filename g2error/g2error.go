package g2error

import (
	"errors"
)

// ----------------------------------------------------------------------------
// Base error
// ----------------------------------------------------------------------------

type G2BaseError struct {
	error
	// ForExampleInt    int
	// ForExampleString string
}

// ----------------------------------------------------------------------------
// "Category" errors - all based directly on G2BaseError
// ----------------------------------------------------------------------------

type G2BadUserInputError struct {
	G2BaseError
}

type G2RetryableError struct {
	G2BaseError
}

type G2UnrecoverableError struct {
	G2BaseError
}

// ----------------------------------------------------------------------------
// Errors based on G2BadUserInputError
// ----------------------------------------------------------------------------

type G2IncompleteRecordError struct {
	G2BadUserInputError
}

type G2MalformedJsonError struct {
	G2BadUserInputError
}

type G2MissingConfigurationError struct {
	G2BadUserInputError
}

type G2MissingDataSourceError struct {
	G2BadUserInputError
}

type G2NotFoundError struct {
	G2BadUserInputError
}

type G2UnacceptableJsonKeyValueError struct {
	G2BadUserInputError
}

// ----------------------------------------------------------------------------
// Errors based on G2RetryableError
// ----------------------------------------------------------------------------

type G2ConfigurationError struct {
	G2RetryableError
}
type G2DatabaseConnectionLostError struct {
	G2RetryableError
}
type G2MessageBufferError struct {
	G2RetryableError
}
type G2RepositoryPurgedError struct {
	G2RetryableError
}
type G2RetryTimeoutExceededError struct {
	G2RetryableError
}

// ----------------------------------------------------------------------------
// Errors based on G2UnrecoverableInputError
// ----------------------------------------------------------------------------

type G2DatabaseError struct {
	G2UnrecoverableError
}

type G2ModuleEmptyMessageError struct {
	G2UnrecoverableError
}

type G2ModuleError struct {
	G2UnrecoverableError
}

type G2ModuleGenericError struct {
	G2UnrecoverableError
}

type G2ModuleInvalidXMLError struct {
	G2UnrecoverableError
}

type G2ModuleLicenseError struct {
	G2UnrecoverableError
}

type G2ModuleNotInitializedError struct {
	G2UnrecoverableError
}

type G2ModuleResolveMissingResEntError struct {
	G2UnrecoverableError
}
type G2UnhandledError struct {
	G2UnrecoverableError
}

// ----------------------------------------------------------------------------
// Functions
// ----------------------------------------------------------------------------

/*
The G2ErrorMessage function returns the string value from the Senzing error message.

Input
  - ctx: A context to control lifecycle.
  - senzingErrorMessage: The message returned from the Senzing engine.
*/
func G2ErrorMessage(senzingErrorMessage string) string {
	result := ""
	return result
}

/*
The G2ErrorCode function returns the integer error code value from the Senzing error message.

Input
  - ctx: A context to control lifecycle.
  - senzingErrorMessage: The message returned from the Senzing engine.
*/
func G2ErrorCode(senzingErrorMessage string) int {
	result := 0
	return result
}

/*
The G2Error function returns the integer error code value from the Senzing error message.

Input
  - ctx: A context to control lifecycle.
  - senzingErrorMessage: The message returned from the Senzing engine.
*/
func G2Error(senzingErrorCode int, message string) error {
	var result error
	if errorTypeId, ok := G2ErrorTypes[senzingErrorCode]; ok {
		switch errorTypeId {

		// Categories

		case G2:
			result = G2BaseError{errors.New(message)}
		case G2BadUserInput:
			tmp := errors.New(message)
			result = tmp.(G2BadUserInputError)
		case G2Retryable:
			result = G2RetryableError{errors.New(message).(G2BaseError)}
		case G2Unrecoverable:
			result = G2UnrecoverableError{errors.New(message).(G2BaseError)}

		// G2BadUserInputError

		case G2IncompleteRecord:
			result = G2IncompleteRecordError{errors.New(message).(G2BadUserInputError)}
		case G2MalformedJson:
			result = G2MalformedJsonError{errors.New(message).(G2BadUserInputError)}
		case G2MissingConfiguration:
			result = G2MissingConfigurationError{errors.New(message).(G2BadUserInputError)}
		case G2MissingDataSource:
			result = G2MissingDataSourceError{errors.New(message).(G2BadUserInputError)}
		case G2NotFound:
			result = G2NotFoundError{errors.New(message).(G2BadUserInputError)}
		case G2UnacceptableJsonKeyValue:
			result = G2UnacceptableJsonKeyValueError{errors.New(message).(G2BadUserInputError)}

		// G2RetryableError

		case G2Configuration:
			result = G2ConfigurationError{errors.New(message).(G2RetryableError)}
		case G2DatabaseConnectionLost:
			result = G2DatabaseConnectionLostError{errors.New(message).(G2RetryableError)}
		case G2MessageBuffer:
			result = G2MessageBufferError{errors.New(message).(G2RetryableError)}
		case G2RepositoryPurged:
			result = G2RepositoryPurgedError{errors.New(message).(G2RetryableError)}
		case G2RetryTimeoutExceeded:
			result = G2RetryTimeoutExceededError{errors.New(message).(G2RetryableError)}

		// G2UnrecoverableError

		case G2Database:
			result = G2DatabaseError{errors.New(message).(G2UnrecoverableError)}
		case G2ModuleEmptyMessage:
			result = G2ModuleEmptyMessageError{errors.New(message).(G2UnrecoverableError)}
		case G2Module:
			result = G2ModuleError{errors.New(message).(G2UnrecoverableError)}
		case G2ModuleGeneric:
			result = G2ModuleGenericError{errors.New(message).(G2UnrecoverableError)}
		case G2ModuleInvalidXML:
			result = G2ModuleInvalidXMLError{errors.New(message).(G2UnrecoverableError)}
		case G2ModuleLicense:
			result = G2ModuleLicenseError{errors.New(message).(G2UnrecoverableError)}
		case G2ModuleNotInitialized:
			result = G2ModuleNotInitializedError{errors.New(message).(G2UnrecoverableError)}
		case G2ModuleResolveMissingResEnt:
			result = G2ModuleResolveMissingResEntError{errors.New(message).(G2UnrecoverableError)}
		case G2Unhandled:
			result = G2UnhandledError{errors.New(message).(G2UnrecoverableError)}

		// Default

		default:
			result = G2BaseError{errors.New(message)}
		}
	}

	return result
}
