package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mohamadafzal06/nextrandomip/cafd"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func extractIpv4(cidr string) []string {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatal(err)
	}

	ips := make([]string, 0, 256)
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	return ips
}

func main() {
	cidr := "192.168.1.0/24"
	ips := extractIpv4(cidr)
	iter := cafd.NewCAFD(extractIpv4(cidr))
	fmt.Println(len(ips))
	for i := 0; i < len(ips); i++ {
		fmt.Println(iter.Next())
		fmt.Println("----------------")
	}
}
