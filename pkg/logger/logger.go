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
	infoPath, debugPath, errPath := "info.log", "debug.log", "error.log"

	var debugfile, err1 = os.Create(debugPath)
	var infofile, err2 = os.Create(infoPath)
	var errfile, err3 = os.Create(errPath)
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Loi moi ", err1, err2, err3)
		panic(err1)
	}
	LogInfo = log.New(debugfile, "", log.Ldate+log.Ltime)
	LogInfo.Println("Debug : " + debugPath)

	LogInfo = log.New(infofile, "", log.Ldate+log.Ltime)
	LogInfo.Println("Info : " + debugPath)

	LogInfo = log.New(errfile, "", log.Ldate+log.Ltime)
	LogInfo.Println("Error : " + debugPath)
}
