package main

import (
	"dial/base"
	"fmt"
)

func main() {
	//fmt.Println(base.Ping("fe80::225:90ff:fec0:1745%eth1"))
	//fmt.Println(base.Ping("fe80::225:90ff:fe39:7f08%eth0"))
	//fmt.Println(base.Ping("1.2.4.2"))
	//fmt.Println(base.Ping("1.1.1.1"))
	fmt.Println(base.Icmp("1.1.1.1"))
}
