package log

import (
	"log"
	"os"
)

const (
	//const LogFile = "/var/log/dial/dial.log"
	//const ConfLog = "/var/log/dial/config.log"
	WorkLog = "./dial.log"
	ConfLog = "./conf.log"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Debug   *log.Logger
	Cfglog  *log.Logger
)

func init() {
	f, err := os.OpenFile(WorkLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}
	cf, err := os.OpenFile(ConfLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	Info = log.New(f, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(f, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(f, "Error:", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(f, "Debug:", log.Ldate|log.Ltime|log.Lshortfile)
	Cfglog = log.New(cf, "Config:", log.Ldate|log.Ltime|log.Lshortfile)
}
