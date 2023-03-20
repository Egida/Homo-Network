package cnc

import (
	"bytes"
	"fmt"
	random "math/rand"
	"net"
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

	login := make([]byte, 1024)

	_, err := conn.Read(login)

	login = bytes.ReplaceAll(login, []byte{0}, []byte{})
	login = bytes.ReplaceAll(login, []byte{13}, []byte{})
	login = bytes.ReplaceAll(login, []byte{10}, []byte{})

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

	capt := GenCaptcha()

	conn.Write([]byte((gradient.Rainbow().Apply("[Homo-Network]") + color.HiWhiteString(" [ "+capt+" ] Enter captcha: "))))

	captcha := make([]byte, 1024)

	_, err = conn.Read(captcha)

	captcha = bytes.ReplaceAll(captcha, []byte{0}, []byte{})
	captcha = bytes.ReplaceAll(captcha, []byte{13}, []byte{})
	captcha = bytes.ReplaceAll(captcha, []byte{10}, []byte{})
	if err != nil {
		fmt.Println(err)
	}

	if string(captcha) != capt {
		Print(color.HiRedString("Access denied\n"), conn)
		DeadSession(conn)
		conn.Close()
		return
	}

	if err != nil {
		conn.Close()
		return
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

func GenCaptcha() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")

	count := make([]rune, 4)
	for i := range count {
		count[i] = letters[random.Intn(len(letters))]
	}

	return string(count)
}
