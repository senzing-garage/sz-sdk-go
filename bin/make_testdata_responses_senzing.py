#! /usr/bin/env python3

# pylint: disable=duplicate-code

"""
For more information, visit https://jsontypedef.com/docs/python-codegen/
"""

import json
import logging
import os
import pathlib
import random

from senzing import SzAbstractFactory, SzEngineFlags, SzError
from senzing_core import SzAbstractFactoryCore

# Logging

logging.basicConfig(
    level=logging.INFO, format="%(asctime)s - %(levelname)s - %(message)s"
)
logger = logging.getLogger(__name__)

# Global variables.

CURRENT_PATH = pathlib.Path(__file__).parent.resolve()
TESTDATA_DIRECTORY = os.path.abspath(f"{CURRENT_PATH}/../testdata")
OUTPUT_DIRECTORY = f"{TESTDATA_DIRECTORY}/responses_senzing"
TRUTHSETS_DIRECTORY = f"{TESTDATA_DIRECTORY}/truthsets"

DEBUG = 0
HR_START = ">" * 80
HR_STOP = "<" * 80
LOADED_ENTITY_IDS = []
LOADED_RECORD_KEYS = []

FLAGS = [
    SzEngineFlags.SZ_WITH_INFO,  # 1
    SzEngineFlags.SZ_ADD_RECORD_DEFAULT_FLAGS,  # 2
    SzEngineFlags.SZ_DELETE_RECORD_DEFAULT_FLAGS,  # 3
    SzEngineFlags.SZ_ENTITY_BRIEF_DEFAULT_FLAGS,  # 4
    SzEngineFlags.SZ_ENTITY_CORE_FLAGS,  # 5
    SzEngineFlags.SZ_ENTITY_DEFAULT_FLAGS,  # 6
    SzEngineFlags.SZ_ENTITY_INCLUDE_ALL_FEATURES,  # 7
    SzEngineFlags.SZ_ENTITY_INCLUDE_ALL_RELATIONS,  # 8
    SzEngineFlags.SZ_ENTITY_INCLUDE_DISCLOSED_RELATIONS,  # 9
    SzEngineFlags.SZ_ENTITY_INCLUDE_ENTITY_NAME,  # 10
    SzEngineFlags.SZ_ENTITY_INCLUDE_FEATURE_STATS,  # 11
    SzEngineFlags.SZ_ENTITY_INCLUDE_INTERNAL_FEATURES,  # 12
    SzEngineFlags.SZ_ENTITY_INCLUDE_NAME_ONLY_RELATIONS,  # 13
    SzEngineFlags.SZ_ENTITY_INCLUDE_POSSIBLY_RELATED_RELATIONS,  # 14
    SzEngineFlags.SZ_ENTITY_INCLUDE_POSSIBLY_SAME_RELATIONS,  # 15
    SzEngineFlags.SZ_ENTITY_INCLUDE_RECORD_DATA,  # 16
    SzEngineFlags.SZ_ENTITY_INCLUDE_RECORD_DATES,  # 17
    SzEngineFlags.SZ_ENTITY_INCLUDE_RECORD_FEATURE_DETAILS,  # 18
    SzEngineFlags.SZ_ENTITY_INCLUDE_RECORD_FEATURE_STATS,  # 19
    SzEngineFlags.SZ_ENTITY_INCLUDE_RECORD_FEATURES,  # 20
    SzEngineFlags.SZ_ENTITY_INCLUDE_RECORD_JSON_DATA,  # 21
    SzEngineFlags.SZ_ENTITY_INCLUDE_RECORD_MATCHING_INFO,  # 22
    SzEngineFlags.SZ_ENTITY_INCLUDE_RECORD_SUMMARY,  # 23
    SzEngineFlags.SZ_ENTITY_INCLUDE_RECORD_TYPES,  # 24
    SzEngineFlags.SZ_ENTITY_INCLUDE_RECORD_UNMAPPED_DATA,  # 25
    SzEngineFlags.SZ_ENTITY_INCLUDE_RELATED_ENTITY_NAME,  # 26
    SzEngineFlags.SZ_ENTITY_INCLUDE_RELATED_MATCHING_INFO,  # 27
    SzEngineFlags.SZ_ENTITY_INCLUDE_RELATED_RECORD_DATA,  # 28
    SzEngineFlags.SZ_ENTITY_INCLUDE_RELATED_RECORD_SUMMARY,  # 29
    SzEngineFlags.SZ_ENTITY_INCLUDE_RELATED_RECORD_TYPES,  # 30
    SzEngineFlags.SZ_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES,  # 31
    SzEngineFlags.SZ_EXPORT_DEFAULT_FLAGS,  # 32
    SzEngineFlags.SZ_EXPORT_INCLUDE_ALL_ENTITIES,  # 33
    SzEngineFlags.SZ_EXPORT_INCLUDE_ALL_HAVING_RELATIONSHIPS,  # 34
    SzEngineFlags.SZ_EXPORT_INCLUDE_DISCLOSED,  # 35
    SzEngineFlags.SZ_EXPORT_INCLUDE_MULTI_RECORD_ENTITIES,  # 36
    SzEngineFlags.SZ_EXPORT_INCLUDE_NAME_ONLY,  # 37
    SzEngineFlags.SZ_EXPORT_INCLUDE_POSSIBLY_RELATED,  # 38
    SzEngineFlags.SZ_EXPORT_INCLUDE_POSSIBLY_SAME,  # 39
    SzEngineFlags.SZ_EXPORT_INCLUDE_SINGLE_RECORD_ENTITIES,  # 40
    SzEngineFlags.SZ_FIND_INTERESTING_ENTITIES_DEFAULT_FLAGS,  # 41
    SzEngineFlags.SZ_FIND_NETWORK_DEFAULT_FLAGS,  # 42
    SzEngineFlags.SZ_FIND_NETWORK_INCLUDE_MATCHING_INFO,  # 43
    SzEngineFlags.SZ_FIND_PATH_DEFAULT_FLAGS,  # 44
    SzEngineFlags.SZ_FIND_PATH_INCLUDE_MATCHING_INFO,  # 45
    SzEngineFlags.SZ_FIND_PATH_STRICT_AVOID,  # 46
    SzEngineFlags.SZ_HOW_ENTITY_DEFAULT_FLAGS,  # 47
    SzEngineFlags.SZ_INCLUDE_FEATURE_SCORES,  # 48
    SzEngineFlags.SZ_INCLUDE_MATCH_KEY_DETAILS,  # 49
    SzEngineFlags.SZ_NO_FLAGS,  # 50
    SzEngineFlags.SZ_RECORD_DEFAULT_FLAGS,  # 51
    SzEngineFlags.SZ_RECORD_PREVIEW_DEFAULT_FLAGS,  # 52
    SzEngineFlags.SZ_REEVALUATE_ENTITY_DEFAULT_FLAGS,  # 53
    SzEngineFlags.SZ_REEVALUATE_RECORD_DEFAULT_FLAGS,  # 54
    SzEngineFlags.SZ_SEARCH_BY_ATTRIBUTES_ALL,  # 55
    SzEngineFlags.SZ_SEARCH_BY_ATTRIBUTES_DEFAULT_FLAGS,  # 56
    SzEngineFlags.SZ_SEARCH_BY_ATTRIBUTES_MINIMAL_ALL,  # 57
    SzEngineFlags.SZ_SEARCH_BY_ATTRIBUTES_MINIMAL_STRONG,  # 58
    SzEngineFlags.SZ_SEARCH_BY_ATTRIBUTES_STRONG,  # 59
    SzEngineFlags.SZ_SEARCH_INCLUDE_ALL_CANDIDATES,  # 60
    SzEngineFlags.SZ_SEARCH_INCLUDE_ALL_ENTITIES,  # 61
    SzEngineFlags.SZ_SEARCH_INCLUDE_NAME_ONLY,  # 62
    SzEngineFlags.SZ_SEARCH_INCLUDE_POSSIBLY_RELATED,  # 63
    SzEngineFlags.SZ_SEARCH_INCLUDE_POSSIBLY_SAME,  # 64
    SzEngineFlags.SZ_SEARCH_INCLUDE_REQUEST_DETAILS,  # 65
    SzEngineFlags.SZ_SEARCH_INCLUDE_REQUEST,  # 66
    SzEngineFlags.SZ_SEARCH_INCLUDE_RESOLVED,  # 67
    SzEngineFlags.SZ_SEARCH_INCLUDE_STATS,  # 68
    SzEngineFlags.SZ_VIRTUAL_ENTITY_DEFAULT_FLAGS,  # 69
    SzEngineFlags.SZ_WHY_ENTITIES_DEFAULT_FLAGS,  # 70
    SzEngineFlags.SZ_WHY_RECORD_IN_ENTITY_DEFAULT_FLAGS,  # 71
    SzEngineFlags.SZ_WHY_RECORDS_DEFAULT_FLAGS,  # 72
    SzEngineFlags.SZ_WHY_SEARCH_DEFAULT_FLAGS,  # 73
]

