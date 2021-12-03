#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>
#include <string.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/time.h>
#include <netinet/in.h>
#include <arpa/inet.h>

#define MYPORT 8100 //连接时使用的端口
#define MAXCLINE 5 //连接队列中的个数
#define BUF_SIZE 200

int fd[MAXCLINE]; //连接的fd
int conn_amount; //当前的连接数
int main(void)
{
    int sock_fd,new_fd; //监听套接字 连接套接字
    struct sockaddr_in server_addr; // 服务器的地址信息
    struct sockaddr_in client_addr; //客户端的地址信息
    socklen_t sin_size;
    int yes = 1;
    char buf[BUF_SIZE];
    int ret;
    int i;

    if((sock_fd = socket(AF_INET,SOCK_STREAM,0))==-1) {
        exit(1);
    }
    //设置套接口的选项 SO_REUSEADDR 允许在同一个端口启动服务器的多个实例
    if(setsockopt(sock_fd, SOL_SOCKET, SO_REUSEADDR, &yes, sizeof(int))==-1) {
        exit(1);
    }
    server_addr.sin_family = AF_INET; //主机字节序
    server_addr.sin_port = htons(MYPORT);
    server_addr.sin_addr.s_addr = INADDR_ANY;//通配IP
    memset(server_addr.sin_zero,'\0',sizeof(server_addr.sin_zero));
    if(bind(sock_fd,(struct sockaddr *)&server_addr,sizeof(server_addr)) == -1) {
        exit(1);
    }
    if(listen(sock_fd,MAXCLINE)==-1) {
        exit(1);
    }
    fd_set fdsr; //文件描述符集的定义
    int maxsock;
    struct timeval tv;
    conn_amount =0;
    sin_size = sizeof(client_addr);
    maxsock = sock_fd;
    while(1) {
        FD_ZERO(&fdsr);        // 所有位清空
        FD_SET(sock_fd,&fdsr); // 将 sock_fd 对应 bit 位设为 1
        tv.tv_sec = 30; // 超时时间，如果不设，阻塞模式一直等到有事件发生
        tv.tv_usec =0;

        //将所有的连接全部加到这个这个集合中，可以监测客户端是否有数据到来
        for(i = 0; i < MAXCLINE; i++) {
            if(fd[i]!= 0) {
                FD_SET(fd[i],&fdsr);
            }
        }
        //如果文件描述符中有连接请求 会做相应的处理，实现I/O的复用 多用户的连接通讯
        ret = select(maxsock +1,&fdsr,NULL,NULL,&tv);
        if(ret <0) { //没有找到有效的连接 失败
            break;
        } else if(ret ==0) {// 指定的时间到，
            continue;
        }
        //下面这个循环是非常必要的，因为你并不知道是哪个连接发过来的数据，所以只有一个一个去找。
        for(i = 0; i < conn_amount; i++) {
            if(FD_ISSET(fd[i],&fdsr)) {
                ret = recv(fd[i],buf,sizeof(buf),0);
                if(ret <=0 ) {//客户端连接关闭，清除文件描述符集中的相应的位
                    close(fd[i]);
                    FD_CLR(fd[i],&fdsr);
                    fd[i]=0;
                    conn_amount--;
                } else {  //否则有相应的数据发送过来 ，进行相应的处理
                    if(ret < BUF_SIZE) {
                        memset(&buf[ret],'\0',1);
                    }
                }
            }
        }
        if(FD_ISSET(sock_fd,&fdsr)) {
            new_fd = accept(sock_fd,(struct sockaddr *)&client_addr,&sin_size);
            if(new_fd <=0) {
                continue;
            }
            // 新连接添加到集合中
            if(conn_amount < MAXCLINE) {
                for(i = 0; i < MAXCLINE; i++) {
                    if(fd[i]==0) {
                        fd[i] = new_fd;
                        break;
                    }
                }
                conn_amount++;
                if(new_fd > maxsock) {
                    maxsock = new_fd;
                }
            } else {
                send(new_fd,"bye",4,0);
                close(new_fd);
                continue;
            }
        }
    }
    for(i=0; i< MAXCLINE; i++)
    {
        if(fd[i]!=0) {
            close(fd[i]);
        }
    }  
    exit(0);
}
