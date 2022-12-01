/*
// The G2engineImpl implementation is a wrapper over the Senzing libg2 library.
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

// G2engineImpl is the default implementation of the G2engine interface.
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

// Make a byte array.
func (g2engine *G2engineImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

// Create a new error.
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

// Get the Logger singleton.
func (g2engine *G2engineImpl) getLogger() messagelogger.MessageLoggerInterface {
	if g2engine.logger == nil {
		g2engine.logger, _ = messagelogger.NewSenzingApiLogger(ProductId, IdMessages, IdStatuses, messagelogger.LevelInfo)
	}
	return g2engine.logger
}

// Trace method entry.
func (g2engine *G2engineImpl) traceEntry(errorNumber int, details ...interface{}) {
	g2engine.getLogger().Log(errorNumber, details...)
}

// Trace method exit.
func (g2engine *G2engineImpl) traceExit(errorNumber int, details ...interface{}) {
	g2engine.getLogger().Log(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The AddRecord method adds a record into the Senzing repository.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
*/
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

/*
The AddRecordWithInfo method adds a record into the Senzing repository and returns information on the affected entities.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    Example: `{"DATA_SOURCE":"TEST","RECORD_ID":"333","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}`
*/
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

/*
The AddRecordWithInfoWithReturnedRecordID method adds a record into the Senzing repository and returns information on the affected entities and the record identifier.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
  - flags: Flags used to control information returned.

Output
  - A JSON document containing the AFFECTED_ENTITIES, INTERESTING_ENTITIES, and RECORD_ID.
    Example: `{"DATA_SOURCE":"TEST","RECORD_ID":"2D4DABB3FAEAFBD452E9487D06FABC22DC69C846","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}`
  - The record identifier.
    Example: 2D4DABB3FAEAFBD452E9487D06FABC22DC69C846
*/
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

/*
The AddRecordWithReturnedRecordID method adds a record into the Senzing repository and returns the record identifier.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.

Output
  - The record identifier.
    Example: 2D4DABB3FAEAFBD452E9487D06FABC22DC69C846
*/
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

/*
The CheckRecord method FIXME:.

Input
  - ctx: A context to control lifecycle.
  - record: A JSON document with the attribute data for the record to check with the "DATA_SOURCE" field.
  - recordQueryList: A JSON document with the datasource codes and recordID's of the records to check against.

Output
  - A JSON document that FIXME:
    Example: `{"CHECK_RECORD_RESPONSE":[{"DSRC_CODE":"TEST","RECORD_ID":"111","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","MATCH_KEY":"","ERRULE_CODE":"","ERRULE_ID":0,"CANDIDATE_MATCH":"N","NON_GENERIC_CANDIDATE_MATCH":"N"}]}`
*/
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

/*
The ClearLastException method erases the last exception message held by the Senzing G2 object.

Input
  - ctx: A context to control lifecycle.
*/
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

/*
The CloseExport method closes the exported document created by ExportJSONEntityReport().
It is part of the ExportJSONEntityReport(), FetchNext(), CloseExport()
lifecycle of a list of sized entities.

Input
  - ctx: A context to control lifecycle.
  - responseHandle: A handle created by ExportJSONEntityReport().
*/
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

/*
The CountRedoRecords method returns the number of records in need of redo-ing.

Input
  - ctx: A context to control lifecycle.

Output
  - The number of redo records in Senzing's redo queue.
*/
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

/*
The DeleteRecord method deletes a record from the Senzing repository.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
    FIXME: How does the "loadID" affect what is deleted?
*/
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

/*
The DeleteRecordWithInfo method deletes a record from the Senzing repository and returns information on the affected entities.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
    FIXME: How does the "loadID" affect what is deleted?
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    Example: `{"DATA_SOURCE":"TEST","RECORD_ID":"111","AFFECTED_ENTITIES":[],"INTERESTING_ENTITIES":{"ENTITIES":[]}}`
*/
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

/*
The Destroy method will destroy and perform cleanup for the Senzing G2 object.
It should be called after all other calls are complete.

Input
  - ctx: A context to control lifecycle.
*/
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

/*
The ExportConfig method returns the Senzing engine configuration.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document.
*/
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

/*
Similar to ExportConfig(), the ExportConfigAndConfigID method returns the Senzing engine configuration and it's identifier.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document.
  - The unique identifier of the configuration.
*/
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

