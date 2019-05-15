package main

import (
	//"dial/gen-go/rpc/dial/yamutech/com"
	"dial/client"
	"dial/log"
	"dial/server"
	"github.com/sevlyar/go-daemon"
	"os"
)

func main() {
	if 2 == len(os.Args) && os.Args[1] == "-b" {
		cntxt := &daemon.Context{
			PidFileName: "dial.pid",
			PidFilePerm: 0644,
			LogFileName: "dial.log",
			LogFilePerm: 0640,
			WorkDir:     "/var/log/dial",
			Umask:       027,
			Args:        []string{"[dial]"},
		}

		d, err := cntxt.Reborn()
		if err != nil {
			log.Error.Println("Unable to run: ", err)
		}
		if d != nil {
			return
		}
		defer cntxt.Release()
	}

	// go server.Work()
	go client.Client.Register()
	server.Server()
	select {}
}
