package orchestrator

import (
	utilhttp "github.com/BoutiqaatREPO/getsetgo/src/common/utils/http"
)

/*
Data structure to store the input and output for each workflow execution node
*/
type WorkFlowIOInterface interface {
	Get(key string) (value interface{}, err error)
	Set(key string, value interface{}) (err error)
	Clone() WorkFlowIOInterface

	GetRequest() (*utilhttp.Request, error)
	SetHeader(key, value string) error
}
