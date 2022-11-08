/*
Package g2config is a Go wrapper over Senzing's G2Config C binding.

To use G2config, the LD_LIBRARY_PATH environment variable must include
a path to Senzing's libraries.  Example:

	export LD_LIBRARY_PATH=/opt/senzing/g2/lib
*/
package g2config

import (
	"context"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2config interface {
	AddDataSource(ctx context.Context, configHandle uintptr, inputJson string) (string, error)
	ClearLastException(ctx context.Context) error
	Close(ctx context.Context, configHandle uintptr) error
	Create(ctx context.Context) (uintptr, error)
	DeleteDataSource(ctx context.Context, configHandle uintptr, inputJson string) error
	Destroy(ctx context.Context) error
	GetLastException(ctx context.Context) (string, error)
	GetLastExceptionCode(ctx context.Context) (int, error)
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error
	ListDataSources(ctx context.Context, configHandle uintptr) (string, error)
	Load(ctx context.Context, configHandle uintptr, jsonConfig string) error
	Save(ctx context.Context, configHandle uintptr) (string, error)
	SetLogLevel(ctx context.Context, logLevel logger.Level) error
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdTemplate = "senzing-6001%04d"

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var IdMessages = map[int]string{
	2001: "Call to G2Config_addDataSource(%v, %s) failed. Return code: %d",
	2002: "Call to G2Config_close(%v) failed. Return code: %d",
	2003: "Call to G2Config_create() failed. Return code: %d",
	2004: "Call to G2Config_deleteDataSource(%v, %s) failed. Return code: %d",
	2005: "Call to G2Config_destroy() failed. Return code: %d",
	2006: "Call to G2Config_init(%s, %s, %d) failed. Return code: %d",
	2007: "Call to G2Config_listDataSources() failed. Return code: %d",
	2008: "Call to G2Config_load(%v, %s) failed. Return code: %d",
	2009: "Call to G2Config_save(%v) failed. Return code: %d",
	2999: "Cannot retrieve last error message.",
	4001: "Enter AddDataSource(%v, %s).",
	4002: "Exit  AddDataSource(%v, %s) returned (%s, %v).",
	4003: "Enter ClearLastException().",
	4004: "Exit  ClearLastException() returned (%v).",
	4005: "Enter Close(%v).",
	4006: "Exit  Close(%v) returned (%v).",
	4007: "Enter Create().",
	4008: "Exit  Create() returned (%v, %v).",
	4009: "Enter DeleteDataSource(%v, %s).",
	4010: "Exit  DeleteDataSource(%v, %s) returned (%v).",
	4011: "Enter Destroy().",
	4012: "Exit  Destroy() returned (%v).",
	4013: "Enter GetLastException().",
	4014: "Exit  GetLastException() returned (%s, %v).",
	4015: "Enter GetLastExceptionCode().",
	4016: "Exit  GetLastExceptionCode() returned (%d, %v).",
	4017: "Enter Init(%s, %s, %d).",
	4018: "Exit  Init(%s, %s, %d) returned (%v).",
	4019: "Enter ListDataSources(%v).",
	4020: "Exit  ListDataSources(%v) returned (%s, %v).",
	4021: "Enter Load(%v, %s).",
	4022: "Exit  Load(%v, %s) returned (%v).",
	4023: "Enter Save(%v).",
	4024: "Exit  Save(%v) returned (%s, %v).",
	4025: "Enter SetLogLevel(%v).",
	4026: "Exit  SetLogLevel(%v) returned (%v).",
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
	2005: logger.LevelErrorName,
	2006: logger.LevelErrorName,
	2007: logger.LevelErrorName,
	2008: logger.LevelErrorName,
	2009: logger.LevelErrorName,
	2999: logger.LevelErrorName,
}
