package g2engine

import (
	"context"
	"fmt"
	"log"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing/go-helpers/g2engineconfigurationjson"
	"github.com/stretchr/testify/assert"
)

const (
	defaultTruncation = 76
)

var (
	g2engineSingleton G2engine
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) G2engine {

	if g2engineSingleton == nil {
		g2engineSingleton = &G2engineImpl{}

		// g2engineSingleton.SetLogLevel(ctx, logger.LevelTrace)
		log.SetFlags(0)

		moduleName := "Test module name"
		verboseLogging := 0
		iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			test.Logf("Cannot construct system configuration. Error: %v", jsonErr)
		}

		initErr := g2engineSingleton.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			test.Logf("Cannot Init. Error: %v", initErr)
		}
	}
	return g2engineSingleton
}

func getG2Engine(ctx context.Context) G2engine {
	g2engine := &G2engineImpl{}
	moduleName := "Test module name"
	verboseLogging := 0
	iniParams, _ := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	g2engine.Init(ctx, moduleName, iniParams, verboseLogging)
	return g2engine
}

func truncate(aString string, length int) string {
	return truncator.Truncate(aString, length, "...", truncator.PositionEnd)
}

func printResult(test *testing.T, title string, result interface{}) {
	if 1 == 0 {
		test.Logf("%s: %v", title, truncate(fmt.Sprintf("%v", result), defaultTruncation))
	}
}

func printActual(test *testing.T, actual interface{}) {
	printResult(test, "Actual", actual)
}

func testError(test *testing.T, ctx context.Context, g2engine G2engine, err error) {
	if err != nil {
		lastException, _ := g2engine.GetLastException(ctx)
		test.Log("Error:", err.Error())
		assert.FailNow(test, lastException)
	}
}

func testErrorNoFail(test *testing.T, ctx context.Context, g2engine G2engine, err error) {
	if err != nil {
		lastException, _ := g2engine.GetLastException(ctx)
		test.Log("Error:", err.Error(), "LastException:", lastException)
	}
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestBuildSimpleSystemConfigurationJson(test *testing.T) {
	actual, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, actual)
	}
	printActual(test, actual)
}

func TestGetObject(test *testing.T) {
	ctx := context.TODO()
	getTestObject(ctx, test)
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

// PurgeRepository() is first to start with a clean database.
func TestG2engineImpl_PurgeRepository(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	err := g2engine.PurgeRepository(ctx)
	testError(test, ctx, g2engine, err)
}

func TestG2engineImpl_AddRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "111", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	err := g2engine.AddRecord(ctx, dataSourceCode, recordID, jsonData, loadID)
	testError(test, ctx, g2engine, err)
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	jsonData2 := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "222", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID2 := "TEST"
	err2 := g2engine.AddRecord(ctx, dataSourceCode2, recordID2, jsonData2, loadID2)
	testError(test, ctx, g2engine, err2)
}

