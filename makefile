TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=do-2021.fr
NAMESPACE=polycode
NAME=polycode
BINARY=terraform-provider-${NAME}
VERSION=v0.3.5-rc9
OS_ARCH=linux_amd64

default: install

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

lint:
	golangci-lint run

format:
	test -z $(gofmt -l .)

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

ci: build lint format test

bump:
	sed -i -E "/^VERSION=/c\VERSION=$(BUILD_VERSION)" makefile