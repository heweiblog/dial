package client

import (
	"dial/config"
	"dial/gen-go/rpc/dial/yamutech/com"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	"strconv"
	"sync"
)

type DialClient struct {
	Client    *com.AgentClient
	Lock      *sync.Mutex
	Transport thrift.TTransport
}

var Client *DialClient

func init() {
	Client = &DialClient{}
	transport, err := thrift.NewTSocket(net.JoinHostPort(config.Cfg.AgentIp, config.Cfg.AgentPort))
	if err != nil {
		// log
		fmt.Println(err)
		os.Exit(1)
	}
	Client.transport = transport
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	Client.Client = com.NewAgentClientFactory(transport, protocolFactory)
	Client.Lock = new(sync.Mutex)
}

func (c *DialClient) Register() {

	if err := c.Transport.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	time.Sleep(2 * time.Second)

	ip := com.NewIpAddr()
	ip.Version = 4
	ip.Addr = common.AgentIp

	c.Lock.Lock()
	for {
		ret, err := c.Client.RegisterModule(com.ModuleType_DIALING)
		if ret != com.RetCode_OK || err != nil {
			time.Sleep(time.Second)
			fmt.Println(err)
			continue
		}
		break
	}
	c.Lock.Unlock()

	fmt.Println("success register ->", config.Cfg.AgentIp, config.Cfg.AgentPort)
}
