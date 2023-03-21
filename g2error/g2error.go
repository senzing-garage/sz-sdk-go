package g2error

import (
	"errors"
)

// ----------------------------------------------------------------------------
// Functions
// ----------------------------------------------------------------------------

/*
The G2ErrorMessage function returns the string value from the Senzing error message.

Input
  - senzingErrorMessage: The message returned from Senzing's G2xxx_getLastException message.
*/
func G2ErrorMessage(senzingErrorMessage string) string {
	result := ""
	return result
}

/*
The G2ErrorCode function returns the integer error code value from the Senzing error message.

Input
  - senzingErrorMessage: The message returned from Senzing's G2xxx_getLastException message.
*/
func G2ErrorCode(senzingErrorMessage string) int {
	result := 0
	return result
}

/*
The G2Error function returns the integer error code value from the Senzing error message.

Input
  - senzingErrorCode: The error integer extracted from Senzing's G2xxx_getLastException message.
  - message: The message to be returned by err.Error().
*/
func G2Error(senzingErrorCode int, message string) error {
	var result error = errors.New(message)
	if errorTypeIds, ok := G2ErrorTypes[senzingErrorCode]; ok {
		for _, errorTypeId := range errorTypeIds {
			switch errorTypeId {

			// Category errors.

			case G2BadUserInput:
				result = G2BadUserInputError{result}
			case G2:
				result = G2BaseError{result}
			case G2Retryable:
				result = G2RetryableError{result}
			case G2Unrecoverable:
				result = G2UnrecoverableError{result}

			// Detail errors.

			case G2Configuration:
				result = G2ConfigurationError{result}
			case G2DatabaseConnectionLost:
				result = G2DatabaseConnectionLostError{result}
			case G2Database:
				result = G2DatabaseError{result}
			case G2IncompleteRecord:
				result = G2IncompleteRecordError{result}
			case G2MalformedJson:
				result = G2MalformedJsonError{result}
			case G2MessageBuffer:
				result = G2MessageBufferError{result}
			case G2MissingConfiguration:
				result = G2MissingConfigurationError{result}
			case G2MissingDataSource:
				result = G2MissingDataSourceError{result}
			case G2ModuleEmptyMessage:
				result = G2ModuleEmptyMessageError{result}
			case G2Module:
				result = G2ModuleError{result}
			case G2ModuleGeneric:
				result = G2ModuleGenericError{result}
			case G2ModuleInvalidXML:
				result = G2ModuleInvalidXMLError{result}
			case G2ModuleLicense:
				result = G2ModuleLicenseError{result}
			case G2ModuleNotInitialized:
				result = G2ModuleNotInitializedError{result}
			case G2ModuleResolveMissingResEnt:
				result = G2ModuleResolveMissingResEntError{result}
			case G2NotFound:
				result = G2NotFoundError{result}
			case G2RepositoryPurged:
				result = G2RepositoryPurgedError{result}
			case G2RetryTimeoutExceeded:
				result = G2RetryTimeoutExceededError{result}
			case G2UnacceptableJsonKeyValue:
				result = G2UnacceptableJsonKeyValueError{result}
			case G2Unhandled:
				result = G2UnhandledError{result}

			// Default error.

			default:
				result = G2BaseError{result}
			}
		}
	}
	return result
}