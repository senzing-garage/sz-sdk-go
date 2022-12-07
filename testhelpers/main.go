package testhelpers

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// A list of data sources.
var TestDataSources = []struct {
	Data string
}{
	{
		Data: `{"DSRC_CODE": "EXAMPLE_DATA_SOURCE"}`,
	},
}

// Must match value in sys_cfg.config_data_id.
var TestConfigDataId = 4175236977

// A list of test records.
var TestRecords = []struct {
	DataSource string
	Id         string
	Data       string
	LoadId     string
}{
	{
		DataSource: "EXAMPLE_DATA_SOURCE",
		Id:         "9001",
		Data:       `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "EXAMPLE_DATA_SOURCE", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "9001", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "JOHNSON", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "TEST",
	},
	{
		DataSource: "EXAMPLE_DATA_SOURCE",
		Id:         "9002",
		Data:       `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "EXAMPLE_DATA_SOURCE", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "9002", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "JOHNSON", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "TEST",
	},
	{
		DataSource: "EXAMPLE_DATA_SOURCE",
		Id:         "9003",
		Data:       `{"ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "ENTITY_TYPE": "EXAMPLE_DATA_SOURCE", "GENDER": "M", "srccode": "MDMPER", "RECORD_ID": "9003", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "Smith", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "TEST",
	},
}
