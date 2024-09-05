package szerror

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// ----------------------------------------------------------------------------
// Private Functions
// ----------------------------------------------------------------------------

/*
mapErrorIDtoError returns the error corresponding to the requested errorTypeID integer.

Input
  - errorTypeID: An integer from An ordered list of error types to wrap the original error.
*/
func mapErrorIDtoError(errorTypeID TypeIDs) error {
	result, ok := SzErrorMap[errorTypeID]
	if !ok {
		result = ErrSz
	}
	return result
}

// ----------------------------------------------------------------------------
// Public Functions
// ----------------------------------------------------------------------------

/*
Function Message returns the string value from the Senzing error message.
The string is defined as the text after the pipe ("|") symbol.

Input
  - senzingErrorMessage: The message returned from Senzing's SzXxx_getLastException message.

Output
  - The text of the message after the pipe symbol.
*/
func Message(senzingErrorMessage string) string {
	result := ""
	splits := strings.Split(senzingErrorMessage, "|")
	if len(splits) > 1 {
		result = strings.TrimSpace(splits[1])
	}
	return result
}

/*
Function Code returns the integer error code value from the Senzing error message.
The integer is defined as the numerical value in the text before the pipe ("|") symbol.

Input
  - senzingErrorMessage: The message returned from Senzing's Szxxx_getLastException message.

Output
  - The integer portion of the message before the pipe symbol.
*/
func Code(senzingErrorMessage string) int {
	result := 0
	if strings.Contains(senzingErrorMessage, "|") {
		splits := strings.Split(senzingErrorMessage, "|")

		// Make a Regex to say we only want numbers.

		regularExpression := regexp.MustCompile("[^0-9]+")
		numericOnlyString := regularExpression.ReplaceAllString(splits[0], "")
		result, err := strconv.Atoi(numericOnlyString)
		if err == nil {
			return result
		}
	}
	return result
}

/*
Function New returns an error based on the error code and message.

Input
  - senzingErrorCode: The error integer extracted from Senzing's Szxxx_getLastException message.
  - message: The message to be returned by err.Error().

Output
  - An error conforming to the error code and message.
*/
func New(senzingErrorCode int, message string) error {
	result := []error{}
	if errorTypeIDs, ok := SzErrorTypes[senzingErrorCode]; ok {
		for _, errorTypeID := range errorTypeIDs {
			result = append(result, mapErrorIDtoError(errorTypeID))
		}
	}
	result = append(result, errors.New(message))
	return errors.Join(result...)
}
