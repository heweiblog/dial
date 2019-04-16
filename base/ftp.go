package base

import (
	"github.com/jlaffaye/ftp"
	"time"
)

// ftp 拨测 addr:"ip:port" 无user,pass只检测ftp端口是否开放，返回时延 0表示失败
func Ftp(addr, user, pass string) int64 {
	t := time.Now()
	c, err := ftp.DialWithOptions(addr, ftp.DialWithTimeout(2*time.Second))
	if err != nil {
		return 0
	}
	defer c.Quit()

	if len(user) > 0 && len(pass) > 0 {
		err = c.Login(user, pass)
		if err != nil {
			return 0
		}
	}

	return time.Since(t).Nanoseconds() / 1000
}
