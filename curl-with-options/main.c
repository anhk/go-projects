#include <arpa/inet.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

int main(int argc, char **argv)
{
    int fd = socket(AF_INET, SOCK_STREAM, 0);
    if (fd < 0) {
        perror("socket:");
        return -1;
    }

    unsigned char opt[16] = {
        0x21, // TAG
        8,    // LEN
        1,    2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
    };

    if (setsockopt(fd, IPPROTO_IP, IP_OPTIONS, (void *)opt, 8) < 0) {
        perror("setsockopt:");
        return -1;
    }

    struct sockaddr_in addr = {
        .sin_family = AF_INET,
        // .sin_addr.s_addr = inet_addr("192.168.64.52"),
        .sin_addr.s_addr = inet_addr("10.244.177.21"),
        // .sin_port = htons(32351),
        .sin_port = htons(80),
    };

    if (connect(fd, (struct sockaddr *)&addr, sizeof(addr)) < 0) {
        perror("connect:");
        return -1;
    }

    char data[] = "GET / HTTP/1.1\r\nHost: 192.168.64.52\r\nConnection: close\r\n\r\n";

    if (write(fd, data, sizeof(data)) != sizeof(data)) {
        perror("write:");
        return -1;
    }

    char response[1024];
    memset(response, 0, sizeof(response));

    int ret;
    if ((ret = read(fd, response, sizeof(response))) < 0) {
        perror("read:");
        return -1;
    }

    printf("%s\n", response);

    return 0;
}