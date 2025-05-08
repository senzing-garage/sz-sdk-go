package senzing_test

import (
	"testing"

	"github.com/senzing-garage/sz-sdk-go/senzing"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name     string
	actual   int64
	expected int64
}{
	{ // 1
		name:     "SZ_EXPORT_INCLUDE_MULTI_RECORD_ENTITIES",
		actual:   senzing.SzExportIncludeMultiRecordEntities,
		expected: 0x0000000000000001,
	},
	{ // 1
		name:     "SZ_SEARCH_INCLUDE_RESOLVED",
		actual:   senzing.SzSearchIncludeResolved,
		expected: 0x0000000000000001,
	},
	{ // 2
		name:     "SZ_EXPORT_INCLUDE_POSSIBLY_SAME",
		actual:   senzing.SzExportIncludePossiblySame,
		expected: 0x0000000000000002,
	},
	{ // 2
		name:     "SZ_SEARCH_INCLUDE_POSSIBLY_SAME",
		actual:   senzing.SzSearchIncludePossiblySame,
		expected: 0x0000000000000002,
	},
	{ // 3
		name:     "SZ_EXPORT_INCLUDE_POSSIBLY_RELATED",
		actual:   senzing.SzExportIncludePossiblyRelated,
		expected: 0x0000000000000004,
	},
	{ // 3
		name:     "SZ_SEARCH_INCLUDE_POSSIBLY_RELATED",
		actual:   senzing.SzSearchIncludePossiblyRelated,
		expected: 0x0000000000000004,
	},
	{ // 4
		name:     "SZ_EXPORT_INCLUDE_NAME_ONLY",
		actual:   senzing.SzExportIncludeNameOnly,
		expected: 0x0000000000000008,
	},
	{ // 4
		name:     "SZ_SEARCH_INCLUDE_NAME_ONLY",
		actual:   senzing.SzSearchIncludeNameOnly,
		expected: 0x0000000000000008,
	},
	{ // 5
		name:     "SZ_EXPORT_INCLUDE_DISCLOSED",
		actual:   senzing.SzExportIncludeDisclosed,
		expected: 0x0000000000000010,
	},
	{ // 6
		name:     "SZ_EXPORT_INCLUDE_SINGLE_RECORD_ENTITIES",
		actual:   senzing.SzExportIncludeSingleRecordEntities,
		expected: 0x0000000000000020,
	},
	{ // 7
		name:     "SZ_ENTITY_INCLUDE_POSSIBLY_SAME_RELATIONS",
		actual:   senzing.SzEntityIncludePossiblySameRelations,
		expected: 0x0000000000000040,
	},
	{ // 8
		name:     "SZ_ENTITY_INCLUDE_POSSIBLY_RELATED_RELATIONS",
		actual:   senzing.SzEntityIncludePossiblyRelatedRelations,
		expected: 0x0000000000000080,
	},
	{ // 9
		name:     "SZ_ENTITY_INCLUDE_NAME_ONLY_RELATIONS",
		actual:   senzing.SzEntityIncludeNameOnlyRelations,
		expected: 0x0000000000000100,
	},
	{ // 10
		name:     "SZ_ENTITY_INCLUDE_DISCLOSED_RELATIONS",
		actual:   senzing.SzEntityIncludeDisclosedRelations,
		expected: 0x0000000000000200,
	},
	{ // 11
		name:     "SZ_ENTITY_INCLUDE_ALL_FEATURES",
		actual:   senzing.SzEntityIncludeAllFeatures,
		expected: 0x0000000000000400,
	},
	{ // 12
		name:     "SZ_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES",
		actual:   senzing.SzEntityIncludeRepresentativeFeatures,
		expected: 0x0000000000000800,
	},
	{ // 13
		name:     "SZ_ENTITY_INCLUDE_ENTITY_NAME",
		actual:   senzing.SzEntityIncludeEntityName,
		expected: 0x0000000000001000,
	},
	{ // 14
		name:     "SZ_ENTITY_INCLUDE_RECORD_SUMMARY",
		actual:   senzing.SzEntityIncludeRecordSummary,
		expected: 0x0000000000002000,
	},
	{ // 15
		name:     "SZ_ENTITY_INCLUDE_RECORD_DATA",
		actual:   senzing.SzEntityIncludeRecordData,
		expected: 0x0000000000004000,
	},
	{ // 16
		name:     "SZ_ENTITY_INCLUDE_RECORD_MATCHING_INFO",
		actual:   senzing.SzEntityIncludeRecordMatchingInfo,
		expected: 0x0000000000008000,
	},
	{ // 17
		name:     "SZ_ENTITY_INCLUDE_RECORD_JSON_DATA",
		actual:   senzing.SzEntityIncludeRecordJSONData,
		expected: 0x0000000000010000,
	},
	{ // 18
		name:     "Bit18",
		actual:   senzing.Bit18,
		expected: 0x0000000000020000,
	},
	{ // 19
		name:     "SZ_ENTITY_INCLUDE_RECORD_FEATURES",
		actual:   senzing.SzEntityIncludeRecordFeatures,
		expected: 0x0000000000040000,
	},
	{ // 20
		name:     "SZ_ENTITY_INCLUDE_RELATED_ENTITY_NAME",
		actual:   senzing.SzEntityIncludeRelatedEntityName,
		expected: 0x0000000000080000,
	},
	{ // 21
		name:     "SZ_ENTITY_INCLUDE_RELATED_MATCHING_INFO",
		actual:   senzing.SzEntityIncludeRelatedMatchingInfo,
		expected: 0x0000000000100000,
	},
	{ // 22
		name:     "SZ_ENTITY_INCLUDE_RELATED_RECORD_SUMMARY",
		actual:   senzing.SzEntityIncludeRelatedRecordSummary,
		expected: 0x0000000000200000,
	},
	{ // 23
		name:     "SZ_ENTITY_INCLUDE_RELATED_RECORD_DATA",
		actual:   senzing.SzEntityIncludeRelatedRecordData,
		expected: 0x0000000000400000,
	},
	{ // 24
		name:     "SZ_ENTITY_INCLUDE_INTERNAL_FEATURES",
		actual:   senzing.SzEntityIncludeInternalFeatures,
		expected: 0x0000000000800000,
	},
	{ // 25
		name:     "SZ_ENTITY_INCLUDE_FEATURE_STATS",
		actual:   senzing.SzEntityIncludeFeatureStats,
		expected: 0x0000000001000000,
	},
	{ // 26
		name:     "SZ_FIND_PATH_STRICT_AVOID",
		actual:   senzing.SzFindPathStrictAvoid,
		expected: 0x0000000002000000,
	},
	{ // 27
		name:     "SZ_INCLUDE_FEATURE_SCORES",
		actual:   senzing.SzIncludeFeatureScores,
		expected: 0x0000000004000000,
	},
	{ // 28
		name:     "SZ_SEARCH_INCLUDE_STATS",
		actual:   senzing.SzSearchIncludeStats,
		expected: 0x0000000008000000,
	},
	{ // 29
		name:     "SZ_ENTITY_INCLUDE_RECORD_TYPES",
		actual:   senzing.SzEntityIncludeRecordTypes,
		expected: 0x0000000010000000,
	},
	{ // 30
		name:     "SZ_ENTITY_INCLUDE_RELATED_RECORD_TYPES",
		actual:   senzing.SzEntityIncludeRelatedRecordTypes,
		expected: 0x0000000020000000,
	},
	{ // 31
		name:     "SZ_FIND_PATH_INCLUDE_MATCHING_INFO",
		actual:   senzing.SzFindPathIncludeMatchingInfo,
		expected: 0x0000000040000000,
	},
	{ // 32
		name:     "SZ_ENTITY_INCLUDE_RECORD_UNMAPPED_DATA",
		actual:   senzing.SzEntityIncludeRecordUnmappedData,
		expected: 0x0000000080000000,
	},
	{ // 33
		name:     "SZ_SEARCH_INCLUDE_ALL_CANDIDATES",
		actual:   senzing.SzSearchIncludeAllCandidates,
		expected: 0x0000000100000000,
	},
	{ // 34
		name:     "SZ_FIND_NETWORK_INCLUDE_MATCHING_INFO",
		actual:   senzing.SzFindNetworkIncludeMatchingInfo,
		expected: 0x0000000200000000,
	},
	{ // 35
		name:     "SZ_INCLUDE_MATCH_KEY_DETAILS",
		actual:   senzing.SzIncludeMatchKeyDetails,
		expected: 0x0000000400000000,
	},
	{ // 36
		name:     "SZ_ENTITY_INCLUDE_RECORD_FEATURE_DETAILS",
		actual:   senzing.SzEntityIncludeRecordFeatureDetails,
		expected: 0x0000000800000000,
	},
	{ // 37
		name:     "SZ_ENTITY_INCLUDE_RECORD_FEATURE_STATS",
		actual:   senzing.SzEntityIncludeRecordFeatureStats,
		expected: 0x0000001000000000,
	},
	{ // 38
		name:     "SZ_SEARCH_INCLUDE_REQUEST",
		actual:   senzing.SzSearchIncludeRequest,
		expected: 0x0000002000000000,
	},
	{ // 39
		name:     "SZ_SEARCH_INCLUDE_REQUEST_DETAILS",
		actual:   senzing.SzSearchIncludeRequestDetails,
		expected: 0x0000004000000000,
	},
	{ // 40
		name:     "Bit40",
		actual:   senzing.Bit40,
		expected: 0x0000008000000000,
	},
	{ // 41
		name:     "Bit41",
		actual:   senzing.Bit41,
		expected: 0x0000010000000000,
	},
	{ // 42
		name:     "Bit42",
		actual:   senzing.Bit42,
		expected: 0x0000020000000000,
	},
	{ // 43
		name:     "Bit43",
		actual:   senzing.Bit43,
		expected: 0x0000040000000000,
	},
	{ // 44
		name:     "Bit44",
		actual:   senzing.Bit44,
		expected: 0x0000080000000000,
	},
	{ // 45
		name:     "Bit45",
		actual:   senzing.Bit45,
		expected: 0x0000100000000000,
	},
	{ // 46
		name:     "Bit46",
		actual:   senzing.Bit46,
		expected: 0x0000200000000000,
	},
	{ // 47
		name:     "Bit47",
		actual:   senzing.Bit47,
		expected: 0x0000400000000000,
	},
	{ // 48
		name:     "Bit48",
		actual:   senzing.Bit48,
		expected: 0x0000800000000000,
	},
	{ // 49
		name:     "Bit49",
		actual:   senzing.Bit49,
		expected: 0x0001000000000000,
	},
	{ // 50
		name:     "Bit50",
		actual:   senzing.Bit50,
		expected: 0x0002000000000000,
	},
	{ // 51
		name:     "Bit51",
		actual:   senzing.Bit51,
		expected: 0x0004000000000000,
	},
	{ // 52
		name:     "Bit52",
		actual:   senzing.Bit52,
		expected: 0x0008000000000000,
	},
	{ // 53
		name:     "Bit53",
		actual:   senzing.Bit53,
		expected: 0x0010000000000000,
	},
	{ // 54
		name:     "Bit54",
		actual:   senzing.Bit54,
		expected: 0x0020000000000000,
	},
	{ // 55
		name:     "Bit55",
		actual:   senzing.Bit55,
		expected: 0x0040000000000000,
	},
	{ // 56
		name:     "Bit56",
		actual:   senzing.Bit56,
		expected: 0x0080000000000000,
	},
	{ // 57
		name:     "Bit57",
		actual:   senzing.Bit57,
		expected: 0x0100000000000000,
	},
	{ // 58
		name:     "Bit58",
		actual:   senzing.Bit58,
		expected: 0x0200000000000000,
	},
	{ // 59
		name:     "Bit59",
		actual:   senzing.Bit59,
		expected: 0x0400000000000000,
	},
	{ // 60
		name:     "Bit60",
		actual:   senzing.Bit60,
		expected: 0x0800000000000000,
	},
	{ // 61
		name:     "Bit61",
		actual:   senzing.Bit61,
		expected: 0x1000000000000000,
	},
	{ // 62
		name:     "Bit62",
		actual:   senzing.Bit62,
		expected: 0x2000000000000000,
	},
	{ // 63
		name:     "SZ_WITH_INFO",
		actual:   senzing.SzWithInfo,
		expected: 0x4000000000000000,
	},
	{
		name:     "SZ_EXPORT_INCLUDE_ALL_ENTITIES",
		actual:   senzing.SzExportIncludeAllEntities,
		expected: 0x0000000000000021,
	},
	{
		name:     "SZ_EXPORT_INCLUDE_ALL_HAVING_RELATIONSHIPS",
		actual:   senzing.SzExportIncludeAllHavingRelationships,
		expected: 0x000000000000001E,
	},
	{
		name:     "SZ_ENTITY_INCLUDE_ALL_RELATIONS",
		actual:   senzing.SzEntityIncludeAllRelations,
		expected: 0x00000000000003C0,
	},
	{
		name:     "SZ_SEARCH_INCLUDE_ALL_ENTITIES",
		actual:   senzing.SzSearchIncludeAllEntities,
		expected: 0x000000000000000F,
	},
	{
		name:     "SZ_RECORD_DEFAULT_FLAGS",
		actual:   senzing.SzRecordDefaultFlags,
		expected: 0x0000000000010000,
	},
	{
		name:     "SZ_ENTITY_DEFAULT_FLAGS",
		actual:   senzing.SzEntityDefaultFlags,
		expected: 0x000000000018FBC0,
	},
	{
		name:     "SZ_ENTITY_BRIEF_DEFAULT_FLAGS",
		actual:   senzing.SzEntityBriefDefaultFlags,
		expected: 0x00000000001083C0,
	},
	{
		name:     "SZ_EXPORT_DEFAULT_FLAGS",
		actual:   senzing.SzExportDefaultFlags,
		expected: 0x000000000018FBE1,
	},
	{
		name:     "SZ_FIND_PATH_DEFAULT_FLAGS",
		actual:   senzing.SzFindPathDefaultFlags,
		expected: 0x0000000040003000,
	},
	{
		name:     "SZ_FIND_NETWORK_DEFAULT_FLAGS",
		actual:   senzing.SzFindNetworkDefaultFlags,
		expected: 0x0000000200003000,
	},
	{
		name:     "SZ_WHY_ENTITIES_DEFAULT_FLAGS",
		actual:   senzing.SzWhyEntitiesDefaultFlags,
		expected: 0x0000000004000000,
	},
	{
		name:     "SZ_WHY_RECORDS_DEFAULT_FLAGS",
		actual:   senzing.SzWhyRecordsDefaultFlags,
		expected: 0x0000000004000000,
	},
	{
		name:     "SZ_WHY_RECORD_IN_ENTITY_DEFAULT_FLAGS",
		actual:   senzing.SzWhyRecordInEntityIDefaultFlags,
		expected: 0x0000000004000000,
	},
	{
		name:     "SZ_HOW_ENTITY_DEFAULT_FLAGS",
		actual:   senzing.SzHowEntityDefaultFlags,
		expected: 0x0000000004000000,
	},
	{
		name:     "SZ_VIRTUAL_ENTITY_DEFAULT_FLAGS",
		actual:   senzing.SzVirtualEntityDefaultFlags,
		expected: 0x000000000000F800,
	},
	{
		name:     "SZ_SEARCH_BY_ATTRIBUTES_ALL",
		actual:   senzing.SzSearchByAttributesAll,
		expected: 0x000000000C00380F,
	},
	{
		name:     "SZ_SEARCH_BY_ATTRIBUTES_STRONG",
		actual:   senzing.SzSearchByAttributesStrong,
		expected: 0x000000000C003803,
	},
	{
		name:     "SZ_SEARCH_BY_ATTRIBUTES_MINIMAL_ALL",
		actual:   senzing.SzSearchByAttributesMinimalAll,
		expected: 0x000000000800000F,
	},
	{
		name:     "SZ_SEARCH_BY_ATTRIBUTES_MINIMAL_STRONG",
		actual:   senzing.SzSearchByAttributesMinimalStrong,
		expected: 0x0000000008000003,
	},
	{
		name:     "SZ_SEARCH_BY_ATTRIBUTES_DEFAULT_FLAGS",
		actual:   senzing.SzSearchByAttributesDefaultFlags,
		expected: 0x000000000C00380F,
	},
	{
		name: "OR 4",
		actual: senzing.SzExportIncludeMultiRecordEntities |
			senzing.SzExportIncludePossiblySame |
			senzing.SzExportIncludePossiblyRelated |
			senzing.SzExportIncludeNameOnly,
		expected: 0x000000000000000F,
	},
	{
		name: "Flags(...)",
		actual: senzing.Flags(
			senzing.SzExportIncludeMultiRecordEntities,
			senzing.SzExportIncludePossiblySame,
			senzing.SzExportIncludePossiblyRelated,
			senzing.SzExportIncludeNameOnly),
		expected: 0x000000000000000F,
	},
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestFlagValues(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			assert.Equal(test, testCase.expected, testCase.actual, testCase.name)
		})
	}
}
