package g2api

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Flags used by the Senzing G2Engine.
// These flags are single-bit flags.
// BIT_NN is for bits numbered as 1..64 (not 0..63).  Comments are for bits numbered as 0..63
const (
	G2_NO_FLAGS                              int64 = 0
	G2_INITIALIZE_WITH_DEFAULT_CONFIGURATION int64 = 0
	G2_EXPORT_INCLUDE_MULTI_RECORD_ENTITIES  int64 = 1 << iota // 0 we should include entities with "resolved" relationships
	G2_EXPORT_INCLUDE_POSSIBLY_SAME                            // 1 we should include entities with "possibly same" relationships
	G2_EXPORT_INCLUDE_POSSIBLY_RELATED                         // 2 we should include entities with "possibly related" relationships
	G2_EXPORT_INCLUDE_NAME_ONLY                                // 3 we should include entities with "name only" relationships
	G2_EXPORT_INCLUDE_DISCLOSED                                // 4 we should include entities with "disclosed" relationships
	G2_EXPORT_INCLUDE_SINGLE_RECORD_ENTITIES                   // 5 we should include singleton entities

	/* flags for outputting entity relation data  */

	G2_ENTITY_INCLUDE_POSSIBLY_SAME_RELATIONS    // 6 include "possibly same" relationships on entities
	G2_ENTITY_INCLUDE_POSSIBLY_RELATED_RELATIONS // 7 include "possibly related" relationships on entities
	G2_ENTITY_INCLUDE_NAME_ONLY_RELATIONS        // 8 include "name only" relationships on entities
	G2_ENTITY_INCLUDE_DISCLOSED_RELATIONS        // 9 include "disclosed" relationships on entities

	/* flags for outputting entity feature data  */

	G2_ENTITY_INCLUDE_ALL_FEATURES            // 10 include all features for entities
	G2_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES // 11 include only representative features on entities

	/* flags for getting extra information about an entity  */

	G2_ENTITY_INCLUDE_ENTITY_NAME            // 12 include the name of the entity
	G2_ENTITY_INCLUDE_RECORD_SUMMARY         // 13 include the record summary of the entity
	G2_ENTITY_INCLUDE_RECORD_DATA            // 14 include the basic record data for the entity
	G2_ENTITY_INCLUDE_RECORD_MATCHING_INFO   // 15 include the record matching info for the entity
	G2_ENTITY_INCLUDE_RECORD_JSON_DATA       // 16 include the record json data for the entity
	BIT_18                                   // 17
	G2_ENTITY_INCLUDE_RECORD_FEATURE_IDS     // 18 include the features identifiers for the records
	G2_ENTITY_INCLUDE_RELATED_ENTITY_NAME    // 19 include the name of the related entities
	G2_ENTITY_INCLUDE_RELATED_MATCHING_INFO  // 20 include the record matching info of the related entities
	G2_ENTITY_INCLUDE_RELATED_RECORD_SUMMARY // 21 include the record summary of the related entities
	G2_ENTITY_INCLUDE_RELATED_RECORD_DATA    // 22 include the basic record of the related entities

	/* flags for extra feature data  */

	G2_ENTITY_OPTION_INCLUDE_INTERNAL_FEATURES // 23 include internal features
	G2_ENTITY_OPTION_INCLUDE_FEATURE_STATS     // 24 include statistics on features

	/* flags for finding entity path data  */

	G2_FIND_PATH_PREFER_EXCLUDE // 25 excluded entities are still allowed, but not preferred

	/* flags for including search result feature scores  */

	G2_INCLUDE_FEATURE_SCORES                  // 26 include feature scores
	G2_SEARCH_INCLUDE_STATS                    // 27 include statistics from search results
	G2_ENTITY_INCLUDE_RECORD_TYPES             // 28 include the record types of the entity
	G2_ENTITY_INCLUDE_RELATED_RECORD_TYPES     // 29 include the record types of the related entities
	G2_FIND_PATH_MATCHING_INFO                 // 30 include matching info on entity paths
	G2_ENTITY_INCLUDE_RECORD_UNMAPPED_DATA     // 31 include the record unmapped data for the entity
	G2_ENTITY_OPTION_INCLUDE_FEATURE_ELEMENTS  // 32 include feature elements
	G2_FIND_NETWORK_MATCHING_INFO              // 33 include matching info on entity networks
	G2_ENTITY_OPTION_INCLUDE_MATCH_KEY_DETAILS // 34 include internal features
	BIT_36                                     // 35
	BIT_37                                     // 36
	BIT_38                                     // 37
	BIT_39                                     // 38
	BIT_40                                     // 39
	BIT_41                                     // 40
	BIT_42                                     // 41
	BIT_43                                     // 42
	BIT_44                                     // 43
	BIT_45                                     // 44
	BIT_46                                     // 45
	BIT_47                                     // 46
	BIT_48                                     // 47
	BIT_49                                     // 48
	BIT_50                                     // 49
	BIT_51                                     // 50
	BIT_52                                     // 51
	BIT_53                                     // 52
	BIT_54                                     // 53
	BIT_55                                     // 54
	BIT_56                                     // 55
	BIT_57                                     // 56
	BIT_58                                     // 57
	BIT_59                                     // 58
	BIT_60                                     // 59
	BIT_61                                     // 60
	BIT_62                                     // 61
	G2_RETURN_INFO                             // 62 return "WithInfo" information
)

