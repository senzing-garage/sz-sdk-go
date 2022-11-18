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

const ProductId = 6006

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var IdMessages = map[int]string{
	1:    "Enter ClearLastException().",
	2:    "Exit  ClearLastException() returned (%v).",
	3:    "Enter Destroy().",
	4:    "Exit  Destroy() returned (%v).",
	5:    "Enter GetLastException().",
	6:    "Exit  GetLastException() returned (%s, %v).",
	7:    "Enter GetLastExceptionCode().",
	8:    "Exit  GetLastExceptionCode() returned (%d, %v).",
	9:    "Enter Init(%s, %s, %s).",
	10:   "Exit  Init(%s, %s, %s) returned (%v).",
	11:   "Enter License().",
	12:   "Exit  License() returned (%s, %v).",
	13:   "Enter SetLogLevel(%v).",
	14:   "Exit  SetLogLevel(%v) returned (%v).",
	15:   "Enter ValidateLicenseFile(%s).",
	16:   "Exit  ValidateLicenseFile(%s) returned (%s, %v).",
	17:   "Enter ValidateLicenseStringBase64(%s).",
	18:   "Exit  ValidateLicenseStringBase64(%s) returned (%s, %v).",
	19:   "Enter Version().",
	20:   "Exit  Version() returned (%s, %v).",
	4001: "Call to G2Product_destroy() failed. Return code: %d",
	4002: "Call to G2Product_getLastException() failed. Return code: %d",
	4003: "Call to G2Product_init(%s, %s, %s) failed. Return code: %d",
	4004: "Call to G2Product_validateLicenseFile(%s) failed. Return code: %d",
	4005: "Call to G2Product_validateLicenseStringBase64(%s) failed. Return code: %d",
}

var IdStatuses = map[int]string{}
