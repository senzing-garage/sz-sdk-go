package g2error

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestG2error_G2BaseErrorRaw(test *testing.T) {
	// nativeError := errors.New("Test message")
	// anError := G2BaseError{nativeError}
	// fmt.Printf("ErrorType: %v; Message: %s\n", reflect.TypeOf(anError), anError.Error())
	// assert.IsType(test, G2BaseError{}, anError)
}

func TestG2error_G2BadUserInputErrorRaw(test *testing.T) {
	// nativeError := errors.New("Test message")
	// baseError := G2BaseError{nativeError}
	// lowError := G2UnrecoverableError{baseError}
	// anError := G2ModuleEmptyMessageError{lowError}
	// fmt.Printf("ErrorType: %v; Message: %s\n", reflect.TypeOf(anError), anError.Error())
	// assert.True(test, errors.Is(anError, G2ModuleEmptyMessageError{}), "Not G2ModuleEmptyMessageError")
	// assert.True(test, errors.Is(anError, G2UnrecoverableError{}), "Not G2UnrecoverableError")
	// assert.True(test, errors.Is(anError, G2BaseError{}), "Not G2BaseError")
	// assert.IsType(test, G2ModuleEmptyMessageError{}, anError)
	// assert.IsType(test, G2UnrecoverableError{}, anError)
	// assert.IsType(test, G2BaseError{}, anError)

}

func TestG2error_G2BadUserInputErrorRaw2(test *testing.T) {
	// anError := G2ModuleEmptyMessageError{G2UnrecoverableError{G2BaseError{errors.New("Test message")}}}
	// testErrors := []error{
	// 	anError,
	// 	anError.G2UnrecoverableError,
	// 	anError.G2UnrecoverableError.G2BaseError,
	// 	anError.G2UnrecoverableError.G2BaseError.error,
	// }

	// for _, testError := range testErrors {
	// 	fmt.Printf("ErrorType: %v; Message: %s\n", reflect.TypeOf(testError), testError.Error())
	// }
}

func TestG2error_G2BadUserInputErrorRaw3(test *testing.T) {
	// var testError error = nil
	// testError = G2ModuleEmptyMessageError{G2UnrecoverableError{G2BaseError{errors.New("Test message")}}}

	// if myError, ok := testError.(G2ModuleEmptyMessageError); ok {
	// 	fmt.Printf("Yahoo! ErrorType: %v; Message: %s\n", reflect.TypeOf(myError), myError.Error())
	// }

	// fmt.Printf("1: ErrorType: %v; Message: %s\n", reflect.TypeOf(testError), testError.Error())

}

func TestG2error_G2BadUserInputErrorRaw4(test *testing.T) {
	// var testError error = nil
	// testError = G2BadUserInputError{G2IncompleteRecordError{errors.New("Test message")}}
	// fmt.Printf("1: ErrorType: %v; Message: %s\n", reflect.TypeOf(testError), testError.Error())

	// if testError2, ok2 := testError.(G2BadUserInputError); ok2 {
	// 	fmt.Printf("2: ErrorType: %v; Message: %s\n", reflect.TypeOf(testError2), testError2.Error())
	// 	if testError3, ok2 := testError2.error.(G2IncompleteRecordError); ok2 {
	// 		fmt.Printf("3: ErrorType: %v; Message: %s\n", reflect.TypeOf(testError3), testError3.Error())
	// 	}
	// }

	// if testError4, ok2 := testError.(G2RetryableError); ok2 {
	// 	fmt.Printf("4: ErrorType: %v; Message: %s\n", reflect.TypeOf(testError4), testError4.Error())
	// }
}

func TestG2error_G2BadUserInputError(test *testing.T) {
	anError := G2Error(99901, "Test message")
	fmt.Printf("1: ErrorType: %v; Message: %s\n", reflect.TypeOf(anError), anError.Error())

	// Find "category" error.

	if errors.As(anError, &G2BadUserInputError{}) {
		fmt.Printf("Yes it is G2BadUserInputError\n")

		// Optionally, find "detail" error.

		if detailedError, ok := anError.(G2BadUserInputError).error.(G2ModuleInvalidXMLError); ok {
			fmt.Printf("2: ErrorType: %v; Message: %s\n", reflect.TypeOf(detailedError), detailedError.Error())
			assert.IsType(test, G2ModuleInvalidXMLError{}, detailedError)
		}
	}
	assert.IsType(test, G2BadUserInputError{}, anError)
}

