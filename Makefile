protoc := protoc --go_out=./go --go-grpc_out=./go --python_out=./python --pyi_out=./python ./proto

protoc-compile:
	$(protoc)/Account.proto

