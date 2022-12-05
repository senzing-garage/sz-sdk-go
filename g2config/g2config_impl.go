// The G2configImpl implementation is a wrapper over the Senzing libg2config library.
package g2config

/*
#include "g2config.h"
#cgo CFLAGS: -g -I/opt/senzing/g2/sdk/c
#cgo LDFLAGS: -L/opt/senzing/g2/lib -lanalytics -ldb2plugin -lG2 -lg2AddressComp -lg2AddressHasher -lg2CloseNames -lg2CompJavaScoreSet -lg2ConfigParseAddr -lg2DateComp -lg2DistinctFeatJava -lg2DLComp -lg2EFeatJava -lg2EmailComp -lg2ExactDomainMatchComp -lg2ExactMatchComp -lg2FeatBuilder -lg2FormatSSN -lg2GenericHasher -lg2GEOLOCComp -lg2GNRNameComp -lg2GroupAssociationComp -lG2Hasher -lg2IDHasher -lg2JVMPlugin -lg2NameHasher -lg2ParseDOB -lg2ParseEmail -lg2ParseGEOLOC -lg2ParseID -lg2ParseName -lg2ParsePhone -lg2PartialAddresses -lg2PartialDates -lg2PartialNames -lg2PhoneComp -lg2PhoneHasher -lG2SSAdm -lg2SSNComp -lg2STBHasher -lg2StdCountry -lg2StdJava -lg2StdTokenizeName -lg2StrictSubsetFelems -lg2StrictSubsetNormalizedFelems -lg2StrictSubsetTokens -lg2StringComp -lmariadbplugin -lmssqlplugin -lNameDataObject -loracleplugin -lpostgresqlplugin -lscoring -lSpaceTimeBoxStandardizer -lsqliteplugin
*/
import "C"

