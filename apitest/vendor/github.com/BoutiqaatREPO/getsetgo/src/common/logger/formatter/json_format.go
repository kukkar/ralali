package formatter

import (
	"encoding/json"
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/common/logger/message"
)

type jsonFormat struct {
}

//GetFormattedLog returns formatted log
func (jf *jsonFormat) GetFormattedLog(msg *message.LogMsg) interface{} {

	if msg.Level != "error" {
		msg.StackTraces = []string{}
	}

	jMsg, err := json.Marshal(msg)
	if err != nil {
		return fmt.Sprintf("\nError In converting to json %+v\n", msg)
	}
	return string(jMsg)
}
