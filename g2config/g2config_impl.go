/*
The G2configImpl implementation...
*/
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
	"unsafe"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messageid"
	"github.com/senzing/go-logging/messagelogger"
	"github.com/senzing/go-logging/messageloglevel"
	"github.com/senzing/go-logging/messagestatus"
	"github.com/senzing/go-logging/messagetext"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2configImpl struct {
	isTrace          bool
	logger           messagelogger.MessageLoggerInterface
	messageGenerator messagelogger.MessageLoggerInterface
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

func (g2config *G2configImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

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
	messageGenerator := g2config.getMessageGenerator(ctx)
	errorMessage, err := messageGenerator.Message(errorNumber, newDetails...)
	if err != nil {
		errorMessage = err.Error()
	}

	return errors.New(errorMessage)
}

func (g2config *G2configImpl) getLogger() messagelogger.MessageLoggerInterface {
	if g2config.logger == nil {
		messageFormat := &messageformat.MessageFormatJson{}
		messageId := &messageid.MessageIdTemplated{
			MessageIdTemplate: MessageIdTemplate,
		}
		messageLogLevel := &messageloglevel.MessageLogLevelByIdRange{
			IdRanges: IdRangesLogLevel,
		}
		messageStatus := &messagestatus.MessageStatusByIdRange{
			IdRanges: IdRanges,
		}
		messageText := &messagetext.MessageTextTemplated{
			IdMessages: IdMessages,
		}
		g2config.logger, _ = messagelogger.New(messageFormat, messageId, messageLogLevel, messageStatus, messageText, messagelogger.LevelInfo)
	}
	return g2config.logger
}

func (g2config *G2configImpl) getMessageGenerator(ctx context.Context) messagelogger.MessageLoggerInterface {
	if g2config.messageGenerator == nil {
		messageFormat := &messageformat.MessageFormatJson{}
		messageId := &messageid.MessageIdTemplated{
			MessageIdTemplate: MessageIdTemplate,
		}
		messageLogLevel := &messageloglevel.MessageLogLevelSenzingApi{
			IdRanges:   IdRanges,
			IdStatuses: IdStatuses,
		}
		messageStatus := &messagestatus.MessageStatusSenzingApi{
			IdRanges:   IdRanges,
			IdStatuses: IdStatuses,
		}
		messageText := &messagetext.MessageTextTemplated{
			IdMessages: IdMessages,
		}
		g2config.messageGenerator, _ = messagelogger.New(messageFormat, messageId, messageLogLevel, messageStatus, messageText, messagelogger.LevelInfo)
	}
	return g2config.messageGenerator
}

func (g2config *G2configImpl) traceEntry(errorNumber int, details ...interface{}) {
	g2config.getLogger().Log(errorNumber, details...)
}

func (g2config *G2configImpl) traceExit(errorNumber int, details ...interface{}) {
	g2config.getLogger().Log(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func (g2config *G2configImpl) AddDataSource(ctx context.Context, configHandle uintptr, inputJson string) (string, error) {
	// _DLEXPORT int G2Config_addDataSource(ConfigHandle configHandle, const char *inputJson, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2config.isTrace {
		g2config.traceEntry(4001, configHandle, inputJson)
	}
	var err error = nil
	inputJsonForC := C.CString(inputJson)
	defer C.free(unsafe.Pointer(inputJsonForC))
	result := C.G2Config_addDataSource_helper(C.uintptr_t(configHandle), inputJsonForC)
	if result.returnCode != 0 {
		err = g2config.newError(ctx, 2001, configHandle, inputJson, result.returnCode, result)
	}
	if g2config.isTrace {
		defer g2config.traceExit(4002, configHandle, inputJson, C.GoString(result.response), err)
	}
	return C.GoString(result.response), err
}

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2config *G2configImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2Config_clearLastException();
	if g2config.isTrace {
		g2config.traceEntry(4003)
	}
	var err error = nil
	C.G2Config_clearLastException()
	if g2config.isTrace {
		defer g2config.traceExit(4004, err)
	}
	return err
}

func (g2config *G2configImpl) Close(ctx context.Context, configHandle uintptr) error {
	// _DLEXPORT int G2Config_close(ConfigHandle configHandle);
	if g2config.isTrace {
		g2config.traceEntry(4005, configHandle)
	}
	var err error = nil
	result := C.G2config_close_helper(C.uintptr_t(configHandle))
	if result != 0 {
		err = g2config.newError(ctx, 2002, configHandle, result)
	}
	if g2config.isTrace {
		defer g2config.traceExit(4006, configHandle, err)
	}
	return err
}

func (g2config *G2configImpl) Create(ctx context.Context) (uintptr, error) {
	// _DLEXPORT int G2Config_create(ConfigHandle* configHandle);
	if g2config.isTrace {
		g2config.traceEntry(4007)
	}
	var err error = nil
	result := C.G2config_create_helper()
	if result.returnCode != 0 {
		err = g2config.newError(ctx, 2003, result.returnCode)
	}
	if g2config.isTrace {
		defer g2config.traceExit(4008, (uintptr)(result.response), err)
	}
	return (uintptr)(result.response), err
}

func (g2config *G2configImpl) DeleteDataSource(ctx context.Context, configHandle uintptr, inputJson string) error {
	// _DLEXPORT int G2Config_deleteDataSource(ConfigHandle configHandle, const char *inputJson);
	if g2config.isTrace {
		g2config.traceEntry(4009, configHandle, inputJson)
	}
	var err error = nil
	inputJsonForC := C.CString(inputJson)
	defer C.free(unsafe.Pointer(inputJsonForC))
	result := C.G2Config_deleteDataSource_helper(C.uintptr_t(configHandle), inputJsonForC)
	if result != 0 {
		err = g2config.newError(ctx, 2004, configHandle, inputJson, result)
	}
	if g2config.isTrace {
		defer g2config.traceExit(4010, configHandle, inputJson, err)
	}
	return err
}

func (g2config *G2configImpl) Destroy(ctx context.Context) error {
	// _DLEXPORT int G2Config_destroy();
	if g2config.isTrace {
		g2config.traceEntry(4011)
	}
	var err error = nil
	result := C.G2Config_destroy()
	if result != 0 {
		err = g2config.newError(ctx, 2005, result)
	}
	if g2config.isTrace {
		defer g2config.traceExit(4012, err)
	}
	return err
}

// GetLastException returns the last exception encountered in the Senzing Engine.
func (g2config *G2configImpl) GetLastException(ctx context.Context) (string, error) {
	// _DLEXPORT int G2Config_getLastException(char *buffer, const size_t bufSize);
	if g2config.isTrace {
		g2config.traceEntry(4013)
	}
	var err error = nil
	stringBuffer := g2config.getByteArray(initialByteArraySize)
	result := C.G2Config_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	if result != 0 {
		err = g2config.newError(ctx, 2005, result)
	}
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if g2config.isTrace {
		defer g2config.traceExit(4014, string(stringBuffer), err)
	}
	return string(stringBuffer), err
}

func (g2config *G2configImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2Config_getLastExceptionCode();
	if g2config.isTrace {
		g2config.traceEntry(4015)
	}
	var err error = nil
	result := int(C.G2Config_getLastExceptionCode())
	if g2config.isTrace {
		defer g2config.traceExit(4016, result, err)
	}
	return result, err
}

func (g2config *G2configImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2Config_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	if g2config.isTrace {
		g2config.traceEntry(4017, moduleName, iniParams, verboseLogging)
	}
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2Config_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))
	if result != 0 {
		err = g2config.newError(ctx, 2007, moduleName, iniParams, verboseLogging, result)
	}
	if g2config.isTrace {
		defer g2config.traceExit(4018, moduleName, iniParams, verboseLogging, err)
	}
	return err
}

