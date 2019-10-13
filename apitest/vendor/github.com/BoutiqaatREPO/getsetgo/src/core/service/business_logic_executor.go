package service

import (
	"fmt"
	"strings"

	"github.com/BoutiqaatREPO/getsetgo/src/core/common/versionmanager"

	"github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"
	workflow "github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/utils/misc"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/utils/orchestratorhelper"
)

type BusinessLogicExecutor struct {
	id string
}

func (n BusinessLogicExecutor) Name() string {
	return "Business Logic Executor"
}

func (n *BusinessLogicExecutor) SetID(id string) {
	n.id = id
}

func (n BusinessLogicExecutor) GetID() (id string, err error) {
	return n.id, nil
}

func (n BusinessLogicExecutor) Execute(data workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	rc, _ := data.ExecContext.Get(constants.RequestContext)
	logger.Info(fmt.Sprintln("entered ", n.Name()), rc)

	resource, version, action, orchBucket, pathParams := getServiceVersion(data)

	data.ExecContext.Set("BASIC-VERSION", versionmanager.BasicVersion{
		Resource: resource,
		Version:  version,
		Action:   action,
		BucketID: orchBucket,
	})

	//Set the resource in exec context. So it can be used when needed.
	data.ExecContext.Set("RESOURCE", resource)

	logger.Debug(fmt.Sprintf("Resource: %s, Version: %s, Action: %s, BucketId: %s, PathParams: %s", resource,
		version, action, orchBucket, pathParams), rc)

	paramsArr := strings.Split(pathParams, ".")
	pathParams = paramsArr[0]
	if len(paramsArr) == 2 {
		data.IOData.Set(constants.ResponseType, n.GetContentType(paramsArr[1]))
	}

	orchestrator, ratelimiter, parameters, oerr := orchestratorhelper.GetOrchestrator(resource, version,
		action, orchBucket, pathParams)
	if oerr != nil {
		data.IOData.Set(constants.APPError, oerr)
		return data, nil
	}

	if ratelimiter != nil {
		if rl := *ratelimiter; rl != nil {
			exceeded, res, err := rl.RateLimit("")
			if err != nil {
				appError := &constants.AppError{
					Code:    constants.RateLimiterInternalError,
					Message: err.Error(),
				}
				data.IOData.Set(constants.APPError, appError)
				return data, nil
			}
			if exceeded {
				appError := &constants.AppError{
					Code:             constants.RateLimitExceeded,
					Message:          fmt.Sprintf("Retry after: %v", res.RetryAfter),
					DeveloperMessage: fmt.Sprintf("Rate limit exceeded"),
				}
				data.IOData.Set(constants.APPError, appError)
				return data, nil
			}
		}
	}

	req, err := misc.GetRequestFromIO(data)
	if err == nil {
		req.PathParameters = parameters
	} else {
		logger.Error("Error in getting request from Workflow IO Data")
	}

	res, err := orchestratorhelper.ExecuteOrchestrator(&data, orchestrator)

	data.IOData.Set(constants.ResponseData, res)
	if err != nil {
		data.IOData.Set(constants.APPError, err)
		return data, nil
	}

	logger.Info(fmt.Sprintln("exiting ", n.Name()), rc)

	return data, nil
}

func (n BusinessLogicExecutor) GetContentType(extension string) string {
	switch extension {
	case "csv":
		return constants.RESPONSE_TYPE_CSV
	case "jpg", "jpeg":
		return constants.RESPONSE_TYPE_IMG_JPG
	case "png":
		return constants.RESPONSE_TYPE_IMG_PNG
	case "gif":
		return constants.RESPONSE_TYPE_IMG_GIF
	default:
		return constants.RESPONSE_TYPE_JSON
	}
}
