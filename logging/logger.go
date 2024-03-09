package logging

import (
	"log"
	"os"
)

var (
	flags                      int = log.LstdFlags | log.Lshortfile
	infoLog, warnLog, errorLog *log.Logger
)

func init() {
	infoLog = log.New(os.Stdout, "INFO: ", flags)
	warnLog = log.New(os.Stdout, "WARN: ", flags)
	errorLog = log.New(os.Stdout, "ERROR: ", flags)
}
func Info(str string) {
	infoLog.Println(str)
}
func Warn(str string) {
	warnLog.Println(str)
}
func Error(str string) {
	errorLog.Println(str)
}
