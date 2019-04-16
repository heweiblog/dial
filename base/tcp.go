package base

import (
	"net"
	"time"
)

// tcp拨测，addr:ip:port，返回时延，返回0表示拨测失败
func Tcp(addr string) int64 {
	t := time.Now()
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		return 0
	}
	defer conn.Close()
	return time.Since(t).Nanoseconds() / 1000
}
