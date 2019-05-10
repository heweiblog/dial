package server

import (
	"dial/config"
	"dial/gen-go/rpc/dial/yamutech/com"
	"dial/log"
	"git.apache.org/thrift.git/lib/go/thrift"
	"os"
)

type Hand struct {
	com.Dial
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
