package server

import (
	"dial/gen-go/rpc/dial/yamutech/com"
	"dial/log"
)

type Auth struct {
	Policys []string
	Records []*com.DialRecord
}

// 添加授权组 策略
func (h Hand) AddHealthGroup(groupName string, policyName string) (com.RetCode, error) {
	if au, ok := Map.Get(groupName).(*Auth); ok {
		au.Policys = append(au.Policys, policyName)
		Map.Set(groupName, au)
		log.Cfglog.Println("add a policy to auth success,group=", groupName, "policy=", policyName)
		return com.RetCode_OK, nil
	} else {
		au := &Auth{}
		au.Policys = append(au.Policys, policyName)
		log.Cfglog.Println("add a new auth success,group=", groupName, "policy=", policyName)
		Map.Set(groupName, au)
		return com.RetCode_OK, nil
	}
	log.Cfglog.Println("add a auth failed,group=", groupName, "policy=", policyName)
	return com.RetCode_FAIL, nil
}

// 删除授权 策略 组
func (h Hand) DelHealthGroup(groupName string, policyName string) (com.RetCode, error) {
	if au, ok := Map.Get(groupName).(*Auth); ok {
		if size := len(au.Policys); size > 1 {
			for i := 0; i < size; i++ {
				if au.Policys[i] == policyName {
					au.Policys = append(au.Policys[:i], au.Policys[i+1:]...)
					log.Cfglog.Println("del a policy from auth success,group=", groupName, "policy=", policyName)
					Map.Set(groupName, au)
					return com.RetCode_OK, nil
				}
			}
		} else {
			Map.Remove(groupName)
			log.Cfglog.Println("del auth success,group=", groupName)
			return com.RetCode_OK, nil
		}
	}
	log.Cfglog.Println("del a policy from auth failed,group=", groupName, "policy=", policyName)
	return com.RetCode_FAIL, nil
}

// 添加记录到授权组
func (h Hand) AddHealthRecord(groupName string, records []*com.DialRecord) (com.RetCode, error) {
	if au, ok := Map.Get(groupName).(*Auth); ok {
		au.Records = append(au.Records, records...)
		log.Cfglog.Println("add records to auth success,group=", groupName, "records=", records)
		return com.RetCode_OK, nil
	}
	log.Cfglog.Println("add records to auth failed,group=", groupName, "records=", records)
	return com.RetCode_FAIL, nil
}

// 从授权组删除记录
func (h Hand) DelHealthRecord(groupName string, records []*com.DialRecord) (com.RetCode, error) {
	if au, ok := Map.Get(groupName).(*Auth); ok {
		for _, r := range records {
			for i := 0; i < len(au.Records); i++ {
				if r.Rid == au.Records[i].Rid {
					au.Records = append(au.Records[:i], au.Records[i+1:]...)
					break
				}
			}
		}
		log.Cfglog.Println("del records from auth success,group=", groupName, "records=", records)
		return com.RetCode_OK, nil
	}
	log.Cfglog.Println("del records from auth failed,group=", groupName, "records=", records)
	return com.RetCode_FAIL, nil
}
