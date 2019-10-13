package misc

import (
	"errors"
	"github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"
	utilhttp "github.com/BoutiqaatREPO/getsetgo/src/common/utils/http"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
)

// GetRequestFromIO returns httpRequest from IO object
func GetRequestFromIO(io orchestrator.WorkFlowData) (*utilhttp.Request, error) {
	httpReq, _ := io.IOData.Get(constants.Request)
	appHTTPReq, ok := httpReq.(*utilhttp.Request)
	if !ok || appHTTPReq == nil {
		logger.Error("GetRequestFromIO() : Bad request.")
		return nil, errors.New("Bad Request")
	}
	return appHTTPReq, nil
}
