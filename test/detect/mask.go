package main

import (
	"dial/base"
	"dial/detect"
	"fmt"
	"net"
	"sync"
)

func Uint32ToString(ipnr uint32) string {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0]).String()
}

func main() {
	b, e := detect.GetAddr("192.168.6.0", 32-26)
	//b, e := detect.GetAddr("115.239.211.112", 32-24)
	fmt.Println(b, Uint32ToString(b))
	fmt.Println(e, Uint32ToString(e))
	fmt.Println(e - b)
	var wg sync.WaitGroup
	for i := b; i < e; i++ {
		ip := Uint32ToString(i)
		//fmt.Println(ip)
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 3; i++ {
				if delay := base.Ping(ip); delay > 0 {
					fmt.Println(ip, delay)
					return
				}
			}
		}()
	}
	wg.Wait()
}
