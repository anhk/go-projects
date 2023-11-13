#pragma once

#define MAX_PROTOCOL_LEN 32
#define INET_DOMAINSTRLEN 128
#define MAX_PATH_FILE_LEN 256
#define MAX_IP_STR_LEN 32

typedef struct {
    char domain[INET_DOMAINSTRLEN];  // 域名或IP
    unsigned short port;             // 端口
    char path[MAX_PATH_FILE_LEN];    // 文件路径
    char protocol[MAX_PROTOCOL_LEN]; // 协议
} URL_RESULT_T;

int parse_url(char *raw_url, URL_RESULT_T *result);