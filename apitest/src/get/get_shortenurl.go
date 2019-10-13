package get

import (
	"github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	florest_constants "github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	reflorest_constants "github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	workflow "github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/utils/misc"
	"github.com/ralali/apitest/src/shortenurl"
)

type ShortenUrl struct {
	id string
}

func (n *ShortenUrl) SetID(id string) {
	n.id = id
}

func (n ShortenUrl) GetID() (id string, err error) {
	return n.id, nil
}

func (a ShortenUrl) Name() string {
	return "ShortenUrl"
}

func (a ShortenUrl) getData() string {
	return "Success Response"
}

func (a ShortenUrl) Execute(io workflow.WorkFlowData) (workflow.WorkFlowData, error) {

	appHTTPReq, err := misc.GetRequestFromIO(io)

	if err != nil {
		appError := constants.AppError{
			Code:    constants.ParamsInValidErrorCode,
			Message: "Invalid request",
		}

		return io, &appError
	}

	code := appHTTPReq.GetPathParameter("shortcode")

	url, svError := shortenurl.GetUrl(code)

	if svError != nil {

		if svError.Code == 1601 {
			return io, &florest_constants.AppError{
				Code:             reflorest_constants.InvalidRequestURI,
				Message:          "Unable to read data from IOData",
				DeveloperMessage: "",
			}
		}
	}
	io.IOData.Set(constants.Result, Response{
		Code: 302,
		Url:  url,
	})
	return io, nil
}
