package main

import (
	"fmt"
	g "github.com/soniah/gosnmp"
	"time"
)

const (
	InterfaceNumOid = ".1.3.6.1.2.1.2.1.0"
)

// GoSNMP v2 默认超时2s
func NewSnmpv2(ip, community string, port uint16) *g.GoSNMP {
	return &g.GoSNMP{
		Target:    ip,
		Port:      port,
		Community: community,
		Version:   g.Version2c,
		Timeout:   time.Duration(2) * time.Second,
	}
}

// 必须connect之后调用
func GetNum(snmp *g.GoSNMP, oid string) int {
	oids := []string{oid}

	result, err := snmp.Get(oids)
	if err != nil {
		return 0
	}

	for _, v := range result.Variables {
		fmt.Printf("oid: %s ", v.Name)
		switch v.Type {
		case g.OctetString:
			fmt.Printf("string: %s\n", string(v.Value.([]byte)))
		default:
			//fmt.Printf("number: %d\n", g.ToBigInt(v.Value))
			fmt.Println(" number:", g.ToBigInt(v.Value))
			//return int(v.Value)
		}
	}
	return 0
}

func SnmpGet(addr, community, oid string, port uint16) {

	snmp := NewSnmpv2(addr, community, port)
	err := snmp.Connect()
	if err != nil {
		return
	}
	defer snmp.Conn.Close()

	num := GetNum(snmp, InterfaceNumOid)
	fmt.Println("interface num =", num)
}

func SnmpWalk(addr, community, oid string, port uint16) {

	snmp := NewSnmpv2(addr, community, port)
	err := snmp.Connect()
	if err != nil {
		return
	}
	defer snmp.Conn.Close()

	fn := func(v g.SnmpPDU) error {
		fmt.Printf("oid: %s, value: ", v.Name)
		switch v.Type {
		case g.OctetString:
			fmt.Printf("%s\n", string(v.Value.([]byte)))
		default:
			fmt.Printf("%d\n", g.ToBigInt(v.Value))
		}
		return nil
	}
	err = snmp.Walk(oid, fn)
	if err != nil {
		fmt.Printf("Walk() err: %vi\n", err)
	}
}

func SnmpWalkAll(addr, community, oid string, port uint16) {

	snmp := NewSnmpv2(addr, community, port)
	err := snmp.Connect()
	if err != nil {
		return
	}
	defer snmp.Conn.Close()

	result, err := snmp.WalkAll(oid)
	if err != nil {
		fmt.Printf("Walk() err: %v\n", err)
	}

	for _, v := range result {
		fmt.Printf("oid: %s, value: ", v.Name)
		switch v.Type {
		case g.OctetString:
			fmt.Printf("%s\n", string(v.Value.([]byte)))
		default:
			fmt.Printf("%d\n", g.ToBigInt(v.Value))
		}
	}
}

func main() {
	SnmpGet("192.168.127.1", "yamu.com", ".1.3.6.1.2.1.2.1.0", 161)
	//SnmpWalk("192.168.127.1", "yamu.com", ".1.3.6.1.2.1.2.2.1.2", 161)
	//SnmpWalkAll("192.168.127.1", "yamu.com", ".1.3.6.1.2.1.2.2.1.2", 161)
}
