package config

import (
	"encoding/json"
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/common/monitor"
	"github.com/BoutiqaatREPO/getsetgo/src/common/ratelimiter"
	"github.com/BoutiqaatREPO/getsetgo/src/common/utils/http"
)

// AppConfig will contain all the app related config data which should be provided at the start of the app
type AppConfig struct {
	AppName              string `json:"AppName"`
	AppVersion           string `json:"AppVersion"`
	ServerPort           string
	LogConfFile          string
	BlackBoxLogs         string
	AuditLogs            string
	AllowedHosts         string
	MonitorConfig        monitor.MConf
	Performance          PerformanceConfigs
	DynamicConfig        DynamicConfigInfo
	HTTPConfig           http.Config `json:"HttpConfig"`
	Profiler             ProfilerConfig
	ResponseHeaders      ResponseHeaderFields
	ApplicationConfig    interface{}
	AppRateLimiterConfig *ratelimiter.Config
}

func (this *AppConfig) String() string {
	str, err := json.Marshal(this)
	if err != nil {
		return "Could NOT Marshal APP config."
	}
	return string(str)
}

func (this *AppConfig) ShowConfig() string {

	s := fmt.Sprintf(
		"AppName      := %s\n"+
			"AppVersion   := %s\n"+
			"ServerPort   := %s\n"+
			"Log File     := %s\n"+
			"BlackBoxLogs := %s\n"+
			"AuditLogs    := %s\n"+
			"AllowedHosts := %s\n"+
			"Performance  := %+v\n"+
			"Dynamic      := %+v\n"+
			"HTTP Config  := %+v\n"+
			"Profiler     := %+v\n"+
			"Resp Headers := %+v\n"+
			"App Config   := %+v\n"+
			"Monitor Conf := %+v\n",
		this.AppName,
		this.AppVersion,
		this.ServerPort,
		this.LogConfFile,
		this.BlackBoxLogs,
		this.AuditLogs,
		this.AllowedHosts,
		this.Performance,
		this.DynamicConfig,
		this.HTTPConfig,
		this.Profiler,
		this.ResponseHeaders,
		this.ApplicationConfig,
		this.MonitorConfig,
	)

	return s
}

// PerformanceConfigs contains Garbage Collector detials, which will determine when the GC will kick
type PerformanceConfigs struct {
	UseCorePercentage float64
	GCPercentage      float64
}

// ResponseHeaderFields
type ResponseHeaderFields struct {
	CacheControl CacheControlHeaders
}

// CacheControlHeaders helps in telling the caller whether to cache the response and up to what time etc.
type CacheControlHeaders struct {
	ResponseType    string
	NoCache         bool
	NoStore         bool
	MaxAgeInSeconds int
}

// Application
type Application struct {
	ResponseHeaders ResponseHeaderFields
}

// DynamicConfigInfo
type DynamicConfigInfo struct {
	Active          bool
	RefreshInterval int
	ConfigKey       string
	CacheKey        string
}

// ProfilerConfig is used to profile the application, like the time taken for a request etc.
type ProfilerConfig struct {
	Enable       bool
	SamplingRate float64
}

// GlobalAppConfig is applicationconfig Singleton
var GlobalAppConfig = new(AppConfig)
