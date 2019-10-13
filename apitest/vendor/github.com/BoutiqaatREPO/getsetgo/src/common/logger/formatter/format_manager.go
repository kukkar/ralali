package formatter

import (
	"errors"
)

const (
	STRING  = "string"
	JSON    = "json"
	CONSOLE = "console"
)

// GetFormatter returns required formatter based on string argument
func GetFormatter(ftype string) (ret FormatInterface, err error) {
	switch ftype {
	case STRING:
		ret = new(stringFormat)
	case JSON:
		ret = new(jsonFormat)
	case CONSOLE:
		ret = new(consoleFormat)
	default:
		err = errors.New("unsupported format type:" + ftype)
	}
	return ret, nil
}
