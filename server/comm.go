package server

import (
	"dial/client"
	"dial/gen-go/rpc/dial/yamutech/com"
	"dial/log"
	"github.com/gogf/gf/g/container/gmap"
)

type Hand struct {
	com.Dial
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
	HeartBeat    *com.HeartBeatState
)

func init() {
	IpsecMap = make(map[string]*Ipsec)
	HeartBeat = com.NewHeartBeatState()
	PolicyMap = make(map[string]*com.HealthPolicyInfo)
	Map = gmap.NewStrAnyMap()
}

func (h Hand) SystemCommand(cmdType com.SysCommand) (com.RetCode, error) {
	return com.RetCode_OK, nil
}

func (h Hand) SetServerState(enable bool) (com.RetCode, error) {
	ServerStatus = enable
	log.Cfglog.Println("SetServerState", ServerStatus)
	return com.RetCode_OK, nil
}

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
