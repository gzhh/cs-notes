# Protobuf Go

Lib
  - https://github.com/protocolbuffers/protobuf-go（新）
    - https://pkg.go.dev/google.golang.org/protobuf
  - https://github.com/golang/protobuf（旧）
    - https://pkg.go.dev/github.com/golang/protobuf
  - https://github.com/gogo/protobuf（Third Party）
  - https://github.com/jhump/protoreflect (Third Party)

Sub Lib
- https://pkg.go.dev/google.golang.org/protobuf/encoding
- https://pkg.go.dev/google.golang.org/protobuf/proto

Docs
- Protocol Buffer Basics: Go https://protobuf.dev/getting-started/gotutorial/


### Troubleshooting
protocol buffer namespace conflict
- https://protobuf.dev/reference/go/faq/#namespace-conflict
- Protobuf Name Conflict 分析与解决 https://blog.ivansli.com/2022/08/07/grpc-conflict/

Python
- ImportError: cannot import name '_message' from 'google.protobuf.pyext'
  - 方案一：c 扩展方案 https://github.com/protocolbuffers/protobuf/issues/8820
  - 方案二：纯 python 实现方案 https://stackoverflow.com/questions/50839667/protofile-proto-a-file-with-this-name-is-already-in-the-pool
- protofile.proto: A file with this name is already in the pool.
  - 方案一：修改 proto name 和 message name
  - 方案二：修改 package name