FLAGS_LEN = len(FLAGS)

SEARCH_RECORDS = [
    {
        "NAME_FULL": "Susan Moony",
        "DATE_OF_BIRTH": "15/6/1998",
        "SSN_NUMBER": "521212123",
    },
    {
        "NAME_FIRST": "Robert",
        "NAME_LAST": "Smith",
        "ADDR_FULL": "123 Main Street Las Vegas NV 89132",
    },
    {
        "NAME_FIRST": "Makio",
        "NAME_LAST": "Yamanaka",
        "ADDR_FULL": "787 Rotary Drive Rotorville FL 78720",
    },
]


# -----------------------------------------------------------------------------
# Add records.
# -----------------------------------------------------------------------------


def add_records(sz_abstract_factory: SzAbstractFactory):
    """Add records to the Senzing repository."""

    # global LOADED_RECORD_KEYS

    debug_records = [  # Format: (data_source, record_id)
        ("CUSTOMER", "0"),
    ]

    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineAddRecordResponse"

    # Add records.

    filenames = [
        "customers.jsonl",
        "reference.jsonl",
        "watchlist.jsonl",
    ]

    test_count = 0
    for filename in filenames:
        file_path = f"{TRUTHSETS_DIRECTORY}/{filename}"
        with open(file_path, "r", encoding="utf-8") as input_file:
            for line in input_file:
                line_as_dict = json.loads(line)
                data_source = line_as_dict.get("DATA_SOURCE")
                record_id = line_as_dict.get("RECORD_ID")

                LOADED_RECORD_KEYS.append(
                    {
                        "data_source": data_source,
                        "record_id": record_id,
                    }
                )
                test_name = f"Add record: {filename}/{data_source}/{record_id}"
                response = sz_engine.add_record(
                    data_source, record_id, line, SzEngineFlags.SZ_WITH_INFO
                )
                if not response:
                    continue
                set_debug((data_source, record_id), debug_records)
                debug(1, f"{HR_START}\n{test_name}; Response:\n{response}\n{HR_STOP}\n")
                test_count += 1
    if not test_count:
        output(0, f"No tests performed for {title}")


