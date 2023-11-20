#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "url.h"

int parse_url(char *raw_url, URL_RESULT_T *result)
{
    memset(result, 0, sizeof(URL_RESULT_T));
    result->port = 0;

    if (strncmp(raw_url, "http://", 7) == 0) {
        strncpy(result->protocol, raw_url, 7);
        raw_url += 7;
    } else if (strncmp(raw_url, "https://", 8) == 0) {
        strncpy(result->protocol, raw_url, 8);
        result->port = 443;
        raw_url += 8;
    }

    char *pos;
    if ((pos = strchr(raw_url, '/')) == NULL) {
        strcpy(result->path, "/");
    } else {
        strcpy(result->path, pos);
        *pos = 0;
    }

    if ((pos = strchr(raw_url, ':')) == NULL) {
        strcpy(result->domain, raw_url);
    } else {
        strncpy(result->domain, raw_url, pos - raw_url);
        result->port = atoi(pos + 1);
    }

    return 0;
}