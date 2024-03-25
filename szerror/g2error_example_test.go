package szerror

import (
	"errors"
	"fmt"
	"reflect"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleCast() {
	originalError := errors.New("Original message")
	senzingErrorMessage := "33E|Test message" // Example message from Senzing G2 engine.
	desiredTypeError := SzError(SzErrorCode(senzingErrorMessage), `{"messageId": 1}`)
	err := Cast(originalError, desiredTypeError)
	fmt.Printf("Error type: %s; Error message: %s", reflect.TypeOf(err), err.Error())
	// Output: Error type: szerror.SzBadInputError; Error message: Original message
}

func ExampleSzErrorMessage() {
	senzingErrorMessage := "33E|Test message" // Example message from Senzing G2 engine.
	result := SzErrorMessage(senzingErrorMessage)
	fmt.Println(result)
	// Output: Test message
}

func ExampleSzErrorCode() {
	senzingErrorMessage := "33E|Test message" // Example message from Senzing G2 engine.
	result := SzErrorCode(senzingErrorMessage)
	fmt.Println(result)
	// Output: 33
}

func ExampleSzError() {
	senzingErrorMessage := "33E|Test message" // Example message from Senzing G2 engine.
	err := SzError(SzErrorCode(senzingErrorMessage), `{"messageId": 1}`)
	fmt.Println(err)
	// Output: {"messageId": 1}
}

func ExampleSzError_typeAssertion() {
	senzingErrorMessage := "33E|Test message" // Example message from Senzing G2 engine.
	err := SzError(SzErrorCode(senzingErrorMessage), `{"messageId": 1}`)
	if errors.As(err, &SzBadInputError{}) {
		fmt.Println("Is a SzBadInputError")
		if _, ok := err.(SzBadInputError).error.(SzNotFoundError); ok {
			fmt.Println("Is a SzNotFoundError")
		}
	}
	// Output:
	// Is a SzBadInputError
	// Is a SzNotFoundError
}

func ExampleIs() {
	senzingErrorMessage := "33E|Test message" // Example message from Senzing G2 engine.
	err := SzError(SzErrorCode(senzingErrorMessage), `{"messageId": 1}`)
	if err != nil {
		if Is(err, SzBadInput) {
			fmt.Println("Is a SzBadInputError")
		}
		if Is(err, SzNotFound) {
			fmt.Println("Is a SzNotFoundError")
		}
		if Is(err, SzUnrecoverable) {
			fmt.Println("Is a SzUnrecoverable")
		}
	}
	// Output:
	// Is a SzBadInputError
	// Is a SzNotFoundError
}

func ExampleIsInList() {
	senzingErrorMessage := "33E|Test message" // Example message from Senzing G2 engine.
	err := SzError(SzErrorCode(senzingErrorMessage), `{"messageId": 1}`)
	if err != nil {
		if IsInList(err, []SzErrorTypeIds{SzLicense, SzRetryTimeoutExceeded, SzBadInput}) {
			fmt.Println("Yes it is one of those listed")
		}
	}
	// Output: Yes it is one of those listed
}
