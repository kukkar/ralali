package formatter

import (
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger/message"
)

// FormatInterface interface methods for formatterss
type FormatInterface interface {
	GetFormattedLog(msg *message.LogMsg) interface{}
}
