package shortenpost

import (
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	florest_constants "github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	reflorest_constants "github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	workflow "github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
	"github.com/ralali/apitest/src/shortenurl"
)

type ShotenPost struct {
	id string
}

func (n *ShotenPost) SetID(id string) {
	n.id = id
}

func (n ShotenPost) GetID() (id string, err error) {
	return n.id, nil
}

func (a ShotenPost) Name() string {
	return "ShotenPostAPI"
}

func (a ShotenPost) getData() string {
	return "Success Response"
}

func (a ShotenPost) Execute(io workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	//Business Logic

	//Business Logic
	appHttpReq, err := io.IOData.GetRequest()
	if err != nil {
		return io, &constants.AppError{
			Code:             constants.ParamsInValidErrorCode,
			Message:          "Could not read request body",
			DeveloperMessage: fmt.Sprintf("Error: %s", err.Error()),
		}
	}

	useRequest := Request{}
	err = appHttpReq.LoadBody(&useRequest)
	if err != nil {
		return io, &constants.AppError{
			Code:             constants.ParamsInValidErrorCode,
			Message:          "Request Unmarshalling failed",
			DeveloperMessage: fmt.Sprintf("Error: %s", err.Error()),
		}
	}

	sc, svError := shortenurl.SaveUrl(useRequest.Url, useRequest.ShortCode)

	if svError != nil {

		if svError.Code == 1505 {
			return io, &florest_constants.AppError{
				Code:             reflorest_constants.FunctionalityNotImplementedErrorCode,
				Message:          "Unable to read data from IOData",
				DeveloperMessage: "",
			}
		} else if svError.Code == 1501 {
			return io, &florest_constants.AppError{
				Code:             reflorest_constants.ResourceErrorCode,
				Message:          "Unable to read data from IOData",
				DeveloperMessage: "",
			}
		}
	}

	io.IOData.Set(constants.Result, Response{
		ShortCode: sc,
	})
	return io, nil

}
