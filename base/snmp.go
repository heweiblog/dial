package base

import (
	"github.com/soniah/gosnmp"
	"time"
)

// snmp拨测，默认2c版本，addr:ip，返回时延，返回0表示拨测失败
func Snmp(addr, community, oid string, port uint16) int64 {
	snmp := &gosnmp.GoSNMP{
		Target:    addr,
		Port:      port,
		Community: community,
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(2) * time.Second,
	}

	t := time.Now()
	err := snmp.Connect()
	if err != nil {
		return 0
	}
	defer snmp.Conn.Close()

	var oids []string
	if len(oid) > 0 {
		oids = []string{oid}
	} else {
		oids = []string{".1.3.6.1.2.1.1.5.0"}
	}
	_, err = snmp.Get(oids)
	if err != nil {
		return 0
	}

	return time.Since(t).Nanoseconds() / 1000
}
