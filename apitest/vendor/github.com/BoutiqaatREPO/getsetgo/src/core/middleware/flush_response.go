package middleware

import (
	"net/http"
)

func Flush(w http.ResponseWriter, r *http.Request) {
	if rw, rok := w.(*GSGResponseWriter); rok {
		rw.Flush()
	}
}
