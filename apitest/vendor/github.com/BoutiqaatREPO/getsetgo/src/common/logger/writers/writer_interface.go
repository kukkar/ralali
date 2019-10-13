package writers

import (
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger/formatter"
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger/message"
)

type LogWriter interface {
	Write(msg *message.LogMsg)
	SetFormatter(formatter.FormatInterface)
}
