package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/BoutiqaatREPO/getsetgo/src/common/logger/message"
	utilHttp "github.com/BoutiqaatREPO/getsetgo/src/common/utils/http"
)

//
// A logger that is compatible with 3rd party application loggers.
//
type GlobalLogInterface interface {
	Debug(a ...interface{})
	Info(a ...interface{})
	Warning(a ...interface{})
	Error(a ...interface{})
	Fatal(a ...interface{})
}

func GetGlobalLogger(ltype string, level int) (*GlobalLogger, error) {

	l, err := GetLoggerHandle(ltype)
	if err != nil {
		return nil, err
	}
	return &GlobalLogger{
		logger: l,
		level:  level,
	}, nil
}

type GlobalLogger struct {
	level  int
	logger LogInterface
}

func (this *GlobalLogger) Debug(a ...interface{}) {
	if !(this.level >= InfoLevel) {
		return
	}
	this.logger.Info(prepareMessage("info", a))
}

func (this *GlobalLogger) Info(a ...interface{}) {
	if !(this.level >= TraceLevel) {
		return
	}
	this.logger.Trace(prepareMessage("trace", a))
}
func (this *GlobalLogger) Warning(a ...interface{}) {
	if !(this.level >= WarningLevel) {
		return
	}
	this.logger.Warning(prepareMessage("warning", a))
}
func (this *GlobalLogger) Error(a ...interface{}) {
	if !(this.level >= ErrLevel) {
		return
	}
	this.logger.Error(prepareMessage("error", a))
}

func (this *GlobalLogger) Fatal(a ...interface{}) {
	if !(this.level >= ErrLevel) {
		return
	}
	this.logger.Error(prepareMessage("fatal", a))
}

func prepareMessage(mtype string, a ...interface{}) message.LogMsg {
	var msg message.LogMsg
	var isRc bool

	l := len(a)

	if l >= 2 {
		_, isRc = a[1].(utilHttp.RequestContext)
	}

	if l >= 2 && isRc {
		msg = Convert(a...)
	} else {
		msg = Convert(fmt.Sprintf(strings.Repeat("%v", l), a...))
	}
	msg.Level = mtype
	msg.StackTraces = getStackTrace()
	msg.TimeStamp = time.Now().Local().Format(time.RFC3339)

	return msg
}