/*
The ExportCSVEntityReport method initializes a cursor over a document of exported entities.
It is part of the ExportCSVEntityReport(), FetchNext(), CloseExport()
lifecycle of a list of entities to export.

Input
  - ctx: A context to control lifecycle.
  - csvColumnList: A comma-separated list of column names for the CSV export.
  - flags: Any combination of G2_EXPORT_ flags to control what is exported.

Output
  - A handle that identifies the document to be scrolled through using FetchNext().
*/
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

/*
The ExportJSONEntityReport method initializes a cursor over a document of exported entities.
It is part of the ExportJSONEntityReport(), FetchNext(), CloseExport()
lifecycle of a list of entities to export.

Input
  - ctx: A context to control lifecycle.
  - flags: Any combination of G2_EXPORT_ flags to control what is exported.

Output
  - A handle that identifies the document to be scrolled through using FetchNext().
*/
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

/*
The FetchNext method is used to scroll through an exported document.
It is part of the ExportJSONEntityReport() or ExportCSVEntityReport(), FetchNext(), CloseExport()
lifecycle of a list of exported entities.

Input
  - ctx: A context to control lifecycle.
  - responseHandle: A handle created by ExportJSONEntityReport() or ExportCSVEntityReport().

Output
  - FIXME:
*/
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

/*
The FindInterestingEntitiesByEntityID method FIXME:

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"INTERESTING_ENTITIES":{"ENTITIES":[]}}`
*/
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

/*
The FindInterestingEntitiesByRecordID method FIXME:

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"INTERESTING_ENTITIES":{"ENTITIES":[]}}`
*/
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

/*
The FindNetworkByEntityID method finds all entities surrounding a requested set of entities.
This includes the requested entities, paths between them, and relations to other nearby entities.
To control output, use FindNetworkByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityList: A JSON document listing entities.
    Example: `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
  - maxDegree: The maximum number of degrees in paths between search entities.
  - buildOutDegree: The number of degrees of relationships to show around each search entity.
  - maxEntities: The maximum number of entities to return in the discovered network.

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"SEAMAN","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-11-29 22:25:18.997","LAST_SEEN_DT":"2022-11-29 22:25:19.005"}],"LAST_SEEN_DT":"2022-11-29 22:25:19.005"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-11-29 22:25:19.009","LAST_SEEN_DT":"2022-11-29 22:25:19.009"}],"LAST_SEEN_DT":"2022-11-29 22:25:19.009"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
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

/*
The FindNetworkByEntityID_V2 method finds all entities surrounding a requested set of entities.
This includes the requested entities, paths between them, and relations to other nearby entities.
It extends FindNetworkByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityList: A JSON document listing entities.
    Example: `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
  - maxDegree: The maximum number of degrees in paths between search entities.
  - buildOutDegree: The number of degrees of relationships to show around each search entity.
  - maxEntities: The maximum number of entities to return in the discovered network.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}}]}`
*/
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

/*
The FindNetworkByRecordID method finds all entities surrounding a requested set of entities identified by record identifiers.
This includes the requested entities, paths between them, and relations to other nearby entities.
To control output, use FindNetworkByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityList: A JSON document listing entities.
    Example: `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
  - maxDegree: The maximum number of degrees in paths between search entities.
  - buildOutDegree: The number of degrees of relationships to show around each search entity.
  - maxEntities: The maximum number of entities to return in the discovered network.

Output
  - A JSON document.
    Example:
*/
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

/*
The FindNetworkByRecordID_V2 method finds all entities surrounding a requested set of entities identified by record identifiers.
This includes the requested entities, paths between them, and relations to other nearby entities.
It extends FindNetworkByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityList: A JSON document listing entities.
    Example: `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
  - maxDegree: The maximum number of degrees in paths between search entities.
  - buildOutDegree: The number of degrees of relationships to show around each search entity.
  - maxEntities: The maximum number of entities to return in the discovered network.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}},{"RESOLVED_ENTITY":{"ENTITY_ID":3}}]}`
*/
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

/*
The FindPathByEntityID method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
To control output, use FindPathByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"SEAMAN","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-11-30 15:10:20.351","LAST_SEEN_DT":"2022-11-30 15:10:20.488"}],"LAST_SEEN_DT":"2022-11-30 15:10:20.488"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-11-30 15:10:20.426","LAST_SEEN_DT":"2022-11-30 15:10:20.426"}],"LAST_SEEN_DT":"2022-11-30 15:10:20.426"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
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

/*
The FindPathByEntityID_V2 method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
It extends FindPathByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}}]}`
*/
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

