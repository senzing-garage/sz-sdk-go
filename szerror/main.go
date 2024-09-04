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
	SzBadInput
	SzConfiguration
	SzDatabase
	SzDatabaseConnectionLost
	SzLicense
	SzNotFound
	SzNotInitialized
	SzReplaceConflict
	SzRetryable
	SzRetryTimeoutExceeded
	SzUnhandled
	SzUnknownDataSource
	SzUnrecoverable
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
	SzBadInput,
	SzBase,
	SzConfiguration,
	SzDatabase,
	SzDatabaseConnectionLost,
	SzLicense,
	SzNotFound,
	SzNotInitialized,
	SzReplaceConflict,
	SzRetryable,
	SzRetryTimeoutExceeded,
	SzUnhandled,
	SzUnknownDataSource,
	SzUnrecoverable,
}

// Map of TypeIDs to corresponding error instances.
var SzErrorMap = map[TypeIDs]error{
	SzBadInput:               ErrSzBadInput,
	SzBase:                   ErrSzBase,
	SzConfiguration:          ErrSzConfiguration,
	SzDatabase:               ErrSzDatabase,
	SzDatabaseConnectionLost: ErrSzDatabaseConnectionLost,
	SzLicense:                ErrSzLicense,
	SzNotFound:               ErrSzNotFound,
	SzNotInitialized:         ErrSzNotInitialized,
	SzReplaceConflict:        ErrSzReplaceConflict,
	SzRetryable:              ErrSzRetryable,
	SzRetryTimeoutExceeded:   ErrSzRetryTimeoutExceeded,
	SzUnhandled:              ErrUnhandled,
	SzUnknownDataSource:      ErrSzUnknownDataSource,
	SzUnrecoverable:          ErrSzUnrecoverable,
}
