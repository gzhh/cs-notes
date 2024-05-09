# Kubernetes
- https://github.com/kubernetes/kubernetes
- https://kubernetes.io

Docs
- https://kubernetes.io/docs/home/
- https://kubernetes.io/docs/concepts/
- https://kubernetes.io/blog/
- https://kubernetes.io/docs/reference/

Courses 
- Kubernetes Crash Course for Absolute Beginners https://www.youtube.com/watch?v=s_o8dwzRlu4
- Kubernetes Roadmap - Complete Step-by-Step Learning Path https://www.youtube.com/watch?v=S8eX0MxfnB4

Kubernetes cluster install (local)
- Minikube | Run Kubernetes locally
  - https://github.com/kubernetes/minikube
  - https://minikube.sigs.k8s.io/docs/
- k8s for Docker Desktop
  - https://github.com/AliyunContainerService/k8s-for-docker-desktop
- kind | Kubernetes IN Docker - local clusters for testing Kubernetes
  - https://github.com/kubernetes-sigs/kind
  - https://kind.sigs.k8s.io/
- Tilt | powers microservice development 
  - https://github.com/tilt-dev/tilt

Kubernetes cluster install (prod)
- 问一问目前生产环境最稳的 k8s 部署方式是啥 https://v2ex.com/t/967147
- 自建
  - minikube https://github.com/kubernetes/minikube
  - kubeadm https://github.com/kubernetes/kubeadm
  - Rancher https://github.com/rancher/rancher
  - Sealos https://github.com/labring/sealos
- 轻量化（生产还是慎用，相对坑多）
  - K3s https://github.com/k3s-io/k3s
    - https://k3d.io
    - https://github.com/k3d-io/k3d
  - MicroK8s https://github.com/canonical/microk8s
    - https://microk8s.io/
- 云平台托管
  - GKE https://cloud.google.com/kubernetes-engine
    - https://cloud.google.com/kubernetes-engine/docs/tutorials
    - https://cloud.google.com/kubernetes-engine/docs/deploy-app-cluster
  - EKS
    - https://aws.amazon.com/eks/

工具
- Kompose - Convert Compose to Kubernetes
  - https://github.com/kubernetes/kompose
- Podman - A tool for managing OCI containers and pods.
  - https://github.com/containers/podman
- nerdctl: Docker-compatible CLI for containerd
  - https://github.com/containerd/nerdctl


## 概念
数据平面（Data Plane）和控制平面（Control Plane）
- https://developer.aliyun.com/article/1395252

### 解决问题
- 容器编排、调度，容器云、集群管理
- 管理网关、水平扩展、监控、备份、灾难恢复
- 容器编排器的自我介绍 https://mp.weixin.qq.com/s/F9g-r4yBYDZ1Q9z6uq5feQ

### 优势功能
- 服务发现
- 扩容
- 负载均衡
- 自恢复
- 领导者的选举

### 架构

架构了解：
- https://www.youtube.com/watch?v=s_o8dwzRlu4
- https://mp.weixin.qq.com/s/qlx6aio53CTBXXBXYjHhVw

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
