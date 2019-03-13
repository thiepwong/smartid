package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	LogInfo  *log.Logger
	LogDebug *log.Logger
	LogErr   *log.Logger
)

func init() {
	// set location of log file
	var (
		infofile, errfile, debugfile *os.File
		err1, err2, err3             error
	)

	infoPath, debugPath, errPath := "logs/info.log", "logs/debug.log", "logs/error.log"

	os.Mkdir("logs", os.ModePerm)
	if _, err := os.Stat(infoPath); os.IsNotExist(err) {
		infofile, err2 = os.Create(infoPath)
		LogInfo = log.New(infofile, "", log.Ldate+log.Ltime)
		LogInfo.Println("Info : " + debugPath)

	} else {
		infofile, err2 = os.OpenFile(infoPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		LogInfo = log.New(infofile, "", log.Ldate+log.Ltime)
	}

	if _, err := os.Stat(errPath); os.IsNotExist(err) {
		errfile, err3 = os.Create(errPath)

		LogErr = log.New(errfile, "", log.Ldate+log.Ltime)
		LogErr.Println("Error : " + errPath)
	} else {
		errfile, err2 = os.OpenFile(errPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		LogErr = log.New(errfile, "", log.Ldate+log.Ltime)
	}

	if _, err := os.Stat(debugPath); os.IsNotExist(err) {
		debugfile, err1 = os.Create(debugPath)
		LogDebug = log.New(debugfile, "", log.Ldate+log.Ltime)
		LogDebug.Println("Debug : " + debugPath)
	} else {
		debugfile, err2 = os.OpenFile(debugPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		LogDebug = log.New(debugfile, "", log.Ldate+log.Ltime)
	}

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Loi moi ", err1, err2, err3)
		panic(err1)
	}

}
