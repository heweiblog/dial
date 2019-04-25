namespace * rpc.dial.yamutech.com

typedef string ObjectId 

exception Xception 
{
  1: i32 errorCode,
  2: string message
}

enum ModuleType
{
  DIALING = 2
}

enum DialStatus
{
  OK = 0,
  FAIL
}

enum RetCode
{
  FAIL = 0, 
  OK = 1
}

enum DialMethod
{
  DIAL_TCPPORT = 0, 
  DIAL_IMCP,
  DIAL_HTTPGET,
  DIAL_DATABASE,
  DIAL_EXTHTTPGET,
  DIAL_EXTTCPPORT,
  DIAL_EXTHTTPPOST,
  DIAL_HTTPCOMMON,
  DIAL_UDPPORT, 
  DIAL_FTP,
  DIAL_SMTP,
  DIAL_SNMP,
  DIAL_ORACLE
}

enum ModuleState
{
  STARTUP = 0,
  REGISTERED
}

enum SysCommand
{
  RestoreConfig = 0
}

enum DialServerType
{
  XPROXY=0,
  REDIRECT,
  XFORWARD,
  DATACENTER
}

struct HeartBeatState
{
  1:ModuleState mState,
  2:bool serverState,
}

struct IpAddr
{
  1: i32 version,
  2: string addr  
}

struct IpsecAddress
{
  1: IpAddr ip,
  2: i32 mask
}

struct SysIpSec
{
  1: string name, 
  2: IpsecAddress ipsec,
  3: string recordId
}

struct DialOption
{
  1: string destUrl,
  2: string testMethod,
  3: list<i32> expectCode,
  4: string expectMatch,
  5: string contentType,
  6: i32 tag
}

struct HealthPolicyInfo
{
  1: string name,
  2: DialMethod method,
  3: optional i16 port,
  4: i32 freq,
  5: i32 times,
  6: i32 passed,
  7: DialOption option
}

struct DialRecord
{
  1: ObjectId rid,
  2: IpAddr ip,
  3: i32 ttl,
  4: i32 priority,
  5: bool enabled
}

struct DialRecordStatus 
{
  1: ObjectId rid,
  2: DialStatus status,
  3: i64 delay
}

struct DialHealthResult
{
  1: string groupName,
  2: string policyName,
  3: list<DialRecordStatus> statusList,
}

struct DialNginxServer
{
  1: string localURL,
  2: i32 priority
}

struct DialNginxStatus
{
  1: DialNginxServer server,
  2: DialStatus status,
  3: i64 delay
}

struct DialNginxResult
{
  1: string groupName,
  2: string policyName,
  3: list<DialNginxStatus> statusList,
}

struct DialServerStatus
{
  1: ObjectId rid,
  2: IpAddr ip,
  3: DialStatus status,
  4: i64 delay
}

struct DialServerResult
{
  1: DialServerStatus status,
  2: DialServerType typ
}

struct DcInfo
{
  1: string id,
  2: IpAddr ip,
  3: list<string> PolicyList
}

struct DialDcResult
{
  1: string id,
  2: string policy,
  3: DialStatus status,
  4: i64 delay
}

enum SnmpDevType
{
  HOST=0,
  ROUTER,
  H3C,
  HUAWEI,
  CISCO
}

struct SnmpGroupInfo
{
  1: bool enable,
  2: string name,
  3: string community,
  4: string user,
  5: string passwd,
  6: i32 version,
  7: i32 interval,
  8: i32 port,
  9: IpAddr ip,
  10: SnmpDevType type
}

struct InterfaceTraffic
{
  1: i32 index,
  2: i64 inoctets,
  3: i64 outoctets
}

struct IpMac
{
  1: i32 index,
  2: IpAddr ip,
  3: string physaddress
}

struct MacTable
{
  1: string macaddress,
  2: i32 index,
  3: string portname
}

struct InterfaceInfo
{
  1: i32 index,
  2: string descr,
  3: i32 type,
  4: i32 status,
  5: i64 speed,
  6: i32 mtu,
  7: string physaddress
}

struct RouteInfo
{
  1: i32 ifindex,
  2: IpAddr destination,
  3: IpAddr gateway,
  4: IpAddr genmask,
  5: i32 type,
  6: i32 proto
}

struct SysInfo
{
  1: i32 load,
  2: i32 usercpu, 
  3: i32 syscpu, 
  4: i32 idlecpu, 
  5: i32 totalmem, 
  6: i32 freemem, 
  7: i32 buffer, 
  8: i32 cache, 
  9: i32 availmem 
}

