.PHONY: server client
server:
	go clean && go build
	./lol-scylla -addr=:50062 -prod=false