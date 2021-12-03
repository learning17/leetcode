/* tcpserv01.c */
#include <sys/socket.h>
#include <strings.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>

void str_echo(int sockfd)
{
    ssize_t        n;
    char           buf[4096];
again:
    while((n = read(sockfd, buf, 4096)) > 0)
        write(sockfd, buf, n);

    if(n < 0 && errno == EINTR)
        goto again;
    else if(n < 0)
    {
        perror("read");
        exit(1);
    }

}

int main(int argc, char **argv)
{
    int                   listenfd, connfd;
    pid_t                 childpid;
    socklen_t             clilen;
    struct sockaddr_in    cliaddr, servaddr;

    if((listenfd = socket(AF_INET, SOCK_STREAM, 0)) < 0)
    {
        perror("socket");
        exit(1);
    }

    bzero(&servaddr, sizeof(servaddr));
    servaddr.sin_family = AF_INET;
    servaddr.sin_addr.s_addr = htonl(INADDR_ANY);
    servaddr.sin_port = htons(9877);

    if(bind(listenfd, (struct sockaddr *)&servaddr, sizeof(servaddr)) < 0)
    {
        perror("bind");
        exit(1);
    }
    if(listen(listenfd, 5) < 0)
    {
        perror("listen");
        exit(1);
    }

    for(;;)
    {
        clilen = sizeof(cliaddr);
        if((connfd = accept(listenfd, (struct sockaddr *)&cliaddr, &clilen)) < 0)
        {
            perror("accept");        
            exit(1);
        }

        if((childpid = fork()) < 0)
        {
            perror("fork");
            exit(1);
        }
        else if(childpid == 0)    /* child process */
        {
            if(close(listenfd) < 0) /* close listening socket */
            {
                perror("child close");
                exit(1);
            }
            str_echo(connfd);    /* process the request */
            exit(0);

        }
        if(close(connfd) < 0) /* parent close connected socket */
        {
            perror("parent close");
            exit(1);
        }
    }
}
