## service

kubernetes 服务是一种为一组功能相同的 pod 提供单一不变的接入点的资源，当服务存在时，它的 IP 地址和端口不会改变。客户端通过 IP 地址和端口号建立连接，这些连接会被路由到提供该服务的任意一个 pod 上。通过这种方式，客户端不需要知道每个单独的提供服务的 pod 的地址，这样这些 pod 就可以在集群中随时被创建或移除。

### 1. 创建服务

**通过 kubectI expose 创建服务**

**通过 YAML 描述文件来创建服务**

- kubectl apply -f kubia-svc.yaml
    
    ```yaml
    apiVersion: v1
    kind: Service
    metadata:
      name: kubia
    spec:
      ports:
      - port: 80
        targetPort: 8080
      selector:
        app: kubia
    ```
    

### 2. **服务发现**

通过环境变量发现服务

通过 DNS 发现服务

通过 FQDN 连接服务

PS：无法 ping 通服务 IP 的原因

因为服务的集群 IP 是一个虚拟 IP，并且只有在与 服务端口结合时才有意义。

### 3.通过 endpoint 连接集群外部的服务

**endpoint介绍**

服务并不是和 pod 直接相连的，相反，有一种资源介于两者之间，它就是 Endpoint 资源。

查看 endpoint

kubectl describe svc svc_name

kubectl get endpoints svc_name

手动创建 endpoint

- kubectl apply -f external-service.yaml
    
    ```yaml
    apiVersion: v1
    kind: Service
    metadata:
      name: external-service
    spec:
      ports:
      - port: 80
    ```
    

### 4.将服务暴露给外部客户端
- Kubernetes service external ip pending https://stackoverflow.com/questions/44110876/kubernetes-service-external-ip-pending
- 对比Kubernetes的Nodeport、Loadbalancer和Ingress，什么时候该用哪种 https://cloud.tencent.com/developer/article/1326529

**4.1使用 NodePort 类型的服务**

kubia-svc-nodeport.yaml

```yaml
apiVersion: v1
kind: Service
metadata:
  name: kubia-nodeport
spec:
  type: NodePort
  ports:
  - port: 80
    targetPort: 8080
    nodePort: 30123
  selector:
    app: kubia
```

**4.2通过负载均衡器将服务暴露出来**

LoadBalancer类型的服务是个具有额外的基础设施提供的负载平衡器NodePort服务。

kubia-svc-loadbalancer.yaml

```yaml
apiVersion: v1
kind: Service
metadata:
  name: kubia-loadbalancer
spec:
  type: LoadBalancer
  ports:
  - port: 80
    targetPort: 8080
  selector:
    app: kubia
```

**4.3通过Ingress暴露服务**
- https://github.com/kubernetes/ingress-nginx
- https://github.com/kubernetes/ingress-nginx/blob/main/docs/deploy/index.md
- https://kubernetes.io/docs/concepts/services-networking/ingress/
- 课程实录 | Ingress Controller 的工作原理（上）https://mp.weixin.qq.com/s/VN-CfvRza3k1IHVxvoUwtg
- Ingress Controller 的工作原理 | K8s Ingress Controller 技术细节探讨 —— 陶辉 https://www.bilibili.com/video/BV1r64y1m72f
- https://github.com/projectcontour/contour

Ingress 介绍

- 不同于 LoadBalancer 服务都需要自己的负载均衡器，以及独有的公有 IP 地址，而 Ingress 只需要一个公网 IP 就能为许多服务提供访问
- 当客户端向 Ingress 发送 HTTP 请求时，Ingress 会根据请求的主机名和路径决定请求转发到的服务

创建 kubia-ingress.yaml

```yaml
apiVersion: extensions/v1beta1 # apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: kubia
spec:
  rules:
  - host: kubia.example.com
    http:
      paths:
      - path: /
        backend:
          serviceName: kubia-nodeport
          servicePort: 80
```

查看 ingress

- kubectl get ingress

PS:

1. 多个服务暴露给单个 ingress 
2. 配置 Ingress 处理 TLS 传输

### 5. 就绪探针

类型

- Eeec 探针
- HTTP GET 探针
- TCP socket 探针

kubia-rc-readinessprobe.yaml
