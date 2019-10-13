package orchestrator

import (
	"errors"
	"fmt"

	utilhttp "github.com/BoutiqaatREPO/getsetgo/src/common/utils/http"
)

type WorkFlowIOInMemoryImpl struct {
	store map[string]interface{}
}

func (io *WorkFlowIOInMemoryImpl) Get(key string) (value interface{}, err error) {
	//Check if the key is already present
	res, found := io.store[key]
	if !found {
		errString := fmt.Sprintf("In memory input output store does not contain key %v", key)
		return nil, errors.New(errString)
	}
	return res, nil
}

func (io *WorkFlowIOInMemoryImpl) Set(key string, value interface{}) (err error) {

	if io.store == nil {
		io.store = make(map[string]interface{})
	}

	io.store[key] = value
	return nil
}

func (io *WorkFlowIOInMemoryImpl) Clone() WorkFlowIOInterface {
	//you cannot generally call methods on pointers directly on values so a pointer to the interface is created.
	ioClone := new(WorkFlowIOInMemoryImpl)
	if io.store == nil {
		return ioClone
	}

	ioClone.store = make(map[string]interface{})
	for k, v := range io.store {
		ioClone.store[k] = v
	}
	return ioClone

}

func (io *WorkFlowIOInMemoryImpl) GetRequest() (*utilhttp.Request, error) {
	httpReq, err := io.Get("REQUEST")
	if err != nil {
		return nil, err
	}
	appHTTPReq, ok := httpReq.(*utilhttp.Request)
	if !ok {
		return nil, fmt.Errorf("WorkFlowIOInMemoryImpl#GetRequest() -> Assertion failed, Expected: httpReq.(*utilhttp.Request), Got: %T", httpReq)
	}
	if appHTTPReq == nil {
		return nil, fmt.Errorf("nil")
	}
	return appHTTPReq, nil
}

func (io *WorkFlowIOInMemoryImpl) SetHeader(key, value string) error {
	var apiresp utilhttp.APIResponse
	httpResp, err := io.Get("API_RESPONSE")
	if err != nil { //we dont have a response defined yet. define it!!
		apiresp = utilhttp.NewAPIResponse()
	} else {
		apiresp, _ = httpResp.(utilhttp.APIResponse)
	}
	apiresp.Headers[key] = value
	io.Set("API_RESPONSE", apiresp)
	return nil
}
