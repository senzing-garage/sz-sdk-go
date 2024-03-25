package szerror

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// See https://github.com/senzing-garage/go-logging/blob/main/messageformat/messageformat_senzing.go
type MessageFormatSenzing struct {
	Errors interface{} `json:"errors,omitempty"` // List of errors.
	Text   interface{} `json:"text,omitempty"`   // Message text.
}

// ----------------------------------------------------------------------------
// Private Functions
// ----------------------------------------------------------------------------

/*
With recursion, extractErrorTexts() parses JSON like:

	"text": "x",
	"errors": [{
		"text": {
			"text": "y",
			"errors": [{
				"text": {
					"text": "z",
					"errors": [{
						"text": "1019E|Datastore schema..."

and returns something like []string{"x", "y", "z", "1019E|Datastore schema..."}
*/
func extractErrorTexts(messageErrors []interface{}, messageTexts []string) ([]string, error) {
	var err error = nil

	// All "text" string values will be aggregated into errorTexts.

	newMessageTexts := []string{}
	for _, messageError := range messageErrors {
		messageText := ""
		switch typedMessageError := messageError.(type) {
		case map[string]interface{}:
			switch typedMessageErrorText := typedMessageError["text"].(type) {
			case map[string]interface{}:
				errorValue, ok := typedMessageErrorText["errors"].([]interface{})
				if ok {
					newMessageTexts, err = extractErrorTexts(errorValue, newMessageTexts)
					if err != nil {
						return append(messageTexts, newMessageTexts...), err
					}
				}
			case string:
				messageText = typedMessageErrorText
			}
		case string:
			messageText = typedMessageError
		}

		if len(messageText) > 0 {
			newMessageTexts = append(newMessageTexts, messageText)
		}
	}
	return append(messageTexts, newMessageTexts...), err
}

/*
extractErrorNumber scans nested messages for a Senzing error code number.
*/
func extractErrorNumber(message string) (int, error) {

	// If non-JSON submitted, inspect the string and return.

	if !isJson(message) {
		return SzErrorCode(message), nil
	}

	// All "text" values will be aggregated into errorTexts.

	errorTexts := []string{}

	// Parse JSON into type-structure.

	messageFormatSenzing := &MessageFormatSenzing{}
	err := json.Unmarshal([]byte(message), &messageFormatSenzing)
	if err != nil {
		return -1, err
	}

	// If exists, add "text" to list.

	if messageFormatSenzing.Text != nil {
		messageText, ok := messageFormatSenzing.Text.(string)
		if ok {
			errorTexts = append(errorTexts, messageText)
		}
	}

	// Recurse through nested "error" JSON stanzas to harvest "text".

	if messageFormatSenzing.Errors != nil {
		errorTexts, err = extractErrorTexts(messageFormatSenzing.Errors.([]interface{}), errorTexts)
		if err != nil {
			return -1, err
		}
	}

	// Loop through harvested "texts" and return the first one that produces a SzErrorCode.

	for _, errorText := range errorTexts {
		result := SzErrorCode(errorText)
		if result > 0 {
			return result, nil
		}
	}

	// No SZErrorCode found.

	return -1, err
}

/*
isIn determines if a SzErrorTypeId is in a list of SzErrorTypeIds.
*/
func isIn(needle SzErrorTypeIds, haystack []SzErrorTypeIds) bool {
	for _, szErrorTypeId := range haystack {
		if needle == szErrorTypeId {
			return true
		}
	}
	return false
}

/*
isJson determines if the string is syntactically JSON.
*/
func isJson(unknownString string) bool {
	unknownStringUnescaped, err := strconv.Unquote(unknownString)
	if err != nil {
		unknownStringUnescaped = unknownString
	}
	var jsonString json.RawMessage
	return json.Unmarshal([]byte(unknownStringUnescaped), &jsonString) == nil
}

