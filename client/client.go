package client

import (
	"dial/config"
	"dial/gen-go/rpc/dial/yamutech/com"
	"dial/log"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	"sync"
	"time"
)

type DialClient struct {
	Client    *com.AgentClient
	Lock      *sync.Mutex
	Transport thrift.TTransport
	Status    bool
}

var Client *DialClient

func init() {
	Client = &DialClient{}
	transport, err := thrift.NewTSocket(net.JoinHostPort(config.Cfg.AgentIp, config.Cfg.AgentPort))
	if err != nil {
		log.Error.Println(err)
		os.Exit(1)
	}
	Client.Transport = transport
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	Client.Client = com.NewAgentClientFactory(transport, protocolFactory)
	Client.Lock = new(sync.Mutex)
}

func (c *DialClient) Register() {

	if err := c.Transport.Open(); err != nil {
		log.Error.Println(err)
		os.Exit(1)
	}

	time.Sleep(2 * time.Second)

	ip := com.NewIpAddr()
	ip.Version = 4
	ip.Addr = config.Cfg.AgentIp

	c.Lock.Lock()
	for {
		ret, err := c.Client.RegisterModule(com.ModuleType_DIALING)
		if ret != com.RetCode_OK || err != nil {
			time.Sleep(time.Second)
			log.Error.Println(err)
			continue
		}
		break
	}
	c.Status = true
	c.Lock.Unlock()

	log.Cfglog.Println("dial client success register agent->", config.Cfg.AgentIp, config.Cfg.AgentPort)
}
