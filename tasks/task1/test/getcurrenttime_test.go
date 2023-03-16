package getcurrenttime

import (
	"task1/pkg/getcurrenttime"
	"testing"
	"time"
)

var server = "0.beevik-ntp.pool.ntp.org"
var server_duration = time.Duration(1 * time.Second)

func TestGetCurrentTime(t *testing.T) {
	systemCurrentTime := time.Now()

	moduleCurrentTime, err := getcurrenttime.GetCurrentTime(server)
	if err != nil {
		t.Error(err)
	}

	if moduleCurrentTime.Sub(systemCurrentTime) <= -server_duration ||
		moduleCurrentTime.Sub(systemCurrentTime) >= server_duration {
		t.Errorf("Time incorrect: (sys) %v != (module) %v", systemCurrentTime, moduleCurrentTime)
	}
}
