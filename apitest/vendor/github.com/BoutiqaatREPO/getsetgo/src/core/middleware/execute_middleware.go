package middleware

import (
	"fmt"
	"net/http"

	"github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	utilhttp "github.com/BoutiqaatREPO/getsetgo/src/common/utils/http"
	workflow "github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/versionmanager"
)

func ExecuteHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		if isPreFlightRequest(req) {
			fn(w, req)
			return
		}
		io, derr := GetData(req)
		if derr != nil {
			fmt.Fprintf(w, "Error %v", derr)
			return
		}

		//
		// Define PreExecution Settings.
		//
		if rw := getCoreWriter(w); rw != nil {
			newRelictx := rw.GetTransaction()
			if newRelictx != nil {
				io.ExecContext.Set(constants.NewRelicTransaction, newRelictx)
			}
		}

		serviceVersion, _, _, gerr := versionmanager.Get("SERVICE", "V1", "GET", constants.OrchestratorBucketDefaultValue, "")

		if gerr != nil {
			fmt.Fprintf(w, "Error %v", gerr)
			return
		}

		serviceOrchestrator, ok := serviceVersion.(workflow.Orchestrator)
		if !ok {
			fmt.Fprintf(w, "Could not get service orchestrator")
			return
		}

		//Start Orchestration.
		output := serviceOrchestrator.Start(io)
		responseI, _ := output.IOData.Get(constants.APIResponse)
		var response utilhttp.APIResponse
		if v, ok := responseI.(utilhttp.APIResponse); ok {
			response = v
		}

		//Attach Response data.
		rw := getCoreWriter(w)
		rw.resp = &response

		//Attach API details.
		resourcec, _ := output.ExecContext.Get(constants.Resource)
		if resourcestr, ok := resourcec.(string); ok {
			rw.creq.resource = resourcestr
		}

		//Attach ReqContext
		if rcon, err := output.ExecContext.GetRequestContext(); err == nil {
			rw.creq.context = rcon
		}

		//Attach version
		versionI, _ := output.ExecContext.Get("BASIC-VERSION")
		if version, ok := versionI.(versionmanager.BasicVersion); ok {
			rw.version = version
		}

		fn(rw, req)
		return
	}
}
