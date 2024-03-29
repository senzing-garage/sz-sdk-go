#! /usr/bin/env python3

"""
Used to generate senzing/unmarshal_test.go
"""

import json
import logging
import os

from test_cases_from_g2_sdk_json_type_definition import TEST_CASES

INPUT_FILE = "bin/response-testcases.json"
OUTPUT_FILE = "bin/response-testcases-verified.json"


# -----------------------------------------------------------------------------
# --- Main
# -----------------------------------------------------------------------------

# Set up logging.

logging.basicConfig(format="%(asctime)s %(message)s", level=logging.INFO)

logging.info("{0}".format("-" * 80))
logging.info("--- {0} - Begin".format(os.path.basename(__file__)))
logging.info("{0}".format("-" * 80))

# aggregation_dict is the dictionary to aggregate different sources of test cases.

aggregation_dict = {}
duplicates = 0

# Prime aggregation_dict with existing test cases (response-test-cases.json).

with open(INPUT_FILE, encoding="utf-8") as input_file:
    response_test_cases = json.load(input_file)

for response_class_name, values in response_test_cases.items():
    aggregation_dict[response_class_name] = {"tests": []}

    metadata = values.get("metadata")
    if metadata:
        aggregation_dict[response_class_name]["metadata"] = metadata

    method_test_cases = values.get("tests", {})
    for _, test_case_json in method_test_cases.items():
        test_case_json_string = json.dumps(test_case_json, sort_keys=True)
        if test_case_json_string not in aggregation_dict[response_class_name]["tests"]:
            aggregation_dict[response_class_name]["tests"].append(test_case_json_string)
        else:
            duplicates += 1

# Add second source of test cases.

for response_class_name, values in TEST_CASES.items():
    for _, test_case_json in values.items():
        test_case_dict = json.loads(test_case_json)
        test_case_json_string = json.dumps(test_case_dict, sort_keys=True)
        if test_case_json_string not in aggregation_dict[response_class_name]["tests"]:
            aggregation_dict[response_class_name]["tests"].append(test_case_json_string)
        else:
            duplicates += 1


# Sort results.

for response_class_name, values in aggregation_dict.items():
    aggregation_dict[response_class_name]["tests"].sort()

# Create final result.

result = {}
for response_class_name, values in aggregation_dict.items():
    result[response_class_name] = {"tests": {}}

    metadata = values.get("metadata")
    if metadata:
        result[response_class_name]["metadata"] = metadata

    test_cases = values.get("tests", [])
    test_case_number = 0
    for test_case in test_cases:
        test_case_number += 1
        key = f"test-{test_case_number:03d}"
        result[response_class_name]["tests"][key] = json.loads(test_case)

# Write output file.

with open(OUTPUT_FILE, "w", encoding="utf-8") as file:
    file.write(json.dumps(result, indent=4))

# Epilog.

logging.info("{0}".format("-" * 80))
logging.info("--- {0} - End".format(os.path.basename(__file__)))
logging.info("---   Duplicates: {0}".format(duplicates))
logging.info("{0}".format("-" * 80))
