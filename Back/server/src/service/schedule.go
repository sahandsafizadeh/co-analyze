package service

import (
	"server/server/src/config"
	"server/server/src/util"
	"time"
)

const LogClearPeriod = 24 * time.Hour
const StatisticsUpdatePeriod = 15 * time.Minute

// RunServices Starts an infinite thread for each background service.
// The first handles clearing log files based on {LogClearPeriod} duration.
// The second re-executes all statistic updates based on {StatisticsUpdatePeriod} duration.
// Thread safety is neither managed nor necessary for log files.
func RunServices() {
	go func() {
		tck := time.NewTicker(LogClearPeriod)
		defer tck.Stop()
		for range tck.C {
			util.ClearFileContent(config.ErrorLogPath)
			util.ClearFileContent(config.ServiceLogPath)
		}
	}()
	go func() {
		tck := time.NewTicker(StatisticsUpdatePeriod)
		defer tck.Stop()
		for range tck.C {
			updateStatistics()
		}
	}()
}
