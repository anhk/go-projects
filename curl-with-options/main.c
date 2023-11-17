#include <arpa/inet.h>
#include <errno.h>
#include <getopt.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

#include "url.h"

static struct option long_options[] = {
    {"ttl", required_argument, NULL, 't'},
    {"option", required_argument, NULL, 'o'},
    {NULL, 0, NULL, 0},
};

int main(int argc, char **argv)
{
    char *url = NULL;
    socklen_t ttl = 0;
    int enable_opt = 1;

    int o;
    int option_index = 0;
    while ((o = getopt_long(argc, argv, "t:o", long_options, &option_index)) >= 0) {
        switch (o) {
        case 't':
            ttl = atoi(optarg);
            break;
        case 'o':
            if (optarg == NULL || strcmp(optarg, "false") == 0) {
                enable_opt = 0;
            }
            break;
        default:
            break;
        }
    }

    if (optind <= argc) {
        url = argv[optind];
    }

    if (url == NULL) {
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
        perror("socket");
        return -1;
    }

    if (ttl > 0) {
        printf("set ttl to %d\n", ttl);
        if (setsockopt(fd, IPPROTO_IP, IP_TTL, &ttl, sizeof(ttl)) < 0) {
            perror("setsockopt ttl");
            return -1;
        }
    }

    if (enable_opt) {
        unsigned char opt[16] = {
            0x21, // TAG
            8,    // LEN
            1,    2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
        };

        printf("add option to ipv4 header\n");
        if (setsockopt(fd, IPPROTO_IP, IP_OPTIONS, (void *)opt, 8) < 0) {
            perror("setsockopt");
            return -1;
        }
    }

    struct sockaddr_in addr = {
        .sin_family = AF_INET,
        .sin_addr.s_addr = inet_addr(result.domain),
        .sin_port = htons(result.port),
    };

    if (connect(fd, (struct sockaddr *)&addr, sizeof(addr)) < 0) {
        printf("connect: [%d] %s\n", errno, strerror(errno));
        return -1;
    }

    char data[] = "GET / HTTP/1.1\r\nHost: 192.168.64.52\r\nConnection: close\r\n\r\n";

    if (write(fd, data, sizeof(data)) != sizeof(data)) {
        perror("write");
        return -1;
    }

    char response[1024];
    memset(response, 0, sizeof(response));

    int ret;
    if ((ret = read(fd, response, sizeof(response))) < 0) {
        perror("read");
        return -1;
    }

    printf("%s\n", response);

    return 0;
}