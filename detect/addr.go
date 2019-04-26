package detect

import (
	"bytes"
	"dial/base"
	"dial/gen-go/rpc/dial/yamutech/com"
	"encoding/binary"
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	"sync"
	"time"
)

var IpsecMap map[string]*com.SysIpSec

func AddIpSec(ipsec *com.SysIpSec, interval int32) (r com.RetCode, err error) {
	IpsecMap[ipsec.RecordId] = ipsec
	go AddrDetect(ipsec, interval)
	return com.RetCode_OK, nil
}

func DelIpSec(ipsecid string) (r com.RetCode, err error) {
	delete(IpsecMap, ipsecid)
	return com.RetCode_OK, nil
}

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

// 地址段在线ip探测 速度较快 适合局域网探测
func LanDetect(ipsec *com.SysIpSec, interval int32) {
	// 判断参数是否合法
	if net.ParseIP(ipsec.Ipsec.IP.Addr) == nil || ipsec.Ipsec.Mask <= 0 || ipsec.Ipsec.Mask > 32 {
		return
	}
	var iplist []*com.IpAddr

	// 计算网段ip区间 使用uint32表示
	begin, end := GetAddr(ipsec.Ipsec.IP.Addr, 32-ipsec.Ipsec.Mask)
	//fmt.Println("begin:", Uint32ToString(begin),"end:", Uint32ToString(end),"size:", end-begin)

	p := fastping.NewPinger()
	p.OnRecv = func(ipaddr *net.IPAddr, rtt time.Duration) {
		//fmt.Printf("IP Addr: %s receive, RTT: %v\n", ipaddr.String(), rtt)
		addr := &com.IpAddr{Addr: ipaddr.IP.String(), Version: 4}
		//fmt.Println(addr)
		//ping成功ip加入上报列表iplist
		iplist = append(iplist, addr)
	}

	for i := begin; i < end; i++ {
		ip := Uint32ToString(i)
		if ra, err := net.ResolveIPAddr("ip4:icmp", ip); err == nil {
			p.AddIPAddr(ra)
		}
	}

	for {
		err := p.Run()
		if err != nil {
			return
		} else {
			fmt.Println(len(iplist))
			// 上报iplist 先检查原ipsec节点是否存在 不存在直接退出
		}

		// clear iplist
		iplist = iplist[:0]

		// interval 为 0 只探测一次
		if interval > 0 {
			time.Sleep(time.Duration(interval) * time.Second)
		} else {
			return
		}
	}
}

// ip段探测 内网外网均可
func AddrDetect(ipsec *com.SysIpSec, interval int32) {
	if net.ParseIP(ipsec.Ipsec.IP.Addr) == nil || ipsec.Ipsec.Mask <= 0 || ipsec.Ipsec.Mask > 32 {
		return
	}
	var (
		iplist []*com.IpAddr
		wg     sync.WaitGroup
		mu     sync.Mutex
	)
	begin, end := GetAddr(ipsec.Ipsec.IP.Addr, 32-ipsec.Ipsec.Mask)
	//fmt.Println("begin:", Uint32ToString(begin), "end:", Uint32ToString(end), "size:", end-begin)

	addrs := make([]string, 0, end-begin)
	for i := begin; i < end; i++ {
		addrs = append(addrs, Uint32ToString(i))
	}

	for {
		for _, v := range addrs {
			wg.Add(1)
			go func(ip string) {
				defer wg.Done()
				for i := 0; i < 3; i++ {
					//if base.Ping(ip) > 0 {
					if base.Icmp(ip) > 0 {
						host := &com.IpAddr{Addr: ip, Version: 4}
						mu.Lock()
						iplist = append(iplist, host)
						mu.Unlock()
						return
					}
				}
			}(v)
		}
		wg.Wait()

		/*
			if _, ok := IpsecMap[ipsec.RecordId]; !ok {
				return
			}
		*/

		//上报
		fmt.Println(len(iplist))
		iplist = iplist[:0]
		if interval > 0 {
			time.Sleep(time.Duration(interval) * time.Second)
		} else {
			return
		}
	}
}
