package main

import (
	"dial/base"
	"fmt"
)

func main() {
	//fmt.Println(base.Snmp("192.168.6.190", "public", ".1.3.6.1.2.1", 161))
	fmt.Println(base.Snmp("192.168.6.191", "public", "", 161))
}
