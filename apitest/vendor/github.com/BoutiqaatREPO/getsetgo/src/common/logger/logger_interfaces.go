package logger

import (
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger/message"
)

type LogInterface interface {
	Trace(msg message.LogMsg)
	Warning(msg message.LogMsg)
	Info(msg message.LogMsg)
	Error(msg message.LogMsg)
	Debug(msg message.LogMsg)
	Profile(msg message.LogMsg)
}
