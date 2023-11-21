IMAGE_TAG ?= latest

.PHONY: clean test build

BIN = dqtools

all: clean go-build

clean:
ifeq ($(BIN), $(wildcard $(BIN)))
	rm $(BIN)
endif

local-test: go-fmt go-lint go-test go-build
	@echo 'Run local-test success!'

go-fmt:
	goimports -l -w -local "dqtools/" .

go-lint:
	golangci-lint run

go-test:  # -race
	go test -v -cover -race dqtools/...

go-test-with-musl:  # -race
	go test -tags 'musl' -v -cover -race dqtools/...

go-build:
	go build -ldflags="-s -w" -gcflags='-l -l' -o dqtools main.go

docker:
	docker build --tag dqworkshop .



