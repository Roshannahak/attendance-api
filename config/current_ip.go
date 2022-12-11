package config

import "net"

func GetCurrentIp() string {
	ifaces, err := net.Interfaces()

	if err != nil {
		println(err)
	}
	var ip net.IP
	for index, interf := range ifaces {
		addrs, _ := interf.Addrs()
		for j, addr := range addrs {
			if index == 2 && j == 0 {
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
					break
				}
				return ip.String()
			}
		}
	}
	return "localhost"
}
