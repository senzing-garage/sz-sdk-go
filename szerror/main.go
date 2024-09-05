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
Senzing error types.
*/
const (
	SzError TypeIDs = iota
	SzBadInputError
	SzConfigurationError
	SzDatabaseConnectionLostError
	SzDatabaseError
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
	ErrSz                       = errors.New(emptyErrorMessage)
	ErrSzBadInput               = errors.New(emptyErrorMessage)
	ErrSzConfiguration          = errors.New(emptyErrorMessage)
	ErrSzDatabase               = errors.New(emptyErrorMessage)
	ErrSzDatabaseConnectionLost = errors.New(emptyErrorMessage)
	ErrSzLicense                = errors.New(emptyErrorMessage)
	ErrSzNotFound               = errors.New(emptyErrorMessage)
	ErrSzNotInitialized         = errors.New(emptyErrorMessage)
	ErrSzReplaceConflict        = errors.New(emptyErrorMessage)
	ErrSzRetryable              = errors.New(emptyErrorMessage)
	ErrSzRetryTimeoutExceeded   = errors.New(emptyErrorMessage)
	ErrSzUnhandled              = errors.New(emptyErrorMessage)
	ErrSzUnknownDataSource      = errors.New(emptyErrorMessage)
	ErrSzUnrecoverable          = errors.New(emptyErrorMessage)
)

// A list of all TypeIDs.
var SzErrorTypesList = []TypeIDs{
	SzBadInputError,
	SzConfigurationError,
	SzDatabaseConnectionLostError,
	SzDatabaseError,
	SzError,
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
	SzConfigurationError:          ErrSzConfiguration,
	SzDatabaseConnectionLostError: ErrSzDatabaseConnectionLost,
	SzDatabaseError:               ErrSzDatabase,
	SzError:                       ErrSz,
	SzLicenseError:                ErrSzLicense,
	SzNotFoundError:               ErrSzNotFound,
	SzNotInitializedError:         ErrSzNotInitialized,
	SzReplaceConflictError:        ErrSzReplaceConflict,
	SzRetryableError:              ErrSzRetryable,
	SzRetryTimeoutExceededError:   ErrSzRetryTimeoutExceeded,
	SzUnhandledError:              ErrSzUnhandled,
	SzUnknownDataSourceError:      ErrSzUnknownDataSource,
	SzUnrecoverableError:          ErrSzUnrecoverable,
}
