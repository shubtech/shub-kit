.PHONY: build test install

build:
	go build -o bin/skit cmd/shub-kit/main.go

test: 
	go test ./...

install: build
	sudo cp bin/skit /usr/local/bin