import (
	"bytes"
	"context"
	"errors"
	"time"
	"unsafe"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagelogger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// G2configImpl is the default implementation of the G2config interface.
type G2configImpl struct {
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
func (g2config *G2configImpl) getByteArrayC(size int) *C.char {
	bytes := C.malloc(C.size_t(size))
	return (*C.char)(bytes)
}

// Make a byte array.
func (g2config *G2configImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

// Create a new error.
func (g2config *G2configImpl) newError(ctx context.Context, errorNumber int, details ...interface{}) error {
	lastException, err := g2config.GetLastException(ctx)
	defer g2config.ClearLastException(ctx)
	message := lastException
	if err != nil {
		message = err.Error()
	}

	var newDetails []interface{}
	newDetails = append(newDetails, details...)
	newDetails = append(newDetails, errors.New(message))
	errorMessage, err := g2config.getLogger().Message(errorNumber, newDetails...)
	if err != nil {
		errorMessage = err.Error()
	}

	return errors.New(errorMessage)
}

// Get the Logger singleton.
func (g2config *G2configImpl) getLogger() messagelogger.MessageLoggerInterface {
	if g2config.logger == nil {
		g2config.logger, _ = messagelogger.NewSenzingApiLogger(ProductId, IdMessages, IdStatuses, messagelogger.LevelInfo)
	}
	return g2config.logger
}

// Trace method entry.
func (g2config *G2configImpl) traceEntry(errorNumber int, details ...interface{}) {
	g2config.getLogger().Log(errorNumber, details...)
}

// Trace method exit.
func (g2config *G2configImpl) traceExit(errorNumber int, details ...interface{}) {
	g2config.getLogger().Log(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The AddDataSource method adds a data source to an existing in-memory configuration.

Input
  - ctx: A context to control lifecycle.
  - configHandle: A pointer to a configuration.
  - inputJson: A JSON document in the format `{"DSRC_CODE": "NAME_OF_DATASOURCE"}`.

Output
  - A string containing a JSON document listing the newly created data source.
    See the example output.
*/
func (g2config *G2configImpl) AddDataSource(ctx context.Context, configHandle uintptr, inputJson string) (string, error) {
	// _DLEXPORT int G2Config_addDataSource(ConfigHandle configHandle, const char *inputJson, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2config.isTrace {
		g2config.traceEntry(1, configHandle, inputJson)
	}
	entryTime := time.Now()
	var err error = nil
	inputJsonForC := C.CString(inputJson)
	defer C.free(unsafe.Pointer(inputJsonForC))
	result := C.G2Config_addDataSource_helper(C.uintptr_t(configHandle), inputJsonForC)
	if result.returnCode != 0 {
		err = g2config.newError(ctx, 4001, configHandle, inputJson, result.returnCode, result, time.Since(entryTime))
	}
	if g2config.isTrace {
		defer g2config.traceExit(2, configHandle, inputJson, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

/*
The ClearLastException method erases the last exception message held by the Senzing G2Config object.

Input
  - ctx: A context to control lifecycle.
*/
func (g2config *G2configImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2Config_clearLastException();
	if g2config.isTrace {
		g2config.traceEntry(3)
	}
	entryTime := time.Now()
	var err error = nil
	C.G2Config_clearLastException()
	if g2config.isTrace {
		defer g2config.traceExit(4, err, time.Since(entryTime))
	}
	return err
}

/*
The Close method cleans up the Senzing G2Config object pointed to by the handle.

Input
  - ctx: A context to control lifecycle.
  - configHandle: A pointer to a configuration.
*/
func (g2config *G2configImpl) Close(ctx context.Context, configHandle uintptr) error {
	// _DLEXPORT int G2Config_close(ConfigHandle configHandle);
	if g2config.isTrace {
		g2config.traceEntry(5, configHandle)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2config_close_helper(C.uintptr_t(configHandle))
	if result != 0 {
		err = g2config.newError(ctx, 4002, configHandle, result, time.Since(entryTime))
	}
	if g2config.isTrace {
		defer g2config.traceExit(6, configHandle, err, time.Since(entryTime))
	}
	return err
}

/*
The Create method creates a stock G2 JSON config from the template config.

Input
  - ctx: A context to control lifecycle.

Output
  - A Pointer to an in-memory configuration.
*/
func (g2config *G2configImpl) Create(ctx context.Context) (uintptr, error) {
	// _DLEXPORT int G2Config_create(ConfigHandle* configHandle);
	if g2config.isTrace {
		g2config.traceEntry(7)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2config_create_helper()
	if result.returnCode != 0 {
		err = g2config.newError(ctx, 4003, result.returnCode, time.Since(entryTime))
	}
	if g2config.isTrace {
		defer g2config.traceExit(8, (uintptr)(result.response), err, time.Since(entryTime))
	}
	return (uintptr)(result.response), err
}

/*
The DeleteDataSource method removes a data source from an existing configuration.

Input
  - ctx: A context to control lifecycle.
  - configHandle: A pointer to a configuration.
  - inputJson: A JSON document in the format `{"DSRC_CODE": "NAME_OF_DATASOURCE"}`.
*/
func (g2config *G2configImpl) DeleteDataSource(ctx context.Context, configHandle uintptr, inputJson string) error {
	// _DLEXPORT int G2Config_deleteDataSource(ConfigHandle configHandle, const char *inputJson);
	if g2config.isTrace {
		g2config.traceEntry(9, configHandle, inputJson)
	}
	entryTime := time.Now()
	var err error = nil
	inputJsonForC := C.CString(inputJson)
	defer C.free(unsafe.Pointer(inputJsonForC))
	result := C.G2Config_deleteDataSource_helper(C.uintptr_t(configHandle), inputJsonForC)
	if result != 0 {
		err = g2config.newError(ctx, 4004, configHandle, inputJson, result, time.Since(entryTime))
	}
	if g2config.isTrace {
		defer g2config.traceExit(10, configHandle, inputJson, err, time.Since(entryTime))
	}
	return err
}

/*
The Destroy method will destroy and perform cleanup for the Senzing G2Config object.
It should be called after all other calls are complete.

Input
  - ctx: A context to control lifecycle.
*/
func (g2config *G2configImpl) Destroy(ctx context.Context) error {
	// _DLEXPORT int G2Config_destroy();
	if g2config.isTrace {
		g2config.traceEntry(11)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Config_destroy()
	if result != 0 {
		err = g2config.newError(ctx, 4005, result, time.Since(entryTime))
	}
	if g2config.isTrace {
		defer g2config.traceExit(12, err, time.Since(entryTime))
	}
	return err
}

/*
The GetLastException method retrieves the last exception thrown in Senzing's G2Config.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the error received from Senzing's G2Config.
*/
func (g2config *G2configImpl) GetLastException(ctx context.Context) (string, error) {
	// _DLEXPORT int G2Config_getLastException(char *buffer, const size_t bufSize);
	if g2config.isTrace {
		g2config.traceEntry(13)
	}
	entryTime := time.Now()
	var err error = nil
	stringBuffer := g2config.getByteArray(initialByteArraySize)
	C.G2Config_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	// if result == 0 { // "result" is length of exception message.
	// 	err = g2config.getLogger().Error(4005, result, time.Since(entryTime))
	// }
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if g2config.isTrace {
		defer g2config.traceExit(14, string(stringBuffer), err, time.Since(entryTime))
	}
	return string(stringBuffer), err
}

/*
The GetLastExceptionCode method retrieves the code of the last exception thrown in Senzing's G2Config.

Input:
  - ctx: A context to control lifecycle.

Output:
  - An int containing the error received from Senzing's G2Config.
*/
func (g2config *G2configImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2Config_getLastExceptionCode();
	if g2config.isTrace {
		g2config.traceEntry(15)
	}
	entryTime := time.Now()
	var err error = nil
	result := int(C.G2Config_getLastExceptionCode())
	if g2config.isTrace {
		defer g2config.traceExit(16, result, err, time.Since(entryTime))
	}
	return result, err
}

/*
The Init method initializes the Senzing G2Config object.
It must be called prior to any other calls.

Input
  - ctx: A context to control lifecycle.
  - moduleName: A name for the auditing node, to help identify it within system logs.
  - iniParams: A JSON string containing configuration paramters.
  - verboseLogging: A flag to enable deeper logging of the G2 processing. 0 for no Senzing logging; 1 for logging.
*/
func (g2config *G2configImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2Config_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	if g2config.isTrace {
		g2config.traceEntry(17, moduleName, iniParams, verboseLogging)
	}
	entryTime := time.Now()
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2Config_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))
	if result != 0 {
		err = g2config.newError(ctx, 4007, moduleName, iniParams, verboseLogging, result, time.Since(entryTime))
	}
	if g2config.isTrace {
		defer g2config.traceExit(18, moduleName, iniParams, verboseLogging, err, time.Since(entryTime))
	}
	return err
}

/*
The ListDataSources method returns a JSON document of data sources.

Input
  - ctx: A context to control lifecycle.
  - configHandle: A pointer to a configuration.

Output
  - A string containing a JSON document listing all of the data sources.
*/
func (g2config *G2configImpl) ListDataSources(ctx context.Context, configHandle uintptr) (string, error) {
	// _DLEXPORT int G2Config_listDataSources(ConfigHandle configHandle, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2config.isTrace {
		g2config.traceEntry(19, configHandle)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Config_listDataSources_helper(C.uintptr_t(configHandle))
	if result.returnCode != 0 {
		err = g2config.newError(ctx, 4008, result.returnCode, result, time.Since(entryTime))
	}
	if g2config.isTrace {
		defer g2config.traceExit(20, configHandle, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

/*
The Load method initializes the Senzing G2Config object from a JSON string.

Input
  - ctx: A context to control lifecycle.
  - configHandle: A pointer to a configuration.
  - jsonConfig: A JSON document containing the Senzing configuration.
*/
func (g2config *G2configImpl) Load(ctx context.Context, configHandle uintptr, jsonConfig string) error {
	// _DLEXPORT int G2Config_load(const char *jsonConfig,ConfigHandle* configHandle);
	if g2config.isTrace {
		g2config.traceEntry(21, configHandle, jsonConfig)
	}
	entryTime := time.Now()
	var err error = nil
	jsonConfigForC := C.CString(jsonConfig)
	defer C.free(unsafe.Pointer(jsonConfigForC))
	result := C.G2Config_load_helper(C.uintptr_t(configHandle), jsonConfigForC)
	if result != 0 {
		err = g2config.newError(ctx, 4009, configHandle, jsonConfig, result, time.Since(entryTime))
	}
	if g2config.isTrace {
		defer g2config.traceExit(22, configHandle, jsonConfig, err, time.Since(entryTime))
	}
	return err
}

/*
The Save method creates a JSON string representation of the Senzing G2Config object.

Input
  - ctx: A context to control lifecycle.
  - configHandle: A pointer to a configuration.

Output
  - A string containing a JSON Document representation of the Senzing G2Config object.
    See the example output.
*/
func (g2config *G2configImpl) Save(ctx context.Context, configHandle uintptr) (string, error) {
	// _DLEXPORT int G2Config_save(ConfigHandle configHandle, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2config.isTrace {
		g2config.traceEntry(23, configHandle)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Config_save_helper(C.uintptr_t(configHandle))
	if result.returnCode != 0 {
		err = g2config.newError(ctx, 4010, configHandle, result.returnCode, result, time.Since(entryTime))
	}
	if g2config.isTrace {
		defer g2config.traceExit(24, configHandle, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

/*
The SetLogLevel method sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevel: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
func (g2config *G2configImpl) SetLogLevel(ctx context.Context, logLevel logger.Level) error {
	if g2config.isTrace {
		g2config.traceEntry(25, logLevel)
	}
	entryTime := time.Now()
	var err error = nil
	g2config.getLogger().SetLogLevel(messagelogger.Level(logLevel))
	g2config.isTrace = (g2config.getLogger().GetLogLevel() == messagelogger.LevelTrace)
	if g2config.isTrace {
		defer g2config.traceExit(26, logLevel, err, time.Since(entryTime))
	}
	return err
}
