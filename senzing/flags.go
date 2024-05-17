package senzing

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const (
	SzNoFlags int64 = 0
)

// Flags used by the Senzing SzEngine.
// These flags are single-bit flags.
// BitNN is for bits numbered as 1..64 (not 0..63).  Comments are for bits numbered as 0..63
const (
	SZ_EXPORT_INCLUDE_MULTI_RECORD_ENTITIES  int64 = 1 << iota // 0 we should include entities with "resolved" relationships
	SZ_EXPORT_INCLUDE_POSSIBLY_SAME                            // 1 we should include entities with "possibly same" relationships
	SZ_EXPORT_INCLUDE_POSSIBLY_RELATED                         // 2 we should include entities with "possibly related" relationships
	SZ_EXPORT_INCLUDE_NAME_ONLY                                // 3 we should include entities with "name only" relationships
	SZ_EXPORT_INCLUDE_DISCLOSED                                // 4 we should include entities with "disclosed" relationships
	SZ_EXPORT_INCLUDE_SINGLE_RECORD_ENTITIES                   // 5 we should include singleton entities

	/* flags for outputting entity relation data  */

	SZ_ENTITY_INCLUDE_POSSIBLY_SAME_RELATIONS    // 6 include "possibly same" relationships on entities
	SZ_ENTITY_INCLUDE_POSSIBLY_RELATED_RELATIONS // 7 include "possibly related" relationships on entities
	SZ_ENTITY_INCLUDE_NAME_ONLY_RELATIONS        // 8 include "name only" relationships on entities
	SZ_ENTITY_INCLUDE_DISCLOSED_RELATIONS        // 9 include "disclosed" relationships on entities

	/* flags for outputting entity feature data  */

	SZ_ENTITY_INCLUDE_ALL_FEATURES            // 10 include all features for entities
	SZ_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES // 11 include only representative features on entities

	/* flags for getting extra information about an entity  */

	SZ_ENTITY_INCLUDE_ENTITY_NAME            // 12 include the name of the entity
	SZ_ENTITY_INCLUDE_RECORD_SUMMARY         // 13 include the record summary of the entity
	SZ_ENTITY_INCLUDE_RECORD_DATA            // 14 include the basic record data for the entity
	SZ_ENTITY_INCLUDE_RECORD_MATCHING_INFO   // 15 include the record matching info for the entity
	SZ_ENTITY_INCLUDE_RECORD_JSON_DATA       // 16 include the record json data for the entity
	Bit18                                    // 17
	SZ_ENTITY_INCLUDE_RECORD_FEATURE_IDS     // 18 include the features identifiers for the records
	SZ_ENTITY_INCLUDE_RELATED_ENTITY_NAME    // 19 include the name of the related entities
	SZ_ENTITY_INCLUDE_RELATED_MATCHING_INFO  // 20 include the record matching info of the related entities
	SZ_ENTITY_INCLUDE_RELATED_RECORD_SUMMARY // 21 include the record summary of the related entities
	SZ_ENTITY_INCLUDE_RELATED_RECORD_DATA    // 22 include the basic record of the related entities

	/* flags for extra feature data  */

	SZ_ENTITY_INCLUDE_INTERNAL_FEATURES // 23 include internal features
	SZ_ENTITY_INCLUDE_FEATURE_STATS     // 24 include statistics on features

	/* flags for finding entity path data  */

	SZ_FIND_PATH_PREFER_EXCLUDE // 25 excluded entities are still allowed, but not preferred

	/* flags for including search result feature scores  */

	SZ_INCLUDE_FEATURE_SCORES              // 26 include feature scores
	SZ_SEARCH_INCLUDE_STATS                // 27 include statistics from search results
	SZ_ENTITY_INCLUDE_RECORD_TYPES         // 28 include the record types of the entity
	SZ_ENTITY_INCLUDE_RELATED_RECORD_TYPES // 29 include the record types of the related entities
	SZ_FIND_PATH_INCLUDE_MATCHING_INFO     // 30 include matching info on entity paths
	SZ_ENTITY_INCLUDE_RECORD_UNMAPPED_DATA // 31 include the record unmapped data for the entity
	SZ_ENTITY_INCLUDE_FEATURE_ELEMENTS     // 32 include feature elements
	SZ_FIND_NETWORK_INCLUDE_MATCHING_INFO  // 33 include matching info on entity networks
	SZ_INCLUDE_MATCH_KEY_DETAILS           // 34 include internal features
	Bit36                                  // 35
	Bit37                                  // 36
	Bit38                                  // 37
	Bit39                                  // 38
	Bit40                                  // 39
	Bit41                                  // 40
	Bit42                                  // 41
	Bit43                                  // 42
	Bit44                                  // 43
	Bit45                                  // 44
	Bit46                                  // 45
	Bit47                                  // 46
	Bit48                                  // 47
	Bit49                                  // 48
	Bit50                                  // 49
	Bit51                                  // 50
	Bit52                                  // 51
	Bit53                                  // 52
	Bit54                                  // 53
	Bit55                                  // 54
	Bit56                                  // 55
	Bit57                                  // 56
	Bit58                                  // 57
	Bit59                                  // 58
	Bit60                                  // 59
	Bit61                                  // 60
	Bit62                                  // 61
	SZ_WITH_INFO                           // 62 return "WithInfo" information
)

