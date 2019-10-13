package profiler

import (
	"time"

	"github.com/BoutiqaatREPO/getsetgo/src/common/monitor"

	"github.com/BoutiqaatREPO/getsetgo/src/common/config"
)

//Niltime used when time is nil or no time captured
const (
	NilTime        = -1
	timeUnit int64 = int64(time.Millisecond)
)

var rate float64 // rate to be used for monitoring

type profiler struct {

	//Key specifies a key for a profiling event
	key string

	//StartTime specifies the timestamp in nano second when a profile event
	//was started
	startTime time.Time
}

// Initprofiler inits the profiler
func InitProfiler(r float64) {
	rate = r
}

//Newprofiler returns a new instance of profiler
func NewProfiler() *profiler {
	if !config.GlobalAppConfig.Profiler.Enable {
		return nil
	}
	return new(profiler)
}

//StartProfile starts a profiling using profiler instance p for key. Key should have the
//suggested name '<package-file-method>'.
func (p *profiler) StartProfile(key string) {
	if !config.GlobalAppConfig.Profiler.Enable {
		return
	}
	p.startTime = time.Now()
	p.key = key
}

//EndProfile ends the profiling using profiler instance p for key k. Return time in MicroSeconds
func (p *profiler) EndProfile() int64 {
	if !config.GlobalAppConfig.Profiler.Enable {
		return NilTime
	}
	d := time.Since(p.startTime).Nanoseconds()
	t := d / timeUnit
	p = nil
	return t
}

//EndProfile ends the profiling starting using profiler instance p for key k
func (p *profiler) EndProfilenRecord(name string, tags []string) int64 {
	t := p.EndProfile()
	if t != NilTime {
		monitor.GetInstance().SendMetric(name, float64(t), tags)
	}
	return t
}

func (p *profiler) EndProfilenRecordEvent(n string, tags map[string]interface{}) int64 {
	t := p.EndProfile()
	if t != NilTime {
		monitor.GetInstance().RecordEvent(n, float64(t), tags)
	}
	return t
}
