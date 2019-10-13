package service

import (
	"github.com/BoutiqaatREPO/getsetgo/src/common/ratelimiter"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/utils/healthcheck"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/versionmanager"
)

type APIOptions struct {
	DisableBlackBoxLogs bool
	DisableAuditLogs    bool
}

type APIInterfaceCustom interface {
	APIInterface
	GetOptions() APIOptions
}

type APIInterface interface {
	GetVersion() versionmanager.Version

	GetOrchestrator() orchestrator.Orchestrator

	GetHealthCheck() healthcheck.HCInterface

	GetRateLimiter() ratelimiter.RateLimiter

	Init()
}
