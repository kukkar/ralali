package audit

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/BoutiqaatREPO/getsetgo/src/core/common/versionmanager"

	"github.com/BoutiqaatREPO/getsetgo/src/common/config"
	"github.com/sanksons/gowraps/filesystem"
)

const IDENTIFIER = "auditlogs"
const FILE_NAME = "audit"

//
// Check if Audit log is enabled.
// @todo: its better to use flushing approach instead of writing to file instantly.
// @todo: Add log roller and zipper.
//
func IsAuditEnabled() bool {
	if config.GlobalAppConfig.AuditLogs == "" {
		return false
	}
	return true
}

//
// get Audit log path.
//
func GetAuditPath() string {
	return config.GlobalAppConfig.AuditLogs
}

//
// Init Auditing Files.
//
func InitAuditing() error {
	return filesystem.CreateDirTree(GetAuditPath()+string(os.PathSeparator)+IDENTIFIER, 0775)
}

//
// Contains Information for which API's BlackBox is disabled.
//
var AuditAPIDisableMap map[versionmanager.BasicVersion]bool = make(map[versionmanager.BasicVersion]bool, 0)

//
// Check if Audit Log is Disabled for API.
//
func IsDisabledForAPI(v versionmanager.BasicVersion) bool {

	if v, ok := AuditAPIDisableMap[v]; ok && v == true {
		return true
	}
	return false
}

//
// Audit Logging.
//
type Audit struct {
	LogTime    string
	Method     string
	StatusCode int
	Path       string
	Duration   int64 //milliseconds
	ReqId      string
	URL        string
}

func (this *Audit) Logit() error {
	data := fmt.Sprintf("%s\t%s\t%d\t%s\t%s\t%s\n",
		this.LogTime,
		strings.ToUpper(this.Method),
		this.StatusCode,
		fmt.Sprintf("%d", this.Duration),
		this.ReqId,
		this.URL,
	)
	//check if file already exists, if not create one.
	return filesystem.AppendToFile(this.getFilePath(), []byte(data))
}

func (this *Audit) getFilePath() string {
	str := time.Now().Local().Format("2006-01-02")
	return this.Path + string(os.PathSeparator) + IDENTIFIER + string(os.PathSeparator) + FILE_NAME + "-" + str + ".log"
}
