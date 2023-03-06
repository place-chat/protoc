COMPILER := protoc --go_out=./go --go-grpc_out=./go --python_out=./python --pyi_out=./python ./proto

protoc-compile:
	$(COMPILER)/AccountService.proto
	$(COMPILER)/MessageService.proto

