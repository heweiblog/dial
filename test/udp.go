package main

import (
	"dial/base"
	"fmt"
)

func main() {
	//fmt.Println(base.Udp("192.168.6.190", 53))
	//fmt.Println(base.Udp("192.168.6.190", 68))
	fmt.Println(base.Udp("192.168.5.30", 53))
	fmt.Println(base.Udp("192.168.6.54", 53))
	//fmt.Println(base.Udp("192.168.6.190:69"))
	//fmt.Println(base.Udp("192.168.6.190:161"))
	fmt.Println(base.Udp("192.168.6.190",2222))
	fmt.Println(base.Udp("192.168.6.195", 22222))
}
