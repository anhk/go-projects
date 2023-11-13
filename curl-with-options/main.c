#include <arpa/inet.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

#include "url.h"

int main(int argc, char **argv)
{
    if (argc != 2) {
        printf("invalid argument\n");
        return -1;
    }

    URL_RESULT_T result;
    if (parse_url(argv[1], &result) != 0) {
        printf("invalid url: %s\n", argv[1]);
        return -1;
    }

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
        .sin_addr.s_addr = inet_addr(result.domain),
        // .sin_port = htons(32351),
        .sin_port = htons(result.port),
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