/*
The FindPathByRecordID method finds single relationship paths between two entities.
The entities are identified by starting and ending records.
Paths are found using known relationships with other entities.
To control output, use FindPathByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"SEAMAN","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-11-30 15:31:32.996","LAST_SEEN_DT":"2022-11-30 15:31:33.128"}],"LAST_SEEN_DT":"2022-11-30 15:31:33.128"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-11-30 15:31:33.068","LAST_SEEN_DT":"2022-11-30 15:31:33.068"}],"LAST_SEEN_DT":"2022-11-30 15:31:33.068"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
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

/*
The FindPathByRecordID_V2 method finds single relationship paths between two entities.
The entities are identified by starting and ending records.
Paths are found using known relationships with other entities.
It extends FindPathByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"SEAMAN","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-11-30 15:31:32.996","LAST_SEEN_DT":"2022-11-30 15:31:33.128"}],"LAST_SEEN_DT":"2022-11-30 15:31:33.128"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-11-30 15:31:33.068","LAST_SEEN_DT":"2022-11-30 15:31:33.068"}],"LAST_SEEN_DT":"2022-11-30 15:31:33.068"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
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

/*
The FindPathExcludingByEntityID method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
In addition, it will find paths that exclude certain entities from being on the path.
To control output, use FindPathExcludingByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"SEAMAN","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-11-30 15:42:04.049","LAST_SEEN_DT":"2022-11-30 15:42:04.209"}],"LAST_SEEN_DT":"2022-11-30 15:42:04.209"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-11-30 15:42:04.148","LAST_SEEN_DT":"2022-11-30 15:42:04.148"}],"LAST_SEEN_DT":"2022-11-30 15:42:04.148"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
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

/*
The FindPathExcludingByEntityID_V2 method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
In addition, it will find paths that exclude certain entities from being on the path.
It extends FindPathExcludingByEntityID() by adding output control flags.

When excluding entities, the user may choose to either strictly exclude the entities,
or prefer to exclude the entities but still include them if no other path is found.
By default, entities will be strictly excluded.
A "preferred exclude" may be done by specifying the G2_FIND_PATH_PREFER_EXCLUDE control flag.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.
  - flags: FIXME:

Output
  - A JSON document.
    Example:
*/
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

/*
The FindPathExcludingByRecordID method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
In addition, it will find paths that exclude certain entities from being on the path.
To control output, use FindPathExcludingByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"SEAMAN","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-11-30 15:57:10.826","LAST_SEEN_DT":"2022-11-30 15:57:10.959"}],"LAST_SEEN_DT":"2022-11-30 15:57:10.959"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-11-30 15:57:10.898","LAST_SEEN_DT":"2022-11-30 15:57:10.898"}],"LAST_SEEN_DT":"2022-11-30 15:57:10.898"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
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

/*
The FindPathExcludingByRecordID_V2 method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
In addition, it will find paths that exclude certain entities from being on the path.
It extends FindPathExcludingByRecordID() by adding output control flags.

When excluding entities, the user may choose to either strictly exclude the entities,
or prefer to exclude the entities but still include them if no other path is found.
By default, entities will be strictly excluded.
A "preferred exclude" may be done by specifying the G2_FIND_PATH_PREFER_EXCLUDE control flag.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}}]}`
*/
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

/*
The FindPathIncludingSourceByEntityID method finds single relationship paths between two entities.
In addition, one of the enties along the path must include a specified data source.
Specific entities may also be excluded,
using the same methodology as the FindPathExcludingByEntityID() and FindPathExcludingByRecordID().
To control output, use FindPathIncludingSourceByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.
  - requiredDsrcs: A JSON document listing data sources that should be included on the path. FIXME:

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"SEAMAN","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-11-30 16:20:53.885","LAST_SEEN_DT":"2022-11-30 16:20:54.020"}],"LAST_SEEN_DT":"2022-11-30 16:20:54.020"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-11-30 16:20:53.956","LAST_SEEN_DT":"2022-11-30 16:20:53.956"}],"LAST_SEEN_DT":"2022-11-30 16:20:53.956"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
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

/*
The FindPathIncludingSourceByEntityID_V2 method finds single relationship paths between two entities.
In addition, one of the enties along the path must include a specified data source.
Specific entities may also be excluded,
using the same methodology as the FindPathExcludingByEntityID_V2() and FindPathExcludingByRecordID_V2().
It extends FindPathIncludingSourceByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.
  - requiredDsrcs: A JSON document listing data sources that should be included on the path. FIXME:
  - flags: FIXME:

Output
  - A JSON document.
    Example: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}}]}
*/
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