// Flags used by the Senzing G2Engine.
// These flags combine single-bit flags.
const (
	/* Flags for exporting entity data.  */

	G2_EXPORT_INCLUDE_ALL_ENTITIES             = G2_EXPORT_INCLUDE_MULTI_RECORD_ENTITIES | G2_EXPORT_INCLUDE_SINGLE_RECORD_ENTITIES                                               // Include all entities.
	G2_EXPORT_INCLUDE_ALL_HAVING_RELATIONSHIPS = G2_EXPORT_INCLUDE_DISCLOSED | G2_EXPORT_INCLUDE_NAME_ONLY | G2_EXPORT_INCLUDE_POSSIBLY_RELATED | G2_EXPORT_INCLUDE_POSSIBLY_SAME // Include all entities with relationships.

	/* Flags for outputting entity relation data  */

	G2_ENTITY_INCLUDE_ALL_RELATIONS = G2_ENTITY_INCLUDE_DISCLOSED_RELATIONS | G2_ENTITY_INCLUDE_NAME_ONLY_RELATIONS | G2_ENTITY_INCLUDE_POSSIBLY_RELATED_RELATIONS | G2_ENTITY_INCLUDE_POSSIBLY_SAME_RELATIONS // Include all relationships on entities.

	/* Flags for searching for entities.  */

	G2_SEARCH_INCLUDE_ALL_ENTITIES      = G2_SEARCH_INCLUDE_NAME_ONLY | G2_SEARCH_INCLUDE_POSSIBLY_RELATED | G2_SEARCH_INCLUDE_POSSIBLY_SAME | G2_SEARCH_INCLUDE_RESOLVED
	G2_SEARCH_INCLUDE_FEATURE_SCORES    = G2_INCLUDE_FEATURE_SCORES                  // Include feature scores from search results
	G2_SEARCH_INCLUDE_MATCH_KEY_DETAILS = G2_ENTITY_OPTION_INCLUDE_MATCH_KEY_DETAILS // Include detailed match key in search results.
	G2_SEARCH_INCLUDE_NAME_ONLY         = G2_EXPORT_INCLUDE_NAME_ONLY
	G2_SEARCH_INCLUDE_POSSIBLY_RELATED  = G2_EXPORT_INCLUDE_POSSIBLY_RELATED
	G2_SEARCH_INCLUDE_POSSIBLY_SAME     = G2_EXPORT_INCLUDE_POSSIBLY_SAME
	G2_SEARCH_INCLUDE_RESOLVED          = G2_EXPORT_INCLUDE_MULTI_RECORD_ENTITIES

	/* Recommended settings for searching by attributes. */

	G2_SEARCH_BY_ATTRIBUTES_ALL            = G2_ENTITY_INCLUDE_ENTITY_NAME | G2_ENTITY_INCLUDE_RECORD_SUMMARY | G2_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES | G2_SEARCH_INCLUDE_ALL_ENTITIES | G2_SEARCH_INCLUDE_FEATURE_SCORES                               // The recommended flag values for searching by attributes, returning all matching entities.
	G2_SEARCH_BY_ATTRIBUTES_MINIMAL_ALL    = G2_SEARCH_INCLUDE_ALL_ENTITIES                                                                                                                                                                                 // Return minimal data with all matches.
	G2_SEARCH_BY_ATTRIBUTES_MINIMAL_STRONG = G2_SEARCH_INCLUDE_POSSIBLY_SAME | G2_SEARCH_INCLUDE_RESOLVED                                                                                                                                                   // Return minimal data with only the strongest matches.
	G2_SEARCH_BY_ATTRIBUTES_STRONG         = G2_ENTITY_INCLUDE_ENTITY_NAME | G2_ENTITY_INCLUDE_RECORD_SUMMARY | G2_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES | G2_SEARCH_INCLUDE_FEATURE_SCORES | G2_SEARCH_INCLUDE_POSSIBLY_SAME | G2_SEARCH_INCLUDE_RESOLVED // The recommended flag values for searching by attributes, returning only strongly matching entities.

	/* Recommended defaults */

	G2_ENTITY_BRIEF_DEFAULT_FLAGS         = G2_ENTITY_INCLUDE_ALL_RELATIONS | G2_ENTITY_INCLUDE_RECORD_MATCHING_INFO | G2_ENTITY_INCLUDE_RELATED_MATCHING_INFO                                                                                                                                                                                                                                   // The recommended default flag values for a brief entity result.
	G2_ENTITY_DEFAULT_FLAGS               = G2_ENTITY_INCLUDE_ALL_RELATIONS | G2_ENTITY_INCLUDE_ENTITY_NAME | G2_ENTITY_INCLUDE_RECORD_DATA | G2_ENTITY_INCLUDE_RECORD_MATCHING_INFO | G2_ENTITY_INCLUDE_RECORD_SUMMARY | G2_ENTITY_INCLUDE_RELATED_ENTITY_NAME | G2_ENTITY_INCLUDE_RELATED_MATCHING_INFO | G2_ENTITY_INCLUDE_RELATED_RECORD_SUMMARY | G2_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES // The recommended default flag values for getting entities.
	G2_EXPORT_DEFAULT_FLAGS               = G2_ENTITY_DEFAULT_FLAGS | G2_EXPORT_INCLUDE_ALL_ENTITIES                                                                                                                                                                                                                                                                                             // The recommended default flag values for exporting entities.
	G2_FIND_NETWORK_DEFAULT_FLAGS         = G2_ENTITY_INCLUDE_ENTITY_NAME | G2_ENTITY_INCLUDE_RECORD_SUMMARY | G2_FIND_NETWORK_MATCHING_INFO                                                                                                                                                                                                                                                     // The recommended default flag values for finding entity paths.
	G2_FIND_PATH_DEFAULT_FLAGS            = G2_ENTITY_INCLUDE_ENTITY_NAME | G2_ENTITY_INCLUDE_RECORD_SUMMARY | G2_FIND_PATH_MATCHING_INFO                                                                                                                                                                                                                                                        // The recommended default flag values for finding entity paths.
	G2_HOW_ENTITY_DEFAULT_FLAGS           = G2_INCLUDE_FEATURE_SCORES                                                                                                                                                                                                                                                                                                                            // The recommended default flag values for how-analysis on entities.
	G2_RECORD_DEFAULT_FLAGS               = G2_ENTITY_INCLUDE_RECORD_JSON_DATA                                                                                                                                                                                                                                                                                                                   // The recommended default flag values for getting records.
	G2_SEARCH_BY_ATTRIBUTES_DEFAULT_FLAGS = G2_SEARCH_BY_ATTRIBUTES_ALL                                                                                                                                                                                                                                                                                                                          // The recommended default flag values for search-by-attributes
	G2_VIRTUAL_ENTITY_DEFAULT_FLAGS       = G2_ENTITY_DEFAULT_FLAGS                                                                                                                                                                                                                                                                                                                              // The recommended default flag values for virtual-entity-analysis on entities.
	G2_WHY_ENTITIES_DEFAULT_FLAGS         = G2_ENTITY_DEFAULT_FLAGS | G2_ENTITY_OPTION_INCLUDE_FEATURE_STATS | G2_ENTITY_OPTION_INCLUDE_INTERNAL_FEATURES | G2_INCLUDE_FEATURE_SCORES
	G2_WHY_RECORD_IN_ENTITY_DEFAULT_FLAGS = G2_ENTITY_DEFAULT_FLAGS | G2_ENTITY_OPTION_INCLUDE_FEATURE_STATS | G2_ENTITY_OPTION_INCLUDE_INTERNAL_FEATURES | G2_INCLUDE_FEATURE_SCORES // The recommended default flag values for why-analysis on entities.
	G2_WHY_RECORDS_DEFAULT_FLAGS          = G2_ENTITY_DEFAULT_FLAGS | G2_ENTITY_OPTION_INCLUDE_FEATURE_STATS | G2_ENTITY_OPTION_INCLUDE_INTERNAL_FEATURES | G2_INCLUDE_FEATURE_SCORES // The recommended default flag values for why-analysis on entities.
)

/*
The Flags function method returns the Senzing engine configuration.

Input
  - flags: A list of "G2_XXX" flags

Output
  - An int64 combining all input "G2_XXX" flags.
*/
func Flags(flags ...int64) int64 {
	var result int64 = 0
	for _, flag := range flags {
		result = result | flag
	}
	return result
}
