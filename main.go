package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/senzing/g2-sdk-go/g2diagnostic"
	"github.com/senzing/g2-sdk-go/g2engine"
	"github.com/senzing/go-helpers/g2engineconfigurationjson"
	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messageid"
	"github.com/senzing/go-logging/messagelevel"
	"github.com/senzing/go-logging/messagelogger"
	"github.com/senzing/go-logging/messagestatus"
	"github.com/senzing/go-logging/messagetext"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdTemplate = "senzing-9999%04d"

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var Messages = map[int]string{
	1:    "%s",
	2:    "WithInfo: %s",
	2999: "Cannot retrieve last error message.",
}

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

// ----------------------------------------------------------------------------
// Internal methods - names begin with lower case
// ----------------------------------------------------------------------------

func getG2diagnostic(ctx context.Context) (g2diagnostic.G2diagnostic, error) {
	g2diagnostic := g2diagnostic.G2diagnosticImpl{}

	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		return &g2diagnostic, err
	}

	err = g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
	return &g2diagnostic, err
}

func getG2engine(ctx context.Context) (g2engine.G2engine, error) {
	g2engine := g2engine.G2engineImpl{}

	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		return &g2engine, err
	}

	err = g2engine.Init(ctx, moduleName, iniParams, verboseLogging)
	return &g2engine, err
}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	ctx := context.TODO()

	// Randomize random number generator.

	rand.Seed(time.Now().UnixNano())

	// Configure the "log" standard library.

	// log.SetFlags(log.Llongfile | log.Ldate | log.Lmicroseconds | log.LUTC)
	log.SetFlags(log.LstdFlags)

	// Configure messagelogger

	messageFormat := &messageformat.MessageFormatJson{}
	messageIdTemplate := &messageid.MessageIdTemplated{
		MessageIdTemplate: MessageIdTemplate,
	}
	messageLevel := &messagelevel.MessageLevelByIdRange{
		IdRanges: map[int]logger.Level{
			0000: logger.LevelInfo,
			1000: logger.LevelWarn,
			2000: logger.LevelError,
			3000: logger.LevelDebug,
			4000: logger.LevelTrace,
			5000: logger.LevelFatal,
			6000: logger.LevelPanic,
		},
	}
	messageStatus := &messagestatus.MessageStatusByIdRange{
		IdRanges: map[int]string{
			0000: logger.LevelInfoName,
			1000: logger.LevelWarnName,
			2000: logger.LevelErrorName,
			3000: logger.LevelDebugName,
			4000: logger.LevelTraceName,
			5000: logger.LevelFatalName,
			6000: logger.LevelPanicName,
		},
	}
	messageText := &messagetext.MessageTextTemplated{
		IdMessages: Messages,
	}
	logger, _ := messagelogger.New(messageFormat, messageIdTemplate, messageLevel, messageStatus, messageText, messagelogger.LevelInfo)

	// Test logger.

	programmMetadataMap := map[string]interface{}{
		"ProgramName":    programName,
		"BuildVersion":   buildVersion,
		"BuildIteration": buildIteration,
	}

	logger.Log(1, "Just a test of logging", programmMetadataMap)

	// Work with G2diagnostic.

	g2diagnostic, err := getG2diagnostic(ctx)
	if err != nil {
		logger.Log(2000, err)
	}

	// Check trace

	actual, err := g2diagnostic.GetPhysicalCores(ctx)
	if err != nil {
		logger.Log(2001, err)
	}
	fmt.Printf("Physical cores: %d\n", actual)

	// g2diagnostic.CheckDBPerf

	// secondsToRun := 1
	// actual, err := g2diagnostic.CheckDBPerf(ctx, secondsToRun)
	// if err != nil {
	// 	logger.Log(2001, err)
	// }
	// fmt.Println(actual)

	// Work with G2engine.

	g2engine, err := getG2engine(ctx)
	if err != nil {
		logger.Log(2002, err)
	}

	// g2engine.AddRecordWithInfo

	dataSourceCode := "TEST"
	recordID := strconv.Itoa(rand.Intn(1000000000))
	jsonData := fmt.Sprintf(
		"%s%s%s",
		`{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "`,
		recordID,
		`", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`)
	loadID := dataSourceCode
	var flags int64 = 0

	withInfo, err := g2engine.AddRecordWithInfo(ctx, dataSourceCode, recordID, jsonData, loadID, flags)
	if err != nil {
		logger.Log(2003, err)
	}

	logger.Log(2, withInfo)
}
