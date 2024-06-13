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
	SzExportIncludeMultiRecordEntities  int64 = 1 << iota // 0 we should include entities with "resolved" relationships
	SzExportIncludePossiblySame                           // 1 we should include entities with "possibly same" relationships
	SzExportIncludePossiblyRelated                        // 2 we should include entities with "possibly related" relationships
	SzExportIncludeNameOnly                               // 3 we should include entities with "name only" relationships
	SzExportIncludeDisclosed                              // 4 we should include entities with "disclosed" relationships
	SzExportIncludeSingleRecordEntities                   // 5 we should include singleton entities

	/* flags for outputting entity relation data  */

	SzEntityIncludePossiblySameRelations    // 6 include "possibly same" relationships on entities
	SzEntityIncludePossiblyRelatedRelations // 7 include "possibly related" relationships on entities
	SzEntityIncludeNameOnlyRelations        // 8 include "name only" relationships on entities
	SzEntityIncludeDisclosedRelations       // 9 include "disclosed" relationships on entities

	/* flags for outputting entity feature data  */

	SzEntityIncludeAllFeatures            // 10 include all features for entities
	SzEntityIncludeRepresentativeFeatures // 11 include only representative features on entities

	/* flags for getting extra information about an entity  */

	SzEntityIncludeEntityName           // 12 include the name of the entity
	SzEntityIncludeRecordSummary        // 13 include the record summary of the entity
	SzEntityIncludeRecordData           // 14 include the basic record data for the entity
	SzEntityIncludeRecordMatchingInfo   // 15 include the record matching info for the entity
	SzEntityIncludeRecordJSONData       // 16 include the record json data for the entity
	Bit18                               // 17
	SzEntityIncludeRecordFeatureIDs     // 18 include the features identifiers for the records
	SzEntityIncludeRelatedEntityName    // 19 include the name of the related entities
	SzEntityIncludeRelatedMatchingInfo  // 20 include the record matching info of the related entities
	SzEntityIncludeRelatedRecordSummary // 21 include the record summary of the related entities
	SzEntityIncludeRelatedRecordData    // 22 include the basic record of the related entities

	/* flags for extra feature data  */

	SzEntityIncludeInternalFeatures // 23 include internal features
	SzEntityIncludeFeatureStats     // 24 include statistics on features

	/* flags for finding entity path data  */

	SzFindPathStrictAvoid // 25 excluded entities are still allowed, but not preferred

	/* flags for including search result feature scores  */

	SzIncludeFeatureScores            // 26 include feature scores
	SzSearchIncludeStats              // 27 include statistics from search results
	SzEntityIncludeRecordTypes        // 28 include the record types of the entity
	SzEntityIncludeRelatedRecordTypes // 29 include the record types of the related entities
	SzFindPathIncludeMatchingInfo     // 30 include matching info on entity paths
	SzEntityIncludeRecordUnmappedData // 31 include the record unmapped data for the entity
	SzEntityIncludeFeatureElements    // 32 include feature elements
	SzFindNetworkIncludeMatchingInfo  // 33 include matching info on entity networks
	SzIncludeMatchKeyDetails          // 34 include internal features
	Bit36                             // 35
	Bit37                             // 36
	Bit38                             // 37
	Bit39                             // 38
	Bit40                             // 39
	Bit41                             // 40
	Bit42                             // 41
	Bit43                             // 42
	Bit44                             // 43
	Bit45                             // 44
	Bit46                             // 45
	Bit47                             // 46
	Bit48                             // 47
	Bit49                             // 48
	Bit50                             // 49
	Bit51                             // 50
	Bit52                             // 51
	Bit53                             // 52
	Bit54                             // 53
	Bit55                             // 54
	Bit56                             // 55
	Bit57                             // 56
	Bit58                             // 57
	Bit59                             // 58
	Bit60                             // 59
	Bit61                             // 60
	Bit62                             // 61
	SzWithInfo                        // 62 return "WithInfo" information
)

