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
	messageGenerator := g2engine.getMessageGenerator(ctx)
	errorMessage, err := messageGenerator.Message(errorNumber, newDetails...)
	if err != nil {
		errorMessage = err.Error()
	}

	return errors.New(errorMessage)
}

func (g2engine *G2engineImpl) getLogger() messagelogger.MessageLoggerInterface {
	if g2engine.logger == nil {
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
		g2engine.logger, _ = messagelogger.New(messageFormat, messageId, messageLogLevel, messageStatus, messageText, messagelogger.LevelInfo)
	}
	return g2engine.logger
}

func (g2engine *G2engineImpl) getMessageGenerator(ctx context.Context) messagelogger.MessageLoggerInterface {
	if g2engine.messageGenerator == nil {
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
		g2engine.messageGenerator, _ = messagelogger.New(messageFormat, messageId, messageLogLevel, messageStatus, messageText, messagelogger.LevelInfo)
	}
	return g2engine.messageGenerator
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
		g2engine.traceEntry(4001, dataSourceCode, recordID, jsonData, loadID)
	}
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
		err = g2engine.getError(ctx, 2001, dataSourceCode, recordID, jsonData, loadID, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4002, dataSourceCode, recordID, jsonData, loadID, err)
	}
	return err
}

func (g2engine *G2engineImpl) AddRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_addRecordWithInfo(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4003, dataSourceCode, recordID, jsonData, loadID, flags)
	}
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
		err = g2engine.getError(ctx, 2002, dataSourceCode, recordID, jsonData, loadID, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4004, dataSourceCode, recordID, jsonData, loadID, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) AddRecordWithInfoWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string, flags int64) (string, string, error) {
	//  _DLEXPORT int G2_addRecordWithInfoWithReturnedRecordID(const char* dataSourceCode, const char* jsonData, const char *loadID, const long long flags, char *recordIDBuf, const size_t recordIDBufSize, char **responseBuf, size_t *responseBufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4005, dataSourceCode, jsonData, loadID, flags)
	}
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))
	result := C.G2_addRecordWithInfoWithReturnedRecordID_helper(dataSourceCodeForC, jsonDataForC, loadIDForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 2003, dataSourceCode, jsonData, loadID, flags, result.returnCode, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4006, dataSourceCode, jsonData, loadID, flags, C.GoString(result.withInfo), C.GoString(result.recordID), err)
	}
	return C.GoString(result.withInfo), C.GoString(result.recordID), err
}

func (g2engine *G2engineImpl) AddRecordWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string) (string, error) {
	//  _DLEXPORT int G2_addRecordWithReturnedRecordID(const char* dataSourceCode, const char* jsonData, const char *loadID, char *recordIDBuf, const size_t bufSize);
	if g2engine.isTrace {
		g2engine.traceEntry(4007, dataSourceCode, jsonData, loadID)
	}
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
		err = g2engine.getError(ctx, 2004, dataSourceCode, jsonData, loadID, result)
	}
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if g2engine.isTrace {
		defer g2engine.traceExit(4008, dataSourceCode, jsonData, loadID, string(stringBuffer), err)
	}
	return string(stringBuffer), err
}

func (g2engine *G2engineImpl) CheckRecord(ctx context.Context, record string, recordQueryList string) (string, error) {
	//  _DLEXPORT int G2_checkRecord(const char *record, const char* recordQueryList, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2engine.isTrace {
		g2engine.traceEntry(4009, record, recordQueryList)
	}
	var err error = nil
	recordForC := C.CString(record)
	defer C.free(unsafe.Pointer(recordForC))
	recordQueryListForC := C.CString(recordQueryList)
	defer C.free(unsafe.Pointer(recordQueryListForC))
	stringBuffer := C.GoString(C.G2_checkRecord_helper(recordForC, recordQueryListForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2005, record, recordQueryList, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4010, record, recordQueryList, stringBuffer, err)
	}
	return stringBuffer, err
}

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2engine *G2engineImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2_clearLastException();
	if g2engine.isTrace {
		g2engine.traceEntry(4011)
	}
	var err error = nil
	C.G2_clearLastException()
	if g2engine.isTrace {
		defer g2engine.traceExit(4012, err)
	}
	return err
}

