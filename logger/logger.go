// code is part of package logger
package logger

//import log and os package
import (
	"log"
	"os"
)

// variable declaration
var (
	flags                      int = log.LstdFlags | log.Lshortfile
	infoLog, warnLog, errorLog *log.Logger
)

// Create a new logger for each type of msg: info, warn, error
func init() {
	infoLog = log.New(os.Stdout, "INFO: ", flags)
	warnLog = log.New(os.Stdout, "WARN: ", flags)
	errorLog = log.New(os.Stdout, "ERROR: ", flags)
}

// function for logging information
func Info(str interface{}) {
	infoLog.Println(str)
}

// function for logging warnings
func Warn(str interface{}) {
	warnLog.Println(str)
}

// function for logging errors
func Error(str interface{}) {
	errorLog.Println(str)
}
