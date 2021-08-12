
.PHONY: proto
proto:
	ls -1 proto/*/*.proto|xargs -L 1 protoc --proto_path=${GOPATH}/src:. --go_out=. --go-grpc_out=paths=source_relative:.
