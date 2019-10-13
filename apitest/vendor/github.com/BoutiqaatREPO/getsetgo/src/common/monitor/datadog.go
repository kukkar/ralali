package monitor

import (
	_ "expvar"
	"fmt"

	"github.com/ooyala/go-dogstatsd"
)

type DataDogConfig struct {
	APIKey      string
	APPKey      string
	AgentServer string
	Enabled     bool
}

//
// Initialize Datadog.
//
func InitializeDatadog(c DataDogConfig) (*DataDogAgentClient, error) {
	clientD, err := dogstatsd.New(c.AgentServer)
	if err != nil {
		return nil, err
	}
	return &DataDogAgentClient{
		client:  clientD,
		enabled: c.Enabled,
	}, nil
}

// DataDogAgent client is the implementation of a client to talk with
// the datadog agent
type DataDogAgentClient struct {
	client  *dogstatsd.Client
	enabled bool
}

// Implementation of IsEnabled()
func (d *DataDogAgentClient) IsEnabled() bool {
	return d.enabled
}

// Implementation of GetRawClient()
func (d *DataDogAgentClient) GetRawClient() interface{} {
	return d.client
}

// Implementation of SendMetric()
func (this *DataDogAgentClient) SendMetric(name string, value float64, tags []string) error {
	if !this.enabled {
		return nil
	}
	return this.client.Gauge(name, value, tags, 1)
}

//Implementation of RecordEvent()
func (this *DataDogAgentClient) RecordEvent(name string, value float64, tags map[string]interface{}) error {
	if !this.enabled {
		return nil
	}
	return fmt.Errorf("To be impl")
}
