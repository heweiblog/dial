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
	addr := &com.IpAddr{Addr: "115.239.211.112", Version: 4}
	ipsec.Ipsec = &com.IpsecAddress{IP: addr, Mask: 22}

	//detect.LanDetect(ipsec, 5)
	//detect.Detect(ipsec, 5)
	detect.NetDetect(ipsec, 5)
}