func TestG2engineImpl_AddRecordWithInfo(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "333"
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "333", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	var flags int64 = 0
	actual, err := g2engine.AddRecordWithInfo(ctx, dataSourceCode, recordID, jsonData, loadID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_AddRecordWithInfoWithReturnedRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	var flags int64 = 0
	actual, actualRecordID, err := g2engine.AddRecordWithInfoWithReturnedRecordID(ctx, dataSourceCode, jsonData, loadID, flags)
	testError(test, ctx, g2engine, err)
	printResult(test, "Actual RecordID", actualRecordID)
	printActual(test, actual)
}

func TestG2engineImpl_AddRecordWithReturnedRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	jsonData := `{"SOCIAL_HANDLE": "bobby", "DATE_OF_BIRTH": "1/2/1983", "ADDR_STATE": "WI", "ADDR_POSTAL_CODE": "54434", "SSN_NUMBER": "987-65-4321", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "Smith", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	actual, err := g2engine.AddRecordWithReturnedRecordID(ctx, dataSourceCode, jsonData, loadID)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_CheckRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	record := `{"DATA_SOURCE": "TEST", "NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "Smith", "NAME_MIDDLE": "M" }], "PASSPORT_NUMBER": "PP11111", "PASSPORT_COUNTRY": "US", "DRIVERS_LICENSE_NUMBER": "DL11111", "SSN_NUMBER": "111-11-1111"}`
	recordQueryList := `{"RECORDS": [{"DATA_SOURCE": "TEST","RECORD_ID": "111"},{"DATA_SOURCE": "TEST","RECORD_ID": "123456789"}]}`
	actual, err := g2engine.CheckRecord(ctx, record, recordQueryList)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_ClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	err := g2engine.ClearLastException(ctx)
	testError(test, ctx, g2engine, err)
}

// FAIL:
func TestG2engineImpl_ExportJSONEntityReport(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	flags := int64(0)
	aHandle, err := g2engine.ExportJSONEntityReport(ctx, flags)
	testError(test, ctx, g2engine, err)
	anEntity, err := g2engine.FetchNext(ctx, aHandle)
	testError(test, ctx, g2engine, err)
	printResult(test, "Entity", anEntity)
	err = g2engine.CloseExport(ctx, aHandle)
	testError(test, ctx, g2engine, err)
}

func TestG2engineImpl_CountRedoRecords(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.CountRedoRecords(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_ExportConfigAndConfigID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actualConfig, actualConfigId, err := g2engine.ExportConfigAndConfigID(ctx)
	testError(test, ctx, g2engine, err)
	printResult(test, "Actual Config", actualConfig)
	printResult(test, "Actual Config ID", actualConfigId)
}

func TestG2engineImpl_ExportConfig(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.ExportConfig(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

//func TestG2engineImpl_ExportCSVEntityReport(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx, test)
//	csvColumnList := ""
//	var flags int64 = 0
//	actual, err := g2engine.ExportCSVEntityReport(ctx, csvColumnList, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}

func TestG2engineImpl_FindInterestingEntitiesByEntityID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	var flags int64 = 0
	actual, err := g2engine.FindInterestingEntitiesByEntityID(ctx, entityID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindInterestingEntitiesByRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	actual, err := g2engine.FindInterestingEntitiesByRecordID(ctx, dataSourceCode, recordID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindNetworkByEntityID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	entityList := `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}]}`
	maxDegree := 2
	buildOutDegree := 1
	maxEntities := 10
	actual, err := g2engine.FindNetworkByEntityID(ctx, entityList, maxDegree, buildOutDegree, maxEntities)
	testErrorNoFail(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindNetworkByEntityID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	entityList := `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}]}`
	maxDegree := 2
	buildOutDegree := 1
	maxEntities := 10
	var flags int64 = 0
	actual, err := g2engine.FindNetworkByEntityID_V2(ctx, entityList, maxDegree, buildOutDegree, maxEntities, flags)
	testErrorNoFail(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindNetworkByRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	recordList := `{"RECORDS": [{"DATA_SOURCE": "TEST", "RECORD_ID": "111"}, {"DATA_SOURCE": "TEST", "RECORD_ID": "222"}, {"DATA_SOURCE": "TEST", "RECORD_ID": "333"}]}`
	maxDegree := 1
	buildOutDegree := 2
	maxEntities := 10
	actual, err := g2engine.FindNetworkByRecordID(ctx, recordList, maxDegree, buildOutDegree, maxEntities)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindNetworkByRecordID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	recordList := `{"RECORDS": [{"DATA_SOURCE": "TEST", "RECORD_ID": "111"}, {"DATA_SOURCE": "TEST", "RECORD_ID": "222"}, {"DATA_SOURCE": "TEST", "RECORD_ID": "333"}]}`
	maxDegree := 1
	buildOutDegree := 2
	maxEntities := 10
	var flags int64 = 0
	actual, err := g2engine.FindNetworkByRecordID_V2(ctx, recordList, maxDegree, buildOutDegree, maxEntities, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindPathByEntityID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	maxDegree := 1
	actual, err := g2engine.FindPathByEntityID(ctx, entityID1, entityID2, maxDegree)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindPathByEntityID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	maxDegree := 1
	var flags int64 = 0
	actual, err := g2engine.FindPathByEntityID_V2(ctx, entityID1, entityID2, maxDegree, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindPathByRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	maxDegree := 1
	actual, err := g2engine.FindPathByRecordID(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindPathByRecordID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	maxDegree := 1
	var flags int64 = 0
	actual, err := g2engine.FindPathByRecordID_V2(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindPathExcludingByEntityID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	maxDegree := 1
	excludedEntities := `{"ENTITIES": [{"ENTITY_ID": 1}]}`
	actual, err := g2engine.FindPathExcludingByEntityID(ctx, entityID1, entityID2, maxDegree, excludedEntities)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindPathExcludingByEntityID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	maxDegree := 1
	excludedEntities := `{"ENTITIES": [{"ENTITY_ID": 1}]}`
	var flags int64 = 0
	actual, err := g2engine.FindPathExcludingByEntityID_V2(ctx, entityID1, entityID2, maxDegree, excludedEntities, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindPathExcludingByRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	maxDegree := 1
	excludedRecords := `{"RECORDS": [{ "DATA_SOURCE": "TEST", "RECORD_ID": "111"}]}`
	actual, err := g2engine.FindPathExcludingByRecordID(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindPathExcludingByRecordID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	maxDegree := 1
	excludedRecords := `{"RECORDS": [{ "DATA_SOURCE": "TEST", "RECORD_ID": "111"}]}`
	var flags int64 = 0
	actual, err := g2engine.FindPathExcludingByRecordID_V2(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindPathIncludingSourceByEntityID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	maxDegree := 1
	excludedEntities := `{"ENTITIES": [{"ENTITY_ID": 1}]}`
	requiredDsrcs := `{"DATA_SOURCES": ["TEST"]}`
	actual, err := g2engine.FindPathIncludingSourceByEntityID(ctx, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindPathIncludingSourceByEntityID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	maxDegree := 1
	excludedEntities := `{"ENTITIES": [{"ENTITY_ID": 1}]}`
	requiredDsrcs := `{"DATA_SOURCES": ["TEST"]}`
	var flags int64 = 0
	actual, err := g2engine.FindPathIncludingSourceByEntityID_V2(ctx, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindPathIncludingSourceByRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	maxDegree := 1
	excludedEntities := `{"ENTITIES": [{"ENTITY_ID": 1}]}`
	requiredDsrcs := `{"DATA_SOURCES": ["TEST"]}`
	actual, err := g2engine.FindPathIncludingSourceByRecordID(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedEntities, requiredDsrcs)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_FindPathIncludingSourceByRecordID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	maxDegree := 1
	excludedEntities := `{"ENTITIES": [{"ENTITY_ID": 1}]}`
	requiredDsrcs := `{"DATA_SOURCES": ["TEST"]}`
	var flags int64 = 0
	actual, err := g2engine.FindPathIncludingSourceByRecordID_V2(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedEntities, requiredDsrcs, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_GetActiveConfigID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.GetActiveConfigID(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_GetEntityByEntityID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	actual, err := g2engine.GetEntityByEntityID(ctx, entityID)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_GetEntityByEntityID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	var flags int64 = 0
	actual, err := g2engine.GetEntityByEntityID_V2(ctx, entityID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_GetEntityByRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	actual, err := g2engine.GetEntityByRecordID(ctx, dataSourceCode, recordID)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_GetEntityByRecordID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	actual, err := g2engine.GetEntityByRecordID_V2(ctx, dataSourceCode, recordID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_GetLastException(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.GetLastException(ctx)
	testErrorNoFail(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_GetLastExceptionCode(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.GetLastExceptionCode(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_GetRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	actual, err := g2engine.GetRecord(ctx, dataSourceCode, recordID)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_GetRecord_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	actual, err := g2engine.GetRecord_V2(ctx, dataSourceCode, recordID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_GetRedoRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.GetRedoRecord(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_GetRepositoryLastModifiedTime(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.GetRepositoryLastModifiedTime(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_GetVirtualEntityByRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	recordList := `{"RECORDS": [{"DATA_SOURCE": "TEST","RECORD_ID": "111"},{"DATA_SOURCE": "TEST","RECORD_ID": "222"}]}`
	actual, err := g2engine.GetVirtualEntityByRecordID(ctx, recordList)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_GetVirtualEntityByRecordID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	recordList := `{"RECORDS": [{"DATA_SOURCE": "TEST","RECORD_ID": "111"},{"DATA_SOURCE": "TEST","RECORD_ID": "222"}]}`
	var flags int64 = 0
	actual, err := g2engine.GetVirtualEntityByRecordID_V2(ctx, recordList, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_HowEntityByEntityID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	actual, err := g2engine.HowEntityByEntityID(ctx, entityID)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_HowEntityByEntityID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	var flags int64 = 0
	actual, err := g2engine.HowEntityByEntityID_V2(ctx, entityID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_Init(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	testError(test, ctx, g2engine, jsonErr)
	err := g2engine.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2engine, err)
}

func TestG2engineImpl_InitWithConfigID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	moduleName := "Test module name"
	var initConfigID int64 = 1
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	testError(test, ctx, g2engine, jsonErr)
	err := g2engine.InitWithConfigID(ctx, moduleName, iniParams, initConfigID, verboseLogging)
	testError(test, ctx, g2engine, err)
}

func TestG2engineImpl_PrimeEngine(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	err := g2engine.PrimeEngine(ctx)
	testError(test, ctx, g2engine, err)
}

func TestG2engineImpl_Process(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	record := `{"DATA_SOURCE": "TEST", "SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "444", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	err := g2engine.Process(ctx, record)
	testError(test, ctx, g2engine, err)
}

func TestG2engineImpl_ProcessRedoRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.ProcessRedoRecord(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_ProcessRedoRecordWithInfo(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var flags int64 = 0
	actual, actualInfo, err := g2engine.ProcessRedoRecordWithInfo(ctx, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
	printResult(test, "Actual Info", actualInfo)
}

func TestG2engineImpl_ProcessWithInfo(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	record := `{"DATA_SOURCE": "TEST", "SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "555", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	var flags int64 = 0
	actual, err := g2engine.ProcessWithInfo(ctx, record, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_ProcessWithResponse(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	record := `{"DATA_SOURCE": "TEST", "SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "666", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	actual, err := g2engine.ProcessWithResponse(ctx, record)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_ProcessWithResponseResize(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	record := `{"DATA_SOURCE": "TEST", "SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "777", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	actual, err := g2engine.ProcessWithResponseResize(ctx, record)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_ReevaluateEntity(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	var flags int64 = 0
	err := g2engine.ReevaluateEntity(ctx, entityID, flags)
	testError(test, ctx, g2engine, err)
}

func TestG2engineImpl_ReevaluateEntityWithInfo(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	var flags int64 = 0
	actual, err := g2engine.ReevaluateEntityWithInfo(ctx, entityID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_ReevaluateRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	err := g2engine.ReevaluateRecord(ctx, dataSourceCode, recordID, flags)
	testError(test, ctx, g2engine, err)
}

func TestG2engineImpl_ReevaluateRecordWithInfo(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	actual, err := g2engine.ReevaluateRecordWithInfo(ctx, dataSourceCode, recordID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_Reinit(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	initConfigID, err := g2engine.GetActiveConfigID(ctx)
	testError(test, ctx, g2engine, err)
	err = g2engine.Reinit(ctx, initConfigID)
	testError(test, ctx, g2engine, err)
	printActual(test, initConfigID)
}

func TestG2engineImpl_ReplaceRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1984", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "111", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	err := g2engine.ReplaceRecord(ctx, dataSourceCode, recordID, jsonData, loadID)
	testError(test, ctx, g2engine, err)
}

func TestG2engineImpl_ReplaceRecordWithInfo(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1985", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "111", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	var flags int64 = 0
	actual, err := g2engine.ReplaceRecordWithInfo(ctx, dataSourceCode, recordID, jsonData, loadID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_SearchByAttributes(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	jsonData := `{"NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "SEAMAN"}], "SSN_NUMBER": "053-39-3251"}`
	actual, err := g2engine.SearchByAttributes(ctx, jsonData)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_SearchByAttributes_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	jsonData := `{"NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "SEAMAN"}], "SSN_NUMBER": "053-39-3251"}`
	var flags int64 = 0
	actual, err := g2engine.SearchByAttributes_V2(ctx, jsonData, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_Stats(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.Stats(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_WhyEntities(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	actual, err := g2engine.WhyEntities(ctx, entityID1, entityID2)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_WhyEntities_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	var flags int64 = 0
	actual, err := g2engine.WhyEntities_V2(ctx, entityID1, entityID2, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_WhyEntityByEntityID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	actual, err := g2engine.WhyEntityByEntityID(ctx, entityID)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_WhyEntityByEntityID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	var flags int64 = 0
	actual, err := g2engine.WhyEntityByEntityID_V2(ctx, entityID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_WhyEntityByRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	actual, err := g2engine.WhyEntityByRecordID(ctx, dataSourceCode, recordID)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_WhyEntityByRecordID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	actual, err := g2engine.WhyEntityByRecordID_V2(ctx, dataSourceCode, recordID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_WhyRecords(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	actual, err := g2engine.WhyRecords(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_WhyRecords_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	var flags int64 = 0
	actual, err := g2engine.WhyRecords_V2(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_DeleteRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	loadID := "TEST"
	err := g2engine.DeleteRecord(ctx, dataSourceCode, recordID, loadID)
	testError(test, ctx, g2engine, err)
}

func TestG2engineImpl_DeleteRecordWithInfo(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	loadID := "TEST"
	var flags int64 = 0
	actual, err := g2engine.DeleteRecordWithInfo(ctx, dataSourceCode, recordID, loadID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestG2engineImpl_Destroy(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	err := g2engine.Destroy(ctx)
	testError(test, ctx, g2engine, err)
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

// PurgeRepository() is first to start with a clean database.
func ExampleG2engineImpl_PurgeRepository() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	g2engine.PurgeRepository(ctx)
	// Output:
}

func ExampleG2engineImpl_AddRecord() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode := "TEST"
	recordID := "111"
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "111", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	g2engine.AddRecord(ctx, dataSourceCode, recordID, jsonData, loadID)
	// Output:
}

func ExampleG2engineImpl_AddRecord_secondRecord() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode := "TEST"
	recordID := "222"
	jsonData := `{"SOCIAL_HANDLE": "flavorh2", "DATE_OF_BIRTH": "6/9/1983", "ADDR_STATE": "WI", "ADDR_POSTAL_CODE": "53543", "SSN_NUMBER": "153-33-5185", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "222", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "OCEANGUY", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	g2engine.AddRecord(ctx, dataSourceCode, recordID, jsonData, loadID)
	// Output:
}

func ExampleG2engineImpl_AddRecordWithInfo() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode := "TEST"
	recordID := "333"
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "333", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	var flags int64 = 0
	result, _ := g2engine.AddRecordWithInfo(ctx, dataSourceCode, recordID, jsonData, loadID, flags)
	fmt.Println(result)
	// Output: {"DATA_SOURCE":"TEST","RECORD_ID":"333","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}
}

func ExampleG2engineImpl_AddRecordWithInfoWithReturnedRecordID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode := "TEST"
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	var flags int64 = 0
	result, _, _ := g2engine.AddRecordWithInfoWithReturnedRecordID(ctx, dataSourceCode, jsonData, loadID, flags)
	fmt.Println(truncate(result, 37))
	// Output: {"DATA_SOURCE":"TEST","RECORD_ID":...
}

func ExampleG2engineImpl_AddRecordWithReturnedRecordID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode := "TEST"
	jsonData := `{"SOCIAL_HANDLE": "bobby", "DATE_OF_BIRTH": "1/2/1983", "ADDR_STATE": "WI", "ADDR_POSTAL_CODE": "54434", "SSN_NUMBER": "987-65-4321", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "Smith", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	result, _ := g2engine.AddRecordWithReturnedRecordID(ctx, dataSourceCode, jsonData, loadID)
	fmt.Printf("Length of record identifier is %d hexidecimal characters.\n", len(result))
	// Output: Length of record identifier is 40 hexidecimal characters.
}

func ExampleG2engineImpl_CheckRecord() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	record := `{"DATA_SOURCE": "TEST", "NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "Smith", "NAME_MIDDLE": "M" }], "PASSPORT_NUMBER": "PP11111", "PASSPORT_COUNTRY": "US", "DRIVERS_LICENSE_NUMBER": "DL11111", "SSN_NUMBER": "111-11-1111"}`
	recordQueryList := `{"RECORDS": [{"DATA_SOURCE": "TEST","RECORD_ID": "111"},{"DATA_SOURCE": "TEST","RECORD_ID": "123456789"}]}`
	result, _ := g2engine.CheckRecord(ctx, record, recordQueryList)
	fmt.Println(truncate(result, 61))
	// Output: {"CHECK_RECORD_RESPONSE":[{"DSRC_CODE":"TEST","RECORD_ID":...
}

func ExampleG2engineImpl_ClearLastException() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	g2engine.ClearLastException(ctx)
	// Output:
}

func ExampleG2engineImpl_CloseExport() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	flags := int64(0)
	responseHandle, _ := g2engine.ExportJSONEntityReport(ctx, flags)
	g2engine.CloseExport(ctx, responseHandle)
	// Output:
}

func ExampleG2engineImpl_CountRedoRecords() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	result, _ := g2engine.CountRedoRecords(ctx)
	fmt.Println(result)
	// Output: 0
}

func ExampleG2engineImpl_DeleteRecord() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode := "TEST"
	recordID := "333"
	loadID := "TEST"
	g2engine.DeleteRecord(ctx, dataSourceCode, recordID, loadID)
	// Output:
}

func ExampleG2engineImpl_DeleteRecordWithInfo() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode := "TEST"
	recordID := "333"
	loadID := "TEST"
	var flags int64 = 0
	result, _ := g2engine.DeleteRecordWithInfo(ctx, dataSourceCode, recordID, loadID, flags)
	fmt.Println(result)
	// Output: {"DATA_SOURCE":"TEST","RECORD_ID":"333","AFFECTED_ENTITIES":[],"INTERESTING_ENTITIES":{"ENTITIES":[]}}
}

func ExampleG2engineImpl_ExportCSVEntityReport() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	csvColumnList := ""
	var flags int64 = 0
	responseHandle, _ := g2engine.ExportCSVEntityReport(ctx, csvColumnList, flags)
	fmt.Println(responseHandle > 0) // Dummy output.
	// Output: true
}

func ExampleG2engineImpl_ExportConfig() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	result, _ := g2engine.ExportConfig(ctx)
	fmt.Println(truncate(result, 42))
	// Output: {"G2_CONFIG":{"CFG_ETYPE":[{"ETYPE_ID":...
}

func ExampleG2engineImpl_ExportConfigAndConfigID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	_, configId, _ := g2engine.ExportConfigAndConfigID(ctx)
	fmt.Println(configId > 0) // Dummy output.
	// Output: true
}

func ExampleG2engineImpl_ExportJSONEntityReport() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	flags := int64(0)
	responseHandle, _ := g2engine.ExportJSONEntityReport(ctx, flags)
	fmt.Println(responseHandle > 0) // Dummy output.
	// Output: true
}

func ExampleG2engineImpl_FetchNext() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	flags := int64(0)
	responseHandle, _ := g2engine.ExportJSONEntityReport(ctx, flags)
	anEntity, _ := g2engine.FetchNext(ctx, responseHandle)
	fmt.Println(len(anEntity) >= 0) // Dummy output.
	// Output: true
}

func ExampleG2engineImpl_FindInterestingEntitiesByEntityID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	var entityID int64 = 1
	var flags int64 = 0
	result, _ := g2engine.FindInterestingEntitiesByEntityID(ctx, entityID, flags)
	fmt.Println(result)
	// Output: {"INTERESTING_ENTITIES":{"ENTITIES":[]}}
}

func ExampleG2engineImpl_FindInterestingEntitiesByRecordID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	result, _ := g2engine.FindInterestingEntitiesByRecordID(ctx, dataSourceCode, recordID, flags)
	fmt.Println(len(result) > 0) // Dummy output.
	// Output: true
}

func ExampleG2engineImpl_FindNetworkByEntityID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	entityList := `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}]}`
	maxDegree := 2
	buildOutDegree := 1
	maxEntities := 10
	result, _ := g2engine.FindNetworkByEntityID(ctx, entityList, maxDegree, buildOutDegree, maxEntities)
	fmt.Println(truncate(result, 124))
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,...
}

func ExampleG2engineImpl_FindNetworkByEntityID_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	entityList := `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}]}`
	maxDegree := 2
	buildOutDegree := 1
	maxEntities := 10
	var flags int64 = 0
	result, _ := g2engine.FindNetworkByEntityID_V2(ctx, entityList, maxDegree, buildOutDegree, maxEntities, flags)
	// fmt.Println(truncate(result, 124))
	fmt.Println(result)
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}},{"RESOLVED_ENTITY":{"ENTITY_ID":3}}]}
}

func ExampleG2engineImpl_FindNetworkByRecordID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	recordList := `{"RECORDS": [{"DATA_SOURCE": "TEST", "RECORD_ID": "111"}, {"DATA_SOURCE": "TEST", "RECORD_ID": "222"}]}`
	maxDegree := 1
	buildOutDegree := 2
	maxEntities := 10
	result, _ := g2engine.FindNetworkByRecordID(ctx, recordList, maxDegree, buildOutDegree, maxEntities)
	fmt.Println(truncate(result, 138))
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":...
}

func ExampleG2engineImpl_FindNetworkByRecordID_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	recordList := `{"RECORDS": [{"DATA_SOURCE": "TEST", "RECORD_ID": "111"}, {"DATA_SOURCE": "TEST", "RECORD_ID": "222"}]}`
	maxDegree := 1
	buildOutDegree := 2
	maxEntities := 10
	var flags int64 = 0
	result, _ := g2engine.FindNetworkByRecordID_V2(ctx, recordList, maxDegree, buildOutDegree, maxEntities, flags)
	fmt.Println(result)
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}},{"RESOLVED_ENTITY":{"ENTITY_ID":3}}]}
}

func ExampleG2engineImpl_FindPathByEntityID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	maxDegree := 1
	result, _ := g2engine.FindPathByEntityID(ctx, entityID1, entityID2, maxDegree)
	fmt.Println(truncate(result, 109))
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":...
}

func ExampleG2engineImpl_FindPathByEntityID_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	maxDegree := 1
	var flags int64 = 0
	result, _ := g2engine.FindPathByEntityID_V2(ctx, entityID1, entityID2, maxDegree, flags)
	fmt.Println(result)
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}}]}
}

func ExampleG2engineImpl_FindPathByRecordID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	maxDegree := 1
	result, _ := g2engine.FindPathByRecordID(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree)
	fmt.Println(truncate(result, 89))
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":...
}

func ExampleG2engineImpl_FindPathExcludingByEntityID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	maxDegree := 1
	excludedEntities := `{"ENTITIES": [{"ENTITY_ID": 1}]}`
	result, _ := g2engine.FindPathExcludingByEntityID(ctx, entityID1, entityID2, maxDegree, excludedEntities)
	fmt.Println(truncate(result, 109))
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":...
}

func ExampleG2engineImpl_FindPathExcludingByEntityID_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	maxDegree := 1
	excludedEntities := `{"ENTITIES": [{"ENTITY_ID": 1}]}`
	var flags int64 = 0
	result, _ := g2engine.FindPathExcludingByEntityID_V2(ctx, entityID1, entityID2, maxDegree, excludedEntities, flags)
	fmt.Println(result)
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}}]}
}

func ExampleG2engineImpl_FindPathExcludingByRecordID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	maxDegree := 1
	excludedRecords := `{"RECORDS": [{ "DATA_SOURCE": "TEST", "RECORD_ID": "111"}]}`
	result, _ := g2engine.FindPathExcludingByRecordID(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords)
	fmt.Println(truncate(result, 109))
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":...
}

func ExampleG2engineImpl_FindPathExcludingByRecordID_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	maxDegree := 1
	excludedRecords := `{"RECORDS": [{ "DATA_SOURCE": "TEST", "RECORD_ID": "111"}]}`
	var flags int64 = 0
	result, _ := g2engine.FindPathExcludingByRecordID_V2(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, flags)
	fmt.Println(result)
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}}]}
}

func ExampleG2engineImpl_FindPathIncludingSourceByEntityID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	maxDegree := 1
	excludedEntities := `{"ENTITIES": [{"ENTITY_ID": 1}]}`
	requiredDsrcs := `{"DATA_SOURCES": ["TEST"]}`
	result, _ := g2engine.FindPathIncludingSourceByEntityID(ctx, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs)
	fmt.Println(truncate(result, 106))
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[]}],"ENTITIES":[{"RESOLVED_ENTITY":...
}

func ExampleG2engineImpl_FindPathIncludingSourceByEntityID_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	maxDegree := 1
	excludedEntities := `{"ENTITIES": [{"ENTITY_ID": 1}]}`
	requiredDsrcs := `{"DATA_SOURCES": ["TEST"]}`
	var flags int64 = 0
	result, _ := g2engine.FindPathIncludingSourceByEntityID_V2(ctx, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, flags)
	fmt.Println(result)
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}}]}
}

