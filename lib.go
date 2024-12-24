package main

/*
#include <stdlib.h>
#include <string.h>

typedef struct {
    char *server;
    int is_valid;
    char *error_message;
} ServerHandle;
*/
import "C"
import (
	"unsafe"
)

//export InitServer
func InitServer(server *C.char) *C.ServerHandle {
	goServer := C.GoString(server)

	// allocate c struct pointer in c
	handle := (*C.ServerHandle)(C.malloc(C.size_t(unsafe.Sizeof(C.ServerHandle{}))))
	if handle == nil {
		return nil
	}

	handle.server = C.CString(goServer)
	handle.is_valid = 0
	handle.error_message = nil

	if goServer == "https://example.com" {
		handle.is_valid = 1
	} else {
		handle.error_message = C.CString("server check failed")
	}
	return handle
}

//export GetLastError
func GetLastError(handle *C.ServerHandle) *C.char {
	return handle.error_message
}

//export CloseServer
func CloseServer(handle *C.ServerHandle) {
	if handle == nil {
		return
	}

	// free all pointer
	C.free(unsafe.Pointer(handle.server))
	C.free(unsafe.Pointer(handle.error_message))
	C.free(unsafe.Pointer(handle))
}

//export GetData
func GetData(handle *C.ServerHandle, input *C.char) *C.char {
	if handle == nil || handle.is_valid == 0 {
		return C.CString("invalid handle")
	}

	goInput := C.GoString(input)
	if goInput != "valid-input" {
		errorMessage := "invalid input"
		handle.error_message = C.CString(errorMessage)
		return nil
	}

	return C.CString("data-from-server")
}

func main() {}
