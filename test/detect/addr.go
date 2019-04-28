package main

import (
	"dial/detect"
	"dial/gen-go/rpc/dial/yamutech/com"
	//"fmt"
)

func main() {
	ipsec := &com.SysIpSec{}
	//addr := &com.IpAddr{Addr: "192.168.6.22", Version: 4}
	//addr := &com.IpAddr{Addr: "192.168.5.55", Version: 4}
	//addr := &com.IpAddr{Addr: "115.239.211.112", Version: 4}
	//ipsec.Ipsec = &com.IpsecAddress{IP: addr, Mask: 24}

	//detect.LanDetect(ipsec, 5)
	//detect.Detect(ipsec, 5)
	//detect.AddrDetect(ipsec, 5)




	addr := &com.IpAddr{Addr: "192.168.6.22", Version: 4}
	//addr := &com.IpAddr{Addr: "115.239.211.112", Version: 4}
	ipsec.Ipsec = &com.IpsecAddress{IP: addr, Mask: 24}
	ipsec.RecordId = "xxx"
	detect.AddIpSec(ipsec,5)
	//detect.IpsecDetect("xxx",detect.IpsecMap["xxx"])
	select{}
}
