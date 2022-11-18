/*
The G2engineImpl implementation...
*/
package g2engine

/*
#include "g2engine.h"
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

type G2engineImpl struct {
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
func (g2engine *G2engineImpl) getByteArrayC(size int) *C.char {
	bytes := C.malloc(C.size_t(size))
	return (*C.char)(bytes)
}

func (g2engine *G2engineImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

func (g2engine *G2engineImpl) newError(ctx context.Context, errorNumber int, details ...interface{}) error {
	lastException, err := g2engine.GetLastException(ctx)
	defer g2engine.ClearLastException(ctx)
	message := lastException
	if err != nil {
		message = err.Error()
	}

	var newDetails []interface{}
	newDetails = append(newDetails, details...)
	newDetails = append(newDetails, errors.New(message))
	errorMessage, err := g2engine.getLogger().Message(errorNumber, newDetails...)
	if err != nil {
		errorMessage = err.Error()
	}

	return errors.New(errorMessage)
}

func (g2engine *G2engineImpl) getLogger() messagelogger.MessageLoggerInterface {
	if g2engine.logger == nil {
		g2engine.logger, _ = messagelogger.NewSenzingApiLogger(ProductId, IdMessages, IdStatuses, messagelogger.LevelInfo)
	}
	return g2engine.logger
}

func (g2engine *G2engineImpl) traceEntry(errorNumber int, details ...interface{}) {
	g2engine.getLogger().Log(errorNumber, details...)
}

func (g2engine *G2engineImpl) traceExit(errorNumber int, details ...interface{}) {
	g2engine.getLogger().Log(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func (g2engine *G2engineImpl) AddRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error {
	//  _DLEXPORT int G2_addRecord(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID);
	if g2engine.isTrace {
		g2engine.traceEntry(1, dataSourceCode, recordID, jsonData, loadID)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))
	result := C.G2_addRecord(dataSourceCodeForC, recordIDForC, jsonDataForC, loadIDForC)
	if result != 0 {
		err = g2engine.newError(ctx, 4001, dataSourceCode, recordID, jsonData, loadID, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(2, dataSourceCode, recordID, jsonData, loadID, err, time.Since(entryTime))
	}
	return err
}

func (g2engine *G2engineImpl) AddRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_addRecordWithInfo(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(3, dataSourceCode, recordID, jsonData, loadID, flags)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))
	result := C.G2_addRecordWithInfo_helper(dataSourceCodeForC, recordIDForC, jsonDataForC, loadIDForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4002, dataSourceCode, recordID, jsonData, loadID, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4, dataSourceCode, recordID, jsonData, loadID, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) AddRecordWithInfoWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string, flags int64) (string, string, error) {
	//  _DLEXPORT int G2_addRecordWithInfoWithReturnedRecordID(const char* dataSourceCode, const char* jsonData, const char *loadID, const long long flags, char *recordIDBuf, const size_t recordIDBufSize, char **responseBuf, size_t *responseBufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(5, dataSourceCode, jsonData, loadID, flags)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))
	result := C.G2_addRecordWithInfoWithReturnedRecordID_helper(dataSourceCodeForC, jsonDataForC, loadIDForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4003, dataSourceCode, jsonData, loadID, flags, result.returnCode, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(6, dataSourceCode, jsonData, loadID, flags, C.GoString(result.withInfo), C.GoString(result.recordID), err, time.Since(entryTime))
	}
	return C.GoString(result.withInfo), C.GoString(result.recordID), err
}

func (g2engine *G2engineImpl) AddRecordWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string) (string, error) {
	//  _DLEXPORT int G2_addRecordWithReturnedRecordID(const char* dataSourceCode, const char* jsonData, const char *loadID, char *recordIDBuf, const size_t bufSize);
	if g2engine.isTrace {
		g2engine.traceEntry(7, dataSourceCode, jsonData, loadID)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))
	stringBuffer := g2engine.getByteArray(250)
	result := C.G2_addRecordWithReturnedRecordID(dataSourceCodeForC, jsonDataForC, loadIDForC, (*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	if result != 0 {
		err = g2engine.newError(ctx, 4004, dataSourceCode, jsonData, loadID, result, time.Since(entryTime))
	}
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if g2engine.isTrace {
		defer g2engine.traceExit(8, dataSourceCode, jsonData, loadID, string(stringBuffer), err, time.Since(entryTime))
	}
	return string(stringBuffer), err
}

func (g2engine *G2engineImpl) CheckRecord(ctx context.Context, record string, recordQueryList string) (string, error) {
	//  _DLEXPORT int G2_checkRecord(const char *record, const char* recordQueryList, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2engine.isTrace {
		g2engine.traceEntry(9, record, recordQueryList)
	}
	entryTime := time.Now()
	var err error = nil
	recordForC := C.CString(record)
	defer C.free(unsafe.Pointer(recordForC))
	recordQueryListForC := C.CString(recordQueryList)
	defer C.free(unsafe.Pointer(recordQueryListForC))
	result := C.G2_checkRecord_helper(recordForC, recordQueryListForC)
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4005, record, recordQueryList, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(10, record, recordQueryList, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2engine *G2engineImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2_clearLastException();
	if g2engine.isTrace {
		g2engine.traceEntry(11)
	}
	entryTime := time.Now()
	var err error = nil
	C.G2_clearLastException()
	if g2engine.isTrace {
		defer g2engine.traceExit(12, err, time.Since(entryTime))
	}
	return err
}

func (g2engine *G2engineImpl) CloseExport(ctx context.Context, responseHandle uintptr) error {
	//  _DLEXPORT int G2_closeExport(ExportHandle responseHandle);
	if g2engine.isTrace {
		g2engine.traceEntry(13, responseHandle)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_closeExport_helper(C.uintptr_t(responseHandle))
	if result != 0 {
		err = g2engine.newError(ctx, 4006, responseHandle, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(14, responseHandle, err, time.Since(entryTime))
	}
	return err
}

func (g2engine *G2engineImpl) CountRedoRecords(ctx context.Context) (int64, error) {
	//  _DLEXPORT long long G2_countRedoRecords();
	if g2engine.isTrace {
		g2engine.traceEntry(15)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_countRedoRecords()
	if g2engine.isTrace {
		defer g2engine.traceExit(16, int64(result), err, time.Since(entryTime))
	}
	return int64(result), err
}

func (g2engine *G2engineImpl) DeleteRecord(ctx context.Context, dataSourceCode string, recordID string, loadID string) error {
	//  _DLEXPORT int G2_deleteRecord(const char* dataSourceCode, const char* recordID, const char* loadID);
	if g2engine.isTrace {
		g2engine.traceEntry(17, dataSourceCode, recordID, loadID)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))
	result := C.G2_deleteRecord(dataSourceCodeForC, recordIDForC, loadIDForC)
	if result != 0 {
		err = g2engine.newError(ctx, 4007, dataSourceCode, recordID, loadID, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(18, dataSourceCode, recordID, loadID, err, time.Since(entryTime))
	}
	return err
}

func (g2engine *G2engineImpl) DeleteRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, loadID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_deleteRecordWithInfo(const char* dataSourceCode, const char* recordID, const char* loadID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(19, dataSourceCode, recordID, loadID, flags)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))
	result := C.G2_deleteRecordWithInfo_helper(dataSourceCodeForC, recordIDForC, loadIDForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4008, dataSourceCode, recordID, loadID, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(20, dataSourceCode, recordID, loadID, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) Destroy(ctx context.Context) error {
	//  _DLEXPORT int G2_destroy();
	if g2engine.isTrace {
		g2engine.traceEntry(21)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_destroy()
	if result != 0 {
		err = g2engine.newError(ctx, 4009, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(22, err, time.Since(entryTime))
	}
	return err
}

func (g2engine *G2engineImpl) ExportConfigAndConfigID(ctx context.Context) (string, int64, error) {
	//  _DLEXPORT int G2_exportConfigAndConfigID(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize), long long* configID );
	if g2engine.isTrace {
		g2engine.traceEntry(23)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_exportConfigAndConfigID_helper()
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4010, result.returnCode, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(24, C.GoString(result.config), int64(C.longlong(result.configID)), err, time.Since(entryTime))
	}
	return C.GoString(result.config), int64(C.longlong(result.configID)), err
}

func (g2engine *G2engineImpl) ExportConfig(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_exportConfig(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2engine.isTrace {
		g2engine.traceEntry(25)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_exportConfig_helper()
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4011, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(26, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) ExportCSVEntityReport(ctx context.Context, csvColumnList string, flags int64) (uintptr, error) {
	//  _DLEXPORT int G2_exportCSVEntityReport(const char* csvColumnList, const long long flags, ExportHandle* responseHandle);
	if g2engine.isTrace {
		g2engine.traceEntry(27, csvColumnList, flags)
	}
	entryTime := time.Now()
	var err error = nil
	csvColumnListForC := C.CString(csvColumnList)
	defer C.free(unsafe.Pointer(csvColumnListForC))
	result := C.G2_exportCSVEntityReport_helper(csvColumnListForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4012, csvColumnList, flags, result.returnCode, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(28, csvColumnList, flags, (uintptr)(result.exportHandle), err, time.Since(entryTime))
	}
	return (uintptr)(result.exportHandle), err
}

func (g2engine *G2engineImpl) ExportJSONEntityReport(ctx context.Context, flags int64) (uintptr, error) {
	//  _DLEXPORT int G2_exportJSONEntityReport(const long long flags, ExportHandle* responseHandle);
	if g2engine.isTrace {
		g2engine.traceEntry(29, flags)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_exportJSONEntityReport_helper(C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4013, flags, result.returnCode, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(30, flags, (uintptr)(result.exportHandle), err, time.Since(entryTime))
	}
	return (uintptr)(result.exportHandle), err
}

func (g2engine *G2engineImpl) FetchNext(ctx context.Context, responseHandle uintptr) (string, error) {
	//  _DLEXPORT int G2_fetchNext(ExportHandle responseHandle, char *responseBuf, const size_t bufSize);
	if g2engine.isTrace {
		g2engine.traceEntry(31, responseHandle)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_fetchNext_helper(C.uintptr_t(responseHandle))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4014, responseHandle, result.returnCode, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(32, responseHandle, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindInterestingEntitiesByEntityID(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_findInterestingEntitiesByEntityID(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(33, entityID, flags)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_findInterestingEntitiesByEntityID_helper(C.longlong(entityID), C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4015, entityID, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(34, entityID, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindInterestingEntitiesByRecordID(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findInterestingEntitiesByRecordID(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(35, dataSourceCode, recordID, flags)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	result := C.G2_findInterestingEntitiesByRecordID_helper(dataSourceCodeForC, recordIDForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4016, dataSourceCode, recordID, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(36, dataSourceCode, recordID, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindNetworkByEntityID(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error) {
	//  _DLEXPORT int G2_findNetworkByEntityID(const char* entityList, const int maxDegree, const int buildOutDegree, const int maxEntities, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(37, entityList, maxDegree, buildOutDegree, maxDegree)
	}
	entryTime := time.Now()
	var err error = nil
	entityListForC := C.CString(entityList)
	defer C.free(unsafe.Pointer(entityListForC))
	result := C.G2_findNetworkByEntityID_helper(entityListForC, C.int(maxDegree), C.int(buildOutDegree), C.int(maxEntities))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4017, entityList, maxDegree, buildOutDegree, maxEntities, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(38, entityList, maxDegree, buildOutDegree, maxDegree, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindNetworkByEntityID_V2(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findNetworkByEntityID_V2(const char* entityList, const int maxDegree, const int buildOutDegree, const int maxEntities, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(39, entityList, maxDegree, buildOutDegree, maxDegree, flags)
	}
	entryTime := time.Now()
	var err error = nil
	entityListForC := C.CString(entityList)
	defer C.free(unsafe.Pointer(entityListForC))
	result := C.G2_findNetworkByEntityID_V2_helper(entityListForC, C.int(maxDegree), C.int(buildOutDegree), C.int(maxEntities), C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4018, entityList, maxDegree, buildOutDegree, maxEntities, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(40, entityList, maxDegree, buildOutDegree, maxDegree, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindNetworkByRecordID(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error) {
	//  _DLEXPORT int G2_findNetworkByRecordID(const char* recordList, const int maxDegree, const int buildOutDegree, const int maxEntities, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(41, recordList, maxDegree, buildOutDegree, maxDegree)
	}
	entryTime := time.Now()
	var err error = nil
	recordListForC := C.CString(recordList)
	defer C.free(unsafe.Pointer(recordListForC))
	result := C.G2_findNetworkByRecordID_helper(recordListForC, C.int(maxDegree), C.int(buildOutDegree), C.int(maxEntities))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4019, recordList, maxDegree, buildOutDegree, maxEntities, recordList, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(42, recordList, maxDegree, buildOutDegree, maxDegree, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindNetworkByRecordID_V2(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findNetworkByRecordID_V2(const char* recordList, const int maxDegree, const int buildOutDegree, const int maxEntities, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(43, recordList, maxDegree, buildOutDegree, maxDegree, flags)
	}
	entryTime := time.Now()
	var err error = nil
	recordListForC := C.CString(recordList)
	defer C.free(unsafe.Pointer(recordListForC))
	result := C.G2_findNetworkByRecordID_V2_helper(recordListForC, C.int(maxDegree), C.int(buildOutDegree), C.int(maxEntities), C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4020, recordList, maxDegree, buildOutDegree, maxEntities, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(44, recordList, maxDegree, buildOutDegree, maxDegree, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindPathByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int) (string, error) {
	//  _DLEXPORT int G2_findPathByEntityID(const long long entityID1, const long long entityID2, const int maxDegree, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(45, entityID1, entityID2, maxDegree)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_findPathByEntityID_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4021, entityID1, entityID2, maxDegree, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(46, entityID1, entityID2, maxDegree, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindPathByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathByEntityID_V2(const long long entityID1, const long long entityID2, const int maxDegree, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(47, entityID1, entityID2, maxDegree, flags)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_findPathByEntityID_V2_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree), C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4022, entityID1, entityID2, maxDegree, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(48, entityID1, entityID2, maxDegree, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindPathByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int) (string, error) {
	//  _DLEXPORT int G2_findPathByRecordID(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(49, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree)
	}
	entryTime := time.Now()
	var err error = nil
	dataSource1CodeForC := C.CString(dataSourceCode1)
	defer C.free(unsafe.Pointer(dataSource1CodeForC))
	recordID1ForC := C.CString(recordID1)
	defer C.free(unsafe.Pointer(recordID1ForC))
	dataSource2CodeForC := C.CString(dataSourceCode2)
	defer C.free(unsafe.Pointer(dataSource2CodeForC))
	recordID2ForC := C.CString(recordID2)
	defer C.free(unsafe.Pointer(recordID2ForC))
	result := C.G2_findPathByRecordID_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC, C.int(maxDegree))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4023, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(50, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindPathByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathByRecordID_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(51, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, flags)
	}
	entryTime := time.Now()
	var err error = nil
	dataSource1CodeForC := C.CString(dataSourceCode1)
	defer C.free(unsafe.Pointer(dataSource1CodeForC))
	recordID1ForC := C.CString(recordID1)
	defer C.free(unsafe.Pointer(recordID1ForC))
	dataSource2CodeForC := C.CString(dataSourceCode2)
	defer C.free(unsafe.Pointer(dataSource2CodeForC))
	recordID2ForC := C.CString(recordID2)
	defer C.free(unsafe.Pointer(recordID2ForC))
	result := C.G2_findPathByRecordID_V2_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC, C.int(maxDegree), C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4024, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(52, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindPathExcludingByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByEntityID(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(53, entityID1, entityID2, maxDegree, excludedEntities)
	}
	entryTime := time.Now()
	var err error = nil
	excludedEntitiesForC := C.CString(excludedEntities)
	defer C.free(unsafe.Pointer(excludedEntitiesForC))
	result := C.G2_findPathExcludingByEntityID_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree), excludedEntitiesForC)
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4025, entityID1, entityID2, maxDegree, excludedEntities, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(54, entityID1, entityID2, maxDegree, excludedEntities, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindPathExcludingByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByEntityID_V2(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(55, entityID1, entityID2, maxDegree, excludedEntities, flags)
	}
	entryTime := time.Now()
	var err error = nil
	excludedEntitiesForC := C.CString(excludedEntities)
	defer C.free(unsafe.Pointer(excludedEntitiesForC))
	result := C.G2_findPathExcludingByEntityID_V2_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree), excludedEntitiesForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4026, entityID1, entityID2, maxDegree, excludedEntities, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(56, entityID1, entityID2, maxDegree, excludedEntities, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindPathExcludingByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByRecordID(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(57, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords)
	}
	entryTime := time.Now()
	var err error = nil
	dataSource1CodeForC := C.CString(dataSourceCode1)
	defer C.free(unsafe.Pointer(dataSource1CodeForC))
	recordID1ForC := C.CString(recordID1)
	defer C.free(unsafe.Pointer(recordID1ForC))
	dataSource2CodeForC := C.CString(dataSourceCode2)
	defer C.free(unsafe.Pointer(dataSource2CodeForC))
	recordID2ForC := C.CString(recordID2)
	defer C.free(unsafe.Pointer(recordID2ForC))
	excludedRecordsForC := C.CString(excludedRecords)
	defer C.free(unsafe.Pointer(excludedRecordsForC))
	result := C.G2_findPathExcludingByRecordID_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC, C.int(maxDegree), excludedRecordsForC)
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4027, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(58, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindPathExcludingByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByRecordID_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(59, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, flags)
	}
	entryTime := time.Now()
	var err error = nil
	dataSource1CodeForC := C.CString(dataSourceCode1)
	defer C.free(unsafe.Pointer(dataSource1CodeForC))
	recordID1ForC := C.CString(recordID1)
	defer C.free(unsafe.Pointer(recordID1ForC))
	dataSource2CodeForC := C.CString(dataSourceCode2)
	defer C.free(unsafe.Pointer(dataSource2CodeForC))
	recordID2ForC := C.CString(recordID2)
	defer C.free(unsafe.Pointer(recordID2ForC))
	excludedRecordsForC := C.CString(excludedRecords)
	defer C.free(unsafe.Pointer(excludedRecordsForC))
	result := C.G2_findPathExcludingByRecordID_V2_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC, C.int(maxDegree), excludedRecordsForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4028, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(60, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, requiredDsrcs string) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByEntityID(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const char* requiredDsrcs, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(61, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs)
	}
	entryTime := time.Now()
	var err error = nil
	excludedEntitiesForC := C.CString(excludedEntities)
	defer C.free(unsafe.Pointer(excludedEntitiesForC))
	requiredDsrcsForC := C.CString(requiredDsrcs)
	defer C.free(unsafe.Pointer(requiredDsrcsForC))
	result := C.G2_findPathIncludingSourceByEntityID_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree), excludedEntitiesForC, requiredDsrcsForC)
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4029, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(62, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, requiredDsrcs string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByEntityID_V2(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const char* requiredDsrcs, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(63, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, flags)
	}
	entryTime := time.Now()
	var err error = nil
	excludedEntitiesForC := C.CString(excludedEntities)
	defer C.free(unsafe.Pointer(excludedEntitiesForC))
	requiredDsrcsForC := C.CString(requiredDsrcs)
	defer C.free(unsafe.Pointer(requiredDsrcsForC))
	result := C.G2_findPathIncludingSourceByEntityID_V2_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree), excludedEntitiesForC, requiredDsrcsForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4030, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(64, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByRecordID(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const char* requiredDsrcs, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(65, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs)
	}
	entryTime := time.Now()
	var err error = nil
	dataSource1CodeForC := C.CString(dataSourceCode1)
	defer C.free(unsafe.Pointer(dataSource1CodeForC))
	recordID1ForC := C.CString(recordID1)
	defer C.free(unsafe.Pointer(recordID1ForC))
	dataSource2CodeForC := C.CString(dataSourceCode2)
	defer C.free(unsafe.Pointer(dataSource2CodeForC))
	recordID2ForC := C.CString(recordID2)
	defer C.free(unsafe.Pointer(recordID2ForC))
	excludedRecordsForC := C.CString(excludedRecords)
	defer C.free(unsafe.Pointer(excludedRecordsForC))
	requiredDsrcsForC := C.CString(requiredDsrcs)
	defer C.free(unsafe.Pointer(requiredDsrcsForC))
	result := C.G2_findPathIncludingSourceByRecordID_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC, C.int(maxDegree), excludedRecordsForC, requiredDsrcsForC)
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4031, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(66, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByRecordID_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const char* requiredDsrcs, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(67, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, flags)
	}
	entryTime := time.Now()
	var err error = nil
	dataSource1CodeForC := C.CString(dataSourceCode1)
	defer C.free(unsafe.Pointer(dataSource1CodeForC))
	recordID1ForC := C.CString(recordID1)
	defer C.free(unsafe.Pointer(recordID1ForC))
	dataSource2CodeForC := C.CString(dataSourceCode2)
	defer C.free(unsafe.Pointer(dataSource2CodeForC))
	recordID2ForC := C.CString(recordID2)
	defer C.free(unsafe.Pointer(recordID2ForC))
	excludedRecordsForC := C.CString(excludedRecords)
	defer C.free(unsafe.Pointer(excludedRecordsForC))
	requiredDsrcsForC := C.CString(requiredDsrcs)
	defer C.free(unsafe.Pointer(requiredDsrcsForC))
	result := C.G2_findPathIncludingSourceByRecordID_V2_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC, C.int(maxDegree), excludedRecordsForC, requiredDsrcsForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4032, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(68, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) GetActiveConfigID(ctx context.Context) (int64, error) {
	//  _DLEXPORT int G2_getActiveConfigID(long long* configID);
	if g2engine.isTrace {
		g2engine.traceEntry(69)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_getActiveConfigID_helper()
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4033, result.returnCode, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(70, int64(C.longlong(result.configID)), err, time.Since(entryTime))
	}
	return int64(C.longlong(result.configID)), err
}

func (g2engine *G2engineImpl) GetEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2_getEntityByEntityID(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(71, entityID)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_getEntityByEntityID_helper(C.longlong(entityID))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4034, entityID, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(72, entityID, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) GetEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_getEntityByEntityID_V2(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(73, entityID, flags)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_getEntityByEntityID_V2_helper(C.longlong(entityID), C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4035, entityID, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(74, entityID, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) GetEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	//  _DLEXPORT int G2_getEntityByRecordID(const char* dataSourceCode, const char* recordID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(75, dataSourceCode, recordID)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	result := C.G2_getEntityByRecordID_helper(dataSourceCodeForC, recordIDForC)
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4036, dataSourceCode, recordID, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(76, dataSourceCode, recordID, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) GetEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_getEntityByRecordID_V2(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(77, dataSourceCode, recordID, flags)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	result := C.G2_getEntityByRecordID_V2_helper(dataSourceCodeForC, recordIDForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4037, dataSourceCode, recordID, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(78, dataSourceCode, recordID, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

// GetLastException returns the last exception encountered in the Senzing Engine.
func (g2engine *G2engineImpl) GetLastException(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_getLastException(char *buffer, const size_t bufSize);
	if g2engine.isTrace {
		g2engine.traceEntry(79)
	}
	entryTime := time.Now()
	var err error = nil
	stringBuffer := g2engine.getByteArray(initialByteArraySize)
	C.G2_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	// if result == 0 { // "result" is length of exception message.
	// 	err = g2engine.getLogger().Error(4038, result, time.Since(entryTime))
	// }
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if g2engine.isTrace {
		defer g2engine.traceExit(80, string(stringBuffer), err, time.Since(entryTime))
	}
	return string(stringBuffer), err
}

func (g2engine *G2engineImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2_getLastExceptionCode();
	if g2engine.isTrace {
		g2engine.traceEntry(81)
	}
	entryTime := time.Now()
	var err error = nil
	result := int(C.G2_getLastExceptionCode())
	if g2engine.isTrace {
		defer g2engine.traceExit(82, result, err, time.Since(entryTime))
	}
	return result, err
}

func (g2engine *G2engineImpl) GetRecord(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	//  _DLEXPORT int G2_getRecord(const char* dataSourceCode, const char* recordID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(83, dataSourceCode, recordID)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	result := C.G2_getRecord_helper(dataSourceCodeForC, recordIDForC)
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4039, dataSourceCode, recordID, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(84, dataSourceCode, recordID, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) GetRecord_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_getRecord_V2(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(85, dataSourceCode, recordID, flags)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	result := C.G2_getRecord_V2_helper(dataSourceCodeForC, recordIDForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4040, dataSourceCode, recordID, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(86, dataSourceCode, recordID, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) GetRedoRecord(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_getRedoRecord(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2engine.isTrace {
		g2engine.traceEntry(87)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_getRedoRecord_helper()
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4041, result.returnCode, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(88, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err

}

func (g2engine *G2engineImpl) GetRepositoryLastModifiedTime(ctx context.Context) (int64, error) {
	//  _DLEXPORT int G2_getRepositoryLastModifiedTime(long long* lastModifiedTime);
	if g2engine.isTrace {
		g2engine.traceEntry(89)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_getRepositoryLastModifiedTime_helper()
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4042, result.returnCode, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(90, int64(result.time), err, time.Since(entryTime))
	}
	return int64(result.time), err
}

func (g2engine *G2engineImpl) GetVirtualEntityByRecordID(ctx context.Context, recordList string) (string, error) {
	//  _DLEXPORT int G2_getVirtualEntityByRecordID(const char* recordList, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(91, recordList)
	}
	entryTime := time.Now()
	var err error = nil
	recordListForC := C.CString(recordList)
	defer C.free(unsafe.Pointer(recordListForC))
	result := C.G2_getVirtualEntityByRecordID_helper(recordListForC)
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4043, recordList, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(92, recordList, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) GetVirtualEntityByRecordID_V2(ctx context.Context, recordList string, flags int64) (string, error) {
	//  _DLEXPORT int G2_getVirtualEntityByRecordID_V2(const char* recordList, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(93, recordList, flags)
	}
	entryTime := time.Now()
	var err error = nil
	recordListForC := C.CString(recordList)
	defer C.free(unsafe.Pointer(recordListForC))
	result := C.G2_getVirtualEntityByRecordID_V2_helper(recordListForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4044, recordList, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(94, recordList, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) HowEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2_howEntityByEntityID(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	entryTime := time.Now()
	if g2engine.isTrace {
		g2engine.traceEntry(95, entityID)
	}
	var err error = nil
	result := C.G2_howEntityByEntityID_helper(C.longlong(entityID))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4045, entityID, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(96, entityID, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) HowEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_howEntityByEntityID_V2(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	entryTime := time.Now()
	if g2engine.isTrace {
		g2engine.traceEntry(97, entityID, flags)
	}
	var err error = nil
	result := C.G2_howEntityByEntityID_V2_helper(C.longlong(entityID), C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4046, entityID, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(98, entityID, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

// Init initializes the Senzing G2diagnosis.
func (g2engine *G2engineImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	if g2engine.isTrace {
		g2engine.traceEntry(99, moduleName, iniParams, verboseLogging)
	}
	entryTime := time.Now()
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))
	if result != 0 {
		err = g2engine.newError(ctx, 4047, moduleName, iniParams, verboseLogging, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(100, moduleName, iniParams, verboseLogging, err, time.Since(entryTime))
	}
	return err
}

func (g2engine *G2engineImpl) InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error {
	//  _DLEXPORT int G2_initWithConfigID(const char *moduleName, const char *iniParams, const long long initConfigID, const int verboseLogging);
	if g2engine.isTrace {
		g2engine.traceEntry(101, moduleName, iniParams, initConfigID, verboseLogging)
	}
	entryTime := time.Now()
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2_initWithConfigID(moduleNameForC, iniParamsForC, C.longlong(initConfigID), C.int(verboseLogging))
	if result != 0 {
		err = g2engine.newError(ctx, 4048, moduleName, iniParams, initConfigID, verboseLogging, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(102, moduleName, iniParams, initConfigID, verboseLogging, err, time.Since(entryTime))
	}
	return err
}

func (g2engine *G2engineImpl) PrimeEngine(ctx context.Context) error {
	//  _DLEXPORT int G2_primeEngine();
	if g2engine.isTrace {
		g2engine.traceEntry(103)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_primeEngine()
	if result != 0 {
		err = g2engine.newError(ctx, 4049, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(104, err, time.Since(entryTime))
	}
	return err
}

func (g2engine *G2engineImpl) Process(ctx context.Context, record string) error {
	//  _DLEXPORT int G2_process(const char *record);
	if g2engine.isTrace {
		g2engine.traceEntry(105, record)
	}
	entryTime := time.Now()
	var err error = nil
	recordForC := C.CString(record)
	defer C.free(unsafe.Pointer(recordForC))
	result := C.G2_process(recordForC)
	if result != 0 {
		err = g2engine.newError(ctx, 4050, record, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(106, record, err, time.Since(entryTime))
	}
	return err
}

func (g2engine *G2engineImpl) ProcessRedoRecord(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_processRedoRecord(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2engine.isTrace {
		g2engine.traceEntry(107)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_processRedoRecord_helper()
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4051, result.returnCode, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(108, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err

}

func (g2engine *G2engineImpl) ProcessRedoRecordWithInfo(ctx context.Context, flags int64) (string, string, error) {
	//  _DLEXPORT int G2_processRedoRecordWithInfo(const long long flags, char **responseBuf, size_t *bufSize, char **infoBuf, size_t *infoBufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(109, flags)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_processRedoRecordWithInfo_helper(C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4052, flags, result.returnCode, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(110, flags, C.GoString(result.response), C.GoString(result.withInfo), err, time.Since(entryTime))
	}
	return C.GoString(result.response), C.GoString(result.withInfo), err
}

func (g2engine *G2engineImpl) ProcessWithInfo(ctx context.Context, record string, flags int64) (string, error) {
	//  _DLEXPORT int G2_processWithInfo(const char *record, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(111, record, flags)
	}
	entryTime := time.Now()
	var err error = nil
	recordForC := C.CString(record)
	defer C.free(unsafe.Pointer(recordForC))
	result := C.G2_processWithInfo_helper(recordForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4053, record, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(112, record, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) ProcessWithResponse(ctx context.Context, record string) (string, error) {
	//  _DLEXPORT int G2_processWithResponse(const char *record, char *responseBuf, const size_t bufSize);
	if g2engine.isTrace {
		g2engine.traceEntry(113, record)
	}
	entryTime := time.Now()
	var err error = nil
	recordForC := C.CString(record)
	defer C.free(unsafe.Pointer(recordForC))
	result := C.G2_processWithResponse_helper(recordForC)
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4054, record, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(114, record, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) ProcessWithResponseResize(ctx context.Context, record string) (string, error) {
	//  _DLEXPORT int G2_processWithResponseResize(const char *record, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2engine.isTrace {
		g2engine.traceEntry(115, record)
	}
	entryTime := time.Now()
	var err error = nil
	recordForC := C.CString(record)
	defer C.free(unsafe.Pointer(recordForC))
	result := C.G2_processWithResponseResize_helper(recordForC)
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4055, record, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(116, record, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) PurgeRepository(ctx context.Context) error {
	//  _DLEXPORT int G2_purgeRepository();
	if g2engine.isTrace {
		g2engine.traceEntry(117)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_purgeRepository()
	if result != 0 {
		err = g2engine.newError(ctx, 4056, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(118, err, time.Since(entryTime))
	}
	return err
}

func (g2engine *G2engineImpl) ReevaluateEntity(ctx context.Context, entityID int64, flags int64) error {
	//  _DLEXPORT int G2_reevaluateEntity(const long long entityID, const long long flags);
	if g2engine.isTrace {
		g2engine.traceEntry(119, entityID, flags)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_reevaluateEntity(C.longlong(entityID), C.longlong(flags))
	if result != 0 {
		err = g2engine.newError(ctx, 4057, entityID, flags, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(120, entityID, flags, err, time.Since(entryTime))
	}
	return err
}

func (g2engine *G2engineImpl) ReevaluateEntityWithInfo(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_reevaluateEntityWithInfo(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(121, entityID, flags)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_reevaluateEntityWithInfo_helper(C.longlong(entityID), C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4058, entityID, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(122, entityID, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) ReevaluateRecord(ctx context.Context, dataSourceCode string, recordID string, flags int64) error {
	//  _DLEXPORT int G2_reevaluateRecord(const char* dataSourceCode, const char* recordID, const long long flags);
	if g2engine.isTrace {
		g2engine.traceEntry(123, dataSourceCode, recordID, flags)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	result := C.G2_reevaluateRecord(dataSourceCodeForC, recordIDForC, C.longlong(flags))
	if result != 0 {
		err = g2engine.newError(ctx, 4059, dataSourceCode, recordID, flags, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(124, dataSourceCode, recordID, flags, err, time.Since(entryTime))
	}
	return err
}

func (g2engine *G2engineImpl) ReevaluateRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_reevaluateRecordWithInfo(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(125, dataSourceCode, recordID, flags)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	result := C.G2_reevaluateRecordWithInfo_helper(dataSourceCodeForC, recordIDForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4060, dataSourceCode, recordID, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(126, dataSourceCode, recordID, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) Reinit(ctx context.Context, initConfigID int64) error {
	//  _DLEXPORT int G2_reinit(const long long initConfigID);
	if g2engine.isTrace {
		g2engine.traceEntry(127, initConfigID)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_reinit(C.longlong(initConfigID))
	if result != 0 {
		err = g2engine.newError(ctx, 4061, initConfigID, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(128, initConfigID, err, time.Since(entryTime))
	}
	return err
}

func (g2engine *G2engineImpl) ReplaceRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error {
	//  _DLEXPORT int G2_replaceRecord(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID);
	if g2engine.isTrace {
		g2engine.traceEntry(129, dataSourceCode, recordID, jsonData, loadID)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))
	result := C.G2_replaceRecord(dataSourceCodeForC, recordIDForC, jsonDataForC, loadIDForC)
	if result != 0 {
		err = g2engine.newError(ctx, 4062, dataSourceCode, recordID, jsonData, loadID, result, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(130, dataSourceCode, recordID, jsonData, loadID, err, time.Since(entryTime))
	}
	return err
}

func (g2engine *G2engineImpl) ReplaceRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_replaceRecordWithInfo(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(131, dataSourceCode, recordID, jsonData, loadID, flags)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))
	result := C.G2_replaceRecordWithInfo_helper(dataSourceCodeForC, recordIDForC, jsonDataForC, loadIDForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4063, dataSourceCode, recordID, jsonData, loadID, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(132, dataSourceCode, recordID, jsonData, loadID, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) SearchByAttributes(ctx context.Context, jsonData string) (string, error) {
	//  _DLEXPORT int G2_searchByAttributes(const char* jsonData, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(133, jsonData)
	}
	entryTime := time.Now()
	var err error = nil
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	result := C.G2_searchByAttributes_helper(jsonDataForC)
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4064, jsonData, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(134, jsonData, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) SearchByAttributes_V2(ctx context.Context, jsonData string, flags int64) (string, error) {
	//  _DLEXPORT int G2_searchByAttributes_V2(const char* jsonData, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(135, jsonData, flags)
	}
	entryTime := time.Now()
	var err error = nil
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	result := C.G2_searchByAttributes_V2_helper(jsonDataForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4065, jsonData, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(136, jsonData, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) SetLogLevel(ctx context.Context, logLevel logger.Level) error {
	if g2engine.isTrace {
		g2engine.traceEntry(137, logLevel)
	}
	entryTime := time.Now()
	var err error = nil
	g2engine.getLogger().SetLogLevel(messagelogger.Level(logLevel))
	g2engine.isTrace = (g2engine.getLogger().GetLogLevel() == messagelogger.LevelTrace)
	if g2engine.isTrace {
		defer g2engine.traceExit(138, logLevel, err, time.Since(entryTime))
	}
	return err
}

func (g2engine *G2engineImpl) Stats(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_stats(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2engine.isTrace {
		g2engine.traceEntry(139)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_stats_helper()
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4066, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(140, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) WhyEntities(ctx context.Context, entityID1 int64, entityID2 int64) (string, error) {
	//  _DLEXPORT int G2_whyEntities(const long long entityID1, const long long entityID2, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(141, entityID1, entityID2)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_whyEntities_helper(C.longlong(entityID1), C.longlong(entityID2))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4067, entityID1, entityID2, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(142, entityID1, entityID2, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) WhyEntities_V2(ctx context.Context, entityID1 int64, entityID2 int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyEntities_V2(const long long entityID1, const long long entityID2, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(143, entityID1, entityID2, flags)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_whyEntities_V2_helper(C.longlong(entityID1), C.longlong(entityID2), C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4068, entityID1, entityID2, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(144, entityID1, entityID2, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) WhyEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2_whyEntityByEntityID(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(145, entityID)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_whyEntityByEntityID_helper(C.longlong(entityID))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4069, entityID, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(146, entityID, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) WhyEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyEntityByEntityID_V2(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(147, entityID, flags)
	}
	entryTime := time.Now()
	var err error = nil
	result := C.G2_whyEntityByEntityID_V2_helper(C.longlong(entityID), C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4070, entityID, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(148, entityID, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) WhyEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	//  _DLEXPORT int G2_whyEntityByRecordID(const char* dataSourceCode, const char* recordID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(149, dataSourceCode, recordID)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	result := C.G2_whyEntityByRecordID_helper(dataSourceCodeForC, recordIDForC)
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4071, dataSourceCode, recordID, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(150, dataSourceCode, recordID, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) WhyEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyEntityByRecordID_V2(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(151, dataSourceCode, recordID, flags)
	}
	entryTime := time.Now()
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	result := C.G2_whyEntityByRecordID_V2_helper(dataSourceCodeForC, recordIDForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4072, dataSourceCode, recordID, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(152, dataSourceCode, recordID, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) WhyRecords(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string) (string, error) {
	//  _DLEXPORT int G2_whyRecords(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(153, dataSourceCode1, recordID1, dataSourceCode2, recordID2)
	}
	entryTime := time.Now()
	var err error = nil
	dataSource1CodeForC := C.CString(dataSourceCode1)
	defer C.free(unsafe.Pointer(dataSource1CodeForC))
	recordID1ForC := C.CString(recordID1)
	defer C.free(unsafe.Pointer(recordID1ForC))
	dataSource2CodeForC := C.CString(dataSourceCode2)
	defer C.free(unsafe.Pointer(dataSource2CodeForC))
	recordID2ForC := C.CString(recordID2)
	defer C.free(unsafe.Pointer(recordID2ForC))
	result := C.G2_whyRecords_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC)
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4073, dataSourceCode1, recordID1, dataSourceCode2, recordID2, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(154, dataSourceCode1, recordID1, dataSourceCode2, recordID2, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) WhyRecords_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyRecords_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(155, dataSourceCode1, recordID1, dataSourceCode2, recordID2, flags)
	}
	entryTime := time.Now()
	var err error = nil
	dataSource1CodeForC := C.CString(dataSourceCode1)
	defer C.free(unsafe.Pointer(dataSource1CodeForC))
	recordID1ForC := C.CString(recordID1)
	defer C.free(unsafe.Pointer(recordID1ForC))
	dataSource2CodeForC := C.CString(dataSourceCode2)
	defer C.free(unsafe.Pointer(dataSource2CodeForC))
	recordID2ForC := C.CString(recordID2)
	defer C.free(unsafe.Pointer(recordID2ForC))
	result := C.G2_whyRecords_V2_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.newError(ctx, 4074, dataSourceCode1, recordID1, dataSourceCode2, recordID2, flags, result.returnCode, time.Since(entryTime))
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(156, dataSourceCode1, recordID1, dataSourceCode2, recordID2, flags, C.GoString(result.response), err, time.Since(entryTime))
	}
	return C.GoString(result.response), err
}
