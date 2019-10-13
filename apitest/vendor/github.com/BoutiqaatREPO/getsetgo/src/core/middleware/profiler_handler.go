package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sanksons/gowraps/util"

	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"

	"github.com/newrelic/go-agent"

	audit "github.com/BoutiqaatREPO/getsetgo/src/common/logger/auditlog"
	"github.com/BoutiqaatREPO/getsetgo/src/common/monitor"

	"github.com/BoutiqaatREPO/getsetgo/src/common/config"
)

func ProfilerHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if isPreFlightRequest(r) {
			fn(w, r)
			return
		}

		profileSTime := time.Now().UnixNano()

		//Check if monitoring is enabled.
		if monitor.IsDefaultEnabled() {
			newRelic, err := monitor.GetNewRelicAgent()
			if rwtmp := getCoreWriter(w); err == nil && rwtmp != nil {
				rwtmp.ResponseWriter = newRelic.GetApp().StartTransaction("", rwtmp.ResponseWriter, r)
				w = rwtmp //Important!!! Overwrite response writer.
			}
		}

		//let the flow pass on.
		fn(w, r)
		// Prepare metrics to push.
		if rw := getCoreWriter(w); rw != nil {
			name := fmt.Sprintf("%s-%s-%s", rw.creq.context.AppName, rw.creq.resource, rw.creq.context.Method)
			if newtx, ok := rw.ResponseWriter.(newrelic.Transaction); ok {
				newtx.SetName(name)
				//@todo: Add developr error message in case of error.
				//@todo: Decide attributes to transactions
				newtx.AddAttribute("Method", rw.creq.context.Method)
				newtx.AddAttribute("StatusCode", rw.resp.HTTPStatus)
				newtx.AddAttribute("Resource", rw.creq.resource)
				newtx.AddAttribute("Service", config.GlobalAppConfig.AppName)
				newtx.AddAttribute("RequestId", rw.creq.context.RequestID)
				if rw.creq.context.TransactionID != "" {
					newtx.AddAttribute("TransactionId", rw.creq.context.TransactionID)
				}
				newtx.AddAttribute("GZipped", rw.doGzip)
				if rw.resp.HTTPStatus != 200 {
					newtx.NoticeError(fmt.Errorf("%+v", string(rw.resp.Body)))
				}
				newtx.End()
			}

			if !audit.IsAuditEnabled() && audit.IsDisabledForAPI(rw.version) {
				return
			}
			profileETime := time.Now().UnixNano()
			durationms := (time.Duration((profileETime - profileSTime)) / time.Millisecond)
			a := audit.Audit{
				LogTime:    time.Now().Format(time.RFC3339),
				Method:     rw.creq.context.Method,
				StatusCode: int(rw.resp.HTTPStatus),
				Path:       audit.GetAuditPath(),
				Duration:   int64(durationms),
				ReqId:      rw.creq.context.RequestID,
				URL:        rw.creq.context.URI,
			}
			//Add profiling details in header.
			//https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Server-Timing
			rw.resp.Headers["Server-Timing"] = fmt.Sprintf("total;dur=%d", a.Duration)
			go func(a audit.Audit) {
				defer util.PanicHandler("Got Panic in Auditlog goroutine")
				err := a.Logit()
				if err != nil {
					logger.Error(fmt.Sprintf("Error in writing Audit logs, Error: %s", err.Error()))
				}
			}(a)

		}
		return
	}
}
