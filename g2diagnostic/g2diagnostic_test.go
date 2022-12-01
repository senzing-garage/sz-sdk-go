package g2diagnostic

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing/g2-sdk-go/g2config"
	"github.com/senzing/g2-sdk-go/g2configmgr"
	"github.com/senzing/g2-sdk-go/g2engine"
	"github.com/senzing/go-helpers/g2engineconfigurationjson"
	"github.com/senzing/go-logging/logger"
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

// Reference: https://medium.com/nerd-for-tech/setup-and-teardown-unit-test-in-go-bd6fa1b785cd
func setupSuite(test *testing.T, ctx context.Context) func(test testing.TB) {

	now := time.Now()
	moduleName := "Test module name"
	iniParams, _ := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	verboseLogging := 0

	// Add Data Sources to in-memory Senzing configuration.

	g2config := &g2config.G2configImpl{}
	err := g2config.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		test.Logf("Cannot Initialize g2config. Error: %v", err)
	}

	configHandle, _ := g2config.Create(ctx)
	for _, testDataSource := range testDataSources {
		_, err := g2config.AddDataSource(ctx, configHandle, testDataSource.Data)
		if err != nil {
			test.Logf("Cannot add data source %s. Error: %v", testDataSource.Data, err)
		}
	}
	configStr, err := g2config.Save(ctx, configHandle)
	if err != nil {
		test.Logf("Cannot save config to string. Error: %v", err)
	}

	err = g2config.Close(ctx, configHandle)
	if err != nil {
		test.Logf("Cannot close configuration handle. Error: %v", err)
	}

	err = g2config.Destroy(ctx)
	if err != nil {
		test.Logf("Cannot destroy g2config. Error: %v", err)
	}

	// Persist the Senzing configuration to the Senzing repository.

	g2configmgr := &g2configmgr.G2configmgrImpl{}
	err = g2configmgr.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		test.Logf("Cannot Initialize g2configmgr. Error: %v", err)
	}

	configComments := fmt.Sprintf("g2diagnostic_test at %s", now.UTC())
	configID, err := g2configmgr.AddConfig(ctx, configStr, configComments)
	if err != nil {
		test.Logf("Cannot add configuration to Senzing repository. Error: %v", err)
	}

	err = g2configmgr.SetDefaultConfigID(ctx, configID)
	if err != nil {
		test.Logf("Cannot set default config ID. Error: %v", err)
	}

	err = g2configmgr.Destroy(ctx)
	if err != nil {
		test.Logf("Cannot destroy g2configmgr. Error: %v", err)
	}

	// Initialize g2engine.

	g2engine := &g2engine.G2engineImpl{}
	err = g2engine.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		test.Logf("Cannot Init. Error: %v", err)
	}

	// Add records.

	for _, testRecord := range testRecords {
		err := g2engine.AddRecord(ctx, testRecord.DataSource, testRecord.Id, testRecord.Data, testRecord.LoadId)
		if err != nil {
			test.Logf("Cannot add test record. Error: %v  Record: %v", err, testRecord)
		}
	}

	// Return a function to teardown the test.
	return func(test testing.TB) {
		// g2engine.PurgeRepository(ctx)
	}
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
	teardownSuite := setupSuite(test, ctx)
	defer teardownSuite(test)
	g2diagnostic := getTestObject(ctx, test)
	entityID := int64(1)
	includeInternalFeatures := 1
	actual, err := g2diagnostic.GetEntityDetails(ctx, entityID, includeInternalFeatures)
	testErrorNoFail(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestG2diagnosticImpl_GetEntityResume(test *testing.T) {
	ctx := context.TODO()
	teardownSuite := setupSuite(test, ctx)
	defer teardownSuite(test)
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
	teardownSuite := setupSuite(test, ctx)
	defer teardownSuite(test)
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
	teardownSuite := setupSuite(test, ctx)
	defer teardownSuite(test)
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
	results, _ := g2diagnostic.CheckDBPerf(ctx, secondsToRun)
	fmt.Println(truncate(results, 25))
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
	aSize := 1000
	entityListBySizeHandle, _ := g2diagnostic.GetEntityListBySize(ctx, aSize)
	anEntity, _ := g2diagnostic.FetchNextEntityBySize(ctx, entityListBySizeHandle)
	fmt.Println(anEntity)
	g2diagnostic.CloseEntityListBySize(ctx, entityListBySizeHandle)
	// Output:
}

func ExampleG2diagnosticImpl_FindEntitiesByFeatureIDs() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	features := `{"ENTITY_ID":1,"LIB_FEAT_IDS":[1,3,4]}`
	result, _ := g2diagnostic.FindEntitiesByFeatureIDs(ctx, features)
	fmt.Println(result)
	// Output: []
}

