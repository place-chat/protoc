GO_DST := "./go"
PY_DST := "./python"

protoc-compile:
	protoc -I=. --go_out=$(GO_DST) --python_out=./$(PY_DST) ./AccountService.proto