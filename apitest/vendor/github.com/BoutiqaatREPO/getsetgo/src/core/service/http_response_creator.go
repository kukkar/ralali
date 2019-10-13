package service

import (
	"bytes"
	"fmt"
	"log"

	"encoding/csv"
	"encoding/json"

	"github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"
	utilhttp "github.com/BoutiqaatREPO/getsetgo/src/common/utils/http"
	workflow "github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
)

type HTTPResponseCreator struct {
	id string
}

func (n HTTPResponseCreator) Name() string {
	return "Http Response Creator"
}

func (n *HTTPResponseCreator) SetID(id string) {
	n.id = id
}

func (n HTTPResponseCreator) GetID() (id string, err error) {
	return n.id, nil
}

func (n *HTTPResponseCreator) Execute(data workflow.WorkFlowData) (workflow.WorkFlowData, error) {

	responseTypeI, _ := data.IOData.Get(constants.ResponseType)
	responseType, ok := responseTypeI.(string)
	if !ok {
		responseType = constants.RESPONSE_TYPE_JSON

	}
	switch responseType {
	case constants.RESPONSE_TYPE_IMG_PNG,
		constants.RESPONSE_TYPE_IMG_JPG,
		constants.RESPONSE_TYPE_IMG_GIF:
		return n.PrepareResponse(data)
	case constants.RESPONSE_TYPE_CSV:
		return n.PrepareCSVResponse(data)
	default:
		return n.PrepareJsonResponse(data)
	}
}

func (n *HTTPResponseCreator) PrepareCSVResponse(data workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	//get error
	resStatus, _ := data.IOData.Get(constants.APPError)
	//get response
	resData, _ := data.IOData.Get(constants.ResponseData)

	appError := new(constants.AppErrors)
	if resStatus != nil {
		if v, ok := resStatus.(*constants.AppError); ok {
			if v != nil { //if v is of type *AppError and is not nil
				appError.Errors = []constants.AppError{*v}
			}
		} else if v, ok := resStatus.(*constants.AppErrors); ok {
			if v != nil { //v is of type *AppErrors and is not nil
				appError = v
			}
		} else {
			appError.Errors = []constants.AppError{{Code: constants.InvalidErrorCode,
				Message: "Invalid App error"}}
		}
	}

	var responseStatusCode constants.HTTPCode
	var responseData []byte

	if appError.Errors == nil {
		switch v := resData.(type) {
		case [][]string:
			bufferedWriter := bytes.NewBuffer(nil)
			csvwriter := csv.NewWriter(bufferedWriter)
			for _, record := range v {
				if err := csvwriter.Write(record); err != nil {
					log.Printf("error writing record to csv: %s", err.Error())
					continue
				}
			}
			csvwriter.Flush()
			if err := csvwriter.Error(); err != nil {
				responseData = []byte(fmt.Sprintf("Could not write csv: %s", err.Error()))
				break
			}
			responseData = bufferedWriter.Bytes()
		case []byte:
			responseData = v
		case string:
			responseData = []byte(v)
		default:
			responseData, _ = json.Marshal(resData)
		}
		responseStatusCode = constants.HTTPStatusSuccessCode
	} else {
		status := constants.GetAppHTTPError(*appError)
		responseStatusCode = status.HTTPStatusCode
		responseData = []byte(appError.Error())
	}

	r, _ := data.IOData.Get(constants.APIResponse)
	apiResponse, _ := r.(utilhttp.APIResponse)
	apiResponse.Body = responseData
	apiResponse.HTTPStatus = responseStatusCode
	data.IOData.Set(constants.APIResponse, apiResponse)

	return data, nil
}

func (n *HTTPResponseCreator) PrepareResponse(data workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	//get error
	resStatus, _ := data.IOData.Get(constants.APPError)
	//get response
	resData, _ := data.IOData.Get(constants.ResponseData)

	appError := new(constants.AppErrors)
	if resStatus != nil {
		if v, ok := resStatus.(*constants.AppError); ok {
			if v != nil { //if v is of type *AppError and is not nil
				appError.Errors = []constants.AppError{*v}
			}
		} else if v, ok := resStatus.(*constants.AppErrors); ok {
			if v != nil { //v is of type *AppErrors and is not nil
				appError = v
			}
		} else {
			appError.Errors = []constants.AppError{{Code: constants.InvalidErrorCode,
				Message: "Invalid App error"}}
		}
	}

	var responseStatusCode constants.HTTPCode
	var responseData []byte

	if appError.Errors == nil {
		switch v := resData.(type) {
		case []byte:
			responseData = v
		case string:
			responseData = []byte(v)
		default:
			responseData, _ = json.Marshal(resData)
		}
		responseStatusCode = constants.HTTPStatusSuccessCode
	} else {
		status := constants.GetAppHTTPError(*appError)
		responseStatusCode = status.HTTPStatusCode
		responseData = []byte(appError.Error())
	}

	r, _ := data.IOData.Get(constants.APIResponse)
	apiResponse, _ := r.(utilhttp.APIResponse)
	apiResponse.Body = responseData
	apiResponse.HTTPStatus = responseStatusCode
	data.IOData.Set(constants.APIResponse, apiResponse)

	return data, nil
}

func (n *HTTPResponseCreator) PrepareJsonResponse(data workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	rc, _ := data.ExecContext.Get(constants.RequestContext)
	logger.Info(fmt.Sprintln("entered ", n.Name()), rc)

	resStatus, _ := data.IOData.Get(constants.APPError)
	resData, _ := data.IOData.Get(constants.ResponseData)

	appError := new(constants.AppErrors)

	if resStatus != nil {
		if v, ok := resStatus.(*constants.AppError); ok {
			if v != nil { //if v is of type *AppError and is not nil
				appError.Errors = []constants.AppError{*v}
			}
		} else if v, ok := resStatus.(*constants.AppErrors); ok {
			if v != nil { //v is of type *AppErrors and is not nil
				appError = v
			}
		} else {
			appError.Errors = []constants.AppError{{Code: constants.InvalidErrorCode,
				Message: "Invalid App error"}}
		}
	}
	status := constants.GetAppHTTPError(*appError)
	debugData, _ := data.ExecContext.GetDebugMsg()

	resource, _, _, _, _ := getServiceVersion(data)

	if status.HTTPStatusCode != constants.HTTPStatusSuccessCode {
		logger.Error(fmt.Sprintf("%s_%v Application Errors : %v", resource, status.HTTPStatusCode, appError), rc)
	}

	var appDebugData []utilhttp.Debug
	for _, d := range debugData {
		if v, ok := d.(workflow.WorkflowDebugDataInMemory); ok {
			appDebugData = append(appDebugData, utilhttp.Debug{Key: v.Key, Value: v.Value})
		}
	}

	m, _ := data.IOData.Get(constants.ResponseMetaData)
	md, _ := m.(*utilhttp.ResponseMetaData)
	appResponse := utilhttp.Response{Status: *status, Data: resData, DebugData: appDebugData, MetaData: md}
	data.IOData.Set(constants.Response, appResponse)
	jsonBody, err := json.Marshal(appResponse)
	if err != nil {
		return data, err
	}
	r, _ := data.IOData.Get(constants.APIResponse)
	apiResponse, _ := r.(utilhttp.APIResponse)
	apiResponse.HTTPStatus = appResponse.Status.HTTPStatusCode
	apiResponse.Body = jsonBody
	data.IOData.Set(constants.APIResponse, apiResponse)

	logger.Info(fmt.Sprintln("exiting ", n.Name()), rc)

	return data, nil
}
