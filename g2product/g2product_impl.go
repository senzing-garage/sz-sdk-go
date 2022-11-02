/*
The G2productImpl implementation...
*/
package g2product

/*
#include "g2product.h"
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

type G2productImpl struct {
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
func (g2product *G2productImpl) getByteArrayC(size int) *C.char {
	bytes := C.malloc(C.size_t(size))
	return (*C.char)(bytes)
}

func (g2product *G2productImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

func (g2product *G2productImpl) getError(ctx context.Context, errorNumber int, details ...interface{}) error {
	lastException, err := g2product.GetLastException(ctx)
	defer g2product.ClearLastException(ctx)
	message := lastException
	if err != nil {
		message = err.Error()
	}

	var newDetails []interface{}
	newDetails = append(newDetails, details...)
	newDetails = append(newDetails, errors.New(message))
	logger := g2product.getLogger(ctx)
	errorMessage, err := logger.Message(errorNumber, newDetails...)
	if err != nil {
		errorMessage = err.Error()
	}

	return errors.New(errorMessage)
}

func (g2product *G2productImpl) getLogger(ctx context.Context) messagelogger.MessageLoggerInterface {
	if g2product.logger == nil {
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
		g2product.logger, _ = messagelogger.New(messageFormat, messageId, messageLogLevel, messageStatus, messageText, messagelogger.LevelInfo)
	}
	return g2product.logger
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2product *G2productImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2Config_clearLastException();
	var err error = nil
	C.G2Product_clearLastException()
	return err
}

func (g2product *G2productImpl) Destroy(ctx context.Context) error {
	// _DLEXPORT int G2Config_destroy();
	var err error = nil
	result := C.G2Product_destroy()
	if result != 0 {
		err = g2product.getError(ctx, 1, result)
	}
	return err
}

// GetLastException returns the last exception encountered in the Senzing Engine.
func (g2product *G2productImpl) GetLastException(ctx context.Context) (string, error) {
	// _DLEXPORT int G2Config_getLastException(char *buffer, const size_t bufSize);
	var err error = nil
	stringBuffer := g2product.getByteArray(initialByteArraySize)
	C.G2Product_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if len(stringBuffer) == 0 {
		logger := g2product.getLogger(ctx)
		err = logger.Error(2999)
	}
	return string(stringBuffer), err
}

func (g2product *G2productImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2Config_getLastExceptionCode();
	var err error = nil
	result := C.G2Product_getLastExceptionCode()
	return int(result), err
}

func (g2product *G2productImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2Config_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2Product_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))
	if result != 0 {
		err = g2product.getError(ctx, 2, moduleName, iniParams, verboseLogging, result)
	}
	return err
}

func (g2product *G2productImpl) License(ctx context.Context) (string, error) {
	// _DLEXPORT char* G2Product_license();
	var err error = nil
	result := C.G2Product_license()
	return C.GoString(result), err
}

func (g2product *G2productImpl) ValidateLicenseFile(ctx context.Context, licenseFilePath string) (string, error) {
	// _DLEXPORT int G2Product_validateLicenseFile(const char* licenseFilePath, char **errorBuf, size_t *errorBufSize, void *(*resizeFunc)(void *ptr,size_t newSize));
	var err error = nil
	licenseFilePathForC := C.CString(licenseFilePath)
	defer C.free(unsafe.Pointer(licenseFilePathForC))
	result := C.G2Product_validateLicenseFile_helper(licenseFilePathForC)
	if result.returnCode != 0 {
		err = g2product.getError(ctx, 3, licenseFilePath, result.returnCode, result)
	}
	return C.GoString(result.response), err
}

func (g2product *G2productImpl) ValidateLicenseStringBase64(ctx context.Context, licenseString string) (string, error) {
	// _DLEXPORT int G2Product_validateLicenseStringBase64(const char* licenseString, char **errorBuf, size_t *errorBufSize, void *(*resizeFunc)(void *ptr,size_t newSize));
	var err error = nil
	licenseStringForC := C.CString(licenseString)
	defer C.free(unsafe.Pointer(licenseStringForC))
	result := C.G2Product_validateLicenseStringBase64_helper(licenseStringForC)
	if result.returnCode != 0 {
		err = g2product.getError(ctx, 4, licenseString, result.returnCode, result)
	}
	return C.GoString(result.response), err
}

func (g2product *G2productImpl) Version(ctx context.Context) (string, error) {
	// _DLEXPORT char* G2Product_license();
	var err error = nil
	result := C.G2Product_version()
	return C.GoString(result), err
}
