package g2product

import (
	"context"
	"fmt"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing/go-helpers/g2engineconfigurationjson"
	"github.com/stretchr/testify/assert"
)

var (
	g2product G2product
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) G2product {

	if g2product == nil {
		g2product = &G2productImpl{}

		moduleName := "Test module name"
		verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
		iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			test.Logf("Cannot construct system configuration. Error: %v", jsonErr)
		}

		initErr := g2product.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			test.Logf("Cannot Init. Error: %v", initErr)
		}
	}
	return g2product
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

func testError(test *testing.T, ctx context.Context, g2product G2product, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		lastException, _ := g2product.GetLastException(ctx)
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

func TestClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx, test)
	g2product.ClearLastException(ctx)
}

func TestGetLastException(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx, test)
	actual, err := g2product.GetLastException(ctx)
	if err == nil {
		printActual(test, actual)
	}
}

func TestGetLastExceptionCode(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx, test)
	actual, err := g2product.GetLastExceptionCode(ctx)
	testError(test, ctx, g2product, err)
	printActual(test, actual)
}

func TestInit(test *testing.T) {
	ctx := context.TODO()
	g2product := &G2productImpl{}
	moduleName := "Test module name"
	verboseLogging := 0
	iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	testError(test, ctx, g2product, jsonErr)
	err := g2product.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2product, err)
}

func TestLicense(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx, test)
	actual, err := g2product.License(ctx)
	testError(test, ctx, g2product, err)
	printActual(test, actual)
}

func TestValidateLicenseFile(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx, test)
	licenseFilePath := ""
	actual, _ := g2product.ValidateLicenseFile(ctx, licenseFilePath)
	// testError(test, ctx, g2product, err)
	printActual(test, actual)
}

func TestValidateLicenseStringBase64(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx, test)
	licenseString := ""
	actual, _ := g2product.ValidateLicenseStringBase64(ctx, licenseString)
	// testError(test, ctx, g2product, err)
	printActual(test, actual)
}

func TestVersion(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx, test)
	actual, err := g2product.Version(ctx)
	testError(test, ctx, g2product, err)
	printActual(test, actual)
}

func TestDestroy(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx, test)
	err := g2product.Destroy(ctx)
	testError(test, ctx, g2product, err)
}
