package szinterface

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name     string
	actual   int64
	expected int64
}{
	{ // 1
		name:     "SZ_EXPORT_INCLUDE_MULTI_RECORD_ENTITIES",
		actual:   SZ_EXPORT_INCLUDE_MULTI_RECORD_ENTITIES,
		expected: 0x0000000000000001,
	},
	{ // 1
		name:     "SZ_SEARCH_INCLUDE_RESOLVED",
		actual:   SZ_SEARCH_INCLUDE_RESOLVED,
		expected: 0x0000000000000001,
	},
	{ // 2
		name:     "SZ_EXPORT_INCLUDE_POSSIBLY_SAME",
		actual:   SZ_EXPORT_INCLUDE_POSSIBLY_SAME,
		expected: 0x0000000000000002,
	},
	{ // 2
		name:     "SZ_SEARCH_INCLUDE_POSSIBLY_SAME",
		actual:   SZ_SEARCH_INCLUDE_POSSIBLY_SAME,
		expected: 0x0000000000000002,
	},
	{ // 3
		name:     "SZ_EXPORT_INCLUDE_POSSIBLY_RELATED",
		actual:   SZ_EXPORT_INCLUDE_POSSIBLY_RELATED,
		expected: 0x0000000000000004,
	},
	{ // 3
		name:     "SZ_SEARCH_INCLUDE_POSSIBLY_RELATED",
		actual:   SZ_SEARCH_INCLUDE_POSSIBLY_RELATED,
		expected: 0x0000000000000004,
	},
	{ // 4
		name:     "SZ_EXPORT_INCLUDE_NAME_ONLY",
		actual:   SZ_EXPORT_INCLUDE_NAME_ONLY,
		expected: 0x0000000000000008,
	},
	{ // 4
		name:     "SZ_SEARCH_INCLUDE_NAME_ONLY",
		actual:   SZ_SEARCH_INCLUDE_NAME_ONLY,
		expected: 0x0000000000000008,
	},
	{ // 5
		name:     "SZ_EXPORT_INCLUDE_DISCLOSED",
		actual:   SZ_EXPORT_INCLUDE_DISCLOSED,
		expected: 0x0000000000000010,
	},
	{ // 6
		name:     "SZ_EXPORT_INCLUDE_SINGLE_RECORD_ENTITIES",
		actual:   SZ_EXPORT_INCLUDE_SINGLE_RECORD_ENTITIES,
		expected: 0x0000000000000020,
	},
	{ // 7
		name:     "SZ_ENTITY_INCLUDE_POSSIBLY_SAME_RELATIONS",
		actual:   SZ_ENTITY_INCLUDE_POSSIBLY_SAME_RELATIONS,
		expected: 0x0000000000000040,
	},
	{ // 8
		name:     "SZ_ENTITY_INCLUDE_POSSIBLY_RELATED_RELATIONS",
		actual:   SZ_ENTITY_INCLUDE_POSSIBLY_RELATED_RELATIONS,
		expected: 0x0000000000000080,
	},
	{ // 9
		name:     "SZ_ENTITY_INCLUDE_NAME_ONLY_RELATIONS",
		actual:   SZ_ENTITY_INCLUDE_NAME_ONLY_RELATIONS,
		expected: 0x0000000000000100,
	},
	{ // 10
		name:     "SZ_ENTITY_INCLUDE_DISCLOSED_RELATIONS",
		actual:   SZ_ENTITY_INCLUDE_DISCLOSED_RELATIONS,
		expected: 0x0000000000000200,
	},
	{ // 11
		name:     "SZ_ENTITY_INCLUDE_ALL_FEATURES",
		actual:   SZ_ENTITY_INCLUDE_ALL_FEATURES,
		expected: 0x0000000000000400,
	},
	{ // 12
		name:     "SZ_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES",
		actual:   SZ_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES,
		expected: 0x0000000000000800,
	},
	{ // 13
		name:     "SZ_ENTITY_INCLUDE_ENTITY_NAME",
		actual:   SZ_ENTITY_INCLUDE_ENTITY_NAME,
		expected: 0x0000000000001000,
	},
	{ // 14
		name:     "SZ_ENTITY_INCLUDE_RECORD_SUMMARY",
		actual:   SZ_ENTITY_INCLUDE_RECORD_SUMMARY,
		expected: 0x0000000000002000,
	},
	{ // 15
		name:     "SZ_ENTITY_INCLUDE_RECORD_DATA",
		actual:   SZ_ENTITY_INCLUDE_RECORD_DATA,
		expected: 0x0000000000004000,
	},
	{ // 16
		name:     "SZ_ENTITY_INCLUDE_RECORD_MATCHING_INFO",
		actual:   SZ_ENTITY_INCLUDE_RECORD_MATCHING_INFO,
		expected: 0x0000000000008000,
	},
	{ // 17
		name:     "SZ_ENTITY_INCLUDE_RECORD_JSON_DATA",
		actual:   SZ_ENTITY_INCLUDE_RECORD_JSON_DATA,
		expected: 0x0000000000010000,
	},
	{ // 18
		name:     "BIT_18",
		actual:   BIT_18,
		expected: 0x0000000000020000,
	},
	{ // 19
		name:     "SZ_ENTITY_INCLUDE_RECORD_FEATURE_IDS",
		actual:   SZ_ENTITY_INCLUDE_RECORD_FEATURE_IDS,
		expected: 0x0000000000040000,
	},
	{ // 20
		name:     "SZ_ENTITY_INCLUDE_RELATED_ENTITY_NAME",
		actual:   SZ_ENTITY_INCLUDE_RELATED_ENTITY_NAME,
		expected: 0x0000000000080000,
	},
	{ // 21
		name:     "SZ_ENTITY_INCLUDE_RELATED_MATCHING_INFO",
		actual:   SZ_ENTITY_INCLUDE_RELATED_MATCHING_INFO,
		expected: 0x0000000000100000,
	},
	{ // 22
		name:     "SZ_ENTITY_INCLUDE_RELATED_RECORD_SUMMARY",
		actual:   SZ_ENTITY_INCLUDE_RELATED_RECORD_SUMMARY,
		expected: 0x0000000000200000,
	},
	{ // 23
		name:     "SZ_ENTITY_INCLUDE_RELATED_RECORD_DATA",
		actual:   SZ_ENTITY_INCLUDE_RELATED_RECORD_DATA,
		expected: 0x0000000000400000,
	},
	{ // 24
		name:     "SZ_ENTITY_OPTION_INCLUDE_INTERNAL_FEATURES",
		actual:   SZ_ENTITY_OPTION_INCLUDE_INTERNAL_FEATURES,
		expected: 0x0000000000800000,
	},
	{ // 25
		name:     "SZ_ENTITY_OPTION_INCLUDE_FEATURE_STATS",
		actual:   SZ_ENTITY_OPTION_INCLUDE_FEATURE_STATS,
		expected: 0x0000000001000000,
	},
	{ // 26
		name:     "SZ_FIND_PATH_PREFER_EXCLUDE",
		actual:   SZ_FIND_PATH_PREFER_EXCLUDE,
		expected: 0x0000000002000000,
	},
	{ // 27
		name:     "SZ_INCLUDE_FEATURE_SCORES",
		actual:   SZ_INCLUDE_FEATURE_SCORES,
		expected: 0x0000000004000000,
	},
	{ // 27
		name:     "SZ_SEARCH_INCLUDE_FEATURE_SCORES",
		actual:   SZ_SEARCH_INCLUDE_FEATURE_SCORES,
		expected: 0x0000000004000000,
	},
	{ // 28
		name:     "SZ_SEARCH_INCLUDE_STATS",
		actual:   SZ_SEARCH_INCLUDE_STATS,
		expected: 0x0000000008000000,
	},
	{ // 29
		name:     "SZ_ENTITY_INCLUDE_RECORD_TYPES",
		actual:   SZ_ENTITY_INCLUDE_RECORD_TYPES,
		expected: 0x0000000010000000,
	},
	{ // 30
		name:     "SZ_ENTITY_INCLUDE_RELATED_RECORD_TYPES",
		actual:   SZ_ENTITY_INCLUDE_RELATED_RECORD_TYPES,
		expected: 0x0000000020000000,
	},
	{ // 31
		name:     "SZ_FIND_PATH_MATCHING_INFO",
		actual:   SZ_FIND_PATH_MATCHING_INFO,
		expected: 0x0000000040000000,
	},
	{ // 32
		name:     "SZ_ENTITY_INCLUDE_RECORD_UNMAPPED_DATA",
		actual:   SZ_ENTITY_INCLUDE_RECORD_UNMAPPED_DATA,
		expected: 0x0000000080000000,
	},
	{ // 33
		name:     "SZ_ENTITY_OPTION_INCLUDE_FEATURE_ELEMENTS",
		actual:   SZ_ENTITY_OPTION_INCLUDE_FEATURE_ELEMENTS,
		expected: 0x0000000100000000,
	},
	{ // 34
		name:     "SZ_FIND_NETWORK_MATCHING_INFO",
		actual:   SZ_FIND_NETWORK_MATCHING_INFO,
		expected: 0x0000000200000000,
	},
	{ // 35
		name:     "SZ_ENTITY_OPTION_INCLUDE_MATCH_KEY_DETAILS",
		actual:   SZ_ENTITY_OPTION_INCLUDE_MATCH_KEY_DETAILS,
		expected: 0x0000000400000000,
	},
	{ // 35
		name:     "SZ_SEARCH_INCLUDE_MATCH_KEY_DETAILS",
		actual:   SZ_SEARCH_INCLUDE_MATCH_KEY_DETAILS,
		expected: 0x0000000400000000,
	},
	{ // 36
		name:     "BIT_36",
		actual:   BIT_36,
		expected: 0x0000000800000000,
	},
	{ // 37
		name:     "BIT_37",
		actual:   BIT_37,
		expected: 0x0000001000000000,
	},
	{ // 38
		name:     "BIT_38",
		actual:   BIT_38,
		expected: 0x0000002000000000,
	},
	{ // 39
		name:     "BIT_39",
		actual:   BIT_39,
		expected: 0x0000004000000000,
	},
	{ // 40
		name:     "BIT_40",
		actual:   BIT_40,
		expected: 0x0000008000000000,
	},
	{ // 41
		name:     "BIT_41",
		actual:   BIT_41,
		expected: 0x0000010000000000,
	},
	{ // 42
		name:     "BIT_42",
		actual:   BIT_42,
		expected: 0x0000020000000000,
	},
	{ // 43
		name:     "BIT_43",
		actual:   BIT_43,
		expected: 0x0000040000000000,
	},
	{ // 44
		name:     "BIT_44",
		actual:   BIT_44,
		expected: 0x0000080000000000,
	},
	{ // 45
		name:     "BIT_45",
		actual:   BIT_45,
		expected: 0x0000100000000000,
	},
	{ // 46
		name:     "BIT_46",
		actual:   BIT_46,
		expected: 0x0000200000000000,
	},
	{ // 47
		name:     "BIT_47",
		actual:   BIT_47,
		expected: 0x0000400000000000,
	},
	{ // 48
		name:     "BIT_48",
		actual:   BIT_48,
		expected: 0x0000800000000000,
	},
	{ // 49
		name:     "BIT_49",
		actual:   BIT_49,
		expected: 0x0001000000000000,
	},
	{ // 50
		name:     "BIT_50",
		actual:   BIT_50,
		expected: 0x0002000000000000,
	},
	{ // 51
		name:     "BIT_51",
		actual:   BIT_51,
		expected: 0x0004000000000000,
	},
	{ // 52
		name:     "BIT_52",
		actual:   BIT_52,
		expected: 0x0008000000000000,
	},
	{ // 53
		name:     "BIT_53",
		actual:   BIT_53,
		expected: 0x0010000000000000,
	},
	{ // 54
		name:     "BIT_54",
		actual:   BIT_54,
		expected: 0x0020000000000000,
	},
	{ // 55
		name:     "BIT_55",
		actual:   BIT_55,
		expected: 0x0040000000000000,
	},
	{ // 56
		name:     "BIT_56",
		actual:   BIT_56,
		expected: 0x0080000000000000,
	},
	{ // 57
		name:     "BIT_57",
		actual:   BIT_57,
		expected: 0x0100000000000000,
	},
	{ // 58
		name:     "BIT_58",
		actual:   BIT_58,
		expected: 0x0200000000000000,
	},
	{ // 59
		name:     "BIT_59",
		actual:   BIT_59,
		expected: 0x0400000000000000,
	},
	{ // 60
		name:     "BIT_60",
		actual:   BIT_60,
		expected: 0x0800000000000000,
	},
	{ // 61
		name:     "BIT_61",
		actual:   BIT_61,
		expected: 0x1000000000000000,
	},
	{ // 62
		name:     "BIT_62",
		actual:   BIT_62,
		expected: 0x2000000000000000,
	},
	{ // 63
		name:     "SZ_WITH_INFO",
		actual:   SZ_WITH_INFO,
		expected: 0x4000000000000000,
	},
	{
		name:     "SZ_EXPORT_INCLUDE_ALL_ENTITIES",
		actual:   SZ_EXPORT_INCLUDE_ALL_ENTITIES,
		expected: 0x0000000000000021,
	},
	{
		name:     "SZ_EXPORT_INCLUDE_ALL_HAVING_RELATIONSHIPS",
		actual:   SZ_EXPORT_INCLUDE_ALL_HAVING_RELATIONSHIPS,
		expected: 0x000000000000001E,
	},
	{
		name:     "SZ_ENTITY_INCLUDE_ALL_RELATIONS",
		actual:   SZ_ENTITY_INCLUDE_ALL_RELATIONS,
		expected: 0x00000000000003C0,
	},
	{
		name:     "SZ_SEARCH_INCLUDE_ALL_ENTITIES",
		actual:   SZ_SEARCH_INCLUDE_ALL_ENTITIES,
		expected: 0x000000000000000F,
	},
	{
		name:     "SZ_RECORD_DEFAULT_FLAGS",
		actual:   SZ_RECORD_DEFAULT_FLAGS,
		expected: 0x0000000000010000,
	},
	{
		name:     "SZ_ENTITY_DEFAULT_FLAGS",
		actual:   SZ_ENTITY_DEFAULT_FLAGS,
		expected: 0x000000000038FBC0,
	},
	{
		name:     "SZ_ENTITY_BRIEF_DEFAULT_FLAGS",
		actual:   SZ_ENTITY_BRIEF_DEFAULT_FLAGS,
		expected: 0x00000000001083C0,
	},
	{
		name:     "SZ_EXPORT_DEFAULT_FLAGS",
		actual:   SZ_EXPORT_DEFAULT_FLAGS,
		expected: 0x000000000038FBE1,
	},
	{
		name:     "SZ_FIND_PATH_DEFAULT_FLAGS",
		actual:   SZ_FIND_PATH_DEFAULT_FLAGS,
		expected: 0x0000000040003000,
	},
	{
		name:     "SZ_FIND_NETWORK_DEFAULT_FLAGS",
		actual:   SZ_FIND_NETWORK_DEFAULT_FLAGS,
		expected: 0x0000000200003000,
	},
	{
		name:     "SZ_WHY_ENTITIES_DEFAULT_FLAGS",
		actual:   SZ_WHY_ENTITIES_DEFAULT_FLAGS,
		expected: 0x0000000005B8FBC0,
	},
	{
		name:     "SZ_WHY_RECORDS_DEFAULT_FLAGS",
		actual:   SZ_WHY_RECORDS_DEFAULT_FLAGS,
		expected: 0x0000000005B8FBC0,
	},
	{
		name:     "SZ_WHY_RECORD_IN_ENTITY_DEFAULT_FLAGS",
		actual:   SZ_WHY_RECORD_IN_ENTITY_DEFAULT_FLAGS,
		expected: 0x0000000005B8FBC0,
	},
	{
		name:     "SZ_HOW_ENTITY_DEFAULT_FLAGS",
		actual:   SZ_HOW_ENTITY_DEFAULT_FLAGS,
		expected: 0x0000000004000000,
	},
	{
		name:     "SZ_VIRTUAL_ENTITY_DEFAULT_FLAGS",
		actual:   SZ_VIRTUAL_ENTITY_DEFAULT_FLAGS,
		expected: 0x000000000038FBC0,
	},
	{
		name:     "SZ_SEARCH_BY_ATTRIBUTES_ALL",
		actual:   SZ_SEARCH_BY_ATTRIBUTES_ALL,
		expected: 0x000000000400380F,
	},
	{
		name:     "SZ_SEARCH_BY_ATTRIBUTES_STRONG",
		actual:   SZ_SEARCH_BY_ATTRIBUTES_STRONG,
		expected: 0x0000000004003803,
	},
	{
		name:     "SZ_SEARCH_BY_ATTRIBUTES_MINIMAL_ALL",
		actual:   SZ_SEARCH_BY_ATTRIBUTES_MINIMAL_ALL,
		expected: 0x000000000000000F,
	},
	{
		name:     "SZ_SEARCH_BY_ATTRIBUTES_MINIMAL_STRONG",
		actual:   SZ_SEARCH_BY_ATTRIBUTES_MINIMAL_STRONG,
		expected: 0x0000000000000003,
	},
	{
		name:     "SZ_SEARCH_BY_ATTRIBUTES_DEFAULT_FLAGS",
		actual:   SZ_SEARCH_BY_ATTRIBUTES_DEFAULT_FLAGS,
		expected: 0x000000000400380F,
	},
	{
		name:     "OR 4",
		actual:   SZ_EXPORT_INCLUDE_MULTI_RECORD_ENTITIES | SZ_EXPORT_INCLUDE_POSSIBLY_SAME | SZ_EXPORT_INCLUDE_POSSIBLY_RELATED | SZ_EXPORT_INCLUDE_NAME_ONLY,
		expected: 0x000000000000000F,
	},
	{
		name:     "Flags(...)",
		actual:   Flags(SZ_EXPORT_INCLUDE_MULTI_RECORD_ENTITIES, SZ_EXPORT_INCLUDE_POSSIBLY_SAME, SZ_EXPORT_INCLUDE_POSSIBLY_RELATED, SZ_EXPORT_INCLUDE_NAME_ONLY),
		expected: 0x000000000000000F,
	},
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestFlagValues(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			assert.Equal(test, testCase.expected, int64(testCase.actual), testCase.name)
		})
	}
}
