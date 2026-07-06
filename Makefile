BINARY_NAME=xtime-conv
VERSION=1.0.0

build:
	go build -o bin/$(BINARY_NAME) ./cmd

clean:
	rm -rf bin/

test:
	go test -count=1 ./...

run:
	go run ./cmd/main.go $(ARGS)

build-win:
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY_NAME)-win-amd64.exe ./cmd
