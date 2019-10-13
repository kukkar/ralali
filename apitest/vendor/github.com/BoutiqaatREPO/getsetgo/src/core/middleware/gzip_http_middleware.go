package middleware

import (
	"net/http"
	"strings"
)

func GzipHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if isPreFlightRequest(r) {
			fn(w, r)
			return
		}

		if rw, rok := w.(*GSGResponseWriter); rok {
			if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
				rw.resp.Headers["Content-Encoding"] = "gzip"
				rw.doGzip = true
				fn(rw, r)
				return
			}
		}
		fn(w, r)
	}
}
