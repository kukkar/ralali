package middleware

import (
	"net/http"
	"strings"

	//"github.com/BoutiqaatREPO/getsetgo/src/common"
	"github.com/BoutiqaatREPO/getsetgo/src/common/utils"

	"github.com/BoutiqaatREPO/getsetgo/src/common/config"
	utilhttp "github.com/BoutiqaatREPO/getsetgo/src/common/utils/http"
)

var AllowedHeaders = strings.Join([]string{
	utilhttp.CustomHeaderMap[utilhttp.SessionID],
	utilhttp.CustomHeaderMap[utilhttp.TransactionID],
	utilhttp.CustomHeaderMap[utilhttp.UserID],
	utilhttp.CustomHeaderMap[utilhttp.TokenID],
	utilhttp.CustomHeaderMap[utilhttp.RequestID],
	"Origin",
	"Accept",
	"Content-Type",
	"Content-Length",
	"Accept-Encoding",
}, ",")

//
// Set Predefined Headers before sending response.
//
func ResponseHeaderHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if rw := getCoreWriter(w); rw != nil {
			origin := r.Header.Get("Origin")
			if rw.resp == nil { // Incase of preflight this will be nil.
				rw.resp = &utilhttp.APIResponse{HTTPStatus: 200}
				rw.resp.Headers = make(map[string]string, 0)
			}
			//!Important: Donot set content-type hedaer here, its already being sent from
			// src/core/common/utils/responseheaders/writer.go
			rw.resp.Headers[utilhttp.CustomHeaderMap[utilhttp.RequestID]] = rw.creq.context.RequestID

			//Control for Access control headers.
			if origin != "" {
				//check if the origin is allowed.
				if hosts := config.GetAllowedHosts(); len(hosts) > 0 {
					for _, host := range hosts {
						if host == origin {
							rw.resp.Headers["Access-Control-Allow-Origin"] = origin
							rw.resp.Headers["Access-Control-Allow-Headers"] = AllowedHeaders
							rw.resp.Headers["Access-Control-Allow-Methods"] = "GET, POST, DELETE, PUT, PATCH, OPTIONS"
							rw.resp.Headers["Access-Control-Allow-Credentials"] = "true"
							rw.resp.Headers["Vary"] = "Origin"
							break
						}
					}
				}
			}
		}
		fn(w, r)
	}
}

//
// Set Mandatory Request headers.
//
func RequestHeaderHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if isPreFlightRequest(r) {
			fn(w, r)
			return
		}
		//check if requestid is set, if not define a reqId.
		reqId := r.Header.Get(utilhttp.CustomHeaderMap[utilhttp.RequestID])
		if reqId == "" {
			reqId = utils.GetNewReqId(config.GlobalAppConfig.AppName)
			r.Header.Set(utilhttp.CustomHeaderMap[utilhttp.RequestID], reqId)
		}
		fn(w, r)
	}
}
