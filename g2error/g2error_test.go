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
	nativeError := errors.New("Test message")
	anError := G2BaseError{nativeError}
	fmt.Printf("ErrorType: %v; Message: %s\n", reflect.TypeOf(anError), anError.Error())
	assert.IsType(test, G2BaseError{}, anError)
}

func TestG2error_G2BadUserInputErrorRaw(test *testing.T) {
	nativeError := errors.New("Test message")
	anError := G2BadUserInputError{nativeError.(G2BaseError)}
	fmt.Printf("ErrorType: %v; Message: %s\n", reflect.TypeOf(anError), anError.Error())
	assert.IsType(test, G2BadUserInputError{}, anError)
}

// func TestG2error_G2BaseError(test *testing.T) {
// 	anError := G2Error(99900, "Test message")
// 	fmt.Printf("ErrorType: %v; Message: %s\n", reflect.TypeOf(anError), anError.Error())
// 	assert.IsType(test, G2BaseError{}, anError)
// }

// func TestG2error_G2BadUserInputError(test *testing.T) {
// 	anError := G2Error(99901, "Test message")
// 	fmt.Printf("Error: %v\n", reflect.TypeOf(anError))
// 	assert.IsType(test, G2BadUserInputError{}, anError)
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