func ExampleG2engineImpl_FindPathIncludingSourceByRecordID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	maxDegree := 1
	excludedEntities := `{"ENTITIES": [{"ENTITY_ID": 1}]}`
	requiredDsrcs := `{"DATA_SOURCES": ["TEST"]}`
	result, _ := g2engine.FindPathIncludingSourceByRecordID(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedEntities, requiredDsrcs)
	fmt.Println(truncate(result, 119))
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":...
}

func ExampleG2engineImpl_FindPathIncludingSourceByRecordID_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	maxDegree := 1
	excludedEntities := `{"ENTITIES": [{"ENTITY_ID": 1}]}`
	requiredDsrcs := `{"DATA_SOURCES": ["TEST"]}`
	var flags int64 = 0
	result, _ := g2engine.FindPathIncludingSourceByRecordID_V2(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedEntities, requiredDsrcs, flags)
	fmt.Println(result)
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}}]}
}

func ExampleG2engineImpl_GetActiveConfigID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	result, _ := g2engine.GetActiveConfigID(ctx)
	fmt.Println(result > 0) // Dummy output.
	// Output: true
}

func ExampleG2engineImpl_GetEntityByEntityID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	var entityID int64 = 1
	result, _ := g2engine.GetEntityByEntityID(ctx, entityID)
	fmt.Println(truncate(result, 51))
	// Output: {"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":...
}

