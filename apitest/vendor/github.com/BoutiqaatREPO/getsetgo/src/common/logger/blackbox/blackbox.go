package blackbox

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/BoutiqaatREPO/getsetgo/src/core/common/versionmanager"

	"github.com/BoutiqaatREPO/getsetgo/src/common/config"
	utilhttp "github.com/BoutiqaatREPO/getsetgo/src/common/utils/http"
)

const REQUEST_FILE_NAME = "request.log"
const RESPONSE_FILE_NAME = "response.log"

const IDENTIFIER = "blackbox"

//
// Factory to get black box log.
//
func GetNewBlackBox(path, reqId string, req string, resp string) *BlackBoxLog {
	bbox := &BlackBoxLog{
		ReqId:        reqId,
		RequestData:  req,
		ResponseData: resp,
	}
	bbox.setFullPath(path)
	return bbox
}

func IsBlackBoxEnabled() bool {
	if config.GlobalAppConfig.BlackBoxLogs == "" {
		return false
	}
	return true
}

func GetBlackBoxPath() string {
	return config.GlobalAppConfig.BlackBoxLogs
}

func GetFormattedRequest(r *http.Request) string {

	dataBytes, err := httputil.DumpRequest(r, true)
	if err == nil {
		return string(dataBytes)
	}
	return ""
}

//
// Prepare a formatted response.
//
func GetFormattedResponse(resp *utilhttp.APIResponse) string {

	var response []string

	statusCode := fmt.Sprintf("%d \n", resp.HTTPStatus)
	response = append(response, statusCode)

	for k, v := range resp.Headers {
		response = append(response, fmt.Sprintf("%s: %s", k, v))
	}

	response = append(response, fmt.Sprintf("\n\n %s", resp.Body))
	return strings.Join(response, "\n")
}

//
// Contains Information for which API's BlackBox is disabled.
//
var BlackBoxAPIDisableMap map[versionmanager.BasicVersion]bool = make(map[versionmanager.BasicVersion]bool, 0)

//
// Check If Disabled for API
//
func IsDisabledForAPI(v versionmanager.BasicVersion) bool {

	if v, ok := BlackBoxAPIDisableMap[v]; ok && v == true {
		return true
	}
	return false
}

//
// Prepare BlackBox
//
type BlackBoxLog struct {
	FullPath     string
	ReqId        string
	RequestData  string
	ResponseData string
}

//
// Prepare full path for Logfile.
//
func (this *BlackBoxLog) setFullPath(path string) {
	now := time.Now()
	nowtime := now.Format("2006-01-02")

	this.FullPath = strings.Join([]string{
		path,
		IDENTIFIER,
		nowtime,
	},
		string(os.PathSeparator),
	)
}

func (this *BlackBoxLog) getFullPath() string {
	return this.FullPath
}

//
// Log the file.
//
func (this *BlackBoxLog) LogIt() error {
	err := os.MkdirAll(this.getFullPath(), 0775)
	if err != nil {
		return err
	}
	filePath := this.getFullPath() + string(os.PathSeparator) + this.ReqId + ".log"
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()
	_, err = w.WriteString(this.RequestData)
	if err != nil {
		return err
	}
	w.WriteString("\n\n\n\n")
	_, err = w.WriteString(this.ResponseData)
	if err != nil {
		return err
	}
	return nil
}

//
// Prepare a formatted request.
//

//
// write request
//
func (this *BlackBoxLog) writeRequest() error {
	requestFile := this.getFullPath() + string(os.PathSeparator) + REQUEST_FILE_NAME
	return this.writeFile(requestFile, this.RequestData)
}

//
// write response
//
func (this *BlackBoxLog) writeResponse() error {
	responseFile := this.getFullPath() + string(os.PathSeparator) + RESPONSE_FILE_NAME
	return this.writeFile(responseFile, this.ResponseData)
}

//
// write File.
//
func (this *BlackBoxLog) writeFile(filePath string, data string) error {

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
