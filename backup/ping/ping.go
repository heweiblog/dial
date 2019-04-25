package base

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Icmp struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func getIcmp(seq uint16) Icmp {
	icmp := Icmp{
		Type:        8,
		Code:        0,
		CheckSum:    0,
		Identifier:  0,
		SequenceNum: seq,
	}
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.CheckSum = Checksum(buffer.Bytes())
	buffer.Reset()
	return icmp
}

func sendIcmpRequest(icmp Icmp, destAddr *net.IPAddr) int64 {
	conn, err := net.DialIP("ip4:icmp", nil, destAddr)
	if err != nil {
		fmt.Printf("Fail to connect to remote host: %s\n", err)
		return 0
	}
	defer conn.Close()
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)
	if _, err := conn.Write(buffer.Bytes()); err != nil {
		return 0
	}
	t := time.Now()
	conn.SetReadDeadline((time.Now().Add(time.Second * 2)))
	recv := make([]byte, 1024)
	receiveCnt, err := conn.Read(recv)
	if err != nil {
		return 0
	}
	//fmt.Printf("%d bytes from %s: seq=%d \n", receiveCnt, destAddr.String(), icmp.SequenceNum)
	return time.Since(t).Nanoseconds() / 1000
}

func Checksum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)
	return uint16(^sum)
}

func IcmpPing(host string) int64 {
	raddr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		//fmt.Printf("Fail to resolve %s, %s\n", host, err)
		return 0
	}
	//fmt.Printf("Ping %s (%s):\n\n", raddr.String(), host)
	for i := 1; i < 3; i++ {
		if delay := sendIcmpRequest(getIcmp(uint16(i)), raddr); delay > 0 {
			return int64(delay)
		}
	}
	return 0
}
