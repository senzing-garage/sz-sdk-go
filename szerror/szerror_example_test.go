package szerror_test

import (
	"fmt"

	"github.com/senzing-garage/sz-sdk-go/szerror"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleMessage() {
	senzingErrorMessage := "SENZ0031E|Test message" // Example message from Senzing Szengine.
	result := szerror.Message(senzingErrorMessage)
	fmt.Println(result)
	// Output: Test message
}

func ExampleCode() {
	senzingErrorMessage := "SENZ0032E|Test message" // Example message from Senzing Szengine.
	result := szerror.Code(senzingErrorMessage)
	fmt.Println(result)
	// Output: 32
}

func ExampleNew() {
	senzingErrorMessage := "SENZ0033E|Test message" // Example message from Senzing Szengine.
	err := szerror.New(szerror.Code(senzingErrorMessage), `{"messageId": 1}`)
	fmt.Println(err)
	// Output: {"messageId": 1}
}
