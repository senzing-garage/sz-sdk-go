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
	1:    "Call to G2Config_destroy() failed. Return code: %d",
	2:    "Call to G2Config_init(%s, %s, %s) failed. Return code: %d",
	3:    "Call to G2Product_validateLicenseFile(%s) failed. Return code: %d",
	4:    "Call to G2Product_validateLicenseStringBase64(%s) failed. Return code: %d",
	2999: "Cannot retrieve last error message.",
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

var IdStatuses = map[int]string{
	1:    logger.LevelErrorName,
	2:    logger.LevelErrorName,
	3:    logger.LevelErrorName,
	4:    logger.LevelErrorName,
	2999: logger.LevelErrorName,
}