def register_datasources(sz_abstract_factory: SzAbstractFactory):
    """Add Data Sources to the Senzing repository."""

    sz_config_manager = sz_abstract_factory.create_configmanager()

    # Register datasources.  TODO: fix underlying database.

    current_config_id = sz_config_manager.get_default_config_id()
    sz_config = sz_config_manager.create_config_from_config_id(current_config_id)

    for data_source in ("CUSTOMERS", "REFERENCE", "WATCHLIST"):
        sz_config.register_data_source(data_source)

    new_config = sz_config.export()
    new_config_id = sz_config_manager.register_config(
        new_config, "sz-sdk-json-type-definition"
    )
    sz_config_manager.replace_default_config_id(current_config_id, new_config_id)
    sz_abstract_factory.reinitialize(new_config_id)


# -----------------------------------------------------------------------------
# Compare
# -----------------------------------------------------------------------------


def compare(sz_abstract_factory: SzAbstractFactory):
    """Aggregate all compare functions."""
    compare_find_interesting_entities_by_entity_id(sz_abstract_factory)
    compare_find_interesting_entities_by_record_id(sz_abstract_factory)
    compare_find_network_by_entity_id(sz_abstract_factory)
    compare_find_network_by_record_id(sz_abstract_factory)
    compare_find_path_by_entity_id(sz_abstract_factory)
    compare_find_path_by_record_id(sz_abstract_factory)
    compare_get_entity_by_entity_id(sz_abstract_factory)
    compare_get_entity_by_record_id(sz_abstract_factory)
    compare_get_feature(sz_abstract_factory)
    compare_get_record_preview(sz_abstract_factory)
    compare_get_record(sz_abstract_factory)
    compare_get_virtual_entity_by_record_id(sz_abstract_factory)
    compare_how_entity_by_entity_id(sz_abstract_factory)
    compare_redo(sz_abstract_factory)
    compare_reevaluate_entity(sz_abstract_factory)
    compare_reevaluate_record(sz_abstract_factory)
    compare_search_by_attributes(sz_abstract_factory)
    compare_static_method_signatures(sz_abstract_factory)
    compare_why_entities(sz_abstract_factory)
    compare_why_record_in_entity(sz_abstract_factory)
    compare_why_records(sz_abstract_factory)
    compare_why_search(sz_abstract_factory)


