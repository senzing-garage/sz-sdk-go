#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>
#include "libg2.h"

typedef void* (*resize_buffer_type)(void*, size_t);

struct G2_addRecordWithInfoWithReturnedRecordID_result {
    char* recordID;
    char* withInfo;
    int returnCode;
};

struct G2_exportConfigAndConfigID_result {
    long long configID;
    char* config;
    int returnCode;
};

struct G2_exportCSVEntityReport_result {
    void* exportHandle;
    int returnCode;
};
struct G2_exportJSONEntityReport_result {
    void* exportHandle;
    int returnCode;
};
struct G2_fetchNext_result {
    char* response;
    int returnCode;
};

struct G2_getActiveConfigID_result {
    long long configID;
    int returnCode;
};

struct G2_getRedoRecord_result {
    char* response;
    int returnCode;
};

struct G2_processRedoRecord_result {
    char* response;
    int returnCode;
};

struct G2_processRedoRecordWithInfo_result {
    char* response;
    char* withInfo;
    int returnCode;
};

void* G2_resizeStringBuffer(void* ptr, size_t size);
char* G2_addRecordWithInfo_helper(const char* dataSourceCode, const char* recordID, const char* jsonData, const char* loadID, const long long flags);
struct G2_addRecordWithInfoWithReturnedRecordID_result G2_addRecordWithInfoWithReturnedRecordID_helper(const char* dataSourceCode, const char* jsonData, const char* loadID, const long long flags);
char* G2_checkRecord_helper(const char* record, const char* recordQueryList);
int G2_closeExport_helper(uintptr_t responseHandle);
char* G2_deleteRecordWithInfo_helper(const char* dataSourceCode, const char* recordID, const char* loadID, const long long flags);
struct G2_exportConfigAndConfigID_result G2_exportConfigAndConfigID_helper();
char* G2_exportConfig_helper();
struct G2_exportCSVEntityReport_result G2_exportCSVEntityReport_helper(const char* csvColumnList, const long long flags);
struct G2_exportJSONEntityReport_result G2_exportJSONEntityReport_helper(const long long flags);
char* G2_findInterestingEntitiesByEntityID_helper(long long entityID, long long flags);
char* G2_findInterestingEntitiesByRecordID_helper(const char* dataSourceCode, const char* recordID, long long flags);
char* G2_findNetworkByEntityID_helper(const char* entityList, const int maxDegree, const int buildOutDegree, const int maxEntities);
char* G2_findNetworkByEntityID_V2_helper(const char* entityList, const int maxDegree, const int buildOutDegree, const int maxEntities, long long flags);
char* G2_findNetworkByRecordID_helper(const char* recordList, const int maxDegree, const int buildOutDegree, const int maxEntities);
char* G2_findNetworkByRecordID_V2_helper(const char* recordList, const int maxDegree, const int buildOutDegree, const int maxEntities, const long long flags);
char* G2_findPathByEntityID_helper(const long long entityID1, const long long entityID2, const int maxDegree);
char* G2_findPathByEntityID_V2_helper(const long long entityID1, const long long entityID2, const int maxDegree, const long long flags);
char* G2_findPathByRecordID_helper(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree);
char* G2_findPathByRecordID_V2_helper(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const long long flags);
char* G2_findPathExcludingByEntityID_helper(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities);
char* G2_findPathExcludingByEntityID_V2_helper(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const long long flags);
char* G2_findPathExcludingByRecordID_helper(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords);
char* G2_findPathExcludingByRecordID_V2_helper(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const long long flags);
char* G2_findPathIncludingSourceByEntityID_helper(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const char* requiredDsrcs);
char* G2_findPathIncludingSourceByEntityID_V2_helper(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const char* requiredDsrcs, const long long flags);
char* G2_findPathIncludingSourceByRecordID_helper(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const char* requiredDsrcs);
char* G2_findPathIncludingSourceByRecordID_V2_helper(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const char* requiredDsrcs, const long long flags);
struct G2_fetchNext_result G2_fetchNext_helper(uintptr_t exportHandle);
struct G2_getActiveConfigID_result G2_getActiveConfigID_helper();
char* G2_getEntityByEntityID_helper(const long long entityID);
char* G2_getEntityByEntityID_V2_helper(const long long entityID, const long long flags);
char* G2_getEntityByRecordID_helper(const char* dataSourceCode, const char* recordID);
char* G2_getEntityByRecordID_V2_helper(const char* dataSourceCode, const char* recordID, const long long flags);
char* G2_getRecord_helper(const char* dataSourceCode, const char* recordID);
char* G2_getRecord_V2_helper(const char* dataSourceCode, const char* recordID, const long long flags);
struct G2_getRedoRecord_result G2_getRedoRecord_helper();
long long G2_getRepositoryLastModifiedTime_helper();
char* G2_getVirtualEntityByRecordID_helper(const char* recordList);
char* G2_getVirtualEntityByRecordID_V2_helper(const char* recordList, const long long flags);
char* G2_howEntityByEntityID_helper(const long long entityID);
char* G2_howEntityByEntityID_V2_helper(const long long entityID, const long long flags);
struct G2_processRedoRecord_result G2_processRedoRecord_helper();
struct G2_processRedoRecordWithInfo_result G2_processRedoRecordWithInfo_helper(const long long flags);
char* G2_processWithInfo_helper(const char* record, const long long flags);
char* G2_processWithResponse_helper(const char* record);
char* G2_processWithResponseResize_helper(const char* record);
char* G2_reevaluateEntityWithInfo_helper(const long long entityID, const long long flags);
char* G2_reevaluateRecordWithInfo_helper(const char* dataSourceCode, const char* recordID, const long long flags);
char* G2_replaceRecordWithInfo_helper(const char* dataSourceCode, const char* recordID, const char* jsonData, const char* loadID, const long long flags);
char* G2_searchByAttributes_helper(const char* jsonData);
char* G2_searchByAttributes_V2_helper(const char* jsonData, const long long flags);
char* G2_stats_helper();
char* G2_whyEntities_helper(const long long entityID1, const long long entityID2);
char* G2_whyEntities_V2_helper(const long long entityID1, const long long entityID2, const long long flags);
char* G2_whyEntityByEntityID_helper(const long long entityID1);
char* G2_whyEntityByEntityID_V2_helper(const long long entityID1, const long long flags);
char* G2_whyEntityByRecordID_helper(const char* dataSourceCode, const char* recordID);
char* G2_whyEntityByRecordID_V2_helper(const char* dataSourceCode, const char* recordID, const long long flags);
char* G2_whyRecords_helper(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2);
char* G2_whyRecords_V2_helper(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const long long flags);
