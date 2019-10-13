package getstats

import (
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"
	"github.com/BoutiqaatREPO/getsetgo/src/common/ratelimiter"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/utils/healthcheck"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/versionmanager"
)

type ShortenStatsAPI struct {
}

func (a *ShortenStatsAPI) GetVersion() versionmanager.Version {
	return versionmanager.Version{
		Resource: "SHORTEN",
		Version:  "V1",
		Action:   "GET",
		BucketID: constants.OrchestratorBucketDefaultValue, //todo - should it be a constant
		Path:     "{shortcode}/stats",
	}
}

func (a *ShortenStatsAPI) GetOrchestrator() orchestrator.Orchestrator {
	logger.Info("Hello World Pipeline Creation begin")

	ShortenUrlOrchestrator := new(orchestrator.Orchestrator)
	ShortenUrlWorkflow := new(orchestrator.WorkFlowDefinition)
	ShortenUrlWorkflow.Create()

	//Creation of the nodes in the workflow definition
	ShortenUrlNode := new(ShortenUrl)
	ShortenUrlNode.SetID("hello world node 1")
	eerr := ShortenUrlWorkflow.AddExecutionNode(ShortenUrlNode)
	if eerr != nil {
		logger.Error(fmt.Sprintln(eerr))
	}

	//Set start node for the search workflow
	ShortenUrlWorkflow.SetStartNode(ShortenUrlNode)

	//Assign the workflow definition to the Orchestrator
	ShortenUrlOrchestrator.Create(ShortenUrlWorkflow)

	logger.Info(ShortenUrlOrchestrator.String())
	logger.Info("Hello World Pipeline Created")
	logger.Info("Hello World Pipeline Created")
	return *ShortenUrlOrchestrator
}

func (a *ShortenStatsAPI) GetHealthCheck() healthcheck.HCInterface {
	return new(ShortenUrlStatsHealthCheck)
}

func (a *ShortenStatsAPI) Init() {
	//api initialization should come here
}

func (a *ShortenStatsAPI) GetRateLimiter() ratelimiter.RateLimiter {
	return nil
}