# -----------------------------------------------------------------------------
# FindInterestingEntities
# -----------------------------------------------------------------------------


def compare_find_interesting_entities_by_entity_id(
    sz_abstract_factory: SzAbstractFactory,
):
    """Compare find_interesting_entities_by_entity_id."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineFindInterestingEntitiesByEntityIdResponse"
    test_cases = ["{}"]
    for entity_id in LOADED_ENTITY_IDS:
        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            response = sz_engine.find_interesting_entities_by_entity_id(entity_id, flag)
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


def compare_find_interesting_entities_by_record_id(
    sz_abstract_factory: SzAbstractFactory,
):
    """Compare find_interesting_entities_by_record_id."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineFindInterestingEntitiesByRecordIdResponse"
    test_cases = ["{}"]
    for record in LOADED_RECORD_KEYS:
        data_source = record.get("data_source", "")
        record_id = record.get("record_id", "")
        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            response = sz_engine.find_interesting_entities_by_record_id(
                data_source, record_id, flag
            )
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# FindNetwork
# -----------------------------------------------------------------------------


def compare_find_network_by_entity_id(sz_abstract_factory: SzAbstractFactory):
    """Compare find_network_by_entity_id."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineFindNetworkByEntityIdResponse"
    test_cases = ["{}"]
    max_degrees = 5
    build_out_degrees = 5
    build_out_max_entities = 5

    for entity_id in LOADED_ENTITY_IDS:

        # Randomize entity_ids.

        entity_ids = [entity_id]
        for _ in range(random.randint(3, 10)):
            new_entity_id = LOADED_ENTITY_IDS[random.randint(0, FLAGS_LEN - 1)]
            if new_entity_id not in entity_ids:
                entity_ids.append(new_entity_id)

        # Compare.

        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            response = sz_engine.find_network_by_entity_id(
                entity_ids,
                max_degrees,
                build_out_degrees,
                build_out_max_entities,
                flag,
            )
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


def compare_find_network_by_record_id(sz_abstract_factory: SzAbstractFactory):
    """Compare find_network_by_record_id."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineFindNetworkByRecordIdResponse"
    test_cases = ["{}"]
    max_degrees = 5
    build_out_degrees = 5
    build_out_max_entities = 5

    for record in LOADED_RECORD_KEYS:

        data_source = record.get("data_source", "")
        record_id = record.get("record_id", "")

        # Randomize record_keys.

        record_keys = [(data_source, record_id)]
        for _ in range(random.randint(3, 10)):
            record_key_dict = LOADED_RECORD_KEYS[random.randint(0, FLAGS_LEN - 1)]
            record_key = (
                record_key_dict.get("data_source", ""),
                record_key_dict.get("record_id", ""),
            )
            if record_key not in record_keys:
                record_keys.append(record_key)

        # Compare.

        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            response = sz_engine.find_network_by_record_id(
                record_keys,
                max_degrees,
                build_out_degrees,
                build_out_max_entities,
                flag,
            )
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# FindPath
# -----------------------------------------------------------------------------


def compare_find_path_by_entity_id(sz_abstract_factory: SzAbstractFactory):
    """Compare find_path_by_entity_id."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineFindPathByEntityIdResponse"
    test_cases = ["{}"]
    max_degrees = 5
    avoid_entity_ids = None
    required_data_sources = None

    for entity_id in LOADED_ENTITY_IDS:
        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            end_entity_id = LOADED_ENTITY_IDS[random.randint(0, FLAGS_LEN - 1)]
            response = sz_engine.find_path_by_entity_id(
                entity_id,
                end_entity_id,
                max_degrees,
                avoid_entity_ids,
                required_data_sources,
                flag,
            )
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


def compare_find_path_by_record_id(sz_abstract_factory: SzAbstractFactory):
    """Compare find_path_by_record_id."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineFindPathByRecordIdResponse"
    test_cases = ["{}"]
    max_degrees = 5
    avoid_record_keys = None
    required_data_sources = None

    for record in LOADED_RECORD_KEYS:
        data_source = record.get("data_source", "")
        record_id = record.get("record_id", "")
        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            end_record = LOADED_RECORD_KEYS[random.randint(0, FLAGS_LEN - 1)]
            response = sz_engine.find_path_by_record_id(
                data_source,
                record_id,
                end_record.get("data_source", ""),
                end_record.get("record_id", ""),
                max_degrees,
                avoid_record_keys,
                required_data_sources,
                flag,
            )
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# GetEntity
# -----------------------------------------------------------------------------


