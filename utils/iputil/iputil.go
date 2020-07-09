package iputil

import (
	"fmt"
	"net"
)

var CurrentIp string

func init ()  {
	CurrentIp = GetCurrentIp()
}

func GetCurrentIp() string {
	interAddrList, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	} else {
		for _, interAddr := range interAddrList {
			if ipNet, ok := interAddr.(*net.IPNet);ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					return ipNet.IP.String()
				}
			}
		}
		return ""
	}
}