// Flags used by the Senzing G2Engine.
// These flags combine single-bit flags.
const (
	/* Flags for exporting entity data.  */

	SzExportIncludeAllEntities            = SzExportIncludeMultiRecordEntities | SzExportIncludeSingleRecordEntities                                          // Include all entities.
	SzExportIncludeAllHavingRelationships = SzExportIncludeDisclosed | SzExportIncludeNameOnly | SzExportIncludePossiblyRelated | SzExportIncludePossiblySame // Include all entities with relationships.

	/* Flags for outputting entity relation data  */

	SzEntityIncludeAllRelations = SzEntityIncludeDisclosedRelations | SzEntityIncludeNameOnlyRelations | SzEntityIncludePossiblyRelatedRelations | SzEntityIncludePossiblySameRelations // Include all relationships on entities.

	/* Flags for searching for entities.  */

	SzSearchIncludeAllEntities     = SzSearchIncludeNameOnly | SzSearchIncludePossiblyRelated | SzSearchIncludePossiblySame | SzSearchIncludeResolved
	SzSearchIncludeNameOnly        = SzExportIncludeNameOnly
	SzSearchIncludePossiblyRelated = SzExportIncludePossiblyRelated
	SzSearchIncludePossiblySame    = SzExportIncludePossiblySame
	SzSearchIncludeResolved        = SzExportIncludeMultiRecordEntities

	/* Recommended settings for searching by attributes. */

	SzSearchByAttributesAll           = SzEntityIncludeEntityName | SzEntityIncludeRecordSummary | SzEntityIncludeRepresentativeFeatures | SzSearchIncludeAllEntities | SzIncludeFeatureScores                            // The recommended flag values for searching by attributes, returning all matching entities.
	SzSearchByAttributesMinimalAll    = SzSearchIncludeAllEntities                                                                                                                                                        // Return minimal data with all matches.
	SzSearchByAttributesMinimalStrong = SzSearchIncludePossiblySame | SzSearchIncludeResolved                                                                                                                             // Return minimal data with only the strongest matches.
	SzSearchByAttributesStrong        = SzEntityIncludeEntityName | SzEntityIncludeRecordSummary | SzEntityIncludeRepresentativeFeatures | SzIncludeFeatureScores | SzSearchIncludePossiblySame | SzSearchIncludeResolved // The recommended flag values for searching by attributes, returning only strongly matching entities.

	/* Recommended defaults */

	SzEntityBriefDefaultFlags        = SzEntityIncludeAllRelations | SzEntityIncludeRecordMatchingInfo | SzEntityIncludeRelatedMatchingInfo                                                                                                                                                                                                         // The recommended default flag values for a brief entity result.
	SzEntityDefaultFlags             = SzEntityIncludeAllRelations | SzEntityIncludeEntityName | SzEntityIncludeRecordData | SzEntityIncludeRecordMatchingInfo | SzEntityIncludeRecordSummary | SzEntityIncludeRelatedEntityName | SzEntityIncludeRelatedMatchingInfo | SzEntityIncludeRelatedRecordSummary | SzEntityIncludeRepresentativeFeatures // The recommended default flag values for getting entities.
	SzExportDefaultFlags             = SzEntityDefaultFlags | SzExportIncludeAllEntities                                                                                                                                                                                                                                                            // The recommended default flag values for exporting entities.
	SzFindNetworkDefaultFlags        = SzEntityIncludeEntityName | SzEntityIncludeRecordSummary | SzFindNetworkIncludeMatchingInfo                                                                                                                                                                                                                  // The recommended default flag values for finding entity paths.
	SzFindPathDefaultFlags           = SzEntityIncludeEntityName | SzEntityIncludeRecordSummary | SzFindPathIncludeMatchingInfo                                                                                                                                                                                                                     // The recommended default flag values for finding entity paths.
	SzHowEntityDefaultFlags          = SzIncludeFeatureScores                                                                                                                                                                                                                                                                                       // The recommended default flag values for how-analysis on entities.
	SzRecordDefaultFlags             = SzEntityIncludeRecordJSONData                                                                                                                                                                                                                                                                                // The recommended default flag values for getting records.
	SzSearchByAttributesDefaultFlags = SzSearchByAttributesAll                                                                                                                                                                                                                                                                                      // The recommended default flag values for search-by-attributes
	SzVirtualEntityDefaultFlags      = SzEntityDefaultFlags                                                                                                                                                                                                                                                                                         // The recommended default flag values for virtual-entity-analysis on entities.
	SzWhyEntitiesDefaultFlags        = SzEntityDefaultFlags | SzEntityIncludeFeatureStats | SzEntityIncludeInternalFeatures | SzIncludeFeatureScores
	SzWhyRecordInEntityIDefaultFlags = SzEntityDefaultFlags | SzEntityIncludeFeatureStats | SzEntityIncludeInternalFeatures | SzIncludeFeatureScores // The recommended default flag values for why-analysis on entities.
	SzWhyRecordsDefaultFlags         = SzEntityDefaultFlags | SzEntityIncludeFeatureStats | SzEntityIncludeInternalFeatures | SzIncludeFeatureScores // The recommended default flag values for why-analysis on entities.
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
