package config

import (
	"github.com/widuu/goini"
)

type Conf struct {
	DialPort    string
	DialIp      string
	AgentPort   string
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
		Health:      5000,
		DelayWeight: 100,
		LostWeight:  2000,
		Count:       5,
		Timeout:     2,
		Interval:    10,
	}
}

func init() {
	Cfg = NewConf()
	conf := goini.SetConfig(CONFPATH)
	Cfg.DialIp = conf.GetValue("dial", "ip")
	Cfg.DialPort = conf.GetValue("dial", "port")
	Cfg.AgentIp = conf.GetValue("agent", "ip")
	Cfg.AgentPort = conf.GetValue("agent", "port")
	Cfg.Dname = conf.GetValue("server", "dname")
}
