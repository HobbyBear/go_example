package main

import (
	"net"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/server2", func(writer http.ResponseWriter, request *http.Request) {
		// var r *http.Request
		var ip = ClientPublicIP(request)
		_, err := writer.Write([]byte(ip))
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8083", nil)

	ch := make(chan int)
	<-ch
}

func ClientPublicIP(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" && !HasLocalIP(net.ParseIP(ip)) {
			return ip
		}
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" && !HasLocalIP(net.ParseIP(ip)) {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		if !HasLocalIP(net.ParseIP(ip)) {
			return ip
		}
	}

	return ""
}

// HasLocalIP 检测 IP 地址是否是内网地址
// 通过直接对比ip段范围效率更高，详见：https://github.com/thinkeridea/go-extend/issues/2
func HasLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}