func (g2engine *G2engineImpl) CloseExport(ctx context.Context, responseHandle uintptr) error {
	//  _DLEXPORT int G2_closeExport(ExportHandle responseHandle);
	if g2engine.isTrace {
		g2engine.traceEntry(4013, responseHandle)
	}
	var err error = nil
	result := C.G2_closeExport_helper(C.uintptr_t(responseHandle))
	if result != 0 {
		err = g2engine.getError(ctx, 2006, responseHandle, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4014, responseHandle, err)
	}
	return err
}

func (g2engine *G2engineImpl) CountRedoRecords(ctx context.Context) (int64, error) {
	//  _DLEXPORT long long G2_countRedoRecords();
	if g2engine.isTrace {
		g2engine.traceEntry(4015)
	}
	var err error = nil
	result := C.G2_countRedoRecords()
	if g2engine.isTrace {
		defer g2engine.traceExit(4016, int64(result), err)
	}
	return int64(result), err
}

func (g2engine *G2engineImpl) DeleteRecord(ctx context.Context, dataSourceCode string, recordID string, loadID string) error {
	//  _DLEXPORT int G2_deleteRecord(const char* dataSourceCode, const char* recordID, const char* loadID);
	if g2engine.isTrace {
		g2engine.traceEntry(4017, dataSourceCode, recordID, loadID)
	}
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))
	result := C.G2_deleteRecord(dataSourceCodeForC, recordIDForC, loadIDForC)
	if result != 0 {
		err = g2engine.getError(ctx, 2007, dataSourceCode, recordID, loadID, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4018, dataSourceCode, recordID, loadID, err)
	}
	return err
}

func (g2engine *G2engineImpl) DeleteRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, loadID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_deleteRecordWithInfo(const char* dataSourceCode, const char* recordID, const char* loadID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4019, dataSourceCode, recordID, loadID, flags)
	}
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
		err = g2engine.getError(ctx, 2008, dataSourceCode, recordID, loadID, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4020, dataSourceCode, recordID, loadID, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) Destroy(ctx context.Context) error {
	//  _DLEXPORT int G2_destroy();
	if g2engine.isTrace {
		g2engine.traceEntry(4021)
	}
	var err error = nil
	result := C.G2_destroy()
	if result != 0 {
		err = g2engine.getError(ctx, 2009, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4022, err)
	}
	return err
}

func (g2engine *G2engineImpl) ExportConfigAndConfigID(ctx context.Context) (string, int64, error) {
	//  _DLEXPORT int G2_exportConfigAndConfigID(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize), long long* configID );
	if g2engine.isTrace {
		g2engine.traceEntry(4023)
	}
	var err error = nil
	result := C.G2_exportConfigAndConfigID_helper()
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 2010, result.returnCode, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4024, C.GoString(result.config), int64(C.longlong(result.configID)), err)
	}
	return C.GoString(result.config), int64(C.longlong(result.configID)), err
}

func (g2engine *G2engineImpl) ExportConfig(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_exportConfig(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2engine.isTrace {
		g2engine.traceEntry(4025)
	}
	var err error = nil
	stringBuffer := C.GoString(C.G2_exportConfig_helper())
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2011, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4026, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) ExportCSVEntityReport(ctx context.Context, csvColumnList string, flags int64) (uintptr, error) {
	//  _DLEXPORT int G2_exportCSVEntityReport(const char* csvColumnList, const long long flags, ExportHandle* responseHandle);
	if g2engine.isTrace {
		g2engine.traceEntry(4027, csvColumnList, flags)
	}
	var err error = nil
	csvColumnListForC := C.CString(csvColumnList)
	defer C.free(unsafe.Pointer(csvColumnListForC))
	result := C.G2_exportCSVEntityReport_helper(csvColumnListForC, C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 2012, csvColumnList, flags, result.returnCode, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4028, csvColumnList, flags, (uintptr)(result.exportHandle), err)
	}
	return (uintptr)(result.exportHandle), err
}

func (g2engine *G2engineImpl) ExportJSONEntityReport(ctx context.Context, flags int64) (uintptr, error) {
	//  _DLEXPORT int G2_exportJSONEntityReport(const long long flags, ExportHandle* responseHandle);
	if g2engine.isTrace {
		g2engine.traceEntry(4029, flags)
	}
	var err error = nil
	result := C.G2_exportJSONEntityReport_helper(C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 2013, flags, result.returnCode, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4030, flags, (uintptr)(result.exportHandle), err)
	}
	return (uintptr)(result.exportHandle), err
}

