.PHONY: server client

server:
	go clean && go build
	./test-server

update:
	go run main.go --update


test:
	gotest -v ./... -count=1