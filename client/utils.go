package main

import (
	"math/rand"
	"strings"
	"time"
)

func RandomInt(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func GetUserAgent() string {

	uag := strings.Join(UserAgents, "\n")
	strs := strings.Split(uag, "\n")
	cStrs := len(strs)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	d := strs[r1.Intn(cStrs)]

	return d

}
