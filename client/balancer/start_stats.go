package balancer

import "errors"

func GetStats() (latency string, cpu int, memory int, err error) {

	latency = GetLatency()
	cpu = getCpuUsage()
	memory = GetMemoryUsage()

	if latency == "" || cpu == 0 || memory == 0 {
		return "", 0, 0, errors.New("Error")
	}

	return latency, cpu, memory, nil
}
