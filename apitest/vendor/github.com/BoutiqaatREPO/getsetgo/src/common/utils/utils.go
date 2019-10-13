package utils

import (
	"fmt"
	"strings"

	"github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	workflow "github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
	newrelic "github.com/newrelic/go-agent"
	uuid "github.com/satori/go.uuid"
)

//
// Utility function to get NewRelic Transaction from Execution Context.
// Note: the transaction will be nil, incase monitoring is disabled.
//
func GetNewRelicTransaction(io workflow.WorkFlowData) newrelic.Transaction {
	val, err := io.ExecContext.Get(constants.NewRelicTransaction)
	if err != nil {
		return nil
	}
	if txn, ok := val.(newrelic.Transaction); ok {
		return txn
	}
	return nil
}

//
// Generate a NewRequest ID.
// To be used while firing requests to external clients.
//
func GetNewReqId(appName string) string {
	return fmt.Sprintf("%s-%s", appName, uuid.Must(uuid.NewV4()).String())
}

//
//Trim to string space form start and end
//
func Trim(nameToTrim string) string {
	return strings.TrimSpace(nameToTrim)
}
