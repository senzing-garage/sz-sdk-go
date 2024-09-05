package szerror

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testCases = []struct {
	expectedCode    int
	expectedMessage string
	expectedError   error
	expectedTypes   []TypeIDs
	falseTypes      []TypeIDs
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
		expectedError:   ErrSz,
		expectedTypes:   []TypeIDs{SzError},
		falseTypes:      []TypeIDs{SzUnrecoverableError},
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
		expectedError:   ErrSzBadInput,
		expectedTypes:   []TypeIDs{SzBadInputError},
		falseTypes:      []TypeIDs{SzUnrecoverableError},
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
		expectedError:   ErrSzRetryable,
		expectedTypes:   []TypeIDs{SzRetryTimeoutExceededError, SzRetryableError},
		falseTypes:      []TypeIDs{SzUnrecoverableError},
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
		expectedError:   ErrSzConfiguration,
		expectedTypes:   []TypeIDs{SzConfigurationError},
		falseTypes:      []TypeIDs{SzUnrecoverableError},
	},
	{
		name:           "szerror-0023",
		senzingMessage: "0023E|Conflicting DATA_SOURCE values 'CUSTOMERS' and 'BOB'",
		message: `{
            "errors": ["0023E|Conflicting DATA_SOURCE values 'CUSTOMERS' and 'BOB'"],
        }`,
		expectedCode:    23,
		expectedMessage: "Conflicting DATA_SOURCE values 'CUSTOMERS' and 'BOB'",
		expectedError:   ErrSzBadInput,
		expectedTypes:   []TypeIDs{SzBadInputError},
		falseTypes:      []TypeIDs{SzUnrecoverableError},
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
		expectedError:   ErrSzBadInput,
		expectedTypes:   []TypeIDs{SzNotFoundError, SzBadInputError},
		falseTypes:      []TypeIDs{SzUnrecoverableError},
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
		expectedError:   ErrSzUnrecoverable,
		expectedTypes:   []TypeIDs{SzNotInitializedError, SzUnrecoverableError},
		falseTypes:      []TypeIDs{SzBadInputError},
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
		expectedError:   ErrSzUnrecoverable,
		expectedTypes:   []TypeIDs{SzDatabaseError, SzUnrecoverableError},
		falseTypes:      []TypeIDs{SzBadInputError},
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
		expectedError:   ErrSzUnrecoverable,
		expectedTypes:   []TypeIDs{SzUnhandledError, SzUnrecoverableError},
		falseTypes:      []TypeIDs{SzBadInputError},
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
		expectedError:   ErrSzUnrecoverable,
		expectedTypes:   []TypeIDs{SzLicenseError, SzUnrecoverableError},
		falseTypes:      []TypeIDs{SzBadInputError},
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
		expectedError:   ErrSzRetryable,
		expectedTypes:   []TypeIDs{SzDatabaseConnectionLostError, SzRetryableError},
		falseTypes:      []TypeIDs{SzUnrecoverableError},
	},
	{
		name:           "szerror-1019",
		senzingMessage: "1019E|Test message",
		message: `{
            "date": "2023-03-23",
            "time": "22:24:53.180659263",
            "level": "FATAL",
            "id": "senzing-60025920",
            "text": "During setup, call to setupSenzingConfig() failed.",
            "location": "In setup() at szconfigmgr_test.go:244",
            "errors": [{
                "text": {
                    "date": "2023-03-23",
                    "time": "22:24:53.180615238",
                    "level": "FATAL",
                    "id": "senzing-60025912",
                    "text": "During setup, call to asconfigmgr.Init() failed.",
                    "location": "In setupSenzingConfig() at szconfigmgr_test.go:185",
                    "errors": [{
                        "text": {
                            "date": "2023-03-23",
                            "time": "22:24:53.180436236",
                            "level": "ERROR",
                            "id": "senzing-60024007",
                            "text": "Call to SzConfigMgr_init(Test module name, {\"PIPELINE\":{\"CONFIGPATH\":\"/etc/opt/senzing\",\"RESOURCEPATH\":\"/opt/senzing/er/resources\",\"SUPPORTPATH\":\"/opt/senzing/data\"},\"SQL\":{\"CONNECTION\":\"postgresql://postgres:postgres@192.168.1.12:5432:G2/?sslmode=disable\"}}, 0) failed. Return code: -2",
                            "duration": 490035005,
                            "location": "In setupSenzingConfig() at szconfigmgr_test.go:183",
                            "errors": [{
                                "text": "1019E|Datastore schema tables not found. [Datastore schema tables not found. [(7:42P01ERROR:  relation \"sys_vars\" does not exist LINE 1: SELECT VAR_VALUE,SYS_LSTUPD_DT FROM SYS_VARS WHERE VAR_GROUP...                                             ^ )]]"
                            }],
                            "details": {
                                "1": "Test module name",
                                "2": {
                                    "PIPELINE": {
                                        "CONFIGPATH": "/etc/opt/senzing",
                                        "RESOURCEPATH": "/opt/senzing/er/resources",
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
		expectedError:   ErrSzConfiguration,
		expectedTypes:   []TypeIDs{SzConfigurationError},
		falseTypes:      []TypeIDs{SzUnrecoverableError},
	},
	{
		name:           "szerror-NoCode",
		senzingMessage: "Test message",
		message: `{
            "errors": [{
                "text": "Test message"
            }]
        }`,
		expectedCode:    0,
		expectedMessage: "",
		expectedError:   ErrSz,
		expectedTypes:   []TypeIDs{SzError},
		falseTypes:      []TypeIDs{SzUnrecoverableError},
	},
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestSzerror_SzErrorMessage(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			actual := Message(testCase.senzingMessage)
			assert.Equal(test, testCase.expectedMessage, actual, testCase.name)
		})
	}
}

func TestSzerror_SzErrorCode(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			actual := Code(testCase.senzingMessage)
			assert.Equal(test, testCase.expectedCode, actual, testCase.name)
		})
	}
}

func TestSzerror_SzError(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			actual := New(Code(testCase.senzingMessage), testCase.message)
			require.ErrorIs(test, actual, testCase.expectedError)
			assert.Equal(test, testCase.message, strings.TrimSpace(actual.Error()))
		})
	}
}

// ----------------------------------------------------------------------------
// Test private functions
// ----------------------------------------------------------------------------

func TestSzerror_mapErrorIDtoError(test *testing.T) {
	err := mapErrorIDtoError(9999)
	require.ErrorIs(test, err, ErrSz)
}