/*
The FindPathIncludingSourceByRecordID method finds single relationship paths between two entities.
In addition, one of the enties along the path must include a specified data source.
Specific entities may also be excluded,
using the same methodology as the FindPathExcludingByEntityID() and FindPathExcludingByRecordID().
To control output, use FindPathIncludingSourceByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.
  - requiredDsrcs: A JSON document listing data sources that should be included on the path. FIXME:

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"SEAMAN","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-11-30 16:38:41.531","LAST_SEEN_DT":"2022-11-30 16:38:41.667"}],"LAST_SEEN_DT":"2022-11-30 16:38:41.667"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-11-30 16:38:41.606","LAST_SEEN_DT":"2022-11-30 16:38:41.606"}],"LAST_SEEN_DT":"2022-11-30 16:38:41.606"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
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

/*
The FindPathIncludingSourceByRecordID method finds single relationship paths between two entities.
In addition, one of the enties along the path must include a specified data source.
Specific entities may also be excluded,
using the same methodology as the FindPathExcludingByEntityID_V2() and FindPathExcludingByRecordID_V2().
It extends FindPathIncludingSourceByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.
  - requiredDsrcs: A JSON document listing data sources that should be included on the path. FIXME:
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}}]}`
*/
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

/*
The GetActiveConfigID method returns the identifier of the loaded Senzing engine configuration.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document.
    Example:
*/
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

/*
The GetEntityByEntityID method returns entity data based on the ID of a resolved identity.
To control output, use GetEntityByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID: The entity ID of the requested data.

Output
  - A JSON document.
    Example: `{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"SEAMAN","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7}]}],"NAME":[{"FEAT_DESC":"SEAMAN","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SEAMAN","LIB_FEAT_ID":1}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-11-30 17:03:40.372","LAST_SEEN_DT":"2022-11-30 17:03:40.503"}],"LAST_SEEN_DT":"2022-11-30 17:03:40.503","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"2E83FCA24AF2996D77569554CC4FD7775F861A8F","ENTITY_DESC":"SEAMAN","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-11-30 17:03:40.372"},{"DATA_SOURCE":"TEST","RECORD_ID":"2D4DABB3FAEAFBD452E9487D06FABC22DC69C846","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"2E83FCA24AF2996D77569554CC4FD7775F861A8F","ENTITY_DESC":"SEAMAN","MATCH_KEY":"+EXACTLY_SAME","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-11-30 17:03:40.503"}]},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-11-30 17:03:40.444","LAST_SEEN_DT":"2022-11-30 17:03:40.444"}],"LAST_SEEN_DT":"2022-11-30 17:03:40.444"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-11-30 17:03:40.510","LAST_SEEN_DT":"2022-11-30 17:03:40.510"}],"LAST_SEEN_DT":"2022-11-30 17:03:40.510"}]}`
*/
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

/*
The GetEntityByEntityID_V2 method returns entity data based on the ID of a resolved identity.
It extends GetEntityByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID: The entity ID of the requested data.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"RESOLVED_ENTITY":{"ENTITY_ID":1}}`
*/
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

/*
The GetEntityByRecordID method returns entity data based on the ID of a record which is a member of the entity.
To control output, use GetEntityByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.

Output
  - A JSON document.
    Example: `{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"SEAMAN","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7}]}],"NAME":[{"FEAT_DESC":"SEAMAN","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SEAMAN","LIB_FEAT_ID":1}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-11-30 17:13:04.312","LAST_SEEN_DT":"2022-11-30 17:13:04.441"}],"LAST_SEEN_DT":"2022-11-30 17:13:04.441","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"2E83FCA24AF2996D77569554CC4FD7775F861A8F","ENTITY_DESC":"SEAMAN","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-11-30 17:13:04.312"},{"DATA_SOURCE":"TEST","RECORD_ID":"2D4DABB3FAEAFBD452E9487D06FABC22DC69C846","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"2E83FCA24AF2996D77569554CC4FD7775F861A8F","ENTITY_DESC":"SEAMAN","MATCH_KEY":"+EXACTLY_SAME","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-11-30 17:13:04.441"}]},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-11-30 17:13:04.384","LAST_SEEN_DT":"2022-11-30 17:13:04.384"}],"LAST_SEEN_DT":"2022-11-30 17:13:04.384"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-11-30 17:13:04.445","LAST_SEEN_DT":"2022-11-30 17:13:04.445"}],"LAST_SEEN_DT":"2022-11-30 17:13:04.445"}]}`
*/
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

