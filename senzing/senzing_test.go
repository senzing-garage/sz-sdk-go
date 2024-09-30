package senzing

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
		actual:   SzExportIncludeMultiRecordEntities,
		expected: 0x0000000000000001,
	},
	{ // 1
		name:     "SZ_SEARCH_INCLUDE_RESOLVED",
		actual:   SzSearchIncludeResolved,
		expected: 0x0000000000000001,
	},
	{ // 2
		name:     "SZ_EXPORT_INCLUDE_POSSIBLY_SAME",
		actual:   SzExportIncludePossiblySame,
		expected: 0x0000000000000002,
	},
	{ // 2
		name:     "SZ_SEARCH_INCLUDE_POSSIBLY_SAME",
		actual:   SzSearchIncludePossiblySame,
		expected: 0x0000000000000002,
	},
	{ // 3
		name:     "SZ_EXPORT_INCLUDE_POSSIBLY_RELATED",
		actual:   SzExportIncludePossiblyRelated,
		expected: 0x0000000000000004,
	},
	{ // 3
		name:     "SZ_SEARCH_INCLUDE_POSSIBLY_RELATED",
		actual:   SzSearchIncludePossiblyRelated,
		expected: 0x0000000000000004,
	},
	{ // 4
		name:     "SZ_EXPORT_INCLUDE_NAME_ONLY",
		actual:   SzExportIncludeNameOnly,
		expected: 0x0000000000000008,
	},
	{ // 4
		name:     "SZ_SEARCH_INCLUDE_NAME_ONLY",
		actual:   SzSearchIncludeNameOnly,
		expected: 0x0000000000000008,
	},
	{ // 5
		name:     "SZ_EXPORT_INCLUDE_DISCLOSED",
		actual:   SzExportIncludeDisclosed,
		expected: 0x0000000000000010,
	},
	{ // 6
		name:     "SZ_EXPORT_INCLUDE_SINGLE_RECORD_ENTITIES",
		actual:   SzExportIncludeSingleRecordEntities,
		expected: 0x0000000000000020,
	},
	{ // 7
		name:     "SZ_ENTITY_INCLUDE_POSSIBLY_SAME_RELATIONS",
		actual:   SzEntityIncludePossiblySameRelations,
		expected: 0x0000000000000040,
	},
	{ // 8
		name:     "SZ_ENTITY_INCLUDE_POSSIBLY_RELATED_RELATIONS",
		actual:   SzEntityIncludePossiblyRelatedRelations,
		expected: 0x0000000000000080,
	},
	{ // 9
		name:     "SZ_ENTITY_INCLUDE_NAME_ONLY_RELATIONS",
		actual:   SzEntityIncludeNameOnlyRelations,
		expected: 0x0000000000000100,
	},
	{ // 10
		name:     "SZ_ENTITY_INCLUDE_DISCLOSED_RELATIONS",
		actual:   SzEntityIncludeDisclosedRelations,
		expected: 0x0000000000000200,
	},
	{ // 11
		name:     "SZ_ENTITY_INCLUDE_ALL_FEATURES",
		actual:   SzEntityIncludeAllFeatures,
		expected: 0x0000000000000400,
	},
	{ // 12
		name:     "SZ_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES",
		actual:   SzEntityIncludeRepresentativeFeatures,
		expected: 0x0000000000000800,
	},
	{ // 13
		name:     "SZ_ENTITY_INCLUDE_ENTITY_NAME",
		actual:   SzEntityIncludeEntityName,
		expected: 0x0000000000001000,
	},
	{ // 14
		name:     "SZ_ENTITY_INCLUDE_RECORD_SUMMARY",
		actual:   SzEntityIncludeRecordSummary,
		expected: 0x0000000000002000,
	},
	{ // 15
		name:     "SZ_ENTITY_INCLUDE_RECORD_DATA",
		actual:   SzEntityIncludeRecordData,
		expected: 0x0000000000004000,
	},
	{ // 16
		name:     "SZ_ENTITY_INCLUDE_RECORD_MATCHING_INFO",
		actual:   SzEntityIncludeRecordMatchingInfo,
		expected: 0x0000000000008000,
	},
	{ // 17
		name:     "SZ_ENTITY_INCLUDE_RECORD_JSON_DATA",
		actual:   SzEntityIncludeRecordJSONData,
		expected: 0x0000000000010000,
	},
	{ // 18
		name:     "Bit18",
		actual:   Bit18,
		expected: 0x0000000000020000,
	},
	{ // 19
		name:     "SZ_ENTITY_INCLUDE_RECORD_FEATURE_IDS",
		actual:   SzEntityIncludeRecordFeatureIDs,
		expected: 0x0000000000040000,
	},
	{ // 20
		name:     "SZ_ENTITY_INCLUDE_RELATED_ENTITY_NAME",
		actual:   SzEntityIncludeRelatedEntityName,
		expected: 0x0000000000080000,
	},
	{ // 21
		name:     "SZ_ENTITY_INCLUDE_RELATED_MATCHING_INFO",
		actual:   SzEntityIncludeRelatedMatchingInfo,
		expected: 0x0000000000100000,
	},
	{ // 22
		name:     "SZ_ENTITY_INCLUDE_RELATED_RECORD_SUMMARY",
		actual:   SzEntityIncludeRelatedRecordSummary,
		expected: 0x0000000000200000,
	},
	{ // 23
		name:     "SZ_ENTITY_INCLUDE_RELATED_RECORD_DATA",
		actual:   SzEntityIncludeRelatedRecordData,
		expected: 0x0000000000400000,
	},
	{ // 24
		name:     "SZ_ENTITY_INCLUDE_INTERNAL_FEATURES",
		actual:   SzEntityIncludeInternalFeatures,
		expected: 0x0000000000800000,
	},
	{ // 25
		name:     "SZ_ENTITY_INCLUDE_FEATURE_STATS",
		actual:   SzEntityIncludeFeatureStats,
		expected: 0x0000000001000000,
	},
	{ // 26
		name:     "SZ_FIND_PATH_STRICT_AVOID",
		actual:   SzFindPathStrictAvoid,
		expected: 0x0000000002000000,
	},
	{ // 27
		name:     "SZ_INCLUDE_FEATURE_SCORES",
		actual:   SzIncludeFeatureScores,
		expected: 0x0000000004000000,
	},
	{ // 28
		name:     "SZ_SEARCH_INCLUDE_STATS",
		actual:   SzSearchIncludeStats,
		expected: 0x0000000008000000,
	},
	{ // 29
		name:     "SZ_ENTITY_INCLUDE_RECORD_TYPES",
		actual:   SzEntityIncludeRecordTypes,
		expected: 0x0000000010000000,
	},
	{ // 30
		name:     "SZ_ENTITY_INCLUDE_RELATED_RECORD_TYPES",
		actual:   SzEntityIncludeRelatedRecordTypes,
		expected: 0x0000000020000000,
	},
	{ // 31
		name:     "SZ_FIND_PATH_INCLUDE_MATCHING_INFO",
		actual:   SzFindPathIncludeMatchingInfo,
		expected: 0x0000000040000000,
	},
	{ // 32
		name:     "SZ_ENTITY_INCLUDE_RECORD_UNMAPPED_DATA",
		actual:   SzEntityIncludeRecordUnmappedData,
		expected: 0x0000000080000000,
	},
	{ // 33
		name:     "Bit33",
		actual:   Bit33,
		expected: 0x0000000100000000,
	},
	{ // 34
		name:     "SZ_FIND_NETWORK_INCLUDE_MATCHING_INFO",
		actual:   SzFindNetworkIncludeMatchingInfo,
		expected: 0x0000000200000000,
	},
	{ // 35
		name:     "SZ_INCLUDE_MATCH_KEY_DETAILS",
		actual:   SzIncludeMatchKeyDetails,
		expected: 0x0000000400000000,
	},
	{ // 36
		name:     "Bit36",
		actual:   Bit36,
		expected: 0x0000000800000000,
	},
	{ // 37
		name:     "Bit37",
		actual:   Bit37,
		expected: 0x0000001000000000,
	},
	{ // 38
		name:     "Bit38",
		actual:   Bit38,
		expected: 0x0000002000000000,
	},
	{ // 39
		name:     "Bit39",
		actual:   Bit39,
		expected: 0x0000004000000000,
	},
	{ // 40
		name:     "Bit40",
		actual:   Bit40,
		expected: 0x0000008000000000,
	},
	{ // 41
		name:     "Bit41",
		actual:   Bit41,
		expected: 0x0000010000000000,
	},
	{ // 42
		name:     "Bit42",
		actual:   Bit42,
		expected: 0x0000020000000000,
	},
	{ // 43
		name:     "Bit43",
		actual:   Bit43,
		expected: 0x0000040000000000,
	},
	{ // 44
		name:     "Bit44",
		actual:   Bit44,
		expected: 0x0000080000000000,
	},
	{ // 45
		name:     "Bit45",
		actual:   Bit45,
		expected: 0x0000100000000000,
	},
	{ // 46
		name:     "Bit46",
		actual:   Bit46,
		expected: 0x0000200000000000,
	},
	{ // 47
		name:     "Bit47",
		actual:   Bit47,
		expected: 0x0000400000000000,
	},
	{ // 48
		name:     "Bit48",
		actual:   Bit48,
		expected: 0x0000800000000000,
	},
	{ // 49
		name:     "Bit49",
		actual:   Bit49,
		expected: 0x0001000000000000,
	},
	{ // 50
		name:     "Bit50",
		actual:   Bit50,
		expected: 0x0002000000000000,
	},
	{ // 51
		name:     "Bit51",
		actual:   Bit51,
		expected: 0x0004000000000000,
	},
	{ // 52
		name:     "Bit52",
		actual:   Bit52,
		expected: 0x0008000000000000,
	},
	{ // 53
		name:     "Bit53",
		actual:   Bit53,
		expected: 0x0010000000000000,
	},
	{ // 54
		name:     "Bit54",
		actual:   Bit54,
		expected: 0x0020000000000000,
	},
	{ // 55
		name:     "Bit55",
		actual:   Bit55,
		expected: 0x0040000000000000,
	},
	{ // 56
		name:     "Bit56",
		actual:   Bit56,
		expected: 0x0080000000000000,
	},
	{ // 57
		name:     "Bit57",
		actual:   Bit57,
		expected: 0x0100000000000000,
	},
	{ // 58
		name:     "Bit58",
		actual:   Bit58,
		expected: 0x0200000000000000,
	},
	{ // 59
		name:     "Bit59",
		actual:   Bit59,
		expected: 0x0400000000000000,
	},
	{ // 60
		name:     "Bit60",
		actual:   Bit60,
		expected: 0x0800000000000000,
	},
	{ // 61
		name:     "Bit61",
		actual:   Bit61,
		expected: 0x1000000000000000,
	},
	{ // 62
		name:     "Bit62",
		actual:   Bit62,
		expected: 0x2000000000000000,
	},
	{ // 63
		name:     "SZ_WITH_INFO",
		actual:   SzWithInfo,
		expected: 0x4000000000000000,
	},
	{
		name:     "SZ_EXPORT_INCLUDE_ALL_ENTITIES",
		actual:   SzExportIncludeAllEntities,
		expected: 0x0000000000000021,
	},
	{
		name:     "SZ_EXPORT_INCLUDE_ALL_HAVING_RELATIONSHIPS",
		actual:   SzExportIncludeAllHavingRelationships,
		expected: 0x000000000000001E,
	},
	{
		name:     "SZ_ENTITY_INCLUDE_ALL_RELATIONS",
		actual:   SzEntityIncludeAllRelations,
		expected: 0x00000000000003C0,
	},
	{
		name:     "SZ_SEARCH_INCLUDE_ALL_ENTITIES",
		actual:   SzSearchIncludeAllEntities,
		expected: 0x000000000000000F,
	},
	{
		name:     "SZ_RECORD_DEFAULT_FLAGS",
		actual:   SzRecordDefaultFlags,
		expected: 0x0000000000010000,
	},
	{
		name:     "SZ_ENTITY_DEFAULT_FLAGS",
		actual:   SzEntityDefaultFlags,
		expected: 0x000000000038FBC0,
	},
	{
		name:     "SZ_ENTITY_BRIEF_DEFAULT_FLAGS",
		actual:   SzEntityBriefDefaultFlags,
		expected: 0x00000000001083C0,
	},
	{
		name:     "SZ_EXPORT_DEFAULT_FLAGS",
		actual:   SzExportDefaultFlags,
		expected: 0x000000000038FBE1,
	},
	{
		name:     "SZ_FIND_PATH_DEFAULT_FLAGS",
		actual:   SzFindPathDefaultFlags,
		expected: 0x0000000040003000,
	},
	{
		name:     "SZ_FIND_NETWORK_DEFAULT_FLAGS",
		actual:   SzFindNetworkDefaultFlags,
		expected: 0x0000000200003000,
	},
	{
		name:     "SZ_WHY_ENTITIES_DEFAULT_FLAGS",
		actual:   SzWhyEntitiesDefaultFlags,
		expected: 0x0000000005B8FBC0,
	},
	{
		name:     "SZ_WHY_RECORDS_DEFAULT_FLAGS",
		actual:   SzWhyRecordsDefaultFlags,
		expected: 0x0000000005B8FBC0,
	},
	{
		name:     "SZ_WHY_RECORD_IN_ENTITY_DEFAULT_FLAGS",
		actual:   SzWhyRecordInEntityIDefaultFlags,
		expected: 0x0000000005B8FBC0,
	},
	{
		name:     "SZ_HOW_ENTITY_DEFAULT_FLAGS",
		actual:   SzHowEntityDefaultFlags,
		expected: 0x0000000004000000,
	},
	{
		name:     "SZ_VIRTUAL_ENTITY_DEFAULT_FLAGS",
		actual:   SzVirtualEntityDefaultFlags,
		expected: 0x000000000038FBC0,
	},
	{
		name:     "SZ_SEARCH_BY_ATTRIBUTES_ALL",
		actual:   SzSearchByAttributesAll,
		expected: 0x000000000400380F,
	},
	{
		name:     "SZ_SEARCH_BY_ATTRIBUTES_STRONG",
		actual:   SzSearchByAttributesStrong,
		expected: 0x0000000004003803,
	},
	{
		name:     "SZ_SEARCH_BY_ATTRIBUTES_MINIMAL_ALL",
		actual:   SzSearchByAttributesMinimalAll,
		expected: 0x000000000000000F,
	},
	{
		name:     "SZ_SEARCH_BY_ATTRIBUTES_MINIMAL_STRONG",
		actual:   SzSearchByAttributesMinimalStrong,
		expected: 0x0000000000000003,
	},
	{
		name:     "SZ_SEARCH_BY_ATTRIBUTES_DEFAULT_FLAGS",
		actual:   SzSearchByAttributesDefaultFlags,
		expected: 0x000000000400380F,
	},
	{
		name:     "OR 4",
		actual:   SzExportIncludeMultiRecordEntities | SzExportIncludePossiblySame | SzExportIncludePossiblyRelated | SzExportIncludeNameOnly,
		expected: 0x000000000000000F,
	},
	{
		name:     "Flags(...)",
		actual:   Flags(SzExportIncludeMultiRecordEntities, SzExportIncludePossiblySame, SzExportIncludePossiblyRelated, SzExportIncludeNameOnly),
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