func (g2engine *G2engineImpl) FetchNext(ctx context.Context, responseHandle uintptr) (string, error) {
	//  _DLEXPORT int G2_fetchNext(ExportHandle responseHandle, char *responseBuf, const size_t bufSize);
	if g2engine.isTrace {
		g2engine.traceEntry(4031, responseHandle)
	}
	var err error = nil
	result := C.G2_fetchNext_helper(C.uintptr_t(responseHandle))
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 2014, responseHandle, result.returnCode, result)
	}
	//    response = bytes.Trim(response, "\x00")
	if g2engine.isTrace {
		defer g2engine.traceExit(4032, responseHandle, C.GoString(result.response), err)
	}
	return C.GoString(result.response), err
}

func (g2engine *G2engineImpl) FindInterestingEntitiesByEntityID(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_findInterestingEntitiesByEntityID(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4033, entityID, flags)
	}
	var err error = nil
	stringBuffer := C.GoString(C.G2_findInterestingEntitiesByEntityID_helper(C.longlong(entityID), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2015, entityID, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4034, entityID, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindInterestingEntitiesByRecordID(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findInterestingEntitiesByRecordID(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4035, dataSourceCode, recordID, flags)
	}
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_findInterestingEntitiesByRecordID_helper(dataSourceCodeForC, recordIDForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2016, dataSourceCode, recordID, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4036, dataSourceCode, recordID, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindNetworkByEntityID(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error) {
	//  _DLEXPORT int G2_findNetworkByEntityID(const char* entityList, const int maxDegree, const int buildOutDegree, const int maxEntities, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4037, entityList, maxDegree, buildOutDegree, maxDegree)
	}
	var err error = nil
	entityListForC := C.CString(entityList)
	defer C.free(unsafe.Pointer(entityListForC))
	stringBuffer := C.GoString(C.G2_findNetworkByEntityID_helper(entityListForC, C.int(maxDegree), C.int(buildOutDegree), C.int(maxEntities)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2017, entityList, maxDegree, buildOutDegree, maxEntities, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4038, entityList, maxDegree, buildOutDegree, maxDegree, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindNetworkByEntityID_V2(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findNetworkByEntityID_V2(const char* entityList, const int maxDegree, const int buildOutDegree, const int maxEntities, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4039, entityList, maxDegree, buildOutDegree, maxDegree, flags)
	}
	var err error = nil
	entityListForC := C.CString(entityList)
	defer C.free(unsafe.Pointer(entityListForC))
	stringBuffer := C.GoString(C.G2_findNetworkByEntityID_V2_helper(entityListForC, C.int(maxDegree), C.int(buildOutDegree), C.int(maxEntities), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2018, entityList, maxDegree, buildOutDegree, maxEntities, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4040, entityList, maxDegree, buildOutDegree, maxDegree, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindNetworkByRecordID(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error) {
	//  _DLEXPORT int G2_findNetworkByRecordID(const char* recordList, const int maxDegree, const int buildOutDegree, const int maxEntities, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4041, recordList, maxDegree, buildOutDegree, maxDegree)
	}
	var err error = nil
	recordListForC := C.CString(recordList)
	defer C.free(unsafe.Pointer(recordListForC))
	stringBuffer := C.GoString(C.G2_findNetworkByRecordID_helper(recordListForC, C.int(maxDegree), C.int(buildOutDegree), C.int(maxEntities)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2019, recordList, maxDegree, buildOutDegree, maxEntities, recordList, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4042, recordList, maxDegree, buildOutDegree, maxDegree, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindNetworkByRecordID_V2(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findNetworkByRecordID_V2(const char* recordList, const int maxDegree, const int buildOutDegree, const int maxEntities, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4043, recordList, maxDegree, buildOutDegree, maxDegree, flags)
	}
	var err error = nil
	recordListForC := C.CString(recordList)
	defer C.free(unsafe.Pointer(recordListForC))
	stringBuffer := C.GoString(C.G2_findNetworkByRecordID_V2_helper(recordListForC, C.int(maxDegree), C.int(buildOutDegree), C.int(maxEntities), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2020, recordList, maxDegree, buildOutDegree, maxEntities, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4044, recordList, maxDegree, buildOutDegree, maxDegree, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int) (string, error) {
	//  _DLEXPORT int G2_findPathByEntityID(const long long entityID1, const long long entityID2, const int maxDegree, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4045, entityID1, entityID2, maxDegree)
	}
	var err error = nil
	stringBuffer := C.GoString(C.G2_findPathByEntityID_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2021, entityID1, entityID2, maxDegree, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4046, entityID1, entityID2, maxDegree, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathByEntityID_V2(const long long entityID1, const long long entityID2, const int maxDegree, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4047, entityID1, entityID2, maxDegree, flags)
	}
	var err error = nil
	stringBuffer := C.GoString(C.G2_findPathByEntityID_V2_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2022, entityID1, entityID2, maxDegree, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4048, entityID1, entityID2, maxDegree, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int) (string, error) {
	//  _DLEXPORT int G2_findPathByRecordID(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4049, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree)
	}
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
		err = g2engine.getError(ctx, 2023, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4050, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathByRecordID_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4051, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, flags)
	}
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
		err = g2engine.getError(ctx, 2024, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4052, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathExcludingByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByEntityID(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4053, entityID1, entityID2, maxDegree, excludedEntities)
	}
	var err error = nil
	excludedEntitiesForC := C.CString(excludedEntities)
	defer C.free(unsafe.Pointer(excludedEntitiesForC))
	stringBuffer := C.GoString(C.G2_findPathExcludingByEntityID_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree), excludedEntitiesForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2025, entityID1, entityID2, maxDegree, excludedEntities, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4054, entityID1, entityID2, maxDegree, excludedEntities, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathExcludingByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByEntityID_V2(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4055, entityID1, entityID2, maxDegree, excludedEntities, flags)
	}
	var err error = nil
	excludedEntitiesForC := C.CString(excludedEntities)
	defer C.free(unsafe.Pointer(excludedEntitiesForC))
	stringBuffer := C.GoString(C.G2_findPathExcludingByEntityID_V2_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree), excludedEntitiesForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2026, entityID1, entityID2, maxDegree, excludedEntities, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4056, entityID1, entityID2, maxDegree, excludedEntities, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathExcludingByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByRecordID(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4057, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords)
	}
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
		err = g2engine.getError(ctx, 2027, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4058, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathExcludingByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByRecordID_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4059, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, flags)
	}
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
		err = g2engine.getError(ctx, 2028, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4060, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, requiredDsrcs string) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByEntityID(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const char* requiredDsrcs, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4061, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs)
	}
	var err error = nil
	excludedEntitiesForC := C.CString(excludedEntities)
	defer C.free(unsafe.Pointer(excludedEntitiesForC))
	requiredDsrcsForC := C.CString(requiredDsrcs)
	defer C.free(unsafe.Pointer(requiredDsrcsForC))
	stringBuffer := C.GoString(C.G2_findPathIncludingSourceByEntityID_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree), excludedEntitiesForC, requiredDsrcsForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2029, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4062, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, requiredDsrcs string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByEntityID_V2(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const char* requiredDsrcs, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4063, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, flags)
	}
	var err error = nil
	excludedEntitiesForC := C.CString(excludedEntities)
	defer C.free(unsafe.Pointer(excludedEntitiesForC))
	requiredDsrcsForC := C.CString(requiredDsrcs)
	defer C.free(unsafe.Pointer(requiredDsrcsForC))
	stringBuffer := C.GoString(C.G2_findPathIncludingSourceByEntityID_V2_helper(C.longlong(entityID1), C.longlong(entityID2), C.int(maxDegree), excludedEntitiesForC, requiredDsrcsForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2030, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4064, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByRecordID(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const char* requiredDsrcs, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4065, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs)
	}
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
		err = g2engine.getError(ctx, 2031, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4066, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByRecordID_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const char* requiredDsrcs, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4067, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, flags)
	}
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
		err = g2engine.getError(ctx, 2032, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4068, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) GetActiveConfigID(ctx context.Context) (int64, error) {
	//  _DLEXPORT int G2_getActiveConfigID(long long* configID);
	if g2engine.isTrace {
		g2engine.traceEntry(4069)
	}
	var err error = nil
	result := C.G2_getActiveConfigID_helper()
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 2033, result.returnCode, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4070, int64(C.longlong(result.configID)), err)
	}
	return int64(C.longlong(result.configID)), err
}

func (g2engine *G2engineImpl) GetEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2_getEntityByEntityID(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4071, entityID)
	}
	var err error = nil
	stringBuffer := C.GoString(C.G2_getEntityByEntityID_helper(C.longlong(entityID)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2034, entityID, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4072, entityID, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) GetEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_getEntityByEntityID_V2(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4073, entityID, flags)
	}
	var err error = nil
	stringBuffer := C.GoString(C.G2_getEntityByEntityID_V2_helper(C.longlong(entityID), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2035, entityID, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4074, entityID, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) GetEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	//  _DLEXPORT int G2_getEntityByRecordID(const char* dataSourceCode, const char* recordID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4075, dataSourceCode, recordID)
	}
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_getEntityByRecordID_helper(dataSourceCodeForC, recordIDForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2036, dataSourceCode, recordID, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4076, dataSourceCode, recordID, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) GetEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_getEntityByRecordID_V2(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4077, dataSourceCode, recordID, flags)
	}
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_getEntityByRecordID_V2_helper(dataSourceCodeForC, recordIDForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2037, dataSourceCode, recordID, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4078, dataSourceCode, recordID, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

// GetLastException returns the last exception encountered in the Senzing Engine.
func (g2engine *G2engineImpl) GetLastException(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_getLastException(char *buffer, const size_t bufSize);
	if g2engine.isTrace {
		g2engine.traceEntry(4079)
	}
	var err error = nil
	stringBuffer := g2engine.getByteArray(initialByteArraySize)
	C.G2_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if len(stringBuffer) == 0 {
		messageGenerator := g2engine.getMessageGenerator(ctx)
		err = messageGenerator.Error(2999)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4080, string(stringBuffer), err)
	}
	return string(stringBuffer), err
}

