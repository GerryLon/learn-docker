
all: go-uts clone-uts setns

go-uts:
	export GOOS=linux; \
	go build -o main.out main.go
	

clone-uts:
	gcc -Wall -o clone-uts.out clone-uts.c

setns:
	gcc -Wall -o setns.out setns.c;
	
