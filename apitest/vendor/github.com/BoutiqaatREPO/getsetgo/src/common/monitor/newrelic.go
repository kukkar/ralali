package monitor

import (
	"os"
	"strings"

	newrelic "github.com/newrelic/go-agent"
)

type NewRelicConfig struct {
	AppId   string
	APIKey  string
	Enabled bool
	Debug   bool
}

//
// Initialize NewRelic.
//
func InitializeNewRelic(c NewRelicConfig) (*NewRelic, error) {
	config := newrelic.NewConfig(c.AppId, c.APIKey)
	if c.Debug {
		config.Logger = newrelic.NewDebugLogger(os.Stdout)
	}
	config.Enabled = c.Enabled
	app, err := newrelic.NewApplication(config)
	if err != nil {
		return nil, err
	}
	return &NewRelic{
		app:     app,
		enabled: c.Enabled,
	}, nil
}

type NewRelic struct {
	enabled bool
	app     newrelic.Application
}

func (this *NewRelic) GetApp() newrelic.Application {
	return this.app
}

// Implementation of GetRawClient()
func (this *NewRelic) GetRawClient() interface{} {
	return this.app
}

//Implementation of IsEnabled()
func (this *NewRelic) IsEnabled() bool {
	return this.enabled
}

//Implementation of SendMetric
func (this *NewRelic) SendMetric(name string, value float64, tags []string) error {
	if !this.enabled {
		return nil
	}
	name = strings.Replace(name, "-", "/", 12)
	return this.app.RecordCustomMetric(name, value)
}

//Implementation of RecordEvent
func (this *NewRelic) RecordEvent(name string, value float64, tags map[string]interface{}) error {
	if !this.enabled {
		return nil
	}
	tags["duration"] = int(value)
	name = strings.Replace(name, "-", ":", 12)
	return this.app.RecordCustomEvent(name, tags)
}
