package config

import (
	"github.com/widuu/goini"
)

type Conf struct {
	DialPort    int
	DialIp      string
	AgentPort   int
	AgentIp     string
	Health      int
	DelayWeight int
	LostWeight  int
	Count       int
	Timeout     int
	Interval    int
	Dname       string
}

const CONFPATH = "/home/heweiwei/go/src/dial/config/dial.ini"

var Cfg *Conf

func NewConf() *Conf {
	return &Conf{
		DialPort:    9092,
		AgentPort:   9294,
		Health:      5000,
		DelayWeight: 100,
		LostWeight:  2000,
		Count:       5,
		Timeout:     2,
		Interval:    10,
		Dname:       "www.baidu.com",
	}
}

func init() {
	Cfg = NewConf()
	conf := goini.SetConfig(CONFPATH)
	Cfg.DialIp = conf.GetValue("dial", "ip")
	Cfg.AgentIp = conf.GetValue("agent", "ip")
}
