#include <stdio.h>
#include <stdlib.h>

// golang lib
#include "libhandleplay.h"

int main() {
    // init handle
    ServerHandle *handle = (ServerHandle*) InitServer("https://example.com");
    if (handle == NULL) {
        printf("Failed to initialize server.\n");
        return 1;
    }
    printf("Server initialized successfully.\n");

    // get data
    const char *result = GetData(handle, "err-input");
    if (result == NULL) {
        const char *err = GetLastError(handle);
        printf("Failed to get data: %s\n", err ? err : "unknown error");
    } else {
        printf("Result: %s\n", result);
    }

    // close
    CloseServer(handle);

    return 0;
}