package szerror

import "errors"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type TypeIDs int

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const emptyErrorMessage = ""

/*
A collection of Senzing error types.
*/
const (
	SzBase TypeIDs = iota
	SzBadInputError
	SzConfigurationError
	SzDatabaseError
	SzDatabaseConnectionLostError
	SzLicenseError
	SzNotFoundError
	SzNotInitializedError
	SzReplaceConflictError
	SzRetryableError
	SzRetryTimeoutExceededError
	SzUnhandledError
	SzUnknownDataSourceError
	SzUnrecoverableError
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

/*
Error instances that follow the [Err prefix] naming convention.

[Err prefix]: https://github.com/mgechev/revive/blob/master/RULES_DESCRIPTIONS.md#error-naming
*/
var (
	ErrSzBadInput               = errors.New(emptyErrorMessage)
	ErrSzBase                   = errors.New(emptyErrorMessage)
	ErrSzConfiguration          = errors.New(emptyErrorMessage)
	ErrSzDatabase               = errors.New(emptyErrorMessage)
	ErrSzDatabaseConnectionLost = errors.New(emptyErrorMessage)
	ErrSzLicense                = errors.New(emptyErrorMessage)
	ErrSzNotFound               = errors.New(emptyErrorMessage)
	ErrSzNotInitialized         = errors.New(emptyErrorMessage)
	ErrSzReplaceConflict        = errors.New(emptyErrorMessage)
	ErrSzRetryable              = errors.New(emptyErrorMessage)
	ErrSzRetryTimeoutExceeded   = errors.New(emptyErrorMessage)
	ErrSzUnknownDataSource      = errors.New(emptyErrorMessage)
	ErrSzUnrecoverable          = errors.New(emptyErrorMessage)
	ErrUnhandled                = errors.New(emptyErrorMessage)
)

// A list of all TypeIDs.
var SzErrorTypesList = []TypeIDs{
	SzBadInputError,
	SzBase,
	SzConfigurationError,
	SzDatabaseError,
	SzDatabaseConnectionLostError,
	SzLicenseError,
	SzNotFoundError,
	SzNotInitializedError,
	SzReplaceConflictError,
	SzRetryableError,
	SzRetryTimeoutExceededError,
	SzUnhandledError,
	SzUnknownDataSourceError,
	SzUnrecoverableError,
}

// Map of TypeIDs to corresponding error instances.
var SzErrorMap = map[TypeIDs]error{
	SzBadInputError:               ErrSzBadInput,
	SzBase:                        ErrSzBase,
	SzConfigurationError:          ErrSzConfiguration,
	SzDatabaseError:               ErrSzDatabase,
	SzDatabaseConnectionLostError: ErrSzDatabaseConnectionLost,
	SzLicenseError:                ErrSzLicense,
	SzNotFoundError:               ErrSzNotFound,
	SzNotInitializedError:         ErrSzNotInitialized,
	SzReplaceConflictError:        ErrSzReplaceConflict,
	SzRetryableError:              ErrSzRetryable,
	SzRetryTimeoutExceededError:   ErrSzRetryTimeoutExceeded,
	SzUnhandledError:              ErrUnhandled,
	SzUnknownDataSourceError:      ErrSzUnknownDataSource,
	SzUnrecoverableError:          ErrSzUnrecoverable,
}
