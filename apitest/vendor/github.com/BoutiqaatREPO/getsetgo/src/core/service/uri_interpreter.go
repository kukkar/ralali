package service

import (
	"fmt"
	"strings"

	"github.com/BoutiqaatREPO/getsetgo/src/common/config"
	"github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"
	utilhttp "github.com/BoutiqaatREPO/getsetgo/src/common/utils/http"
	workflow "github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
)

type URIInterpreter struct {
	id string
}

func (u URIInterpreter) Name() string {
	return "URL Interpreter"
}

func (u *URIInterpreter) SetID(id string) {
	u.id = id
}

func (u URIInterpreter) GetID() (id string, err error) {
	return u.id, nil
}

func (u URIInterpreter) Execute(data workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	rc, _ := data.ExecContext.Get(constants.RequestContext)

	logger.Info(fmt.Sprintln("Entered ", u.Name()), rc)

	resource, version, action, pathParams, queryString := u.getResource(data)
	data.IOData.Set(constants.Resource, resource)
	data.IOData.Set(constants.Version, version)
	data.IOData.Set(constants.Action, action)
	data.IOData.Set(constants.PathParams, pathParams)
	data.IOData.Set(constants.QueryString, queryString)
	data.IOData.Set(constants.ResponseMetaData, utilhttp.NewResponseMetaData())

	logger.Info(fmt.Sprintln("Exiting ", u.Name()), rc)
	return data, nil
}

func (u URIInterpreter) getResource(data workflow.WorkFlowData) (resource string,
	version string,
	action string, pathParams string, queryString map[string]string) {

	rc, _ := data.ExecContext.Get(constants.RequestContext)
	uridata, _ := data.IOData.Get(constants.URI)
	actiondata, _ := data.IOData.Get(constants.HTTPVerb)

	var uri string
	if v, ok := uridata.(string); ok {
		uri = v
	}
	logger.Info(fmt.Sprintln("uri is ", uri), rc)

	//prepare query string data
	queryString = u.getQueryParams(uri)

	// remove query parameter if any
	if index := strings.Index(uri, "?"); index != -1 {
		uri = uri[:index]
	}
	// split on basis of separator
	uriArr := strings.Split(uri[1:], "/")

	//its an health check api if url: appname/healthckeck
	if len(uriArr) >= 2 &&
		uriArr[0] == config.GlobalAppConfig.AppName &&
		strings.ToUpper(uriArr[1]) == constants.HealthCheckAPI {
		resource = constants.HealthCheckAPI
		version = ""
		pathParams = ""
		//else if url of format appname/version/resource its a normal url.
	} else if len(uriArr) >= 3 && uriArr[0] == config.GlobalAppConfig.AppName {
		resource = strings.ToUpper(uriArr[2])
		version = strings.ToUpper(uriArr[1])
		if len(uriArr) > 3 {
			pathParams = strings.Join(uriArr[3:], "/")
		}
	} else {
		//Badly formed URI
		resource = ""
		version = ""
		pathParams = ""
	}

	if v, ok := actiondata.(utilhttp.Method); ok {
		action = string(v)
	}

	return resource, version, action, pathParams, queryString
}

func (u URIInterpreter) getQueryParams(uri string) map[string]string {
	m := make(map[string]string)
	if index := strings.Index(uri, "?"); index != -1 {
		uri = uri[index+1:]
		//split on &
		uriArr := strings.Split(uri, "&")
		if len(uriArr) > 0 {
			for _, data := range uriArr {
				dataArr := strings.Split(data, "=")
				if len(dataArr) == 2 { //check if we have exactly two elements var=data
					m[dataArr[0]] = dataArr[1]
				}
			}
		}
	}
	return m
}
