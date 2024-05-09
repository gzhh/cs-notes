# RPC

## gRPC

gRPC是Google公司基于Protobuf开发的跨语言的开源RPC框架。gRPC基于HTTP/2协议设计，可以基于一个HTTP/2链接提供多个服务。

grpc 一般用来在内部服务之间做通信

grpc gateway 可以对外提供 rest 接口

### gRPC & protobuf

生成 grpc 的 go 代码

```
protoc --go_out=. --go_opt=paths=source_relative \\n

--go-grpc_out=. --go-grpc_opt=paths=source_relative \\n

proto/helloworld.proto
```

例如

`protoc --go_out=. --go-grpc_out=. hello.proto`
