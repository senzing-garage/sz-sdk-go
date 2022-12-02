/*
The G2diagnosticImpl implementation is a wrapper over the Senzing libg2diagnostic library.
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
	"github.com/senzing/go-logging/messagelogger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// G2diagnosticImpl is the default implementation of the G2diagnostic interface.
type G2diagnosticImpl struct {
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
func (g2diagnostic *G2diagnosticImpl) getByteArrayC(size int) *C.char {
	bytes := C.malloc(C.size_t(size))
	return (*C.char)(bytes)
}

// Make a byte array.
func (g2diagnostic *G2diagnosticImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

// Create a new error.
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
	errorMessage, err := g2diagnostic.getLogger().Message(errorNumber, newDetails...)
	if err != nil {
		errorMessage = err.Error()
	}

	return errors.New(errorMessage)
}

// Get the Logger singleton.
func (g2diagnostic *G2diagnosticImpl) getLogger() messagelogger.MessageLoggerInterface {
	if g2diagnostic.logger == nil {
		g2diagnostic.logger, _ = messagelogger.NewSenzingApiLogger(ProductId, IdMessages, IdStatuses, messagelogger.LevelInfo)
	}
	return g2diagnostic.logger
}

// Trace method entry.
func (g2diagnostic *G2diagnosticImpl) traceEntry(errorNumber int, details ...interface{}) {
	g2diagnostic.getLogger().Log(errorNumber, details...)
}

// Trace method exit.
func (g2diagnostic *G2diagnosticImpl) traceExit(errorNumber int, details ...interface{}) {
	g2diagnostic.getLogger().Log(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The CheckDBPerf method performs inserts to determine rate of insertion.

Input
  - ctx: A context to control lifecycle.
  - secondsToRun: Duration of the test in seconds.

Output

  - A string containing a JSON document.
    Example: `{"numRecordsInserted":0,"insertTime":0}`
*/
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

/*
The ClearLastException method erases the last exception message held by the Senzing G2Config object.

Input
  - ctx: A context to control lifecycle.
*/
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

/*
The CloseEntityListBySize method closes the list created by GetEntityListBySize().
It is part of the GetEntityListBySize(), FetchNextEntityBySize(), CloseEntityListBySize()
lifecycle of a list of sized entities.

Input
  - ctx: A context to control lifecycle.
  - entityListBySizeHandle: A handle created by GetEntityListBySize().
*/
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

/*
The Destroy method will destroy and perform cleanup for the Senzing G2Diagnostic object.
It should be called after all other calls are complete.

Input
  - ctx: A context to control lifecycle.
*/
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

/*
The FetchNextEntityBySize method gets the next section of the list created by GetEntityListBySize().
It is part of the GetEntityListBySize(), FetchNextEntityBySize(), CloseEntityListBySize()
lifecycle of a list of sized entities.

Input
  - ctx: A context to control lifecycle.
  - entityListBySizeHandle: A handle created by GetEntityListBySize().

Output
  - A string containing a JSON document.
    Example: `[{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","ENT_SRC_KEY":"2E83FCA24AF2996D77569554CC4FD7775F861A8F","ENT_SRC_DESC":"SEAMAN","RECORD_ID":"9001","JSON_DATA":"{\"SOCIAL_HANDLE\":\"flavorh\",\"DATE_OF_BIRTH\":\"4/8/1983\",\"ADDR_STATE\":\"LA\",\"ADDR_POSTAL_CODE\":\"71232\",\"SSN_NUMBER\":\"053-39-3251\",\"GENDER\":\"F\",\"srccode\":\"MDMPER\",\"CC_ACCOUNT_NUMBER\":\"5534202208773608\",\"ADDR_CITY\":\"Delhi\",\"DRIVERS_LICENSE_STATE\":\"DE\",\"PHONE_NUMBER\":\"225-671-0796\",\"NAME_LAST\":\"SEAMAN\",\"entityid\":\"284430058\",\"ADDR_LINE1\":\"772 Armstrong RD\",\"DATA_SOURCE\":\"TEST_G2DIAGNOSTIC\",\"ENTITY_TYPE\":\"GENERIC\",\"DSRC_ACTION\":\"A\",\"RECORD_ID\":\"9001\"}","OBS_ENT_ID":1,"ER_ID":0},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","ENT_SRC_KEY":"2E83FCA24AF2996D77569554CC4FD7775F861A8F","ENT_SRC_DESC":"SEAMAN","RECORD_ID":"9002","JSON_DATA":"{\"SOCIAL_HANDLE\":\"flavorh\",\"DATE_OF_BIRTH\":\"4/8/1983\",\"ADDR_STATE\":\"LA\",\"ADDR_POSTAL_CODE\":\"71232\",\"SSN_NUMBER\":\"053-39-3251\",\"GENDER\":\"F\",\"srccode\":\"MDMPER\",\"CC_ACCOUNT_NUMBER\":\"5534202208773608\",\"ADDR_CITY\":\"Delhi\",\"DRIVERS_LICENSE_STATE\":\"DE\",\"PHONE_NUMBER\":\"225-671-0796\",\"NAME_LAST\":\"SEAMAN\",\"entityid\":\"284430058\",\"ADDR_LINE1\":\"772 Armstrong RD\",\"DATA_SOURCE\":\"TEST_G2DIAGNOSTIC\",\"ENTITY_TYPE\":\"GENERIC\",\"DSRC_ACTION\":\"A\",\"RECORD_ID\":\"9002\"}","OBS_ENT_ID":1,"ER_ID":0}]`
*/
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

