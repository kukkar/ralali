package hello

import (
	reflorest_constants "github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	workflow "github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
)

type HelloWorld struct {
	id string
}

func (n *HelloWorld) SetID(id string) {
	n.id = id
}

func (n HelloWorld) GetID() (id string, err error) {
	return n.id, nil
}

func (a HelloWorld) Name() string {
	return "HelloWorld"
}

func (a HelloWorld) getData() string {
	return "Success Response"
}

func (a HelloWorld) Execute(io workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	//Business Logic
	io.IOData.Set(reflorest_constants.Result, a.getData())
	return io, nil
	// return io, &reflorest_constants.AppError{
	// 	Code:    errors.FunctionalityNotImplementedErrorCode,
	// 	Message: "Sample Error",
	// }
}
