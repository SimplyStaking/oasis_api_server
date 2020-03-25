package logger

import (
	"io"
	"log"
)

// 3 Types of loggers to be used
var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

// SetLogger creates loggers that will be used through out API
func SetLogger(
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
