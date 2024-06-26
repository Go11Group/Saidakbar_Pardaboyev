CURRENT_DIR := $(shell pwd)

run-server:
	@go run server/main.go

run-client:
	@go run client/main.go

gen-proto:
	@scripts/gen-proto.sh $(CURRENT_DIR)

tidy:
	@go mod tidy
	@go mod vendor