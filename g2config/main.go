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
	2001: "Call to G2Config_addDataSource(%s) failed. Return code: %d",
	2002: "Call to G2Config_close() failed. Return code: %d",
	2003: "Call to G2Config_create() failed. Return code: %d",
	2004: "Call to G2Config_deleteDataSource(%s) failed. Return code: %d",
	2005: "Call to G2Config_destroy() failed. Return code: %d",
	2006: "Call to G2Config_init(%s, %s, %d) failed. Return code: %d",
	2007: "Call to G2Config_listDataSources() failed. Return code: %d",
	2008: "Call to G2Config_load(%s) failed. Return code: %d",
	2009: "Call to G2Config_save() failed. Return code: %d",
	2999: "Cannot retrieve last error message.",
	4001: "Enter G2Config_addDataSource(%s).",
	4002: "Exit  G2Config_addDataSource(%s).",
	4003: "Enter G2Config_close().",
	4004: "Exit  G2Config_close().",
	4005: "Enter G2Config_create().",
	4006: "Exit  G2Config_create().",
	4007: "Enter G2Config_deleteDataSource(%s).",
	4008: "Exit  G2Config_deleteDataSource(%s).",
	4009: "Enter G2Config_destroy().",
	4010: "Exit  G2Config_destroy().",
	4011: "Enter G2Config_init(%s, %s, %d).",
	4012: "Exit  G2Config_init(%s, %s, %d).",
	4013: "Enter G2Config_listDataSources().",
	4014: "Exit  G2Config_listDataSources().",
	4015: "Enter G2Config_load(%s).",
	4016: "Exit  G2Config_load(%s).",
	4017: "Enter G2Config_save().",
	4018: "Exit  G2Config_save().",
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
