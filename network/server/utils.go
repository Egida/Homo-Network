package server

import "net"

func RemoveConn(array []net.Conn, elem net.Conn) []net.Conn {

	ret := make([]net.Conn, 0)

	for _, s := range array {
		if s != elem {
			ret = append(ret, s)
		}
	}
	return ret
}
