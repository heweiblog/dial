package base

/*
#include <sys/socket.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <fcntl.h>
#include <string.h>
#include <stdlib.h>
#include <netinet/ip_icmp.h>

int new_raw_socket(){
		int fd = socket(AF_INET, SOCK_RAW,IPPROTO_ICMP);
		if(fd < 0){
				return -1;
		}
		struct timeval timeout;
		timeout.tv_sec = 0;
		timeout.tv_usec = 500*1000;
		int rtn = setsockopt(fd,SOL_SOCKET,SO_RCVTIMEO,(char *)&timeout,sizeof(struct timeval));
		if(rtn < 0){
				close(fd);
				return -1;
		}
		return fd;
}

int new_udp_block_fd(){
		int fd = socket(AF_INET, SOCK_DGRAM, 0);
		if(fd < 0){
				return -1;
		}
		struct timeval timeout;
		timeout.tv_sec = 0;
		timeout.tv_usec= 500*1000;
		int rtn = setsockopt(fd,SOL_SOCKET,SO_RCVTIMEO,&timeout,sizeof(struct timeval));
		if(rtn < 0){
				close(fd);
				return -1;
		}
		return fd;
}

#define SUB_TIME(x,y)  ((x.tv_sec * 1000*1000 + x.tv_usec) - (y.tv_sec * 1000*1000 + y.tv_usec))

int do_udp_dial(char*ip,unsigned short port){
		char answer[500] = {'\0'};
		char msg[8] = {'\0'};
		char remote_ip[32] = {'\0'};
		int anslen = 0,rtn = 0;
		socklen_t rlen = 0;
		struct timeval t_start;
		struct timeval t_end;

		struct sockaddr_in servaddr;
		memset(&servaddr, 0, sizeof(servaddr));
		servaddr.sin_family = AF_INET;
		servaddr.sin_port = htons(port);
		inet_pton(AF_INET,ip,&servaddr.sin_addr.s_addr);

		int fd = new_udp_block_fd();
		if(fd <= 0){
				return 0;
		}
		int raw_fd = new_raw_socket();
		if(raw_fd <= 0){
				close(fd);
				return 0;
		}

		gettimeofday(&t_start,NULL);
		//for(i = 0 ; i < 3 ; i++){
		rtn = sendto(fd,msg,0,0,(struct sockaddr *)&servaddr,sizeof(struct sockaddr_in));
		if(rtn < 0){
				close(fd);
				close(raw_fd);
				return 0;
				//break;
		}

		//memset(answer,0,sizeof(answer));
		anslen = recvfrom(raw_fd,answer,sizeof(answer),0,NULL,NULL);

		if(anslen > 0 && !strcmp(ip,inet_ntop(AF_INET,answer+12,remote_ip,32)) &&
		(*(uint8_t*)(answer+20) == 3) && (*(uint8_t*)(answer+21) == 3)){
				close(fd);
				close(raw_fd);
				return 0;
				//break;
		}
		//}
		gettimeofday(&t_end,NULL);

		close(fd);
		close(raw_fd);
		return SUB_TIME(t_end,t_start);
}
*/
import "C"

import (
	"unsafe"
)

// udp 拨测 调用c函数
func Udp(ip string, port uint16) int64 {
	if Ping(ip) == 0 {
		return 0
	}

	addr := C.CString(ip)
	defer C.free(unsafe.Pointer(addr))

	return int64(C.do_udp_dial(addr, C.ushort(port)))
}
