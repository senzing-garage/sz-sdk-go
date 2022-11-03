/*
The G2configmgrImpl implementation...
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

type G2configmgrImpl struct {
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
func (g2configmgr *G2configmgrImpl) getByteArrayC(size int) *C.char {
	bytes := C.malloc(C.size_t(size))
	return (*C.char)(bytes)
}

func (g2configmgr *G2configmgrImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

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
	messageGenerator := g2configmgr.getMessageGenerator(ctx)
	errorMessage, err := messageGenerator.Message(errorNumber, newDetails...)
	if err != nil {
		errorMessage = err.Error()
	}

	return errors.New(errorMessage)
}

func (g2configmgr *G2configmgrImpl) getLogger() messagelogger.MessageLoggerInterface {
	if g2configmgr.logger == nil {
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
		g2configmgr.logger, _ = messagelogger.New(messageFormat, messageId, messageLogLevel, messageStatus, messageText, messagelogger.LevelInfo)
	}
	return g2configmgr.logger
}

func (g2configmgr *G2configmgrImpl) getMessageGenerator(ctx context.Context) messagelogger.MessageLoggerInterface {
	if g2configmgr.messageGenerator == nil {
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
		g2configmgr.messageGenerator, _ = messagelogger.New(messageFormat, messageId, messageLogLevel, messageStatus, messageText, messagelogger.LevelInfo)
	}
	return g2configmgr.messageGenerator
}

func (g2configmgr *G2configmgrImpl) traceEntry(errorNumber int, details ...interface{}) {
	g2configmgr.getLogger().Log(errorNumber, details...)
}

func (g2configmgr *G2configmgrImpl) traceExit(errorNumber int, details ...interface{}) {
	g2configmgr.getLogger().Log(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func (g2configmgr *G2configmgrImpl) AddConfig(ctx context.Context, configStr string, configComments string) (int64, error) {
	// _DLEXPORT int G2ConfigMgr_addConfig(const char* configStr, const char* configComments, long long* configID);
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(4001, configStr, configComments)
	}
	var err error = nil
	configStrForC := C.CString(configStr)
	defer C.free(unsafe.Pointer(configStrForC))
	configCommentsForC := C.CString(configComments)
	defer C.free(unsafe.Pointer(configCommentsForC))
	result := C.G2ConfigMgr_addConfig_helper(configStrForC, configCommentsForC)
	if result.returnCode != 0 {
		err = g2configmgr.newError(ctx, 2001, configStr, configComments, result.returnCode, result)
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(4002, configStr, configComments, int64(C.longlong(result.configID)), err)
	}
	return int64(C.longlong(result.configID)), err
}

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2configmgr *G2configmgrImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2Config_clearLastException();
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(4003)
	}
	var err error = nil
	C.G2ConfigMgr_clearLastException()
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(4004, err)
	}
	return err
}

func (g2configmgr *G2configmgrImpl) Destroy(ctx context.Context) error {
	// _DLEXPORT int G2Config_destroy();
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(4005)
	}
	var err error = nil
	result := C.G2ConfigMgr_destroy()
	if result != 0 {
		err = g2configmgr.newError(ctx, 2002, result)
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(4006, err)
	}
	return err
}

func (g2configmgr *G2configmgrImpl) GetConfig(ctx context.Context, configID int64) (string, error) {
	// _DLEXPORT int G2ConfigMgr_getConfig(const long long configID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(4007, configID)
	}
	var err error = nil
	result := C.G2ConfigMgr_getConfig_helper(C.longlong(configID))
	if result.returnCode != 0 {
		err = g2configmgr.newError(ctx, 2003, configID, result.returnCode, result)
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(4008, configID, C.GoString(result.config), err)
	}
	return C.GoString(result.config), err
}

func (g2configmgr *G2configmgrImpl) GetConfigList(ctx context.Context) (string, error) {
	// _DLEXPORT int G2ConfigMgr_getConfigList(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(4009)
	}
	var err error = nil
	result := C.G2ConfigMgr_getConfigList_helper()
	if result.returnCode != 0 {
		err = g2configmgr.newError(ctx, 2004, result.returnCode, result)
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(4010, C.GoString(result.configList), err)
	}
	return C.GoString(result.configList), err
}

func (g2configmgr *G2configmgrImpl) GetDefaultConfigID(ctx context.Context) (int64, error) {
	//  _DLEXPORT int G2ConfigMgr_getDefaultConfigID(long long* configID);
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(4011)
	}
	var err error = nil
	result := C.G2ConfigMgr_getDefaultConfigID_helper()
	if result.returnCode != 0 {
		err = g2configmgr.newError(ctx, 2005, result.returnCode, result)
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(4012, int64(C.longlong(result.configID)), err)
	}
	return int64(C.longlong(result.configID)), err
}

// GetLastException returns the last exception encountered in the Senzing Engine.
func (g2configmgr *G2configmgrImpl) GetLastException(ctx context.Context) (string, error) {
	// _DLEXPORT int G2Config_getLastException(char *buffer, const size_t bufSize);
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(4013)
	}
	var err error = nil
	stringBuffer := g2configmgr.getByteArray(initialByteArraySize)
	C.G2ConfigMgr_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if len(stringBuffer) == 0 {
		messageGenerator := g2configmgr.getMessageGenerator(ctx)
		err = messageGenerator.Error(2999)
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(4014)
	}
	return string(stringBuffer), err
}

func (g2configmgr *G2configmgrImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2Config_getLastExceptionCode();
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(4015)
	}
	var err error = nil
	result := int(C.G2ConfigMgr_getLastExceptionCode())
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(4016, result, err)
	}
	return result, err
}

func (g2configmgr *G2configmgrImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2Config_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(4017, moduleName, iniParams, verboseLogging)
	}
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2ConfigMgr_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))
	if result != 0 {
		err = g2configmgr.newError(ctx, 2006, moduleName, iniParams, verboseLogging, result)
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(4018, moduleName, iniParams, verboseLogging, err)
	}
	return err
}

// Very much like a "compare-and-swap" instruction to serialize concurrent editing of configuration.
// To simply set the default configuration ID, use SetDefaultConfigID().
func (g2configmgr *G2configmgrImpl) ReplaceDefaultConfigID(ctx context.Context, oldConfigID int64, newConfigID int64) error {
	// _DLEXPORT int G2ConfigMgr_replaceDefaultConfigID(const long long oldConfigID, const long long newConfigID);
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(4019, oldConfigID, newConfigID)
	}
	var err error = nil
	result := C.G2ConfigMgr_replaceDefaultConfigID(C.longlong(oldConfigID), C.longlong(newConfigID))
	if result != 0 {
		err = g2configmgr.newError(ctx, 2007, oldConfigID, newConfigID, result)
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(4020, oldConfigID, newConfigID, err)
	}
	return err
}

func (g2configmgr *G2configmgrImpl) SetDefaultConfigID(ctx context.Context, configID int64) error {
	// _DLEXPORT int G2ConfigMgr_setDefaultConfigID(const long long configID);
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(4021, configID)
	}
	var err error = nil
	result := C.G2ConfigMgr_setDefaultConfigID(C.longlong(configID))
	if result != 0 {
		err = g2configmgr.newError(ctx, 2008, configID, result)
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(4022, configID, err)
	}
	return err
}

func (g2configmgr *G2configmgrImpl) SetLogLevel(ctx context.Context, logLevel logger.Level) error {
	if g2configmgr.isTrace {
		g2configmgr.traceEntry(4023, logLevel)
	}
	var err error = nil
	g2configmgr.getLogger().SetLogLevel(messagelogger.Level(logLevel))

	if g2configmgr.getLogger().GetLogLevel() == messagelogger.LevelTrace {
		g2configmgr.isTrace = true
	} else {
		g2configmgr.isTrace = false
	}
	if g2configmgr.isTrace {
		defer g2configmgr.traceExit(4024, logLevel, err)
	}
	return err
}
