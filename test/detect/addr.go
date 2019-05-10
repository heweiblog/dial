package main

import (
	"dial/detect"
	"dial/gen-go/rpc/dial/yamutech/com"
	//"fmt"
	"time"
)

func main() {
	//ipsec := &com.SysIpSec{}
	//addr := &com.IpAddr{Addr: "192.168.6.22", Version: 4}
	//addr := &com.IpAddr{Addr: "192.168.5.55", Version: 4}
	//addr := &com.IpAddr{Addr: "115.239.211.112", Version: 4}
	//ipsec.Ipsec = &com.IpsecAddress{IP: addr, Mask: 24}

	//detect.LanDetect(ipsec, 5)
	//detect.Detect(ipsec, 5)
	//detect.AddrDetect(ipsec, 5)

	ipsec := &com.SysIpSec{}
	addr := &com.IpAddr{Addr: "192.168.6.22", Version: 4}
	//addr := &com.IpAddr{Addr: "115.239.211.112", Version: 4}
	ipsec.Ipsec = &com.IpsecAddress{IP: addr, Mask: 24}
	ipsec.RecordId = "xxx"
	detect.AddIpSec(ipsec, 5)
	//detect.IpsecDetect("xxx",detect.IpsecMap["xxx"])
	time.Sleep(10 * time.Second)
	detect.DelIpSec("xxx")
	ips := &com.SysIpSec{}
	ad := &com.IpAddr{Addr: "115.239.211.112", Version: 4}
	ips.Ipsec = &com.IpsecAddress{IP: ad, Mask: 24}
	ipsec.RecordId = "xxx"
	detect.AddIpSec(ips, 5)
	select {}
}
