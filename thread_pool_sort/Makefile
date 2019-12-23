export GOPATH=$(pwd)
export GOBIN=$(pwd)/bin
go-install:
	mkdir -p results && GOPATH=$(shell pwd) GOBIN=$(shell pwd)/bin go install src/main.go
run:
	./bin/main
