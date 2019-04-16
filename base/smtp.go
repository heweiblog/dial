package base

import (
	"net"
	"net/smtp"
	"time"
)

// smtp拨测，返回时延，返回0表示拨测失败
func Smtp(addr string) int64 {
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		return 0
	}
	defer conn.Close()

	t := time.Now()
	client, err := smtp.Dial(addr)
	if err != nil {
		return 0
	}
	defer client.Close()

	err = client.Hello("localhost")
	if err != nil {
		return 0
	}
	return time.Since(t).Nanoseconds() / 1000
}
