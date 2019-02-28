.PHONY: build, clean

build:
	if [ ! -f .env ]; then \
		cp .env.example .env; \
	fi
	go build -o ./api;

clean:
	rm -rf ./api

docker-build:
	docker build . -t arya