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

type G2engineImpl struct {
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
func (g2engine *G2engineImpl) getByteArrayC(size int) *C.char {
	bytes := C.malloc(C.size_t(size))
	return (*C.char)(bytes)
}

func (g2engine *G2engineImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

func (g2engine *G2engineImpl) getError(ctx context.Context, errorNumber int, details ...interface{}) error {
	lastException, err := g2engine.GetLastException(ctx)
	defer g2engine.ClearLastException(ctx)
	message := lastException
	if err != nil {
		message = err.Error()
	}

	var newDetails []interface{}
	newDetails = append(newDetails, details...)
	newDetails = append(newDetails, errors.New(message))
	logger := g2engine.getLogger(ctx)
	errorMessage, err := logger.Message(errorNumber, newDetails...)
	if err != nil {
		errorMessage = err.Error()
	}

	return errors.New(errorMessage)
}

func (g2engine *G2engineImpl) getLogger(ctx context.Context) messagelogger.MessageLoggerInterface {
	if g2engine.logger == nil {
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
		g2engine.logger, _ = messagelogger.New(messageFormat, messageId, messageLogLevel, messageStatus, messageText, messagelogger.LevelInfo)
	}
	return g2engine.logger
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func (g2engine *G2engineImpl) AddRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error {
	//  _DLEXPORT int G2_addRecord(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID);
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
		err = g2engine.getError(ctx, 1, dataSourceCode, recordID, jsonData, loadID, result)
	}
	return err
}

func (g2engine *G2engineImpl) AddRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_addRecordWithInfo(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))
	stringBuffer := C.GoString(C.G2_addRecordWithInfo_helper(dataSourceCodeForC, recordIDForC, jsonDataForC, loadIDForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2, dataSourceCode, recordID, jsonData, loadID, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) AddRecordWithInfoWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string, flags int64) (string, string, error) {
	//  _DLEXPORT int G2_addRecordWithInfoWithReturnedRecordID(const char* dataSourceCode, const char* jsonData, const char *loadID, const long long flags, char *recordIDBuf, const size_t recordIDBufSize, char **responseBuf, size_t *responseBufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))
	result := C.G2_addRecordWithInfoWithReturnedRecordID_helper(dataSourceCodeForC, jsonDataForC, loadIDForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 3, dataSourceCode, jsonData, loadID, flags, result.returnCode, result)
	}
	return C.GoString(result.withInfo), C.GoString(result.recordID), err
}

func (g2engine *G2engineImpl) AddRecordWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string) (string, error) {
	//  _DLEXPORT int G2_addRecordWithReturnedRecordID(const char* dataSourceCode, const char* jsonData, const char *loadID, char *recordIDBuf, const size_t bufSize);
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
		err = g2engine.getError(ctx, 4, dataSourceCode, jsonData, loadID, result)
	}
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	return string(stringBuffer), err
}

func (g2engine *G2engineImpl) CheckRecord(ctx context.Context, record string, recordQueryList string) (string, error) {
	//  _DLEXPORT int G2_checkRecord(const char *record, const char* recordQueryList, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	recordForC := C.CString(record)
	defer C.free(unsafe.Pointer(recordForC))
	recordQueryListForC := C.CString(recordQueryList)
	defer C.free(unsafe.Pointer(recordQueryListForC))
	stringBuffer := C.GoString(C.G2_checkRecord_helper(recordForC, recordQueryListForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 5, record, recordQueryList, returnCode)
	}
	return stringBuffer, err
}

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2engine *G2engineImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2_clearLastException();
	var err error = nil
	C.G2_clearLastException()
	return err
}

func (g2engine *G2engineImpl) CloseExport(ctx context.Context, responseHandle uintptr) error {
	//  _DLEXPORT int G2_closeExport(ExportHandle responseHandle);
	var err error = nil
	result := C.G2_closeExport_helper(C.uintptr_t(responseHandle))
	if result != 0 {
		err = g2engine.getError(ctx, 6, result)
	}
	return err
}

