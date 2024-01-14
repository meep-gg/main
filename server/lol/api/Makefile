.PHONY: server client
server:
	go clean && go build
	./lol-api -addr=:50061 -prod=false -redis=:6379