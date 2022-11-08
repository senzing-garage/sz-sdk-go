/*
Package G2configmgr is a Go wrapper over Senzing's G2Configmgr C binding.

To use G2configmgr, the LD_LIBRARY_PATH environment variable must include
a path to Senzing's libraries.  Example:

	export LD_LIBRARY_PATH=/opt/senzing/g2/lib
*/
package g2configmgr

import (
	"context"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2configmgr interface {
	AddConfig(ctx context.Context, configStr string, configComments string) (int64, error)
	ClearLastException(ctx context.Context) error
	Destroy(ctx context.Context) error
	GetConfig(ctx context.Context, configID int64) (string, error)
	GetConfigList(ctx context.Context) (string, error)
	GetDefaultConfigID(ctx context.Context) (int64, error)
	GetLastException(ctx context.Context) (string, error)
	GetLastExceptionCode(ctx context.Context) (int, error)
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error
	ReplaceDefaultConfigID(ctx context.Context, oldConfigID int64, newConfigID int64) error
	SetDefaultConfigID(ctx context.Context, configID int64) error
	SetLogLevel(ctx context.Context, logLevel logger.Level) error
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdTemplate = "senzing-6002%04d"

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var IdMessages = map[int]string{
	2001: "Call to G2ConfigMgr_addConfig(%s, %s) failed. Return code: %d",
	2002: "Call to G2Config_destroy() failed. Return code: %d",
	2003: "Call to G2ConfigMgr_getConfig(%d) failed. Return code: %d",
	2004: "Call to G2ConfigMgr_getConfigList() failed. Return code: %d",
	2005: "Call to G2ConfigMgr_getDefaultConfigID() failed. Return code: %d",
	2006: "Call to G2Config_getLastException() failed. Return code: %d",
	2007: "Call to G2Config_init(%s, %s, %d) failed. Return code: %d",
	2008: "Call to G2ConfigMgr_replaceDefaultConfigID(%d, %d) failed. Return code: %d",
	2009: "Call to G2ConfigMgr_setDefaultConfigID(%d) failed. Return code: %d",
	4001: "Enter AddConfig(%s, %s).",
	4002: "Exit  AddConfig(%s, %s) returned (%d, %v).",
	4003: "Enter ClearLastException().",
	4004: "Exit  ClearLastException() returned (%v).",
	4005: "Enter Destroy().",
	4006: "Exit  Destroy() returned (%v).",
	4007: "Enter GetConfig(%d).",
	4008: "Exit  GetConfig(%d) returned (%s, %v).",
	4009: "Enter GetConfigList().",
	4010: "Exit  GetConfigList() returned (%s, %v).",
	4011: "Enter GetDefaultConfigID().",
	4012: "Exit  GetDefaultConfigID() returned (%d, %v).",
	4013: "Enter GetLastException().",
	4014: "Exit  GetLastException() returned (%s, %v).",
	4015: "Enter GetLastExceptionCode().",
	4016: "Exit  GetLastExceptionCode() returned (%d, %v).",
	4017: "Enter Init(%s, %s, %d).",
	4018: "Exit  Init(%s, %s, %d) returned (%v).",
	4019: "Enter ReplaceDefaultConfigID(%d, %d).",
	4020: "Exit  ReplaceDefaultConfigID(%d, %d) returned (%v).",
	4021: "Enter SetDefaultConfigID(%d).",
	4022: "Exit  SetDefaultConfigID(%d) returned (%v).",
	4023: "Enter SetLogLevel(%v).",
	4024: "Exit  SetLogLevel(%v) returned (%v).",
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
	2999: logger.LevelErrorName,
}
