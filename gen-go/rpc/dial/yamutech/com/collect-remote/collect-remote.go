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
	fmt.Fprintln(os.Stderr, "  RetCode registerModule(i32 moduleId, IpAddr ip, i32 port)")
	fmt.Fprintln(os.Stderr, "  RetCode unRegisterModule(i32 moduleId)")
	fmt.Fprintln(os.Stderr, "  RetCode heartBeat(i32 moduleId)")
	fmt.Fprintln(os.Stderr, "  RetCode reportTaskProcess(i32 moduleId, string taskId, TaskProcessArgs arg)")
	fmt.Fprintln(os.Stderr, "  RetCode reportResult(i32 moduleId, string taskId, string batchno,  resultList)")
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
	client := com.NewCollectClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "registerModule":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "RegisterModule requires 3 args")
			flag.Usage()
		}
		tmp0, err367 := (strconv.Atoi(flag.Arg(1)))
		if err367 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		arg368 := flag.Arg(2)
		mbTrans369 := thrift.NewTMemoryBufferLen(len(arg368))
		defer mbTrans369.Close()
		_, err370 := mbTrans369.WriteString(arg368)
		if err370 != nil {
			Usage()
			return
		}
		factory371 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt372 := factory371.GetProtocol(mbTrans369)
		argvalue1 := com.NewIpAddr()
		err373 := argvalue1.Read(jsProt372)
		if err373 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		tmp2, err374 := (strconv.Atoi(flag.Arg(3)))
		if err374 != nil {
			Usage()
			return
		}
		argvalue2 := int32(tmp2)
		value2 := argvalue2
		fmt.Print(client.RegisterModule(value0, value1, value2))
		fmt.Print("\n")
		break
	case "unRegisterModule":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "UnRegisterModule requires 1 args")
			flag.Usage()
		}
		tmp0, err375 := (strconv.Atoi(flag.Arg(1)))
		if err375 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.UnRegisterModule(value0))
		fmt.Print("\n")
		break
	case "heartBeat":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "HeartBeat requires 1 args")
			flag.Usage()
		}
		tmp0, err376 := (strconv.Atoi(flag.Arg(1)))
		if err376 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.HeartBeat(value0))
		fmt.Print("\n")
		break
	case "reportTaskProcess":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "ReportTaskProcess requires 3 args")
			flag.Usage()
		}
		tmp0, err377 := (strconv.Atoi(flag.Arg(1)))
		if err377 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		arg379 := flag.Arg(3)
		mbTrans380 := thrift.NewTMemoryBufferLen(len(arg379))
		defer mbTrans380.Close()
		_, err381 := mbTrans380.WriteString(arg379)
		if err381 != nil {
			Usage()
			return
		}
		factory382 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt383 := factory382.GetProtocol(mbTrans380)
		argvalue2 := com.NewTaskProcessArgs_()
		err384 := argvalue2.Read(jsProt383)
		if err384 != nil {
			Usage()
			return
		}
		value2 := argvalue2
		fmt.Print(client.ReportTaskProcess(value0, value1, value2))
		fmt.Print("\n")
		break
	case "reportResult":
		if flag.NArg()-1 != 4 {
			fmt.Fprintln(os.Stderr, "ReportResult_ requires 4 args")
			flag.Usage()
		}
		tmp0, err385 := (strconv.Atoi(flag.Arg(1)))
		if err385 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		arg388 := flag.Arg(4)
		mbTrans389 := thrift.NewTMemoryBufferLen(len(arg388))
		defer mbTrans389.Close()
		_, err390 := mbTrans389.WriteString(arg388)
		if err390 != nil {
			Usage()
			return
		}
		factory391 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt392 := factory391.GetProtocol(mbTrans389)
		containerStruct3 := com.NewCollectReportResultArgs()
		err393 := containerStruct3.ReadField4(jsProt392)
		if err393 != nil {
			Usage()
			return
		}
		argvalue3 := containerStruct3.ResultList
		value3 := argvalue3
		fmt.Print(client.ReportResult_(value0, value1, value2, value3))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
