package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type data struct {
	name string
}

func main()  {
	ip, ipNet, err := net.ParseCIDR("fe01:0:0:0:0:ffff:192.168.1.10/120")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n %s \n%s \n%s\n", ip.String(), ipNet.String(), ipNet.IP.String(), ipNet.Mask.String())

	ipv6, _, err := net.ParseCIDR("240e:5a:6c01:1::/64")
	if err != nil {
		panic(err)
	}

	vpcId := 1000000000
	vpcIdHex := strconv.FormatInt(int64(vpcId), 16)
	fmt.Println(vpcIdHex)
	i := 11
	for vpcId > 0 {
		ipv6[i] = byte(vpcId % 256)
		vpcId >>= 8
		i--
	}

	startIP := "192.168.1.0"
	endIP := "192.168.1.255"
	gateway := "192.168.1.1"

	startIpv4 := net.ParseIP(startIP)
	for i := 12; i < 16; i++ {
		ipv6[i] = startIpv4[i]
	}
	startIPv6 := ipv6.String()
	fmt.Printf("startIPv6= %s\n", startIPv6)

	endIpv4 := net.ParseIP(endIP)
	for i := 12; i < 16; i++ {
		ipv6[i] = endIpv4[i]
	}
	endIPv6 := ipv6.String()
	fmt.Printf("endIPv6= %s\n", endIPv6)

	ipv4Gateway := net.ParseIP(gateway)
	for i := 12; i < 16; i++ {
		ipv6[i] = ipv4Gateway[i]
	}
	ipv6Gateway := ipv6.String()
	fmt.Printf("ipv6Gateway= %s\n", ipv6Gateway)

	fmt.Printf("%s\n", ipv6)

	strs := strings.Split("abcdef", ",")
	for _, str := range strs {
		fmt.Println(str)
	}
	return
}
