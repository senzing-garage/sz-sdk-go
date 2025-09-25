package response_test

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing-garage/sz-sdk-go/response"
	"github.com/stretchr/testify/require"
)

const (
	defaultTruncation = 127
	printResults      = false
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func closeFile(t *testing.T, file *os.File) {
	t.Helper()

	err := file.Close()
	if err != nil {
		t.Fatalf("Could not close file: %v", err)
	}
}

func createScanner(fileName string) (*bufio.Scanner, *os.File) {
	testFilePath := createTestFilePath(fileName)

	file, err := os.Open(testFilePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 1024*1024)

	return scanner, file
}

func createTestFilePath(filename string) string {
	return filepath.Join("..", "testdata", "responses_senzing", filename)
}

func truncate(aString string, length int) string {
	return truncator.Truncate(aString, length, "...", truncator.PositionEnd)
}

func printResult(t *testing.T, title string, result interface{}) {
	t.Helper()

	if printResults {
		t.Logf("%s: %+v", title, truncate(fmt.Sprintf("%+v", result), defaultTruncation))
	}
}

func printActual(t *testing.T, actual interface{}) {
	t.Helper()
	printResult(t, "Actual", actual)
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

// --- Config -----------------------------------------------------------------

func TestSzConfigExport(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzConfigExportResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzConfigExport(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzConfigGetDataSourceRegistry(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzConfigGetDataSourceRegistryResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzConfigGetDataSourceRegistry(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzConfigRegisterDataSource(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzConfigRegisterDataSourceResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzConfigRegisterDataSource(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzConfigUnregisterDataSource(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzConfigUnregisterDataSourceResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		if len(strings.TrimSpace(jsonString)) == 0 {
			continue
		}

		result, err := response.SzConfigUnregisterDataSource(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

// --- ConfigManager ----------------------------------------------------------

func TestSzConfigManagerGetConfigRegistry(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzConfigManagerGetConfigRegistryResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzConfigGetDataSourceRegistry(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

// --- Diagnostic -------------------------------------------------------------

func TestSzDiagnosticCheckRepositoryPerformance(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzDiagnosticCheckRepositoryPerformanceResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzDiagnosticCheckRepositoryPerformance(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzDiagnosticGetFeature(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzDiagnosticGetFeatureResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzDiagnosticGetFeature(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzDiagnosticGetRepositoryInfo(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzDiagnosticGetRepositoryInfoResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzDiagnosticGetRepositoryInfo(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

// --- Engine -----------------------------------------------------------------

// func TestSzEngineAddRecord(test *testing.T) {
// }

func TestSzEngineDeleteRecord(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineDeleteRecordResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzDiagnosticGetFeature(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineFindInterestingEntitiesByEntityID(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineFindInterestingEntitiesByEntityIdResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineFindInterestingEntitiesByEntityID(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineFindInterestingEntitiesByRecordID(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineFindInterestingEntitiesByRecordIdResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineFindInterestingEntitiesByRecordID(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineFindNetworkByEntityID(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineFindNetworkByEntityIdResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineFindNetworkByEntityID(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineFindNetworkByRecordID(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineFindNetworkByRecordIdResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineFindNetworkByRecordID(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineFindPathByEntityID(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineFindPathByEntityIdResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineFindPathByEntityID(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineFindPathByRecordID(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineFindPathByRecordIdResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineFindPathByRecordID(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineGetEntityByEntityID(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineGetEntityByEntityIdResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineGetEntityByEntityID(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineGetEntityByRecordID(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineGetEntityByRecordIdResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineGetEntityByRecordID(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineGetRecordPreview(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineGetRecordPreviewResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineGetRecordPreview(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineGetRecord(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineGetRecordResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineGetRecord(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineGetRedoRecord(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineGetRedoRecordResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineGetRedoRecord(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineGetStats(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineGetStatsResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineGetStats(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineGetVirtualEntityByRecordID(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineGetVirtualEntityByRecordIdResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineGetVirtualEntityByRecordID(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineHowEntityByEntityID(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineHowEntityByEntityIdResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineHowEntityByEntityID(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineProcessRedoRecord(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineProcessRedoRecordResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineProcessRedoRecord(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineReevaluateEntity(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineReevaluateEntityResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineReevaluateEntity(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineReevaluateRecord(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineReevaluateRecordResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineReevaluateRecord(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineSearchByAttributes(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineSearchByAttributesResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineSearchByAttributes(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineWhyEntities(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineWhyEntitiesResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineWhyEntities(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineWhyRecordInEntity(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineWhyRecordInEntityResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineWhyRecordInEntity(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineWhyRecords(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineWhyRecordsResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineWhyRecords(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzEngineWhySearch(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzEngineWhySearchResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzEngineWhySearch(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

// --- Product ----------------------------------------------------------------

func TestSzProductGetLicense(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzProductGetLicenseResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzProductGetLicense(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}

func TestSzProductGetVersion(test *testing.T) {
	test.Parallel()

	ctx := context.TODO()

	scanner, file := createScanner("SzProductGetVersionResponse.jsonl")
	defer closeFile(test, file)

	for scanner.Scan() {
		jsonString := scanner.Text()
		result, err := response.SzProductGetVersion(ctx, jsonString)
		require.NoError(test, err)
		printActual(test, result)
	}

	require.NoError(test, scanner.Err())
}
