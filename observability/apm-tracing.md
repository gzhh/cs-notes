# APM
- https://en.wikipedia.org/wiki/Application_performance_management

介绍
- APM 原理与框架选型 https://cloud.tencent.com/developer/article/1347919
- APM应用性能监控 https://help.aliyun.com/zh/es/use-cases/best-practices-of-apm


## ELK


## Distributed Systems Tracining 链路追踪
- Dapper, a Large-Scale Distributed Systems Tracing Infrastructure https://static.googleusercontent.com/media/research.google.com/zh-CN//archive/papers/dapper-2010-1.pdf
- 链路追踪 https://icyfenix.cn/distribution/observability/tracing.html
- 什么是链路追踪？分布式系统如何实现链路追踪？https://zhuanlan.zhihu.com/p/344020712
- 全链路分布式跟踪系统与APM https://zhuanlan.zhihu.com/p/40623010
- 基础篇丨链路追踪（Tracing）其实很简单 https://mp.weixin.qq.com/s/00aiWY5bX6RnAKL8UpAZyw
- 分布式链路跟踪中的traceid和spanid代表什么？https://cloud.tencent.com/developer/article/1832719

### OpenTelemetry
- https://opentelemetry.io
- https://github.com/open-telemetry
- https://github.com/open-telemetry/opentelemetry-go
- https://github.com/open-telemetry/opentelemetry-collector
- https://github.com/open-telemetry/opentelemetry-specification
- OpenTelemetry Operator for Kubernetes https://github.com/open-telemetry/opentelemetry-operator
- Collection of extensions for OpenTelemetry-Go. https://github.com/open-telemetry/opentelemetry-go-contrib/

Docs
- OpenTelemetry Specification https://opentelemetry.io/docs/specs/otel/

Demo
- https://github.com/open-telemetry/opentelemetry-demo
- https://opentelemetry.io/docs/demo/
- Docker deployment https://opentelemetry.io/docs/demo/docker-deployment/

Best Practice
- 基於 Open Telemetry 的自助遙測解決方案 https://blog.wildcat.io/2024/01/self-hosted-telemetry-solution-based-on-otel-zh/
- 李文周
  - OpenTelemetry 介绍 https://www.liwenzhou.com/posts/Go/otel/
  - OpenTelemetry Go快速指南 https://www.liwenzhou.com/posts/Go/openTelemetry-go/
  - 基于OTel的HTTP链路追踪 https://www.liwenzhou.com/posts/Go/go-http-otel/
  - gRPC的链路追踪 https://www.liwenzhou.com/posts/Go/go-grpc-otel/
  - GORM配置链路追踪 https://www.liwenzhou.com/posts/Go/gorm-otel/
  - go-redis配置链路追踪 https://www.liwenzhou.com/posts/Go/redis-otel/
  - zap日志库配置链路追踪 https://www.liwenzhou.com/posts/Go/zap-otel/
- OpenTelemetry: A Guide to Observability with Go https://www.lucavall.in/blog/opentelemetry-a-guide-to-observability-with-go
- OTel 数据库查询分析：开发者的性能优化神器 https://mp.weixin.qq.com/s/RJxSx2ICEw2OJrD5q6SnDg
- Practical Tracing for Go Apps with OpenTelemetry (Beginner's Guide) https://betterstack.com/community/guides/observability/opentelemetry-go/

### Jaeger
- https://www.jaegertracing.io
- https://github.com/jaegertracing
- https://github.com/jaegertracing/jaeger
- Jaeger Operator for Kubernetes https://github.com/jaegertracing/jaeger-operator
- Jaeger快速指南 https://www.liwenzhou.com/posts/Go/jaeger/
- Jaeger Deployment https://www.jaegertracing.io/docs/1.18/deployment/

### Zipkin
- https://zipkin.io
- https://github.com/openzipkin/zipkin
- https://github.com/openzipkin/zipkin-go

### SkyWalking
- https://skywalking.apache.org
- https://github.com/apache/skywalking

### Uptrace
- https://uptrace.dev
- https://github.com/uptrace/uptrace
- https://github.com/uptrace/opentelemetry-go-extra

### OpenTracing(archived)
- https://opentracing.io
- https://github.com/opentracing
- https://github.com/opentracing/opentracing-go