struct ProcessInfo
{
  1: string name,
  2: bool existflag, 
  3: i32 pid,
  4: i32 cputime, 
  5: i32 usedmem
}



enum TaskEvent
{
  IDLE =0,
  RUNNING = 1,
  SUSPEND = 2,
  FINISHED =3,
  CANCELED =4
}

struct TaskProcessArgs
{
  1: TaskEvent event,
  2: string batchno,
  3: double percent,
  4: double dialLocalRate,
  5: i32 dialAvgDelay,
  6: double detectLocalRate,
  7: double detectAvailRate,
  8: i32 detectAvgDelay,
  9: i32 totalAvgDelay,
  10: bool closed
}

enum DomainType
{
	A = 1,
	NS = 2,
	CNAME = 5,
	SOA = 6,
	PTR = 12,
	HINFO = 13,
	MX = 15,
	TXT = 16,
	AAAA = 28,
	SRV = 33,
	A6 = 38,
	ANY = 255	
}

struct DomainRecord
{
  1: string dname, 
  2: DomainType dtype
} 

enum DialPlatformMethod
{
  Dig = 0,
  DigAndPing = 1,
  DigAndHttp = 2,
  DigAndWeb = 3,
  DigAndVideo = 4,
  FocusDomain = 5,
  RefreshCache = 6,
  DomainSchedul = 7
}

struct VideoResult
{
  1: string  url,
  2: bool    available,
  3: i32  speed
}

struct IpResult
{
  1: IpAddr  ip,
  2: bool    local,
  3: i32  delay,
  4: bool    available,
  5: list<VideoResult> videoResults,
  6: i32 downloadspeed
}

enum FocusDomainResultStatus
{
  noerror = 0,
  formerr = 1,
  servfail = 2,
  nxdomain = 3,
  notimpl = 4,  
  refused = 5,
  others = 6
}

struct FocusDomainResultItem
{
  1: i32 priority,
  2: string value
}

struct FocusDomainResult
{
  1: FocusDomainResultStatus status,
  2: i32 delay,
  3: list<FocusDomainResultItem> results
}

struct DomainResult
{
  1: string     dname,
  2: DomainType dtype,
  3: bool       available,
  4: list<IpResult> results,
  5: bool       local,
  6: i32     delay,
  7: FocusDomainResult   fdr
}


struct IpSec
{  
  1: i32 version,
  2: IpAddr ip,
  3: i32 mask,
  4: string carrier,
  5: bool local
}

struct DomainTarget
{
  1: string targetid,
  2: string taskid,
  3: string batchno,  
  4: bool       available,  
  5: bool       local,
  6: i32     delay,
  7: list<IpResult> results,
  8: i32     avgIpDelay,
  9: i32     avgVideoSpeed,
  10: i32     totalDelay,
  11: i64   updated
}

struct AnalysisResult
{
  1: string     dname,
  2: DomainType dtype,   
  3: DomainTarget  home,
  4: DomainTarget  suggest,
  5: list<DomainTarget> totals
}

service Agent
{
  RetCode         registerModule(1: ModuleType typ) throws(1: Xception ex),
  RetCode         updateHealthStatus(1: list<DialHealthResult> results) throws(1: Xception ex),
  RetCode         updateServerStatus(1: list<DialServerResult> results) throws(1: Xception ex),
  RetCode         updateDcStatus(1: list<DialDcResult> results) throws(1: Xception ex),
  RetCode         updateNginxStatus(1: list<DialNginxResult> results) throws(1: Xception ex),
  RetCode         updateSysInfo(1: string snmp,2: SysInfo sysinfo) throws(1: Xception ex),
  RetCode         updateInterfaceInfo(1: string snmp,2: list<InterfaceInfo> interfaces) throws(1: Xception ex),
  RetCode         updateInterfaceTraffic(1: string snmp,2: list<InterfaceTraffic> traffic) throws(1: Xception ex),
  RetCode         updateInterfaceIpMac(1: string snmp,2: list<IpMac> ipmac) throws(1: Xception ex),
  RetCode         updateRouteInfo(1: string snmp,2: list<RouteInfo> routeinfo) throws(1: Xception ex),
  RetCode         updateProcessInfo(1: string snmp,2: ProcessInfo processinfo) throws(1: Xception ex),
  RetCode         updateIpSecOnlineIp(1: string ipsecid,2: list<IpAddr> iplist) throws(1: Xception ex),
  RetCode         updateMacTable(1: string snmp,2: list<MacTable> mactable) throws(1: Xception ex)

  RetCode         registerDialModule(1: i32 moduleId,2:IpAddr ip,3:i32 port) throws(1: Xception ex),
  RetCode         unRegisterDialModule(1: i32 moduleId) throws(1: Xception ex),
  RetCode         heartBeat(1: i32 moduleId) throws(1: Xception ex)
}


