
all: build-readlink build-memory-limit

build-readlink:
	gcc -Wall -o readlink.out readlink.c

build-memory-limit:
	go build -o mem.out ./memory-limit.go