func (g2engine *G2engineImpl) CountRedoRecords(ctx context.Context) (int64, error) {
	//  _DLEXPORT long long G2_countRedoRecords();
	var err error = nil
	result := C.G2_countRedoRecords()
	return int64(result), err
}

func (g2engine *G2engineImpl) DeleteRecord(ctx context.Context, dataSourceCode string, recordID string, loadID string) error {
	//  _DLEXPORT int G2_deleteRecord(const char* dataSourceCode, const char* recordID, const char* loadID);
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))
	result := C.G2_deleteRecord(dataSourceCodeForC, recordIDForC, loadIDForC)
	if result != 0 {
		err = g2engine.getError(ctx, 7, dataSourceCode, recordID, loadID, result)
	}
	return err
}

func (g2engine *G2engineImpl) DeleteRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, loadID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_deleteRecordWithInfo(const char* dataSourceCode, const char* recordID, const char* loadID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))
	stringBuffer := C.GoString(C.G2_deleteRecordWithInfo_helper(dataSourceCodeForC, recordIDForC, loadIDForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 8, dataSourceCode, recordID, loadID, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) Destroy(ctx context.Context) error {
	//  _DLEXPORT int G2_destroy();
	var err error = nil
	result := C.G2_destroy()
	if result != 0 {
		err = g2engine.getError(ctx, 9, result)
	}
	return err
}

func (g2engine *G2engineImpl) ExportConfigAndConfigID(ctx context.Context) (string, int64, error) {
	//  _DLEXPORT int G2_exportConfigAndConfigID(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize), long long* configID );
	var err error = nil
	result := C.G2_exportConfigAndConfigID_helper()
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 10, result.returnCode, result)
	}
	return C.GoString(result.config), int64(C.longlong(result.configID)), err
}

func (g2engine *G2engineImpl) ExportConfig(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_exportConfig(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2_exportConfig_helper())
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 11, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) ExportCSVEntityReport(ctx context.Context, csvColumnList string, flags int64) (uintptr, error) {
	//  _DLEXPORT int G2_exportCSVEntityReport(const char* csvColumnList, const long long flags, ExportHandle* responseHandle);
	var err error = nil
	csvColumnListForC := C.CString(csvColumnList)
	defer C.free(unsafe.Pointer(csvColumnListForC))
	result := C.G2_exportCSVEntityReport_helper(csvColumnListForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 12, csvColumnList, flags, result.returnCode, result)
	}
	return (uintptr)(result.exportHandle), err
}

func (g2engine *G2engineImpl) ExportJSONEntityReport(ctx context.Context, flags int64) (uintptr, error) {
	//  _DLEXPORT int G2_exportJSONEntityReport(const long long flags, ExportHandle* responseHandle);
	var err error = nil
	result := C.G2_exportJSONEntityReport_helper(C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 13, flags, result.returnCode, result)
	}
	return (uintptr)(result.exportHandle), err
}

