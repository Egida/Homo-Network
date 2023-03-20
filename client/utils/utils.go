package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
func RandomInt(strlen int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "1234567890"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	conv, _ := strconv.Atoi(string(result))

	return conv
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