func ExampleG2engineImpl_GetEntityByEntityID_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	var entityID int64 = 1
	var flags int64 = 0
	result, _ := g2engine.GetEntityByEntityID_V2(ctx, entityID, flags)
	fmt.Println(result)
	// Output: {"RESOLVED_ENTITY":{"ENTITY_ID":1}}
}

func ExampleG2engineImpl_GetEntityByRecordID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode := "TEST"
	recordID := "111"
	result, _ := g2engine.GetEntityByRecordID(ctx, dataSourceCode, recordID)
	fmt.Println(truncate(result, 35))
	// Output: {"RESOLVED_ENTITY":{"ENTITY_ID":...
}

func ExampleG2engineImpl_GetEntityByRecordID_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	result, _ := g2engine.GetEntityByRecordID_V2(ctx, dataSourceCode, recordID, flags)
	fmt.Println(result)
	// Output: {"RESOLVED_ENTITY":{"ENTITY_ID":1}}
}

func ExampleG2engineImpl_GetLastException() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	result, _ := g2engine.GetLastException(ctx)
	fmt.Println(result)
	// Output:
}

func ExampleG2engineImpl_GetLastExceptionCode() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	result, _ := g2engine.GetLastExceptionCode(ctx)
	fmt.Println(result)
	// Output: 0
}