def compare_get_entity_by_entity_id(sz_abstract_factory: SzAbstractFactory):
    """Compare get_entity_by_entity_id."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineGetEntityByEntityIdResponse"
    test_cases = ["{}"]

    for entity_id in LOADED_ENTITY_IDS:
        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            response = sz_engine.get_entity_by_entity_id(entity_id, flag)
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


def compare_get_entity_by_record_id(sz_abstract_factory: SzAbstractFactory):
    """Compare get_entity_by_record_id."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineGetEntityByRecordIdResponse"
    test_cases = ["{}"]

    for record in LOADED_RECORD_KEYS:
        data_source = record.get("data_source", "")
        record_id = record.get("record_id", "")

        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            response = sz_engine.get_entity_by_record_id(data_source, record_id, flag)
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# GetFeature
# -----------------------------------------------------------------------------


def compare_get_feature(sz_abstract_factory: SzAbstractFactory):
    """Compare get_feature."""
    sz_diagnostic = sz_abstract_factory.create_diagnostic()
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzDiagnosticGetFeatureResponse"
    test_cases = ["{}"]

    # Extract feature IDs.

    feature_ids = []
    for entity_id in LOADED_ENTITY_IDS:
        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            response = sz_engine.get_entity_by_entity_id(entity_id, flag)
            response_dict = json.loads(response)
            feature = response_dict.get("RESOLVED_ENTITY", {}).get("FEATURES", {})
            for feature_values in feature.values():
                for feature_value in feature_values:
                    feature_id = feature_value.get("LIB_FEAT_ID", None)
                    if feature_id:
                        if feature_id not in feature_ids:
                            feature_ids.append(feature_id)

    # Test get_feature.

    for feature_id in feature_ids:
        response = sz_diagnostic.get_feature(feature_id)
        if response not in test_cases:
            test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# GetRecord
# -----------------------------------------------------------------------------


def compare_get_record(sz_abstract_factory: SzAbstractFactory):
    """Compare get_record."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineGetRecordResponse"
    test_cases = ["{}"]

    for record in LOADED_RECORD_KEYS:
        data_source = record.get("data_source", "")
        record_id = record.get("record_id", "")

        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            response = sz_engine.get_record(data_source, record_id, flag)
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# GetRecordPreview
# -----------------------------------------------------------------------------


def compare_get_record_preview(sz_abstract_factory: SzAbstractFactory):
    """Compare get_record_preview."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineGetRecordPreviewResponse"
    test_cases = ["{}"]

    for record_dict in LOADED_RECORD_KEYS:

        # Fetch actual record.

        data_source = record_dict.get("data_source", "")
        record_id = record_dict.get("record_id", "")
        record_str = sz_engine.get_record(data_source, record_id)
        record_dict = json.loads(record_str)
        record_json_dict = record_dict.get("JSON_DATA")
        record_json = json.dumps(record_json_dict)

        # Preview record with various flags.

        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            response = sz_engine.get_record_preview(record_json, flag)
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# GetVirtualEntity
# -----------------------------------------------------------------------------


def compare_get_virtual_entity_by_record_id(sz_abstract_factory: SzAbstractFactory):
    """Compare get_virtual_entity_by_record_id."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineGetVirtualEntityByRecordIdResponse"
    test_cases = ["{}"]

    for record in LOADED_RECORD_KEYS:
        data_source = record.get("data_source", "")
        record_id = record.get("record_id", "")

        # Randomize record_keys.

        record_keys = [(data_source, record_id)]
        for _ in range(random.randint(3, 10)):
            record_key_dict = LOADED_RECORD_KEYS[random.randint(0, FLAGS_LEN - 1)]
            record_key = (
                record_key_dict.get("data_source", ""),
                record_key_dict.get("record_id", ""),
            )
            if record_key not in record_keys:
                record_keys.append(record_key)

        # Compare.

        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            response = sz_engine.get_virtual_entity_by_record_id(record_keys, flag)
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# HowEntity
# -----------------------------------------------------------------------------


def compare_how_entity_by_entity_id(sz_abstract_factory: SzAbstractFactory):
    """Compare how_entity_by_entity_id."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineHowEntityByEntityIdResponse"
    test_cases = ["{}"]

    for entity_id in LOADED_ENTITY_IDS:
        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            response = sz_engine.how_entity_by_entity_id(entity_id, flag)
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# Redo
# -----------------------------------------------------------------------------


