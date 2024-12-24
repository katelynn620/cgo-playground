GO_LIBNAME := handleplay
OUT := handle-play

all: go-static c-static

go-static:
	CGO_ENABLED=1 go build -o lib$(GO_LIBNAME).a -buildmode=c-archive -ldflags '-w -s -linkmode external -extldflags "-static"' lib.go

c-static:
	gcc -o $(OUT)-static.out main.c -L. -Wl,-Bstatic -l$(GO_LIBNAME) -Wl,-Bdynamic -lpthread

.PHONY:clean
clean:
	rm -f *.so *.a *.out