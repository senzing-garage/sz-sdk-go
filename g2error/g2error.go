package g2error

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// ----------------------------------------------------------------------------
// Private Functions
// ----------------------------------------------------------------------------

func isIn(needle G2ErrorTypeIds, haystack []G2ErrorTypeIds) bool {
	for _, g2ErrorTypeId := range haystack {
		if needle == g2ErrorTypeId {
			return true
		}
	}
	return false
}

// ----------------------------------------------------------------------------
// Public Functions
// ----------------------------------------------------------------------------

/*
The G2ErrorMessage function returns the string value from the Senzing error message.

Input
  - senzingErrorMessage: The message returned from Senzing's G2xxx_getLastException message.
*/
func G2ErrorMessage(senzingErrorMessage string) string {
	result := ""
	splits := strings.Split(senzingErrorMessage, "|")
	if len(splits) > 1 {
		result = strings.TrimSpace(splits[1])
	}
	return result
}

/*
The G2ErrorCode function returns the integer error code value from the Senzing error message.
Example Senzing error message: "0037E|Unknown resolved entity value '-4'"

Input
  - senzingErrorMessage: The message returned from Senzing's G2xxx_getLastException message.
*/
func G2ErrorCode(senzingErrorMessage string) int {
	result := 0
	splits := strings.Split(senzingErrorMessage, "|")
	if len(splits) > 0 {

		// Make a Regex to say we only want numbers.

		regularExpression, err := regexp.Compile("[^0-9]+")
		if err != nil {
			return result
		}
		numericOnlyString := regularExpression.ReplaceAllString(splits[0], "")
		result, err := strconv.Atoi(numericOnlyString)
		if err == nil {
			return result
		}
	}
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
				result = G2BadUserInputError{
					error:          result,
					G2ErrorTypeIds: errorTypeIds,
				}
			case G2Base:
				result = G2BaseError{
					error:          result,
					G2ErrorTypeIds: errorTypeIds,
				}
			case G2Retryable:
				result = G2RetryableError{
					error:          result,
					G2ErrorTypeIds: errorTypeIds,
				}
			case G2Unrecoverable:
				result = G2UnrecoverableError{
					error:          result,
					G2ErrorTypeIds: errorTypeIds,
				}

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
				result = G2BaseError{
					error:          result,
					G2ErrorTypeIds: errorTypeIds,
				}
			}
		}
	}
	return result
}

/*
The Is function determines if an error is of a certain type.

Input
  - err: The error to be tested.
  - errorType: The error type desired.
*/
func Is(err error, errorType G2ErrorTypeIds) bool {
	if errors.As(err, &G2BadUserInputError{}) {
		return isIn(errorType, err.(G2BadUserInputError).G2ErrorTypeIds)
	}
	if errors.As(err, &G2BaseError{}) {
		return isIn(errorType, err.(G2BaseError).G2ErrorTypeIds)
	}
	if errors.As(err, &G2RetryableError{}) {
		return isIn(errorType, err.(G2RetryableError).G2ErrorTypeIds)
	}
	if errors.As(err, &G2UnrecoverableError{}) {
		return isIn(errorType, err.(G2UnrecoverableError).G2ErrorTypeIds)
	}
	return false
}

/*
The IsInList function determines if an error is of a certain type in a list.
This is a convenience function to avoid calling Is() repeatedly.

Input
  - err: The error to be tested.
  - errorType: A list of error types desired.
*/
func IsInList(err error, errorType []G2ErrorTypeIds) bool {
	result := false
	for _, g2ErrorTypeId := range errorType {
		if Is(err, g2ErrorTypeId) {
			return true
		}
	}
	return result
}
