BINARY_NAME := xtime-conv
VERSION     := 0.1.0
LDFLAGS     := -s -w \
	-X main.version=$(VERSION) \
	-X main.commit=$(shell git rev-parse --short HEAD 2>/dev/null || echo none) \
	-X main.date=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)

.PHONY: build install run test clean build-win help

build: ## Собрать бинарник в bin/
	go build -ldflags "$(LDFLAGS)" -o bin/$(BINARY_NAME) ./cmd/xtime-conv

install: ## Установить в $GOPATH/bin или $GOBIN
	go install -ldflags "$(LDFLAGS)" ./cmd/xtime-conv

run: ## Запуск: make run ARGS="-t 1710000000"
	go run ./cmd/xtime-conv $(ARGS)

test: ## Запустить тесты
	go test -count=1 ./...

clean: ## Удалить bin/
	rm -rf bin/

build-win: ## Кросс-компиляция для Windows
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o bin/$(BINARY_NAME)-win-amd64.exe ./cmd/xtime-conv

help: ## Показать доступные команды
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'
