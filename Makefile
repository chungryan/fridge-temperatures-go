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

cover: test
	go tool cover -html=c.out

package:
	aws cloudformation package \
        --template-file cloudformation.yml \
        --output-template-file packaged.yml \
        --s3-bucket ${AWS_BUCKET} \
		--s3-prefix FridgeTemperatureSensors

deploy:
	aws cloudformation deploy \
		--region ap-southeast-2 \
		--template-file packaged.yml \
		--stack-name FridgeTemperatureSensors \
		--capabilities CAPABILITY_IAM

publish: cibuild package deploy

local-invoke: cibuild
	sam local invoke ProcessFunction -t cloudformation.yml -e event.json
