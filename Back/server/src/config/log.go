package config

import (
	"github.com/labstack/echo"
	"log"

	"io"
	"os"
)

const (
	filePermissions   = 0200
	serviceLogFlags   = log.LstdFlags | log.Lshortfile
	fileOperationMode = os.O_RDWR | os.O_CREATE | os.O_APPEND
	errorLogPath      = "server/err_log.txt"
	serviceLogPath    = "server/srv_log.txt"
	//errorLogPath      = "/var/log/co-analyze/err_log.txt"
	//requestLogPath    = "/var/log/co-analyze/srv_log.txt"
)

var ServiceLogger *log.Logger

// InitLogger Builds and configures application's service and error logger.
// Causes panic if any of the two log files is not successfully opened.
func InitLogger(e *echo.Echo) {
	serviceLogFile, err := openLogFile(serviceLogPath)
	if err != nil {
		panic("Unable to open service log file")
	}
	errorLogFile, err := openLogFile(errorLogPath)
	if err != nil {
		panic("Unable to open error log file")
	}
	e.Logger.SetOutput(errorLogFile)
	ServiceLogger = log.New(serviceLogFile, log.Prefix(), serviceLogFlags)
}

// ----- helpers

func openLogFile(path string) (io.Writer, error) {
	fileWriter, err := os.OpenFile(path, fileOperationMode, filePermissions)
	if err != nil {
		return fileWriter, err
	}
	return fileWriter, nil
}