func ExampleG2diagnosticImpl_GetAvailableMemory() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	result, _ := g2diagnostic.GetAvailableMemory(ctx)
	fmt.Println(result > 0)
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
	fmt.Println(truncate(result, 134))
	// Output: [{"RES_ENT_ID":1,"OBS_ENT_ID":1,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":...
}

func ExampleG2diagnosticImpl_GetEntityListBySize() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	entitySize := 1000
	entityListBySizeHandle, _ := g2diagnostic.GetEntityListBySize(ctx, entitySize)
	fmt.Println(entityListBySizeHandle > 0)
	// Output: true
}

func ExampleG2diagnosticImpl_GetEntityResume() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	entityID := int64(1)
	result, _ := g2diagnostic.GetEntityResume(ctx, entityID)
	fmt.Println(truncate(result, 177))
	// Output: [{"RES_ENT_ID":1,"REL_ENT_ID":0,"ERRULE_CODE":"","MATCH_KEY":"","DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","RECORD_ID":"9001","ENT_SRC_DESC":"SEAMAN","JSON_DATA":...
}

func ExampleG2diagnosticImpl_GetEntitySizeBreakdown() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	minimumEntitySize := 1
	includeInternalFeatures := 1
	result, _ := g2diagnostic.GetEntitySizeBreakdown(ctx, minimumEntitySize, includeInternalFeatures)
	fmt.Println(truncate(result, 19))
	// Output: [{"ENTITY_SIZE":...
}

func ExampleG2diagnosticImpl_GetFeature() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	libFeatID := int64(1)
	result, _ := g2diagnostic.GetFeature(ctx, libFeatID)
	fmt.Println(truncate(result, 95))
	// Output: {"LIB_FEAT_ID":1,"FTYPE_CODE":"NAME","ELEMENTS":[{"FELEM_CODE":"TOKENIZED_NM","FELEM_VALUE":...
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
	fmt.Println(result > 0)
	// Output: true
}

func ExampleG2diagnosticImpl_GetMappingStatistics() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	includeInternalFeatures := 1
	result, _ := g2diagnostic.GetMappingStatistics(ctx, includeInternalFeatures)
	fmt.Println(truncate(result, 169))
	// Output: [{"DSRC_CODE":"TEST_G2DIAGNOSTIC","ETYPE_CODE":"GENERIC","DERIVED":"No","FTYPE_CODE":"NAME","USAGE_TYPE":"","REC_COUNT":2,"REC_PCT":1.0,"UNIQ_COUNT":2,"UNIQ_PCT":1.0,...
}

func ExampleG2diagnosticImpl_GetPhysicalCores() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	result, _ := g2diagnostic.GetPhysicalCores(ctx)
	fmt.Println(result > 0)
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
	// Output: FIXME:
}

func ExampleG2diagnosticImpl_GetResolutionStatistics() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	result, _ := g2diagnostic.GetResolutionStatistics(ctx)
	fmt.Println(result)
	// Output: [{"MATCH_LEVEL":1,"MATCH_KEY":"+NAME+DOB+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM-GENDER","RAW_MATCH_KEYS":[{"MATCH_KEY":"+NAME+DOB+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM-GENDER"}],"ERRULE_ID":140,"ERRULE_CODE":"MSTAB_LNAME_SF1","IS_AMBIGUOUS":"No","RECORD_COUNT":1,"MIN_RES_ENT_ID":1,"MAX_RES_ENT_ID":1,"MIN_RES_REL_ID":0,"MAX_RES_REL_ID":0}]
}

func ExampleG2diagnosticImpl_GetTotalSystemMemory() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/g2diagnostic_test.go
	ctx := context.TODO()
	g2diagnostic := getG2Diagnostic(ctx)
	result, _ := g2diagnostic.GetTotalSystemMemory(ctx)
	fmt.Println(result > 0)
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
