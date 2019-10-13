package shortenpost

import (
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"
	"github.com/BoutiqaatREPO/getsetgo/src/common/ratelimiter"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/utils/healthcheck"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/versionmanager"
)

type ShotenPostAPI struct {
}

func (a *ShotenPostAPI) GetVersion() versionmanager.Version {
	return versionmanager.Version{
		Resource: "SHORTEN",
		Version:  "V1",
		Action:   "POST",
		BucketID: constants.OrchestratorBucketDefaultValue, //todo - should it be a constant
		Path:     "",
	}
}

func (a *ShotenPostAPI) GetOrchestrator() orchestrator.Orchestrator {
	logger.Info("Hello World Pipeline Creation begin")

	helloWorldOrchestrator := new(orchestrator.Orchestrator)
	helloWorldWorkflow := new(orchestrator.WorkFlowDefinition)
	helloWorldWorkflow.Create()

	//Creation of the nodes in the workflow definition
	helloWorldNode := new(ShotenPost)
	helloWorldNode.SetID("hello world node 1")
	eerr := helloWorldWorkflow.AddExecutionNode(helloWorldNode)
	if eerr != nil {
		logger.Error(fmt.Sprintln(eerr))
	}

	//Set start node for the search workflow
	helloWorldWorkflow.SetStartNode(helloWorldNode)

	//Assign the workflow definition to the Orchestrator
	helloWorldOrchestrator.Create(helloWorldWorkflow)

	logger.Info(helloWorldOrchestrator.String())
	logger.Info("Hello World Pipeline Created")
	logger.Info("Hello World Pipeline Created")
	return *helloWorldOrchestrator
}

func (a *ShotenPostAPI) GetHealthCheck() healthcheck.HCInterface {
	return new(ShotenPostAPIHealthCheck)
}

func (a *ShotenPostAPI) Init() {
	//api initialization should come here
}

func (a *ShotenPostAPI) GetRateLimiter() ratelimiter.RateLimiter {
	return nil
}
