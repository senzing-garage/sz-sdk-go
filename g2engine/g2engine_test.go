package g2engine

import (
	"context"
	"fmt"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing/go-helpers/g2engineconfigurationjson"
	"github.com/stretchr/testify/assert"
)

var (
	g2engine G2engine
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) G2engine {

	if g2engine == nil {
		g2engine = &G2engineImpl{}

		moduleName := "Test module name"
		verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
		iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			test.Logf("Cannot construct system configuration. Error: %v", jsonErr)
		}

		initErr := g2engine.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			test.Logf("Cannot Init. Error: %v", initErr)
		}
	}
	return g2engine
}

func truncate(aString string) string {
	return truncator.Truncate(aString, 50, "...", truncator.PositionEnd)
}

func printResult(test *testing.T, title string, result interface{}) {
	if 1 == 0 {
		test.Logf("%s: %v", title, truncate(fmt.Sprintf("%v", result)))
	}
}

func printActual(test *testing.T, actual interface{}) {
	printResult(test, "Actual", actual)
}

func testError(test *testing.T, ctx context.Context, g2engine G2engine, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		lastException, _ := g2engine.GetLastException(ctx)
		assert.FailNow(test, lastException)
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

func TestAddRecord(test *testing.T) {
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

func TestAddRecordWithInfo(test *testing.T) {
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

func TestAddRecordWithInfoWithReturnedRecordID(test *testing.T) {
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

func TestAddRecordWithReturnedRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	jsonData := `{"SOCIAL_HANDLE": "bobby", "DATE_OF_BIRTH": "1/2/1983", "ADDR_STATE": "WI", "ADDR_POSTAL_CODE": "54434", "SSN_NUMBER": "987-65-4321", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "Smith", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	actual, err := g2engine.AddRecordWithReturnedRecordID(ctx, dataSourceCode, jsonData, loadID)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestCheckRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	record := `{"DATA_SOURCE": "TEST", "NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "Smith", "NAME_MIDDLE": "M" }], "PASSPORT_NUMBER": "PP11111", "PASSPORT_COUNTRY": "US", "DRIVERS_LICENSE_NUMBER": "DL11111", "SSN_NUMBER": "111-11-1111"}`
	recordQueryList := `{"RECORDS": [{"DATA_SOURCE": "TEST","RECORD_ID": "111"},{"DATA_SOURCE": "TEST","RECORD_ID": "123456789"}]}`
	actual, err := g2engine.CheckRecord(ctx, record, recordQueryList)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	err := g2engine.ClearLastException(ctx)
	testError(test, ctx, g2engine, err)
}

// FAIL:
func TestExportJSONEntityReport(test *testing.T) {
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

func TestCountRedoRecords(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.CountRedoRecords(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestExportConfigAndConfigID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actualConfig, actualConfigId, err := g2engine.ExportConfigAndConfigID(ctx)
	testError(test, ctx, g2engine, err)
	printResult(test, "Actual Config", actualConfig)
	printResult(test, "Actual Config ID", actualConfigId)
}

func TestExportConfig(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.ExportConfig(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

//func TestExportCSVEntityReport(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx, test)
//	csvColumnList := ""
//	var flags int64 = 0
//	actual, err := g2engine.ExportCSVEntityReport(ctx, csvColumnList, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//

func TestFindInterestingEntitiesByEntityID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	var flags int64 = 0
	actual, err := g2engine.FindInterestingEntitiesByEntityID(ctx, entityID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestFindInterestingEntitiesByRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	actual, err := g2engine.FindInterestingEntitiesByRecordID(ctx, dataSourceCode, recordID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestFindNetworkByEntityID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	entityList := `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
	maxDegree := 2
	buildOutDegree := 1
	maxEntities := 10
	actual, err := g2engine.FindNetworkByEntityID(ctx, entityList, maxDegree, buildOutDegree, maxEntities)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestFindNetworkByEntityID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	entityList := `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
	maxDegree := 2
	buildOutDegree := 1
	maxEntities := 10
	var flags int64 = 0
	actual, err := g2engine.FindNetworkByEntityID_V2(ctx, entityList, maxDegree, buildOutDegree, maxEntities, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestFindNetworkByRecordID(test *testing.T) {
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

func TestFindNetworkByRecordID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	recordList := ""
	maxDegree := 1
	buildOutDegree := 2
	maxEntities := 10
	var flags int64 = 0
	actual, err := g2engine.FindNetworkByRecordID_V2(ctx, recordList, maxDegree, buildOutDegree, maxEntities, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestFindPathByEntityID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	maxDegree := 1
	actual, err := g2engine.FindPathByEntityID(ctx, entityID1, entityID2, maxDegree)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestFindPathByEntityID_V2(test *testing.T) {
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

func TestFindPathByRecordID(test *testing.T) {
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

func TestFindPathByRecordID_V2(test *testing.T) {
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

func TestFindPathExcludingByEntityID(test *testing.T) {
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

func TestFindPathExcludingByEntityID_V2(test *testing.T) {
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

func TestFindPathExcludingByRecordID(test *testing.T) {
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

func TestFindPathExcludingByRecordID_V2(test *testing.T) {
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

func TestFindPathIncludingSourceByEntityID(test *testing.T) {
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

func TestFindPathIncludingSourceByEntityID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	maxDegree := 1
	excludedEntities := ""
	requiredDsrcs := ""
	var flags int64 = 0
	actual, err := g2engine.FindPathIncludingSourceByEntityID_V2(ctx, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestFindPathIncludingSourceByRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	maxDegree := 1
	excludedRecords := ""
	requiredDsrcs := ""
	actual, err := g2engine.FindPathIncludingSourceByRecordID(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestFindPathIncludingSourceByRecordID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode1 := "TEST"
	recordID1 := "111"
	dataSourceCode2 := "TEST"
	recordID2 := "222"
	maxDegree := 1
	excludedRecords := ""
	requiredDsrcs := ""
	var flags int64 = 0
	actual, err := g2engine.FindPathIncludingSourceByRecordID_V2(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestGetActiveConfigID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.GetActiveConfigID(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestGetEntityByEntityID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	actual, err := g2engine.GetEntityByEntityID(ctx, entityID)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestGetEntityByEntityID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	var flags int64 = 0
	actual, err := g2engine.GetEntityByEntityID_V2(ctx, entityID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestGetEntityByRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	actual, err := g2engine.GetEntityByRecordID(ctx, dataSourceCode, recordID)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestGetEntityByRecordID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	actual, err := g2engine.GetEntityByRecordID_V2(ctx, dataSourceCode, recordID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestGetLastException(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.GetLastException(ctx)
	if err == nil {
		printActual(test, actual)
	}
}

func TestGetLastExceptionCode(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.GetLastExceptionCode(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestGetRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	actual, err := g2engine.GetRecord(ctx, dataSourceCode, recordID)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestGetRecord_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	actual, err := g2engine.GetRecord_V2(ctx, dataSourceCode, recordID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestGetRedoRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.GetRedoRecord(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestGetRepositoryLastModifiedTime(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.GetRepositoryLastModifiedTime(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestGetVirtualEntityByRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	recordList := `{"RECORDS": [{"DATA_SOURCE": "TEST","RECORD_ID": "111"},{"DATA_SOURCE": "TEST","RECORD_ID": "222"}]}`
	actual, err := g2engine.GetVirtualEntityByRecordID(ctx, recordList)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestGetVirtualEntityByRecordID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	recordList := `{"RECORDS": [{"DATA_SOURCE": "TEST","RECORD_ID": "111"},{"DATA_SOURCE": "TEST","RECORD_ID": "222"}]}`
	var flags int64 = 0
	actual, err := g2engine.GetVirtualEntityByRecordID_V2(ctx, recordList, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestHowEntityByEntityID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	actual, err := g2engine.HowEntityByEntityID(ctx, entityID)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestHowEntityByEntityID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	var flags int64 = 0
	actual, err := g2engine.HowEntityByEntityID_V2(ctx, entityID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestInit(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	testError(test, ctx, g2engine, jsonErr)
	err := g2engine.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2engine, err)
}

func TestInitWithConfigID(test *testing.T) {
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

func TestPrimeEngine(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	err := g2engine.PrimeEngine(ctx)
	testError(test, ctx, g2engine, err)
}

func TestProcess(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	record := `{"DATA_SOURCE": "TEST", "SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "444", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	err := g2engine.Process(ctx, record)
	testError(test, ctx, g2engine, err)
}

func TestProcessRedoRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.ProcessRedoRecord(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestProcessRedoRecordWithInfo(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var flags int64 = 0
	actual, actualInfo, err := g2engine.ProcessRedoRecordWithInfo(ctx, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
	printResult(test, "Actual Info", actualInfo)
}

func TestProcessWithInfo(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	record := `{"DATA_SOURCE": "TEST", "SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "555", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	var flags int64 = 0
	actual, err := g2engine.ProcessWithInfo(ctx, record, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestProcessWithResponse(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	record := `{"DATA_SOURCE": "TEST", "SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "666", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	actual, err := g2engine.ProcessWithResponse(ctx, record)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestProcessWithResponseResize(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	record := `{"DATA_SOURCE": "TEST", "SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "777", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	actual, err := g2engine.ProcessWithResponseResize(ctx, record)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestReevaluateEntity(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	var flags int64 = 0
	err := g2engine.ReevaluateEntity(ctx, entityID, flags)
	testError(test, ctx, g2engine, err)
}

func TestReevaluateEntityWithInfo(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	var flags int64 = 0
	actual, err := g2engine.ReevaluateEntityWithInfo(ctx, entityID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestReevaluateRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	err := g2engine.ReevaluateRecord(ctx, dataSourceCode, recordID, flags)
	testError(test, ctx, g2engine, err)
}

func TestReevaluateRecordWithInfo(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	actual, err := g2engine.ReevaluateRecordWithInfo(ctx, dataSourceCode, recordID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestReinit(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	initConfigID, err := g2engine.GetActiveConfigID(ctx)
	testError(test, ctx, g2engine, err)
	err = g2engine.Reinit(ctx, initConfigID)
	testError(test, ctx, g2engine, err)
	printActual(test, initConfigID)
}

func TestReplaceRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1984", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "111", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	err := g2engine.ReplaceRecord(ctx, dataSourceCode, recordID, jsonData, loadID)
	testError(test, ctx, g2engine, err)
}

func TestReplaceRecordWithInfo(test *testing.T) {
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

func TestSearchByAttributes(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	jsonData := `{"NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "SEAMAN"}], "SSN_NUMBER": "053-39-3251"}`
	actual, err := g2engine.SearchByAttributes(ctx, jsonData)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestSearchByAttributes_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	jsonData := `{"NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "SEAMAN"}], "SSN_NUMBER": "053-39-3251"}`
	var flags int64 = 0
	actual, err := g2engine.SearchByAttributes_V2(ctx, jsonData, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestStats(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	actual, err := g2engine.Stats(ctx)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestWhyEntities(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	actual, err := g2engine.WhyEntities(ctx, entityID1, entityID2)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestWhyEntities_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID1 int64 = 1
	var entityID2 int64 = 2
	var flags int64 = 0
	actual, err := g2engine.WhyEntities_V2(ctx, entityID1, entityID2, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestWhyEntityByEntityID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	actual, err := g2engine.WhyEntityByEntityID(ctx, entityID)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestWhyEntityByEntityID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	var entityID int64 = 1
	var flags int64 = 0
	actual, err := g2engine.WhyEntityByEntityID_V2(ctx, entityID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestWhyEntityByRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	actual, err := g2engine.WhyEntityByRecordID(ctx, dataSourceCode, recordID)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestWhyEntityByRecordID_V2(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	var flags int64 = 0
	actual, err := g2engine.WhyEntityByRecordID_V2(ctx, dataSourceCode, recordID, flags)
	testError(test, ctx, g2engine, err)
	printActual(test, actual)
}

func TestWhyRecords(test *testing.T) {
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

func TestWhyRecords_V2(test *testing.T) {
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

func TestDeleteRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	dataSourceCode := "TEST"
	recordID := "111"
	loadID := "TEST"
	err := g2engine.DeleteRecord(ctx, dataSourceCode, recordID, loadID)
	testError(test, ctx, g2engine, err)
}

func TestDeleteRecordWithInfo(test *testing.T) {
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

func TestPurgeRepository(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	err := g2engine.PurgeRepository(ctx)
	testError(test, ctx, g2engine, err)
}

func TestDestroy(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx, test)
	err := g2engine.Destroy(ctx)
	testError(test, ctx, g2engine, err)
}
