package monitor

//
// Configuration for Monitor.
//
type MConf struct {
	Platform       string
	Enabled        bool
	Debugging      bool
	NewRelicConfig NewRelicConfig
	DataDogConfig  DataDogConfig
}
