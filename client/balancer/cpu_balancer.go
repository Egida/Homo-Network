package balancer

import (
	"log"
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func getCpuUsage() int {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatal(err)
	}
	return int(math.Ceil(percent[0]))
}
