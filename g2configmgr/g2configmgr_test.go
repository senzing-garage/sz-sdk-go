package g2configmgr

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing/g2-sdk-go/g2config"
	"github.com/senzing/go-helpers/g2engineconfigurationjson"
	"github.com/senzing/go-logging/logger"
	"github.com/stretchr/testify/assert"
)

const (
	defaultTruncation = 76
)

var (
	g2configmgrSingleton G2configmgr
	g2configSingleton    g2config.G2config
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) G2configmgr {

	if g2configmgrSingleton == nil {
		g2configmgrSingleton = &G2configmgrImpl{}

		// g2configmgrSingleton.SetLogLevel(ctx, logger.LevelTrace)
		log.SetFlags(0)

		moduleName := "Test module name"
		verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
		iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			test.Logf("Cannot construct system configuration. Error: %v", jsonErr)
		}

		initErr := g2configmgrSingleton.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			test.Logf("Cannot Init. Error: %v", initErr)
		}
	}
	return g2configmgrSingleton
}

func getG2Configmgr(ctx context.Context) G2configmgr {
	g2configmgr := &G2configmgrImpl{}
	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, _ := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	g2configmgr.Init(ctx, moduleName, iniParams, verboseLogging)
	return g2configmgr
}

func getG2Config(ctx context.Context, test *testing.T) g2config.G2config {

	if g2configSingleton == nil {
		g2configSingleton = &g2config.G2configImpl{}

		moduleName := "Test module name"
		verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
		iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			test.Logf("Cannot construct system configuration. Error: %v", jsonErr)
		}

		initErr := g2configSingleton.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			test.Logf("Cannot Init. Error: %v", initErr)
		}
	}
	return g2configSingleton
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

func testError(test *testing.T, ctx context.Context, g2configmgr G2configmgr, err error) {
	if err != nil {
		lastException, _ := g2configmgr.GetLastException(ctx)
		test.Log("Error:", err.Error())
		assert.FailNow(test, lastException)
	}
}

func testErrorNoFail(test *testing.T, ctx context.Context, g2configmgr G2configmgr, err error) {
	if err != nil {
		lastException, _ := g2configmgr.GetLastException(ctx)
		test.Log("Error:", err.Error(), "LastException:", lastException)
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleG2configmgrImpl_AddConfig() {
	// Create an in-memory configuration.
	ctx := context.TODO()
	g2config := &g2config.G2configImpl{}
	configHandle, _ := g2config.Create(ctx)

	g2configmgr := getG2Configmgr(ctx) // See https://github.com/Senzing/g2-sdk-go/blob/main/g2configmgr/g2configmgr_test.go
	configStr, _ := g2config.Save(ctx, configHandle)
	configComments := "Example configuration"
	configID, _ := g2configmgr.AddConfig(ctx, configStr, configComments)
	fmt.Println(configID > 0)
	// Output: true
}

func ExampleG2configmgrImpl_ClearLastException() {
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx) // See https://github.com/Senzing/g2-sdk-go/blob/main/g2configmgr/g2configmgr_test.go
	g2configmgr.ClearLastException(ctx)
	// Output:
}

func ExampleG2configmgrImpl_Destroy() {
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx) // See https://github.com/Senzing/g2-sdk-go/blob/main/g2configmgr/g2configmgr_test.go
	g2configmgr.Destroy(ctx)
	// Output:
}

func ExampleG2configmgrImpl_GetConfig() {
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx) // See https://github.com/Senzing/g2-sdk-go/blob/main/g2configmgr/g2configmgr_test.go
	configID, _ := g2configmgr.GetDefaultConfigID(ctx)
	configStr, _ := g2configmgr.GetConfig(ctx, configID)
	fmt.Println(truncate(configStr, defaultTruncation))
	// Output: {"G2_CONFIG":{"CFG_ATTR":[{"ATTR_ID":1001,"ATTR_CODE":"DATA_SOURCE","ATTR...
}

func ExampleG2configmgrImpl_GetConfigList() {
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx) // See https://github.com/Senzing/g2-sdk-go/blob/main/g2configmgr/g2configmgr_test.go
	jsonConfigList, _ := g2configmgr.GetConfigList(ctx)
	fmt.Println(truncate(jsonConfigList, 28))
	// Output: {"CONFIGS":[{"CONFIG_ID":...
}

func ExampleG2configmgrImpl_GetDefaultConfigID() {
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx) // See https://github.com/Senzing/g2-sdk-go/blob/main/g2configmgr/g2configmgr_test.go
	configID, _ := g2configmgr.GetDefaultConfigID(ctx)
	fmt.Println(configID > 0)
	// Output: true
}

func ExampleG2configmgrImpl_GetLastException() {
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx) // See https://github.com/Senzing/g2-sdk-go/blob/main/g2configmgr/g2configmgr_test.go
	result, _ := g2configmgr.GetLastException(ctx)
	fmt.Println(result)
	// Output:
}

func ExampleG2configmgrImpl_GetLastExceptionCode() {
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx) // See https://github.com/Senzing/g2-sdk-go/blob/main/g2configmgr/g2configmgr_test.go
	result, _ := g2configmgr.GetLastExceptionCode(ctx)
	fmt.Println(result)
	// Output: 0
}

func ExampleG2configmgrImpl_Init() {
	g2configmgr := &G2configmgrImpl{}
	ctx := context.TODO()
	moduleName := "Test module name"
	iniParams, _ := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("") // See https://pkg.go.dev/github.com/senzing/go-helpers
	verboseLogging := 0                                                              // 0 for no Senzing logging; 1 for logging
	g2configmgr.Init(ctx, moduleName, iniParams, verboseLogging)
	// Output:
}

