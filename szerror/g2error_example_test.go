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
	senzingErrorMessage := "99911E|Test message" // Example message from Senzing G2 engine.
	desiredTypeError := G2Error(G2ErrorCode(senzingErrorMessage), `{"messageId": 1}`)
	err := Cast(originalError, desiredTypeError)
	fmt.Printf("Error type: %s; Error message: %s", reflect.TypeOf(err), err.Error())
	// Output: Error type: szerror.G2BadInputError; Error message: Original message
}

func ExampleG2ErrorMessage() {
	senzingErrorMessage := "99911E|Test message" // Example message from Senzing G2 engine.
	result := G2ErrorMessage(senzingErrorMessage)
	fmt.Println(result)
	// Output: Test message
}

func ExampleG2ErrorCode() {
	senzingErrorMessage := "99911E|Test message" // Example message from Senzing G2 engine.
	result := G2ErrorCode(senzingErrorMessage)
	fmt.Println(result)
	// Output: 99911
}

func ExampleG2Error() {
	senzingErrorMessage := "99911E|Test message" // Example message from Senzing G2 engine.
	err := G2Error(G2ErrorCode(senzingErrorMessage), `{"messageId": 1}`)
	fmt.Println(err)
	// Output: {"messageId": 1}
}

func ExampleG2Error_typeAssertion() {
	senzingErrorMessage := "99911E|Test message" // Example message from Senzing G2 engine.
	err := G2Error(G2ErrorCode(senzingErrorMessage), `{"messageId": 1}`)
	if errors.As(err, &G2BadInputError{}) {
		fmt.Println("Is a G2BadInputError")
		if _, ok := err.(G2BadInputError).error.(G2NotFoundError); ok {
			fmt.Println("Is a G2NotFoundError")
		}
	}
	// Output:
	// Is a G2BadInputError
	// Is a G2NotFoundError
}

func ExampleIs() {
	senzingErrorMessage := "99911E|Test message" // Example message from Senzing G2 engine.
	err := G2Error(G2ErrorCode(senzingErrorMessage), `{"messageId": 1}`)
	if err != nil {
		if Is(err, G2BadInput) {
			fmt.Println("Is a G2BadInputError")
		}
		if Is(err, G2NotFound) {
			fmt.Println("Is a G2NotFoundError")
		}
		if Is(err, G2Unrecoverable) {
			fmt.Println("Is a G2Unrecoverable")
		}
	}
	// Output:
	// Is a G2BadInputError
	// Is a G2NotFoundError
}

func ExampleIsInList() {
	senzingErrorMessage := "99911E|Test message" // Example message from Senzing G2 engine.
	err := G2Error(G2ErrorCode(senzingErrorMessage), `{"messageId": 1}`)
	if err != nil {
		if IsInList(err, []G2ErrorTypeIds{G2License, G2RetryTimeoutExceeded, G2BadInput}) {
			fmt.Println("Yes it is one of those listed")
		}
	}
	// Output: Yes it is one of those listed
}
