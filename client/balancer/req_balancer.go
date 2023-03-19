package balancer

import (
	"strconv"
	"time"

	"github.com/go-ping/ping"
)

func GetLatency() string {
	pinger, err := ping.NewPinger("www.google.com")
	if err != nil {
		return ""
	}
	pinger.Count = 1
	pinger.Timeout = 1 * time.Second
	pinger.Debug = false
	pinger.Run()
	if pinger.Statistics().Rtts == nil {
		return ""
	}
	return strconv.Itoa(int(pinger.Statistics().Rtts[0].Milliseconds()))
}
