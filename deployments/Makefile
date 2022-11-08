TARGET_SERVER = server
GO = go

# PID file will keep the process id of the server
PID1 := /tmp/.$(TARGET_SERVER).pid

GOFILES := $(PWD)/..

# Make is verbose in Linux. Make it silent.
#MAKEFLAGS += --silent

.PHONY: help
all: help
help: Makefile
	@echo
	@echo "Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo


all: install compile

## start: Run development mode.
start: start-auth
## stop: Stop development mode.
stop: stop-server

compile: build-server

start-server:
	@echo "  >  $(TARGET_SERVER) is available at localhost:$(SERVER_PORT)"
	@-$(GOFILES)/bin/server 2>&1 & echo $$! > $(PID1)
	@cat $(PID1) | sed "/^/s/^/  \>  PID: /"

stop-server:
	@-touch $(PID1)
	@-kill `cat $(PID1)` 2> /dev/null || true
	@-rm $(PID1)

build-server:
	@echo "  >  Building binary for $(TARGET_SERVER)..."
	$(GO) build $(FLAGS) -o $(GOFILES)/bin/server $(GOFILES)/cmd/server

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	$(GO) get -t

run-tests:
	@echo "  > Running tests..."
	$(GO) generate $(GOFILES)/...
	$(GO) test `go list all | grep "mb_api" | grep -v "mocks"` -coverprofile=coverage.out.tmp -cover $(GOFILES)/...
	cat coverage.out.tmp | grep -v ".*_easyjson.go.*" > ./coverage.out
	$(GO) tool cover -func=./coverage.out

run-linters:
	@echo "  > Running linters..."
	golangci-lint run $(GOFILES)/...

install:
	@echo "  >  Installing dependencies..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.35.0
	go install github.com/golang/mock/mockgen

