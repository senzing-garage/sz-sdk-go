package g2config

import (
	"context"
	"fmt"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing/go-helpers/g2engineconfigurationjson"
	"github.com/stretchr/testify/assert"
)

var (
	g2config G2config
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) G2config {

	if g2config == nil {
		g2config = &G2configImpl{}

		moduleName := "Test module name"
		verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
		iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			test.Logf("Cannot construct system configuration. Error: %v", jsonErr)
		}

		initErr := g2config.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			test.Logf("Cannot Init. Error: %v", initErr)
		}
	}
	return g2config
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

func testError(test *testing.T, ctx context.Context, g2config G2config, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		lastException, _ := g2config.GetLastException(ctx)
		assert.FailNow(test, lastException)
	}
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestAddDataSource(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	aHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	inputJson := `{"DSRC_CODE": "GO_TEST"}`
	actual, err := g2config.AddDataSource(ctx, aHandle, inputJson)
	testError(test, ctx, g2config, err)
	printActual(test, actual)
	err = g2config.Close(ctx, aHandle)
	testError(test, ctx, g2config, err)
}

func TestClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	err := g2config.ClearLastException(ctx)
	testError(test, ctx, g2config, err)
}

func TestClose(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	aHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	err = g2config.Close(ctx, aHandle)
	testError(test, ctx, g2config, err)
}

func TestCreate(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	actual, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	printActual(test, actual)
}

func TestDeleteDataSource(test *testing.T) {

	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	aHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)

	actual, err := g2config.ListDataSources(ctx, aHandle)
	testError(test, ctx, g2config, err)
	printResult(test, "Original", actual)

	inputJson := `{"DSRC_CODE": "GO_TEST"}`
	_, err = g2config.AddDataSource(ctx, aHandle, inputJson)
	testError(test, ctx, g2config, err)

	actual, err = g2config.ListDataSources(ctx, aHandle)
	testError(test, ctx, g2config, err)
	printResult(test, "     Add", actual)

	err = g2config.DeleteDataSource(ctx, aHandle, inputJson)
	testError(test, ctx, g2config, err)

	actual, err = g2config.ListDataSources(ctx, aHandle)
	testError(test, ctx, g2config, err)
	printResult(test, "  Delete", actual)

	err = g2config.Close(ctx, aHandle)
	testError(test, ctx, g2config, err)
}

func TestGetLastException(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	actual, err := g2config.GetLastException(ctx)
	if err != nil {
		test.Log("Error:", err.Error())
	} else {
		printActual(test, actual)
	}
}

func TestGetLastExceptionCode(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	actual, err := g2config.GetLastExceptionCode(ctx)
	testError(test, ctx, g2config, err)
	printActual(test, actual)
}

func TestInit(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	testError(test, ctx, g2config, jsonErr)
	err := g2config.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2config, err)
}

func TestListDataSources(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	aHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	actual, err := g2config.ListDataSources(ctx, aHandle)
	testError(test, ctx, g2config, err)
	printActual(test, actual)
	err = g2config.Close(ctx, aHandle)
	testError(test, ctx, g2config, err)
}

func TestLoad(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)

	aHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)

	actual, err := g2config.Save(ctx, aHandle)
	testError(test, ctx, g2config, err)

	err = g2config.Load(ctx, aHandle, actual)
	testError(test, ctx, g2config, err)

}

func TestSave(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	aHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	actual, err := g2config.Save(ctx, aHandle)
	testError(test, ctx, g2config, err)
	printActual(test, actual)
}

func TestDestroy(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	err := g2config.Destroy(ctx)
	testError(test, ctx, g2config, err)
}
