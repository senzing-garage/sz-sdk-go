package szerror_test

import (
	"strings"
	"testing"

	"github.com/senzing-garage/sz-sdk-go/szerror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testCases = []struct {
	expectedCode    int
	expectedMessage string
	expectedError   error
	expectedTypes   []szerror.TypeIDs
	falseTypes      []szerror.TypeIDs
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
		expectedError:   szerror.ErrSz,
		expectedTypes:   []szerror.TypeIDs{szerror.SzError},
		falseTypes:      []szerror.TypeIDs{szerror.SzUnrecoverableError},
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
		expectedError:   szerror.ErrSzBadInput,
		expectedTypes:   []szerror.TypeIDs{szerror.SzBadInputError},
		falseTypes:      []szerror.TypeIDs{szerror.SzUnrecoverableError},
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
		expectedError:   szerror.ErrSzRetryable,
		expectedTypes:   []szerror.TypeIDs{szerror.SzRetryTimeoutExceededError, szerror.SzRetryableError},
		falseTypes:      []szerror.TypeIDs{szerror.SzUnrecoverableError},
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
		expectedError:   szerror.ErrSzConfiguration,
		expectedTypes:   []szerror.TypeIDs{szerror.SzConfigurationError},
		falseTypes:      []szerror.TypeIDs{szerror.SzUnrecoverableError},
	},
	{
		name:           "szerror-0023",
		senzingMessage: "0023E|Conflicting DATA_SOURCE values 'CUSTOMERS' and 'BOB'",
		message: `{
            "errors": ["0023E|Conflicting DATA_SOURCE values 'CUSTOMERS' and 'BOB'"],
        }`,
		expectedCode:    23,
		expectedMessage: "Conflicting DATA_SOURCE values 'CUSTOMERS' and 'BOB'",
		expectedError:   szerror.ErrSzBadInput,
		expectedTypes:   []szerror.TypeIDs{szerror.SzBadInputError},
		falseTypes:      []szerror.TypeIDs{szerror.SzUnrecoverableError},
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
		expectedError:   szerror.ErrSzBadInput,
		expectedTypes:   []szerror.TypeIDs{szerror.SzNotFoundError, szerror.SzBadInputError},
		falseTypes:      []szerror.TypeIDs{szerror.SzUnrecoverableError},
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
		expectedError:   szerror.ErrSzUnrecoverable,
		expectedTypes:   []szerror.TypeIDs{szerror.SzNotInitializedError, szerror.SzUnrecoverableError},
		falseTypes:      []szerror.TypeIDs{szerror.SzBadInputError},
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
		expectedError:   szerror.ErrSzUnrecoverable,
		expectedTypes:   []szerror.TypeIDs{szerror.SzDatabaseError, szerror.SzUnrecoverableError},
		falseTypes:      []szerror.TypeIDs{szerror.SzBadInputError},
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
		expectedError:   szerror.ErrSzUnrecoverable,
		expectedTypes:   []szerror.TypeIDs{szerror.SzUnhandledError, szerror.SzUnrecoverableError},
		falseTypes:      []szerror.TypeIDs{szerror.SzBadInputError},
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
		expectedError:   szerror.ErrSzUnrecoverable,
		expectedTypes:   []szerror.TypeIDs{szerror.SzLicenseError, szerror.SzUnrecoverableError},
		falseTypes:      []szerror.TypeIDs{szerror.SzBadInputError},
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
		expectedError:   szerror.ErrSzRetryable,
		expectedTypes:   []szerror.TypeIDs{szerror.SzDatabaseConnectionLostError, szerror.SzRetryableError},
		falseTypes:      []szerror.TypeIDs{szerror.SzUnrecoverableError},
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
                            "text": "Call to SzConfigMgr_init(Test module name,
								{\"PIPELINE\":
									{\"CONFIGPATH\":\"/etc/opt/senzing\",
									\"RESOURCEPATH\":\"/opt/senzing/er/resources\",
									\"SUPPORTPATH\":\"/opt/senzing/data\"},
								\"SQL\":
									{\"CONNECTION\":
										\"postgresql://postgres:postgres@192.168.1.12:5432:G2/?sslmode=disable\"}}, 0) failed. Return code: -2",
                            "duration": 490035005,
                            "location": "In setupSenzingConfig() at szconfigmgr_test.go:183",
                            "errors": [{
                                "text": "1019E|Datastore schema tables not found. [Datastore schema tables not found.
								[(7:42P01ERROR:  relation \"sys_vars\" does not exist LINE 1:
								SELECT VAR_VALUE,SYS_LSTUPD_DT FROM SYS_VARS WHERE VAR_GROUP...^ )]]"
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
                                        "CONNECTION":
										"postgresql://postgres:postgres@192.168.1.12:5432:G2/?sslmode=disable"
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
		expectedError:   szerror.ErrSzConfiguration,
		expectedTypes:   []szerror.TypeIDs{szerror.SzConfigurationError},
		falseTypes:      []szerror.TypeIDs{szerror.SzUnrecoverableError},
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
		expectedError:   szerror.ErrSz,
		expectedTypes:   []szerror.TypeIDs{szerror.SzError},
		falseTypes:      []szerror.TypeIDs{szerror.SzUnrecoverableError},
	},
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestSzerror_SzErrorMessage(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			actual := szerror.Message(testCase.senzingMessage)
			assert.Equal(test, testCase.expectedMessage, actual, testCase.name)
		})
	}
}

func TestSzerror_SzErrorCode(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			actual := szerror.Code(testCase.senzingMessage)
			assert.Equal(test, testCase.expectedCode, actual, testCase.name)
		})
	}
}

func TestSzerror_SzError(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			actual := szerror.New(szerror.Code(testCase.senzingMessage), testCase.message)
			require.ErrorIs(test, actual, testCase.expectedError)
			assert.Equal(test, testCase.message, strings.TrimSpace(actual.Error()))
		})
	}
}

func TestSzerror_SzError_BadErrorCode(test *testing.T) {
	test.Parallel()

	err := szerror.New(999999999, "Fake message")
	require.Error(test, err)
}