/*
The GetEntityByRecordID_V2 method returns entity data based on the ID of a record which is a member of the entity.
It extends GetEntityByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"RESOLVED_ENTITY":{"ENTITY_ID":1}}`
*/
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

/*
The GetLastException method retrieves the last exception thrown in Senzing's G2.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the error received from Senzing's G2Product.
*/
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

/*
The GetLastExceptionCode method retrieves the code of the last exception thrown in Senzing's G2.

Input:
  - ctx: A context to control lifecycle.

Output:
  - An int containing the error received from Senzing's G2Product.
*/
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

/*
The GetRecord method returns a JSON document of a single record from the Senzing repository.
To control output, use GetRecord_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.

Output
  - A JSON document.
    Example: `{"DATA_SOURCE":"TEST","RECORD_ID":"111","JSON_DATA":{"SOCIAL_HANDLE":"flavorh","DATE_OF_BIRTH":"4/8/1983","ADDR_STATE":"LA","ADDR_POSTAL_CODE":"71232","SSN_NUMBER":"053-39-3251","GENDER":"F","srccode":"MDMPER","CC_ACCOUNT_NUMBER":"5534202208773608","ADDR_CITY":"Delhi","DRIVERS_LICENSE_STATE":"DE","PHONE_NUMBER":"225-671-0796","NAME_LAST":"SEAMAN","entityid":"284430058","ADDR_LINE1":"772 Armstrong RD","DATA_SOURCE":"TEST","ENTITY_TYPE":"TEST","DSRC_ACTION":"A","RECORD_ID":"111"}}`
*/
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

/*
The GetRecord_V2 method returns a JSON document of a single record from the Senzing repository.
It extends GetRecord() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"DATA_SOURCE":"TEST","RECORD_ID":"111"}`
*/
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

/*
The GetRedoRecord method returns the next internally queued maintenance record from the Senzing repository.
Usually, the ProcessRedoRecord() or ProcessRedoRecordWithInfo() method is called to process the maintenance record
retrieved by GetRedoRecord().

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document.
*/
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

/*
The  method returns ordWithInfo() method retrieves the last modified time of the Senzing repository,
measured in the number of seconds between the last modified time and January 1, 1970 12:00am GMT (epoch time).

Input
  - ctx: A context to control lifecycle.

Output
  - A Unix TimestampA JSON document.
*/
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

/*
The GetVirtualEntityByRecordID method FIXME:
To control output, use GetVirtualEntityByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - recordList: A JSON document.
    Example: `{"RECORDS": [{"DATA_SOURCE": "TEST","RECORD_ID": "111"},{"DATA_SOURCE": "TEST","RECORD_ID": "222"}]}`

Output
  - A JSON document.
    Example: `{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":...`
*/
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

/*
The GetVirtualEntityByRecordID_V2 method FIXME:
It extends GetVirtualEntityByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - recordList: A JSON document.
    Example: `{"RECORDS": [{"DATA_SOURCE": "TEST","RECORD_ID": "111"},{"DATA_SOURCE": "TEST","RECORD_ID": "222"}]}`
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"RESOLVED_ENTITY":{"ENTITY_ID":1}}`
*/
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

/*
The HowEntityByEntityID method FIXME:
To control output, use HowEntityByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID: The entity ID of the requested data.

Output
  - A JSON document.
    Example: `{"HOW_RESULTS":{"RESOLUTION_STEPS":[],"FINAL_STATE":{"NEED_REEVALUATION":0,"VIRTUAL_ENTITIES":[{"VIRTUAL_ENTITY_ID":"V1","MEMBER_RECORDS":[{"INTERNAL_ID":1,"RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"},{"DATA_SOURCE":"TEST","RECORD_ID":"2D4DABB3FAEAFBD452E9487D06FABC22DC69C846"}]}]}]}}}`
*/
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

