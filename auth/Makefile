MAIN_VERSION:=$(shell git describe --always --abbrev=0 --tags || echo "1.0")
VERSION:=${MAIN_VERSION}
PACKAGES:=$(shell go list ./... | sed -n '1!p' | grep -v /vendor/)
BINARY:=arya_auth

default: run

.PHONY: build, clean, test, docker-build, run

test:
	go test -p=1 -cover -covermode=count ${PACKAGES}

cover:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES), \
	echo ${pkg}; \
	go test -coverprofile=coverage.out -covermode=count ${pkg}; \
	tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out

clean:
	rm -rf coverage.out coverage-all.out ${BINARY}

build: clean test cover build
	if [ ! -f .env ]; then \
		cp .env.example .env; \
	fi
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -a -o ${BINARY} cmd/api/main.go

run:
	./${BINARY}

docker-build:
	docker build . -t arya