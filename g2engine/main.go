package g2engine

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Log message prefix.
const Prefix = "g2engine."

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for g2engine implementations.
var IdMessages = map[int]string{
	1:    "Enter " + Prefix + "AddRecord(%s, %s, %s, %d).",
	2:    "Exit  " + Prefix + "AddRecord(%s, %s, %s, %d) returned (%s, %v).",
	11:   "Enter " + Prefix + "ClearLastException().",
	12:   "Exit  " + Prefix + "ClearLastException() returned (%v).",
	13:   "Enter " + Prefix + "CloseExport(%v).",
	14:   "Exit  " + Prefix + "CloseExport(%v) returned (%v).",
	15:   "Enter " + Prefix + "CountRedoRecords().",
	16:   "Exit  " + Prefix + "CountRedoRecords() returned (%d, %v).",
	17:   "Enter " + Prefix + "DeleteRecord(%s, %s, %d).",
	18:   "Exit  " + Prefix + "DeleteRecord(%s, %s, %d) returned (%s, %v).",
	21:   "Enter " + Prefix + "Destroy().",
	22:   "Exit  " + Prefix + "Destroy() returned (%v).",
	27:   "Enter " + Prefix + "ExportCsvEntityReport(%s, %d).",
	28:   "Exit  " + Prefix + "ExportCsvEntityReport(%s, %d) returned (%v, %v).",
	29:   "Enter " + Prefix + "ExportJsonEntityReport(%d).",
	30:   "Exit  " + Prefix + "ExportJsonEntityReport(%d) returned (%v, %v).",
	31:   "Enter " + Prefix + "FetchNext(%v).",
	32:   "Exit  " + Prefix + "FetchNext(%v) returned (%s, %v).",
	37:   "Enter " + Prefix + "FindNetworkByEntityId(%s, %d, %d, %d, %d).",
	38:   "Exit  " + Prefix + "FindNetworkByEntityId(%s, %d, %d, %d, %d) returned (%s, %v).",
	41:   "Enter " + Prefix + "FindNetworkByRecordId(%s, %d, %d, %d, %d).",
	42:   "Exit  " + Prefix + "FindNetworkByRecordId(%s, %d, %d, %d, %d) returned (%s, %v).",
	45:   "Enter " + Prefix + "FindPathByEntityId(%d, %d, %d, %s, %s, %d).",
	46:   "Exit  " + Prefix + "FindPathByEntityId(%d, %d, %d, %s, %s, %d) returned (%s, %v).",
	49:   "Enter " + Prefix + "FindPathByRecordId(%s, %s, %s, %s, %d, %s, %s, %d).",
	50:   "Exit  " + Prefix + "FindPathByRecordId(%s, %s, %s, %s, %d, %s, %s, %d) returned (%s, %v).",
	69:   "Enter " + Prefix + "GetActiveConfigId().",
	70:   "Exit  " + Prefix + "GetActiveConfigId() returned (%d, %v).",
	71:   "Enter " + Prefix + "GetEntityByEntityId(%d, %d).",
	72:   "Exit  " + Prefix + "GetEntityByEntityId(%d, %d) returned (%s, %v).",
	75:   "Enter " + Prefix + "GetEntityByRecordId(%s, %s, %d).",
	76:   "Exit  " + Prefix + "GetEntityByRecordId(%s, %s, %d) returned (%s, %v).",
	79:   "Enter " + Prefix + "GetLastException().",
	80:   "Exit  " + Prefix + "GetLastException() returned (%s, %v).",
	81:   "Enter " + Prefix + "GetLastExceptionCode().",
	82:   "Exit  " + Prefix + "GetLastExceptionCode() returned (%d, %v).",
	83:   "Enter " + Prefix + "GetRecord(%s, %s, %d).",
	84:   "Exit  " + Prefix + "GetRecord(%s, %s, %d) returned (%s, %v).",
	87:   "Enter " + Prefix + "GetRedoRecord().",
	88:   "Exit  " + Prefix + "GetRedoRecord() returned (%s, %v).",
	89:   "Enter " + Prefix + "GetRepositoryLastModifiedTime().",
	90:   "Exit  " + Prefix + "GetRepositoryLastModifiedTime() returned (%d, %v).",
	91:   "Enter " + Prefix + "GetVirtualEntityByRecordId(%s, %d).",
	92:   "Exit  " + Prefix + "GetVirtualEntityByRecordId(%s, %d) returned (%s, %v).",
	95:   "Enter " + Prefix + "HowEntityByEntityId(%d, %d).",
	96:   "Exit  " + Prefix + "HowEntityByEntityId(%d, %d) returned (%s, %v).",
	99:   "Enter " + Prefix + "Initialize(%s, %s, %d, %d).",
	100:  "Exit  " + Prefix + "Initialize(%s, %s, %d, %d) returned (%v).",
	103:  "Enter " + Prefix + "PrimeEngine().",
	104:  "Exit  " + Prefix + "PrimeEngine() returned (%v).",
	107:  "Enter " + Prefix + "ProcessRedoRecord(%s, %d).",
	108:  "Exit  " + Prefix + "ProcessRedoRecord(%s, %d) returned (%s, %v).",
	119:  "Enter " + Prefix + "ReevaluateEntity(%d, %d).",
	120:  "Exit  " + Prefix + "ReevaluateEntity(%d, %d) returned (%s, %v).",
	123:  "Enter " + Prefix + "ReevaluateRecord(%s, %s, %d).",
	124:  "Exit  " + Prefix + "ReevaluateRecord(%s, %s, %d) returned (%s, %v).",
	127:  "Enter " + Prefix + "Reinitialize(%d).",
	128:  "Exit  " + Prefix + "Reinitialize(%d) returned (%v).",
	129:  "Enter " + Prefix + "ReplaceRecord(%s, %s, %s, %d).",
	130:  "Exit  " + Prefix + "ReplaceRecord(%s, %s, %s, %d) returned (%s, %v).",
	133:  "Enter " + Prefix + "SearchByAttributes(%s, %s, %d).",
	134:  "Exit  " + Prefix + "SearchByAttributes(%s, %s, %d) returned (%s, %v).",
	137:  "Enter " + Prefix + "SetLogLevel(%v).",
	138:  "Exit  " + Prefix + "SetLogLevel(%v) returned (%v).",
	139:  "Enter " + Prefix + "GetStats().",
	140:  "Exit  " + Prefix + "GetStats() returned (%s, %v).",
	141:  "Enter " + Prefix + "WhyEntities(%d, %d, %d).",
	142:  "Exit  " + Prefix + "WhyEntities(%d, %d, %d) returned (%s, %v).",
	143:  "Enter " + Prefix + "WhyRecordInEntity(%s, %s, %d).",
	144:  "Exit  " + Prefix + "WhyRecordInEntity(%s, %s, %d) returned (%s, %v).",
	153:  "Enter " + Prefix + "WhyRecords(%s, %s, %s, %s, %d).",
	154:  "Exit  " + Prefix + "WhyRecords(%s, %s, %s, %s, %d) returned (%s, %v).",
	157:  "Enter " + Prefix + "RegisterObserver(%s).",
	158:  "Exit  " + Prefix + "RegisterObserver(%s) returned (%v).",
	159:  "Enter " + Prefix + "UnregisterObserver(%s).",
	160:  "Exit  " + Prefix + "UnregisterObserver(%s) returned (%v).",
	161:  "Enter " + Prefix + "GetSdkId().",
	162:  "Exit  " + Prefix + "GetSdkId() returned (%s).",
	163:  "Enter " + Prefix + "ExportCsvEntityReportIterator(%s, %d).",
	164:  "Exit  " + Prefix + "ExportCsvEntityReportIterator(%s, %d) returned (%v, %v).",
	165:  "Enter " + Prefix + "ExportJsonEntityReportIterator(%d).",
	166:  "Exit  " + Prefix + "ExportJsonEntityReportIterator(%d) returned (%v, %v).",
	4001: Prefix + "G2_addRecord(%s, %s, %s, %s) failed. Return code: %d",
	4002: Prefix + "G2_addRecordWithInfo(%s, %s, %s, %s, %d) failed. Return code: %d",
	4006: Prefix + "G2_closeExport(%v) failed. Return code: %d",
	4007: Prefix + "G2_deleteRecord(%s, %s, %s) failed. Return code: %d",
	4008: Prefix + "G2_deleteRecordWithInfo(%s, %s, %s, %d) failed. Return code: %d",
	4009: Prefix + "G2_destroy() failed. Return code: %d",
	4010: Prefix + "G2_exportConfigAndConfigID() failed. Return code: %d",
	4011: Prefix + "G2_exportConfig() failed. Return code: %d",
	4012: Prefix + "G2_exportCSVEntityReport(%s, %d) failed. Return code: %d",
	4013: Prefix + "G2_exportJSONEntityReport(%d) failed. Return code: %d",
	4014: Prefix + "G2_fetchNext(%v) failed. Return code: %d",
	4015: Prefix + "G2_findInterestingEntitiesByEntityID(%d, %d) failed. Return code: %d",
	4016: Prefix + "G2_findInterestingEntitiesByRecordID(%s, %s, %d) failed. Return code: %d",
	4017: Prefix + "G2_findNetworkByEntityID(%s, %d, %d, %d) failed. Return code: %d",
	4018: Prefix + "G2_findNetworkByEntityID_V2(%s, %d, %d, %d, %d) failed. Return code: %d",
	4019: Prefix + "G2_findNetworkByRecordID(%s, %d, %d, %d) failed. Return code: %d",
	4020: Prefix + "G2_findNetworkByRecordID_V2(%s, %d, %d, %d, %d) failed. Return code: %d",
	4021: Prefix + "G2_findPathByEntityID(%d, %d, %d) failed. Return code: %d",
	4022: Prefix + "G2_findPathByEntityID_V2(%d, %d, %d, %d) failed. Return code: %d",
	4023: Prefix + "G2_findPathByRecordID(%s, %s, %s, %s, %d) failed. Return code: %d",
	4024: Prefix + "G2_findPathByRecordID_V2(%s, %s, %s, %s, %d, %d) failed. Return code: %d",
	4025: Prefix + "G2_findPathExcludingByEntityID(%d, %d, %d, %s) failed. Return code: %d",
	4026: Prefix + "G2_findPathExcludingByEntityID_V2(%d, %d, %d, %s, %d) failed. Return code: %d",
	4027: Prefix + "G2_findPathExcludingByRecordID(%s, %s, %s, %s %d, %s) failed. Return code: %d",
	4028: Prefix + "G2_findPathExcludingByRecordID_V2(%s, %s, %s, %s %d, %s, %d) failed. Return code: %d",
	4029: Prefix + "G2_findPathIncludingSourceByEntityID(%d, %d, %d, %s, %s) failed. Return code: %d",
	4030: Prefix + "G2_findPathIncludingSourceByEntityID_V2(%d, %d, %d, %s, %s, %d) failed. Return code: %d",
	4031: Prefix + "G2_findPathIncludingSourceByRecordID(%s, %s, %s, %s %d, %s, %s) failed. Return code: %d",
	4032: Prefix + "G2_findPathIncludingSourceByRecordID_V2(%s, %s, %s, %s %d, %s, %s, %d) failed. Return code: %d",
	4033: Prefix + "G2_getActiveConfigID() failed. Return code: %d",
	4034: Prefix + "G2_getEntityByEntityID(%d) failed. Return code: %d",
	4035: Prefix + "G2_getEntityByEntityID_V2(%d, %d) failed. Return code: %d",
	4036: Prefix + "G2_getEntityByRecordID(%s, %s) failed. Return code: %d",
	4037: Prefix + "G2_getEntityByRecordID_V2(%s, %s, %d) failed. Return code: %d",
	4038: Prefix + "G2_getLastException() failed. Return code: %d",
	4039: Prefix + "G2_getRecord(%s, %s) failed. Return code: %d",
	4040: Prefix + "G2_getRecord_V2(%s, %s, %d) failed. Return code: %d",
	4041: Prefix + "G2_getRedoRecord() failed. Return code: %d",
	4042: Prefix + "G2_getRepositoryLastModifiedTime() failed. Return code: %d",
	4043: Prefix + "G2_getVirtualEntityByRecordID(%s) failed. Return code: %d",
	4044: Prefix + "G2_getVirtualEntityByRecordID_V2(%s, %d) failed. Return code: %d",
	4045: Prefix + "G2_howEntityByEntityID(%d) failed. Return code: %d",
	4046: Prefix + "G2_howEntityByEntityID_V2(%d, %d) failed. Return code: %d",
	4047: Prefix + "G2_init(%s, %s, %d) failed. Return code: %d",
	4048: Prefix + "G2_initWithConfigID(%s, %s, %d, %d) failed. Return code: %d",
	4049: Prefix + "G2_primeEngine() failed. Return code: %d",
	4051: Prefix + "G2_processRedoRecord() failed. Return code: %d",
	4052: Prefix + "G2_processRedoRecordWithInfo(%d) failed. Return code: %d",
	4056: Prefix + "G2Diagnostic_purgeRepository() failed. Return code: %d",
	4057: Prefix + "G2_reevaluateEntity(%d, %d) failed. Return code: %d",
	4058: Prefix + "G2_reevaluateEntityWithInfo(%d, %d) failed. Return code: %d",
	4059: Prefix + "G2_reevaluateRecord(%s, %s, %d) failed. Return code: %d",
	4060: Prefix + "G2_reevaluateRecordWithInfo(%s, %s, %d) failed. Return code: %d",
	4061: Prefix + "G2_reinit(%d) failed. Return code: %d",
	4062: Prefix + "G2_replaceRecord(%s, %s, %s, %s) failed. Return code: %d",
	4063: Prefix + "G2_replaceRecordWithInfo(%s, %s, %s, %s, %d) failed. Return code: %d",
	4064: Prefix + "G2_searchByAttributes(%s) failed. Return code: %d",
	4065: Prefix + "G2_searchByAttributes_V2(%s, %d) failed. Return code: %d",
	4066: Prefix + "G2_stats() failed. Return code: %d",
	4067: Prefix + "G2_whyEntities(%d, %d) failed. Return code: %d",
	4068: Prefix + "G2_whyEntities_V2(%d, %d, %d) failed. Return code: %d",
	4073: Prefix + "G2_whyRecords(%s, %s, %s, %s) failed. Return code: %d",
	4074: Prefix + "G2_whyRecords_V2(%s, %s, %s, %s, %d) failed. Return code: %d",
	5901: "During test setup, call to messagelogger.NewSenzingApiLogger() failed.",
	5902: "During test setup, call to g2engineconfigurationjson.BuildSimpleSystemConfigurationJsonViaMap() failed.",
	5903: "During test setup, call to g2engine.Initialize() failed.",
	5904: "During test setup, call to g2diagnostic.PurgeRepository() failed.",
	5905: "During test setup, call to g2engine.Destroy() failed.",
	5906: "During test setup, call to g2config.Initialize() failed.",
	5907: "During test setup, call to g2config.Create() failed.",
	5908: "During test setup, call to g2config.AddDataSource() failed.",
	5909: "During test setup, call to g2config.GetJsonString() failed.",
	5910: "During test setup, call to g2config.Close() failed.",
	5911: "During test setup, call to g2config.Destroy() failed.",
	5912: "During test setup, call to g2configmgr.Initialize() failed.",
	5913: "During test setup, call to g2configmgr.AddConfig() failed.",
	5914: "During test setup, call to g2configmgr.SetDefaultConfigID() failed.",
	5915: "During test setup, call to g2configmgr.Destroy() failed.",
	5916: "During test setup, call to g2engine.Initialize() failed.",
	5917: "During test setup, call to g2engine.AddRecord() failed.",
	5918: "During test setup, call to g2engine.Destroy() failed.",
	5920: "During test setup, call to setupSenzingConfig() failed.",
	5921: "During test setup, call to setupPurgeRepository() failed.",
	5922: "During test setup, call to setupAddRecords() failed.",
	5931: "During test setup, call to g2engine.Initialize() failed.",
	5932: "During test setup, call to g2diagnostic.PurgeRepository() failed.",
	5933: "During test setup, call to g2engine.Destroy() failed.",
	8001: Prefix + "AddRecord",
	8006: Prefix + "CloseExport",
	8007: Prefix + "CountRedoRecords",
	8008: Prefix + "DeleteRecord",
	8010: Prefix + "Destroy",
	8013: Prefix + "ExportCSVEntityReport",
	8014: Prefix + "ExportJSONEntityReport",
	8015: Prefix + "FetchNext",
	8018: Prefix + "FindNetworkByEntityID",
	8020: Prefix + "FindNetworkByRecordID",
	8022: Prefix + "FindPathByEntityID",
	8024: Prefix + "FindPathByRecordID",
	8034: Prefix + "GetActiveConfigID",
	8035: Prefix + "GetEntityByEntityID",
	8037: Prefix + "GetEntityByRecordID",
	8039: Prefix + "GetRecord",
	8041: Prefix + "GetRedoRecord",
	8042: Prefix + "GetRepositoryLastModifiedTime",
	8043: Prefix + "GetVirtualEntityByRecordID",
	8045: Prefix + "HowEntityByEntityID",
	8047: Prefix + "Initialize",
	8049: Prefix + "PrimeEngine",
	8051: Prefix + "ProcessRedoRecord",
	8057: Prefix + "ReevaluateEntity",
	8059: Prefix + "ReevaluateRecord",
	8061: Prefix + "Reinitialize",
	8062: Prefix + "ReplaceRecord",
	8064: Prefix + "SearchByAttributes",
	8066: Prefix + "GetStats",
	8067: Prefix + "WhyEntities",
	8073: Prefix + "WhyRecords",
	8074: Prefix + "WhyRecordInEntity",
	8075: Prefix + "GetSdkId",
	8076: Prefix + "RegisterObserver",
	8077: Prefix + "SetLogLevel",
	8078: Prefix + "UnregisterObserver",
	8079: Prefix + "ExportCSVEntityReport",
	8080: Prefix + "ExportJSONEntityReport",
}

// Status strings for specific g2engine messages.
var IdStatuses = map[int]string{}
