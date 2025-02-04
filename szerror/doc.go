/*
Package szerror manages types of errors issued by Senzing.

Note that type names have an "Error" suffix
and instance names have an "Err" prefix.

The following is the error type hierarchy:

	SzError
	├── SzBadInputError
	│	├── SzNotFoundError
	│	└── SzUnknownDataSourceError
	├── SzGeneralError
	│	├── SzConfigurationError
	│	├── SzSdkError
	│	└── SzReplaceConflictError
	├── SzRetryableError
	│	├── SzDatabaseConnectionLostError
	│	└── SzRetryTimeoutExceededError
	└── SzUnrecoverableError
		├── SzDatabaseError
		├── SzLicenseError
		├── SzNotInitializedError
		└── SzUnhandledError
*/
package szerror
