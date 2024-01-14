.PHONY: server client
server:
	go clean && go build
	./riot-api -addr=:50051 -redis=:6379 -prod=false