/*
The FindEntitiesByFeatureIDs method finds entities having any of the lib feat id specified in the "features" JSON document.
The "features" also contains an entity id.
This entity is ignored in the returned values.

Input
  - ctx: A context to control lifecycle.
  - features: A JSON document having the format: `{"ENTITY_ID":<entity id>,"LIB_FEAT_IDS":[<id1>,<id2>,...<idn>]}` where ENTITY_ID specifies the entity to ignore in the returns and <id#> are the lib feat ids used to query for entities.

Output
  - A string containing a JSON document.
    Example: `[{"LIB_FEAT_ID":4,"USAGE_TYPE":"","RES_ENT_ID":2}]`
*/
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

/*
The GetAvailableMemory method returns the available memory, in bytes, on the host system.

Input
  - ctx: A context to control lifecycle.

Output
  - Number of bytes of available memory.
*/
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

/*
The GetDataSourceCounts method returns information about data sources.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document enumerating data sources.
    Example: `[{"DSRC_ID":1001,"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_ID":3,"ETYPE_CODE":"GENERIC","OBS_ENT_COUNT":2,"DSRC_RECORD_COUNT":3}]`
*/
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

/*
The GetDBInfo method returns information about the database connection.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document enumerating data sources.
    Example: `{"Hybrid Mode":false,"Database Details":[{"Name":"0.0.0.0","Type":"postgresql"}]}`
*/
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

/*
The GetEntityDetails method returns information about the database connection.

Input
  - ctx: A context to control lifecycle.
  - entityID: Unique number identifying an entity.
  - includeInternalFeatures: FIXME:

Output
  - A JSON document enumerating FIXME:.
    Example: `[{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"NAME","USAGE_TYPE":"","FEAT_DESC":"SEAMAN"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"DOB","USAGE_TYPE":"","FEAT_DESC":"4/8/1983"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"GENDER","USAGE_TYPE":"","FEAT_DESC":"F"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"ADDRESS","USAGE_TYPE":"","FEAT_DESC":"772 Armstrong RD Delhi LA 71232"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"PHONE","USAGE_TYPE":"","FEAT_DESC":"225-671-0796"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"SSN","USAGE_TYPE":"","FEAT_DESC":"053-39-3251"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"LOGIN_ID","USAGE_TYPE":"","FEAT_DESC":"flavorh"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"ACCT_NUM","USAGE_TYPE":"CC","FEAT_DESC":"5534202208773608"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN|SSN=3251"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN|DOB=80804"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN|PHONE.PHONE_LAST_5=10796"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN|DOB.MMDD_HASH=0804"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN|POST=71232"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN|ADDRESS.CITY_STD=TL"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN|DOB.MMYY_HASH=0483"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"ADDR_KEY","USAGE_TYPE":"","FEAT_DESC":"772|ARMSTRNK||TL"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"ADDR_KEY","USAGE_TYPE":"","FEAT_DESC":"772|ARMSTRNK||71232"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"ID_KEY","USAGE_TYPE":"","FEAT_DESC":"ACCT_NUM=5534202208773608"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"ID_KEY","USAGE_TYPE":"","FEAT_DESC":"SSN=053-39-3251"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"PHONE_KEY","USAGE_TYPE":"","FEAT_DESC":"2256710796"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"SEARCH_KEY","USAGE_TYPE":"","FEAT_DESC":"LOGIN_ID:FLAVORH|"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"SEARCH_KEY","USAGE_TYPE":"","FEAT_DESC":"SSN:3251|80804|"}]`
*/
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

