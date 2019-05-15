package server

import (
	"dial/gen-go/rpc/dial/yamutech/com"
	"dial/log"
)

func (h Hand) AddDcInfo(dc *com.DcInfo) (com.RetCode, error) {
	Map.Set(dc.ID, dc)
	log.Cfglog.Println("add dc success", dc)
	return com.RetCode_OK, nil
}

func (h Hand) DelDcInfo(id string) (com.RetCode, error) {
	Map.Remove(id)
	log.Cfglog.Println("del dc success", id)
	return com.RetCode_OK, nil
}