def compare_redo(sz_abstract_factory: SzAbstractFactory):
    """Compare get_redo_record."""
    sz_engine = sz_abstract_factory.create_engine()

    # Get all redo records.

    title = "SzEngineGetRedoRecordResponse"
    test_cases = ["{}"]

    redo_records = []
    redo_record_count = 0
    while True:
        redo_record_count += 1
        response = sz_engine.get_redo_record()
        if not response:
            break
        redo_records.append(response)
        if response not in test_cases:
            test_cases.append(response)
    output_file(title, test_cases)

    # Process redo records.

    title = "SzEngineProcessRedoRecordResponse"

    redo_record_count = 0
    for redo_record in redo_records:
        redo_record_count += 1
        response = sz_engine.process_redo_record(
            redo_record, SzEngineFlags.SZ_WITH_INFO
        )
        if response not in test_cases:
            test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# Reevaluate
# -----------------------------------------------------------------------------


def compare_reevaluate_entity(sz_abstract_factory: SzAbstractFactory):
    """Compare reevaluate_entity."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineReevaluateEntityResponse"
    test_cases = ["{}"]

    for entity_id in LOADED_ENTITY_IDS:
        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            flag = flag | SzEngineFlags.SZ_WITH_INFO
            response = sz_engine.reevaluate_entity(entity_id, flag)
            if not response:
                continue
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


def compare_reevaluate_record(sz_abstract_factory: SzAbstractFactory):
    """Compare reevaluate_record."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineReevaluateRecordResponse"
    test_cases = ["{}"]

    for record in LOADED_RECORD_KEYS:
        data_source = record.get("data_source", "")
        record_id = record.get("record_id", "")
        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            flag = flag | SzEngineFlags.SZ_WITH_INFO
            response = sz_engine.reevaluate_record(data_source, record_id, flag)
            if not response:
                continue
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# SearchByAttributes
# -----------------------------------------------------------------------------


def compare_search_by_attributes(sz_abstract_factory: SzAbstractFactory):
    """Compare search_by_attributes."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineSearchByAttributesResponse"
    test_cases = ["{}"]
    search_profile = ""

    search_record_count = 0
    for search_record in SEARCH_RECORDS:
        search_record_count += 1
        attributes = json.dumps(search_record)
        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            response = sz_engine.search_by_attributes(attributes, flag, search_profile)
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# Static method signature calls
# -----------------------------------------------------------------------------


def compare_static_method_signatures(sz_abstract_factory: SzAbstractFactory):
    """Compare methods without variable parameters."""

    sz_config_manager = sz_abstract_factory.create_configmanager()
    sz_config = sz_config_manager.create_config_from_template()
    sz_diagnostic = sz_abstract_factory.create_diagnostic()
    sz_engine = sz_abstract_factory.create_engine()
    sz_product = sz_abstract_factory.create_product()

    # For linter

    _ = sz_config
    _ = sz_engine
    _ = sz_diagnostic
    _ = sz_product

    # Define testcases

    testcases = [
        {
            "testcase": "sz_config.export()",
            "response": "SzConfigExportResponse",
        },
        {
            "testcase": "sz_config.get_data_source_registry()",
            "response": "SzConfigGetDataSourceRegistryResponse",
        },
        {
            "testcase": 'sz_config.register_data_source("A_DATASOURCE_NAME")',
            "response": "SzConfigRegisterDataSourceResponse",
        },
        {
            "testcase": 'sz_config.unregister_data_source("A_DATASOURCE_NAME")',
            "response": "SzConfigUnregisterDataSourceResponse",
        },
        {
            "testcase": "sz_config_manager.get_config_registry()",
            "response": "SzConfigManagerGetConfigRegistryResponse",
        },
        {
            "testcase": "sz_diagnostic.check_repository_performance(2)",
            "response": "SzDiagnosticCheckRepositoryPerformanceResponse",
        },
        {
            "testcase": "sz_diagnostic.get_repository_info()",
            "response": "SzDiagnosticGetRepositoryInfoResponse",
        },
        # {
        #     "testcase": "sz_diagnostic.get_feature(1)",
        #     "response": "SzDiagnosticGetFeatureResponse",
        # },
        {
            "testcase": "sz_engine.get_stats()",
            "response": "SzEngineGetStatsResponse",
        },
        {
            "testcase": "sz_product.get_license()",
            "response": "SzProductGetLicenseResponse",
        },
        {
            "testcase": "sz_product.get_version()",
            "response": "SzProductGetVersionResponse",
        },
    ]

    for testcase in testcases:
        response = eval(testcase.get("testcase", ""))
        output_file(f"{testcase.get('response')}", ["{}", response])


# -----------------------------------------------------------------------------
# WhyEntities
# -----------------------------------------------------------------------------


def compare_why_entities(sz_abstract_factory: SzAbstractFactory):
    """Compare why_entities."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineWhyEntitiesResponse"
    test_cases = ["{}"]

    for entity_id in LOADED_ENTITY_IDS:
        entity_id_2 = LOADED_ENTITY_IDS[random.randint(0, FLAGS_LEN - 1)]
        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            response = sz_engine.why_entities(entity_id, entity_id_2, flag)
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# WhyRecordInEntity
# -----------------------------------------------------------------------------


