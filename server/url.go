package server

import (
	"dial/gen-go/rpc/dial/yamutech/com"
	"dial/log"
)

type Nginx struct {
	Policys []string
	Urls    []*com.DialNginxServer
}

// 添加url组 策略
func (h Hand) AddNginxGroup(groupName string, policyName string) (com.RetCode, error) {
	if ng, ok := Map.Get(groupName).(*Nginx); ok {
		ng.Policys = append(ng.Policys, policyName)
		Map.Set(groupName, ng)
		log.Cfglog.Println("add a policy to url success,group=", groupName, "policy=", policyName)
		return com.RetCode_OK, nil
	} else {
		ng := &Nginx{}
		ng.Policys = append(ng.Policys, policyName)
		log.Cfglog.Println("add a new url success,group=", groupName, "policy=", policyName)
		Map.Set(groupName, ng)
		return com.RetCode_OK, nil
	}
	log.Cfglog.Println("add a url failed,group=", groupName, "policy=", policyName)
	return com.RetCode_FAIL, nil
}

// 删除url 策略 组
func (h Hand) DelNginxGroup(groupName string, policyName string) (com.RetCode, error) {
	if ng, ok := Map.Get(groupName).(*Nginx); ok {
		if size := len(ng.Policys); size > 1 {
			for i := 0; i < size; i++ {
				if ng.Policys[i] == policyName {
					ng.Policys = append(ng.Policys[:i], ng.Policys[i+1:]...)
					log.Cfglog.Println("del a policy from url success,group=", groupName, "policy=", policyName)
					Map.Set(groupName, ng)
					return com.RetCode_OK, nil
				}
			}
		} else {
			Map.Remove(groupName)
			log.Cfglog.Println("del url success,group=", groupName)
			return com.RetCode_OK, nil
		}
	}
	log.Cfglog.Println("del a policy from url failed,group=", groupName, "policy=", policyName)
	return com.RetCode_FAIL, nil
}

// 添加server到url组
func (h Hand) AddNginxServer(groupName string, servers []*com.DialNginxServer) (com.RetCode, error) {
	if ng, ok := Map.Get(groupName).(*Nginx); ok {
		ng.Urls = append(ng.Urls, servers...)
		log.Cfglog.Println("add servers to url success,group=", groupName, "servers=", servers)
		return com.RetCode_OK, nil
	}
	log.Cfglog.Println("add servers to url failed,group=", groupName, "servers=", servers)
	return com.RetCode_FAIL, nil
}

// 从url组删除server
func (h Hand) DelNginxServer(groupName string, servers []*com.DialNginxServer) (com.RetCode, error) {
	if ng, ok := Map.Get(groupName).(*Nginx); ok {
		for _, r := range servers {
			for i := 0; i < len(ng.Urls); i++ {
				if r.LocalURL == ng.Urls[i].LocalURL {
					ng.Urls = append(ng.Urls[:i], ng.Urls[i+1:]...)
					break
				}
			}
		}
		log.Cfglog.Println("del servers from url success,group=", groupName, "servers=", servers)
		return com.RetCode_OK, nil
	}
	log.Cfglog.Println("del servers from url failed,group=", groupName, "servers=", servers)
	return com.RetCode_FAIL, nil
}
