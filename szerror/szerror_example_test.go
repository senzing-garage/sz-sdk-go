package szerror

import (
	"fmt"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleMessage() {
	senzingErrorMessage := "33E|Test message" // Example message from Senzing Szengine.
	result := Message(senzingErrorMessage)
	fmt.Println(result)
	// Output: Test message
}

func ExampleCode() {
	senzingErrorMessage := "33E|Test message" // Example message from Senzing Szengine.
	result := Code(senzingErrorMessage)
	fmt.Println(result)
	// Output: 33
}

func ExampleNew() {
	senzingErrorMessage := "33E|Test message" // Example message from Senzing Szengine.
	err := New(Code(senzingErrorMessage), `{"messageId": 1}`)
	fmt.Println(err)
	// Output: {"messageId": 1}
}
