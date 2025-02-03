package ip

import (
	"net"
	"strconv"
	"strings"
)

type IPPortServer struct {
	ip   string
	port string
}

func CreatePortServer(port int) *IPPortServer {
	ip := &IPPortServer{port: strconv.FormatInt(int64(port), 32)}
	inters, _ := net.Interfaces()
	for _, i := range inters {
		if i.Flags&net.FlagLoopback > 0 ||
			i.Flags&net.FlagRunning == 0 ||
			strings.Contains(i.Name, "usb") ||
			strings.Contains(i.Name, "vmnet") {
			continue
		}

		addrs, _ := i.Addrs()
		if len(addrs) == 0 {
			panic("no addrs")
		}
		ip.ip = strings.Split(addrs[0].String(), "/")[0]
	}
	return ip
}

func (ip *IPPortServer) GetIP() string {
	return ip.ip
}

func (ip *IPPortServer) GetPort() string {
	return ip.port
}