func ExampleG2configmgrImpl_ReplaceDefaultConfigID() {
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx) // See https://github.com/Senzing/g2-sdk-go/blob/main/g2configmgr/g2configmgr_test.go
	oldConfigID, _ := g2configmgr.GetDefaultConfigID(ctx)

	// Create an example configuration.
	g2config := &g2config.G2configImpl{}
	configHandle, _ := g2config.Create(ctx)
	configStr, _ := g2config.Save(ctx, configHandle)
	configComments := "Example configuration"
	newConfigID, _ := g2configmgr.AddConfig(ctx, configStr, configComments)

	g2configmgr.ReplaceDefaultConfigID(ctx, oldConfigID, newConfigID)
	// Output:
}

func ExampleG2configmgrImpl_SetDefaultConfigID() {
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx)                 // See https://github.com/Senzing/g2-sdk-go/blob/main/g2configmgr/g2configmgr_test.go
	configID, _ := g2configmgr.GetDefaultConfigID(ctx) // For example purposes only. Normally would use output from GetConfigList()
	g2configmgr.SetDefaultConfigID(ctx, configID)
	// Output:
}

func ExampleG2configmgrImpl_SetLogLevel() {
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx) // See https://github.com/Senzing/g2-sdk-go/blob/main/g2configmgr/g2configmgr_test.go
	g2configmgr.SetLogLevel(ctx, logger.LevelInfo)
	// Output:
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

func TestG2configmgrImpl_AddConfig(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	now := time.Now()

	// Create an in-memory configuration.

	g2config := getG2Config(ctx, test)
	configHandle, err1 := g2config.Create(ctx)
	if err1 != nil {
		test.Log("Error:", err1.Error())
		assert.FailNow(test, "g2config.Create()")
	}

	// Modify the in-memory configuration so it is different from the created configuration.
	// If not, on Save Senzing will detect that it is the same and no Save occurs.

	inputJson := `{"DSRC_CODE": "GO_TEST_` + strconv.FormatInt(now.Unix(), 10) + `"}`
	_, err2 := g2config.AddDataSource(ctx, configHandle, inputJson)
	if err2 != nil {
		test.Log("Error:", err2.Error())
		assert.FailNow(test, "g2config.AddDataSource()")
	}

	// Create a JSON string from the in-memory version of the configuration.

	configStr, err3 := g2config.Save(ctx, configHandle)
	if err3 != nil {
		test.Log("Error:", err2.Error())
		assert.FailNow(test, configStr)
	}

	// Perform the test.

	configComments := fmt.Sprintf("g2configmgr_test at %s", now.UTC())
	actual, err := g2configmgr.AddConfig(ctx, configStr, configComments)
	testError(test, ctx, g2configmgr, err)
	printActual(test, actual)
}

func TestG2configmgrImpl_ClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	err := g2configmgr.ClearLastException(ctx)
	testError(test, ctx, g2configmgr, err)
}

func TestG2configmgrImpl_GetConfig(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)

	// Get a ConfigID.

	configID, err1 := g2configmgr.GetDefaultConfigID(ctx)
	if err1 != nil {
		test.Log("Error:", err1.Error())
		assert.FailNow(test, "g2configmgr.GetDefaultConfigID()")
	}

	actual, err := g2configmgr.GetConfig(ctx, configID)
	testError(test, ctx, g2configmgr, err)
	printActual(test, actual)
}

func TestG2configmgrImpl_GetConfigList(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	actual, err := g2configmgr.GetConfigList(ctx)
	testError(test, ctx, g2configmgr, err)
	printActual(test, actual)
}

func TestG2configmgrImpl_GetDefaultConfigID(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	actual, err := g2configmgr.GetDefaultConfigID(ctx)
	testError(test, ctx, g2configmgr, err)
	printActual(test, actual)
}

func TestG2configmgrImpl_GetLastException(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	actual, err := g2configmgr.GetLastException(ctx)
	testErrorNoFail(test, ctx, g2configmgr, err)
	printActual(test, actual)
}

func TestG2configmgrImpl_GetLastExceptionCode(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	actual, err := g2configmgr.GetLastExceptionCode(ctx)
	testError(test, ctx, g2configmgr, err)
	printActual(test, actual)
}

func TestG2configmgrImpl_Init(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	if jsonErr != nil {
		test.Fatalf("Cannot construct system configuration: %v", jsonErr)
	}
	err := g2configmgr.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2configmgr, err)
}

func TestG2configmgrImpl_ReplaceDefaultConfigID(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)

	oldConfigID, err1 := g2configmgr.GetDefaultConfigID(ctx)
	if err1 != nil {
		test.Log("Error:", err1.Error())
		assert.FailNow(test, "g2configmgr.GetDefaultConfigID()")
	}

	// FIXME: This is kind of a cheeter.

	newConfigID, err2 := g2configmgr.GetDefaultConfigID(ctx)
	if err2 != nil {
		test.Log("Error:", err2.Error())
		assert.FailNow(test, "g2configmgr.GetDefaultConfigID()-2")
	}

	err := g2configmgr.ReplaceDefaultConfigID(ctx, oldConfigID, newConfigID)
	testError(test, ctx, g2configmgr, err)
}

func TestG2configmgrImpl_SetDefaultConfigID(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	configID, err1 := g2configmgr.GetDefaultConfigID(ctx)
	if err1 != nil {
		test.Log("Error:", err1.Error())
		assert.FailNow(test, "g2configmgr.GetDefaultConfigID()")
	}
	err := g2configmgr.SetDefaultConfigID(ctx, configID)
	testError(test, ctx, g2configmgr, err)
}

func TestG2configmgrImpl_Destroy(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	err := g2configmgr.Destroy(ctx)
	testError(test, ctx, g2configmgr, err)
}
