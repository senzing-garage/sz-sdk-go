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
  - errorTypeIds: An integer from An ordered list of error types to wrap the original error.
*/
func mapErrorIDtoError(errorTypeID TypeIDs) error {
	result, ok := SzErrorMap[errorTypeID]
	if !ok {
		result = ErrSzBase
	}
	return result
}

// ----------------------------------------------------------------------------
// Public Functions
// ----------------------------------------------------------------------------

/*
The Message function returns the string value from the Senzing error message.

Input
  - senzingErrorMessage: The message returned from Senzing's G2xxx_getLastException message.
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
The Code function returns the integer error code value from the Senzing error message.
Example Senzing error message: "0037E|Unknown resolved entity value '-4'"

Input
  - senzingErrorMessage: The message returned from Senzing's G2xxx_getLastException message.
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
The New function returns the integer error code value from the Senzing error message.

Input
  - senzingErrorCode: The error integer extracted from Senzing's G2xxx_getLastException message.
  - message: The message to be returned by err.Error().
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
