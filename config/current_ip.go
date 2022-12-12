package config

import (
	"fmt"
	"net"
)

func GetCurrentIp() string {
	ifaces, err := net.Interfaces()

	if err != nil {
		println(err)
	}
	var ip net.IP
	for index, interf := range ifaces {
		addrs, _ := interf.Addrs()
		fmt.Printf("Full object : %+v \n", interf)
		for j, addr := range addrs {
			v := addr.(*net.IPNet)
			if index == 1 && j == 0 && v != nil {
				ip = v.IP
				return ip.String()
			} else if index == 2 && j == 0 && v != nil {
				ip = v.IP
				return ip.String()
			}

		}
	}
	return "localhost"
}
