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
	"time"
	"unsafe"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagelevel"
	"github.com/senzing/go-logging/messagelocation"
	"github.com/senzing/go-logging/messagelogger"
	"github.com/senzing/go-logging/messagestatus"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2productImpl struct {
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
func (g2product *G2productImpl) getByteArrayC(size int) *C.char {
	bytes := C.malloc(C.size_t(size))
	return (*C.char)(bytes)
}

func (g2product *G2productImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

func (g2product *G2productImpl) newError(ctx context.Context, errorNumber int, details ...interface{}) error {
	lastException, err := g2product.GetLastException(ctx)
	defer g2product.ClearLastException(ctx)
	message := lastException
	if err != nil {
		message = err.Error()
	}

	var newDetails []interface{}
	newDetails = append(newDetails, details...)
	newDetails = append(newDetails, errors.New(message))
	messageGenerator := g2product.getMessageGenerator()
	errorMessage, err := messageGenerator.Message(errorNumber, newDetails...)
	if err != nil {
		errorMessage = err.Error()
	}

	return errors.New(errorMessage)
}

func (g2product *G2productImpl) getLogger() messagelogger.MessageLoggerInterface {
	if g2product.logger == nil {
		messageLevel := &messagelevel.MessageLevelSenzingApi{
			IdRanges:   IdRanges,
			IdStatuses: IdStatuses,
		}

		messageStatus := &messagestatus.MessageStatusSenzingApi{
			IdRanges: IdRanges,
		}

		messageLocation := &messagelocation.MessageLocationSenzing{
			CallerSkip: 4,
		}

		g2product.logger, _ = messagelogger.NewSenzingLogger(ProductId, IdMessages, messageLevel, messageStatus, messageLocation, messagelogger.LevelInfo)
	}
	return g2product.logger
}

func (g2product *G2productImpl) getMessageGenerator() messagelogger.MessageLoggerInterface {
	if g2product.messageGenerator == nil {
		messageLevel := &messagelevel.MessageLevelSenzingApi{
			IdRanges:   IdRanges,
			IdStatuses: IdStatuses,
		}

		messageStatus := &messagestatus.MessageStatusSenzingApi{
			IdRanges: IdRanges,
		}

		messageLocation := &messagelocation.MessageLocationSenzing{
			CallerSkip: 4,
		}

		g2product.messageGenerator, _ = messagelogger.NewSenzingLogger(ProductId, IdMessages, messageLevel, messageStatus, messageLocation, messagelogger.LevelInfo)
	}
	return g2product.messageGenerator
}

func (g2product *G2productImpl) traceEntry(errorNumber int, details ...interface{}) {
	g2product.getLogger().Log(errorNumber, details...)
}

func (g2product *G2productImpl) traceExit(errorNumber int, details ...interface{}) {
	g2product.getLogger().Log(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2product *G2productImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2Config_clearLastException();
	if g2product.isTrace {
		g2product.traceEntry(1)
	}
	entryTime := time.Now()
	var err error = nil
	C.G2Product_clearLastException()
	if g2product.isTrace {
		defer g2product.traceExit(2, err, time.Since(entryTime))
	}
	return err
}

func (g2product *G2productImpl) Destroy(ctx context.Context) error {
	// _DLEXPORT int G2Config_destroy();
	if g2product.isTrace {
		g2product.traceEntry(3)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Product_destroy()
	if result != 0 {
		err = g2product.newError(ctx, 4001, result, time.Since(entryTime))
	}
	if g2product.isTrace {
		defer g2product.traceExit(4, err, time.Since(entryTime))
	}
	return err
}

// GetLastException returns the last exception encountered in the Senzing Engine.
func (g2product *G2productImpl) GetLastException(ctx context.Context) (string, error) {
	// _DLEXPORT int G2Config_getLastException(char *buffer, const size_t bufSize);
	if g2product.isTrace {
		g2product.traceEntry(5)
	}
	entryTime := time.Now()
	var err error = nil
	stringBuffer := g2product.getByteArray(initialByteArraySize)
	result := C.G2Product_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	if result == 0 {
		messageGenerator := g2product.getMessageGenerator()
		err = messageGenerator.Error(4002, result, time.Since(entryTime))
	}
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if g2product.isTrace {
		defer g2product.traceExit(6, string(stringBuffer), err, time.Since(entryTime))
	}
	return string(stringBuffer), err
}

func (g2product *G2productImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2Config_getLastExceptionCode();
	if g2product.isTrace {
		g2product.traceEntry(7)
	}
	entryTime := time.Now()
	var err error = nil
	result := int(C.G2Product_getLastExceptionCode())
	if g2product.isTrace {
		defer g2product.traceExit(8, result, err, time.Since(entryTime))
	}
	return result, err
}

func (g2product *G2productImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2Config_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	if g2product.isTrace {
		g2product.traceEntry(9, moduleName, iniParams, verboseLogging)
	}
	entryTime := time.Now()
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2Product_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))
	if result != 0 {
		err = g2product.newError(ctx, 4003, moduleName, iniParams, verboseLogging, result, time.Since(entryTime))
	}
	if g2product.isTrace {
		defer g2product.traceExit(10, moduleName, iniParams, verboseLogging, err, time.Since(entryTime))
	}
	return err
}

func (g2product *G2productImpl) License(ctx context.Context) (string, error) {
	// _DLEXPORT char* G2Product_license();
	if g2product.isTrace {
		g2product.traceEntry(11)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Product_license()
	if g2product.isTrace {
		defer g2product.traceExit(12, C.GoString(result), err, time.Since(entryTime))
	}
	return C.GoString(result), err
}

func (g2product *G2productImpl) SetLogLevel(ctx context.Context, logLevel logger.Level) error {
	if g2product.isTrace {
		g2product.traceEntry(13, logLevel)
	}
	entryTime := time.Now()
	var err error = nil
	g2product.getLogger().SetLogLevel(messagelogger.Level(logLevel))
	g2product.isTrace = g2product.getLogger().GetLogLevel() == messagelogger.LevelTrace
	if g2product.isTrace {
		defer g2product.traceExit(14, logLevel, err, time.Since(entryTime))
	}
	return err
}

func (g2product *G2productImpl) ValidateLicenseFile(ctx context.Context, licenseFilePath string) (string, error) {
	// _DLEXPORT int G2Product_validateLicenseFile(const char* licenseFilePath, char **errorBuf, size_t *errorBufSize, void *(*resizeFunc)(void *ptr,size_t newSize));
	if g2product.isTrace {
		g2product.traceEntry(15, licenseFilePath)
	}
	entryTime := time.Now()
	var err error = nil
	licenseFilePathForC := C.CString(licenseFilePath)
	defer C.free(unsafe.Pointer(licenseFilePathForC))
	result := C.G2Product_validateLicenseFile_helper(licenseFilePathForC)
	if result.returnCode != 0 {
		err = g2product.newError(ctx, 4004, licenseFilePath, result.returnCode, result, time.Since(entryTime))
	}
	if g2product.isTrace {
		defer g2product.traceExit(16, licenseFilePath, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2product *G2productImpl) ValidateLicenseStringBase64(ctx context.Context, licenseString string) (string, error) {
	// _DLEXPORT int G2Product_validateLicenseStringBase64(const char* licenseString, char **errorBuf, size_t *errorBufSize, void *(*resizeFunc)(void *ptr,size_t newSize));
	if g2product.isTrace {
		g2product.traceEntry(17, licenseString)
	}
	entryTime := time.Now()
	var err error = nil
	licenseStringForC := C.CString(licenseString)
	defer C.free(unsafe.Pointer(licenseStringForC))
	result := C.G2Product_validateLicenseStringBase64_helper(licenseStringForC)
	if result.returnCode != 0 {
		err = g2product.newError(ctx, 4005, licenseString, result.returnCode, result, time.Since(entryTime))
	}
	if g2product.isTrace {
		defer g2product.traceExit(18, licenseString, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2product *G2productImpl) Version(ctx context.Context) (string, error) {
	// _DLEXPORT char* G2Product_license();
	if g2product.isTrace {
		g2product.traceEntry(19)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Product_version()
	if g2product.isTrace {
		defer g2product.traceExit(20, C.GoString(result), err, time.Since(entryTime))
	}
	return C.GoString(result), err
}