/*
wrapError return an error that has nested errors.

Input
  - originalError: The error containing the message to be maintained (i.e. err.Error()).
  - errorTypeIds: An ordered list of error types to wrap the original error.
*/
func wrapError(originalError error, errorTypeIds []SzErrorTypeIds) error {
	result := originalError
	for _, errorTypeId := range errorTypeIds {
		switch errorTypeId {

		// Category errors.

		case SzBadInput:
			result = SzBadInputError{
				error:          result,
				SzErrorTypeIds: errorTypeIds,
			}
		case SzBase:
			result = SzBaseError{
				error:          result,
				SzErrorTypeIds: errorTypeIds,
			}
		case SzConfiguration:
			result = SzConfigurationError{
				error:          result,
				SzErrorTypeIds: errorTypeIds,
			}
		case SzRetryable:
			result = SzRetryableError{
				error:          result,
				SzErrorTypeIds: errorTypeIds,
			}
		case SzUnrecoverable:
			result = SzUnrecoverableError{
				error:          result,
				SzErrorTypeIds: errorTypeIds,
			}

			// Detail errors.

		case SzDatabase:
			result = SzDatabaseError{result}
		case SzDatabaseConnectionLost:
			result = SzDatabaseConnectionLostError{result}
		case SzLicense:
			result = SzLicenseError{result}
		case SzNotFound:
			result = SzNotFoundError{result}
		case SzNotInitialized:
			result = SzNotInitializedError{result}
		case SzRetryTimeoutExceeded:
			result = SzRetryTimeoutExceededError{result}
		case SzUnhandled:
			result = SzUnhandledError{result}
		case SzUnknownDatasource:
			result = SzUnhandledError{result}

		// Default error.

		default:
			result = SzBaseError{
				error:          result,
				SzErrorTypeIds: errorTypeIds,
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
	var errorTypeIds []SzErrorTypeIds

	// Handle case of nil.

	if originalError == nil || desiredTypeError == nil {
		return originalError
	}

	result := originalError

	// Get the desiredTypeError's SzErrorTypeIds value.

	switch {
	case errors.As(desiredTypeError, &SzBadInputError{}):
		errorTypeIds = desiredTypeError.(SzBadInputError).SzErrorTypeIds
	case errors.As(desiredTypeError, &SzBaseError{}):
		errorTypeIds = desiredTypeError.(SzBaseError).SzErrorTypeIds
	case errors.As(desiredTypeError, &SzConfigurationError{}):
		errorTypeIds = desiredTypeError.(SzConfigurationError).SzErrorTypeIds
	case errors.As(desiredTypeError, &SzRetryableError{}):
		errorTypeIds = desiredTypeError.(SzRetryableError).SzErrorTypeIds
	case errors.As(desiredTypeError, &SzUnrecoverableError{}):
		errorTypeIds = desiredTypeError.(SzUnrecoverableError).SzErrorTypeIds
	}

	// Cast.

	if len(errorTypeIds) > 0 {
		if IsInList(originalError, AllSzErrorTypes) {
			result = errors.New(result.Error())
		}
		result = wrapError(result, errorTypeIds)
	}
	return result
}

/*
The Convert function uses the error message from the originalError to determine
the appropriate szerror type hierarchy.

Input
  - originalError: The error containing the message to be analyzed.
*/
func Convert(originalError error) error {

	result := originalError
	if result != nil {
		extractedErrorNumber, err := extractErrorNumber(originalError.Error())
		if err != nil {
			return originalError
		}
		if extractedErrorNumber < 1 {
			return originalError
		}
		result = SzError(extractedErrorNumber, originalError.Error())
	}
	return result
}

/*
The SzErrorMessage function returns the string value from the Senzing error message.

Input
  - senzingErrorMessage: The message returned from Senzing's G2xxx_getLastException message.
*/
func SzErrorMessage(senzingErrorMessage string) string {
	result := ""
	splits := strings.Split(senzingErrorMessage, "|")
	if len(splits) > 1 {
		result = strings.TrimSpace(splits[1])
	}
	return result
}

/*
The SzErrorCode function returns the integer error code value from the Senzing error message.
Example Senzing error message: "0037E|Unknown resolved entity value '-4'"

Input
  - senzingErrorMessage: The message returned from Senzing's G2xxx_getLastException message.
*/
func SzErrorCode(senzingErrorMessage string) int {

	result := 0
	if strings.Contains(senzingErrorMessage, "|") {
		splits := strings.Split(senzingErrorMessage, "|")

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
The SzError function returns the integer error code value from the Senzing error message.

Input
  - senzingErrorCode: The error integer extracted from Senzing's G2xxx_getLastException message.
  - message: The message to be returned by err.Error().
*/
func SzError(senzingErrorCode int, message string) error {
	var result error = errors.New(message)
	if errorTypeIds, ok := SzErrorTypes[senzingErrorCode]; ok {
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
func Is(err error, errorType SzErrorTypeIds) bool {
	if errors.As(err, &SzBadInputError{}) {
		return isIn(errorType, err.(SzBadInputError).SzErrorTypeIds)
	}
	if errors.As(err, &SzBaseError{}) {
		return isIn(errorType, err.(SzBaseError).SzErrorTypeIds)
	}
	if errors.As(err, &SzConfigurationError{}) {
		return isIn(errorType, err.(SzConfigurationError).SzErrorTypeIds)
	}
	if errors.As(err, &SzRetryableError{}) {
		return isIn(errorType, err.(SzRetryableError).SzErrorTypeIds)
	}
	if errors.As(err, &SzUnrecoverableError{}) {
		return isIn(errorType, err.(SzUnrecoverableError).SzErrorTypeIds)
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
func IsInList(err error, errorType []SzErrorTypeIds) bool {
	result := false
	for _, szErrorTypeId := range errorType {
		if Is(err, szErrorTypeId) {
			return true
		}
	}
	return result
}
