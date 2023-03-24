package g2error

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// See https://github.com/Senzing/go-logging/blob/main/messageformat/messageformat_senzing.go
type MessageFormatSenzing struct {
	Errors interface{} `json:"errors,omitempty"` // List of errors.
	Text   interface{} `json:"text,omitempty"`   // Message text.
}

// ----------------------------------------------------------------------------
// Private Functions
// ----------------------------------------------------------------------------

func extractErrorTexts(messageErrors []interface{}, messageTexts []string) ([]string, error) {
	var err error = nil

	fmt.Printf(">>>>> messageErrors: %#v\n", messageErrors)

	// All "text" values will be aggregated into errorTexts.

	newMessageTexts := []string{}

	for _, messageError := range messageErrors {

		// Parse XXX into type-structure.

		fmt.Printf(">>>>> messageError: %#v\n", messageError)

		abc, ok := messageError.(map[string]interface{})
		if !ok {
			fmt.Printf(">>>>> abc failed.\n")
			return append(newMessageTexts, messageTexts...), nil
		}

		def1, ok := abc["text"].(string)
		if ok {
			fmt.Printf(">>>>> def1 succeeded.\n")
			newMessageTexts = append(newMessageTexts, def1)
		}

		def, ok := abc["text"].(map[string]interface{})
		if !ok {
			fmt.Printf(">>>>> def failed.\n")
			return append(newMessageTexts, messageTexts...), nil
		}

		ghi, ok := def["text"].(string)
		if ok {
			fmt.Printf(">>>>> ghi succeeded.\n")
			newMessageTexts = append(newMessageTexts, ghi)
		}

		jkl, ok := def["errors"].([]interface{})
		if !ok {
			return append(newMessageTexts, messageTexts...), nil
		}

		fmt.Printf(">>>>> jkl succeeded.\n")
		moreMessageTexts, err := extractErrorTexts(jkl, messageTexts)
		if err != nil {
			fmt.Printf(">>>>> extractErrorTexts failed.\n")
		}
		newMessageTexts = append(newMessageTexts, moreMessageTexts...)

		// fmt.Printf(">>>>> abc: %#v\n", abc)

		// xyz, ok := abc.(MessageFormatSenzing)
		// if !ok {
		// 	fmt.Printf(">>>>> xyz failed.\n")
		// 	return append(newMessageTexts, messageTexts...), nil
		// }

		// if xyz.Text != nil {
		// 	newMessageTexts = append(newMessageTexts, xyz.Text.(string))
		// }

		// if xyz.Errors != nil {
		// 	newMessageTexts, err = extractErrorTexts(xyz.Errors.([]interface{}), newMessageTexts)
		// }
	}

	return append(messageTexts, newMessageTexts...), err
}

func extractErrorNumber(message string) (int, error) {
	result := 0
	var err error = nil
	fmt.Printf(">>> message:  %s\n", message)

	if !isJson(message) {
		return result, errors.New("not JSON format")
	}

	// All "text" values will be aggregated into errorTexts.

	errorTexts := []string{}

	// Parse JSON into type-structure.

	messageFormatSenzing := &MessageFormatSenzing{}
	err = json.Unmarshal([]byte(message), &messageFormatSenzing)
	if err != nil {
		fmt.Printf(">>>>> Error:  %s\n", err)
		return -1, err
	}

	// If exist, add "text" to list.

	if messageFormatSenzing.Text != nil {
		messageText, ok := messageFormatSenzing.Text.(string)
		if ok {
			errorTexts = append(errorTexts, messageText)
		}
	}

	// Recurse through "error" JSON stanzas to harvest "text".

	if messageFormatSenzing.Errors != nil {
		errorTexts, err = extractErrorTexts(messageFormatSenzing.Errors.([]interface{}), errorTexts)
	}

	for _, errorText := range errorTexts {
		fmt.Printf(">>>>> errorText: %s\n", errorText)
		result := G2ErrorCode(errorText)
		if result > 0 {
			fmt.Printf(">>>>> Success: %d\n", result)
			return result, err
		}
	}
	fmt.Printf(">>>>> Dropped out bottom\n")
	return -1, err
}

func isIn(needle G2ErrorTypeIds, haystack []G2ErrorTypeIds) bool {
	for _, g2ErrorTypeId := range haystack {
		if needle == g2ErrorTypeId {
			return true
		}
	}
	return false
}

func isJson(unknownString string) bool {
	unknownStringUnescaped, err := strconv.Unquote(unknownString)
	if err != nil {
		unknownStringUnescaped = unknownString
	}
	var jsonString json.RawMessage
	return json.Unmarshal([]byte(unknownStringUnescaped), &jsonString) == nil
}

func wrapError(originalError error, errorTypeIds []G2ErrorTypeIds) error {
	result := originalError
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
	return result
}

// ----------------------------------------------------------------------------
// Public Functions
// ----------------------------------------------------------------------------

/*
The Cast function will cast an originalError into the nested types specified by the desiredTypeError.
This essentially creates a new error having the message from the originalError and the type
structure of the desiredTypeError.

Input
  - originalError: The error containing the message to be maintained (i.e. err.Error()).
  - desiredTypeError: An error having the nested types desired in the resulting error.
*/
func Cast(originalError error, desiredTypeError error) error {
	var errorTypeIds []G2ErrorTypeIds
	result := originalError

	// Get the desiredTypeError's G2ErrorTypeIds value.

	switch {
	case errors.As(desiredTypeError, &G2BadUserInputError{}):
		errorTypeIds = desiredTypeError.(G2BadUserInputError).G2ErrorTypeIds
	case errors.As(desiredTypeError, &G2BaseError{}):
		errorTypeIds = desiredTypeError.(G2BaseError).G2ErrorTypeIds
	case errors.As(desiredTypeError, &G2RetryableError{}):
		errorTypeIds = desiredTypeError.(G2RetryableError).G2ErrorTypeIds
	case errors.As(desiredTypeError, &G2UnrecoverableError{}):
		errorTypeIds = desiredTypeError.(G2UnrecoverableError).G2ErrorTypeIds
	}

	// Cast.

	if len(errorTypeIds) > 0 {
		if IsInList(originalError, AllG2ErrorTypes) {
			result = errors.New(result.Error())
		}
		result = wrapError(result, errorTypeIds)
	}
	return result
}

/*
The Convert function uses the error message from the originalError to determine
the appropriate g2error type hierarchy.

Input
  - originalError: The error containing the message to be analyzed.
*/
func Convert(originalError error) error {
	extractedErrorNumber, err := extractErrorNumber(originalError.Error())
	if err != nil {
		return originalError
	}
	if extractedErrorNumber < 1 {
		return originalError
	}
	return G2Error(extractedErrorNumber, originalError.Error())
}

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
		result = wrapError(result, errorTypeIds)
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