/*
The GetEntityListBySize method gets the next section of the list created by GetEntityListBySize().
It is part of the GetEntityListBySize(), FetchNextEntityBySize(), CloseEntityListBySize()
lifecycle of a list of sized entities.

Input
  - ctx: A context to control lifecycle.
  - entitySize: FIXME:

Output
  - A handle to an entity list to be used with FetchNextEntityBySize() and CloseEntityListBySize().
*/
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

/*
The GetEntityResume method FIXME:

Input
  - ctx: A context to control lifecycle.
  - entityID: Unique number identifying an entity.

Output
  - A string containing a JSON document.
    Example: `[{"RES_ENT_ID":1,"REL_ENT_ID":0,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","ENT_SRC_DESC":"SEAMAN","JSON_DATA":"{\"SOCIAL_HANDLE\":\"flavorh\",\"DATE_OF_BIRTH\":\"4/8/1983\",\"ADDR_STATE\":\"LA\",\"ADDR_POSTAL_CODE\":\"71232\",\"SSN_NUMBER\":\"053-39-3251\",\"GENDER\":\"F\",\"srccode\":\"MDMPER\",\"CC_ACCOUNT_NUMBER\":\"5534202208773608\",\"ADDR_CITY\":\"Delhi\",\"DRIVERS_LICENSE_STATE\":\"DE\",\"PHONE_NUMBER\":\"225-671-0796\",\"NAME_LAST\":\"SEAMAN\",\"entityid\":\"284430058\",\"ADDR_LINE1\":\"772 Armstrong RD\",\"DATA_SOURCE\":\"TEST_G2DIAGNOSTIC\",\"ENTITY_TYPE\":\"GENERIC\",\"DSRC_ACTION\":\"A\",\"RECORD_ID\":\"9001\"}"},{"RES_ENT_ID":1,"REL_ENT_ID":0,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9002","ENT_SRC_DESC":"SEAMAN","JSON_DATA":"{\"SOCIAL_HANDLE\":\"flavorh\",\"DATE_OF_BIRTH\":\"4/8/1983\",\"ADDR_STATE\":\"LA\",\"ADDR_POSTAL_CODE\":\"71232\",\"SSN_NUMBER\":\"053-39-3251\",\"GENDER\":\"F\",\"srccode\":\"MDMPER\",\"CC_ACCOUNT_NUMBER\":\"5534202208773608\",\"ADDR_CITY\":\"Delhi\",\"DRIVERS_LICENSE_STATE\":\"DE\",\"PHONE_NUMBER\":\"225-671-0796\",\"NAME_LAST\":\"SEAMAN\",\"entityid\":\"284430058\",\"ADDR_LINE1\":\"772 Armstrong RD\",\"DATA_SOURCE\":\"TEST_G2DIAGNOSTIC\",\"ENTITY_TYPE\":\"GENERIC\",\"DSRC_ACTION\":\"A\",\"RECORD_ID\":\"9002\"}"},{"RES_ENT_ID":1,"REL_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9003","ENT_SRC_DESC":"Smith","JSON_DATA":"{\"ADDR_STATE\":\"LA\",\"ADDR_POSTAL_CODE\":\"71232\",\"GENDER\":\"M\",\"srccode\":\"MDMPER\",\"ADDR_CITY\":\"Delhi\",\"PHONE_NUMBER\":\"225-671-0796\",\"NAME_LAST\":\"Smith\",\"entityid\":\"284430058\",\"ADDR_LINE1\":\"772 Armstrong RD\",\"DATA_SOURCE\":\"TEST_G2DIAGNOSTIC\",\"ENTITY_TYPE\":\"GENERIC\",\"DSRC_ACTION\":\"A\",\"RECORD_ID\":\"9003\"}"}]`
*/
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

