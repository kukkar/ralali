package formatter

import (
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/common/logger/message"
)

type stringFormat struct {
}

//formatString is format of the log in string formattype configuration
var formatString = "[level : %s, message : %s, tId : %s, reqId : %s, appId : %s, sessionId : %s, userId : %s, stackTraces : %s, timestamp : %s, uri : %s]"

//GetFormattedLog returns formatted log as a string interface
func (sf *stringFormat) GetFormattedLog(msg *message.LogMsg) interface{} {
	var stack []string
	if msg.Level == "error" {
		stack = msg.StackTraces
	}

	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",

		msg.TimeStamp,
		msg.Level,
		msg.Message,
		msg.TransactionID,
		msg.RequestID,
		msg.AppID,
		msg.SessionID,
		msg.UserID,
		msg.URI,
		stack,
	)
}