/*
The HowEntityByEntityID_V2 method FIXME:
It extends HowEntityByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID: The entity ID of the requested data.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"HOW_RESULTS":{"RESOLUTION_STEPS":[],"FINAL_STATE":{"NEED_REEVALUATION":0,"VIRTUAL_ENTITIES":[{"VIRTUAL_ENTITY_ID":"V1","MEMBER_RECORDS":[{"INTERNAL_ID":1,"RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"},{"DATA_SOURCE":"TEST","RECORD_ID":"2D4DABB3FAEAFBD452E9487D06FABC22DC69C846"}]}]}]}}}`
*/
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

/*
The Init method initializes the Senzing G2 object.
It must be called prior to any other calls.

Input
  - ctx: A context to control lifecycle.
  - moduleName: A name for the auditing node, to help identify it within system logs.
  - iniParams: A JSON string containing configuration paramters.
  - verboseLogging: A flag to enable deeper logging of the G2 processing. 0 for no Senzing logging; 1 for logging.
*/
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

/*
The InitWithConfigID method initializes the Senzing G2 object with a non-default configuration ID.
It must be called prior to any other calls.

Input
  - ctx: A context to control lifecycle.
  - moduleName: A name for the auditing node, to help identify it within system logs.
  - iniParams: A JSON string containing configuration paramters.
  - initConfigID: The configuration ID used for the initialization.
  - verboseLogging: A flag to enable deeper logging of the G2 processing. 0 for no Senzing logging; 1 for logging.
*/
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

/*
The PrimeEngine method pre-initializes some of the heavier weight internal resources of the G2 engine.
The G2 Engine uses "lazy initialization".
PrimeEngine() forces initialization.

Input
  - ctx: A context to control lifecycle.
*/
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

/*
The Process method FIXME:

Input
  - ctx: A context to control lifecycle.
  - record: A JSON document containing the record to be added to the Senzing repository.
*/
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

/*
The ProcessRedoRecord method processes the next redo record and returns it.
Calling ProcessRedoRecord() has the potential to create more redo records in certian situations.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document.
*/
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

/*
The ProcessRedoRecordWithInfo method processes the next redo record and returns it and affected entities.
Calling ProcessRedoRecordWithInfo() has the potential to create more redo records in certian situations.

Input
  - ctx: A context to control lifecycle.
  - flags: FIXME:

Output
  - A JSON document with the record that was re-done.
  - A JSON document with affected entities.
*/
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

/*
The Process method FIXME:

Input
  - ctx: A context to control lifecycle.
  - record: A JSON document containing the record to be added to the Senzing repository.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"DATA_SOURCE":"TEST","RECORD_ID":"555","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}`
*/
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

/*
The ProcessWithResponse method FIXME:

Input
  - ctx: A context to control lifecycle.
  - record: A JSON document containing the record to be added to the Senzing repository.

Output
  - A JSON document.
    Example: `{"MESSAGE": "ER SKIPPED - DUPLICATE RECORD IN G2"}`
*/
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

/*
The ProcessWithResponseResize method FIXME:

Input
  - ctx: A context to control lifecycle.
  - record: A JSON document containing the record to be added to the Senzing repository.

Output
  - A JSON document.
    Example: `{"MESSAGE": "ER SKIPPED - DUPLICATE RECORD IN G2"}`
*/
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

/*
The PurgeRepository method removes every record in the Senzing repository.

Before calling purgeRepository() all other instances of the Senzing API
(whether in custom code, REST API, stream-loader, redoer, G2Loader, etc)
MUST be destroyed or shutdown.

Input
  - ctx: A context to control lifecycle.
*/
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

/*
The ReevaluateEntity method FIXME:

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: FIXME:
*/
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

/*
The ReevaluateEntityWithInfo method FIXME:

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: FIXME:

Output

  - A JSON document.
    Example: `{"DATA_SOURCE":"TEST","RECORD_ID":"111","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}`
*/
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

/*
The ReevaluateRecord method FIXME:

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: FIXME:
*/
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

/*
The ReevaluateRecordWithInfo method FIXME:

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: FIXME:

Output

  - A JSON document.
    Example: `{"DATA_SOURCE":"TEST","RECORD_ID":"111","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}`
*/
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