func (g2engine *G2engineImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2_getLastExceptionCode();
	if g2engine.isTrace {
		g2engine.traceEntry(4081)
	}
	var err error = nil
	result := int(C.G2_getLastExceptionCode())
	if g2engine.isTrace {
		defer g2engine.traceExit(4082, result, err)
	}
	return result, err
}

func (g2engine *G2engineImpl) GetRecord(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	//  _DLEXPORT int G2_getRecord(const char* dataSourceCode, const char* recordID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4083, dataSourceCode, recordID)
	}
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_getRecord_helper(dataSourceCodeForC, recordIDForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2038, dataSourceCode, recordID, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4084, dataSourceCode, recordID, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) GetRecord_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_getRecord_V2(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4085, dataSourceCode, recordID, flags)
	}
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_getRecord_V2_helper(dataSourceCodeForC, recordIDForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2039, dataSourceCode, recordID, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4086, dataSourceCode, recordID, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) GetRedoRecord(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_getRedoRecord(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2engine.isTrace {
		g2engine.traceEntry(4087)
	}
	var err error = nil
	result := C.G2_getRedoRecord_helper()
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 2040, result.returnCode, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4088, C.GoString(result.response), err)
	}
	return C.GoString(result.response), err

}

func (g2engine *G2engineImpl) GetRepositoryLastModifiedTime(ctx context.Context) (int64, error) {
	//  _DLEXPORT int G2_getRepositoryLastModifiedTime(long long* lastModifiedTime);
	if g2engine.isTrace {
		g2engine.traceEntry(4089)
	}
	var err error = nil
	result := C.G2_getRepositoryLastModifiedTime_helper()
	// FIXME:
	// if result.returnCode != 0 {
	// 	err = g2engine.getError(ctx, 2041, result.returnCode, result)
	// }
	if g2engine.isTrace {
		defer g2engine.traceExit(4090, int64(result), err)
	}
	return int64(result), err
}

