#安装grpc
```
curl -fsSL https://goo.gl/getgrpc | bash -s -- --with-plugins
或
brew tap grpc/grpc
brew install --with-plugins grpc
```
#更新依赖的go项目
```
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```
#编译命令
```
protoc -I . --go_out=plugins=grpc:. ./hello.proto
```
#帮助查看
```
protoc --help
```
#查看范例
```
protoc --proto_path=IMPORT_PATH --cpp_out=DST_DIR --java_out=DST_DIR --python_out=DST_DIR --go_out=DST_DIR --ruby_out=DST_DIR --javanano_out=DST_DIR --objc_out=DST_DIR --csharp_out=DST_DIR path/to/file.proto
```