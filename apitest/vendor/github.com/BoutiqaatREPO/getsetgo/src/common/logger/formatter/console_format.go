package formatter

import (
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/common/logger/message"
)

type consoleFormat struct {
}

//Define color specific constants.
const greenColor = "\x1b[32m"
const redColor = "\x1b[91m"
const yellowColor = "\x1b[33m"
const blueColor = "\x1b[34m"
const pinkColor = "\x1b[91m"
const purpleColor = "\x1b[35m"
const lightBlueColor = "\x1b[36m"
const defaultStyle = "\x1b[0m"
const lightGrayColor = "\x1b[30m"

//GetFormattedLog returns formatted log as a string interface
func (this *consoleFormat) GetFormattedLog(msg *message.LogMsg) interface{} {

	return fmt.Sprintf("%s %s %s %s %s %s %s \n %s",
		this.getTime(msg.TimeStamp),
		this.getLevel(msg.Level),
		this.getURI(msg.URI),
		msg.TransactionID,
		msg.RequestID,
		msg.SessionID,
		msg.UserID,
		msg.Message)

}

func (this *consoleFormat) getTime(time string) string {
	return fmt.Sprintf("%s %s %s", pinkColor, time, defaultStyle)
}

func (this *consoleFormat) getURI(uri string) string {
	return fmt.Sprintf("%s %s %s", lightBlueColor, uri, defaultStyle)
}

func (this *consoleFormat) getLevel(level string) string {

	switch level {
	case "error":
		return fmt.Sprintf("%s ERR %s", "\x1b[41m", defaultStyle)
	case "warning":
		return fmt.Sprintf("%s WRN %s", "\x1b[43m", defaultStyle)
	case "debug":
		return fmt.Sprintf("%s DBG %s", "\x1b[42m", defaultStyle)
	case "trace":
		return fmt.Sprintf("%s TRE %s", "\x1b[7m", defaultStyle)
	case "info":
		return fmt.Sprintf("%s INF %s", "\x1b[7m", defaultStyle)
	default:
		return fmt.Sprintf(level)
	}

}
