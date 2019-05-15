package server

import (
	"dial/client"
	"dial/config"
	"dial/gen-go/rpc/dial/yamutech/com"
	"dial/log"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/gogf/gf/g/container/gmap"
	"os"
)

type Hand struct {
	com.Dial
}

type Srv struct {
	IP   *com.IpAddr
	Type com.DialServerType
}

type Ipsec struct {
	Ipsec    *com.SysIpSec
	Interval int32
}

var (
	ServerStatus bool
	Map          *gmap.StrAnyMap
	PolicyMap    map[string]*com.HealthPolicyInfo
	IpsecMap     map[string]*Ipsec
	SnmpMap      map[string]*com.SnmpGroupInfo
	HeartBeat    *com.HeartBeatState
)

func init() {
	IpsecMap = make(map[string]*Ipsec)
	SnmpMap = make(map[string]*com.SnmpGroupInfo)
	HeartBeat = com.NewHeartBeatState()
	PolicyMap = make(map[string]*com.HealthPolicyInfo)
	Map = gmap.NewStrAnyMap()
}

func (h Hand) SystemCommand(cmdType com.SysCommand) (com.RetCode, error) {
	return com.RetCode_OK, nil
}

// 是否启动服务
func (h Hand) SetServerState(enable bool) (com.RetCode, error) {
	ServerStatus = enable
	log.Cfglog.Println("SetServerState", ServerStatus)
	return com.RetCode_OK, nil
}

// 心跳
func (h Hand) HeartBeat() (*com.HeartBeatState, error) {
	if client.Client.Status {
		HeartBeat.MState = com.ModuleState_REGISTERED
	} else {
		HeartBeat.MState = com.ModuleState_STARTUP
	}
	HeartBeat.ServerState = ServerStatus
	log.Debug.Println("HeartBeat ServerStatus =", ServerStatus)
	return HeartBeat, nil
}

// 添加服务器组
func (h Hand) AddDialServer(rid com.ObjectId, ip *com.IpAddr, typ com.DialServerType) (com.RetCode, error) {
	s := &Srv{IP: ip, Type: typ}
	Map.Set(string(rid), s)
	log.Cfglog.Println("add a new server success,rid=", rid, "ip=", ip, "type=", typ)
	return com.RetCode_OK, nil
}

// 删除服务器组
func (h Hand) DelDialServer(rid com.ObjectId) (com.RetCode, error) {
	Map.Remove(string(rid))
	log.Cfglog.Println("del a server success,rid=", rid)
	return com.RetCode_OK, nil
}

func Server() {
	log.Cfglog.Println("Dial listen:", config.Cfg.DialIp+":"+config.Cfg.DialPort)
	serverTransport, err := thrift.NewTServerSocket(config.Cfg.DialIp + ":" + config.Cfg.DialPort)
	if err != nil {
		log.Error.Println("Error!", err)
		os.Exit(1)
	}
	processor := com.NewDialProcessor(Hand{})
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	log.Cfglog.Println("Dial thrift server success in", config.Cfg.DialIp+":"+config.Cfg.DialPort)

	server.Serve()
}
