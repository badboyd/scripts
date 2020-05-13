package main

/*
URL: https://github.com/mccoyst/myip/blob/master/myip.go
URL: http://changsijay.com/2013/07/28/golang-get-ip-address/
*/

import (
	"net"
	"os"
	"fmt"
)

func getMyIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		switch v := addr.(type) {
		case *net.IPNet, *net.IPAddr:
			return v.String()
		}
	}
	return ""
}

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}

	fmt.Println(getMyIP())
}