func ExampleG2engineImpl_GetRecord() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode := "TEST"
	recordID := "111"
	result, _ := g2engine.GetRecord(ctx, dataSourceCode, recordID)
	fmt.Println(result)
	// Output: {"DATA_SOURCE":"TEST","RECORD_ID":"111","JSON_DATA":{"SOCIAL_HANDLE":"flavorh","DATE_OF_BIRTH":"4/8/1983","ADDR_STATE":"LA","ADDR_POSTAL_CODE":"71232","SSN_NUMBER":"053-39-3251","GENDER":"F","srccode":"MDMPER","CC_ACCOUNT_NUMBER":"5534202208773608","ADDR_CITY":"Delhi","DRIVERS_LICENSE_STATE":"DE","PHONE_NUMBER":"225-671-0796","NAME_LAST":"SEAMAN","entityid":"284430058","ADDR_LINE1":"772 Armstrong RD","DATA_SOURCE":"TEST","ENTITY_TYPE":"TEST","DSRC_ACTION":"A","RECORD_ID":"111"}}
}

func ExampleG2engineImpl_GetRecord_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	result, _ := g2engine.GetRecord_V2(ctx, dataSourceCode, recordID, flags)
	fmt.Println(result)
	// Output: {"DATA_SOURCE":"TEST","RECORD_ID":"111"}
}

