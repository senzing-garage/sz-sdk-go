package senzing

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

/*
A flag with no flags set.
*/
const (
	SzNoFlags int64 = 0 // Used to indicate no flags being set
)

/*
Flags used by the Senzing SzEngine.
These flags are single-bit flags.
BitNN is for bits numbered as 1..64 (not 0..63).  Comments are for bits numbered as 0..63
The are referred to as "SzXxx" flags.
*/
const (
	SzExportIncludeMultiRecordEntities  int64 = 1 << iota // 0 Include entities with "resolved" relationships
	SzExportIncludePossiblySame                           // 1 Include entities with "possibly same" relationships
	SzExportIncludePossiblyRelated                        // 2 Include entities with "possibly related" relationships
	SzExportIncludeNameOnly                               // 3 Include entities with "name only" relationships
	SzExportIncludeDisclosed                              // 4 Include entities with "disclosed" relationships
	SzExportIncludeSingleRecordEntities                   // 5 Include singleton entities

	/* flags for outputting entity relation data  */

	SzEntityIncludePossiblySameRelations    // 6 Include "possibly same" relationships on entities
	SzEntityIncludePossiblyRelatedRelations // 7 Include "possibly related" relationships on entities
	SzEntityIncludeNameOnlyRelations        // 8 Include "name only" relationships on entities
	SzEntityIncludeDisclosedRelations       // 9 Include "disclosed" relationships on entities

	/* flags for outputting entity feature data  */

	SzEntityIncludeAllFeatures            // 10 Include all features for entities
	SzEntityIncludeRepresentativeFeatures // 11 Include only representative features on entities

	/* flags for getting extra information about an entity  */

	SzEntityIncludeEntityName           // 12 Include the name of the entity
	SzEntityIncludeRecordSummary        // 13 Include the record summary of the entity
	SzEntityIncludeRecordData           // 14 Include the basic record data for the entity
	SzEntityIncludeRecordMatchingInfo   // 15 Include the record matching info for the entity
	SzEntityIncludeRecordJSONData       // 16 Include the record json data for the entity
	Bit18                               // 17
	SzEntityIncludeRecordFeatureIDs     // 18 Include the features identifiers for the records
	SzEntityIncludeRelatedEntityName    // 19 Include the name of the related entities
	SzEntityIncludeRelatedMatchingInfo  // 20 Include the record matching info of the related entities
	SzEntityIncludeRelatedRecordSummary // 21 Include the record summary of the related entities
	SzEntityIncludeRelatedRecordData    // 22 Include the basic record of the related entities

	/* flags for extra feature data  */

	SzEntityIncludeInternalFeatures // 23 Include internal features
	SzEntityIncludeFeatureStats     // 24 Include statistics on features

	/* flags for finding entity path data  */

	SzFindPathStrictAvoid // 25 excluded entities are still allowed, but not preferred

	/* flags for including search result feature scores  */

	SzIncludeFeatureScores              // 26 Include feature scores
	SzSearchIncludeStats                // 27 Include statistics from search results
	SzEntityIncludeRecordTypes          // 28 Include the record types of the entity
	SzEntityIncludeRelatedRecordTypes   // 29 Include the record types of the related entities
	SzFindPathIncludeMatchingInfo       // 30 Include matching info on entity paths
	SzEntityIncludeRecordUnmappedData   // 31 Include the record unmapped data for the entity
	SzSearchIncludeAllCandidates        // 32
	SzFindNetworkIncludeMatchingInfo    // 33 Include matching info on entity networks
	SzIncludeMatchKeyDetails            // 34 Include internal features
	SzEntityIncludeRecordFeatureDetails // 35 Include attributes section
	SzEntityIncludeRecordFeatureStats   // 36 Include the feature statistics
	SzSearchIncludeRequest              // 37
	SzSearchIncludeRequestDetails       // 38

	/* Reserved  */

	Bit40 // 39
	Bit41 // 40
	Bit42 // 41
	Bit43 // 42
	Bit44 // 43
	Bit45 // 44
	Bit46 // 45
	Bit47 // 46
	Bit48 // 47
	Bit49 // 48
	Bit50 // 49
	Bit51 // 50
	Bit52 // 51
	Bit53 // 52
	Bit54 // 53
	Bit55 // 54
	Bit56 // 55
	Bit57 // 56
	Bit58 // 57
	Bit59 // 58
	Bit60 // 59
	Bit61 // 60
	Bit62 // 61

	/* Reserved for use by SDKs */

	SzWithInfo // 62 return "WithInfo" information
)

// Compound Flags used by the Senzing SzEngine.

/*
Flags for exporting entity data.
*/
const (
	SzExportIncludeAllEntities            = SzExportIncludeMultiRecordEntities | SzExportIncludeSingleRecordEntities                                          // Include all entities.
	SzExportIncludeAllHavingRelationships = SzExportIncludeDisclosed | SzExportIncludeNameOnly | SzExportIncludePossiblyRelated | SzExportIncludePossiblySame // Include all entities with relationships.
)