func TestG2error_G2BadUserInputError2(test *testing.T) {
	var anError error = nil
	anError = G2Error(99901, "Test message")
	if anError != nil {
		if errors.As(anError, &G2BadUserInputError{}) {
			fmt.Printf("Yes it is G2BadUserInputError.  So I have to correct my input.\n")
			if errors.As(anError.(G2BadUserInputError).error, &G2ModuleInvalidXMLError{}) {
				fmt.Printf("Yes it is G2ModuleInvalidXMLError.  So it's the XML that's at issue.\n")
			}
		}
	}
	assert.IsType(test, G2BadUserInputError{}, anError)
}

func TestG2error_G2BadUserInputError3(test *testing.T) {
	var anError error = nil
	anError = G2Error(99901, "Test message")
	if anError != nil {
		if Is(anError, G2BadUserInput) {
			fmt.Printf("Yes it is G2BadUserInputError.  So I have to correct my input.\n")
		}
		if Is(anError, G2ModuleInvalidXML) {
			fmt.Printf("Yes it is G2ModuleInvalidXMLError.  So it's the XML that's at issue.\n")
		}
		if Is(anError, G2ModuleEmptyMessage) {
			fmt.Printf("Yes it is G2ModuleEmptyMessageError.  This shouldn't print.\n")
		}
	}
	assert.IsType(test, G2BadUserInputError{}, anError)
}

func TestG2error_G2BadUserInputError4(test *testing.T) {
	var anError error = nil
	anError = G2Error(99901, "Test message")
	if anError != nil {
		if IsInList(anError, []G2ErrorTypeIds{G2ModuleInvalidXML, G2ModuleEmptyMessage, G2MalformedJson}) {
			fmt.Printf("Yes it is one of those listed.\n")
		}
	}
	assert.IsType(test, G2BadUserInputError{}, anError)
}

func TestG2error_Unwrap(test *testing.T) {
	var anError error = nil
	anError = G2Error(99901, "Test message")

	for anError != nil {
		fmt.Printf("ErrorType: %v; Message: %s\n", reflect.TypeOf(anError), anError.Error())
		anError = errors.Unwrap(anError)
	}
}

// func TestG2error_G2BadUserInputErrorRaw4(test *testing.T) {
// 	anError := G2ModuleEmptyMessageError{G2UnrecoverableError{G2BaseError{errors.New("Test message")}}}

// 	testErrors := []error{
// 		anError,
// 		anError.G2UnrecoverableError,
// 		anError.G2UnrecoverableError.G2BaseError,
// 		anError.G2UnrecoverableError.G2BaseError,
// 	}

// 	for _, testError := range testErrors {
// 		fmt.Printf("1: ErrorType: %v; Message: %s\n", reflect.TypeOf(testError), testError.Error())
// 	}

// 	// fmt.Println(typeOf, "has", typeOf.NumMethod(), "methods:")
// 	// for i := 0; i < typeOf.NumMethod(); i++ {
// 	// 	fmt.Print(" method#", i, ": ", typeOf.Method(i).Name, "\n")
// 	// }

// 	// fmt.Printf("ErrorType: %v; Message: %s\n", typeOf, anError.Error())
// 	assert.True(test, errors.Is(anError, G2ModuleEmptyMessageError{}), "Not G2ModuleEmptyMessageError")
// 	assert.True(test, errors.Is(anError, G2UnrecoverableError{}), "Not G2UnrecoverableError")
// 	assert.True(test, errors.Is(anError, G2BaseError{}), "Not G2BaseError")
// 	assert.IsType(test, G2ModuleEmptyMessageError{}, anError)
// 	assert.IsType(test, G2UnrecoverableError{}, anError)
// 	assert.IsType(test, G2BaseError{}, anError)

// }

// func TestG2error_G2BaseError(test *testing.T) {
// 	anError := G2Error(99900, "Test message")
// 	fmt.Printf("ErrorType: %v; Message: %s\n", reflect.TypeOf(anError), anError.Error())
// 	assert.IsType(test, G2BaseError{}, anError)
// }

// func TestG2error_T2(test *testing.T) {
// 	anError := G2BadUserInputError{errors.New("Test message")}
// 	fmt.Printf("Error: %v\n", reflect.TypeOf(anError))
// 	assert.IsType(test, G2BadUserInputError{}, anError)
// }

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

// func ExampleG2config_AddDataSource() {
// 	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2error/g2error_test.go
// 	fmt.Println("bob")
// 	// Output: bob
// }
