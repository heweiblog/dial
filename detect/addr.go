package detect

import (
	"bytes"
	"dial/gen-go/rpc/dial/yamutech/com"
	"encoding/binary"
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	//"sync"
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

// 网络地址段探测
func AddrDetect(ipnet *com.SysIpSec, interval int32) {
	var (
		//wg     sync.WaitGroup
		//mu     sync.Mutex
		iplist []*com.IpAddr
	)
	// 计算网段ip区间 使用uint32表示
	//fmt.Println("ip:", ipnet.Ipsec.IP.Addr, "32-mask:", 32-ipnet.Ipsec.Mask)
	begin, end := GetAddr(ipnet.Ipsec.IP.Addr, 32-ipnet.Ipsec.Mask)
	/*
		fmt.Println("begin:", Uint32ToString(begin))
		fmt.Println("end:", Uint32ToString(end))
		fmt.Println("size:", end-begin)
	*/

	p := fastping.NewPinger()
	p.OnRecv = func(ipaddr *net.IPAddr, rtt time.Duration) {
		//fmt.Printf("IP Addr: %s receive, RTT: %v\n", ipaddr.String(), rtt)
		addr := &com.IpAddr{Addr: ipaddr.IP.String(), Version: 4}
		//fmt.Println(addr)
		//mu.Lock()
		//ping成功ip加入上报列表iplist
		iplist = append(iplist, addr)
		//mu.Unlock()
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
			//fmt.Println(len(iplist))
			// 上报iplist 先检查原ipsec节点是否存在 不存在直接退出
		}

		// clear iplist
		iplist = iplist[:0]

		// interval 为 0 只拨测一次
		if interval > 0 {
			time.Sleep(time.Duration(interval) * time.Second)
		} else {
			return
		}
	}
}
