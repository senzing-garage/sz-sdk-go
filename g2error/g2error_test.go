package g2error

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	expectedCode    int
	expectedMessage string
	expectedType    error
	expectedTypes   []G2ErrorTypeIds
	falseTypes      []G2ErrorTypeIds
	message         string
	name            string
	senzingMessage  string
}{
	{
		name:           "g2error-99900",
		senzingMessage: "99900I|Test message",
		message: `{
			"errors": [
				{
					"id": "senzing-60044001",
					"text": "Not a Senzing message"
				},
				{
					"id": "senzing-60044001",
					"text": "99900I|Test message"
				}
			]
		}`,
		expectedCode:    99900,
		expectedMessage: "Test message",
		expectedType:    G2BaseError{},
		expectedTypes:   []G2ErrorTypeIds{G2Base},
		falseTypes:      []G2ErrorTypeIds{G2ModuleEmptyMessage},
	},
	{
		name:           "g2error-99901",
		senzingMessage: "99901W|Test message",
		message: `{
			"errors": [{
				"text": "99901W|Test message",
				"status": "Warning"
			}]
		}`,
		expectedCode:    99901,
		expectedMessage: "Test message",
		expectedType:    G2BadUserInputError{},
		expectedTypes:   []G2ErrorTypeIds{G2BadUserInput},
		falseTypes:      []G2ErrorTypeIds{G2ModuleEmptyMessage},
	},
	{
		name:           "g2error-99902",
		senzingMessage: "99902W|Test message",
		message: `{
			"errors": [{
				"text": "99902W|Test message"
			}]
		}`,
		expectedCode:    99902,
		expectedMessage: "Test message",
		expectedType:    G2RetryableError{},
		expectedTypes:   []G2ErrorTypeIds{G2Retryable},
		falseTypes:      []G2ErrorTypeIds{G2ModuleEmptyMessage},
	},
	{
		name:           "g2error-99903",
		senzingMessage: "99903E|Test message",
		message: `{
			"errors": [{
				"text": "99903E|Test message"
			}]
		}`,
		expectedCode:    99903,
		expectedMessage: "Test message",
		expectedType:    G2UnrecoverableError{},
		expectedTypes:   []G2ErrorTypeIds{G2Unrecoverable},
		falseTypes:      []G2ErrorTypeIds{G2ModuleEmptyMessage},
	},
	{
		name:           "g2error-99904",
		senzingMessage: "99904E|Test message",
		message: `{
			"errors": [{
				"text": "99904E|Test message"
			}]
		}`,
		expectedCode:    99904,
		expectedMessage: "Test message",
		expectedType:    G2BadUserInputError{},
		expectedTypes:   []G2ErrorTypeIds{G2BadUserInput, G2ModuleInvalidXML},
		falseTypes:      []G2ErrorTypeIds{G2ModuleEmptyMessage},
	},
	{
		name:           "g2error-99905",
		senzingMessage: "99905E|Test message",
		message: `{
			"errors": [{
				"text": "99905E|Test message"
			}]
		}`,
		expectedCode:    99905,
		expectedMessage: "Test message",
		expectedType:    G2RetryableError{},
		expectedTypes:   []G2ErrorTypeIds{G2Retryable, G2Configuration},
		falseTypes:      []G2ErrorTypeIds{G2ModuleEmptyMessage},
	},
	{
		name:           "g2error-99906",
		senzingMessage: "99906E|Test message",
		message: `{
			"errors": [{
				"text": "99906E|Test message"
			}]
		}`,
		expectedCode:    99906,
		expectedMessage: "Test message",
		expectedType:    G2UnrecoverableError{},
		expectedTypes:   []G2ErrorTypeIds{G2Unrecoverable, G2ModuleLicense},
		falseTypes:      []G2ErrorTypeIds{G2ModuleEmptyMessage},
	},
	{
		name:           "g2error-1019",
		senzingMessage: "99906E|Test message",
		message: `
		{
			"date": "2023-03-23",
			"time": "22:24:53.180659263",
			"level": "FATAL",
			"id": "senzing-60025920",
			"text": "During setup, call to setupSenzingConfig() failed.",
			"location": "In setup() at g2configmgr_test.go:244",
			"errors": [{
				"text": {
					"date": "2023-03-23",
					"time": "22:24:53.180615238",
					"level": "FATAL",
					"id": "senzing-60025912",
					"text": "During setup, call to g2configmgr.Init() failed.",
					"location": "In setupSenzingConfig() at g2configmgr_test.go:185",
					"errors": [{
						"text": {
							"date": "2023-03-23",
							"time": "22:24:53.180436236",
							"level": "ERROR",
							"id": "senzing-60024007",
							"text": "Call to G2ConfigMgr_init(Test module name, {\"PIPELINE\":{\"CONFIGPATH\":\"/etc/opt/senzing\",\"RESOURCEPATH\":\"/opt/senzing/g2/resources\",\"SUPPORTPATH\":\"/opt/senzing/data\"},\"SQL\":{\"CONNECTION\":\"postgresql://postgres:postgres@192.168.1.12:5432:G2/?sslmode=disable\"}}, 0) failed. Return code: -2",
							"duration": 490035005,
							"location": "In setupSenzingConfig() at g2configmgr_test.go:183",
							"errors": [{
								"text": "1019E|Datastore schema tables not found. [Datastore schema tables not found. [(7:42P01ERROR:  relation \"sys_vars\" does not exist LINE 1: SELECT VAR_VALUE,SYS_LSTUPD_DT FROM SYS_VARS WHERE VAR_GROUP...                                             ^ )]]"
							}],
							"details": {
								"1": "Test module name",
								"2": {
									"PIPELINE": {
										"CONFIGPATH": "/etc/opt/senzing",
										"RESOURCEPATH": "/opt/senzing/g2/resources",
										"SUPPORTPATH": "/opt/senzing/data"
									},
									"SQL": {
										"CONNECTION": "postgresql://postgres:postgres@192.168.1.12:5432:G2/?sslmode=disable"
									}
								},
								"3": 0,
								"4": -2,
								"5": 490035005
							}
						}
					}]
				}
			}]
		}`,
		expectedCode:    1019,
		expectedMessage: "Test message",
		expectedType:    G2UnrecoverableError{},
		expectedTypes:   []G2ErrorTypeIds{G2Unrecoverable, G2ModuleLicense},
		falseTypes:      []G2ErrorTypeIds{G2ModuleEmptyMessage},
	},
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestG2error_Cast(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			originalError := errors.New(testCase.message)
			desiredTypeError := G2Error(G2ErrorCode(testCase.senzingMessage), testCase.message)
			actual := Cast(originalError, desiredTypeError)
			assert.NotNil(test, actual)
			assert.IsType(test, testCase.expectedType, actual)
			assert.Equal(test, testCase.message, actual.Error())
			for _, g2ErrorTypeId := range testCase.expectedTypes {
				assert.True(test, Is(actual, g2ErrorTypeId), g2ErrorTypeId)
			}
			for _, g2ErrorTypeId := range testCase.falseTypes {
				assert.False(test, Is(actual, g2ErrorTypeId), g2ErrorTypeId)
			}
		})
	}
}

