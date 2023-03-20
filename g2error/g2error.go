package g2error

// ----------------------------------------------------------------------------
// Base error
// ----------------------------------------------------------------------------

type G2Error struct {
	error
	ForExampleInt    int
	ForExampleString string
}

// ----------------------------------------------------------------------------
// "Category" errors - all based directly on G2Error
// ----------------------------------------------------------------------------

type G2UnrecoverableInputError struct {
	G2Error
}

type G2RetryableError struct {
	G2Error
}

type G2BadUserInputError struct {
	G2Error
}

// ----------------------------------------------------------------------------
// Errors based on G2BadUserInputError
// ----------------------------------------------------------------------------

type G2IncompleteRecordError struct {
	G2BadUserInputError
}

type G2MalformedJsonError struct {
	G2BadUserInputError
}

type G2MissingConfigurationException struct {
	G2BadUserInputError
}

type G2MissingDataSourceError struct {
	G2BadUserInputError
}

type G2NotFoundError struct {
	G2BadUserInputError
}

type G2UnacceptableJsonKeyValueError struct {
	G2BadUserInputError
}

// ----------------------------------------------------------------------------
// Errors based on G2RetryableError
// ----------------------------------------------------------------------------

type G2ConfigurationError struct {
	G2RetryableError
}
type G2DatabaseConnectionLostError struct {
	G2RetryableError
}
type G2MessageBufferError struct {
	G2RetryableError
}
type G2RepositoryPurgedError struct {
	G2RetryableError
}
type G2RetryTimeoutExceededError struct {
	G2RetryableError
}

// ----------------------------------------------------------------------------
// Errors based on G2UnrecoverableInputError
// ----------------------------------------------------------------------------

type G2DatabaseError struct {
	G2UnrecoverableInputError
}

type G2ModuleEmptyMessageError struct {
	G2UnrecoverableInputError
}

type G2ModuleError struct {
	G2UnrecoverableInputError
}

type G2ModuleGenericError struct {
	G2UnrecoverableInputError
}

type G2ModuleInvalidXMLError struct {
	G2UnrecoverableInputError
}

type G2ModuleLicenseError struct {
	G2UnrecoverableInputError
}

type G2ModuleNotInitializedError struct {
	G2UnrecoverableInputError
}

type G2ModuleResolveMissingResEntError struct {
	G2UnrecoverableInputError
}
type G2UnhandledError struct {
	G2UnrecoverableInputError
}
