package server

import (
	"dial/gen-go/rpc/dial/yamutech/com"
	"dial/log"
)

func (h Hand) AddSnmpGroupInfo(snmp *com.SnmpGroupInfo) (com.RetCode, error) {
	SnmpMap[snmp.Name] = snmp
	if s, ok := SnmpMap[snmp.Name]; ok {
		// go SnmpDetect(s)
		log.Cfglog.Println("add a new snmp success,snmp=", s)
		return com.RetCode_OK, nil
	}
	log.Cfglog.Println("add a new snmp failed,snmp=", snmp)
	return com.RetCode_FAIL, nil
}

func (h Hand) DelSnmpGroupInfo(snmp string) (com.RetCode, error) {
	delete(SnmpMap, snmp)
	log.Cfglog.Println("del a snmp success,snmp=", snmp)
	return com.RetCode_OK, nil
}
