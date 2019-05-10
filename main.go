package main

import (
	//"dial/client"
	//"dial/gen-go/rpc/dial/yamutech/com"
	"dial/log"
	"dial/server"
)

func main() {
	log.Info.Println("飞雪无情的博客:", "http://www.flysnow.org")
	log.Warning.Printf("飞雪无情的微信公众号：%s\n", "flysnow_org")
	log.Error.Println("欢迎关注留言")

	server.Server()
	select {}
}
