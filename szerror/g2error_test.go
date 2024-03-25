package szerror

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
	expectedTypes   []SzErrorTypeIds
	falseTypes      []SzErrorTypeIds
	message         string
	name            string
	senzingMessage  string
}{
	{
		name:           "szerror-0005",
		senzingMessage: "5I|Test message",
		message: `{
			"errors": [
				{
					"id": "senzing-60044001",
					"text": "Not a Senzing message"
				},
				{
					"id": "senzing-60044001",
					"text": "5I|Test message"
				}
			]
		}`,
		expectedCode:    5,
		expectedMessage: "Test message",
		expectedType:    SzBaseError{},
		expectedTypes:   []SzErrorTypeIds{SzBase},
		falseTypes:      []SzErrorTypeIds{SzUnrecoverable},
	},
	{
		name:           "szerror-0007",
		senzingMessage: "7W|Test message",
		message: `{
			"errors": [{
				"text": "7W|Test message",
				"status": "Warning"
			}]
		}`,
		expectedCode:    7,
		expectedMessage: "Test message",
		expectedType:    SzBadInputError{},
		expectedTypes:   []SzErrorTypeIds{SzBadInput},
		falseTypes:      []SzErrorTypeIds{SzUnrecoverable},
	},
	{
		name:           "szerror-0010",
		senzingMessage: "10E|Test message",
		message: `{
			"errors": [{
				"text": "10E|Test message"
			}]
		}`,
		expectedCode:    10,
		expectedMessage: "Test message",
		expectedType:    SzRetryableError{},
		expectedTypes:   []SzErrorTypeIds{SzRetryTimeoutExceeded, SzRetryable},
		falseTypes:      []SzErrorTypeIds{SzUnrecoverable},
	},
	{
		name:           "szerror-0014",
		senzingMessage: "14W|Test message",
		message: `{
			"errors": [{
				"text": "14W|Test message"
			}]
		}`,
		expectedCode:    14,
		expectedMessage: "Test message",
		expectedType:    SzConfigurationError{},
		expectedTypes:   []SzErrorTypeIds{SzConfiguration},
		falseTypes:      []SzErrorTypeIds{SzUnrecoverable},
	},
	{
		name:           "szerror-0023",
		senzingMessage: "0023E|Conflicting DATA_SOURCE values 'CUSTOMERS' and 'BOB'",
		message: `{
			"errors": ["0023E|Conflicting DATA_SOURCE values 'CUSTOMERS' and 'BOB'"],
		}`,
		expectedCode:    23,
		expectedMessage: "Conflicting DATA_SOURCE values 'CUSTOMERS' and 'BOB'",
		expectedType:    SzBadInputError{},
		expectedTypes:   []SzErrorTypeIds{SzBadInput},
		falseTypes:      []SzErrorTypeIds{SzUnrecoverable},
	},
	{
		name:           "szerror-0027",
		senzingMessage: "27E|Test message",
		message: `{
			"errors": [{
				"text": "27E|Test message"
			}]
		}`,
		expectedCode:    27,
		expectedMessage: "Test message",
		expectedType:    SzBadInputError{},
		expectedTypes:   []SzErrorTypeIds{SzUnknownDatasource, SzBadInput},
		falseTypes:      []SzErrorTypeIds{SzUnrecoverable},
	},
	{
		name:           "szerror-0033",
		senzingMessage: "33E|Test message",
		message: `{
			"errors": [{
				"text": "33E|Test message"
			}]
		}`,
		expectedCode:    33,
		expectedMessage: "Test message",
		expectedType:    SzBadInputError{},
		expectedTypes:   []SzErrorTypeIds{SzNotFound, SzBadInput},
		falseTypes:      []SzErrorTypeIds{SzUnrecoverable},
	},
	{
		name:           "szerror-0048",
		senzingMessage: "48E|Test message",
		message: `{
			"errors": [{
				"text": "48E|Test message"
			}]
		}`,
		expectedCode:    48,
		expectedMessage: "Test message",
		expectedType:    SzUnrecoverableError{},
		expectedTypes:   []SzErrorTypeIds{SzNotInitialized, SzUnrecoverable},
		falseTypes:      []SzErrorTypeIds{SzBadInput},
	},
	{
		name:           "szerror-0054",
		senzingMessage: "54E|Test message",
		message: `{
			"errors": [{
				"text": "54E|Test message"
			}]
		}`,
		expectedCode:    54,
		expectedMessage: "Test message",
		expectedType:    SzUnrecoverableError{},
		expectedTypes:   []SzErrorTypeIds{SzDatabase, SzUnrecoverable},
		falseTypes:      []SzErrorTypeIds{SzBadInput},
	},
	{
		name:           "szerror-00087",
		senzingMessage: "87E|Test message",
		message: `{
			"errors": [{
				"text": "87E|Test message"
			}]
		}`,
		expectedCode:    87,
		expectedMessage: "Test message",
		expectedType:    SzUnrecoverableError{},
		expectedTypes:   []SzErrorTypeIds{SzUnhandled, SzUnrecoverable},
		falseTypes:      []SzErrorTypeIds{SzBadInput},
	},
	{
		name:           "szerror-0999",
		senzingMessage: "999E|Test message",
		message: `{
			"errors": [{
				"text": "999E|Test message"
			}]
		}`,
		expectedCode:    999,
		expectedMessage: "Test message",
		expectedType:    SzUnrecoverableError{},
		expectedTypes:   []SzErrorTypeIds{SzLicense, SzUnrecoverable},
		falseTypes:      []SzErrorTypeIds{SzBadInput},
	},
	{
		name:           "szerror-1006",
		senzingMessage: "1006E|Test message",
		message: `{
			"errors": [{
				"text": "1006E|Test message"
			}]
		}`,
		expectedCode:    1006,
		expectedMessage: "Test message",
		expectedType:    SzRetryableError{},
		expectedTypes:   []SzErrorTypeIds{SzDatabaseConnectionLost, SzRetryable},
		falseTypes:      []SzErrorTypeIds{SzUnrecoverable},
	},
	{
		name:           "szerror-1019",
		senzingMessage: "1019E|Test message",
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
		expectedType:    SzConfigurationError{},
		expectedTypes:   []SzErrorTypeIds{SzConfiguration},
		falseTypes:      []SzErrorTypeIds{SzUnrecoverable},
	},
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestSzerror_Cast(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			originalError := errors.New(testCase.message)
			desiredTypeError := SzError(SzErrorCode(testCase.senzingMessage), testCase.message)
			actual := Cast(originalError, desiredTypeError)
			assert.NotNil(test, actual)
			assert.IsType(test, testCase.expectedType, actual)
			assert.Equal(test, testCase.message, actual.Error())
			for _, g2ErrorTypeId := range testCase.expectedTypes {
				assert.True(test, Is(actual, g2ErrorTypeId), fmt.Sprintf("%d should be %d", g2ErrorTypeId, actual))
			}
			for _, g2ErrorTypeId := range testCase.falseTypes {
				assert.False(test, Is(actual, g2ErrorTypeId), g2ErrorTypeId)
			}
		})
	}
}

