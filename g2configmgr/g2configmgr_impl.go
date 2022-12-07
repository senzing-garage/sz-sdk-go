/*
The G2configmgrImpl implementation is a wrapper over the Senzing libg2configmgr library.
*/
package g2configmgr

/*
#include "g2configmgr.h"
#cgo CFLAGS: -g -I/opt/senzing/g2/sdk/c
#cgo LDFLAGS: -L/opt/senzing/g2/lib -lanalytics -ldb2plugin -lG2 -lg2AddressComp -lg2AddressHasher -lg2CloseNames -lg2CompJavaScoreSet -lg2ConfigParseAddr -lg2DateComp -lg2DistinctFeatJava -lg2DLComp -lg2EFeatJava -lg2EmailComp -lg2ExactDomainMatchComp -lg2ExactMatchComp -lg2FeatBuilder -lg2FormatSSN -lg2GenericHasher -lg2GEOLOCComp -lg2GNRNameComp -lg2GroupAssociationComp -lG2Hasher -lg2IDHasher -lg2JVMPlugin -lg2NameHasher -lg2ParseDOB -lg2ParseEmail -lg2ParseGEOLOC -lg2ParseID -lg2ParseName -lg2ParsePhone -lg2PartialAddresses -lg2PartialDates -lg2PartialNames -lg2PhoneComp -lg2PhoneHasher -lG2SSAdm -lg2SSNComp -lg2STBHasher -lg2StdCountry -lg2StdJava -lg2StdTokenizeName -lg2StrictSubsetFelems -lg2StrictSubsetNormalizedFelems -lg2StrictSubsetTokens -lg2StringComp -lmariadbplugin -lmssqlplugin -lNameDataObject -loracleplugin -lpostgresqlplugin -lscoring -lSpaceTimeBoxStandardizer -lsqliteplugin
*/
import "C"

