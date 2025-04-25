# 微服务
- SOA https://en.wikipedia.org/wiki/Service-oriented_architecture
- Microservices https://en.wikipedia.org/wiki/Microservices
- SOA 和微服务有什么区别？https://aws.amazon.com/cn/compare/the-difference-between-soa-microservices/
- What are microservices? https://microservices.io

介绍
- Service Mesh, Service Discovery and API Gateways https://www.express-gateway.io/service-mesh-service-discovery-and-api-gateways/


## Best Practice
- 《微服务架构设计模式》的一点总结 https://mp.weixin.qq.com/s/TLZR252J7EHcR8h_vEsXrA
- 单体架构比微服务架构更落后吗？https://mp.weixin.qq.com/s/HvMlNYrJFQkqewD9Bwn8_w
- 学习微服务首先要了解为什么使用微服务 https://www.imooc.com/article/22737
- 微服务接口设计原则 https://mp.weixin.qq.com/s/3UNL1EZfgfGkQcdbwuMj8Q
- 微服务讨论 https://maimai.cn/web/gossip_detail/30655604?src=app&webid=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6MzA2NTU2MDQsInUiOjE5NDEzNTc0MCwiZWdpZCI6ImRkMzEyOGYzYTUzYTQzZTdiNzE5ZDQ0ZGEyMjVhZWM1In0.tGHl6tUYY8hKLdsq6N33pbRLzGy_sTZA0RfcKWcIf-I&share_channel=2&_share_channel=wechat
- vivo 海量微服务架构最新实践 https://mp.weixin.qq.com/s/yNbr7LXhWODdgVKfdnba0g
- 从单体到微服务，腾讯文档微服务网关工程化的演进实践 https://mp.weixin.qq.com/s/tqvaCNcUJtWmmGNmDEan3Q
- 回归单体成为潮流？腾讯文档如何实现灵活架构切换 https://mp.weixin.qq.com/s/qtQTZky4SG4d-x0XE97FIg
- TSF微服务治理实战系列
  - （一）——治理蓝图 https://mp.weixin.qq.com/s/VCwiUQjbtof93Gyo3RufDw
  - （二）——服务路由 https://mp.weixin.qq.com/s/EKfji9KbM7j3OPXx_QzVIA
  - （三）——服务限流 https://mp.weixin.qq.com/s/em2-gWgYq2JwH9XiWSHkIg
  - （四）——服务安全 https://mp.weixin.qq.com/s/yT16s3GnGLRvL_pyo5v77w
- 万字长文分享腾讯云原生微服务治理实践及企业落地建议 https://mp.weixin.qq.com/s/TKbYv_yGga7cDqLMt12dYw
- 深入微服务核心：从架构设计到规模化 https://mp.weixin.qq.com/s/M9hfVaMPeLjER4l3Mi3Eww
- 微服务与分布式系统设计 https://mp.weixin.qq.com/s/wg_EkeogSkjGaChvsLsaVw
- 微服务高可用容灾架构设计 https://mp.weixin.qq.com/s/WLO4JI46JkaXYK1cqTDnLg
- 业务系统架构实践总结 https://mp.weixin.qq.com/s/pNfC7klCZTKhXwC4t5V7BA
- DEATH BY A THOUSAND MICROSERVICES https://renegadeotter.com/2023/09/10/death-by-a-thousand-microservices.html

微服务带来的问题
- 2024 年了，之前搞微服务的公司你们还好么 https://v2ex.com/t/1057052
- 微服务到底在哪个方面让开发、维护简单了？https://v2ex.com/t/1082359#r_15417959


## kubernetes 容器编排

k8s让运维只关注硬件底层设备、不关注具体的应用部署，开发者可以自行管理应用的发布和部署。


## 微服务

微服务 → 服务拆分、独立发布部署（解耦）、更新扩容 → k8s管理发布部署

### 单体应用到微服务架构的演变

单体应用（更容易部署、但难扩容和维护）

↓

多服务系统（更容易扩容、但稍难部署）

### 微服务拆分基本原则
- 由粗到细，避免过度拆分，遵循渐进式演进的原则
- 不同服务之间应该是正交的，不要你中有我我中有你
- 避免环形依赖，服务依赖关系应该是有向无环图
- 避免不同服务之间共享同一个数据库

### 微服务缺点

调试、debug

DevOps（开发、QA、运维）

在整个软件开发的生命周期结合在一起（数据层不处理）

### 最佳实践
![microservice best practices](images/microservice-best-practices.webp)
