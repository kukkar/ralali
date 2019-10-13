package stdoutwriter

import (
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/common/logger/formatter"
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger/message"
)

//FileWriter is a file logger structure
type StdoutWriter struct {
	// formatter
	myFormat formatter.FormatInterface
}

//GetNewObj returns a file logger with log file name fname, having configuration
//specified in conf and allowedLogLevel specifies the log level that are actually to
//be logged
func GetNewObj() *StdoutWriter {
	obj := new(StdoutWriter)

	return obj
}

// Write write message to file
func (fw *StdoutWriter) Write(msg *message.LogMsg) {
	str, _ := fw.myFormat.GetFormattedLog(msg).(string)
	fmt.Printf("%s\n", str)
}

// SetFormatter get formatted object
func (fw *StdoutWriter) SetFormatter(format formatter.FormatInterface) {
	fw.myFormat = format
}