/*
The GetEntitySizeBreakdown method FIXME:

Input
  - ctx: A context to control lifecycle.
  - minimumEntitySize: FIXME:
  - includeInternalFeatures: FIXME:

Output
  - A string containing a JSON document.
    Example: `[{"ENTITY_SIZE": 1,"ENTITY_COUNT": 2,"NAME": 1.00,"DOB": 0.50,"GENDER": 1.00,"ADDRESS": 1.00,"PHONE": 1.00,"SSN": 0.50,"LOGIN_ID": 0.50,"ACCT_NUM": 0.50,"NAME_KEY": 6.00,"ADDR_KEY": 2.00,"ID_KEY": 1.00,"PHONE_KEY": 1.00,"SEARCH_KEY": 1.00,"MIN_RES_ENT_ID": 1,"MAX_RES_ENT_ID": 2}]`
*/
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

/*
The GetFeature method retrieves a stored feature.

Input
  - ctx: A context to control lifecycle.
  - libFeatID: The identifier of the feature requested in the search.

Output
  - A string containing a JSON document.
    Example: `{"LIB_FEAT_ID":1,"FTYPE_CODE":"NAME","ELEMENTS":[{"FELEM_CODE":"TOKENIZED_NM","FELEM_VALUE":"SEAMAN"},{"FELEM_CODE":"CATEGORY","FELEM_VALUE":"PERSON"},{"FELEM_CODE":"CULTURE","FELEM_VALUE":"ANGLO"},{"FELEM_CODE":"SUR_NAME","FELEM_VALUE":"SEAMAN"},{"FELEM_CODE":"FULL_NAME","FELEM_VALUE":"SEAMAN"}]}`
*/
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

/*
The GetGenericFeatures method retrieves a stored feature.

Input
  - ctx: A context to control lifecycle.
  - featureType: FIXME:
  - maximumEstimatedCount: FIXME:

Output
  - A string containing a JSON document.
    Example: `{FIXME:}`
*/
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

/*
The GetLastException method retrieves the last exception thrown in Senzing's G2Diagnostic.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the error received from Senzing's G2Config.
*/
func (g2diagnostic *G2diagnosticImpl) GetLastException(ctx context.Context) (string, error) {
	// _DLEXPORT int G2Config_getLastException(char *buffer, const size_t bufSize);
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(31)
	}
	entryTime := time.Now()
	var err error = nil
	stringBuffer := g2diagnostic.getByteArray(initialByteArraySize)
	C.G2Diagnostic_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	// if result == 0 { // "result" is length of exception message.
	// 	err = g2diagnostic.getLogger().Error(4014, result, time.Since(entryTime))
	// }
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(32, string(stringBuffer), err, time.Since(entryTime))
	}
	return string(stringBuffer), err
}

/*
The GetLastExceptionCode method retrieves the code of the last exception thrown in Senzing's G2Diagnostic.

Input:
  - ctx: A context to control lifecycle.

Output:
  - An int containing the error received from Senzing's G2Config.
*/
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

/*
The GetLogicalCores method returns the number of logical cores on the host system.

Input
  - ctx: A context to control lifecycle.

Output
  - Number of logical cores.
*/
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

