package responseheaders

import (
	"fmt"
	"strings"

	"github.com/BoutiqaatREPO/getsetgo/src/core/common/utils/misc"

	"github.com/BoutiqaatREPO/getsetgo/src/common/config"
	"github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	utilhttp "github.com/BoutiqaatREPO/getsetgo/src/common/utils/http"
	workflow "github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
)

const (
	publicResponse  string = "public"
	privateResponse string = "private"
	noCache         string = "no-cache"
	noStore         string = "no-store"

	//http headers
	cacheControl string = "Cache-Control"
	contentType  string = "Content-Type"
	maxAge       string = "max-age"
)

//IdentifierExecutor joins the result retrieved from multiple nodes
type Writer struct {
	id string
}

func (r *Writer) SetID(id string) {
	r.id = id
}

func (r *Writer) GetID() (id string, err error) {
	return r.id, nil
}

func (r *Writer) Name() string {
	return "Response Header Writer"
}

func (r *Writer) Execute(io workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	rh, _ := io.IOData.Get(constants.ResponseHeadersConfig)
	rdata, err1 := io.IOData.Get(constants.APIResponse)

	var res utilhttp.APIResponse
	if err1 == nil {
		res = rdata.(utilhttp.APIResponse)
	} else {
		res = utilhttp.NewAPIResponse()
	}

	if rh != nil {
		if resHeaderConf, ok := rh.(config.ResponseHeaderFields); ok {
			cHeaders := resHeaderConf.CacheControl
			res.Headers[cacheControl] = getCacheControlHeader(cHeaders)
		}
	}

	io.IOData.Set(constants.APIResponse, res)

	rcontentType := r.GetContentType(io)
	io.IOData.Set(constants.ResponseType, rcontentType)

	res.Headers[contentType] = r.GetContentTypeHeader(rcontentType)

	return io, nil
}

func (r *Writer) GetContentType(io workflow.WorkFlowData) string {
	rtype, _ := io.IOData.Get(constants.ResponseType)
	if responseType, ok := rtype.(string); ok {
		return responseType
	}
	req, _ := misc.GetRequestFromIO(io)
	accept := strings.ToLower(req.Headers.Accept)
	switch accept {
	case "text/csv":
		return constants.RESPONSE_TYPE_CSV
	case "image/jpeg":
		return constants.RESPONSE_TYPE_IMG_JPG
	case "image/gif":
		return constants.RESPONSE_TYPE_IMG_GIF
	case "image/png":
		return constants.RESPONSE_TYPE_IMG_PNG
	default:
		return constants.RESPONSE_TYPE_JSON
	}
}

func (r *Writer) GetContentTypeHeader(ctype string) string {
	masterData := map[string]string{
		constants.RESPONSE_TYPE_CSV:     "text/csv",
		constants.RESPONSE_TYPE_JSON:    "application/json",
		constants.RESPONSE_TYPE_IMG_JPG: "image/jpeg",
		constants.RESPONSE_TYPE_IMG_GIF: "image/gif",
		constants.RESPONSE_TYPE_IMG_PNG: "image/png",
	}

	if val, ok := masterData[ctype]; ok {
		return val
	}
	return masterData[constants.RESPONSE_TYPE_JSON]

}

func getCacheControlHeader(cHeaders config.CacheControlHeaders) string {
	c := ""

	var params []string

	if cHeaders.ResponseType == publicResponse {
		params = append(params, publicResponse)
	} else if cHeaders.ResponseType == privateResponse {
		params = append(params, privateResponse)
	}

	if cHeaders.NoCache {
		params = append(params, noCache)
	}

	if cHeaders.NoStore {
		params = append(params, noStore)
	}

	if cHeaders.MaxAgeInSeconds > 0 {
		params = append(params, fmt.Sprintf("%s=%d", maxAge, cHeaders.MaxAgeInSeconds))
	}

	if len(params) > 0 {
		c = strings.Join(params, ", ")
	}
	return c
}
