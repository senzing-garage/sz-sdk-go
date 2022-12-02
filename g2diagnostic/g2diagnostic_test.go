package g2diagnostic

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing/g2-sdk-go/g2config"
	"github.com/senzing/g2-sdk-go/g2configmgr"
	"github.com/senzing/g2-sdk-go/g2engine"
	"github.com/senzing/go-helpers/g2engineconfigurationjson"
	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagelogger"
	"github.com/stretchr/testify/assert"
)

const (
	defaultTruncation = 76
)

var (
	g2diagnosticSingleton G2diagnostic
)

var testDataSources = []struct {
	Data string
}{
	{
		Data: `{"DSRC_CODE": "TEST_G2DIAGNOSTIC"}`,
	},
}

var testRecords = []struct {
	DataSource string
	Id         string
	Data       string
	LoadId     string
}{
	{
		DataSource: "TEST_G2DIAGNOSTIC",
		Id:         "9001",
		Data:       `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST_G2DIAGNOSTIC", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "9001", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "TEST",
	},
	{
		DataSource: "TEST_G2DIAGNOSTIC",
		Id:         "9002",
		Data:       `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST_G2DIAGNOSTIC", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "9002", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "TEST",
	},
	{
		DataSource: "TEST_G2DIAGNOSTIC",
		Id:         "9003",
		Data:       `{"ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "ENTITY_TYPE": "TEST_G2DIAGNOSTIC", "GENDER": "M", "srccode": "MDMPER", "RECORD_ID": "9003", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "Smith", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "TEST",
	},
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) G2diagnostic {
	if g2diagnosticSingleton == nil {
		g2diagnosticSingleton = &G2diagnosticImpl{}

		// g2diagnosticSingleton.SetLogLevel(ctx, logger.LevelTrace)
		log.SetFlags(0)

		moduleName := "Test module name"
		verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
		iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			test.Logf("Cannot construct system configuration. Error: %v", jsonErr)
		}
		initErr := g2diagnosticSingleton.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			test.Logf("Cannot Init. Error: %v", initErr)
		}
	}
	return g2diagnosticSingleton
}

func getG2Diagnostic(ctx context.Context) G2diagnostic {
	g2diagnostic := &G2diagnosticImpl{}
	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, _ := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
	return g2diagnostic
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

func testError(test *testing.T, ctx context.Context, g2diagnostic G2diagnostic, err error) {
	if err != nil {
		lastException, _ := g2diagnostic.GetLastException(ctx)
		test.Log("Error:", err.Error())
		assert.FailNow(test, lastException)
	}
}

func testErrorNoFail(test *testing.T, ctx context.Context, g2diagnostic G2diagnostic, err error) {
	if err != nil {
		lastException, _ := g2diagnostic.GetLastException(ctx)
		test.Log("Error:", err.Error(), "LastException:", lastException)
	}
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	code := m.Run()
	err = teardown()
	if err != nil {
		fmt.Print(err)
	}
	os.Exit(code)
}

func setup() error {
	var err error = nil
	ctx := context.TODO()
	now := time.Now()
	moduleName := "Test module name"
	iniParams, _ := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	verboseLogging := 0
	logger, _ := messagelogger.NewSenzingApiLogger(ProductId, IdMessages, IdStatuses, messagelogger.LevelInfo)

	// Purge repository.

	aG2engine := &g2engine.G2engineImpl{}
	err = aG2engine.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		return logger.Error(5901, err)
	}

	err = aG2engine.PurgeRepository(ctx)
	if err != nil {
		return logger.Error(5902, err)
	}

	err = aG2engine.Destroy(ctx)
	if err != nil {
		return logger.Error(5903, err)
	}

	// Add Data Sources to in-memory Senzing configuration.

	aG2config := &g2config.G2configImpl{}
	err = aG2config.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		return logger.Error(5904, err)
	}

	configHandle, _ := aG2config.Create(ctx)
	for _, testDataSource := range testDataSources {
		_, err := aG2config.AddDataSource(ctx, configHandle, testDataSource.Data)
		if err != nil {
			return logger.Error(5905, err)
		}
	}
	configStr, err := aG2config.Save(ctx, configHandle)
	if err != nil {
		return logger.Error(5906, err)
	}

	err = aG2config.Close(ctx, configHandle)
	if err != nil {
		return logger.Error(5907, err)
	}

	err = aG2config.Destroy(ctx)
	if err != nil {
		return logger.Error(5908, err)
	}

	// Persist the Senzing configuration to the Senzing repository.

	aG2configmgr := &g2configmgr.G2configmgrImpl{}
	err = aG2configmgr.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		return logger.Error(5909, err)
	}

	configComments := fmt.Sprintf("Created by g2diagnostic_test at %s", now.UTC())
	configID, err := aG2configmgr.AddConfig(ctx, configStr, configComments)
	if err != nil {
		return logger.Error(5910, err)
	}

	err = aG2configmgr.SetDefaultConfigID(ctx, configID)
	if err != nil {
		return logger.Error(5911, err)
	}

	err = aG2configmgr.Destroy(ctx)
	if err != nil {
		return logger.Error(5912, err)
	}

	// Add records.

	aG2engine = &g2engine.G2engineImpl{}
	err = aG2engine.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		return logger.Error(5913, err)
	}

	for _, testRecord := range testRecords {
		err := aG2engine.AddRecord(ctx, testRecord.DataSource, testRecord.Id, testRecord.Data, testRecord.LoadId)
		if err != nil {
			return logger.Error(5914, err)
		}
	}

	err = aG2engine.Destroy(ctx)
	if err != nil {
		return logger.Error(5914, err)
	}

	return err
}

func teardown() error {
	var err error = nil
	// ctx := context.TODO()
	// moduleName := "Test module name2"
	// iniParams, _ := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	// verboseLogging := 0
	// logger, _ := messagelogger.NewSenzingApiLogger(ProductId, IdMessages, IdStatuses, messagelogger.LevelInfo)

	// // Purge repository.

	// aG2engine := &g2engine.G2engineImpl{}
	// err = aG2engine.Init(ctx, moduleName, iniParams, verboseLogging)
	// if err != nil {
	// 	return logger.Error(5931, err)
	// }

	// err = aG2engine.PurgeRepository(ctx)
	// if err != nil {
	// 	return logger.Error(5932, err)
	// }

	// err = aG2engine.Destroy(ctx)
	// if err != nil {
	// 	return logger.Error(5933, err)
	// }
	return err
}

func TestG2diagnosticImpl_BuildSimpleSystemConfigurationJson(test *testing.T) {
	actual, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, actual)
	}
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestG2diagnosticImpl_CheckDBPerf(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	secondsToRun := 1
	actual, err := g2diagnostic.CheckDBPerf(ctx, secondsToRun)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestG2diagnosticImpl_ClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	g2diagnostic.ClearLastException(ctx)
}

func TestG2diagnosticImpl_EntityListBySize(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	aSize := 1000
	aHandle, err := g2diagnostic.GetEntityListBySize(ctx, aSize)
	testError(test, ctx, g2diagnostic, err)
	anEntity, err := g2diagnostic.FetchNextEntityBySize(ctx, aHandle)
	testError(test, ctx, g2diagnostic, err)
	printResult(test, "Entity", anEntity)
	err = g2diagnostic.CloseEntityListBySize(ctx, aHandle)
	testError(test, ctx, g2diagnostic, err)
}

func TestG2diagnosticImpl_FindEntitiesByFeatureIDs(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	features := "{\"ENTITY_ID\":1,\"LIB_FEAT_IDS\":[1,3,4]}"
	actual, err := g2diagnostic.FindEntitiesByFeatureIDs(ctx, features)
	testError(test, ctx, g2diagnostic, err)
	printResult(test, "len(Actual)", len(actual))
}

func TestG2diagnosticImpl_GetAvailableMemory(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	actual, err := g2diagnostic.GetAvailableMemory(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, int64(0))
	printActual(test, actual)
}

func TestG2diagnosticImpl_GetDataSourceCounts(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	actual, err := g2diagnostic.GetDataSourceCounts(ctx)
	testError(test, ctx, g2diagnostic, err)
	printResult(test, "Data Source counts", actual)
}

func TestG2diagnosticImpl_GetDBInfo(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	actual, err := g2diagnostic.GetDBInfo(ctx)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestG2diagnosticImpl_GetEntityDetails(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	entityID := int64(1)
	includeInternalFeatures := 1
	actual, err := g2diagnostic.GetEntityDetails(ctx, entityID, includeInternalFeatures)
	testErrorNoFail(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestG2diagnosticImpl_GetEntityResume(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	entityID := int64(1)
	actual, err := g2diagnostic.GetEntityResume(ctx, entityID)
	testErrorNoFail(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestG2diagnosticImpl_GetEntitySizeBreakdown(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	minimumEntitySize := 1
	includeInternalFeatures := 1
	actual, err := g2diagnostic.GetEntitySizeBreakdown(ctx, minimumEntitySize, includeInternalFeatures)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestG2diagnosticImpl_GetFeature(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	libFeatID := int64(1)
	actual, err := g2diagnostic.GetFeature(ctx, libFeatID)
	testErrorNoFail(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestG2diagnosticImpl_GetGenericFeatures(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	featureType := "PHONE"
	maximumEstimatedCount := 10
	actual, err := g2diagnostic.GetGenericFeatures(ctx, featureType, maximumEstimatedCount)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestG2diagnosticImpl_GetLastException(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	actual, err := g2diagnostic.GetLastException(ctx)
	testErrorNoFail(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestG2diagnosticImpl_GetLastExceptionCode(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	actual, err := g2diagnostic.GetLastExceptionCode(ctx)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestG2diagnosticImpl_GetLogicalCores(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	actual, err := g2diagnostic.GetLogicalCores(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, 0)
	printActual(test, actual)
}

func TestG2diagnosticImpl_GetMappingStatistics(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	includeInternalFeatures := 1
	actual, err := g2diagnostic.GetMappingStatistics(ctx, includeInternalFeatures)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestG2diagnosticImpl_GetPhysicalCores(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	actual, err := g2diagnostic.GetPhysicalCores(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, 0)
	printActual(test, actual)
}

func TestG2diagnosticImpl_GetRelationshipDetails(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	relationshipID := int64(1)
	includeInternalFeatures := 1
	actual, err := g2diagnostic.GetRelationshipDetails(ctx, relationshipID, includeInternalFeatures)
	testErrorNoFail(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestG2diagnosticImpl_GetResolutionStatistics(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	actual, err := g2diagnostic.GetResolutionStatistics(ctx)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestG2diagnosticImpl_GetTotalSystemMemory(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	actual, err := g2diagnostic.GetTotalSystemMemory(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, int64(0))
	printActual(test, actual)
}

func TestG2diagnosticImpl_Init(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := &G2diagnosticImpl{}
	moduleName := "Test module name"
	verboseLogging := 0
	iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	testError(test, ctx, g2diagnostic, jsonErr)
	err := g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2diagnostic, err)
}

func TestG2diagnosticImpl_InitWithConfigID(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := &G2diagnosticImpl{}
	moduleName := "Test module name"
	initConfigID := int64(1)
	verboseLogging := 0
	iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	testError(test, ctx, g2diagnostic, jsonErr)
	err := g2diagnostic.InitWithConfigID(ctx, moduleName, iniParams, initConfigID, verboseLogging)
	testError(test, ctx, g2diagnostic, err)
}

func TestG2diagnosticImpl_Reinit(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := &G2diagnosticImpl{}
	initConfigID := int64(3173568616) // Must match value in sys_cfg.config_data_id.
	err := g2diagnostic.Reinit(ctx, initConfigID)
	testErrorNoFail(test, ctx, g2diagnostic, err)
}

func TestG2diagnosticImpl_Destroy(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx, test)
	err := g2diagnostic.Destroy(ctx)
	testError(test, ctx, g2diagnostic, err)
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleG2diagnosticImpl_CheckDBPerf() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	secondsToRun := 1
	result, _ := g2diagnostic.CheckDBPerf(ctx, secondsToRun)
	fmt.Println(truncate(result, 25))
	// Output: {"numRecordsInserted":...
}

func ExampleG2diagnosticImpl_ClearLastException() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	g2diagnostic.ClearLastException(ctx)
	// Output:
}

func ExampleG2diagnosticImpl_CloseEntityListBySize() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	aSize := 1000
	entityListBySizeHandle, _ := g2diagnostic.GetEntityListBySize(ctx, aSize)
	g2diagnostic.CloseEntityListBySize(ctx, entityListBySizeHandle)
	// Output:
}

func ExampleG2diagnosticImpl_Destroy() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	g2diagnostic.Destroy(ctx)
	// Output:
}

func ExampleG2diagnosticImpl_FetchNextEntityBySize() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	aSize := 1
	entityListBySizeHandle, _ := g2diagnostic.GetEntityListBySize(ctx, aSize)
	anEntity, _ := g2diagnostic.FetchNextEntityBySize(ctx, entityListBySizeHandle)
	g2diagnostic.CloseEntityListBySize(ctx, entityListBySizeHandle)
	fmt.Println(anEntity)
	// Output: [{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","ENT_SRC_KEY":"2E83FCA24AF2996D77569554CC4FD7775F861A8F","ENT_SRC_DESC":"SEAMAN","RECORD_ID":"9001","JSON_DATA":"{\"SOCIAL_HANDLE\":\"flavorh\",\"DATE_OF_BIRTH\":\"4/8/1983\",\"ADDR_STATE\":\"LA\",\"ADDR_POSTAL_CODE\":\"71232\",\"SSN_NUMBER\":\"053-39-3251\",\"GENDER\":\"F\",\"srccode\":\"MDMPER\",\"CC_ACCOUNT_NUMBER\":\"5534202208773608\",\"ADDR_CITY\":\"Delhi\",\"DRIVERS_LICENSE_STATE\":\"DE\",\"PHONE_NUMBER\":\"225-671-0796\",\"NAME_LAST\":\"SEAMAN\",\"entityid\":\"284430058\",\"ADDR_LINE1\":\"772 Armstrong RD\",\"DATA_SOURCE\":\"TEST_G2DIAGNOSTIC\",\"ENTITY_TYPE\":\"GENERIC\",\"DSRC_ACTION\":\"A\",\"RECORD_ID\":\"9001\"}","OBS_ENT_ID":1,"ER_ID":0},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","ENT_SRC_KEY":"2E83FCA24AF2996D77569554CC4FD7775F861A8F","ENT_SRC_DESC":"SEAMAN","RECORD_ID":"9002","JSON_DATA":"{\"SOCIAL_HANDLE\":\"flavorh\",\"DATE_OF_BIRTH\":\"4/8/1983\",\"ADDR_STATE\":\"LA\",\"ADDR_POSTAL_CODE\":\"71232\",\"SSN_NUMBER\":\"053-39-3251\",\"GENDER\":\"F\",\"srccode\":\"MDMPER\",\"CC_ACCOUNT_NUMBER\":\"5534202208773608\",\"ADDR_CITY\":\"Delhi\",\"DRIVERS_LICENSE_STATE\":\"DE\",\"PHONE_NUMBER\":\"225-671-0796\",\"NAME_LAST\":\"SEAMAN\",\"entityid\":\"284430058\",\"ADDR_LINE1\":\"772 Armstrong RD\",\"DATA_SOURCE\":\"TEST_G2DIAGNOSTIC\",\"ENTITY_TYPE\":\"GENERIC\",\"DSRC_ACTION\":\"A\",\"RECORD_ID\":\"9002\"}","OBS_ENT_ID":1,"ER_ID":0}]
}

func ExampleG2diagnosticImpl_FindEntitiesByFeatureIDs() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	features := `{"ENTITY_ID":1,"LIB_FEAT_IDS":[1,3,4]}`
	result, _ := g2diagnostic.FindEntitiesByFeatureIDs(ctx, features)
	fmt.Println(result)
	// Output: [{"LIB_FEAT_ID":4,"USAGE_TYPE":"","RES_ENT_ID":2}]
}

func ExampleG2diagnosticImpl_GetAvailableMemory() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	result, _ := g2diagnostic.GetAvailableMemory(ctx)
	fmt.Println(result > 0) // Dummy output.
	// Output: true
}

func ExampleG2diagnosticImpl_GetDataSourceCounts() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	result, _ := g2diagnostic.GetDataSourceCounts(ctx)
	fmt.Println(result)
	// Output: [{"DSRC_ID":1001,"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_ID":3,"ETYPE_CODE":"GENERIC","OBS_ENT_COUNT":2,"DSRC_RECORD_COUNT":3}]
}

func ExampleG2diagnosticImpl_GetDBInfo() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	result, _ := g2diagnostic.GetDBInfo(ctx)
	fmt.Println(truncate(result, 52))
	// Output: {"Hybrid Mode":false,"Database Details":[{"Name":...
}

func ExampleG2diagnosticImpl_GetEntityDetails() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	entityID := int64(1)
	includeInternalFeatures := 1
	result, _ := g2diagnostic.GetEntityDetails(ctx, entityID, includeInternalFeatures)
	fmt.Println(result)
	// Output: [{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"NAME","USAGE_TYPE":"","FEAT_DESC":"SEAMAN"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"DOB","USAGE_TYPE":"","FEAT_DESC":"4/8/1983"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"GENDER","USAGE_TYPE":"","FEAT_DESC":"F"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"ADDRESS","USAGE_TYPE":"","FEAT_DESC":"772 Armstrong RD Delhi LA 71232"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"PHONE","USAGE_TYPE":"","FEAT_DESC":"225-671-0796"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"SSN","USAGE_TYPE":"","FEAT_DESC":"053-39-3251"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"LOGIN_ID","USAGE_TYPE":"","FEAT_DESC":"flavorh"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"No","FTYPE_CODE":"ACCT_NUM","USAGE_TYPE":"CC","FEAT_DESC":"5534202208773608"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN|SSN=3251"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN|DOB=80804"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN|PHONE.PHONE_LAST_5=10796"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN|DOB.MMDD_HASH=0804"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN|POST=71232"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN|ADDRESS.CITY_STD=TL"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN|DOB.MMYY_HASH=0483"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","FEAT_DESC":"SMN"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"ADDR_KEY","USAGE_TYPE":"","FEAT_DESC":"772|ARMSTRNK||TL"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"ADDR_KEY","USAGE_TYPE":"","FEAT_DESC":"772|ARMSTRNK||71232"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"ID_KEY","USAGE_TYPE":"","FEAT_DESC":"ACCT_NUM=5534202208773608"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"ID_KEY","USAGE_TYPE":"","FEAT_DESC":"SSN=053-39-3251"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"PHONE_KEY","USAGE_TYPE":"","FEAT_DESC":"2256710796"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"SEARCH_KEY","USAGE_TYPE":"","FEAT_DESC":"LOGIN_ID:FLAVORH|"},{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","DERIVED":"Yes","FTYPE_CODE":"SEARCH_KEY","USAGE_TYPE":"","FEAT_DESC":"SSN:3251|80804|"}]
}

func ExampleG2diagnosticImpl_GetEntityListBySize() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	entitySize := 1000
	entityListBySizeHandle, _ := g2diagnostic.GetEntityListBySize(ctx, entitySize)
	fmt.Println(entityListBySizeHandle > 0) // Dummy output.
	// Output: true
}

func ExampleG2diagnosticImpl_GetEntityResume() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	entityID := int64(1)
	result, _ := g2diagnostic.GetEntityResume(ctx, entityID)
	fmt.Println(result)
	// Output:  [{"RES_ENT_ID":1,"REL_ENT_ID":0,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","ENT_SRC_DESC":"SEAMAN","JSON_DATA":"{\"SOCIAL_HANDLE\":\"flavorh\",\"DATE_OF_BIRTH\":\"4/8/1983\",\"ADDR_STATE\":\"LA\",\"ADDR_POSTAL_CODE\":\"71232\",\"SSN_NUMBER\":\"053-39-3251\",\"GENDER\":\"F\",\"srccode\":\"MDMPER\",\"CC_ACCOUNT_NUMBER\":\"5534202208773608\",\"ADDR_CITY\":\"Delhi\",\"DRIVERS_LICENSE_STATE\":\"DE\",\"PHONE_NUMBER\":\"225-671-0796\",\"NAME_LAST\":\"SEAMAN\",\"entityid\":\"284430058\",\"ADDR_LINE1\":\"772 Armstrong RD\",\"DATA_SOURCE\":\"TEST_G2DIAGNOSTIC\",\"ENTITY_TYPE\":\"GENERIC\",\"DSRC_ACTION\":\"A\",\"RECORD_ID\":\"9001\"}"},{"RES_ENT_ID":1,"REL_ENT_ID":0,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9002","ENT_SRC_DESC":"SEAMAN","JSON_DATA":"{\"SOCIAL_HANDLE\":\"flavorh\",\"DATE_OF_BIRTH\":\"4/8/1983\",\"ADDR_STATE\":\"LA\",\"ADDR_POSTAL_CODE\":\"71232\",\"SSN_NUMBER\":\"053-39-3251\",\"GENDER\":\"F\",\"srccode\":\"MDMPER\",\"CC_ACCOUNT_NUMBER\":\"5534202208773608\",\"ADDR_CITY\":\"Delhi\",\"DRIVERS_LICENSE_STATE\":\"DE\",\"PHONE_NUMBER\":\"225-671-0796\",\"NAME_LAST\":\"SEAMAN\",\"entityid\":\"284430058\",\"ADDR_LINE1\":\"772 Armstrong RD\",\"DATA_SOURCE\":\"TEST_G2DIAGNOSTIC\",\"ENTITY_TYPE\":\"GENERIC\",\"DSRC_ACTION\":\"A\",\"RECORD_ID\":\"9002\"}"},{"RES_ENT_ID":1,"REL_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9003","ENT_SRC_DESC":"Smith","JSON_DATA":"{\"ADDR_STATE\":\"LA\",\"ADDR_POSTAL_CODE\":\"71232\",\"GENDER\":\"M\",\"srccode\":\"MDMPER\",\"ADDR_CITY\":\"Delhi\",\"PHONE_NUMBER\":\"225-671-0796\",\"NAME_LAST\":\"Smith\",\"entityid\":\"284430058\",\"ADDR_LINE1\":\"772 Armstrong RD\",\"DATA_SOURCE\":\"TEST_G2DIAGNOSTIC\",\"ENTITY_TYPE\":\"GENERIC\",\"DSRC_ACTION\":\"A\",\"RECORD_ID\":\"9003\"}"}]
}

func ExampleG2diagnosticImpl_GetEntitySizeBreakdown() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	minimumEntitySize := 1
	includeInternalFeatures := 1
	result, _ := g2diagnostic.GetEntitySizeBreakdown(ctx, minimumEntitySize, includeInternalFeatures)
	fmt.Println(result)
	// Output: [{"ENTITY_SIZE": 1,"ENTITY_COUNT": 2,"NAME": 1.00,"DOB": 0.50,"GENDER": 1.00,"ADDRESS": 1.00,"PHONE": 1.00,"SSN": 0.50,"LOGIN_ID": 0.50,"ACCT_NUM": 0.50,"NAME_KEY": 6.00,"ADDR_KEY": 2.00,"ID_KEY": 1.00,"PHONE_KEY": 1.00,"SEARCH_KEY": 1.00,"MIN_RES_ENT_ID": 1,"MAX_RES_ENT_ID": 2}]
}

func ExampleG2diagnosticImpl_GetFeature() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	libFeatID := int64(1)
	result, _ := g2diagnostic.GetFeature(ctx, libFeatID)
	fmt.Println(result)
	// Output: {"LIB_FEAT_ID":1,"FTYPE_CODE":"NAME","ELEMENTS":[{"FELEM_CODE":"TOKENIZED_NM","FELEM_VALUE":"SEAMAN"},{"FELEM_CODE":"CATEGORY","FELEM_VALUE":"PERSON"},{"FELEM_CODE":"CULTURE","FELEM_VALUE":"ANGLO"},{"FELEM_CODE":"SUR_NAME","FELEM_VALUE":"SEAMAN"},{"FELEM_CODE":"FULL_NAME","FELEM_VALUE":"SEAMAN"}]}
}

func ExampleG2diagnosticImpl_GetGenericFeatures() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	featureType := "PHONE"
	maximumEstimatedCount := 10
	result, _ := g2diagnostic.GetGenericFeatures(ctx, featureType, maximumEstimatedCount)
	fmt.Println(result)
	// Output: []
}

func ExampleG2diagnosticImpl_GetLastException() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	result, _ := g2diagnostic.GetLastException(ctx)
	fmt.Println(result)
	// Output:
}

func ExampleG2diagnosticImpl_GetLastExceptionCode() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	result, _ := g2diagnostic.GetLastExceptionCode(ctx)
	fmt.Println(result)
	// Output: 0
}

func ExampleG2diagnosticImpl_GetLogicalCores() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	result, _ := g2diagnostic.GetLogicalCores(ctx)
	fmt.Println(result > 0) // Dummy output.
	// Output: true
}

func ExampleG2diagnosticImpl_GetMappingStatistics() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	includeInternalFeatures := 1
	result, _ := g2diagnostic.GetMappingStatistics(ctx, includeInternalFeatures)
	fmt.Println(result)
	// Output: [{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"NAME","USAGE_TYPE":"","REC_COUNT":2,"REC_PCT":1.0,"UNIQ_COUNT":2,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"SEAMAN","MAX_FEAT_DESC":"Smith"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"DOB","USAGE_TYPE":"","REC_COUNT":1,"REC_PCT":0.5,"UNIQ_COUNT":1,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"4/8/1983","MAX_FEAT_DESC":"4/8/1983"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"GENDER","USAGE_TYPE":"","REC_COUNT":2,"REC_PCT":1.0,"UNIQ_COUNT":2,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"F","MAX_FEAT_DESC":"M"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"ADDRESS","USAGE_TYPE":"","REC_COUNT":2,"REC_PCT":1.0,"UNIQ_COUNT":1,"UNIQ_PCT":0.5,"MIN_FEAT_DESC":"772 Armstrong RD Delhi LA 71232","MAX_FEAT_DESC":"772 Armstrong RD Delhi LA 71232"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"PHONE","USAGE_TYPE":"","REC_COUNT":2,"REC_PCT":1.0,"UNIQ_COUNT":1,"UNIQ_PCT":0.5,"MIN_FEAT_DESC":"225-671-0796","MAX_FEAT_DESC":"225-671-0796"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"SSN","USAGE_TYPE":"","REC_COUNT":1,"REC_PCT":0.5,"UNIQ_COUNT":1,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"053-39-3251","MAX_FEAT_DESC":"053-39-3251"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"LOGIN_ID","USAGE_TYPE":"","REC_COUNT":1,"REC_PCT":0.5,"UNIQ_COUNT":1,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"flavorh","MAX_FEAT_DESC":"flavorh"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"ACCT_NUM","USAGE_TYPE":"CC","REC_COUNT":1,"REC_PCT":0.5,"UNIQ_COUNT":1,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"5534202208773608","MAX_FEAT_DESC":"5534202208773608"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"Yes","FTYPE_CODE":"NAME_KEY","USAGE_TYPE":"","REC_COUNT":12,"REC_PCT":6.0,"UNIQ_COUNT":12,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"SM0","MAX_FEAT_DESC":"SMN|SSN=3251"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"Yes","FTYPE_CODE":"ADDR_KEY","USAGE_TYPE":"","REC_COUNT":4,"REC_PCT":2.0,"UNIQ_COUNT":2,"UNIQ_PCT":0.5,"MIN_FEAT_DESC":"772|ARMSTRNK||71232","MAX_FEAT_DESC":"772|ARMSTRNK||TL"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"Yes","FTYPE_CODE":"ID_KEY","USAGE_TYPE":"","REC_COUNT":2,"REC_PCT":1.0,"UNIQ_COUNT":2,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"ACCT_NUM=5534202208773608","MAX_FEAT_DESC":"SSN=053-39-3251"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"Yes","FTYPE_CODE":"PHONE_KEY","USAGE_TYPE":"","REC_COUNT":2,"REC_PCT":1.0,"UNIQ_COUNT":1,"UNIQ_PCT":0.5,"MIN_FEAT_DESC":"2256710796","MAX_FEAT_DESC":"2256710796"},{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"Yes","FTYPE_CODE":"SEARCH_KEY","USAGE_TYPE":"","REC_COUNT":2,"REC_PCT":1.0,"UNIQ_COUNT":2,"UNIQ_PCT":1.0,"MIN_FEAT_DESC":"LOGIN_ID:FLAVORH|","MAX_FEAT_DESC":"SSN:3251|80804|"}]
}

func ExampleG2diagnosticImpl_GetPhysicalCores() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	result, _ := g2diagnostic.GetPhysicalCores(ctx)
	fmt.Println(result > 0) // Dummy output.
	// Output: true
}

func ExampleG2diagnosticImpl_GetRelationshipDetails() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	relationshipID := int64(1)
	includeInternalFeatures := 1
	result, _ := g2diagnostic.GetRelationshipDetails(ctx, relationshipID, includeInternalFeatures)
	fmt.Println(result)
	// Output: [{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME","FEAT_DESC":"SEAMAN"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"DOB","FEAT_DESC":"4/8/1983"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"GENDER","FEAT_DESC":"F"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ADDRESS","FEAT_DESC":"772 Armstrong RD Delhi LA 71232"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"PHONE","FEAT_DESC":"225-671-0796"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"SSN","FEAT_DESC":"053-39-3251"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"LOGIN_ID","FEAT_DESC":"flavorh"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ACCT_NUM","FEAT_DESC":"5534202208773608"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN|SSN=3251"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN|DOB=80804"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN|PHONE.PHONE_LAST_5=10796"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN|DOB.MMDD_HASH=0804"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN|POST=71232"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN|ADDRESS.CITY_STD=TL"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN|DOB.MMYY_HASH=0483"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SMN"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ADDR_KEY","FEAT_DESC":"772|ARMSTRNK||TL"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ADDR_KEY","FEAT_DESC":"772|ARMSTRNK||71232"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ID_KEY","FEAT_DESC":"ACCT_NUM=5534202208773608"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ID_KEY","FEAT_DESC":"SSN=053-39-3251"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"PHONE_KEY","FEAT_DESC":"2256710796"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"SEARCH_KEY","FEAT_DESC":"LOGIN_ID:FLAVORH|"},{"RES_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"SEARCH_KEY","FEAT_DESC":"SSN:3251|80804|"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME","FEAT_DESC":"Smith"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"GENDER","FEAT_DESC":"M"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ADDRESS","FEAT_DESC":"772 Armstrong RD Delhi LA 71232"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"PHONE","FEAT_DESC":"225-671-0796"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SM0|POST=71232"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SM0|ADDRESS.CITY_STD=TL"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SM0|PHONE.PHONE_LAST_5=10796"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"NAME_KEY","FEAT_DESC":"SM0"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ADDR_KEY","FEAT_DESC":"772|ARMSTRNK||TL"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"ADDR_KEY","FEAT_DESC":"772|ARMSTRNK||71232"},{"RES_ENT_ID":2,"ERRULE_CODE":"MFF","MATCH_KEY":"+ADDRESS+PHONE-GENDER","FTYPE_CODE":"PHONE_KEY","FEAT_DESC":"2256710796"}]
}

func ExampleG2diagnosticImpl_GetResolutionStatistics() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	result, _ := g2diagnostic.GetResolutionStatistics(ctx)
	fmt.Println(result)
	// Output: [{"MATCH_LEVEL":3,"MATCH_KEY":"+ADDRESS+PHONE-GENDER","RAW_MATCH_KEYS":[{"MATCH_KEY":"+ADDRESS+PHONE-GENDER"}],"ERRULE_ID":200,"ERRULE_CODE":"MFF","IS_AMBIGUOUS":"No","RECORD_COUNT":1,"MIN_RES_ENT_ID":1,"MAX_RES_ENT_ID":2,"MIN_RES_REL_ID":1,"MAX_RES_REL_ID":1}]
}

func ExampleG2diagnosticImpl_GetTotalSystemMemory() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	result, _ := g2diagnostic.GetTotalSystemMemory(ctx)
	fmt.Println(result > 0) // Dummy output.
	// Output: true
}

func ExampleG2diagnosticImpl_Init() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	g2diagnostic := &G2diagnosticImpl{}
	ctx := context.TODO()
	moduleName := "Test module name"
	iniParams, _ := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("") // See https://pkg.go.dev/github.com/senzing/go-helpers
	verboseLogging := 0                                                              // 0 for no Senzing logging; 1 for logging
	g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
	// Output:
}

func ExampleG2diagnosticImpl_InitWithConfigID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	g2diagnostic := &G2diagnosticImpl{}
	ctx := context.TODO()
	moduleName := "Test module name"
	iniParams, _ := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	initConfigID := int64(1)
	verboseLogging := 0
	g2diagnostic.InitWithConfigID(ctx, moduleName, iniParams, initConfigID, verboseLogging)
	// Output:
}

func ExampleG2diagnosticImpl_Reinit() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	g2diagnostic := &G2diagnosticImpl{}
	ctx := context.TODO()
	initConfigID := int64(1)
	g2diagnostic.Reinit(ctx, initConfigID)
	// Output:
}

func ExampleG2diagnosticImpl_SetLogLevel() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	g2diagnostic := &G2diagnosticImpl{}
	ctx := context.TODO()
	g2diagnostic.SetLogLevel(ctx, logger.LevelInfo)
	// Output:
}
