package balancer

import (
	"log"
	"math"

	"github.com/shirou/gopsutil/mem"
)

func GetMemoryUsage() int {
	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}
	return int(math.Ceil(memory.UsedPercent))
}
