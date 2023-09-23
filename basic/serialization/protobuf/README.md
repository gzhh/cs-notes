# Protobuf
- https://protobuf.dev/
- https://en.wikipedia.org/wiki/Protocol_Buffers


## Protocol Buffers
Protocol Buffers are language-neutral, platform-neutral extensible mechanisms for serializing structured data.

### Proto definition
proto3
- https://protobuf.dev/programming-guides/proto3/

person.proto
```
syntax = "proto3"

message Person {
  optional string name = 1;
  optional int32 id = 2;
  optional string email = 3;
}
```


## Protocol Buffer Compiler
The protocol buffer compiler is used to compile .proto files, which contain service and message definitions.

### protoc version and language runtime version
Version Support
- https://protobuf.dev/support/version-support/
- 2022 年之前 protoc 和 language runtime 共用一个版本号，带有大版本例如 3.7.1。
- 现在将 protoc 和 language runtime 的版本发型拆开，让各个语言都能单独发布自己的版本更新。

protoc version
- 不带大版本号，例如 24.3

language runtime version
- 带有大版本号，例如 3.24.3

PS:
1. 一般来说 protoc 和 language runtime 的版本需要有一定的兼容，否则会出问题，兼容的例子如下：
    - protoc version 为 libprotoc 23.3，go 的 package version 为 v1.31.0
    - protoc version 为 libprotoc 3.7.1，python 的 protobuf package version 为 3.19.4

### Install protoc
Mac
- brew
  - `brew install protobuf`
- install package
    ```
    PROTOC_ZIP=protoc-3.14.0-osx-x86_64.zip
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/$PROTOC_ZIP
    sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
    ```

### Code generator
Python
- `protoc --proto_path=your_src_path --python_out=your_dest_path your.proto`

Golang
- `protoc --go_out=paths=source_relative:. your.proto`

Trouble
- Go
  - protoc 编译生成 .pb.go 文件后，可能会出现 Import cycle not allowed，此时需要注意 .proto 文件中不应有相互引用的情况。
- protoc version
  - 不同版本的 protoc 生成的目标语言文件可能不兼容


### Decode
without proto file
- `protoc --decode_raw < pbbin`
- ui tool: https://deprotobuf.com

with proto file
- `protoc --decode=packageName.messageName your.proto < pbbin`
- example `protoc --decode=api.Hello hello.proto < pbbin`
- `protoc --decode=MESSAGE_TYPE --proto_path=your_proto_dir your.proto < pbbin`
- example `protoc --decode=api.Hello --proto_path=. hello.proto < pbbin`


### Buf cli
The buf CLI is a tool for working with Protocol Buffers.
- https://github.com/bufbuild/buf


## HTTP & Protobuf
### HTTP Content-Type
- `application/x-protobuf` preferred choice
- `application/vnd.google.protobuf`

Ref
- https://stackoverflow.com/questions/30505408/what-is-the-correct-protobuf-content-type
- https://groups.google.com/g/protobuf/c/VAoJ-HtgpAI


## 使用参考
李文周的博客 - protocol buffers使用指南
- https://www.liwenzhou.com/posts/Go/protobuf/
