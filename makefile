PREFIX    :=/usr/local
# export GOPATH := $(shell pwd)
# export PATH := $(GOPATH)/bin:$(PATH)

build: LDFLAGS   += $(shell GOPATH=${GOPATH} src/build/ldflags.sh)
build:
	@echo "--> Building..."
	@mkdir -p bin/
	go build -v -o bin/xenon    --ldflags '$(LDFLAGS)' src/xenon/xenon.go
	go build -v -o bin/xenoncli --ldflags '$(LDFLAGS)' src/cli/cli.go
	@chmod 755 bin/*

clean:
	@echo "--> Cleaning..."
	@mkdir -p bin/
	@go clean
	@rm -f bin/*
	@rm -f coverage*

install:
	@echo "--> Installing..."
	@install bin/xenon bin/xenonctl $(PREFIX)/sbin/

fmt:
	go fmt ./...

test:
	@echo "--> Testing..."
	@$(MAKE) testcommon
	@$(MAKE) testlog
	@$(MAKE) testrpc
	@$(MAKE) testconfig
	@$(MAKE) testmysql
	@$(MAKE) testmysqld
	@$(MAKE) testserver
	@$(MAKE) testraft
	@$(MAKE) testcli
	@$(MAKE) testctl

testcommon:
	go test -v github.com/radondb/xenon/src/xbase/common
testlog:
	go test -v github.com/radondb/xenon/src/xbase/xlog
testrpc:
	go test -v github.com/radondb/xenon/src/xbase/xrpc
testconfig:
	go test -v github.com/radondb/xenon/src/config
testmysql:
	go test -v github.com/radondb/xenon/src/mysql
testmysqld:
	go test -v github.com/radondb/xenon/src/mysqld
testserver:
	go test -v github.com/radondb/xenon/src/server
testraft:
	go test -v github.com/radondb/xenon/src/raft
testcli:
	go test -v github.com/radondb/xenon/src/cli/cmd
testctl:
	go test -v github.com/radondb/xenon/src/ctl/v1

COVPKGS = xbase/common\
		  xbase/xlog\
		  xbase/xrpc\
		  config\
		  mysql\
		  mysqld\
		  raft\
		  server\
		  ctl/v1/
vet:
	go vet $(COVPKGS)

coverage:
	go build -v -o bin/gotestcover \
	src/vendor/github.com/pierrre/gotestcover/*.go;
	bin/gotestcover -coverprofile=coverage.out -v $(COVPKGS)
	go tool cover -html=coverage.out