import (
	"bytes"
	"context"
	"errors"
	"runtime"
	"time"
	"unsafe"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagelogger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// G2configmgrImpl is the default implementation of the G2configmgr interface.
type G2configmgrImpl struct {
	isTrace bool
	logger  messagelogger.MessageLoggerInterface
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const initialByteArraySize = 65535

// ----------------------------------------------------------------------------
// Internal methods - names begin with lower case
// ----------------------------------------------------------------------------

// Get space for an array of bytes of a given size.
func (g2configmgr *G2configmgrImpl) getByteArrayC(size int) *C.char {
	bytes := C.malloc(C.size_t(size))
	return (*C.char)(bytes)
}

// Make a byte array.
func (g2configmgr *G2configmgrImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

// Create a new error.
func (g2configmgr *G2configmgrImpl) newError(ctx context.Context, errorNumber int, details ...interface{}) error {
	lastException, err := g2configmgr.GetLastException(ctx)
	defer g2configmgr.ClearLastException(ctx)
	message := lastException
	if err != nil {
		message = err.Error()
	}

	var newDetails []interface{}
	newDetails = append(newDetails, details...)
	newDetails = append(newDetails, errors.New(message))
	errorMessage, err := g2configmgr.getLogger().Message(errorNumber, newDetails...)
	if err != nil {
		errorMessage = err.Error()
	}

	return errors.New(errorMessage)
}

// Get the Logger singleton.
func (g2configmgr *G2configmgrImpl) getLogger() messagelogger.MessageLoggerInterface {
	if g2configmgr.logger == nil {
		g2configmgr.logger, _ = messagelogger.NewSenzingApiLogger(ProductId, IdMessages, IdStatuses, messagelogger.LevelInfo)
	}
	return g2configmgr.logger
}

// Trace method entry.
func (g2configmgr *G2configmgrImpl) traceEntry(errorNumber int, details ...interface{}) {
	g2configmgr.getLogger().Log(errorNumber, details...)
}

// Trace method exit.
func (g2configmgr *G2configmgrImpl) traceExit(errorNumber int, details ...interface{}) {
	g2configmgr.getLogger().Log(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The AddConfig method adds a Senzing configuration JSON document to the Senzing database.

Input
  - ctx: A context to control lifecycle.
  - configStr: The Senzing configuration JSON document.
  - configComments: A free-form string of comments describing the configuration document.

Output
  - A configuration identifier.
*/
func (g2configmgr *G2configmgrImpl) AddConfig(ctx context.Context, configStr string, configComments string) (int64, error) {
	// _DLEXPORT int G2ConfigMgr_addConfig(const char* configStr, const char* configComments, long long* configID);
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(1, configStr, configComments)
	}
	entryTime := time.Now()
	var err error = nil
	configStrForC := C.CString(configStr)
	defer C.free(unsafe.Pointer(configStrForC))
	configCommentsForC := C.CString(configComments)
	defer C.free(unsafe.Pointer(configCommentsForC))
	result := C.G2ConfigMgr_addConfig_helper(configStrForC, configCommentsForC)
	if result.returnCode != 0 {
		err = g2configmgr.newError(ctx, 4001, configStr, configComments, result.returnCode, result, time.Since(entryTime))
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(2, configStr, configComments, int64(C.longlong(result.configID)), err, time.Since(entryTime))
	}
	return int64(C.longlong(result.configID)), err
}

/*
The ClearLastException method erases the last exception message held by the Senzing G2ConfigMgr object.

Input
  - ctx: A context to control lifecycle.
*/
func (g2configmgr *G2configmgrImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2Config_clearLastException()
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(3)
	}
	entryTime := time.Now()
	var err error = nil
	C.G2ConfigMgr_clearLastException()
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(4, err, time.Since(entryTime))
	}
	return err
}

/*
The Destroy method will destroy and perform cleanup for the Senzing G2ConfigMgr object.
It should be called after all other calls are complete.

Input
  - ctx: A context to control lifecycle.
*/
func (g2configmgr *G2configmgrImpl) Destroy(ctx context.Context) error {
	// _DLEXPORT int G2Config_destroy();
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(5)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2ConfigMgr_destroy()
	if result != 0 {
		err = g2configmgr.newError(ctx, 4002, result, time.Since(entryTime))
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(6, err, time.Since(entryTime))
	}
	return err
}

/*
The GetConfig method retrieves a specific Senzing configuration JSON document from the Senzing database.

Input
  - ctx: A context to control lifecycle.
  - configID: The configuration identifier of the desired Senzing Engine configuration JSON document to retrieve.

Output
  - A JSON document containing the Senzing configuration.
    See the example output.
*/
func (g2configmgr *G2configmgrImpl) GetConfig(ctx context.Context, configID int64) (string, error) {
	// _DLEXPORT int G2ConfigMgr_getConfig(const long long configID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(7, configID)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2ConfigMgr_getConfig_helper(C.longlong(configID))
	if result.returnCode != 0 {
		err = g2configmgr.newError(ctx, 4003, configID, result.returnCode, result, time.Since(entryTime))
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(8, configID, C.GoString(result.config), err, time.Since(entryTime))
	}
	return C.GoString(result.config), err
}

/*
The GetConfigList method retrieves a list of Senzing configurations from the Senzing database.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document containing Senzing configurations.
    See the example output.
*/
func (g2configmgr *G2configmgrImpl) GetConfigList(ctx context.Context) (string, error) {
	// _DLEXPORT int G2ConfigMgr_getConfigList(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(9)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2ConfigMgr_getConfigList_helper()
	if result.returnCode != 0 {
		err = g2configmgr.newError(ctx, 4004, result.returnCode, result, time.Since(entryTime))
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(10, C.GoString(result.configList), err, time.Since(entryTime))
	}
	return C.GoString(result.configList), err
}

/*
The GetDefaultConfigID method retrieves from the Senzing database the configuration identifier of the default Senzing configuration.

Input
  - ctx: A context to control lifecycle.

Output
  - A configuration identifier which identifies the current configuration in use.
*/
func (g2configmgr *G2configmgrImpl) GetDefaultConfigID(ctx context.Context) (int64, error) {
	//  _DLEXPORT int G2ConfigMgr_getDefaultConfigID(long long* configID);
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(11)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2ConfigMgr_getDefaultConfigID_helper()
	if result.returnCode != 0 {
		err = g2configmgr.newError(ctx, 4005, result.returnCode, result, time.Since(entryTime))
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(12, int64(C.longlong(result.configID)), err, time.Since(entryTime))
	}
	return int64(C.longlong(result.configID)), err
}

/*
The GetLastException method retrieves the last exception thrown in Senzing's G2ConfigMgr.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the error received from Senzing's G2ConfigMgr.
*/
func (g2configmgr *G2configmgrImpl) GetLastException(ctx context.Context) (string, error) {
	// _DLEXPORT int G2Config_getLastException(char *buffer, const size_t bufSize);
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(13)
	}
	entryTime := time.Now()
	var err error = nil
	stringBuffer := g2configmgr.getByteArray(initialByteArraySize)
	C.G2ConfigMgr_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	// if result == 0 { // "result" is length of exception message.
	// 	err = g2configmgr.getLogger().Error(4006, result, time.Since(entryTime))
	// }
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(14, string(stringBuffer), err, time.Since(entryTime))
	}
	return string(stringBuffer), err
}

