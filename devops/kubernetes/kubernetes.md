# Kubernetes
## 概念
### 解决问题
- 容器编排、调度，容器云、集群管理
- 管理网关、水平扩展、监控、备份、灾难恢复

### 优势功能
- 服务发现
- 扩容
- 负载均衡
- 自恢复
- 领导者的选举

### 架构

架构了解：[https://www.youtube.com/watch?v=s_o8dwzRlu4](https://www.youtube.com/watch?v=s_o8dwzRlu4)

- master 主节点
    - API Server
        > k8s 集群的入口点
        - API kube-apiserver 负责 API 服务
        - UI
        - CLI
    - Scheduler
        > kube-shceduler 负责 pod 调度
    - Controller Manager
        > controller-manager 负责容器编排
    - etcd
        > k8s 集群的备份存储，可以持久化集群配置
- node 工作节点
    - kubelet
        - 主要负责同容器运行时（比如 Docker 项目）打交道，依赖 CRI (Container Runtime Interface) 远程调用接口
        - 还通过 grpc 同一个叫做 Device Plugin 的插件进行交互
        - 还通过网络插件和存储插件为容器配置网络和持久化存储
    - 容器 runtime
    - kube-proxy
- master 与 工作节点的关系
    node → pod → container
    ![架构图](images/kubernetes-architecture.png)
- pod 副本数量
    可手动调整副本量或由 kubernetes 自动调整

### 概念

ReplicationController
- [https://kubernetes.io/zh/docs/concepts/workloads/controllers/replicationcontroller/](https://kubernetes.io/zh/docs/concepts/workloads/controllers/replicationcontroller/)

### 原理
```
container {

    Linux Namespace

    cgroup (linux)

}
```

每个容器使用自己的网络接口（namespace）

限制进程使用的系统资源（cgroup）

```
k8s支持容器runtime {

    Docker

    CRI-O

    containerd

}
```


### k8s cluster

> master node（管理 woker node）

1.调度器

2.API server

3.控制管理

4.etcd

> woker node

1.container

2.kubelet

3.kube-proxy

> 在k8s中运行一个应用（实现+理论）

1.打包发布镜像到注册商

2.发布 k8s 的 API Server 的描述（各镜像容器直接的关系）


### 概念

- Deployment - Maintains a set of identical pods, ensuring that they have the correct config and that the right number of them exist.
- Pods - runs one or more closely related containers
- Services - sets up networking in a Kubernetes cluster

[pod](concept-pod.md)

[副本机制和其他控制器](concept-replica.md)

[service](concept-service.md)

[Deployment](concept-deployment.md)

[StatefulSet](concept-statefulset.md)
