#! /usr/bin/env python3

"""
Used to generate senzing/unmarshal_test.go
"""

import json
import logging
import os
import re

from test_cases_from_g2_sdk_json_type_definition import TEST_CASES

IS_DEBUG = False
FINAL_RESULT = {}

OUTPUT_FILE = "./senzing/unmarshal_test.go"

NAME_MAP = {
    "G2diagnosticCheckDbperfResponse": "G2diagnosticCheckDBPerfResponse",
    "G2engineExportConfigAndConfigIdresponse": "G2engineExportConfigAndConfigIDResponse",
    "G2engineFindInterestingEntitiesByEntityIdresponse": "G2engineFindInterestingEntitiesByEntityIDResponse",
    "G2engineFindInterestingEntitiesByRecordIdresponse": "G2engineFindInterestingEntitiesByRecordIDResponse",
    "G2engineFindNetworkByEntityIdresponse": "G2engineFindNetworkByEntityIDResponse",
    "G2engineFindNetworkByEntityIdv2response": "G2engineFindNetworkByEntityIDV2Response",
    "G2engineFindNetworkByRecordIdresponse": "G2engineFindNetworkByRecordIDResponse",
    "G2engineFindNetworkByRecordIdv2response": "G2engineFindNetworkByRecordIDV2Response",
    "G2engineFindPathByEntityIdresponse": "G2engineFindPathByEntityIDResponse",
    "G2engineFindPathByEntityIdv2response": "G2engineFindPathByEntityIDV2Response",
    "G2engineFindPathByRecordIdresponse": "G2engineFindPathByRecordIDResponse",
    "G2engineFindPathByRecordIdv2response": "G2engineFindPathByRecordIDV2Response",
    "G2engineFindPathExcludingByEntityIdresponse": "G2engineFindPathExcludingByEntityIDResponse",
    "G2engineFindPathExcludingByEntityIdv2response": "G2engineFindPathExcludingByEntityIDV2Response",
    "G2engineFindPathExcludingByRecordIdresponse": "G2engineFindPathExcludingByRecordIDResponse",
    "G2engineFindPathExcludingByRecordIdv2response": "G2engineFindPathExcludingByRecordIDV2Response",
    "G2engineFindPathIncludingSourceByEntityIdresponse": "G2engineFindPathIncludingSourceByEntityIDResponse",
    "G2engineFindPathIncludingSourceByEntityIdv2response": "G2engineFindPathIncludingSourceByEntityIDV2Response",
    "G2engineFindPathIncludingSourceByRecordIdresponse": "G2engineFindPathIncludingSourceByRecordIDResponse",
    "G2engineFindPathIncludingSourceByRecordIdv2response": "G2engineFindPathIncludingSourceByRecordIDV2Response",
    "G2engineGetEntityByEntityIdresponse": "G2engineGetEntityByEntityIDResponse",
    "G2engineGetEntityByEntityIdv2response": "G2engineGetEntityByEntityIDV2Response",
    "G2engineGetEntityByRecordIdresponse": "G2engineGetEntityByRecordIDResponse",
    "G2engineGetEntityByRecordIdv2response": "G2engineGetEntityByRecordIDV2Response",
    "G2engineGetRecordV2response": "G2engineGetRecordV2Response",
    "G2engineGetVirtualEntityByRecordIdresponse": "G2engineGetVirtualEntityByRecordIDResponse",
    "G2engineGetVirtualEntityByRecordIdv2response": "G2engineGetVirtualEntityByRecordIDV2Response",
    "G2engineHowEntityByEntityIdresponse": "G2engineHowEntityByEntityIDResponse",
    "G2engineHowEntityByEntityIdv2response": "G2engineHowEntityByEntityIDV2Response",
    "G2engineSearchByAttributesV2response": "G2engineSearchByAttributesV2Response",
    "G2engineSearchByAttributesV3response": "G2engineSearchByAttributesV3Response",
    "G2engineWhyEntitiesV2response": "G2engineWhyEntitiesV2Response",
    "G2engineWhyRecordsV2response": "G2engineWhyRecordsV2Response",
}


# -----------------------------------------------------------------------------
# --- Helpers
# -----------------------------------------------------------------------------


def canonical_json(json_string):
    """Create compact JSON.  No spaces."""
    json_object = json.loads(json_string)
    result = json.dumps(json_object, sort_keys=True, separators=(",", ":"))
    return result


# -----------------------------------------------------------------------------
# --- Main
# -----------------------------------------------------------------------------

# Set up logging.

logging.basicConfig(format="%(asctime)s %(message)s", level=logging.INFO)

logging.info("{0}".format("-" * 80))
logging.info("--- {0} - Begin".format(os.path.basename(__file__)))
logging.info("{0}".format("-" * 80))

# Create multi-line strings for output.

# noqa: E101
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

"""  # noqa: E101, W191

TEST_FUNCTION_TEMPLATE = f"""
	ctx := context.TODO()
	jsonString := `{{json}}`
	result, err := Unmarshal{{struct}}(ctx, jsonString)
	testError(test, ctx, err)
	printActual(test, result)
"""  # noqa: E101,F541,W191

OUTPUT_FOOTER = """
"""

with open(OUTPUT_FILE, "w", encoding="utf-8") as file:
    file.write(OUTPUT_HEADER)
    for senzing_api_class, method_test_cases in TEST_CASES.items():
        for test_case_name, test_case_json in method_test_cases.items():
            file.write("\n\n")
            better_test_case_name = re.sub(
                "[^0-9a-zA-Z]+", "", test_case_name
            ).capitalize()
            canonical_test_case_json = canonical_json(test_case_json)
            canonical_senzing_api_class = NAME_MAP.get(
                senzing_api_class, senzing_api_class
            )
            file.write(
                "func Test{0}{1}(test *testing.T) {{".format(
                    senzing_api_class, better_test_case_name
                )
            )
            file.write(
                TEST_FUNCTION_TEMPLATE.format(
                    json=canonical_test_case_json, struct=canonical_senzing_api_class
                )
            )
            file.write("}")
    file.write(OUTPUT_FOOTER)

# Epilog.

logging.info("{0}".format("-" * 80))
logging.info("--- {0} - End".format(os.path.basename(__file__)))
logging.info("{0}".format("-" * 80))
