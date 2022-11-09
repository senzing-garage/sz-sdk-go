/*
Package g2product is a Go wrapper over Senzing's G2product C binding.

To use G2product, the LD_LIBRARY_PATH environment variable must include
a path to Senzing's libraries.  Example:

	export LD_LIBRARY_PATH=/opt/senzing/g2/lib
*/
package g2product

import (
	"context"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2product interface {
	ClearLastException(ctx context.Context) error
	Destroy(ctx context.Context) error
	GetLastException(ctx context.Context) (string, error)
	GetLastExceptionCode(ctx context.Context) (int, error)
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error
	License(ctx context.Context) (string, error)
	SetLogLevel(ctx context.Context, logLevel logger.Level) error
	ValidateLicenseFile(ctx context.Context, licenseFilePath string) (string, error)
	ValidateLicenseStringBase64(ctx context.Context, licenseString string) (string, error)
	Version(ctx context.Context) (string, error)
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdTemplate = "senzing-6006%04d"

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var IdMessages = map[int]string{
	2001: "Call to G2Product_destroy() failed. Return code: %d",
	2002: "Call to G2Product_getLastException() failed. Return code: %d",
	2003: "Call to G2Product_init(%s, %s, %s) failed. Return code: %d",
	2004: "Call to G2Product_validateLicenseFile(%s) failed. Return code: %d",
	2005: "Call to G2Product_validateLicenseStringBase64(%s) failed. Return code: %d",
	4001: "Enter ClearLastException().",
	4002: "Exit  ClearLastException() returned (%v).",
	4003: "Enter Destroy().",
	4004: "Exit  Destroy() returned (%v).",
	4005: "Enter GetLastException().",
	4006: "Exit  GetLastException() returned (%s, %v).",
	4007: "Enter GetLastExceptionCode().",
	4008: "Exit  GetLastExceptionCode() returned (%d, %v).",
	4009: "Enter Init(%s, %s, %s).",
	4010: "Exit  Init(%s, %s, %s) returned (%v).",
	4011: "Enter License().",
	4012: "Exit  License() returned (%s, %v).",
	4013: "Enter SetLogLevel(%v).",
	4014: "Exit  SetLogLevel(%v) returned (%v).",
	4015: "Enter ValidateLicenseFile(%s).",
	4016: "Exit  ValidateLicenseFile(%s) returned (%s, %v).",
	4017: "Enter ValidateLicenseStringBase64(%s).",
	4018: "Exit  ValidateLicenseStringBase64(%s) returned (%s, %v).",
	4019: "Enter Version().",
	4020: "Exit  Version() returned (%s, %v).",
}

var IdRanges = map[int]string{
	0000: logger.LevelInfoName,
	1000: logger.LevelWarnName,
	2000: logger.LevelErrorName,
	3000: logger.LevelDebugName,
	4000: logger.LevelTraceName,
	5000: logger.LevelFatalName,
	6000: logger.LevelPanicName,
}

var IdRangesLogLevel = map[int]logger.Level{
	0000: logger.LevelInfo,
	1000: logger.LevelWarn,
	2000: logger.LevelError,
	3000: logger.LevelDebug,
	4000: logger.LevelTrace,
	5000: logger.LevelFatal,
	6000: logger.LevelPanic,
}

var IdStatuses = map[int]string{
	2001: logger.LevelErrorName,
	2002: logger.LevelErrorName,
	2003: logger.LevelErrorName,
	2004: logger.LevelErrorName,
	2999: logger.LevelErrorName,
}
