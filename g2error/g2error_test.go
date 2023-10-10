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
		falseTypes:      []G2ErrorTypeIds{G2Unrecoverable},
	},
	{
		name:           "g2error-99910",
		senzingMessage: "99910W|Test message",
		message: `{
			"errors": [{
				"text": "99910W|Test message",
				"status": "Warning"
			}]
		}`,
		expectedCode:    99910,
		expectedMessage: "Test message",
		expectedType:    G2BadInputError{},
		expectedTypes:   []G2ErrorTypeIds{G2BadInput},
		falseTypes:      []G2ErrorTypeIds{G2Unrecoverable},
	},
	{
		name:           "g2error-99911",
		senzingMessage: "99911E|Test message",
		message: `{
			"errors": [{
				"text": "99911E|Test message"
			}]
		}`,
		expectedCode:    99911,
		expectedMessage: "Test message",
		expectedType:    G2BadInputError{},
		expectedTypes:   []G2ErrorTypeIds{G2NotFound, G2BadInput},
		falseTypes:      []G2ErrorTypeIds{G2Unrecoverable},
	},
	{
		name:           "g2error-99912",
		senzingMessage: "99912E|Test message",
		message: `{
			"errors": [{
				"text": "99912E|Test message"
			}]
		}`,
		expectedCode:    99912,
		expectedMessage: "Test message",
		expectedType:    G2BadInputError{},
		expectedTypes:   []G2ErrorTypeIds{G2UnknownDatasource, G2BadInput},
		falseTypes:      []G2ErrorTypeIds{G2Unrecoverable},
	},
	{
		name:           "g2error-99920",
		senzingMessage: "99920W|Test message",
		message: `{
			"errors": [{
				"text": "99920W|Test message"
			}]
		}`,
		expectedCode:    99920,
		expectedMessage: "Test message",
		expectedType:    G2ConfigurationError{},
		expectedTypes:   []G2ErrorTypeIds{G2Configuration},
		falseTypes:      []G2ErrorTypeIds{G2Unrecoverable},
	},
	{
		name:           "g2error-99930",
		senzingMessage: "99930E|Test message",
		message: `{
			"errors": [{
				"text": "99930E|Test message"
			}]
		}`,
		expectedCode:    99930,
		expectedMessage: "Test message",
		expectedType:    G2RetryableError{},
		expectedTypes:   []G2ErrorTypeIds{G2Retryable},
		falseTypes:      []G2ErrorTypeIds{G2Unrecoverable},
	},
	{
		name:           "g2error-99931",
		senzingMessage: "99931E|Test message",
		message: `{
			"errors": [{
				"text": "99931E|Test message"
			}]
		}`,
		expectedCode:    99931,
		expectedMessage: "Test message",
		expectedType:    G2RetryableError{},
		expectedTypes:   []G2ErrorTypeIds{G2DatabaseConnectionLost, G2Retryable},
		falseTypes:      []G2ErrorTypeIds{G2Unrecoverable},
	},
	{
		name:           "g2error-99932",
		senzingMessage: "99932E|Test message",
		message: `{
			"errors": [{
				"text": "99932E|Test message"
			}]
		}`,
		expectedCode:    99932,
		expectedMessage: "Test message",
		expectedType:    G2RetryableError{},
		expectedTypes:   []G2ErrorTypeIds{G2RetryTimeoutExceeded, G2Retryable},
		falseTypes:      []G2ErrorTypeIds{G2Unrecoverable},
	},
	{
		name:           "g2error-99940",
		senzingMessage: "99940E|Test message",
		message: `{
			"errors": [{
				"text": "99940E|Test message"
			}]
		}`,
		expectedCode:    99940,
		expectedMessage: "Test message",
		expectedType:    G2UnrecoverableError{},
		expectedTypes:   []G2ErrorTypeIds{G2Unrecoverable},
		falseTypes:      []G2ErrorTypeIds{G2BadInput},
	},
	{
		name:           "g2error-99941",
		senzingMessage: "99941E|Test message",
		message: `{
			"errors": [{
				"text": "99941E|Test message"
			}]
		}`,
		expectedCode:    99941,
		expectedMessage: "Test message",
		expectedType:    G2UnrecoverableError{},
		expectedTypes:   []G2ErrorTypeIds{G2Database, G2Unrecoverable},
		falseTypes:      []G2ErrorTypeIds{G2BadInput},
	},
	{
		name:           "g2error-99942",
		senzingMessage: "99942E|Test message",
		message: `{
			"errors": [{
				"text": "99942E|Test message"
			}]
		}`,
		expectedCode:    99942,
		expectedMessage: "Test message",
		expectedType:    G2UnrecoverableError{},
		expectedTypes:   []G2ErrorTypeIds{G2License, G2Unrecoverable},
		falseTypes:      []G2ErrorTypeIds{G2BadInput},
	},
	{
		name:           "g2error-99943",
		senzingMessage: "99943E|Test message",
		message: `{
			"errors": [{
				"text": "99943E|Test message"
			}]
		}`,
		expectedCode:    99943,
		expectedMessage: "Test message",
		expectedType:    G2UnrecoverableError{},
		expectedTypes:   []G2ErrorTypeIds{G2NotInitialized, G2Unrecoverable},
		falseTypes:      []G2ErrorTypeIds{G2BadInput},
	},
	{
		name:           "g2error-99944",
		senzingMessage: "99944E|Test message",
		message: `{
			"errors": [{
				"text": "99944E|Test message"
			}]
		}`,
		expectedCode:    99944,
		expectedMessage: "Test message",
		expectedType:    G2UnrecoverableError{},
		expectedTypes:   []G2ErrorTypeIds{G2Unhandled, G2Unrecoverable},
		falseTypes:      []G2ErrorTypeIds{G2BadInput},
	},

	// ------------------------------------------------------------------------

	{
		name:           "g2error-0023",
		senzingMessage: "0023E|Conflicting DATA_SOURCE values 'CUSTOMERS' and 'BOB'",
		message: `{
			"errors": ["0023E|Conflicting DATA_SOURCE values 'CUSTOMERS' and 'BOB'"],
		}`,
		expectedCode:    23,
		expectedMessage: "Conflicting DATA_SOURCE values 'CUSTOMERS' and 'BOB'",
		expectedType:    G2BadInputError{},
		expectedTypes:   []G2ErrorTypeIds{G2BadInput},
		falseTypes:      []G2ErrorTypeIds{G2Unrecoverable},
	},

	{
		name:           "g2error-1019",
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
		expectedType:    G2ConfigurationError{},
		expectedTypes:   []G2ErrorTypeIds{G2Configuration},
		falseTypes:      []G2ErrorTypeIds{G2Unrecoverable},
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
				assert.True(test, Is(actual, g2ErrorTypeId), fmt.Sprintf("%d should be %d", g2ErrorTypeId, actual))
			}
			for _, g2ErrorTypeId := range testCase.falseTypes {
				assert.False(test, Is(actual, g2ErrorTypeId), g2ErrorTypeId)
			}
		})
	}
}

func TestG2error_Cast_nil(test *testing.T) {
	actual := Convert(nil)
	assert.Nil(test, actual)

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			desiredTypeError := G2Error(G2ErrorCode(testCase.senzingMessage), testCase.message)
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

func TestG2error_Convert_nil(test *testing.T) {
	actual := Convert(nil)
	assert.Nil(test, actual)
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
