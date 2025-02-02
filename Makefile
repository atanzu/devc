PREFIX?=/usr/local
_INSTDIR=${DESTDIR}${PREFIX}
BINDIR?=${_INSTDIR}/bin
MANDIR?=${_INSTDIR}/share/man
APP=devc
TARGET=""

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	TARGET=linux
endif
ifeq ($(UNAME_S),Darwin)
	TARGET=darwin
endif

.PHONY: all
all: build

.PHONY: build
## build: Build the application
build:
	@echo "Building..."
	@env GOOS=${TARGET} GOARCH=amd64 go build -mod vendor -o build/${APP}-${TARGET}-amd64 main.go

.PHONY: check
## check: Check that the build is working
check:
	@./${APP}

.PHONY: install
## install: Install the application
install:
	@echo "Installing..."
	@cp build/${APP}-${TARGET}-amd64 build/${APP}
	@mkdir -p ${BINDIR}
	@install -t ${BINDIR}/ build/${APP}

.PHONY: uninstall
## uninstall: Uninstall the application
uninstall:
	@echo "Uninstalling..."
	@rm -rf ${BINDIR}/${APP}

.PHONY: run
## run: Runs go run main.go
run:
	go run -race main.go

.PHONY: clean
## clean: Cleans the binary
clean:
	@echo "Cleaning..."
	@rm -rf ${APP}

.PHONY: setup
## setup: Setup go modules
setup:
	@-go mod init
	@go mod tidy
	@go mod vendor

.PHONY: lint
## lint: Runs golint linter on the project
lint:
	@golint .

.PHONY: format
## format: Runs goimports on the project
format:
	@goimports -l -w .

.PHONY: test
## test: Runs go test
test:
	@go test ./...

.PHONY: help
## help: Prints this help message
help:
	@echo -e "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
