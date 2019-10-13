package middleware

import (
	"fmt"
	"net/http"

	"github.com/sanksons/gowraps/util"

	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"

	"github.com/BoutiqaatREPO/getsetgo/src/common/logger/blackbox"
)

//
// Log Request Response incase blackbox path is supplied.
//
func LoggerHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var requestData, responseData string

		fn(w, r)

		if isPreFlightRequest(r) || !blackbox.IsBlackBoxEnabled() {
			return
		}
		if rw := getCoreWriter(w); rw != nil {

			//Check if blackbox logs are disabled for specific api.
			if blackbox.IsDisabledForAPI(rw.version) {
				return
			}

			reqId := rw.creq.context.RequestID

			//set response data.
			responseData = blackbox.GetFormattedResponse(rw.resp)
			requestData = blackbox.GetFormattedRequest(r)
			bbox := blackbox.GetNewBlackBox(blackbox.GetBlackBoxPath(), reqId, requestData, responseData)
			go func(b *blackbox.BlackBoxLog) {
				defer util.PanicHandler("Got Panic in BlackBox goroutine")
				err := b.LogIt()
				if err != nil {
					logger.Error(fmt.Sprintf("Could Not write BlackBox Logs, Error: %s", err.Error()))
				}
			}(bbox)
		}

	}
}
