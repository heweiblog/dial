
#include <sys/socket.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <fcntl.h>
#include <string.h>
#include <stdio.h>
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

int do_udp_dial(char*ip,unsigned short port){
		char answer[500] = {'\0'};
		char msg[8] = {'\0'};
		char remote_ip[32] = {'\0'};
		int anslen = 0,i = 0,res = 0,rtn = 0;
		socklen_t rlen = 0;

		struct sockaddr_in servaddr;
		memset(&servaddr, 0, sizeof(servaddr));
		servaddr.sin_family = AF_INET;
		servaddr.sin_port = htons(port);
		inet_pton(AF_INET,ip,&servaddr.sin_addr.s_addr);

		int fd = new_udp_block_fd();
		int raw_fd = new_raw_socket();
		if(fd <= 0 || raw_fd <= 0){
				return -1;
		}

		for(i = 0 ; i < 3 ; i++){
				rtn = sendto(fd,msg,0,0,(struct sockaddr *)&servaddr,sizeof(struct sockaddr_in));
				if(rtn < 0){
						res = -1;
						break;
				}

				memset(answer,0,sizeof(answer));
				anslen = recvfrom(raw_fd,answer,sizeof(answer),0,NULL,NULL);
				printf("%d-%d-%s\n",rtn,anslen,inet_ntop(AF_INET,answer+12,remote_ip,32));

				if(anslen > 0 && !strcmp(ip,inet_ntop(AF_INET,answer+12,remote_ip,32)) &&
				(*(uint8_t*)(answer+20) == 3) && (*(uint8_t*)(answer+21) == 3)){
						res = -1;
						break;
				}
		}

		close(fd);
		close(raw_fd);
		return res;
}

int main(){
	int res = do_udp_dial("192.168.5.30",535);
	printf("%d\n",res);
	res = do_udp_dial("192.168.6.54",53);
	printf("%d\n",res);
	res = do_udp_dial("192.168.6.190",688);
	printf("%d\n",res);
	res = do_udp_dial("192.168.6.195",688);
	printf("%d\n",res);
}