func (g2engine *G2engineImpl) GetVirtualEntityByRecordID(ctx context.Context, recordList string) (string, error) {
	//  _DLEXPORT int G2_getVirtualEntityByRecordID(const char* recordList, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4091, recordList)
	}
	var err error = nil
	recordListForC := C.CString(recordList)
	defer C.free(unsafe.Pointer(recordListForC))
	stringBuffer := C.GoString(C.G2_getVirtualEntityByRecordID_helper(recordListForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2042, recordList, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4092, recordList, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) GetVirtualEntityByRecordID_V2(ctx context.Context, recordList string, flags int64) (string, error) {
	//  _DLEXPORT int G2_getVirtualEntityByRecordID_V2(const char* recordList, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4093, recordList, flags)
	}
	var err error = nil
	recordListForC := C.CString(recordList)
	defer C.free(unsafe.Pointer(recordListForC))
	stringBuffer := C.GoString(C.G2_getVirtualEntityByRecordID_V2_helper(recordListForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2043, recordList, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4094, recordList, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) HowEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2_howEntityByEntityID(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4095, entityID)
	}
	var err error = nil
	stringBuffer := C.GoString(C.G2_howEntityByEntityID_helper(C.longlong(entityID)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2044, entityID, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4096, entityID, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) HowEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_howEntityByEntityID_V2(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4097, entityID, flags)
	}
	var err error = nil
	stringBuffer := C.GoString(C.G2_howEntityByEntityID_V2_helper(C.longlong(entityID), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2045, entityID, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4098, entityID, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

// Init initializes the Senzing G2diagnosis.
func (g2engine *G2engineImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	if g2engine.isTrace {
		g2engine.traceEntry(4099, moduleName, iniParams, verboseLogging)
	}
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))
	if result != 0 {
		err = g2engine.getError(ctx, 2046, moduleName, iniParams, verboseLogging, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4100, moduleName, iniParams, verboseLogging, err)
	}
	return err
}

func (g2engine *G2engineImpl) InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error {
	//  _DLEXPORT int G2_initWithConfigID(const char *moduleName, const char *iniParams, const long long initConfigID, const int verboseLogging);
	if g2engine.isTrace {
		g2engine.traceEntry(4101, moduleName, iniParams, initConfigID, verboseLogging)
	}
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2_initWithConfigID(moduleNameForC, iniParamsForC, C.longlong(initConfigID), C.int(verboseLogging))
	if result != 0 {
		err = g2engine.getError(ctx, 2047, moduleName, iniParams, initConfigID, verboseLogging, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4102, moduleName, iniParams, initConfigID, verboseLogging, err)
	}
	return err
}

func (g2engine *G2engineImpl) PrimeEngine(ctx context.Context) error {
	//  _DLEXPORT int G2_primeEngine();
	if g2engine.isTrace {
		g2engine.traceEntry(4103)
	}
	var err error = nil
	result := C.G2_primeEngine()
	if result != 0 {
		err = g2engine.getError(ctx, 2048, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4104, err)
	}
	return err
}

func (g2engine *G2engineImpl) Process(ctx context.Context, record string) error {
	//  _DLEXPORT int G2_process(const char *record);
	if g2engine.isTrace {
		g2engine.traceEntry(4105, record)
	}
	var err error = nil
	recordForC := C.CString(record)
	defer C.free(unsafe.Pointer(recordForC))
	result := C.G2_process(recordForC)
	if result != 0 {
		err = g2engine.getError(ctx, 2049, record, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4106, record, err)
	}
	return err
}

func (g2engine *G2engineImpl) ProcessRedoRecord(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_processRedoRecord(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2engine.isTrace {
		g2engine.traceEntry(4107)
	}
	var err error = nil
	result := C.G2_processRedoRecord_helper()
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 2050, result.returnCode, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4108, C.GoString(result.response), err)
	}
	return C.GoString(result.response), err

}

