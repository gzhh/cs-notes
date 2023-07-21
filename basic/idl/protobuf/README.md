## Protobuf
- https://en.wikipedia.org/wiki/Protocol_Buffers

### Code generator
Python
- `protoc --proto_path=your_src_path --python_out=your_dest_path your.proto`

Golang
- `protoc --go_out=paths=source_relative:. your.proto`
