protoc := protoc --go_out=./out --go-grpc_out=./out --python_out=./out --pyi_out=./out ./proto

protoc-compile:
	$(protoc)/Account.proto

