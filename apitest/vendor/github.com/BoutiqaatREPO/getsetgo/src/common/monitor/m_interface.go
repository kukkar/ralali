package monitor

var _ MInterface = (*DataDogAgentClient)(nil)
var _ MInterface = (*NewRelic)(nil)

//
// Monitor Interface to be used by External clients.
//
type MInterface interface {

	// Get Access to raw client.
	GetRawClient() interface{}

	// Check if monitoring is enabled.
	IsEnabled() bool

	// Push Metric data.
	// name should be provided seperated by "/"
	// sample: service-GET-Resource (only alphanumeric supported)
	SendMetric(name string, value float64, tags []string) error

	// Record and event.
	// name should be provided seperated by "/"
	// sample: service-GET-Resource (only alphanumeric supported)
	RecordEvent(name string, value float64, tags map[string]interface{}) error
}
