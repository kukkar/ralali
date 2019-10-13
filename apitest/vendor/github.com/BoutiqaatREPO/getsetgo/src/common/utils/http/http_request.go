package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Method string

const (
	GET    Method = "GET"
	PUT    Method = "PUT"
	POST   Method = "POST"
	DELETE Method = "DELETE"
	PATCH  Method = "PATCH"
)

// Request represents all the request related data
type Request struct {
	HTTPVerb        Method
	URI             string
	OriginalRequest *http.Request
	Headers         RequestHeader
	PathParameters  *map[string]string
}

func getMethod(method string) (Method, error) {
	switch strings.ToUpper(method) {
	case "GET":
		return GET, nil
	case "PUT":
		return PUT, nil
	case "POST":
		return POST, nil
	case "DELETE":
		return DELETE, nil
	case "PATCH":
		return PATCH, nil
	}
	return "", errors.New("Incorrect HTTP Method")
}

func GetRequest(r *http.Request) (Request, error) {
	httpVerb, verr := getMethod(r.Method)
	if verr != nil {
		return Request{}, verr
	}

	return Request{
		HTTPVerb:        httpVerb,
		URI:             r.URL.String(),
		OriginalRequest: r,
		Headers:         GetReqHeader(r)}, nil
}

// GetBodyParameter returns the Body Parameters from the OriginalRequest as string
func (req *Request) GetBodyParameter() (string, error) {
	return getBodyParam(req.OriginalRequest)
}

// GetPathParameter returns the value corresponding to the key in Path parameter map as string
func (req *Request) GetPathParameter(key string) string {
	pathParams := req.PathParameters
	if pathParams != nil {
		if val, ok := (*pathParams)[key]; ok {
			return val
		}
	}
	return ""
}

// GetHeaderParameter returns the value for given header key
func (req *Request) GetHeaderParameter(key string) string {
	return req.OriginalRequest.Header.Get(key)
}

// Get Body Data without draining body.
func (this *Request) GetBody() ([]byte, error) {
	b := this.OriginalRequest.Body

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(b); err != nil {
		this.OriginalRequest.Body = b
		return nil, err
	}
	if err := b.Close(); err != nil {
		this.OriginalRequest.Body = b
		return nil, err
	}
	this.OriginalRequest.Body = ioutil.NopCloser(bytes.NewReader(buf.Bytes()))
	return buf.Bytes(), nil
}

// Load Body into supplied struct.
func (this *Request) LoadBody(payload interface{}) error {

	dataBytes, err := this.GetBody()
	if err != nil {
		return fmt.Errorf("*Request#LoadBody -> %s", err.Error())
	}
	return json.Unmarshal(dataBytes, payload)
}
