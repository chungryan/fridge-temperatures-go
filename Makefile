.PHONY: build

BUILD_CMD = "go build -v -o"
CI_BUILD_FLAGS = "GOOS=linux GOARCH=amd64"
BINARIES = processFunction

deps:
	go get -v -t -d ./...

clean:
	rm -rf ./build

build: clean deps
	for bin in ${BINARIES}; do eval "${BUILD_CMD} build/$$bin ./$$bin/main.go"; done

cibuild: clean deps
	for bin in ${BINARIES}; do eval "${CI_BUILD_FLAGS} ${BUILD_CMD} build/$$bin ./$$bin/main.go"; done

test:
	go test -coverprofile c.out -v ./...
