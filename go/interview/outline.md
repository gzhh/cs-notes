# Go interview outline
### 面试准备项

- Golang 语言：基础 + 原理 + Coding 能力
- 项目：围绕项目展开技术细节和业务，要有亮点，行业宏观展望
- 八股文：适当准备常见即可，需要对简历中的技术点及细节要熟掌于心
- 算法和数据结构：leetcode hot 100 即可
- 系统设计

PS：此外 “沟通” 极其重要

### 八股文题

- nil
    - https://zhuanlan.zhihu.com/p/105554073
- 类型断言
    - https://juejin.cn/post/6844904153056034823
- 听说go有什么什么的缺陷，你怎么看

### 技术相关

- 基础相关
    - 结构体、接口、组合
    - reflect
- 数据结构
    - slice
    - map
    - channel
    - interface
- 核心机制原理
    - 内存管理
    - GC
    - GMP
    - reflect
- 并发相关
    - goroutine
    - Mutex
    - Context
    - Select
    - 并发原语 sync 包等
- 性能调优相关
    - CPU、内存、IO、网络
    - pprof
    - linux: top、netstat、
    - https://www.cnblogs.com/jiujuan/p/14588185.html
- 常用框架
    - Gin
        - middleware
            - log
            - auth
            - rate limit
    - Gorm
    - gRPC
- 中间件相关
    - mysql
    - redis
    - mq
    - nginx
    - elasticsearch
    - mongodb
- 微服务、分布式相关
    - 微服务服务治理
        - ServiceMesh、istio、etcd 等
    - 分布式概念
        - cap、raft
- DevOps 工程效率相关
    - 脚本自动化
        - linux command
        - makefile
        - 自动化测试
    - CI/CD
        - Github action
        - Gitlab ci
        - jenkins
        - 阿里云效
    - Docker、K8S
        - docker
            - 原理
            - 使用
        - k8s
            - 原理
            - 使用
    - Prometheus、Grafana
