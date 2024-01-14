# Root directory (assuming the Makefile is in the root)
ROOT_DIR := .

# Pattern of the generated files
FILE_PATTERN := *.pb.go

clean:
	@echo "Cleaning proto files in the project"
	@find $(ROOT_DIR) -name "$(FILE_PATTERN)" -exec rm {} \;
	@echo "All generated proto files have been deleted."


# Root directory (assuming the Makefile is in the root)
ROOT_DIR := .

# Output directory for generated Go files
GO_OUT_DIR := .

# Go module name
GO_MODULE_NAME := meep.gg/protos

build:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	@export PATH="$PATH:$(go env GOPATH)/bin"
	@make clean
	@for proto in $$(find $(ROOT_DIR) -name '*.proto'); do \
		protoc -I=$(ROOT_DIR) \
		       --go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative \
		       --go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=module=$(GO_MODULE_NAME) \
		       $$proto; \
		echo "Generated $$proto"; \
	done
	@echo "All proto files have been compiled."