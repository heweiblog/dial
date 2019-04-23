package main

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	"os"
	"time"
)

func main() {
	ip := "192.168.6.195"
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ip)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	ra, err = net.ResolveIPAddr("ip4:icmp", "192.168.6.55")
	p.AddIPAddr(ra)
	ra, err = net.ResolveIPAddr("ip4:icmp", "1.1.1.1")
	p.AddIPAddr(ra)
	ra, err = net.ResolveIPAddr("ip4:icmp", "192.168.6.190")
	p.AddIPAddr(ra)

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
}
