package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/BoutiqaatREPO/getsetgo/src/common/config"
	"github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"
	utilhttp "github.com/BoutiqaatREPO/getsetgo/src/common/utils/http"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
)

func getAppRequest(r *http.Request) (*utilhttp.Request, error) {
	req, rerr := utilhttp.GetRequest(r)
	if rerr != nil {
		return nil, rerr
	}

	return &req, nil
}

func getBucketsMap(bucketsList string) map[string]string {
	buckets := strings.Split(bucketsList, constants.FieldSeperator)
	bucketMap := make(map[string]string, len(buckets))
	for _, v := range buckets {
		bKV := strings.Split(v, constants.KeyValueSeperator)
		if len(bKV) < 2 { //invalid bucket
			continue
		}
		bucketMap[bKV[0]] = bKV[1]
	}
	return bucketMap
}

//Get the Service WorkFlow Data
func GetData(r *http.Request) (*orchestrator.WorkFlowData, error) {
	serviceInputOutput := new(orchestrator.WorkFlowIOInMemoryImpl)
	appReq, rerr := getAppRequest(r)
	if rerr != nil {
		return nil, rerr
	}

	serviceInputOutput.Set(constants.URI, appReq.URI)
	serviceInputOutput.Set(constants.HTTPVerb, appReq.HTTPVerb)
	serviceInputOutput.Set(constants.Request, appReq)

	logger.Info(fmt.Sprintf("Service Input Output %v", serviceInputOutput))

	serviceEcContext := new(orchestrator.WorkFlowECInMemoryImpl)
	serviceEcContext.SetUserId(appReq.Headers.UserID)
	serviceEcContext.Set(utilhttp.CustomHeaderMap[utilhttp.SessionID], appReq.Headers.SessionID)

	// Get Token from URL Params
	// if not found get from headers
	// if not found get from cookie
	token := func() string {
		t := appReq.OriginalRequest.URL.Query().Get("token")
		if t != "" {
			return t
		}
		if appReq.Headers.AuthToken != "" {
			return appReq.Headers.AuthToken
		}
		cook, _ := appReq.OriginalRequest.Cookie("token")
		if cook != nil && cook.Value != "" {
			return cook.Value
		}
		return ""
	}()
	serviceEcContext.SetToken(token)

	serviceEcContext.Set(constants.UserAgent, appReq.Headers.UserAgent)
	serviceEcContext.Set(constants.HTTPReferrer, appReq.Headers.Referrer)
	serviceEcContext.Set(utilhttp.CustomHeaderMap[utilhttp.RequestID], appReq.Headers.RequestID)
	serviceEcContext.SetBuckets(getBucketsMap(appReq.Headers.BucketsList))
	serviceEcContext.SetDebugFlag(appReq.Headers.Debug)

	serviceEcContext.SetRequestContext(utilhttp.RequestContext{
		AppName:       config.GlobalAppConfig.AppName,
		UserID:        appReq.Headers.UserID,
		SessionID:     appReq.Headers.SessionID,
		RequestID:     appReq.Headers.RequestID,
		TransactionID: appReq.Headers.TransactionID,
		URI:           appReq.URI,
		Method:        string(appReq.HTTPVerb),
		ClientAppID:   appReq.Headers.ClientAppID,
	})

	logger.Info(fmt.Sprintf("Service Execution Context %v", serviceInputOutput))

	serviceWorkFlowData := new(orchestrator.WorkFlowData)
	serviceWorkFlowData.Create(serviceInputOutput, serviceEcContext)

	return serviceWorkFlowData, nil
}
