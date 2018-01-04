package logutil

import (
	"github.com/dyweb/gommon/log"
)

//var Logger = log.NewPackageLogger()

// TODO: allow forge package name?
// TODo: have a logger that is parent of all the loggers in this project
func NewLogger() *log.Logger {
	return log.NewPackageLoggerWithSkip(1)
}

func init() {

}
