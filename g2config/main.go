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

const ProductId = 6001

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var IdMessages = map[int]string{
	1:    "Enter AddDataSource(%v, %s).",
	2:    "Exit  AddDataSource(%v, %s) returned (%s, %v).",
	3:    "Enter ClearLastException().",
	4:    "Exit  ClearLastException() returned (%v).",
	5:    "Enter Close(%v).",
	6:    "Exit  Close(%v) returned (%v).",
	7:    "Enter Create().",
	8:    "Exit  Create() returned (%v, %v).",
	9:    "Enter DeleteDataSource(%v, %s).",
	10:   "Exit  DeleteDataSource(%v, %s) returned (%v).",
	11:   "Enter Destroy().",
	12:   "Exit  Destroy() returned (%v).",
	13:   "Enter GetLastException().",
	14:   "Exit  GetLastException() returned (%s, %v).",
	15:   "Enter GetLastExceptionCode().",
	16:   "Exit  GetLastExceptionCode() returned (%d, %v).",
	17:   "Enter Init(%s, %s, %d).",
	18:   "Exit  Init(%s, %s, %d) returned (%v).",
	19:   "Enter ListDataSources(%v).",
	20:   "Exit  ListDataSources(%v) returned (%s, %v).",
	21:   "Enter Load(%v, %s).",
	22:   "Exit  Load(%v, %s) returned (%v).",
	23:   "Enter Save(%v).",
	24:   "Exit  Save(%v) returned (%s, %v).",
	25:   "Enter SetLogLevel(%v).",
	26:   "Exit  SetLogLevel(%v) returned (%v).",
	4001: "Call to G2Config_addDataSource(%v, %s) failed. Return code: %d",
	4002: "Call to G2Config_close(%v) failed. Return code: %d",
	4003: "Call to G2Config_create() failed. Return code: %d",
	4004: "Call to G2Config_deleteDataSource(%v, %s) failed. Return code: %d",
	4005: "Call to G2Config_getLastException() failed. Return code: %d",
	4006: "Call to G2Config_destroy() failed. Return code: %d",
	4007: "Call to G2Config_init(%s, %s, %d) failed. Return code: %d",
	4008: "Call to G2Config_listDataSources() failed. Return code: %d",
	4009: "Call to G2Config_load(%v, %s) failed. Return code: %d",
	4010: "Call to G2Config_save(%v) failed. Return code: %d",
}

var XIdLevelRangeAsString = map[int]string{
	0000: logger.LevelTraceName,
	1000: logger.LevelDebugName,
	2000: logger.LevelInfoName,
	3000: logger.LevelWarnName,
	4000: logger.LevelErrorName,
	5000: logger.LevelFatalName,
	6000: logger.LevelPanicName,
}

var IdLevelRanges = map[int]logger.Level{
	0000: logger.LevelTrace,
	1000: logger.LevelDebug,
	2000: logger.LevelInfo,
	3000: logger.LevelWarn,
	4000: logger.LevelError,
	5000: logger.LevelFatal,
	6000: logger.LevelPanic,
}

var IdStatuses = map[int]string{
	4001: logger.LevelErrorName,
	4002: logger.LevelErrorName,
	4003: logger.LevelErrorName,
	4004: logger.LevelErrorName,
	4005: logger.LevelErrorName,
	4006: logger.LevelErrorName,
	4007: logger.LevelErrorName,
	4008: logger.LevelErrorName,
	4009: logger.LevelErrorName,
	4010: logger.LevelErrorName,
}