func TestG2error_Convert(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			originalError := errors.New(testCase.message)
			actual := Convert(originalError)
			assert.NotNil(test, actual)
			assert.IsType(test, testCase.expectedType, actual)
			assert.Equal(test, testCase.message, actual.Error())
			for _, g2ErrorTypeId := range testCase.expectedTypes {
				assert.True(test, Is(actual, g2ErrorTypeId), g2ErrorTypeId)
			}
			for _, g2ErrorTypeId := range testCase.falseTypes {
				assert.False(test, Is(actual, g2ErrorTypeId), g2ErrorTypeId)
			}
		})
	}
}

func TestG2error_G2ErrorMessage(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			actual := G2ErrorMessage(testCase.senzingMessage)
			assert.Equal(test, testCase.expectedMessage, actual, testCase.name)
		})
	}
}

func TestG2error_G2ErrorCode(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			actual := G2ErrorCode(testCase.senzingMessage)
			assert.Equal(test, testCase.expectedCode, actual, testCase.name)
		})
	}
}

func TestG2error_G2Error(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			actual := G2Error(G2ErrorCode(testCase.senzingMessage), testCase.message)
			assert.NotNil(test, actual)
			assert.IsType(test, testCase.expectedType, actual)
			assert.Equal(test, testCase.message, actual.Error())
		})
	}
}

