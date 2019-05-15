package server

import (
	"dial/gen-go/rpc/dial/yamutech/com"
	"dial/log"
)

func (h Hand) AddHealthPolicy(policy *com.HealthPolicyInfo) (com.RetCode, error) {
	PolicyMap[policy.Name] = policy
	if _, ok := PolicyMap[policy.Name]; ok {
		log.Cfglog.Println("add success->", policy)
		return com.RetCode_OK, nil
	}
	log.Cfglog.Println("add failed->", policy)
	return com.RetCode_FAIL, nil
}

func (h Hand) ModHealthPolicy(policy *com.HealthPolicyInfo) (com.RetCode, error) {
	PolicyMap[policy.Name] = policy
	if _, ok := PolicyMap[policy.Name]; ok {
		log.Cfglog.Println("mod success->", policy)
		return com.RetCode_OK, nil
	}
	log.Cfglog.Println("mod failed->", policy)
	return com.RetCode_FAIL, nil
}

func (h Hand) DelHealthPolicy(policy *com.HealthPolicyInfo) (com.RetCode, error) {
	delete(PolicyMap, policy.Name)
	if _, ok := PolicyMap[policy.Name]; ok {
		log.Cfglog.Println("del failed->", policy)
		return com.RetCode_FAIL, nil
	}
	log.Cfglog.Println("del success->", policy)
	return com.RetCode_OK, nil
}
