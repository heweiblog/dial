package main

import (
	"dial/base"
	"fmt"
)

func main() {
	fmt.Println(base.Mysql("192.168.6.53:33066", "root", "root", "mysql", "select * from user"))
	//fmt.Println(base.Mysql("192.168.6.53:33066", "root", "root", "mysql", ""))
	//fmt.Println(base.Mysql("192.168.6.190:13636", "root", "root", "mysql", "select * from user"))
}