/*
The Reinit method re-initializes the Senzing G2Engine object using a specified configuration identifier.

Input
  - ctx: A context to control lifecycle.
  - initConfigID: The configuration ID used for the initialization.
*/
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

/*
The ReplaceRecord method updates/replaces a record in the Senzing repository.
If record doesn't exist, a new record is added to the data repository.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
*/
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

/*
The ReplaceRecordWithInfo method updates/replaces a record in the Senzing repository and returns information on the affected entities.
If record doesn't exist, a new record is added to the data repository.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"DATA_SOURCE":"TEST","RECORD_ID":"111","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}`
*/
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

/*
The SearchByAttributes method retrieves entity data based on a user-specified set of entity attributes.
To control output, use SearchByAttributes_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.

Output
  - A JSON document.
    Example: `{"DATA_SOURCE":"TEST","RECORD_ID":"111","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}`
*/
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

/*
The SearchByAttributes_V2 method retrieves entity data based on a user-specified set of entity attributes.
It extends SearchByAttributes() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"RESOLVED_ENTITIES":[{"MATCH_INFO":{"MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","MATCH_KEY":"+NAME+SSN","ERRULE_CODE":"SF1_PNAME_CSTAB"},"ENTITY":{"RESOLVED_ENTITY":{"ENTITY_ID":1}}}]}`
*/
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

/*
The SetLogLevel method sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevel: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
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

/*
The Stats method retrieves workload statistics for the current process.
These statistics will automatically reset after retrieval.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document.
    Example: { "workload": { "loadedRecords":...
*/
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

/*
The WhyEntities method explains why records belong to their resolved entities.
WhyEntities() will compare the record data within an entity
against the rest of the entity data and show why they are connected.
This is calculated based on the features that record data represents.
To control output, use WhyEntities_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"ENTITY_ID":1,"ENTITY_ID_2":2,"MATCH_INFO":{"WHY_KEY":...`
*/
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

/*
The WhyEntities_V2 method explains why records belong to their resolved entities.
WhyEntities_V2() will compare the record data within an entity
against the rest of the entity data and show why they are connected.
This is calculated based on the features that record data represents.
It extends WhyEntities() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"ENTITY_ID":1,"ENTITY_ID_2":2,"MATCH_INFO":{"WHY_KEY":"+PHONE+ACCT_NUM-SSN","WHY_ERRULE_CODE":"SF1","MATCH_LEVEL_CODE":"POSSIBLY_RELATED"}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}}]}`
*/
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

/*
The WhyEntityByEntityID method explains why records belong to their resolved entities.
To control output, use WhyEntityByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID: The entity ID for the starting entity of the search path.

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":...`
*/
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

/*
The WhyEntityByEntityID_V2 method explains why records belong to their resolved entities.
It extends WhyEntityByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID: The entity ID for the starting entity of the search path.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":...`
*/
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

/*
The WhyEntityByRecordID method explains why records belong to their resolved entities.
To control output, use WhyEntityByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":...`
*/
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

/*
The WhyEntityByRecordID_V2 method explains why records belong to their resolved entities.
It extends WhyEntityByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":...`
*/
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

/*
The WhyRecords method explains why records belong to their resolved entities.
To control output, use WhyRecords_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the data.
  - recordID1: The unique identifier within the records of the same data source.
  - dataSourceCode2: Identifies the provenance of the data.
  - recordID2: The unique identifier within the records of the same data source.

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":100001,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"}],...`
*/
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

/*
The WhyRecords_V2 method explains why records belong to their resolved entities.
It extends WhyRecords() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the data.
  - recordID1: The unique identifier within the records of the same data source.
  - dataSourceCode2: Identifies the provenance of the data.
  - recordID2: The unique identifier within the records of the same data source.
  - flags: FIXME:

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":100001,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"}],"INTERNAL_ID_2":2,"ENTITY_ID_2":2,"FOCUS_RECORDS_2":[{"DATA_SOURCE":"TEST","RECORD_ID":"222"}],"MATCH_INFO":{"WHY_KEY":"+PHONE+ACCT_NUM-DOB-SSN","WHY_ERRULE_CODE":"SF1","MATCH_LEVEL_CODE":"POSSIBLY_RELATED"}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}}]}`
*/
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
