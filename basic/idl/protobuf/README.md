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

### Code generator
Python
- `protoc --proto_path=your_src_path --python_out=your_dest_path your.proto`

Golang
- `protoc --go_out=paths=source_relative:. your.proto`


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
