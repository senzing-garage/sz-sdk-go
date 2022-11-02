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
	logger messagelogger.MessageLoggerInterface
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

func (g2config *G2configImpl) getError(ctx context.Context, errorNumber int, details ...interface{}) error {
	lastException, err := g2config.GetLastException(ctx)
	defer g2config.ClearLastException(ctx)
	message := lastException
	if err != nil {
		message = err.Error()
	}

	var newDetails []interface{}
	newDetails = append(newDetails, details...)
	newDetails = append(newDetails, errors.New(message))
	logger := g2config.getLogger(ctx)
	errorMessage, err := logger.Message(errorNumber, newDetails...)
	if err != nil {
		errorMessage = err.Error()
	}

	return errors.New(errorMessage)
}

func (g2config *G2configImpl) getLogger(ctx context.Context) messagelogger.MessageLoggerInterface {
	if g2config.logger == nil {
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
		g2config.logger, _ = messagelogger.New(messageFormat, messageId, messageLogLevel, messageStatus, messageText, messagelogger.LevelInfo)
	}
	return g2config.logger
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func (g2config *G2configImpl) AddDataSource(ctx context.Context, configHandle uintptr, inputJson string) (string, error) {
	// _DLEXPORT int G2Config_addDataSource(ConfigHandle configHandle, const char *inputJson, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	inputJsonForC := C.CString(inputJson)
	defer C.free(unsafe.Pointer(inputJsonForC))
	result := C.G2Config_addDataSource_helper(C.uintptr_t(configHandle), inputJsonForC)
	if result.returnCode != 0 {
		err = g2config.getError(ctx, 1, inputJson, result.returnCode, result)
	}
	return C.GoString(result.response), err
}

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2config *G2configImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2Config_clearLastException();
	var err error = nil
	C.G2Config_clearLastException()
	return err
}

func (g2config *G2configImpl) Close(ctx context.Context, configHandle uintptr) error {
	// _DLEXPORT int G2Config_close(ConfigHandle configHandle);
	var err error = nil
	result := C.G2config_close_helper(C.uintptr_t(configHandle))
	if result != 0 {
		err = g2config.getError(ctx, 2, result)
	}
	return err
}

func (g2config *G2configImpl) Create(ctx context.Context) (uintptr, error) {
	// _DLEXPORT int G2Config_create(ConfigHandle* configHandle);
	var err error = nil
	result := C.G2config_create_helper()
	returnCode := 0 // FIXME:
	if result == nil {
		err = g2config.getError(ctx, 3, returnCode)
	}
	return (uintptr)(result), err
}

func (g2config *G2configImpl) DeleteDataSource(ctx context.Context, configHandle uintptr, inputJson string) error {
	// _DLEXPORT int G2Config_deleteDataSource(ConfigHandle configHandle, const char *inputJson);
	var err error = nil
	inputJsonForC := C.CString(inputJson)
	defer C.free(unsafe.Pointer(inputJsonForC))
	result := C.G2Config_deleteDataSource_helper(C.uintptr_t(configHandle), inputJsonForC)
	if result != 0 {
		err = g2config.getError(ctx, 4, inputJson, result)
	}
	return err
}

func (g2config *G2configImpl) Destroy(ctx context.Context) error {
	// _DLEXPORT int G2Config_destroy();
	var err error = nil
	result := C.G2Config_destroy()
	if result != 0 {
		err = g2config.getError(ctx, 5, result)
	}
	return err
}

// GetLastException returns the last exception encountered in the Senzing Engine.
func (g2config *G2configImpl) GetLastException(ctx context.Context) (string, error) {
	// _DLEXPORT int G2Config_getLastException(char *buffer, const size_t bufSize);
	var err error = nil
	stringBuffer := g2config.getByteArray(initialByteArraySize)
	C.G2Config_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if len(stringBuffer) == 0 {
		logger := g2config.getLogger(ctx)
		err = logger.Error(2999)
	}
	return string(stringBuffer), err
}

func (g2config *G2configImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2Config_getLastExceptionCode();
	var err error = nil
	result := C.G2Config_getLastExceptionCode()
	return int(result), err
}

func (g2config *G2configImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2Config_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2Config_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))
	if result != 0 {
		err = g2config.getError(ctx, 6, moduleName, iniParams, verboseLogging, result)
	}
	return err
}

func (g2config *G2configImpl) ListDataSources(ctx context.Context, configHandle uintptr) (string, error) {
	// _DLEXPORT int G2Config_listDataSources(ConfigHandle configHandle, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	result := C.G2Config_listDataSources_helper(C.uintptr_t(configHandle))
	if result.returnCode != 0 {
		err = g2config.getError(ctx, 7, result.returnCode, result)
	}
	return C.GoString(result.response), err
}

func (g2config *G2configImpl) Load(ctx context.Context, configHandle uintptr, jsonConfig string) error {
	// _DLEXPORT int G2Config_load(const char *jsonConfig,ConfigHandle* configHandle);
	var err error = nil
	jsonConfigForC := C.CString(jsonConfig)
	defer C.free(unsafe.Pointer(jsonConfigForC))
	result := C.G2Config_load_helper(C.uintptr_t(configHandle), jsonConfigForC)
	if result != 0 {
		err = g2config.getError(ctx, 8, jsonConfig, result)
	}
	return err
}

func (g2config *G2configImpl) Save(ctx context.Context, configHandle uintptr) (string, error) {
	// _DLEXPORT int G2Config_save(ConfigHandle configHandle, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	result := C.G2Config_save_helper(C.uintptr_t(configHandle))
	if result.returnCode != 0 {
		err = g2config.getError(ctx, 9, result.returnCode, result)
	}
	return C.GoString(result.response), err
}