func (g2engine *G2engineImpl) ProcessRedoRecordWithInfo(ctx context.Context, flags int64) (string, string, error) {
	//  _DLEXPORT int G2_processRedoRecordWithInfo(const long long flags, char **responseBuf, size_t *bufSize, char **infoBuf, size_t *infoBufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4109, flags)
	}
	var err error = nil
	result := C.G2_processRedoRecordWithInfo_helper(C.longlong(flags))
	if result.returnCode != 0 {
		err = g2engine.getError(ctx, 2051, flags, result.returnCode, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4110, flags, C.GoString(result.response), C.GoString(result.withInfo), err)
	}
	return C.GoString(result.response), C.GoString(result.withInfo), err
}

func (g2engine *G2engineImpl) ProcessWithInfo(ctx context.Context, record string, flags int64) (string, error) {
	//  _DLEXPORT int G2_processWithInfo(const char *record, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4111, record, flags)
	}
	var err error = nil
	recordForC := C.CString(record)
	defer C.free(unsafe.Pointer(recordForC))
	stringBuffer := C.GoString(C.G2_processWithInfo_helper(recordForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2052, record, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4112, record, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) ProcessWithResponse(ctx context.Context, record string) (string, error) {
	//  _DLEXPORT int G2_processWithResponse(const char *record, char *responseBuf, const size_t bufSize);
	if g2engine.isTrace {
		g2engine.traceEntry(4113, record)
	}
	var err error = nil
	recordForC := C.CString(record)
	defer C.free(unsafe.Pointer(recordForC))
	stringBuffer := C.GoString(C.G2_processWithResponse_helper(recordForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2053, record, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4114, record, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) ProcessWithResponseResize(ctx context.Context, record string) (string, error) {
	//  _DLEXPORT int G2_processWithResponseResize(const char *record, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2engine.isTrace {
		g2engine.traceEntry(4115, record)
	}
	var err error = nil
	recordForC := C.CString(record)
	defer C.free(unsafe.Pointer(recordForC))
	stringBuffer := C.GoString(C.G2_processWithResponseResize_helper(recordForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2054, record, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4116, record, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) PurgeRepository(ctx context.Context) error {
	//  _DLEXPORT int G2_purgeRepository();
	if g2engine.isTrace {
		g2engine.traceEntry(4117)
	}
	var err error = nil
	result := C.G2_purgeRepository()
	if result != 0 {
		err = g2engine.getError(ctx, 2055, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4118, err)
	}
	return err
}

func (g2engine *G2engineImpl) ReevaluateEntity(ctx context.Context, entityID int64, flags int64) error {
	//  _DLEXPORT int G2_reevaluateEntity(const long long entityID, const long long flags);
	if g2engine.isTrace {
		g2engine.traceEntry(4119, entityID, flags)
	}
	var err error = nil
	result := C.G2_reevaluateEntity(C.longlong(entityID), C.longlong(flags))
	if result != 0 {
		err = g2engine.getError(ctx, 2056, entityID, flags, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4120, entityID, flags, err)
	}
	return err
}

func (g2engine *G2engineImpl) ReevaluateEntityWithInfo(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_reevaluateEntityWithInfo(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4121, entityID, flags)
	}
	var err error = nil
	stringBuffer := C.GoString(C.G2_reevaluateEntityWithInfo_helper(C.longlong(entityID), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2057, entityID, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4122, entityID, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) ReevaluateRecord(ctx context.Context, dataSourceCode string, recordID string, flags int64) error {
	//  _DLEXPORT int G2_reevaluateRecord(const char* dataSourceCode, const char* recordID, const long long flags);
	if g2engine.isTrace {
		g2engine.traceEntry(4123, dataSourceCode, recordID, flags)
	}
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	result := C.G2_reevaluateRecord(dataSourceCodeForC, recordIDForC, C.longlong(flags))
	if result != 0 {
		err = g2engine.getError(ctx, 2058, dataSourceCode, recordID, flags, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4124, dataSourceCode, recordID, flags, err)
	}
	return err
}

func (g2engine *G2engineImpl) ReevaluateRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_reevaluateRecordWithInfo(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4125, dataSourceCode, recordID, flags)
	}
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_reevaluateRecordWithInfo_helper(dataSourceCodeForC, recordIDForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2059, dataSourceCode, recordID, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4126, dataSourceCode, recordID, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) Reinit(ctx context.Context, initConfigID int64) error {
	//  _DLEXPORT int G2_reinit(const long long initConfigID);
	if g2engine.isTrace {
		g2engine.traceEntry(4127, initConfigID)
	}
	var err error = nil
	result := C.G2_reinit(C.longlong(initConfigID))
	if result != 0 {
		err = g2engine.getError(ctx, 2060, initConfigID, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4128, initConfigID, err)
	}
	return err
}

func (g2engine *G2engineImpl) ReplaceRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error {
	//  _DLEXPORT int G2_replaceRecord(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID);
	if g2engine.isTrace {
		g2engine.traceEntry(4129, dataSourceCode, recordID, jsonData, loadID)
	}
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
		err = g2engine.getError(ctx, 2061, dataSourceCode, recordID, jsonData, loadID, result)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4130, dataSourceCode, recordID, jsonData, loadID, err)
	}
	return err
}

func (g2engine *G2engineImpl) ReplaceRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_replaceRecordWithInfo(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4131, dataSourceCode, recordID, jsonData, loadID, flags)
	}
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
		err = g2engine.getError(ctx, 2062, dataSourceCode, recordID, jsonData, loadID, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4132, dataSourceCode, recordID, jsonData, loadID, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) SearchByAttributes(ctx context.Context, jsonData string) (string, error) {
	//  _DLEXPORT int G2_searchByAttributes(const char* jsonData, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4133, jsonData)
	}
	var err error = nil
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	stringBuffer := C.GoString(C.G2_searchByAttributes_helper(jsonDataForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2063, jsonData, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4134, jsonData, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) SearchByAttributes_V2(ctx context.Context, jsonData string, flags int64) (string, error) {
	//  _DLEXPORT int G2_searchByAttributes_V2(const char* jsonData, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4135, jsonData, flags)
	}
	var err error = nil
	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))
	stringBuffer := C.GoString(C.G2_searchByAttributes_V2_helper(jsonDataForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2064, jsonData, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4136, jsonData, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) Stats(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_stats(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	if g2engine.isTrace {
		g2engine.traceEntry(4137)
	}
	var err error = nil
	stringBuffer := C.GoString(C.G2_stats_helper())
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2065, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4138, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyEntities(ctx context.Context, entityID1 int64, entityID2 int64) (string, error) {
	//  _DLEXPORT int G2_whyEntities(const long long entityID1, const long long entityID2, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4139, entityID1, entityID2)
	}
	var err error = nil
	stringBuffer := C.GoString(C.G2_whyEntities_helper(C.longlong(entityID1), C.longlong(entityID2)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2066, entityID1, entityID2, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4140, entityID1, entityID2, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyEntities_V2(ctx context.Context, entityID1 int64, entityID2 int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyEntities_V2(const long long entityID1, const long long entityID2, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4141, entityID1, entityID2, flags)
	}
	var err error = nil
	stringBuffer := C.GoString(C.G2_whyEntities_V2_helper(C.longlong(entityID1), C.longlong(entityID2), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2067, entityID1, entityID2, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4142, entityID1, entityID2, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2_whyEntityByEntityID(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4143, entityID)
	}
	var err error = nil
	stringBuffer := C.GoString(C.G2_whyEntityByEntityID_helper(C.longlong(entityID)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2068, entityID, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4144, entityID, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyEntityByEntityID_V2(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4145, entityID, flags)
	}
	var err error = nil
	stringBuffer := C.GoString(C.G2_whyEntityByEntityID_V2_helper(C.longlong(entityID), C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2069, entityID, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4146, entityID, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	//  _DLEXPORT int G2_whyEntityByRecordID(const char* dataSourceCode, const char* recordID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4147, dataSourceCode, recordID)
	}
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_whyEntityByRecordID_helper(dataSourceCodeForC, recordIDForC))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2070, dataSourceCode, recordID, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4148, dataSourceCode, recordID, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyEntityByRecordID_V2(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4149, dataSourceCode, recordID, flags)
	}
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))
	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))
	stringBuffer := C.GoString(C.G2_whyEntityByRecordID_V2_helper(dataSourceCodeForC, recordIDForC, C.longlong(flags)))
	returnCode := 0 // FIXME:
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2071, dataSourceCode, recordID, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4150, dataSourceCode, recordID, flags, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyRecords(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string) (string, error) {
	//  _DLEXPORT int G2_whyRecords(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4151, dataSourceCode1, recordID1, dataSourceCode2, recordID2)
	}
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
		err = g2engine.getError(ctx, 2072, dataSourceCode1, recordID1, dataSourceCode2, recordID2, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4152, dataSourceCode1, recordID1, dataSourceCode2, recordID2, stringBuffer, err)
	}
	return stringBuffer, err
}

func (g2engine *G2engineImpl) WhyRecords_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyRecords_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	if g2engine.isTrace {
		g2engine.traceEntry(4153, dataSourceCode1, recordID1, dataSourceCode2, recordID2, flags)
	}
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
		err = g2engine.getError(ctx, 2073, dataSourceCode1, recordID1, dataSourceCode2, recordID2, flags, returnCode)
	}
	if g2engine.isTrace {
		defer g2engine.traceExit(4154, dataSourceCode1, recordID1, dataSourceCode2, recordID2, flags, stringBuffer, err)
	}
	return stringBuffer, err
}
