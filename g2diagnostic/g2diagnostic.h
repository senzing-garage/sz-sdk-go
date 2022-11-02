#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include "libg2diagnostic.h"

//typedef void* EntityListBySizeHandle;
typedef void* (*resize_buffer_type)(void*, size_t);

void* G2Diagnostic_resizeStringBuffer(void* ptr, size_t size);
char* G2Diagnostic_checkDBPerf_helper(int secondsToRun);
int G2Diagnostic_closeEntityListBySize_helper(uintptr_t entityListBySizeHandle);
int G2Diagnostic_fetchNextEntityBySize_helper(uintptr_t configHandle, char* responseBuf, const size_t bufSize);
char* G2Diagnostic_findEntitiesByFeatureIDs_helper(const char* features);
char* G2Diagnostic_getDataSourceCounts_helper();
char* G2Diagnostic_getDBInfo_helper();
char* G2Diagnostic_getEntityDetails_helper(const long long entityID, const int includeInternalFeatures);
void* G2Diagnostic_getEntityListBySize_helper(const size_t entitySize);
char* G2Diagnostic_getEntityResume_helper(const long long entityID);
char* G2Diagnostic_getEntitySizeBreakdown_helper(const size_t minimumEntitySize, const int includeInternalFeatures);
char* G2Diagnostic_getFeature_helper(const long long libFeatID);
char* G2Diagnostic_getGenericFeatures_helper(const char* featureType, const size_t maximumEstimatedCount);
char* G2Diagnostic_getMappingStatistics_helper(const int includeInternalFeatures);
char* G2Diagnostic_getRelationshipDetails_helper(const long long relationshipID, const int includeInternalFeatures);
char* G2Diagnostic_getResolutionStatistics_helper();
