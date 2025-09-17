.PHONY: build run docker

build:
	go build -v -o bin/sysinfo ./...

run:
	go run ./...

docker:
	docker build -t sysinfo:latest .