/*
The GetLastExceptionCode method retrieves the code of the last exception thrown in Senzing's G2ConfigMgr.

Input:
  - ctx: A context to control lifecycle.

Output:
  - An int containing the error received from Senzing's G2ConfigMgr.
*/
func (g2configmgr *G2configmgrImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2Config_getLastExceptionCode();
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(15)
	}
	entryTime := time.Now()
	var err error = nil
	result := int(C.G2ConfigMgr_getLastExceptionCode())
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(16, result, err, time.Since(entryTime))
	}
	return result, err
}

/*
The Init method initializes the Senzing G2ConfigMgr object.
It must be called prior to any other calls.

Input
  - ctx: A context to control lifecycle.
  - moduleName: A name for the auditing node, to help identify it within system logs.
  - iniParams: A JSON string containing configuration paramters.
  - verboseLogging: A flag to enable deeper logging of the G2 processing. 0 for no Senzing logging; 1 for logging.
*/
func (g2configmgr *G2configmgrImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2Config_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(17, moduleName, iniParams, verboseLogging)
	}
	entryTime := time.Now()
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2ConfigMgr_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))
	if result != 0 {
		err = g2configmgr.newError(ctx, 4007, moduleName, iniParams, verboseLogging, result, time.Since(entryTime))
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(18, moduleName, iniParams, verboseLogging, err, time.Since(entryTime))
	}
	return err
}

/*
The ReplaceDefaultConfigID method replaces the old configuration identifier with a new configuration identifier in the Senzing database.
It is like a "compare-and-swap" instruction to serialize concurrent editing of configuration.
If oldConfigID is no longer the "old configuration identifier", the operation will fail.
To simply set the default configuration ID, use SetDefaultConfigID().

Input
  - ctx: A context to control lifecycle.
  - oldConfigID: The configuration identifier to replace.
  - newConfigID: The configuration identifier to use as the default.
*/
func (g2configmgr *G2configmgrImpl) ReplaceDefaultConfigID(ctx context.Context, oldConfigID int64, newConfigID int64) error {
	// _DLEXPORT int G2ConfigMgr_replaceDefaultConfigID(const long long oldConfigID, const long long newConfigID);
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(19, oldConfigID, newConfigID)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2ConfigMgr_replaceDefaultConfigID(C.longlong(oldConfigID), C.longlong(newConfigID))
	if result != 0 {
		err = g2configmgr.newError(ctx, 4008, oldConfigID, newConfigID, result, time.Since(entryTime))
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(20, oldConfigID, newConfigID, err, time.Since(entryTime))
	}
	return err
}

/*
The SetDefaultConfigID method replaces the sets a new configuration identifier in the Senzing database.
To serialize modifying of the configuration identifier, see ReplaceDefaultConfigID().

Input
  - ctx: A context to control lifecycle.
  - configID: The configuration identifier of the Senzing Engine configuration to use as the default.
*/
func (g2configmgr *G2configmgrImpl) SetDefaultConfigID(ctx context.Context, configID int64) error {
	// _DLEXPORT int G2ConfigMgr_setDefaultConfigID(const long long configID);
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(21, configID)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2ConfigMgr_setDefaultConfigID(C.longlong(configID))
	if result != 0 {
		err = g2configmgr.newError(ctx, 4009, configID, result, time.Since(entryTime))
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(22, configID, err, time.Since(entryTime))
	}
	return err
}

/*
The SetLogLevel method sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevel: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
func (g2configmgr *G2configmgrImpl) SetLogLevel(ctx context.Context, logLevel logger.Level) error {
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(23, logLevel)
	}
	entryTime := time.Now()
	var err error = nil
	g2configmgr.getLogger().SetLogLevel(messagelogger.Level(logLevel))
	g2configmgr.isTrace = (g2configmgr.getLogger().GetLogLevel() == messagelogger.LevelTrace)
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(24, logLevel, err, time.Since(entryTime))
	}
	return err
}
