package utils

import (
	"net/http"
	"strings"
)

func Proxy(r *http.Request) []string {
	if ips := r.Header.Get("X-Forwarded-For"); ips != "" {
		return strings.Split(ips, ",")
	}
	return []string{}
}

func GetIp(r *http.Request) string {

	ips := Proxy(r)
	if len(ips) > 0 && ips[0] != "" {
		return ips[0]
	}
	ip := strings.Split(r.RemoteAddr, ":")
	if len(ip) > 0 {
		return ip[0]
	}

	//ip := r.Header.Get("RemoteIp")
	//if len(ip) > 0 {
	//	return ip
	//}
	return "0.0.0.0"
}
