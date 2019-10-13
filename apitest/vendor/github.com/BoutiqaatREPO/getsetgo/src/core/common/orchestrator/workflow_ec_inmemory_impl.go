package orchestrator

import (
	"errors"
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/common/utils/http"
)

const (
	buckets        string = "BUCKETS"
	threadID       string = "THREAD_ID"
	debugFlag      string = "IS_DEBUG"
	debugMsg       string = "DEBUG_MSG"
	requestContext string = "REQUEST_CONTEXT"
)

type WorkFlowECInMemoryImpl struct {
	store map[string]interface{}
}

type WorkflowDebugDataInMemory struct {
	Key   string
	Value string
}

func (ec *WorkFlowECInMemoryImpl) Get(key string) (value interface{}, err error) {
	//Check if the key is already present
	res, found := ec.store[key]
	if !found {
		errString := fmt.Sprintln("In memory ExecutionContext store does not contain key: ", key)
		return nil, errors.New(errString)
	}
	return res, nil
}

func (ec *WorkFlowECInMemoryImpl) Set(key string, value interface{}) (err error) {
	if ec.store == nil {
		ec.store = make(map[string]interface{})
	}
	ec.store[key] = value
	return nil
}

func (ec *WorkFlowECInMemoryImpl) SetBuckets(bucketIDMap map[string]string) (err error) {
	return ec.Set(buckets, bucketIDMap)
}

func (ec *WorkFlowECInMemoryImpl) GetBuckets() (bucketIDMap map[string]string, err error) {
	res, err := ec.Get(buckets)
	if v, ok := res.(map[string]string); ok {
		bucketIDMap = v
	}
	return bucketIDMap, err
}

func (ec *WorkFlowECInMemoryImpl) GetExecuteThreadID() (executeThreadID string, err error) {
	res, err := ec.Get(threadID)
	if v, ok := res.(string); ok {
		executeThreadID = v
	}
	return executeThreadID, err
}

func (ec *WorkFlowECInMemoryImpl) SetDebugFlag(flag bool) (err error) {
	return ec.Set(debugFlag, flag)
}

func (ec *WorkFlowECInMemoryImpl) SetDebugMsg(msgkey string, msgData string) (err error) {
	if ec.store == nil {
		ec.store = make(map[string]interface{})
	}

	isDebugVal, isDebugSet := ec.store[debugFlag]
	if !isDebugSet {
		return nil
	}
	isDebug, ok := isDebugVal.(bool)
	if !ok {
		return errors.New("Incorrect data stored in debug flag")
	}
	if !isDebug {
		return nil
	}

	dMsg, found := ec.store[debugMsg]

	//This is the first debug msg
	if !found {
		newDebugMsg := WorkflowDebugDataInMemory{Key: msgkey, Value: msgData}
		ec.store[debugMsg] = []WorkflowDebugDataInMemory{newDebugMsg}
		return nil
	}

	//Debug Messages are already present
	v, ok := dMsg.([]WorkflowDebugDataInMemory)
	if ok {
		v = append(v, WorkflowDebugDataInMemory{Key: msgkey, Value: msgData})
	}
	ec.store[debugMsg] = v
	return nil
}

func (ec *WorkFlowECInMemoryImpl) GetDebugMsg() (msg []interface{}, err error) {
	msgData, err := ec.Get(debugMsg)
	if v, ok := msgData.([]WorkflowDebugDataInMemory); ok {
		msg = make([]interface{}, len(v))
		for i, val := range v {
			msg[i] = val
		}
	}
	return msg, err
}

func (ec *WorkFlowECInMemoryImpl) GetRequestContext() (http.RequestContext, error) {

	rcI, err := ec.Get(requestContext)
	if err != nil {
		return http.RequestContext{}, err
	}
	rc, ok := rcI.(http.RequestContext)
	if !ok {
		return http.RequestContext{}, fmt.Errorf("WorkFlowECInMemoryImpl#GetRequestContext, Expected http.RequestContext, Got: %T", rcI)
	}
	return rc, nil
}

func (ec *WorkFlowECInMemoryImpl) SetRequestContext(rc http.RequestContext) error {
	return ec.Set(requestContext, rc)
}

func (ec *WorkFlowECInMemoryImpl) GetToken() string {

	tokenI, _ := ec.Get(http.CustomHeaderMap[http.TokenID])
	if v, ok := tokenI.(string); ok {
		return v
	}
	return ""
}

func (ec *WorkFlowECInMemoryImpl) SetToken(token string) error {
	return ec.Set(http.CustomHeaderMap[http.TokenID], token)
}

func (ec *WorkFlowECInMemoryImpl) GetUserId() string {
	userI, _ := ec.Get(http.CustomHeaderMap[http.UserID])
	if v, ok := userI.(string); ok {
		return v
	}
	return ""
}

func (ec *WorkFlowECInMemoryImpl) SetUserId(user string) error {
	return ec.Set(http.CustomHeaderMap[http.UserID], user)
}