func TestG2error_Is(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			actual := G2Error(G2ErrorCode(testCase.senzingMessage), testCase.message)
			assert.NotNil(test, actual)
			for _, g2ErrorTypeId := range testCase.expectedTypes {
				assert.True(test, Is(actual, g2ErrorTypeId), g2ErrorTypeId)
			}
			for _, g2ErrorTypeId := range testCase.falseTypes {
				assert.False(test, Is(actual, g2ErrorTypeId), g2ErrorTypeId)
			}
		})
	}
}

func TestG2error_IsInList(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			actual := G2Error(G2ErrorCode(testCase.senzingMessage), testCase.message)
			assert.NotNil(test, actual)
			assert.True(test, IsInList(actual, testCase.expectedTypes))
			assert.False(test, IsInList(actual, testCase.falseTypes))
		})
	}
}

func TestG2error_Unwrap(test *testing.T) {
	expectedWrapCount := 1
	actualWrapCount := 0
	err := G2Error(99901, "Test message")
	for err != nil {
		actualWrapCount += 1
		err = errors.Unwrap(err)
	}
	assert.Equal(test, expectedWrapCount, actualWrapCount)
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleCast() {
	originalError := errors.New("Original message")
	senzingErrorMessage := "99904E|Test message" // Example message from Senzing G2 engine.
	desiredTypeError := G2Error(G2ErrorCode(senzingErrorMessage), `{"messageId": 1}`)
	err := Cast(originalError, desiredTypeError)
	fmt.Println(err)
	// Output: Original message
}

func ExampleG2ErrorMessage() {
	senzingErrorMessage := "99904E|Test message" // Example message from Senzing G2 engine.
	result := G2ErrorMessage(senzingErrorMessage)
	fmt.Println(result)
	// Output: Test message
}

func ExampleG2ErrorCode() {
	senzingErrorMessage := "99904E|Test message" // Example message from Senzing G2 engine.
	result := G2ErrorCode(senzingErrorMessage)
	fmt.Println(result)
	// Output: 99904
}

func ExampleG2Error() {
	senzingErrorMessage := "99904E|Test message" // Example message from Senzing G2 engine.
	err := G2Error(G2ErrorCode(senzingErrorMessage), `{"messageId": 1}`)
	fmt.Println(err)
	// Output: {"messageId": 1}
}

func ExampleG2Error_typeAssertion() {
	senzingErrorMessage := "99904E|Test message" // Example message from Senzing G2 engine.
	err := G2Error(G2ErrorCode(senzingErrorMessage), `{"messageId": 1}`)
	if errors.As(err, &G2BadUserInputError{}) {
		fmt.Println("Is a G2BadUserInputError")
		if _, ok := err.(G2BadUserInputError).error.(G2ModuleInvalidXMLError); ok {
			fmt.Println("Is a G2ModuleInvalidXMLError")
		}
	}
	// Output:
	// Is a G2BadUserInputError
	// Is a G2ModuleInvalidXMLError
}

func ExampleIs() {
	senzingErrorMessage := "99904E|Test message" // Example message from Senzing G2 engine.
	err := G2Error(G2ErrorCode(senzingErrorMessage), `{"messageId": 1}`)
	if err != nil {
		if Is(err, G2BadUserInput) {
			fmt.Println("Is a G2BadUserInputError")
		}
		if Is(err, G2ModuleInvalidXML) {
			fmt.Println("Is a G2ModuleInvalidXMLError")
		}
		if Is(err, G2ModuleEmptyMessage) {
			fmt.Println("Is a G2ModuleEmptyMessageError.")
		}
	}
	// Output:
	// Is a G2BadUserInputError
	// Is a G2ModuleInvalidXMLError
}

func ExampleIsInList() {
	senzingErrorMessage := "99904E|Test message" // Example message from Senzing G2 engine.
	err := G2Error(G2ErrorCode(senzingErrorMessage), `{"messageId": 1}`)
	if err != nil {
		if IsInList(err, []G2ErrorTypeIds{G2ModuleInvalidXML, G2ModuleEmptyMessage, G2MalformedJson}) {
			fmt.Println("Yes it is one of those listed")
		}
	}
	// Output: Yes it is one of those listed
}
