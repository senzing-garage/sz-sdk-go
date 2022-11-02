/*
The G2diagnosticImpl implementation...
*/
package g2diagnostic

/*
#include "g2diagnostic.h"
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

type G2diagnosticImpl struct {
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
func (g2diagnostic *G2diagnosticImpl) getByteArrayC(size int) *C.char {
	bytes := C.malloc(C.size_t(size))
	return (*C.char)(bytes)
}

func (g2diagnostic *G2diagnosticImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

func (g2diagnostic *G2diagnosticImpl) getError(ctx context.Context, errorNumber int, details ...interface{}) error {
	lastException, err := g2diagnostic.GetLastException(ctx)
	defer g2diagnostic.ClearLastException(ctx)
	message := lastException
	if err != nil {
		message = err.Error()
	}

	var newDetails []interface{}
	newDetails = append(newDetails, details...)
	newDetails = append(newDetails, errors.New(message))
	logger := g2diagnostic.getLogger(ctx)
	errorMessage, err := logger.Message(errorNumber, newDetails...)
	if err != nil {
		errorMessage = err.Error()
	}

	return errors.New(errorMessage)
}

func (g2diagnostic *G2diagnosticImpl) getLogger(ctx context.Context) messagelogger.MessageLoggerInterface {
	if g2diagnostic.logger == nil {
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
		g2diagnostic.logger, _ = messagelogger.New(messageFormat, messageId, messageLogLevel, messageStatus, messageText, messagelogger.LevelInfo)
	}
	return g2diagnostic.logger
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// CheckDBPerf performs inserts to determine rate of insertion.
func (g2diagnostic *G2diagnosticImpl) CheckDBPerf(ctx context.Context, secondsToRun int) (string, error) {
	// _DLEXPORT int G2Diagnostic_checkDBPerf(int secondsToRun, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_checkDBPerf_helper(C.int(secondsToRun)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 1, secondsToRun, returnCode)
	}
	return stringBuffer, err
}

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2Diagnostic_clearLastException();
	var err error = nil
	C.G2Diagnostic_clearLastException()
	return err
}

func (g2diagnostic *G2diagnosticImpl) CloseEntityListBySize(ctx context.Context, entityListBySizeHandle uintptr) error {
	//  _DLEXPORT int G2Diagnostic_closeEntityListBySize(EntityListBySizeHandle entityListBySizeHandle);
	var err error = nil
	result := C.G2Diagnostic_closeEntityListBySize_helper(C.uintptr_t(entityListBySizeHandle))
	if result != 0 {
		err = g2diagnostic.getError(ctx, 2, result)
	}
	return err
}

func (g2diagnostic *G2diagnosticImpl) Destroy(ctx context.Context) error {
	//  _DLEXPORT int G2Diagnostic_destroy();
	var err error = nil
	result := C.G2Diagnostic_destroy()
	if result != 0 {
		err = g2diagnostic.getError(ctx, 3, result)
	}
	return err
}

func (g2diagnostic *G2diagnosticImpl) FetchNextEntityBySize(ctx context.Context, entityListBySizeHandle uintptr) (string, error) {
	//  _DLEXPORT int G2Diagnostic_fetchNextEntityBySize(EntityListBySizeHandle entityListBySizeHandle, char *responseBuf, const size_t bufSize);
	var err error = nil
	stringBuffer := g2diagnostic.getByteArray(initialByteArraySize)
	result := C.G2Diagnostic_fetchNextEntityBySize_helper(C.uintptr_t(entityListBySizeHandle), (*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	if result != 0 {
		err = g2diagnostic.getError(ctx, 4, result)
	}
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	return string(stringBuffer), err
}

func (g2diagnostic *G2diagnosticImpl) FindEntitiesByFeatureIDs(ctx context.Context, features string) (string, error) {
	//  _DLEXPORT int G2Diagnostic_findEntitiesByFeatureIDs(const char *features, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	featuresForC := C.CString(features)
	defer C.free(unsafe.Pointer(featuresForC))
	stringBuffer := C.GoString(C.G2Diagnostic_findEntitiesByFeatureIDs_helper(featuresForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 5, features, returnCode)
	}
	return stringBuffer, err
}

// GetAvailableMemory returns the available memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) GetAvailableMemory(ctx context.Context) (int64, error) {
	// _DLEXPORT long long G2Diagnostic_getAvailableMemory();
	var err error = nil
	result := C.G2Diagnostic_getAvailableMemory()
	return int64(result), err
}

func (g2diagnostic *G2diagnosticImpl) GetDataSourceCounts(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getDataSourceCounts(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getDataSourceCounts_helper())
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 6, returnCode)
	}
	return stringBuffer, err
}

// GetDBInfo returns information about the database connection.
func (g2diagnostic *G2diagnosticImpl) GetDBInfo(ctx context.Context) (string, error) {
	// _DLEXPORT int G2Diagnostic_getDBInfo(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getDBInfo_helper())
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 7, returnCode)
	}
	return stringBuffer, err
}

func (g2diagnostic *G2diagnosticImpl) GetEntityDetails(ctx context.Context, entityID int64, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getEntityDetails(const long long entityID, const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getEntityDetails_helper(C.longlong(entityID), C.int(includeInternalFeatures)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 8, entityID, includeInternalFeatures, returnCode)
	}
	return stringBuffer, err
}

func (g2diagnostic *G2diagnosticImpl) GetEntityListBySize(ctx context.Context, entitySize int) (uintptr, error) {
	//  _DLEXPORT int G2Diagnostic_getEntityListBySize(const size_t entitySize, EntityListBySizeHandle* entityListBySizeHandle);
	var err error = nil
	result := C.G2Diagnostic_getEntityListBySize_helper(C.size_t(entitySize))
	returnCode := 0 // FIXME:
	if result == nil {
		err = g2diagnostic.getError(ctx, 9, entitySize, returnCode)
	}
	return (uintptr)(result), err
}

func (g2diagnostic *G2diagnosticImpl) GetEntityResume(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getEntityResume(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getEntityResume_helper(C.longlong(entityID)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 10, entityID, returnCode)
	}
	return stringBuffer, err
}

func (g2diagnostic *G2diagnosticImpl) GetEntitySizeBreakdown(ctx context.Context, minimumEntitySize int, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getEntitySizeBreakdown(const size_t minimumEntitySize, const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getEntitySizeBreakdown_helper(C.size_t(minimumEntitySize), C.int(includeInternalFeatures)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 11, minimumEntitySize, includeInternalFeatures, returnCode)
	}
	return stringBuffer, err
}

func (g2diagnostic *G2diagnosticImpl) GetFeature(ctx context.Context, libFeatID int64) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getFeature(const long long libFeatID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getFeature_helper(C.longlong(libFeatID)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 12, libFeatID, returnCode)
	}
	return stringBuffer, err
}

func (g2diagnostic *G2diagnosticImpl) GetGenericFeatures(ctx context.Context, featureType string, maximumEstimatedCount int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getGenericFeatures(const char* featureType, const size_t maximumEstimatedCount, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	featureTypeForC := C.CString(featureType)
	defer C.free(unsafe.Pointer(featureTypeForC))
	stringBuffer := C.GoString(C.G2Diagnostic_getGenericFeatures_helper(featureTypeForC, C.size_t(maximumEstimatedCount)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 13, featureType, maximumEstimatedCount, returnCode)
	}
	return stringBuffer, err
}

// GetLastException returns the last exception encountered in the Senzing Engine.
func (g2diagnostic *G2diagnosticImpl) GetLastException(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getLastException(char *buffer, const size_t bufSize);
	var err error = nil
	stringBuffer := g2diagnostic.getByteArray(initialByteArraySize)
	C.G2Diagnostic_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if len(stringBuffer) == 0 {
		logger := g2diagnostic.getLogger(ctx)
		err = logger.Error(2999)
	}
	return string(stringBuffer), err
}

func (g2diagnostic *G2diagnosticImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2Diagnostic_getLastExceptionCode();
	var err error = nil
	result := C.G2Diagnostic_getLastExceptionCode()
	return int(result), err
}

// GetLogicalCores returns the number of logical cores on the host system.
func (g2diagnostic *G2diagnosticImpl) GetLogicalCores(ctx context.Context) (int, error) {
	// _DLEXPORT int G2Diagnostic_getLogicalCores();
	var err error = nil
	result := C.G2Diagnostic_getLogicalCores()
	return int(result), err
}

func (g2diagnostic *G2diagnosticImpl) GetMappingStatistics(ctx context.Context, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getMappingStatistics(const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getMappingStatistics_helper(C.int(includeInternalFeatures)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 14, includeInternalFeatures, returnCode)
	}
	return stringBuffer, err
}

// GetPhysicalCores returns the number of physical cores on the host system.
func (g2diagnostic *G2diagnosticImpl) GetPhysicalCores(ctx context.Context) (int, error) {
	// _DLEXPORT int G2Diagnostic_getPhysicalCores();
	var err error = nil
	result := C.G2Diagnostic_getPhysicalCores()
	return int(result), err
}

func (g2diagnostic *G2diagnosticImpl) GetRelationshipDetails(ctx context.Context, relationshipID int64, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getRelationshipDetails(const long long relationshipID, const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getRelationshipDetails_helper(C.longlong(relationshipID), C.int(includeInternalFeatures)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 15, relationshipID, includeInternalFeatures, returnCode)
	}
	return stringBuffer, err
}

func (g2diagnostic *G2diagnosticImpl) GetResolutionStatistics(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getResolutionStatistics(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getResolutionStatistics_helper())
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 16, returnCode)
	}
	return stringBuffer, err
}

// GetTotalSystemMemory returns the total memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) GetTotalSystemMemory(ctx context.Context) (int64, error) {
	// _DLEXPORT long long G2Diagnostic_getTotalSystemMemory();
	var err error = nil
	result := C.G2Diagnostic_getTotalSystemMemory()
	return int64(result), err
}

// Init initializes the Senzing G2diagnosis.
func (g2diagnostic *G2diagnosticImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2Diagnostic_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2Diagnostic_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))
	if result != 0 {
		err = g2diagnostic.getError(ctx, 17, moduleName, iniParams, verboseLogging, result)
	}
	return err
}

func (g2diagnostic *G2diagnosticImpl) InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error {
	//  _DLEXPORT int G2Diagnostic_initWithConfigID(const char *moduleName, const char *iniParams, const long long initConfigID, const int verboseLogging);
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2Diagnostic_initWithConfigID(moduleNameForC, iniParamsForC, C.longlong(initConfigID), C.int(verboseLogging))
	if result != 0 {
		err = g2diagnostic.getError(ctx, 18, moduleName, iniParams, initConfigID, verboseLogging, result)
	}
	return err
}

// Null shows how to report a BUG inline.
func (g2diagnostic *G2diagnosticImpl) Null(ctx context.Context) (int64, error) {
	// BUG(mjd): Just an example of how to show bugs in GoDoc.
	var err error = nil
	return 1, err
}

func (g2diagnostic *G2diagnosticImpl) Reinit(ctx context.Context, initConfigID int64) error {
	//  _DLEXPORT int G2Diagnostic_reinit(const long long initConfigID);
	var err error = nil
	result := C.G2Diagnostic_reinit(C.longlong(initConfigID))
	if result != 0 {
		err = g2diagnostic.getError(ctx, 19, initConfigID, result)
	}
	return err
}
