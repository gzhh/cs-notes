## 用例
- https://github.com/jxlwqq/kubernetes-examples

### 部署一个web应用

部署nodejs应用

kubectl run kubia --image=luksa/kubia --port=8080

运行一个 ReplicationController

kubectl apply -f kubia-rc.yaml

- replicationcontroller/kubia created

查看 ReplicationController状态

kubectl describe replicationcontrollers/kubia

```yaml
apiVersion: v1
kind: ReplicationController
metadata:
  name: kubia
spec:
  replicas: 1
  selector:
    app: kubia
  template:
    metadata:
      labels:
        app: kubia
    spec:
      containers:
      - name: kubia
        image: luksa/kubia
        ports:
        - containerPort: 8080
```

查看pod及其状态（是否Running）

kubctl get pods

创建一个服务对象（需要告知Kubernetes对外暴露的ReplicationController）

kubectl expose rc kubia --type=LoadBalancer --name kubia-http

查看服务（是否分配了外部IP，通过 EXTERNAL_IP:PORT 即可访问 web 应用）

kubectl get services

- Minikube 不 支持 LoadBalancer 类型的服务，因此服务不会有外部 IP。但是可以通过外部端口访问服务 。
    
    minikube service kubia-http
    

增加期望的副本数（扩容pod，非常简单，让k8s自行去处理）

- kubectl scale rc kubia --replicas=3
    
    

### 系统的逻辑，ReplicationController、pod 和服务是如何组合在一起的

- pod
    
    一个 pod 可以包含任意个容器，pod 有自己独立的私有 IP 地址和主机名
    
- ReplicationController
    
    确保始终存在一个运行中的 pod 实例
    
    通常 rc 用于复制 pod（即创建 pod 的多个副本），如果你的 pod 消失了，那么 rc 将创建一个新的 pod
    
- service
    
    解决 pod 不停删除重建后 ip 地址不断变换的问题，以及在一个固定的 IP 和端口对上对外暴露多个 pod