service Dial
{
  RetCode         systemCommand(1: SysCommand cmdType) throws(1: Xception ex),
  RetCode         addHealthGroup(1: string groupName,2: string policyName) throws(1: Xception ex),
  RetCode         delHealthGroup(1: string groupName,2: string policyName) throws(1: Xception ex),
  RetCode         addHealthRecord(1: string groupName,2:list<DialRecord> records) throws(1: Xception ex),
  RetCode         delHealthRecord(1: string groupName,2:list<DialRecord> records) throws(1: Xception ex),
  RetCode         addHealthPolicy(1: HealthPolicyInfo policy) throws(1: Xception ex),
  RetCode         modHealthPolicy(1: HealthPolicyInfo policy) throws(1: Xception ex),
  RetCode         delHealthPolicy(1: HealthPolicyInfo policy) throws(1: Xception ex),
  RetCode         addDialServer(1:ObjectId rid, 2: IpAddr ip,3: DialServerType typ) throws(1: Xception ex),
  RetCode         delDialServer(1:ObjectId rid) throws(1: Xception ex),
  RetCode         addNginxGroup(1: string groupName,2: string policyName) throws(1:Xception ex),
  RetCode         delNginxGroup(1: string groupName,2: string policyName) throws(1:Xception ex),
  RetCode         addNginxServer(1: string groupName,2: list<DialNginxServer> servers) throws(1:Xception ex),
  RetCode         delNginxServer(1: string groupName,2: list<DialNginxServer> servers) throws(1:Xception ex),
  HeartBeatState  heartBeat() throws(1: Xception ex),
  RetCode   	  setServerState(1:bool enable) throws(1: Xception ex)
  RetCode         addSnmpGroupInfo(1: SnmpGroupInfo snmp) throws(1: Xception ex),
  RetCode         delSnmpGroupInfo(1: string snmp) throws(1: Xception ex)
  RetCode         addSnmpProcessInfo(1: string snmp,2: string processname) throws(1: Xception ex),
  RetCode         delSnmpProcessInfo(1: string snmp,2: string processname) throws(1: Xception ex),
  RetCode         addIpSec(1: SysIpSec ipsec,2: i32 interval) throws(1: Xception ex),
  RetCode         delIpSec(1: string ipsecid) throws(1: Xception ex),
  RetCode         addDcInfo(1: DcInfo dc) throws(1: Xception ex),
  RetCode         delDcInfo(1: string id) throws(1: Xception ex)

  RetCode         heartBeatDial()throws(1: Xception ex),
  RetCode         resetModule()throws(1: Xception ex),
  RetCode		  addDialIpSec(1:list<IpSec> ipSecList) throws(1: Xception ex),
  RetCode		  removeDialIpSec(1:list<IpSec> ipSecList) throws(1: Xception ex),
  RetCode		  clearDialPlatformIpSec() throws(1: Xception ex),
  RetCode		  addDialDomain(1:string groupId,2:list<DomainRecord> DomainList) throws(1: Xception ex),
  RetCode		  removeDialDomain(1:string groupId,2:list<DomainRecord> DomainList) throws(1: Xception ex),
  RetCode		  clearDialDomain(1:string groupId) throws(1: Xception ex),
  RetCode		  addDialTask(1:string taskId,2:DialMethod method,3:list<IpAddr> targetList,4:IpAddr sourceip,5:i32 interval,6:string domainGroupId) throws(1: Xception ex),
  RetCode		  removeDialTask(1:string taskId) throws(1: Xception ex)
}

service Collect
{
  RetCode         registerModule(1: i32 moduleId,2:IpAddr ip,3:i32 port) throws(1: Xception ex),
  RetCode         unRegisterModule(1: i32 moduleId) throws(1: Xception ex),
  RetCode         heartBeat(1: i32 moduleId) throws(1: Xception ex),
  RetCode		  reportTaskProcess(1:i32 moduleId,2: string taskId,3:TaskProcessArgs arg) throws(1: Xception ex),
  RetCode		  reportResult(1:i32 moduleId,2:string taskId,3:string batchno,4:list<DomainResult> resultList) throws(1: Xception ex)
}

