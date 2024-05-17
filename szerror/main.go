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

const (
	SzBase TypeIDs = iota
	SzBadInput
	SzConfiguration
	SzDatabase
	SzDatabaseConnectionLost
	SzLicense
	SzNotFound
	SzNotInitialized
	SzRetryable
	SzRetryTimeoutExceeded
	SzUnhandled
	SzUnknownDataSource
	SzUnrecoverable
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// For "Err" prefix, see https://github.com/mgechev/revive/blob/master/RULES_DESCRIPTIONS.md#error-naming
var (
	ErrSzBadInput               = errors.New(emptyErrorMessage)
	ErrSzBase                   = errors.New(emptyErrorMessage)
	ErrSzConfiguration          = errors.New(emptyErrorMessage)
	ErrSzDatabaseConnectionLost = errors.New(emptyErrorMessage)
	ErrSzDatabase               = errors.New(emptyErrorMessage)
	ErrSzLicense                = errors.New(emptyErrorMessage)
	ErrSzNotFound               = errors.New(emptyErrorMessage)
	ErrSzNotInitialized         = errors.New(emptyErrorMessage)
	ErrSzRetryable              = errors.New(emptyErrorMessage)
	ErrSzRetryTimeoutExceeded   = errors.New(emptyErrorMessage)
	ErrSzUnknownDataSource      = errors.New(emptyErrorMessage)
	ErrSzUnrecoverable          = errors.New(emptyErrorMessage)
	ErrUnhandled                = errors.New(emptyErrorMessage)
)

// A list of all SzErrorTypeIds.
var SzErrorTypesList = []TypeIDs{
	SzBadInput,
	SzBase,
	SzConfiguration,
	SzDatabase,
	SzDatabaseConnectionLost,
	SzLicense,
	SzNotFound,
	SzNotInitialized,
	SzRetryable,
	SzRetryTimeoutExceeded,
	SzUnhandled,
	SzUnknownDataSource,
	SzUnrecoverable,
}

// Map of SzErrorTypeIds to corresponding error.
var SzErrorMap = map[TypeIDs]error{
	SzBadInput:               ErrSzBadInput,
	SzBase:                   ErrSzBase,
	SzConfiguration:          ErrSzConfiguration,
	SzRetryable:              ErrSzRetryable,
	SzUnrecoverable:          ErrSzUnrecoverable,
	SzDatabase:               ErrSzDatabase,
	SzDatabaseConnectionLost: ErrSzDatabaseConnectionLost,
	SzLicense:                ErrSzLicense,
	SzNotFound:               ErrSzNotFound,
	SzNotInitialized:         ErrSzNotInitialized,
	SzRetryTimeoutExceeded:   ErrSzRetryTimeoutExceeded,
	SzUnhandled:              ErrUnhandled,
	SzUnknownDataSource:      ErrSzUnknownDataSource,
}
