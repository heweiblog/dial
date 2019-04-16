// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"rpc/dial/yamutech/com"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  RetCode registerModule(ModuleType typ)")
	fmt.Fprintln(os.Stderr, "  RetCode updateHealthStatus( results)")
	fmt.Fprintln(os.Stderr, "  RetCode updateServerStatus( results)")
	fmt.Fprintln(os.Stderr, "  RetCode updateDcStatus( results)")
	fmt.Fprintln(os.Stderr, "  RetCode updateNginxStatus( results)")
	fmt.Fprintln(os.Stderr, "  RetCode updateSysInfo(string snmp, SysInfo sysinfo)")
	fmt.Fprintln(os.Stderr, "  RetCode updateInterfaceInfo(string snmp,  interfaces)")
	fmt.Fprintln(os.Stderr, "  RetCode updateInterfaceTraffic(string snmp,  traffic)")
	fmt.Fprintln(os.Stderr, "  RetCode updateInterfaceIpMac(string snmp,  ipmac)")
	fmt.Fprintln(os.Stderr, "  RetCode updateRouteInfo(string snmp,  routeinfo)")
	fmt.Fprintln(os.Stderr, "  RetCode updateProcessInfo(string snmp, ProcessInfo processinfo)")
	fmt.Fprintln(os.Stderr, "  RetCode updateIpSecOnlineIp(string ipsecid,  iplist)")
	fmt.Fprintln(os.Stderr, "  RetCode updateMacTable(string snmp,  mactable)")
	fmt.Fprintln(os.Stderr, "  RetCode registerDialModule(i32 moduleId, IpAddr ip, i32 port)")
	fmt.Fprintln(os.Stderr, "  RetCode unRegisterDialModule(i32 moduleId)")
	fmt.Fprintln(os.Stderr, "  RetCode heartBeat(i32 moduleId)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := com.NewAgentClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "registerModule":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RegisterModule requires 1 args")
			flag.Usage()
		}
		tmp0, err := (strconv.Atoi(flag.Arg(1)))
		if err != nil {
			Usage()
			return
		}
		argvalue0 := com.ModuleType(tmp0)
		value0 := argvalue0
		fmt.Print(client.RegisterModule(value0))
		fmt.Print("\n")
		break
	case "updateHealthStatus":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "UpdateHealthStatus requires 1 args")
			flag.Usage()
		}
		arg53 := flag.Arg(1)
		mbTrans54 := thrift.NewTMemoryBufferLen(len(arg53))
		defer mbTrans54.Close()
		_, err55 := mbTrans54.WriteString(arg53)
		if err55 != nil {
			Usage()
			return
		}
		factory56 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt57 := factory56.GetProtocol(mbTrans54)
		containerStruct0 := com.NewAgentUpdateHealthStatusArgs()
		err58 := containerStruct0.ReadField1(jsProt57)
		if err58 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Results
		value0 := argvalue0
		fmt.Print(client.UpdateHealthStatus(value0))
		fmt.Print("\n")
		break
	case "updateServerStatus":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "UpdateServerStatus requires 1 args")
			flag.Usage()
		}
		arg59 := flag.Arg(1)
		mbTrans60 := thrift.NewTMemoryBufferLen(len(arg59))
		defer mbTrans60.Close()
		_, err61 := mbTrans60.WriteString(arg59)
		if err61 != nil {
			Usage()
			return
		}
		factory62 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt63 := factory62.GetProtocol(mbTrans60)
		containerStruct0 := com.NewAgentUpdateServerStatusArgs()
		err64 := containerStruct0.ReadField1(jsProt63)
		if err64 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Results
		value0 := argvalue0
		fmt.Print(client.UpdateServerStatus(value0))
		fmt.Print("\n")
		break
	case "updateDcStatus":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "UpdateDcStatus requires 1 args")
			flag.Usage()
		}
		arg65 := flag.Arg(1)
		mbTrans66 := thrift.NewTMemoryBufferLen(len(arg65))
		defer mbTrans66.Close()
		_, err67 := mbTrans66.WriteString(arg65)
		if err67 != nil {
			Usage()
			return
		}
		factory68 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt69 := factory68.GetProtocol(mbTrans66)
		containerStruct0 := com.NewAgentUpdateDcStatusArgs()
		err70 := containerStruct0.ReadField1(jsProt69)
		if err70 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Results
		value0 := argvalue0
		fmt.Print(client.UpdateDcStatus(value0))
		fmt.Print("\n")
		break
	case "updateNginxStatus":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "UpdateNginxStatus requires 1 args")
			flag.Usage()
		}
		arg71 := flag.Arg(1)
		mbTrans72 := thrift.NewTMemoryBufferLen(len(arg71))
		defer mbTrans72.Close()
		_, err73 := mbTrans72.WriteString(arg71)
		if err73 != nil {
			Usage()
			return
		}
		factory74 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt75 := factory74.GetProtocol(mbTrans72)
		containerStruct0 := com.NewAgentUpdateNginxStatusArgs()
		err76 := containerStruct0.ReadField1(jsProt75)
		if err76 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Results
		value0 := argvalue0
		fmt.Print(client.UpdateNginxStatus(value0))
		fmt.Print("\n")
		break
	case "updateSysInfo":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "UpdateSysInfo requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		arg78 := flag.Arg(2)
		mbTrans79 := thrift.NewTMemoryBufferLen(len(arg78))
		defer mbTrans79.Close()
		_, err80 := mbTrans79.WriteString(arg78)
		if err80 != nil {
			Usage()
			return
		}
		factory81 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt82 := factory81.GetProtocol(mbTrans79)
		argvalue1 := com.NewSysInfo()
		err83 := argvalue1.Read(jsProt82)
		if err83 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.UpdateSysInfo(value0, value1))
		fmt.Print("\n")
		break
	case "updateInterfaceInfo":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "UpdateInterfaceInfo requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		arg85 := flag.Arg(2)
		mbTrans86 := thrift.NewTMemoryBufferLen(len(arg85))
		defer mbTrans86.Close()
		_, err87 := mbTrans86.WriteString(arg85)
		if err87 != nil {
			Usage()
			return
		}
		factory88 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt89 := factory88.GetProtocol(mbTrans86)
		containerStruct1 := com.NewAgentUpdateInterfaceInfoArgs()
		err90 := containerStruct1.ReadField2(jsProt89)
		if err90 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Interfaces
		value1 := argvalue1
		fmt.Print(client.UpdateInterfaceInfo(value0, value1))
		fmt.Print("\n")
		break
	case "updateInterfaceTraffic":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "UpdateInterfaceTraffic requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		arg92 := flag.Arg(2)
		mbTrans93 := thrift.NewTMemoryBufferLen(len(arg92))
		defer mbTrans93.Close()
		_, err94 := mbTrans93.WriteString(arg92)
		if err94 != nil {
			Usage()
			return
		}
		factory95 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt96 := factory95.GetProtocol(mbTrans93)
		containerStruct1 := com.NewAgentUpdateInterfaceTrafficArgs()
		err97 := containerStruct1.ReadField2(jsProt96)
		if err97 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Traffic
		value1 := argvalue1
		fmt.Print(client.UpdateInterfaceTraffic(value0, value1))
		fmt.Print("\n")
		break
	case "updateInterfaceIpMac":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "UpdateInterfaceIpMac requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		arg99 := flag.Arg(2)
		mbTrans100 := thrift.NewTMemoryBufferLen(len(arg99))
		defer mbTrans100.Close()
		_, err101 := mbTrans100.WriteString(arg99)
		if err101 != nil {
			Usage()
			return
		}
		factory102 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt103 := factory102.GetProtocol(mbTrans100)
		containerStruct1 := com.NewAgentUpdateInterfaceIpMacArgs()
		err104 := containerStruct1.ReadField2(jsProt103)
		if err104 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Ipmac
		value1 := argvalue1
		fmt.Print(client.UpdateInterfaceIpMac(value0, value1))
		fmt.Print("\n")
		break
	case "updateRouteInfo":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "UpdateRouteInfo requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		arg106 := flag.Arg(2)
		mbTrans107 := thrift.NewTMemoryBufferLen(len(arg106))
		defer mbTrans107.Close()
		_, err108 := mbTrans107.WriteString(arg106)
		if err108 != nil {
			Usage()
			return
		}
		factory109 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt110 := factory109.GetProtocol(mbTrans107)
		containerStruct1 := com.NewAgentUpdateRouteInfoArgs()
		err111 := containerStruct1.ReadField2(jsProt110)
		if err111 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Routeinfo
		value1 := argvalue1
		fmt.Print(client.UpdateRouteInfo(value0, value1))
		fmt.Print("\n")
		break
	case "updateProcessInfo":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "UpdateProcessInfo requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		arg113 := flag.Arg(2)
		mbTrans114 := thrift.NewTMemoryBufferLen(len(arg113))
		defer mbTrans114.Close()
		_, err115 := mbTrans114.WriteString(arg113)
		if err115 != nil {
			Usage()
			return
		}
		factory116 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt117 := factory116.GetProtocol(mbTrans114)
		argvalue1 := com.NewProcessInfo()
		err118 := argvalue1.Read(jsProt117)
		if err118 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.UpdateProcessInfo(value0, value1))
		fmt.Print("\n")
		break
	case "updateIpSecOnlineIp":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "UpdateIpSecOnlineIp requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		arg120 := flag.Arg(2)
		mbTrans121 := thrift.NewTMemoryBufferLen(len(arg120))
		defer mbTrans121.Close()
		_, err122 := mbTrans121.WriteString(arg120)
		if err122 != nil {
			Usage()
			return
		}
		factory123 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt124 := factory123.GetProtocol(mbTrans121)
		containerStruct1 := com.NewAgentUpdateIpSecOnlineIpArgs()
		err125 := containerStruct1.ReadField2(jsProt124)
		if err125 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Iplist
		value1 := argvalue1
		fmt.Print(client.UpdateIpSecOnlineIp(value0, value1))
		fmt.Print("\n")
		break
	case "updateMacTable":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "UpdateMacTable requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		arg127 := flag.Arg(2)
		mbTrans128 := thrift.NewTMemoryBufferLen(len(arg127))
		defer mbTrans128.Close()
		_, err129 := mbTrans128.WriteString(arg127)
		if err129 != nil {
			Usage()
			return
		}
		factory130 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt131 := factory130.GetProtocol(mbTrans128)
		containerStruct1 := com.NewAgentUpdateMacTableArgs()
		err132 := containerStruct1.ReadField2(jsProt131)
		if err132 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Mactable
		value1 := argvalue1
		fmt.Print(client.UpdateMacTable(value0, value1))
		fmt.Print("\n")
		break
	case "registerDialModule":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "RegisterDialModule requires 3 args")
			flag.Usage()
		}
		tmp0, err133 := (strconv.Atoi(flag.Arg(1)))
		if err133 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		arg134 := flag.Arg(2)
		mbTrans135 := thrift.NewTMemoryBufferLen(len(arg134))
		defer mbTrans135.Close()
		_, err136 := mbTrans135.WriteString(arg134)
		if err136 != nil {
			Usage()
			return
		}
		factory137 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt138 := factory137.GetProtocol(mbTrans135)
		argvalue1 := com.NewIpAddr()
		err139 := argvalue1.Read(jsProt138)
		if err139 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		tmp2, err140 := (strconv.Atoi(flag.Arg(3)))
		if err140 != nil {
			Usage()
			return
		}
		argvalue2 := int32(tmp2)
		value2 := argvalue2
		fmt.Print(client.RegisterDialModule(value0, value1, value2))
		fmt.Print("\n")
		break
	case "unRegisterDialModule":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "UnRegisterDialModule requires 1 args")
			flag.Usage()
		}
		tmp0, err141 := (strconv.Atoi(flag.Arg(1)))
		if err141 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.UnRegisterDialModule(value0))
		fmt.Print("\n")
		break
	case "heartBeat":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "HeartBeat requires 1 args")
			flag.Usage()
		}
		tmp0, err142 := (strconv.Atoi(flag.Arg(1)))
		if err142 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.HeartBeat(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