func ExampleG2engineImpl_GetRedoRecord() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	result, _ := g2engine.GetRedoRecord(ctx)
	fmt.Println(result)
	// Output:
}

func ExampleG2engineImpl_GetRepositoryLastModifiedTime() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	result, _ := g2engine.GetRepositoryLastModifiedTime(ctx)
	fmt.Println(result > 0) // Dummy output.
	// Output: true
}

func ExampleG2engineImpl_GetVirtualEntityByRecordID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	recordList := `{"RECORDS": [{"DATA_SOURCE": "TEST","RECORD_ID": "111"},{"DATA_SOURCE": "TEST","RECORD_ID": "222"}]}`
	result, _ := g2engine.GetVirtualEntityByRecordID(ctx, recordList)
	fmt.Println(truncate(result, 51))
	// Output: {"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":...
}

func ExampleG2engineImpl_GetVirtualEntityByRecordID_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	recordList := `{"RECORDS": [{"DATA_SOURCE": "TEST","RECORD_ID": "111"},{"DATA_SOURCE": "TEST","RECORD_ID": "222"}]}`
	var flags int64 = 0
	result, _ := g2engine.GetVirtualEntityByRecordID_V2(ctx, recordList, flags)
	fmt.Println(result)
	// Output: {"RESOLVED_ENTITY":{"ENTITY_ID":1}}
}