/*
The GetMappingStatistics method FIXME:

Input
  - ctx: A context to control lifecycle.
  - includeInternalFeatures: FIXME:

Output
  - A string containing a JSON document.
    Example: `[{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"NAME","USAGE_TYPE":"","REC_COUNT":2,"REC_PCT":1.0,"UNIQ_COUNT":2,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"SEAMAN","MAX_FEAT_DESC":"Smith"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"DOB","USAGE_TYPE":"","REC_COUNT":1,"REC_PCT":0.5,"UNIQ_COUNT":1,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"4/8/1983","MAX_FEAT_DESC":"4/8/1983"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"GENDER","USAGE_TYPE":"","REC_COUNT":2,"REC_PCT":1.0,"UNIQ_COUNT":2,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"F","MAX_FEAT_DESC":"M"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"ADDRESS","USAGE_TYPE":"","REC_COUNT":2,"REC_PCT":1.0,"UNIQ_COUNT":1,"UNIQ_PCT":0.5,"MIN_FEAT_DESC":"772 Armstrong RD Delhi LA 71232","MAX_FEAT_DESC":"772 Armstrong RD Delhi LA 71232"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"PHONE","USAGE_TYPE":"","REC_COUNT":2,"REC_PCT":1.0,"UNIQ_COUNT":1,"UNIQ_PCT":0.5,"MIN_FEAT_DESC":"225-671-0796","MAX_FEAT_DESC":"225-671-0796"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"SSN","USAGE_TYPE":"","REC_COUNT":1,"REC_PCT":0.5,"UNIQ_COUNT":1,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"053-39-3251","MAX_FEAT_DESC":"053-39-3251"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"LOGIN_ID","USAGE_TYPE":"","REC_COUNT":1,"REC_PCT":0.5,"UNIQ_COUNT":1,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"flavorh","MAX_FEAT_DESC":"flavorh"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"ACCT_NUM","USAGE_TYPE":"CC","REC_COUNT":1,"REC_PCT":0.5,"UNIQ_COUNT":1,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"5534202208773608","MAX_FEAT_DESC":"5534202208773608"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","REC_COUNT":12,"REC_PCT":6.0,"UNIQ_COUNT":12,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"SM0","MAX_FEAT_DESC":"SMN|SSN=3251"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"Yes","FTYPE_CODE":"ADDR_KEY","USAGE_TYPE":"","REC_COUNT":4,"REC_PCT":2.0,"UNIQ_COUNT":2,"UNIQ_PCT":0.5,"MIN_FEAT_DESC":"772|ARMSTRNK||71232","MAX_FEAT_DESC":"772|ARMSTRNK||TL"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"Yes","FTYPE_CODE":"ID_KEY","USAGE_TYPE":"","REC_COUNT":2,"REC_PCT":1.0,"UNIQ_COUNT":2,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"ACCT_NUM=5534202208773608","MAX_FEAT_DESC":"SSN=053-39-3251"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"Yes","FTYPE_CODE":"PHONE_KEY","USAGE_TYPE":"","REC_COUNT":2,"REC_PCT":1.0,"UNIQ_COUNT":1,"UNIQ_PCT":0.5,"MIN_FEAT_DESC":"2256710796","MAX_FEAT_DESC":"2256710796"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"Yes","FTYPE_CODE":"SEARCH_KEY","USAGE_TYPE":"","REC_COUNT":2,"REC_PCT":1.0,"UNIQ_COUNT":2,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"LOGIN_ID:FLAVORH|","MAX_FEAT_DESC":"SSN:3251|80804|"}]`
*/
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

/*
The GetPhysicalCores method returns the number of physical cores on the host system.

Input
  - ctx: A context to control lifecycle.

Output
  - Number of physical cores.
*/
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

