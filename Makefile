protoc := protoc --go_out=./out --go-grpc_out=./out --python_out=./out --pyi_out=./out

build:
	@bash ./build.sh "$(protoc)"