func (g2engine *G2engineImpl) FetchNext(ctx context.Context, responseHandle uintptr) (string, error) {
	//  _DLEXPORT int G2_fetchNext(ExportHandle responseHandle, char *responseBuf, const size_t bufSize);
	var err error = nil
	result := C.G2_fetchNext_helper(C.uintptr_t(responseHandle))
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 14, result.returnCode, result)
	}
	//    response = bytes.Trim(response, "\x00")
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindInterestingEntitiesByEntityID(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_findInterestingEntitiesByEntityID(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	stringBuffer := C.GoString(C.G2_findInterestingEntitiesByEntityID_helper(C.longlong(entityID), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 15, entityID, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindInterestingEntitiesByRecordID(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findInterestingEntitiesByRecordID(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_findInterestingEntitiesByRecordID_helper(dataSourceCodeForC, recordIDForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 16, dataSourceCode, recordID, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindNetworkByEntityID(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error) {
	//  _DLEXPORT int G2_findNetworkByEntityID(const char* entityList, const int maxDegree, const int buildOutDegree, const int maxEntities, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	entityListForC := C.CString(entityList)
	defer C.free(unsafe.Pointer(entityListForC))
	stringBuffer := C.GoString(C.G2_findNetworkByEntityID_helper(entityListForC, C.int(maxDegree), C.int(buildOutDegree), C.int(maxEntities)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 17, entityList, maxDegree, buildOutDegree, maxEntities, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindNetworkByEntityID_V2(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findNetworkByEntityID_V2(const char* entityList, const int maxDegree, const int buildOutDegree, const int maxEntities, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	entityListForC := C.CString(entityList)
	defer C.free(unsafe.Pointer(entityListForC))
	stringBuffer := C.GoString(C.G2_findNetworkByEntityID_V2_helper(entityListForC, C.int(maxDegree), C.int(buildOutDegree), C.int(maxEntities), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 18, entityList, maxDegree, buildOutDegree, maxEntities, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindNetworkByRecordID(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error) {
	//  _DLEXPORT int G2_findNetworkByRecordID(const char* recordList, const int maxDegree, const int buildOutDegree, const int maxEntities, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	recordListForC := C.CString(recordList)
	defer C.free(unsafe.Pointer(recordListForC))
	stringBuffer := C.GoString(C.G2_findNetworkByRecordID_helper(recordListForC, C.int(maxDegree), C.int(buildOutDegree), C.int(maxEntities)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 19, recordList, maxDegree, buildOutDegree, maxEntities, recordList, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindNetworkByRecordID_V2(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findNetworkByRecordID_V2(const char* recordList, const int maxDegree, const int buildOutDegree, const int maxEntities, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	recordListForC := C.CString(recordList)
	defer C.free(unsafe.Pointer(recordListForC))
	stringBuffer := C.GoString(C.G2_findNetworkByRecordID_V2_helper(recordListForC, C.int(maxDegree), C.int(buildOutDegree), C.int(maxEntities), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 20, recordList, maxDegree, buildOutDegree, maxEntities, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int) (string, error) {
	//  _DLEXPORT int G2_findPathByEntityID(const long long entityID1, const long long entityID2, const int maxDegree, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	stringBuffer := C.GoString(C.G2_findPathByEntityID_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 21, entityID1, entityID2, maxDegree, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathByEntityID_V2(const long long entityID1, const long long entityID2, const int maxDegree, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	stringBuffer := C.GoString(C.G2_findPathByEntityID_V2_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 22, entityID1, entityID2, maxDegree, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int) (string, error) {
	//  _DLEXPORT int G2_findPathByRecordID(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSource1CodeForC := C.CString(dataSourceCode1)
	defer C.free(unsafe.Pointer(dataSource1CodeForC))
	recordID1ForC := C.CString(recordID1)
	defer C.free(unsafe.Pointer(recordID1ForC))
	dataSource2CodeForC := C.CString(dataSourceCode2)
	defer C.free(unsafe.Pointer(dataSource2CodeForC))
	recordID2ForC := C.CString(recordID2)
	defer C.free(unsafe.Pointer(recordID2ForC))
	stringBuffer := C.GoString(C.G2_findPathByRecordID_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC, C.int(maxDegree)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 23, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathByRecordID_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSource1CodeForC := C.CString(dataSourceCode1)
	defer C.free(unsafe.Pointer(dataSource1CodeForC))
	recordID1ForC := C.CString(recordID1)
	defer C.free(unsafe.Pointer(recordID1ForC))
	dataSource2CodeForC := C.CString(dataSourceCode2)
	defer C.free(unsafe.Pointer(dataSource2CodeForC))
	recordID2ForC := C.CString(recordID2)
	defer C.free(unsafe.Pointer(recordID2ForC))
	stringBuffer := C.GoString(C.G2_findPathByRecordID_V2_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC, C.int(maxDegree), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 24, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathExcludingByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByEntityID(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	excludedEntitiesForC := C.CString(excludedEntities)
	defer C.free(unsafe.Pointer(excludedEntitiesForC))
	stringBuffer := C.GoString(C.G2_findPathExcludingByEntityID_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree), excludedEntitiesForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 25, entityID1, entityID2, maxDegree, excludedEntities, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathExcludingByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByEntityID_V2(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	excludedEntitiesForC := C.CString(excludedEntities)
	defer C.free(unsafe.Pointer(excludedEntitiesForC))
	stringBuffer := C.GoString(C.G2_findPathExcludingByEntityID_V2_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree), excludedEntitiesForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 26, entityID1, entityID2, maxDegree, excludedEntities, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathExcludingByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByRecordID(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
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
	stringBuffer := C.GoString(C.G2_findPathExcludingByRecordID_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC, C.int(maxDegree), excludedRecordsForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 27, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathExcludingByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByRecordID_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
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
	stringBuffer := C.GoString(C.G2_findPathExcludingByRecordID_V2_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC, C.int(maxDegree), excludedRecordsForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 28, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, requiredDsrcs string) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByEntityID(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const char* requiredDsrcs, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	excludedEntitiesForC := C.CString(excludedEntities)
	defer C.free(unsafe.Pointer(excludedEntitiesForC))
	requiredDsrcsForC := C.CString(requiredDsrcs)
	defer C.free(unsafe.Pointer(requiredDsrcsForC))
	stringBuffer := C.GoString(C.G2_findPathIncludingSourceByEntityID_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree), excludedEntitiesForC, requiredDsrcsForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 29, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, requiredDsrcs string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByEntityID_V2(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const char* requiredDsrcs, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	excludedEntitiesForC := C.CString(excludedEntities)
	defer C.free(unsafe.Pointer(excludedEntitiesForC))
	requiredDsrcsForC := C.CString(requiredDsrcs)
	defer C.free(unsafe.Pointer(requiredDsrcsForC))
	stringBuffer := C.GoString(C.G2_findPathIncludingSourceByEntityID_V2_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree), excludedEntitiesForC, requiredDsrcsForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 30, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByRecordID(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const char* requiredDsrcs, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
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
	stringBuffer := C.GoString(C.G2_findPathIncludingSourceByRecordID_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC, C.int(maxDegree), excludedRecordsForC, requiredDsrcsForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 31, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByRecordID_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const char* requiredDsrcs, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
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
	stringBuffer := C.GoString(C.G2_findPathIncludingSourceByRecordID_V2_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC, C.int(maxDegree), excludedRecordsForC, requiredDsrcsForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 32, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) GetActiveConfigID(ctx context.Context) (int64, error) {
	//  _DLEXPORT int G2_getActiveConfigID(long long* configID);
	var err error = nil
	result := C.G2_getActiveConfigID_helper()
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 33, result.returnCode, result)
	}
	return int64(C.longlong(result.configID)), err
}

func (g2engine *G2engineImpl) GetEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2_getEntityByEntityID(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	stringBuffer := C.GoString(C.G2_getEntityByEntityID_helper(C.longlong(entityID)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 34, entityID, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) GetEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_getEntityByEntityID_V2(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	stringBuffer := C.GoString(C.G2_getEntityByEntityID_V2_helper(C.longlong(entityID), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 35, entityID, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) GetEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	//  _DLEXPORT int G2_getEntityByRecordID(const char* dataSourceCode, const char* recordID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_getEntityByRecordID_helper(dataSourceCodeForC, recordIDForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 36, dataSourceCode, recordID, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) GetEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_getEntityByRecordID_V2(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_getEntityByRecordID_V2_helper(dataSourceCodeForC, recordIDForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 37, dataSourceCode, recordID, flags, returnCode)
	}
	return stringBuffer, err
}

// GetLastException returns the last exception encountered in the Senzing Engine.
func (g2engine *G2engineImpl) GetLastException(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_getLastException(char *buffer, const size_t bufSize);
	var err error = nil
	stringBuffer := g2engine.getByteArray(initialByteArraySize)
	C.G2_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if len(stringBuffer) == 0 {
		logger := g2engine.getLogger(ctx)
		err = logger.Error(2999)
	}
	return string(stringBuffer), err
}

func (g2engine *G2engineImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2_getLastExceptionCode();
	var err error = nil
	result := C.G2_getLastExceptionCode()
	return int(result), err
}

func (g2engine *G2engineImpl) GetRecord(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	//  _DLEXPORT int G2_getRecord(const char* dataSourceCode, const char* recordID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_getRecord_helper(dataSourceCodeForC, recordIDForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 38, dataSourceCode, recordID, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) GetRecord_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_getRecord_V2(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_getRecord_V2_helper(dataSourceCodeForC, recordIDForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 39, dataSourceCode, recordID, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) GetRedoRecord(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_getRedoRecord(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	result := C.G2_getRedoRecord_helper()
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 40, result.returnCode, result)
	}
	return C.GoString(result.response), err

}

func (g2engine *G2engineImpl) GetRepositoryLastModifiedTime(ctx context.Context) (int64, error) {
	//  _DLEXPORT int G2_getRepositoryLastModifiedTime(long long* lastModifiedTime);
	var err error = nil
	result := C.G2_getRepositoryLastModifiedTime_helper()
	// FIXME:
	// if result.returnCode != 0 {
	// 	err = g2engine.getError(ctx, 41, result.returnCode, result)
	// }
	return int64(result), err
}

func (g2engine *G2engineImpl) GetVirtualEntityByRecordID(ctx context.Context, recordList string) (string, error) {
	//  _DLEXPORT int G2_getVirtualEntityByRecordID(const char* recordList, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	recordListForC := C.CString(recordList)
	defer C.free(unsafe.Pointer(recordListForC))
	stringBuffer := C.GoString(C.G2_getVirtualEntityByRecordID_helper(recordListForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 42, recordList, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) GetVirtualEntityByRecordID_V2(ctx context.Context, recordList string, flags int64) (string, error) {
	//  _DLEXPORT int G2_getVirtualEntityByRecordID_V2(const char* recordList, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	recordListForC := C.CString(recordList)
	defer C.free(unsafe.Pointer(recordListForC))
	stringBuffer := C.GoString(C.G2_getVirtualEntityByRecordID_V2_helper(recordListForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 43, recordList, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) HowEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2_howEntityByEntityID(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	stringBuffer := C.GoString(C.G2_howEntityByEntityID_helper(C.longlong(entityID)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 44, entityID, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) HowEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_howEntityByEntityID_V2(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	stringBuffer := C.GoString(C.G2_howEntityByEntityID_V2_helper(C.longlong(entityID), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 45, entityID, flags, returnCode)
	}
	return stringBuffer, err
}

// Init initializes the Senzing G2diagnosis.
func (g2engine *G2engineImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))
	if result != 0 {
		err = g2engine.getError(ctx, 46, moduleName, iniParams, verboseLogging, result)
	}
	return err
}

func (g2engine *G2engineImpl) InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error {
	//  _DLEXPORT int G2_initWithConfigID(const char *moduleName, const char *iniParams, const long long initConfigID, const int verboseLogging);
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2_initWithConfigID(moduleNameForC, iniParamsForC, C.longlong(initConfigID), C.int(verboseLogging))
	if result != 0 {
		err = g2engine.getError(ctx, 47, moduleName, iniParams, initConfigID, verboseLogging, result)
	}
	return err
}

func (g2engine *G2engineImpl) PrimeEngine(ctx context.Context) error {
	//  _DLEXPORT int G2_primeEngine();
	var err error = nil
	result := C.G2_primeEngine()
	if result != 0 {
		err = g2engine.getError(ctx, 48, result)
	}
	return err
}

func (g2engine *G2engineImpl) Process(ctx context.Context, record string) error {
	//  _DLEXPORT int G2_process(const char *record);
	var err error = nil
	recordForC := C.CString(record)
	defer C.free(unsafe.Pointer(recordForC))
	result := C.G2_process(recordForC)
	if result != 0 {
		err = g2engine.getError(ctx, 49, record, result)
	}
	return err
}

func (g2engine *G2engineImpl) ProcessRedoRecord(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_processRedoRecord(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	result := C.G2_processRedoRecord_helper()
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 50, result.returnCode, result)
	}
	return C.GoString(result.response), err

}

func (g2engine *G2engineImpl) ProcessRedoRecordWithInfo(ctx context.Context, flags int64) (string, string, error) {
	//  _DLEXPORT int G2_processRedoRecordWithInfo(const long long flags, char **responseBuf, size_t *bufSize, char **infoBuf, size_t *infoBufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	result := C.G2_processRedoRecordWithInfo_helper(C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 51, flags, result.returnCode, result)
	}
	return C.GoString(result.response), C.GoString(result.withInfo), err
}

func (g2engine *G2engineImpl) ProcessWithInfo(ctx context.Context, record string, flags int64) (string, error) {
	//  _DLEXPORT int G2_processWithInfo(const char *record, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	recordForC := C.CString(record)
	defer C.free(unsafe.Pointer(recordForC))
	stringBuffer := C.GoString(C.G2_processWithInfo_helper(recordForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 52, record, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) ProcessWithResponse(ctx context.Context, record string) (string, error) {
	//  _DLEXPORT int G2_processWithResponse(const char *record, char *responseBuf, const size_t bufSize);
	var err error = nil
	recordForC := C.CString(record)
	defer C.free(unsafe.Pointer(recordForC))
	stringBuffer := C.GoString(C.G2_processWithResponse_helper(recordForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 53, record, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) ProcessWithResponseResize(ctx context.Context, record string) (string, error) {
	//  _DLEXPORT int G2_processWithResponseResize(const char *record, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	recordForC := C.CString(record)
	defer C.free(unsafe.Pointer(recordForC))
	stringBuffer := C.GoString(C.G2_processWithResponseResize_helper(recordForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 54, record, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) PurgeRepository(ctx context.Context) error {
	//  _DLEXPORT int G2_purgeRepository();
	var err error = nil
	result := C.G2_purgeRepository()
	if result != 0 {
		err = g2engine.getError(ctx, 55, result)
	}
	return err
}

func (g2engine *G2engineImpl) ReevaluateEntity(ctx context.Context, entityID int64, flags int64) error {
	//  _DLEXPORT int G2_reevaluateEntity(const long long entityID, const long long flags);
	var err error = nil
	result := C.G2_reevaluateEntity(C.longlong(entityID), C.longlong(flags))
	if result != 0 {
		err = g2engine.getError(ctx, 56, entityID, flags, result)
	}
	return err
}

func (g2engine *G2engineImpl) ReevaluateEntityWithInfo(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_reevaluateEntityWithInfo(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	stringBuffer := C.GoString(C.G2_reevaluateEntityWithInfo_helper(C.longlong(entityID), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 57, entityID, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) ReevaluateRecord(ctx context.Context, dataSourceCode string, recordID string, flags int64) error {
	//  _DLEXPORT int G2_reevaluateRecord(const char* dataSourceCode, const char* recordID, const long long flags);
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	result := C.G2_reevaluateRecord(dataSourceCodeForC, recordIDForC, C.longlong(flags))
	if result != 0 {
		err = g2engine.getError(ctx, 58, dataSourceCode, recordID, flags, result)
	}
	return err
}

func (g2engine *G2engineImpl) ReevaluateRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_reevaluateRecordWithInfo(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_reevaluateRecordWithInfo_helper(dataSourceCodeForC, recordIDForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 59, dataSourceCode, recordID, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) Reinit(ctx context.Context, initConfigID int64) error {
	//  _DLEXPORT int G2_reinit(const long long initConfigID);
	var err error = nil
	result := C.G2_reinit(C.longlong(initConfigID))
	if result != 0 {
		err = g2engine.getError(ctx, 60, initConfigID, result)
	}
	return err
}

func (g2engine *G2engineImpl) ReplaceRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error {
	//  _DLEXPORT int G2_replaceRecord(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID);
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
		err = g2engine.getError(ctx, 61, dataSourceCode, recordID, jsonData, loadID, result)
	}
	return err
}

func (g2engine *G2engineImpl) ReplaceRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_replaceRecordWithInfo(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))
	stringBuffer := C.GoString(C.G2_replaceRecordWithInfo_helper(dataSourceCodeForC, recordIDForC, jsonDataForC, loadIDForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 62, dataSourceCode, recordID, jsonData, loadID, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) SearchByAttributes(ctx context.Context, jsonData string) (string, error) {
	//  _DLEXPORT int G2_searchByAttributes(const char* jsonData, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	stringBuffer := C.GoString(C.G2_searchByAttributes_helper(jsonDataForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 63, jsonData, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) SearchByAttributes_V2(ctx context.Context, jsonData string, flags int64) (string, error) {
	//  _DLEXPORT int G2_searchByAttributes_V2(const char* jsonData, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	stringBuffer := C.GoString(C.G2_searchByAttributes_V2_helper(jsonDataForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 64, jsonData, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) Stats(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_stats(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2_stats_helper())
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 65, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyEntities(ctx context.Context, entityID1 int64, entityID2 int64) (string, error) {
	//  _DLEXPORT int G2_whyEntities(const long long entityID1, const long long entityID2, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	stringBuffer := C.GoString(C.G2_whyEntities_helper(C.longlong(entityID1), C.longlong(entityID2)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 66, entityID1, entityID2, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyEntities_V2(ctx context.Context, entityID1 int64, entityID2 int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyEntities_V2(const long long entityID1, const long long entityID2, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	stringBuffer := C.GoString(C.G2_whyEntities_V2_helper(C.longlong(entityID1), C.longlong(entityID2), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 67, entityID1, entityID2, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2_whyEntityByEntityID(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	stringBuffer := C.GoString(C.G2_whyEntityByEntityID_helper(C.longlong(entityID)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 68, entityID, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyEntityByEntityID_V2(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	stringBuffer := C.GoString(C.G2_whyEntityByEntityID_V2_helper(C.longlong(entityID), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 69, entityID, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	//  _DLEXPORT int G2_whyEntityByRecordID(const char* dataSourceCode, const char* recordID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_whyEntityByRecordID_helper(dataSourceCodeForC, recordIDForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 70, dataSourceCode, recordID, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyEntityByRecordID_V2(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_whyEntityByRecordID_V2_helper(dataSourceCodeForC, recordIDForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 71, dataSourceCode, recordID, flags, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyRecords(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string) (string, error) {
	//  _DLEXPORT int G2_whyRecords(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSource1CodeForC := C.CString(dataSourceCode1)
	defer C.free(unsafe.Pointer(dataSource1CodeForC))
	recordID1ForC := C.CString(recordID1)
	defer C.free(unsafe.Pointer(recordID1ForC))
	dataSource2CodeForC := C.CString(dataSourceCode2)
	defer C.free(unsafe.Pointer(dataSource2CodeForC))
	recordID2ForC := C.CString(recordID2)
	defer C.free(unsafe.Pointer(recordID2ForC))
	stringBuffer := C.GoString(C.G2_whyRecords_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 72, dataSourceCode1, recordID1, dataSourceCode2, recordID2, returnCode)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyRecords_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyRecords_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSource1CodeForC := C.CString(dataSourceCode1)
	defer C.free(unsafe.Pointer(dataSource1CodeForC))
	recordID1ForC := C.CString(recordID1)
	defer C.free(unsafe.Pointer(recordID1ForC))
	dataSource2CodeForC := C.CString(dataSourceCode2)
	defer C.free(unsafe.Pointer(dataSource2CodeForC))
	recordID2ForC := C.CString(recordID2)
	defer C.free(unsafe.Pointer(recordID2ForC))
	stringBuffer := C.GoString(C.G2_whyRecords_V2_helper(dataSource1CodeForC, recordID1ForC, dataSource2CodeForC, recordID2ForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 73, dataSourceCode1, recordID1, dataSourceCode2, recordID2, flags, returnCode)
	}
	return stringBuffer, err
}
