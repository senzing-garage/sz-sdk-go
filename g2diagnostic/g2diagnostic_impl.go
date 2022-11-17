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

type G2diagnosticImpl struct {
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
func (g2diagnostic *G2diagnosticImpl) getByteArrayC(size int) *C.char {
	bytes := C.malloc(C.size_t(size))
	return (*C.char)(bytes)
}

func (g2diagnostic *G2diagnosticImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

func (g2diagnostic *G2diagnosticImpl) newError(ctx context.Context, errorNumber int, details ...interface{}) error {
	lastException, err := g2diagnostic.GetLastException(ctx)
	defer g2diagnostic.ClearLastException(ctx)
	message := lastException
	if err != nil {
		message = err.Error()
	}

	var newDetails []interface{}
	newDetails = append(newDetails, details...)
	newDetails = append(newDetails, errors.New(message))
	messageGenerator := g2diagnostic.getMessageGenerator()
	errorMessage, err := messageGenerator.Message(errorNumber, newDetails...)
	if err != nil {
		errorMessage = err.Error()
	}

	return errors.New(errorMessage)
}

func (g2diagnostic *G2diagnosticImpl) getLogger() messagelogger.MessageLoggerInterface {
	if g2diagnostic.logger == nil {
		messageLevel := &messagelevel.MessageLevelSenzingApi{
			IdLevelRanges: messagelevel.IdLevelRanges,
			IdStatuses:    messagestatus.IdLevelRangesAsString,
		}
		messageStatus := &messagestatus.MessageStatusSenzingApi{
			IdStatuses: messagestatus.IdLevelRangesAsString,
		}
		messageLocation := &messagelocation.MessageLocationSenzing{
			CallerSkip: 4,
		}
		g2diagnostic.logger, _ = messagelogger.NewSenzingLogger(ProductId, IdMessages, messageLevel, messageStatus, messageLocation, messagelogger.LevelInfo)
	}
	return g2diagnostic.logger
}

func (g2diagnostic *G2diagnosticImpl) getMessageGenerator() messagelogger.MessageLoggerInterface {
	if g2diagnostic.messageGenerator == nil {
		messageLevel := &messagelevel.MessageLevelSenzingApi{
			IdLevelRanges: messagelevel.IdLevelRanges,
			IdStatuses:    messagestatus.IdLevelRangesAsString,
		}
		messageStatus := &messagestatus.MessageStatusSenzingApi{
			IdStatuses: messagestatus.IdLevelRangesAsString,
		}
		messageLocation := &messagelocation.MessageLocationSenzing{
			CallerSkip: 4,
		}
		g2diagnostic.messageGenerator, _ = messagelogger.NewSenzingLogger(ProductId, IdMessages, messageLevel, messageStatus, messageLocation, messagelogger.LevelInfo)
	}
	return g2diagnostic.messageGenerator
}

func (g2diagnostic *G2diagnosticImpl) traceEntry(errorNumber int, details ...interface{}) {
	g2diagnostic.getLogger().Log(errorNumber, details...)
}

func (g2diagnostic *G2diagnosticImpl) traceExit(errorNumber int, details ...interface{}) {
	g2diagnostic.getLogger().Log(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// CheckDBPerf performs inserts to determine rate of insertion.
func (g2diagnostic *G2diagnosticImpl) CheckDBPerf(ctx context.Context, secondsToRun int) (string, error) {
	// _DLEXPORT int G2Diagnostic_checkDBPerf(int secondsToRun, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(1, secondsToRun)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Diagnostic_checkDBPerf_helper(C.int(secondsToRun))
	if result.returnCode != 0 {
		err = g2diagnostic.newError(ctx, 4001, secondsToRun, result.returnCode, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(2, secondsToRun, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err

}

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2Diagnostic_clearLastException();
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(3)
	}
	entryTime := time.Now()
	var err error = nil
	C.G2Diagnostic_clearLastException()
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(4, err, time.Since(entryTime))
	}
	return err
}

func (g2diagnostic *G2diagnosticImpl) CloseEntityListBySize(ctx context.Context, entityListBySizeHandle uintptr) error {
	//  _DLEXPORT int G2Diagnostic_closeEntityListBySize(EntityListBySizeHandle entityListBySizeHandle);
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(5)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Diagnostic_closeEntityListBySize_helper(C.uintptr_t(entityListBySizeHandle))
	if result != 0 {
		err = g2diagnostic.newError(ctx, 4002, result, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(6, err, time.Since(entryTime))
	}
	return err
}

func (g2diagnostic *G2diagnosticImpl) Destroy(ctx context.Context) error {
	//  _DLEXPORT int G2Diagnostic_destroy();
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(7)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Diagnostic_destroy()
	if result != 0 {
		err = g2diagnostic.newError(ctx, 4003, result, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(8, err, time.Since(entryTime))
	}
	return err
}

func (g2diagnostic *G2diagnosticImpl) FetchNextEntityBySize(ctx context.Context, entityListBySizeHandle uintptr) (string, error) {
	//  _DLEXPORT int G2Diagnostic_fetchNextEntityBySize(EntityListBySizeHandle entityListBySizeHandle, char *responseBuf, const size_t bufSize);
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(9)
	}
	entryTime := time.Now()
	var err error = nil
	stringBuffer := g2diagnostic.getByteArray(initialByteArraySize)
	result := C.G2Diagnostic_fetchNextEntityBySize_helper(C.uintptr_t(entityListBySizeHandle), (*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	if result != 0 {
		err = g2diagnostic.newError(ctx, 4004, result, time.Since(entryTime))
	}
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(10, string(stringBuffer), err, time.Since(entryTime))
	}
	return string(stringBuffer), err
}

func (g2diagnostic *G2diagnosticImpl) FindEntitiesByFeatureIDs(ctx context.Context, features string) (string, error) {
	//  _DLEXPORT int G2Diagnostic_findEntitiesByFeatureIDs(const char *features, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(11, features)
	}
	entryTime := time.Now()
	var err error = nil
	featuresForC := C.CString(features)
	defer C.free(unsafe.Pointer(featuresForC))
	result := C.G2Diagnostic_findEntitiesByFeatureIDs_helper(featuresForC)
	if result.returnCode != 0 {
		err = g2diagnostic.newError(ctx, 4005, features, result.returnCode, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(12, features, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

// GetAvailableMemory returns the available memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) GetAvailableMemory(ctx context.Context) (int64, error) {
	// _DLEXPORT long long G2Diagnostic_getAvailableMemory();
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(13)
	}
	entryTime := time.Now()
	var err error = nil
	result := int64(C.G2Diagnostic_getAvailableMemory())
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(14, result, err, time.Since(entryTime))
	}
	return result, err
}

func (g2diagnostic *G2diagnosticImpl) GetDataSourceCounts(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getDataSourceCounts(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(15)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Diagnostic_getDataSourceCounts_helper()
	if result.returnCode != 0 {
		err = g2diagnostic.newError(ctx, 4006, result.returnCode, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(16, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

// GetDBInfo returns information about the database connection.
func (g2diagnostic *G2diagnosticImpl) GetDBInfo(ctx context.Context) (string, error) {
	// _DLEXPORT int G2Diagnostic_getDBInfo(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(17)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Diagnostic_getDBInfo_helper()
	if result.returnCode != 0 {
		err = g2diagnostic.newError(ctx, 4007, result.returnCode, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(18, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2diagnostic *G2diagnosticImpl) GetEntityDetails(ctx context.Context, entityID int64, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getEntityDetails(const long long entityID, const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(19, entityID, includeInternalFeatures)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Diagnostic_getEntityDetails_helper(C.longlong(entityID), C.int(includeInternalFeatures))
	if result.returnCode != 0 {
		err = g2diagnostic.newError(ctx, 4008, entityID, includeInternalFeatures, result.returnCode, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(20, entityID, includeInternalFeatures, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2diagnostic *G2diagnosticImpl) GetEntityListBySize(ctx context.Context, entitySize int) (uintptr, error) {
	//  _DLEXPORT int G2Diagnostic_getEntityListBySize(const size_t entitySize, EntityListBySizeHandle* entityListBySizeHandle);
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(21, entitySize)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Diagnostic_getEntityListBySize_helper(C.size_t(entitySize))
	if result.returnCode != 0 {
		err = g2diagnostic.newError(ctx, 4009, entitySize, result.returnCode, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(22, entitySize, (uintptr)(result.response), err, time.Since(entryTime))
	}
	return (uintptr)(result.response), err
}

func (g2diagnostic *G2diagnosticImpl) GetEntityResume(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getEntityResume(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(23, entityID)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Diagnostic_getEntityResume_helper(C.longlong(entityID))
	if result.returnCode != 0 {
		err = g2diagnostic.newError(ctx, 4010, entityID, result.returnCode, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(24, entityID, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2diagnostic *G2diagnosticImpl) GetEntitySizeBreakdown(ctx context.Context, minimumEntitySize int, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getEntitySizeBreakdown(const size_t minimumEntitySize, const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(25, minimumEntitySize, includeInternalFeatures)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Diagnostic_getEntitySizeBreakdown_helper(C.size_t(minimumEntitySize), C.int(includeInternalFeatures))
	if result.returnCode != 0 {
		err = g2diagnostic.newError(ctx, 4011, minimumEntitySize, includeInternalFeatures, result.returnCode, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(26, minimumEntitySize, includeInternalFeatures, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2diagnostic *G2diagnosticImpl) GetFeature(ctx context.Context, libFeatID int64) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getFeature(const long long libFeatID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(27, libFeatID)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Diagnostic_getFeature_helper(C.longlong(libFeatID))
	if result.returnCode != 0 {
		err = g2diagnostic.newError(ctx, 4012, libFeatID, result.returnCode, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(28, libFeatID, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2diagnostic *G2diagnosticImpl) GetGenericFeatures(ctx context.Context, featureType string, maximumEstimatedCount int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getGenericFeatures(const char* featureType, const size_t maximumEstimatedCount, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(29, featureType, maximumEstimatedCount)
	}
	entryTime := time.Now()
	var err error = nil
	featureTypeForC := C.CString(featureType)
	defer C.free(unsafe.Pointer(featureTypeForC))
	result := C.G2Diagnostic_getGenericFeatures_helper(featureTypeForC, C.size_t(maximumEstimatedCount))
	if result.returnCode != 0 {
		err = g2diagnostic.newError(ctx, 4013, featureType, maximumEstimatedCount, result.returnCode, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(30, featureType, maximumEstimatedCount, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

// GetLastException returns the last exception encountered in the Senzing Engine.
func (g2diagnostic *G2diagnosticImpl) GetLastException(ctx context.Context) (string, error) {
	// _DLEXPORT int G2Config_getLastException(char *buffer, const size_t bufSize);
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(31)
	}
	entryTime := time.Now()
	var err error = nil
	stringBuffer := g2diagnostic.getByteArray(initialByteArraySize)
	result := C.G2Diagnostic_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	if result == 0 {
		messageGenerator := g2diagnostic.getMessageGenerator()
		err = messageGenerator.Error(4014, result, time.Since(entryTime))
	}
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(32, string(stringBuffer), err, time.Since(entryTime))
	}
	return string(stringBuffer), err
}

func (g2diagnostic *G2diagnosticImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2Diagnostic_getLastExceptionCode();
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(33)
	}
	entryTime := time.Now()
	var err error = nil
	result := int(C.G2Diagnostic_getLastExceptionCode())
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(34, result, err, time.Since(entryTime))
	}
	return result, err
}

// GetLogicalCores returns the number of logical cores on the host system.
func (g2diagnostic *G2diagnosticImpl) GetLogicalCores(ctx context.Context) (int, error) {
	// _DLEXPORT int G2Diagnostic_getLogicalCores();
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(35)
	}
	entryTime := time.Now()
	var err error = nil
	result := int(C.G2Diagnostic_getLogicalCores())
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(36, result, err, time.Since(entryTime))
	}
	return result, err
}

func (g2diagnostic *G2diagnosticImpl) GetMappingStatistics(ctx context.Context, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getMappingStatistics(const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(37, includeInternalFeatures)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Diagnostic_getMappingStatistics_helper(C.int(includeInternalFeatures))
	if result.returnCode != 0 {
		err = g2diagnostic.newError(ctx, 4015, includeInternalFeatures, result.returnCode, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(38, includeInternalFeatures, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

// GetPhysicalCores returns the number of physical cores on the host system.
func (g2diagnostic *G2diagnosticImpl) GetPhysicalCores(ctx context.Context) (int, error) {
	// _DLEXPORT int G2Diagnostic_getPhysicalCores();
	var result int
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(39)
	}
	entryTime := time.Now()
	var err error = nil
	result = int(C.G2Diagnostic_getPhysicalCores())
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(40, result, err, time.Since(entryTime))
	}
	return result, err
}

func (g2diagnostic *G2diagnosticImpl) GetRelationshipDetails(ctx context.Context, relationshipID int64, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getRelationshipDetails(const long long relationshipID, const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(41, relationshipID, includeInternalFeatures)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Diagnostic_getRelationshipDetails_helper(C.longlong(relationshipID), C.int(includeInternalFeatures))
	if result.returnCode != 0 {
		err = g2diagnostic.newError(ctx, 4016, relationshipID, includeInternalFeatures, result.returnCode, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(42, relationshipID, includeInternalFeatures, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2diagnostic *G2diagnosticImpl) GetResolutionStatistics(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getResolutionStatistics(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(43)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Diagnostic_getResolutionStatistics_helper()
	if result.returnCode != 0 {
		err = g2diagnostic.newError(ctx, 4017, result.returnCode, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(44, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

// GetTotalSystemMemory returns the total memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) GetTotalSystemMemory(ctx context.Context) (int64, error) {
	// _DLEXPORT long long G2Diagnostic_getTotalSystemMemory();
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(57)
	}
	entryTime := time.Now()
	var err error = nil
	result := int64(C.G2Diagnostic_getTotalSystemMemory())
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(46, result, err, time.Since(entryTime))
	}
	return result, err
}

// Init initializes the Senzing G2diagnosis.
func (g2diagnostic *G2diagnosticImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2Diagnostic_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(47, moduleName, iniParams, verboseLogging)
	}
	entryTime := time.Now()
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2Diagnostic_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))
	if result != 0 {
		err = g2diagnostic.newError(ctx, 4018, moduleName, iniParams, verboseLogging, result, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(48, moduleName, iniParams, verboseLogging, err, time.Since(entryTime))
	}
	return err
}

func (g2diagnostic *G2diagnosticImpl) InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error {
	//  _DLEXPORT int G2Diagnostic_initWithConfigID(const char *moduleName, const char *iniParams, const long long initConfigID, const int verboseLogging);
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(49, moduleName, iniParams, initConfigID, verboseLogging)
	}
	entryTime := time.Now()
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2Diagnostic_initWithConfigID(moduleNameForC, iniParamsForC, C.longlong(initConfigID), C.int(verboseLogging))
	if result != 0 {
		err = g2diagnostic.newError(ctx, 4019, moduleName, iniParams, initConfigID, verboseLogging, result, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(50, moduleName, iniParams, initConfigID, verboseLogging, err, time.Since(entryTime))
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
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(51, initConfigID)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2Diagnostic_reinit(C.longlong(initConfigID))
	if result != 0 {
		err = g2diagnostic.newError(ctx, 4020, initConfigID, result, time.Since(entryTime))
	}
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(52, initConfigID, err, time.Since(entryTime))
	}
	return err
}

func (g2diagnostic *G2diagnosticImpl) SetLogLevel(ctx context.Context, logLevel logger.Level) error {
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(53, logLevel)
	}
	entryTime := time.Now()
	var err error = nil
	g2diagnostic.getLogger().SetLogLevel(messagelogger.Level(logLevel))
	g2diagnostic.isTrace = g2diagnostic.getLogger().GetLogLevel() == messagelogger.LevelTrace
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(54, logLevel, err, time.Since(entryTime))
	}
	return err
}
