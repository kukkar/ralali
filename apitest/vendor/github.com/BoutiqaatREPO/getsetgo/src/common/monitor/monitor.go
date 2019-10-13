package monitor

import (
	"fmt"
)

// monitorObj singleton object
var monitorObj MInterface

//stores if the default monitor is enabled or not.
var isDefaultEnabled bool

// GetInstance returns a singleton instance of type MInterface
func GetInstance() MInterface {
	// return
	return monitorObj
}

func IsDefaultEnabled() bool {
	return isDefaultEnabled
}

//
// Exposed Method to directly connect with NewRelic Agent.
//
func GetNewRelicAgent() (*NewRelic, error) {
	if monitorObj == nil {
		return nil, fmt.Errorf("GetNewRelicAgent()->New Relic Monitor is not initialized")
	}
	newRelic, ok := monitorObj.(*NewRelic)
	if !ok {
		return nil, fmt.Errorf("GetNewRelicAgent()->Expected Monitoring object of Type (*NewRelic), Got:  %T", monitorObj)
	}
	if newRelic == nil {
		return nil, fmt.Errorf("GetNewRelicAgent()->New Relic Object is nil")
	}
	return newRelic, nil
}

func InitializeDefault(cnfg *MConf) error {
	if !cnfg.Enabled {
		return nil
	}
	i, err := Initialize(cnfg)
	if err != nil {
		return err
	}
	monitorObj = i
	if monitorObj.IsEnabled() {
		isDefaultEnabled = true
	}
	return nil
}

// Initialize initialize the monitor object
func Initialize(cnfg *MConf) (MInterface, error) {

	switch cnfg.Platform {

	case PLATFORM_TYPE_DATADOG:
		ddogConf := cnfg.DataDogConfig
		ddogConf.Enabled = cnfg.Enabled
		return InitializeDatadog(ddogConf)

	case PLATFORM_TYPE_NEWRELIC:
		newRConf := cnfg.NewRelicConfig
		newRConf.Enabled = cnfg.Enabled
		newRConf.Debug = cnfg.Debugging
		return InitializeNewRelic(newRConf)

	default:
		return nil, fmt.Errorf("Invalid Monitoring Platform supplied")
	}

}
