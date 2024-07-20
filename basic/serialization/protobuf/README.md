# Protobuf
- https://github.com/protocolbuffers
- https://protobuf.dev
- https://protobuf.dev/programming-guides/
- https://protobuf.dev/getting-started/
- https://protobuf.dev/support/version-support/
- Go
  - https://github.com/protocolbuffers/protobuf-go
  - https://protobuf.dev/getting-started/gotutorial/
  - https://protobuf.dev/reference/go/
  - https://protobuf.dev/reference/go/go-generated/
  - https://protobuf.dev/reference/go/faq/
- Python
  - https://protobuf.dev/getting-started/pythontutorial/
  - https://protobuf.dev/reference/python/
  - https://protobuf.dev/reference/python/python-generated/

Buf
- https://github.com/bufbuild
- https://github.com/bufbuild/buf
- https://github.com/bufbuild/protoc-gen-validate
- https://buf.build
- https://buf.build/docs/introduction
- https://buf.build/docs/tutorials/getting-started-with-buf-cli/

接口定义语言，类似于 json，xml
- https://en.wikipedia.org/wiki/Protocol_Buffers
- https://github.com/protocolbuffers/protobuf
- https://protobuf.dev/


## 最佳实践
- Google API Protos https://github.com/googleapis/googleapis/tree/master/google/api
- 高效的数据压缩编码方式 Protobuf https://halfrost.com/protobuf_encode/
- 高效的序列化/反序列化数据方式 Protobuf https://halfrost.com/protobuf_decode/
- Protobuf 编码&避坑指南 https://www.luozhiyun.com/archives/800
- Protobuf 终极教程 https://colobu.com/2019/10/03/protobuf-ultimate-tutorial-in-go/
- 腾讯
  - 数据序列化工具Protobuf编码&避坑指南 https://mp.weixin.qq.com/s/xH9v4Al3B2vPbZIp2yqZpQ
  - Protobuf编码原理及优化技巧探讨 https://mp.weixin.qq.com/s/hAfrPlPD2KBCWxpIuGkQTQ
  - 更快更小！ProtoBuf 入门详解 https://mp.weixin.qq.com/s/zKE3kqkJNQDKby78ddjWBA
- 【Go微服务】一文带你玩转ProtoBuf https://juejin.cn/post/7144948875613339685
- Go API Design With Protocol Buffers and gRPC https://betterprogramming.pub/go-api-design-with-protocol-buffers-and-grpc-991838e4852b
- Supercharge your REST APIs with Protobuf https://medium.com/swlh/supercharge-your-rest-apis-with-protobuf-b38d3d7a28d3

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

proto 文件解释
```
syntax = "proto3";

// 类似于 c++ java 中的包命名，防止不同 namespace 中相同类型的冲突
package proto;

//option go_package = "path;name";
//option go_package = "name";
//path 表示生成的go文件的存放地址，会自动生成目录的。
//name 表示生成的go文件所属的包名
option go_package="./;search";
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
- install
  - go install google.golang.org/protobuf/cmd/protoc-gen-go
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
