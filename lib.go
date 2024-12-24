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
func GetLastError(handle unsafe.Pointer) *C.char {
	h := (*C.ServerHandle)(handle)
	return h.error_message
}

//export CloseServer
func CloseServer(handle unsafe.Pointer) {
	h := (*C.ServerHandle)(handle)
	if h == nil {
		return
	}

	// free all pointer
	C.free(unsafe.Pointer(h.server))
	C.free(unsafe.Pointer(h.error_message))
	C.free(unsafe.Pointer(h))
}

//export GetData
func GetData(handle unsafe.Pointer, input *C.char) *C.char {
	h := (*C.ServerHandle)(handle)
	if h == nil || h.is_valid == 0 {
		return C.CString("invalid handle")
	}

	goInput := C.GoString(input)
	if goInput == "valid-input" {
		return C.CString("data-from-server")
	}
	return C.CString("invalid input")
}

func main() {}
