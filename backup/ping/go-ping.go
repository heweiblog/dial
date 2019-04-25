package main

import (
	"bytes"
	//"dial/base"
	"dial/gen-go/rpc/dial/yamutech/com"
	"encoding/binary"
	"fmt"
	"github.com/sparrc/go-ping"
	"net"
	"sync"
	"time"
)

// 获取地址段开始结束ip  sip：eg 192.168.5.3 mask：eg 24 返回IP段起始结束ip
func GetAddr(sip string, mask int32) (uint32, uint32) {
	var hmask uint32 = 1
	var i int32
	for i = 0; i < mask; i++ {
		hmask *= 2
	}
	hmask -= 1
	hmask = ^hmask
	//fmt.Println(hmask)

	var ip uint32
	binary.Read(bytes.NewBuffer(net.ParseIP(sip).To4()), binary.BigEndian, &ip)

	//fmt.Println(ip)

	return ip&hmask + 1, (^hmask) | ip
}

// uint32 to ip string
func Uint32ToString(ipnr uint32) string {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0]).String()
}

func detect(ipnet *com.SysIpSec, interval int32) {
	if net.ParseIP(ipnet.Ipsec.IP.Addr) == nil || ipnet.Ipsec.Mask <= 0 || ipnet.Ipsec.Mask > 32 {
		return
	}
	var (
		iplist []*com.IpAddr
		//wg     sync.WaitGroup
		mu sync.Mutex
	)
	begin, end := GetAddr(ipnet.Ipsec.IP.Addr, 32-ipnet.Ipsec.Mask)
	fmt.Println("begin:", Uint32ToString(begin), "end:", Uint32ToString(end), "size:", end-begin)

	//size := 0
	for {
		for i := begin; i < end; i++ {
			//wg.Add(1)
			go func(ip string) {
				//defer wg.Done()
				pinger, err := ping.NewPinger(ip)
				if err != nil {
					return
				}
				pinger.SetPrivileged(true)
				pinger.Count = 3
				pinger.Timeout = 2 * time.Second
				pinger.Run()                 // blocks until finished
				stats := pinger.Statistics() // get send/receive/rtt stats
				//fmt.Println(ip, stats)
				if stats.AvgRtt.Nanoseconds() > 0 {
					addr := &com.IpAddr{Addr: ip, Version: 4}
					mu.Lock()
					iplist = append(iplist, addr)
					mu.Unlock()
				}
			}(Uint32ToString(i))
		}
		//wg.Wait()
		time.Sleep(10 * time.Second)
		fmt.Println("iplist=", len(iplist))
		iplist = iplist[:0]
		if interval > 0 {
			time.Sleep(time.Duration(interval) * time.Second)
		} else {
			return
		}
	}
}

func Ping(ip string, mu *sync.Mutex, wg *sync.WaitGroup, iplist *[]*com.IpAddr) {
	defer wg.Done()
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return
	}
	pinger.SetPrivileged(true)
	pinger.Count = 3
	pinger.Timeout = 2 * time.Second
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	fmt.Println(ip, stats)
	if stats.AvgRtt.Nanoseconds() > 0 {
		addr := &com.IpAddr{Addr: ip, Version: 4}
		mu.Lock()
		*iplist = append(*iplist, addr)
		mu.Unlock()
	}
}

func main() {
	/*
		ip := "1.2.3.4"
		pinger, err := ping.NewPinger(ip)
		if err != nil {
			return
		}
		pinger.SetPrivileged(true)
		pinger.Count = 1
		pinger.Timeout = 2 * time.Second
		pinger.Run()                 // blocks until finished
		stats := pinger.Statistics() // get send/receive/rtt stats
		fmt.Println(stats)
	*/

	ipsec := &com.SysIpSec{}
	addr := &com.IpAddr{Addr: "192.168.6.22", Version: 4}
	//addr := &com.IpAddr{Addr: "192.168.5.55", Version: 4}
	//addr := &com.IpAddr{Addr: "115.239.211.112", Version: 4}
	ipsec.Ipsec = &com.IpsecAddress{IP: addr, Mask: 24}

	detect(ipsec, 5)

}