func ExampleG2engineImpl_HowEntityByEntityID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	var entityID int64 = 1
	result, _ := g2engine.HowEntityByEntityID(ctx, entityID)
	fmt.Println(result)
	// Output: {"HOW_RESULTS":{"RESOLUTION_STEPS":[],"FINAL_STATE":{"NEED_REEVALUATION":0,"VIRTUAL_ENTITIES":[{"VIRTUAL_ENTITY_ID":"V1","MEMBER_RECORDS":[{"INTERNAL_ID":1,"RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"},{"DATA_SOURCE":"TEST","RECORD_ID":"2D4DABB3FAEAFBD452E9487D06FABC22DC69C846"}]}]}]}}}
}

func ExampleG2engineImpl_HowEntityByEntityID_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	ctx := context.TODO()
	g2engine := getG2Engine(ctx)
	var entityID int64 = 1
	var flags int64 = 0
	result, _ := g2engine.HowEntityByEntityID_V2(ctx, entityID, flags)
	fmt.Println(result)
	// Output: {"HOW_RESULTS":{"RESOLUTION_STEPS":[],"FINAL_STATE":{"NEED_REEVALUATION":0,"VIRTUAL_ENTITIES":[{"VIRTUAL_ENTITY_ID":"V1","MEMBER_RECORDS":[{"INTERNAL_ID":1,"RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"},{"DATA_SOURCE":"TEST","RECORD_ID":"2D4DABB3FAEAFBD452E9487D06FABC22DC69C846"}]}]}]}}}
}

func ExampleG2engineImpl_Init() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	moduleName := "Test module name"
	iniParams, _ := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	verboseLogging := 0
	g2engine.Init(ctx, moduleName, iniParams, verboseLogging)
	// Output:
}

func ExampleG2engineImpl_InitWithConfigID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	moduleName := "Test module name"
	iniParams, _ := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	initConfigID := int64(1)
	verboseLogging := 0
	g2engine.InitWithConfigID(ctx, moduleName, iniParams, initConfigID, verboseLogging)
	// Output:
}

func ExampleG2engineImpl_PrimeEngine() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	g2engine.PrimeEngine(ctx)
	// Output:
}

func ExampleG2engineImpl_Process() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	record := `{"DATA_SOURCE": "TEST", "SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "444", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	g2engine.Process(ctx, record)
	// Output:
}

func ExampleG2engineImpl_ProcessRedoRecord() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	result, _ := g2engine.ProcessRedoRecord(ctx)
	fmt.Println(result)
	// Output:
}

func ExampleG2engineImpl_ProcessRedoRecordWithInfo() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	var flags int64 = 0
	_, result, _ := g2engine.ProcessRedoRecordWithInfo(ctx, flags)
	fmt.Println(result)
	// Output:
}

func ExampleG2engineImpl_ProcessWithInfo() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	record := `{"DATA_SOURCE": "TEST", "SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "555", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	var flags int64 = 0
	result, _ := g2engine.ProcessWithInfo(ctx, record, flags)
	fmt.Println(result)
	// Output: {"DATA_SOURCE":"TEST","RECORD_ID":"555","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}
}

func ExampleG2engineImpl_ProcessWithResponse() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	record := `{"DATA_SOURCE": "TEST", "SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "666", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	result, _ := g2engine.ProcessWithResponse(ctx, record)
	fmt.Println(result)
	// Output: {"MESSAGE": "ER SKIPPED - DUPLICATE RECORD IN G2"}
}

func ExampleG2engineImpl_ProcessWithResponseResize() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	record := `{"DATA_SOURCE": "TEST", "SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "777", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	result, _ := g2engine.ProcessWithResponseResize(ctx, record)
	fmt.Println(result)
	// Output: {"MESSAGE": "ER SKIPPED - DUPLICATE RECORD IN G2"}
}

func ExampleG2engineImpl_ReevaluateEntityWithInfo() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	var entityID int64 = 1
	var flags int64 = 0
	result, _ := g2engine.ReevaluateEntityWithInfo(ctx, entityID, flags)
	fmt.Println(result)
	// Output: {"DATA_SOURCE":"TEST","RECORD_ID":"111","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}
}

