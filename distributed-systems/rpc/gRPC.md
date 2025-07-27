# gRPC
- https://grpc.io
- https://github.com/grpc
- https://github.com/grpc/grpc

Why gRPC
- gRPC is a modern open source high performance Remote Procedure Call (RPC) framework that can run in any environment.
- It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication.
- It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.

Advantage
- Simple service definition: Define your service using Protocol Buffers
- Start quickly and scale: dev & deply
- Works across languages and platforms
- Bi-directional streaming and integrated

Docs
- https://grpc.io/docs/
- Technical Docs
  - https://github.com/grpc/grpc/tree/master/doc
  - https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
  - https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md
  -
- Concepts
  - https://grpc.io/docs/what-is-grpc/
  - https://grpc.io/docs/what-is-grpc/introduction/
  - https://grpc.io/docs/what-is-grpc/core-concepts/
  - RPC life cycle https://grpc.io/docs/what-is-grpc/core-concepts/#rpc-life-cycle
  - https://github.com/grpc/grpc/blob/master/CONCEPTS.md
  - https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md
- Lnaguages
  - https://grpc.io/docs/languages/
  - https://grpc.io/docs/languages/go/
  - https://grpc.io/docs/languages/go/quickstart/
  - gRPC 完整特性 demo 示例 https://grpc.io/docs/languages/go/basics/
- Guides
  - https://grpc.io/docs/guides/
  - https://grpc.io/docs/guides/auth/

客户端与服务器通信流程
- 服务端
  - 实现 RPC 接口
  - 监听 TCP 端口
  - 注册服务
  - 启动服务
- 客户端
  - 建立 TCP 连接
  - 创建 Client
  - 执行 RPC 调用
  - 得到返回消息

Best Practice
- 部署
  - Scaling gRPC With Kubernetes (Using Go) https://nyadgar.com/posts/scaling-grpc-with-kubernetes-using-go/
- 谷歌开源、高性能RPC框架：gRPC 使用体验 https://mp.weixin.qq.com/s/6XXJfbnIaKzSFtXyDDB72g
- 深入浅出 gRPC https://time.geekbang.org/column/intro/100005601
- gRPC 教程 https://www.liwenzhou.com/posts/Go/gRPC/
- Blog https://grpc.io/blog/
- Awesome gRPC
  - https://github.com/grpc-ecosystem/awesome-grpc
- JetBrains Guide - Go
  - https://github.com/heraldofsolace/go-grpc-demo
  - https://www.jetbrains.com/guide/go/tutorials/grpc_part_one/
  - https://www.jetbrains.com/guide/go/tutorials/grpc_part_two/
  - https://www.jetbrains.com/guide/go/tutorials/grpc_part_three/
- gPRC Stream 最佳实践
  - https://grpc.io/docs/languages/go/basics/
    - Server-side streaming RPC
    - Client-side streaming RPC
    - Bidirectional streaming RPC
  - gRPC Streaming https://eddycjy.com/posts/go/grpc/2018-09-24-stream-client-server/
  - gRPC流式示例 https://www.liwenzhou.com/posts/Go/gRPC/#c-0-5-7
  - gRPC(Go)教程(三)---Stream 推送流 https://www.lixueduan.com/posts/grpc/03-stream/
  - go: 如何正确使用grpc stream https://www.superpig.win/blog/details/tppdvipf


## 关键特性
- 基于 HTTP2
  - 支持更高效的流控制、数据分片、头部压缩等特性
- 跨语言支持
- 使用 Protobuf 作为接口定义语言和消息传输格式，二进制协议
- 多种通信模式
  - Unary RPC
  - Server streaming RPC
  - Client streaming RPC
  - Bidirectional streaming RPC
- 超时和取消
- 强大的工具和生态系统
  - 性能分析工具、错误追踪、调试
- 内置身份验证、错误处理
- 支持服务发现、负载均衡

Name Resolution
- https://grpc.io/docs/guides/custom-name-resolution/
- https://github.com/grpc/grpc/blob/master/doc/naming.md

Service Config
- https://grpc.io/docs/guides/service-config/
- https://github.com/grpc/grpc/blob/master/doc/service_config.md

Load Balancing
  - 问题：由于 gRPC 是基于 HTTP2.0 实现的，导致使用 k8s 部署时负载均衡失效，所有请求都复用一个 http2.0 连接，导致请求没有负载均衡。所以需要引入客户端负载均衡。
  - https://grpc.io/blog/grpc-load-balancing/
  - https://grpc.io/docs/guides/custom-load-balancing/
  - https://github.com/grpc/grpc/blob/master/doc/load-balancing.md
  - gRPC负载均衡（客户端负载均衡）https://www.cnblogs.com/FireworksEasyCool/p/12912839.html
  - gRPC中的名称解析和负载均衡 https://www.liwenzhou.com/posts/Go/name-resolving-and-load-balancing-in-grpc/
  - 基于 gRPC 的服务发现与负载均衡（基础篇）https://pandaychen.github.io/2019/07/11/GRPC-SERVICE-DISCOVERY/


## gRPC-Go
- https://grpc.io/docs/languages/go/
- https://pkg.go.dev/google.golang.org/grpc
- https://github.com/grpc/grpc-go
- https://github.com/grpc/grpc-go/tree/master/Documentation
- https://github.com/grpc/grpc-go/tree/master/examples
- https://github.com/grpc/grpc-go/tree/master/examples/features

### Go gRPC Middleware
- https://github.com/grpc-ecosystem/go-grpc-middleware


## gRPC Web
- https://github.com/grpc/grpc-web


## gRPC Ecosystem
- https://github.com/grpc-ecosystem

### gRPC-Gateway
- https://github.com/grpc-ecosystem/grpc-gateway
- https://grpc-ecosystem.github.io/grpc-gateway/
- https://grpc-ecosystem.github.io/grpc-gateway/docs/operations/health_check/
- https://grpc-ecosystem.github.io/grpc-gateway/docs/operations/tracing/

### grpc_health_probe
- https://github.com/grpc-ecosystem/grpc-health-probe


## gPRC debugging
- https://github.com/grpc/grpc-go/tree/master/examples/features/debugging
- https://grpc.io/blog/a-short-introduction-to-channelz/
- https://github.com/grpc/grpc-experiments/tree/master/gdebug


## gRPC tools
- https://github.com/fullstorydev/grpcurl
- https://github.com/fullstorydev/grpcui
- https://github.com/bloomrpc/bloomrpc
