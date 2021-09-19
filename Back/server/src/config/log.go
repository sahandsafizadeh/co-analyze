package config

import (
	"github.com/labstack/echo"
	"log"

	"server/server/src/util"
)

const (
	serviceLogFlags = log.LstdFlags | log.Lshortfile
	ErrorLogPath    = "server/err_log.txt"
	ServiceLogPath  = "server/srv_log.txt"
)

var ServiceLogger *log.Logger

// InitLogger Builds and configures application's service and error logger.
// Causes panic if any of the two log files is not successfully opened.
func InitLogger(e *echo.Echo) {
	serviceLogFile, err := util.OpenFile(ServiceLogPath)
	if err != nil {
		panic("Unable to open service log file")
	}
	errorLogFile, err := util.OpenFile(ErrorLogPath)
	if err != nil {
		panic("Unable to open error log file")
	}
	e.Logger.SetOutput(errorLogFile)
	ServiceLogger = log.New(serviceLogFile, log.Prefix(), serviceLogFlags)
}