func ExampleG2engineImpl_ReevaluateRecord() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	g2engine.ReevaluateRecord(ctx, dataSourceCode, recordID, flags)
	// Output:
}

func ExampleG2engineImpl_ReevaluateRecordWithInfo() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	result, _ := g2engine.ReevaluateRecordWithInfo(ctx, dataSourceCode, recordID, flags)
	fmt.Println(result)
	// Output: {"DATA_SOURCE":"TEST","RECORD_ID":"111","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}

}

func ExampleG2engineImpl_Reinit() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	initConfigID, _ := g2engine.GetActiveConfigID(ctx) // Example initConfigID.
	g2engine.Reinit(ctx, initConfigID)
	// Output:
}

func ExampleG2engineImpl_ReplaceRecordWithInfo() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	dataSourceCode := "TEST"
	recordID := "111"
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1985", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "111", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	var flags int64 = 0
	result, _ := g2engine.ReplaceRecordWithInfo(ctx, dataSourceCode, recordID, jsonData, loadID, flags)
	fmt.Println(result)
	// Output: {"DATA_SOURCE":"TEST","RECORD_ID":"111","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}
}

func ExampleG2engineImpl_Destroy() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	g2engine.Destroy(ctx)
	// Output:
}

func ExampleG2engineImpl_SearchByAttributes() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	jsonData := `{"NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "SEAMAN"}], "SSN_NUMBER": "053-39-3251"}`
	result, _ := g2engine.SearchByAttributes(ctx, jsonData)
	fmt.Println(truncate(result, 54))
	// Output: {"RESOLVED_ENTITIES":[{"MATCH_INFO":{"MATCH_LEVEL":...
}

func ExampleG2engineImpl_SearchByAttributes_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	jsonData := `{"NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "SEAMAN"}], "SSN_NUMBER": "053-39-3251"}`
	var flags int64 = 0
	result, _ := g2engine.SearchByAttributes_V2(ctx, jsonData, flags)
	fmt.Println(result)
	// Output: {"RESOLVED_ENTITIES":[{"MATCH_INFO":{"MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","MATCH_KEY":"+NAME+SSN","ERRULE_CODE":"SF1_PNAME_CSTAB"},"ENTITY":{"RESOLVED_ENTITY":{"ENTITY_ID":1}}}]}
}

func ExampleG2engineImpl_Stats() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	result, _ := g2engine.Stats(ctx)
	fmt.Println(truncate(result, 35))
	// Output: { "workload": { "loadedRecords":...
}

func ExampleG2engineImpl_WhyEntities() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	result, _ := g2engine.WhyEntities(ctx, entityID1, entityID2)
	fmt.Println(truncate(result, 74))
	// Output: {"WHY_RESULTS":[{"ENTITY_ID":1,"ENTITY_ID_2":2,"MATCH_INFO":{"WHY_KEY":...
}

func ExampleG2engineImpl_WhyEntities_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	var flags int64 = 0
	result, _ := g2engine.WhyEntities_V2(ctx, entityID1, entityID2, flags)
	fmt.Println(result)
	// Output: {"WHY_RESULTS":[{"ENTITY_ID":1,"ENTITY_ID_2":2,"MATCH_INFO":{"WHY_KEY":"+PHONE+ACCT_NUM-SSN","WHY_ERRULE_CODE":"SF1","MATCH_LEVEL_CODE":"POSSIBLY_RELATED"}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}}]}
}

func ExampleG2engineImpl_WhyEntityByEntityID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	var entityID int64 = 1
	result, _ := g2engine.WhyEntityByEntityID(ctx, entityID)
	fmt.Println(truncate(result, 101))
	// Output: {"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":...
}

func ExampleG2engineImpl_WhyEntityByEntityID_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	var entityID int64 = 1
	var flags int64 = 0
	result, _ := g2engine.WhyEntityByEntityID_V2(ctx, entityID, flags)
	fmt.Println(truncate(result, 101))
	// Output: {"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":...
}

func ExampleG2engineImpl_WhyEntityByRecordID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	dataSourceCode := "TEST"
	recordID := "111"
	result, _ := g2engine.WhyEntityByRecordID(ctx, dataSourceCode, recordID)
	fmt.Println(truncate(result, 101))
	// Output: {"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":...
}

func ExampleG2engineImpl_WhyEntityByRecordID_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	result, _ := g2engine.WhyEntityByRecordID_V2(ctx, dataSourceCode, recordID, flags)
	fmt.Println(truncate(result, 101))
	// Output: {"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":...
}

func ExampleG2engineImpl_WhyRecords() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	result, _ := g2engine.WhyRecords(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2)
	fmt.Println(truncate(result, 114))
	// Output: {"WHY_RESULTS":[{"INTERNAL_ID":100001,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"}],...
}

func ExampleG2engineImpl_WhyRecords_V2() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/g2engine_test.go
	g2engine := &G2engineImpl{}
	ctx := context.TODO()
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	var flags int64 = 0
	result, _ := g2engine.WhyRecords_V2(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, flags)
	fmt.Println(result)
	// Output: {"WHY_RESULTS":[{"INTERNAL_ID":100001,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"}],"INTERNAL_ID_2":2,"ENTITY_ID_2":2,"FOCUS_RECORDS_2":[{"DATA_SOURCE":"TEST","RECORD_ID":"222"}],"MATCH_INFO":{"WHY_KEY":"+PHONE+ACCT_NUM-DOB-SSN","WHY_ERRULE_CODE":"SF1","MATCH_LEVEL_CODE":"POSSIBLY_RELATED"}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}},{"RESOLVED_ENTITY":{"ENTITY_ID":2}}]}
}
