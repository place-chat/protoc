COMPILER := protoc --go_out=./go --python_out=./python ./proto/

protoc-compile:
	$(COMPILER)/AccountService.proto
	$(COMPILER)/MessageService.proto

