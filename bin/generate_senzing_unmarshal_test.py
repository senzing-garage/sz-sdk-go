#! /usr/bin/env python3

"""
Used to generate senzing/unmarshal_test.go
"""

import logging
import json
import os
import re
from test_cases_from_g2_sdk_json_type_definition import TEST_CASES

IS_DEBUG = False
FINAL_RESULT = {}

OUTPUT_FILE = "./senzing/unmarshal_test.go"

NAME_MAP = {
    "DiagnosticCheckDbperfResponse": "DiagnosticCheckDBPerfResponse",
    "DiagnosticGetDbinfoResponse": "DiagnosticGetDBInfoResponse",
    "EngineAddRecordWithInfoWithReturnedRecordIdresponse": "EngineAddRecordWithInfoWithReturnedRecordIDResponse",
    "EngineExportConfigAndConfigIdresponse": "EngineExportConfigAndConfigIDResponse",
    "EngineFindInterestingEntitiesByEntityIdresponse": "EngineFindInterestingEntitiesByEntityIDResponse",
    "EngineFindInterestingEntitiesByRecordIdresponse": "EngineFindInterestingEntitiesByRecordIDResponse",
    "EngineFindNetworkByEntityIdresponse": "EngineFindNetworkByEntityIDResponse",
    "EngineFindNetworkByEntityIdv2response": "EngineFindNetworkByEntityIDV2Response",
    "EngineFindNetworkByRecordIdresponse": "EngineFindNetworkByRecordIDResponse",
    "EngineFindNetworkByRecordIdv2response": "EngineFindNetworkByRecordIDV2Response",
    "EngineFindPathByEntityIdresponse": "EngineFindPathByEntityIDResponse",
    "EngineFindPathByEntityIdv2response": "EngineFindPathByEntityIDV2Response",
    "EngineFindPathByRecordIdresponse": "EngineFindPathByRecordIDResponse",
    "EngineFindPathByRecordIdv2response": "EngineFindPathByRecordIDV2Response",
    "EngineFindPathExcludingByEntityIdresponse": "EngineFindPathExcludingByEntityIDResponse",
    "EngineFindPathExcludingByEntityIdv2response": "EngineFindPathExcludingByEntityIDV2Response",
    "EngineFindPathExcludingByRecordIdresponse": "EngineFindPathExcludingByRecordIDResponse",
    "EngineFindPathExcludingByRecordIdv2response": "EngineFindPathExcludingByRecordIDV2Response",
    "EngineFindPathIncludingSourceByEntityIdresponse": "EngineFindPathIncludingSourceByEntityIDResponse",
    "EngineFindPathIncludingSourceByEntityIdv2response": "EngineFindPathIncludingSourceByEntityIDV2Response",
    "EngineFindPathIncludingSourceByRecordIdresponse": "EngineFindPathIncludingSourceByRecordIDResponse",
    "EngineFindPathIncludingSourceByRecordIdv2response": "EngineFindPathIncludingSourceByRecordIDV2Response",
    "EngineGetEntityByEntityIdresponse": "EngineGetEntityByEntityIDResponse",
    "EngineGetEntityByEntityIdv2response": "EngineGetEntityByEntityIDV2Response",
    "EngineGetEntityByRecordIdresponse": "EngineGetEntityByRecordIDResponse",
    "EngineGetEntityByRecordIdv2response": "EngineGetEntityByRecordIDV2Response",
    "EngineGetVirtualEntityByRecordIdresponse": "EngineGetVirtualEntityByRecordIDResponse",
    "EngineGetVirtualEntityByRecordIdv2response": "EngineGetVirtualEntityByRecordIDV2Response",
    "EngineHowEntityByEntityIdresponse": "EngineHowEntityByEntityIDResponse",
    "EngineHowEntityByEntityIdv2response": "EngineHowEntityByEntityIDV2Response",
    "EngineSearchByAttributesV2response": "EngineSearchByAttributesV2Response",
    "EngineWhyEntitiesV2response": "EngineWhyEntitiesV2Response",
    "EngineWhyEntityByEntityIdresponse": "EngineWhyEntityByEntityIDResponse",
    "EngineWhyEntityByEntityIdv2response": "EngineWhyEntityByEntityIDV2Response",
    "EngineWhyEntityByRecordIdresponse": "EngineWhyEntityByRecordIDResponse",
    "EngineWhyEntityByRecordIdv2response": "EngineWhyEntityByRecordIDV2Response",
    "EngineWhyRecordsV2response": "EngineWhyRecordsV2Response",
    "DiagnosticFindEntitiesByFeatureIdsResponse": "DiagnosticFindEntitiesByFeatureIDsResponse",
    "EngineGetRecordV2response": "EngineGetRecordV2Response",
    "EngineSearchByAttributesV3response": "EngineSearchByAttributesV3Response"
}


# -----------------------------------------------------------------------------
# --- Helpers
# -----------------------------------------------------------------------------

def canonical_json(json_string):
    """Create compact JSON.  No spaces."""
    json_object = json.loads(json_string)
    result = json.dumps(json_object, sort_keys=True, separators=(',', ':'))
    return result


# -----------------------------------------------------------------------------
# --- Main
# -----------------------------------------------------------------------------

# Set up logging.

logging.basicConfig(format='%(asctime)s %(message)s', level=logging.INFO)

logging.info("{0}".format("-" * 80))
logging.info("--- {0} - Begin".format(os.path.basename(__file__)))
logging.info("{0}".format("-" * 80))

# Create multi-line strings for output.

OUTPUT_HEADER = """// Code generated by generate_senzing_unmarshal_test.py. DO NOT EDIT.

package senzing

import (
	"context"
	"fmt"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/stretchr/testify/assert"
)

const (
	defaultTruncation = 127
	printResults      = false
)

var ()

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func truncate(aString string, length int) string {
	return truncator.Truncate(aString, length, "...", truncator.PositionEnd)
}

func printResult(test *testing.T, title string, result interface{}) {
	if printResults {
		test.Logf("%s: %+v", title, truncate(fmt.Sprintf("%+v", result), defaultTruncation))
	}
}

func printActual(test *testing.T, actual interface{}) {
	printResult(test, "Actual", actual)
}

func testError(test *testing.T, ctx context.Context, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

"""

TEST_FUNCTION_TEMPLATE = f"""
    ctx := context.TODO()
	jsonString := `{{json}}`
    result, err := Unmarshal{{struct}}(ctx, jsonString)
	testError(test, ctx, err)
    printActual(test, result)
"""

OUTPUT_FOOTER = """
"""

with open(OUTPUT_FILE, "w", encoding="utf-8") as file:
    file.write(OUTPUT_HEADER)
    for senzing_api_class, method_test_cases in TEST_CASES.items():
        for test_case_name, test_case_json in method_test_cases.items():
            better_test_case_name = re.sub('[^0-9a-zA-Z]+', '', test_case_name).capitalize()
            canonical_test_case_json = canonical_json(test_case_json)
            canonical_senzing_api_class = NAME_MAP.get(senzing_api_class, senzing_api_class)
            file.write("func Test{0}{1}(test *testing.T) {{".format(senzing_api_class, better_test_case_name))
            file.write(TEST_FUNCTION_TEMPLATE.format(json=canonical_test_case_json, struct=canonical_senzing_api_class))
            file.write("}\n\n")
    file.write(OUTPUT_FOOTER)

# Epilog.

logging.info("{0}".format("-" * 80))
logging.info("--- {0} - End".format(os.path.basename(__file__)))
logging.info("{0}".format("-" * 80))