def compare_why_record_in_entity(sz_abstract_factory: SzAbstractFactory):
    """Compare why_record_in_entity."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineWhyRecordInEntityResponse"
    test_cases = ["{}"]

    for record in LOADED_RECORD_KEYS:
        data_source = record.get("data_source", "")
        record_id = record.get("record_id", "")
        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            response = sz_engine.why_record_in_entity(data_source, record_id, flag)
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# WhyRecords
# -----------------------------------------------------------------------------


def compare_why_records(sz_abstract_factory: SzAbstractFactory):
    """Compare why_records."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineWhyRecordsResponse"
    test_cases = ["{}"]

    for record in LOADED_RECORD_KEYS:

        data_source = record.get("data_source", "")
        record_id = record.get("record_id", "")

        # Randomize record.

        record_key_dict = LOADED_RECORD_KEYS[random.randint(0, FLAGS_LEN - 1)]
        data_source_2 = record_key_dict.get("data_source", "")
        record_id_2 = record_key_dict.get("record_id", "")

        # Compare.

        flag_count = 0
        for flag in FLAGS:
            flag_count += 1
            response = sz_engine.why_records(
                data_source, record_id, data_source_2, record_id_2, flag
            )
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# WhySearch
# -----------------------------------------------------------------------------


def compare_why_search(sz_abstract_factory: SzAbstractFactory):
    """Compare why_search."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineWhySearchResponse"
    test_cases = ["{}"]
    search_profile = ""

    for entity_id in LOADED_ENTITY_IDS:
        search_record_count = 0
        for search_record in SEARCH_RECORDS:
            search_record_count += 1
            attributes = json.dumps(search_record)
            flag_count = 0
            for flag in FLAGS:
                flag_count += 1
                response = sz_engine.why_search(
                    attributes, entity_id, flag, search_profile
                )
            if response not in test_cases:
                test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# Delete records
# -----------------------------------------------------------------------------


def delete_records(sz_abstract_factory: SzAbstractFactory):
    """Compare delete_record."""
    sz_engine = sz_abstract_factory.create_engine()
    title = "SzEngineDeleteRecordResponse"
    test_cases = ["{}"]
    record_count = 0
    for record in LOADED_RECORD_KEYS:
        record_count += 1
        data_source = record.get("data_source", "")
        record_id = record.get("record_id", "")
        response = sz_engine.delete_record(
            data_source, record_id, SzEngineFlags.SZ_WITH_INFO
        )
        if response not in test_cases:
            test_cases.append(response)
    output_file(title, test_cases)


# -----------------------------------------------------------------------------
# Utility functions
# -----------------------------------------------------------------------------


def create_sz_abstract_factory() -> SzAbstractFactory:
    """Create an SzAbstractFactory."""
    instance_name = "Example"
    settings = {
        "PIPELINE": {
            "CONFIGPATH": "/etc/opt/senzing",
            "RESOURCEPATH": "/opt/senzing/er/resources",
            "SUPPORTPATH": "/opt/senzing/data",
        },
        "SQL": {"CONNECTION": "sqlite3://na:na@/tmp/sqlite/G2C.db"},
    }

    try:
        sz_abstract_factory = SzAbstractFactoryCore(instance_name, settings)
    except SzError as err:
        logger.error("%s", err)

    return sz_abstract_factory


def debug(level, message):
    """If appropriate, print debug statement."""
    if DEBUG >= level:
        logger.debug(message)


def error_message(test_name, json_path, message, schema, fragment):
    """Create an error message."""
    output(0, test_name)
    output(1, f"Path: {json_path}")
    output(2, "Error:")
    output(3, message)
    output(3, f"schema: {json.dumps(schema)}")
    output(3, f"  json: {json.dumps(fragment)}")


def get_entity_ids(sz_abstract_factory: SzAbstractFactory):
    """Get a list of entity IDs."""
    result = []
    sz_engine = sz_abstract_factory.create_engine()

    for record in LOADED_RECORD_KEYS:
        response = sz_engine.get_entity_by_record_id(
            record.get("data_source", ""), record.get("record_id", "")
        )
        response_dict = json.loads(response)
        entity_id = response_dict.get("RESOLVED_ENTITY", {}).get("ENTITY_ID", 0)
        if entity_id not in result:
            result.append(entity_id)

    return result


def is_json_subset(subset_json, full_json):
    """
    Checks if one JSON object (subset_json) is a subset of another (full_json).

    Args:
        subset_json (dict or list): The potential subset JSON.
        full_json (dict or list): The JSON to check against.

    Returns:
        bool: True if subset_json is a subset of full_json, False otherwise.
    """
    if isinstance(subset_json, dict):
        if not isinstance(full_json, dict):
            return False
        for key, value in subset_json.items():
            if key not in full_json:
                return False
            if not is_json_subset(value, full_json[key]):
                return False
        return True

    if isinstance(subset_json, list):
        if not isinstance(full_json, list):
            return False
        for item in subset_json:
            if item not in full_json:
                return False
        return True

    # Primitive types (int, str, bool, float, None)
    return subset_json == full_json


def normalize_files(directory):
    """Deduplicate and sort JSON lines."""
    for root, _, files in os.walk(directory):
        for file in files:
            remove_duplicate_lines(f"{root}/{file}")


def output(indentation, message):
    """Create an indented message."""
    print(f"{'    ' * indentation}{message}")


def output_file(filename, response):
    """Write response to a file."""
    with open(f"{OUTPUT_DIRECTORY}/{filename}.jsonl", "w", encoding="utf-8") as file:
        for line in response:
            file.write(f"{line}\n")


def path_to_testdata(filename: str) -> str:
    """Determine the path to the test data."""
    result = f"{TESTDATA_DIRECTORY}/{filename}"
    return result


def remove_duplicate_lines(input_filepath, output_filepath=None):
    """
    Removes duplicate lines from a text file.

    Args:
        input_filepath (str): The path to the input file.
        output_filepath (str, optional): The path to the output file.
        If None, the input file will be overwritten.
    """
    unique_lines = set()
    try:
        with open(input_filepath, "r", encoding="utf-8") as infile:
            for line in infile:
                line = line.strip()
                if len(line) > 0:
                    line_as_dict = json.loads(line)
                    unique_lines.add(json.dumps(line_as_dict, sort_keys=True))
    except FileNotFoundError:
        logger.warning("Error: Input file '%s' not found.", input_filepath)
        return

    if output_filepath is None:
        output_filepath = input_filepath

    try:
        with open(output_filepath, "w", encoding="utf-8") as outfile:
            for line in sorted(list(unique_lines)):
                outfile.write(f"{line}\n")
        logger.debug("Duplicates removed in '%s'.", output_filepath)
    except IOError:
        logger.error("Error: Could not write to output file '%s'.", output_filepath)


def set_debug(needle, haystack):
    """Determine if debug should be set."""
    global DEBUG

    DEBUG = 0
    if needle in haystack:
        DEBUG = 1


# -----------------------------------------------------------------------------
# Main
# -----------------------------------------------------------------------------


if __name__ == "__main__":

    logger.info("Begin %s", os.path.basename(__file__))

    # Create SzAbstractFactory.

    the_sz_abstract_factory = create_sz_abstract_factory()

    # Insert test data.

    register_datasources(the_sz_abstract_factory)
    add_records(the_sz_abstract_factory)
    LOADED_ENTITY_IDS = get_entity_ids(the_sz_abstract_factory)

    # Make comparisons.

    compare(the_sz_abstract_factory)

    # Delete test data.

    delete_records(the_sz_abstract_factory)

    # Deduplicate and organize output files.

    normalize_files(OUTPUT_DIRECTORY)

    # Epilog.

    logger.info("End   %s", os.path.basename(__file__))