// Flags used by the Senzing G2Engine.
// These flags combine single-bit flags.
const (
	/* Flags for exporting entity data.  */

	SZ_EXPORT_INCLUDE_ALL_ENTITIES             = SZ_EXPORT_INCLUDE_MULTI_RECORD_ENTITIES | SZ_EXPORT_INCLUDE_SINGLE_RECORD_ENTITIES                                               // Include all entities.
	SZ_EXPORT_INCLUDE_ALL_HAVING_RELATIONSHIPS = SZ_EXPORT_INCLUDE_DISCLOSED | SZ_EXPORT_INCLUDE_NAME_ONLY | SZ_EXPORT_INCLUDE_POSSIBLY_RELATED | SZ_EXPORT_INCLUDE_POSSIBLY_SAME // Include all entities with relationships.

	/* Flags for outputting entity relation data  */

	SZ_ENTITY_INCLUDE_ALL_RELATIONS = SZ_ENTITY_INCLUDE_DISCLOSED_RELATIONS | SZ_ENTITY_INCLUDE_NAME_ONLY_RELATIONS | SZ_ENTITY_INCLUDE_POSSIBLY_RELATED_RELATIONS | SZ_ENTITY_INCLUDE_POSSIBLY_SAME_RELATIONS // Include all relationships on entities.

	/* Flags for searching for entities.  */

	SZ_SEARCH_INCLUDE_ALL_ENTITIES     = SZ_SEARCH_INCLUDE_NAME_ONLY | SZ_SEARCH_INCLUDE_POSSIBLY_RELATED | SZ_SEARCH_INCLUDE_POSSIBLY_SAME | SZ_SEARCH_INCLUDE_RESOLVED
	SZ_SEARCH_INCLUDE_NAME_ONLY        = SZ_EXPORT_INCLUDE_NAME_ONLY
	SZ_SEARCH_INCLUDE_POSSIBLY_RELATED = SZ_EXPORT_INCLUDE_POSSIBLY_RELATED
	SZ_SEARCH_INCLUDE_POSSIBLY_SAME    = SZ_EXPORT_INCLUDE_POSSIBLY_SAME
	SZ_SEARCH_INCLUDE_RESOLVED         = SZ_EXPORT_INCLUDE_MULTI_RECORD_ENTITIES

	/* Recommended settings for searching by attributes. */

	SZ_SEARCH_BY_ATTRIBUTES_ALL            = SZ_ENTITY_INCLUDE_ENTITY_NAME | SZ_ENTITY_INCLUDE_RECORD_SUMMARY | SZ_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES | SZ_SEARCH_INCLUDE_ALL_ENTITIES | SZ_INCLUDE_FEATURE_SCORES                               // The recommended flag values for searching by attributes, returning all matching entities.
	SZ_SEARCH_BY_ATTRIBUTES_MINIMAL_ALL    = SZ_SEARCH_INCLUDE_ALL_ENTITIES                                                                                                                                                                          // Return minimal data with all matches.
	SZ_SEARCH_BY_ATTRIBUTES_MINIMAL_STRONG = SZ_SEARCH_INCLUDE_POSSIBLY_SAME | SZ_SEARCH_INCLUDE_RESOLVED                                                                                                                                            // Return minimal data with only the strongest matches.
	SZ_SEARCH_BY_ATTRIBUTES_STRONG         = SZ_ENTITY_INCLUDE_ENTITY_NAME | SZ_ENTITY_INCLUDE_RECORD_SUMMARY | SZ_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES | SZ_INCLUDE_FEATURE_SCORES | SZ_SEARCH_INCLUDE_POSSIBLY_SAME | SZ_SEARCH_INCLUDE_RESOLVED // The recommended flag values for searching by attributes, returning only strongly matching entities.

	/* Recommended defaults */

	SZ_ENTITY_BRIEF_DEFAULT_FLAGS         = SZ_ENTITY_INCLUDE_ALL_RELATIONS | SZ_ENTITY_INCLUDE_RECORD_MATCHING_INFO | SZ_ENTITY_INCLUDE_RELATED_MATCHING_INFO                                                                                                                                                                                                                                   // The recommended default flag values for a brief entity result.
	SZ_ENTITY_DEFAULT_FLAGS               = SZ_ENTITY_INCLUDE_ALL_RELATIONS | SZ_ENTITY_INCLUDE_ENTITY_NAME | SZ_ENTITY_INCLUDE_RECORD_DATA | SZ_ENTITY_INCLUDE_RECORD_MATCHING_INFO | SZ_ENTITY_INCLUDE_RECORD_SUMMARY | SZ_ENTITY_INCLUDE_RELATED_ENTITY_NAME | SZ_ENTITY_INCLUDE_RELATED_MATCHING_INFO | SZ_ENTITY_INCLUDE_RELATED_RECORD_SUMMARY | SZ_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES // The recommended default flag values for getting entities.
	SZ_EXPORT_DEFAULT_FLAGS               = SZ_ENTITY_DEFAULT_FLAGS | SZ_EXPORT_INCLUDE_ALL_ENTITIES                                                                                                                                                                                                                                                                                             // The recommended default flag values for exporting entities.
	SZ_FIND_NETWORK_DEFAULT_FLAGS         = SZ_ENTITY_INCLUDE_ENTITY_NAME | SZ_ENTITY_INCLUDE_RECORD_SUMMARY | SZ_FIND_NETWORK_INCLUDE_MATCHING_INFO                                                                                                                                                                                                                                             // The recommended default flag values for finding entity paths.
	SZ_FIND_PATH_DEFAULT_FLAGS            = SZ_ENTITY_INCLUDE_ENTITY_NAME | SZ_ENTITY_INCLUDE_RECORD_SUMMARY | SZ_FIND_PATH_INCLUDE_MATCHING_INFO                                                                                                                                                                                                                                                // The recommended default flag values for finding entity paths.
	SZ_HOW_ENTITY_DEFAULT_FLAGS           = SZ_INCLUDE_FEATURE_SCORES                                                                                                                                                                                                                                                                                                                            // The recommended default flag values for how-analysis on entities.
	SZ_RECORD_DEFAULT_FLAGS               = SZ_ENTITY_INCLUDE_RECORD_JSON_DATA                                                                                                                                                                                                                                                                                                                   // The recommended default flag values for getting records.
	SZ_SEARCH_BY_ATTRIBUTES_DEFAULT_FLAGS = SZ_SEARCH_BY_ATTRIBUTES_ALL                                                                                                                                                                                                                                                                                                                          // The recommended default flag values for search-by-attributes
	SZ_VIRTUAL_ENTITY_DEFAULT_FLAGS       = SZ_ENTITY_DEFAULT_FLAGS                                                                                                                                                                                                                                                                                                                              // The recommended default flag values for virtual-entity-analysis on entities.
	SZ_WHY_ENTITIES_DEFAULT_FLAGS         = SZ_ENTITY_DEFAULT_FLAGS | SZ_ENTITY_INCLUDE_FEATURE_STATS | SZ_ENTITY_INCLUDE_INTERNAL_FEATURES | SZ_INCLUDE_FEATURE_SCORES
	SZ_WHY_RECORD_IN_ENTITY_DEFAULT_FLAGS = SZ_ENTITY_DEFAULT_FLAGS | SZ_ENTITY_INCLUDE_FEATURE_STATS | SZ_ENTITY_INCLUDE_INTERNAL_FEATURES | SZ_INCLUDE_FEATURE_SCORES // The recommended default flag values for why-analysis on entities.
	SZ_WHY_RECORDS_DEFAULT_FLAGS          = SZ_ENTITY_DEFAULT_FLAGS | SZ_ENTITY_INCLUDE_FEATURE_STATS | SZ_ENTITY_INCLUDE_INTERNAL_FEATURES | SZ_INCLUDE_FEATURE_SCORES // The recommended default flag values for why-analysis on entities.
)

/*
The Flags function method returns the Senzing engine configuration.

Input
  - flags: A list of "SZ_XXX" flags

Output
  - An int64 combining all input "SZ_XXX" flags.
*/
func Flags(flags ...int64) int64 {
	var result int64
	for _, flag := range flags {
		result |= flag
	}
	return result
}
