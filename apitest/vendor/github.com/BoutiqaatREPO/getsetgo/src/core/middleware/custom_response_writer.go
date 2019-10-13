package middleware

import (
	"compress/gzip"
	"net/http"

	"github.com/BoutiqaatREPO/getsetgo/src/core/common/versionmanager"

	"github.com/newrelic/go-agent"

	utilhttp "github.com/BoutiqaatREPO/getsetgo/src/common/utils/http"
)

type CRequest struct {
	context  utilhttp.RequestContext
	resource string
	method   string
}

type GSGResponseWriter struct {
	http.ResponseWriter
	resp    *utilhttp.APIResponse
	doGzip  bool
	creq    *CRequest
	version versionmanager.BasicVersion
}

func (this *GSGResponseWriter) Flush() (int, error) {

	for key, val := range this.resp.Headers {
		this.Header().Set(key, val)
	}
	this.WriteHeader(int(this.resp.HTTPStatus))
	if !this.doGzip {
		return this.Write(this.resp.Body)
	}
	//gzip the response
	// We need to gzip the response.
	// @todo: instead of initializing a new writer each time, we should ideally
	// create a pool of gzip writers and use them.
	//
	//
	// Note: The application auto adjusts itself to correct the Content-Length header, incase its gzip.
	//
	//
	gzw := gzip.NewWriter(this.ResponseWriter)
	defer gzw.Close()
	return gzw.Write(this.resp.Body)
}

func (this *GSGResponseWriter) GetTransaction() newrelic.Transaction {
	if txn, ok := this.ResponseWriter.(newrelic.Transaction); ok && txn != nil {
		return txn
	}
	return nil
}

func EntryHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cw := &GSGResponseWriter{w, nil, false, &CRequest{}, versionmanager.BasicVersion{}}
		fn(cw, req)
	}
}