func TestSzerror_Cast_nil(test *testing.T) {
	actual := Convert(nil)
	assert.Nil(test, actual)

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			desiredTypeError := SzError(SzErrorCode(testCase.senzingMessage), testCase.message)
			actual := Cast(nil, desiredTypeError)
			assert.Nil(test, actual, "Nil actual")
		})
	}
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			originalError := errors.New(testCase.message)
			actual := Cast(originalError, nil)
			assert.NotNil(test, actual, "Nil desired type")
		})
	}
}

func TestSzerror_Convert(test *testing.T) {
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

func TestSzerror_Convert_nil(test *testing.T) {
	actual := Convert(nil)
	assert.Nil(test, actual)
}

func TestSzerror_SzErrorMessage(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			actual := SzErrorMessage(testCase.senzingMessage)
			assert.Equal(test, testCase.expectedMessage, actual, testCase.name)
		})
	}
}

func TestSzerror_SzErrorCode(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			actual := SzErrorCode(testCase.senzingMessage)
			assert.Equal(test, testCase.expectedCode, actual, testCase.name)
		})
	}
}

func TestSzerror_SzError(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			actual := SzError(SzErrorCode(testCase.senzingMessage), testCase.message)
			assert.NotNil(test, actual)
			assert.IsType(test, testCase.expectedType, actual)
			assert.Equal(test, testCase.message, actual.Error())
		})
	}
}

func TestSzerror_Is(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			actual := SzError(SzErrorCode(testCase.senzingMessage), testCase.message)
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

func TestSzerror_IsInList(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			actual := SzError(SzErrorCode(testCase.senzingMessage), testCase.message)
			assert.NotNil(test, actual)
			assert.True(test, IsInList(actual, testCase.expectedTypes))
			assert.False(test, IsInList(actual, testCase.falseTypes))
		})
	}
}

func TestSzerror_Unwrap(test *testing.T) {
	expectedWrapCount := 1
	actualWrapCount := 0
	err := SzError(99901, "Test message")
	for err != nil {
		actualWrapCount += 1
		err = errors.Unwrap(err)
	}
	assert.Equal(test, expectedWrapCount, actualWrapCount)
}
