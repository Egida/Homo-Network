package balancer

import (
	"strconv"
)

var BalanceCh = make(chan bool)

func Balancer(cpu int, latency string, memory int) {

	go func() {
		for {
			cp := getCpuUsage()
			mem := GetMemoryUsage()
			lat := GetLatency()

			lts, _ := strconv.Atoi(latency)
			ltcheck, _ := strconv.Atoi(lat)

			if cpu-cp < 0 || memory-mem < 0 || lts-ltcheck < 0 {
				BalanceCh <- true
				//
			}
		}
	}()

}