/*
Flags for outputting entity relation data
*/
const (
	SzEntityIncludeAllRelations = SzEntityIncludeDisclosedRelations | SzEntityIncludeNameOnlyRelations | SzEntityIncludePossiblyRelatedRelations | SzEntityIncludePossiblySameRelations // Include all relationships on entities.
)

/*
Flags for searching for entities.
*/
const (
	SzSearchIncludeAllEntities     = SzSearchIncludeNameOnly | SzSearchIncludePossiblyRelated | SzSearchIncludePossiblySame | SzSearchIncludeResolved
	SzSearchIncludeNameOnly        = SzExportIncludeNameOnly
	SzSearchIncludePossiblyRelated = SzExportIncludePossiblyRelated
	SzSearchIncludePossiblySame    = SzExportIncludePossiblySame
	SzSearchIncludeResolved        = SzExportIncludeMultiRecordEntities
)

/*
Recommended settings for searching by attributes.
*/
const (
	SzSearchByAttributesAll           = SzEntityIncludeEntityName | SzEntityIncludeRecordSummary | SzEntityIncludeRepresentativeFeatures | SzSearchIncludeAllEntities | SzIncludeFeatureScores                            // The recommended flag values for searching by attributes, returning all matching entities.
	SzSearchByAttributesMinimalAll    = SzSearchIncludeAllEntities                                                                                                                                                        // Return minimal data with all matches.
	SzSearchByAttributesMinimalStrong = SzSearchIncludePossiblySame | SzSearchIncludeResolved                                                                                                                             // Return minimal data with only the strongest matches.
	SzSearchByAttributesStrong        = SzEntityIncludeEntityName | SzEntityIncludeRecordSummary | SzEntityIncludeRepresentativeFeatures | SzIncludeFeatureScores | SzSearchIncludePossiblySame | SzSearchIncludeResolved // The recommended flag values for searching by attributes, returning only strongly matching entities.
)

/*
Recommended defaults
*/
const (
	SzEntityBriefDefaultFlags        = SzEntityIncludeAllRelations | SzEntityIncludeRecordMatchingInfo | SzEntityIncludeRelatedMatchingInfo // The recommended default flag values for a brief entity result.
	SzEntityCoreFlags                = SzEntityIncludeEntityName | SzEntityIncludeRecordData | SzEntityIncludeRecordMatchingInfo | SzEntityIncludeRecordSummary | SzEntityIncludeRepresentativeFeatures
	SzEntityDefaultFlags             = SzEntityCoreFlags | SzEntityIncludeAllRelations | SzEntityIncludeRecordSummary | SzEntityIncludeRelatedEntityName | SzEntityIncludeRelatedMatchingInfo // The recommended default flag values for getting entities.
	SzExportDefaultFlags             = SzEntityDefaultFlags | SzExportIncludeAllEntities                                                                                                      // The recommended default flag values for exporting entities.
	SzFindNetworkDefaultFlags        = SzEntityIncludeEntityName | SzEntityIncludeRecordSummary | SzFindNetworkIncludeMatchingInfo                                                            // The recommended default flag values for finding entity paths.
	SzFindPathDefaultFlags           = SzEntityIncludeEntityName | SzEntityIncludeRecordSummary | SzFindPathIncludeMatchingInfo                                                               // The recommended default flag values for finding entity paths.
	SzHowEntityDefaultFlags          = SzIncludeFeatureScores                                                                                                                                 // The recommended default flag values for how-analysis on entities.
	SzRecordDefaultFlags             = SzEntityIncludeRecordJSONData                                                                                                                          // The recommended default flag values for getting records.
	SzSearchByAttributesDefaultFlags = SzSearchByAttributesAll                                                                                                                                // The recommended default flag values for search-by-attributes
	SzVirtualEntityDefaultFlags      = SzEntityCoreFlags                                                                                                                                      // The recommended default flag values for virtual-entity-analysis on entities.
	SzWhyEntitiesDefaultFlags        = SzEntityDefaultFlags | SzEntityIncludeFeatureStats | SzEntityIncludeInternalFeatures | SzIncludeFeatureScores
	SzWhyRecordInEntityIDefaultFlags = SzEntityDefaultFlags | SzEntityIncludeFeatureStats | SzEntityIncludeInternalFeatures | SzIncludeFeatureScores // The recommended default flag values for why-analysis on entities.
	SzWhyRecordsDefaultFlags         = SzEntityDefaultFlags | SzEntityIncludeFeatureStats | SzEntityIncludeInternalFeatures | SzIncludeFeatureScores // The recommended default flag values for why-analysis on entities.
)

/*
Function Flags returns a flag that is the combination of input parameter flags.

Input
  - flags: A list of "SzXxx" flags

Output
  - An int64 combining all input "SzXxx" flags.
*/
func Flags(flags ...int64) int64 {
	var result int64
	for _, flag := range flags {
		result |= flag
	}
	return result
}
