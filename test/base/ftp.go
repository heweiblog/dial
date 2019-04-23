package main

import (
	"dial/base"
	"fmt"
)

func main() {
	//fmt.Println(base.Ftp("192.168.8.253:21", "anonymous", "anonymous"))
	fmt.Println(base.Ftp("192.168.8.253:21", "", ""))
	//fmt.Println(base.Ftp("149.20.1.49:21", "anonymous", "anonymous"))
	//fmt.Println(base.Ftp("149.20.1.49:21", "", ""))
}