/*
The GetRelationshipDetails method FIXME:

Input
  - ctx: A context to control lifecycle.
  - relationshipID: FIXME:
  - includeInternalFeatures: FIXME:

Output
  - A string containing a JSON document.
    Example: `[{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME","FEAT_DESC":"SEAMAN"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"DOB","FEAT_DESC":"4/8/1983"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"GENDER","FEAT_DESC":"F"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ADDRESS","FEAT_DESC":"772 Armstrong RD Delhi LA 71232"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"PHONE","FEAT_DESC":"225-671-0796"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"SSN","FEAT_DESC":"053-39-3251"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"LOGIN_ID","FEAT_DESC":"flavorh"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ACCT_NUM","FEAT_DESC":"5534202208773608"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN|SSN=3251"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN|DOB=80804"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN|PHONE.PHONE_LAST_5=10796"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN|DOB.MMDD_HASH=0804"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN|POST=71232"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN|ADDRESS.CITY_STD=TL"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN|DOB.MMYY_HASH=0483"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ADDR_KEY","FEAT_DESC":"772|ARMSTRNK||TL"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ADDR_KEY","FEAT_DESC":"772|ARMSTRNK||71232"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ID_KEY","FEAT_DESC":"ACCT_NUM=5534202208773608"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ID_KEY","FEAT_DESC":"SSN=053-39-3251"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"PHONE_KEY","FEAT_DESC":"2256710796"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"SEARCH_KEY","FEAT_DESC":"LOGIN_ID:FLAVORH|"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"SEARCH_KEY","FEAT_DESC":"SSN:3251|80804|"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME","FEAT_DESC":"Smith"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"GENDER","FEAT_DESC":"M"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ADDRESS","FEAT_DESC":"772 Armstrong RD Delhi LA 71232"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"PHONE","FEAT_DESC":"225-671-0796"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SM0|POST=71232"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SM0|ADDRESS.CITY_STD=TL"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SM0|PHONE.PHONE_LAST_5=10796"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SM0"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ADDR_KEY","FEAT_DESC":"772|ARMSTRNK||TL"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ADDR_KEY","FEAT_DESC":"772|ARMSTRNK||71232"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"PHONE_KEY","FEAT_DESC":"2256710796"}]`
*/
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

/*
The GetResolutionStatistics method FIXME:

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing a JSON document.
    Example: `[{"MATCH_LEVEL":3,"MATCH_KEY":"+ADDRESS+PHONE-GENDER","RAW_MATCH_KEYS":[{"MATCH_KEY":"+ADDRESS+PHONE-GENDER"}],"ERRULE_ID":200,"ERRULE_CODE":"MFF","IS_AMBIGUOUS":"No","RECORD_COUNT":1,"MIN_RES_ENT_ID":1,"MAX_RES_ENT_ID":2,"MIN_RES_REL_ID":1,"MAX_RES_REL_ID":1}]`
*/
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

/*
The GetTotalSystemMemory method returns the total memory, in bytes, on the host system.

Input
  - ctx: A context to control lifecycle.

Output
  - Number of bytes of memory.
*/
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

/*
The Init method initializes the Senzing G2Diagnosis object.
It must be called prior to any other calls.

Input
  - ctx: A context to control lifecycle.
  - moduleName: A name for the auditing node, to help identify it within system logs.
  - iniParams: A JSON string containing configuration paramters.
  - verboseLogging: A flag to enable deeper logging of the G2 processing. 0 for no Senzing logging; 1 for logging.
*/
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

/*
The InitWithConfigID method initializes the Senzing G2Diagnosis object with a non-default configuration ID.
It must be called prior to any other calls.

Input
  - ctx: A context to control lifecycle.
  - moduleName: A name for the auditing node, to help identify it within system logs.
  - iniParams: A JSON string containing configuration paramters.
  - initConfigID: The configuration ID used for the initialization.
  - verboseLogging: A flag to enable deeper logging of the G2 processing. 0 for no Senzing logging; 1 for logging.
*/
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

/*
The Reinit method re-initializes the Senzing G2Diagnosis object.

Input
  - ctx: A context to control lifecycle.
  - initConfigID: The configuration ID used for the initialization.
*/
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

/*
The SetLogLevel method sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevel: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
func (g2diagnostic *G2diagnosticImpl) SetLogLevel(ctx context.Context, logLevel logger.Level) error {
	if g2diagnostic.isTrace {
		g2diagnostic.traceEntry(53, logLevel)
	}
	entryTime := time.Now()
	var err error = nil
	g2diagnostic.getLogger().SetLogLevel(messagelogger.Level(logLevel))
	g2diagnostic.isTrace = (g2diagnostic.getLogger().GetLogLevel() == messagelogger.LevelTrace)
	if g2diagnostic.isTrace {
		defer g2diagnostic.traceExit(54, logLevel, err, time.Since(entryTime))
	}
	return err
}