func (g2config *G2configImpl) ListDataSources(ctx context.Context, configHandle uintptr) (string, error) {
	// _DLEXPORT int G2Config_listDataSources(ConfigHandle configHandle, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2config.isTrace {
		g2config.traceEntry(4019, configHandle)
	}
	var err error = nil
	result := C.G2Config_listDataSources_helper(C.uintptr_t(configHandle))
	if result.returnCode != 0 {
		err = g2config.newError(ctx, 2008, result.returnCode, result)
	}
	if g2config.isTrace {
		defer g2config.traceExit(4020, configHandle, C.GoString(result.response), err)
	}
	return C.GoString(result.response), err
}

func (g2config *G2configImpl) Load(ctx context.Context, configHandle uintptr, jsonConfig string) error {
	// _DLEXPORT int G2Config_load(const char *jsonConfig,ConfigHandle* configHandle);
	if g2config.isTrace {
		g2config.traceEntry(4021, configHandle, jsonConfig)
	}
	var err error = nil
	jsonConfigForC := C.CString(jsonConfig)
	defer C.free(unsafe.Pointer(jsonConfigForC))
	result := C.G2Config_load_helper(C.uintptr_t(configHandle), jsonConfigForC)
	if result != 0 {
		err = g2config.newError(ctx, 2009, configHandle, jsonConfig, result)
	}
	if g2config.isTrace {
		defer g2config.traceExit(4022, configHandle, jsonConfig, err)
	}
	return err
}

func (g2config *G2configImpl) Save(ctx context.Context, configHandle uintptr) (string, error) {
	// _DLEXPORT int G2Config_save(ConfigHandle configHandle, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2config.isTrace {
		g2config.traceEntry(4023, configHandle)
	}
	var err error = nil
	result := C.G2Config_save_helper(C.uintptr_t(configHandle))
	if result.returnCode != 0 {
		err = g2config.newError(ctx, 2010, configHandle, result.returnCode, result)
	}
	if g2config.isTrace {
		defer g2config.traceExit(4024, configHandle, C.GoString(result.response), err)
	}
	return C.GoString(result.response), err
}

func (g2config *G2configImpl) SetLogLevel(ctx context.Context, logLevel logger.Level) error {
	if g2config.isTrace {
		g2config.traceEntry(4025, logLevel)
	}
	var err error = nil
	g2config.getLogger().SetLogLevel(messagelogger.Level(logLevel))

	if g2config.getLogger().GetLogLevel() == messagelogger.LevelTrace {
		g2config.isTrace = true
	} else {
		g2config.isTrace = false
	}
	if g2config.isTrace {
		defer g2config.traceExit(4026, logLevel, err)
	}
	return err
}
