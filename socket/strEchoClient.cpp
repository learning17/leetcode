/* tcpcli01.c */
#include <stdio.h>
#include <strings.h>
#include <string.h>
#include <stdlib.h>
#include <errno.h>
#include <unistd.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <sys/socket.h>

int main(int argc, char **argv)
{
    int                   sockfd;
    struct sockaddr_in    servaddr;
    
    if(argc != 2)
    {
        printf("usage: tcpcli <IPaddress> ");
        exit(0);
    }
    
    if((sockfd = socket(AF_INET, SOCK_STREAM, 0)) < 0)
    {
        perror("socket");
        exit(1);
    }
    
    bzero(&servaddr, sizeof(servaddr));
    servaddr.sin_family = AF_INET;
    servaddr.sin_port = htons(9877);
    if(inet_pton(AF_INET, argv[1], &servaddr.sin_addr) < 0)
    {
        perror("inet_pton");    
        exit(1);
    }
    
    if(connect(sockfd, (struct sockaddr *)&servaddr, sizeof(servaddr)) < 0)
    {
        perror("connect");
        exit(1);
    }
    
    str_cli(stdin, sockfd);    /* do it all */

    exit(0);
}
