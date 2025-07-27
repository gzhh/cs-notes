## pod

### 创建、日志、暴露端口
创建流程
- kubectl `POST /pods`
- API Server
- etcd `add pod`
- scheduler `select node`
- etcd `bind node and pod`
- cri `run cri, etc Docker`

检查现有 pod 的 YAML 描述文件

- kubectl get pod pod_name -o yaml

使用 kubectl 来创建 pod 资源

- kubectl create -f kubia-manual.yaml
    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: kubia-manual
    spec:
      containers:
      - name: kubia
        image: luksa/kubia
        ports:
        - containerPort: 8080
          protocol: TCP
    ```

查看应用程序日志

- kubectl logs pod_name
- kubectl logs pod_name -c container_name

向 pod 发送请求（不通过 services）

- kubectl port-forward pod_name host_port:pod_port

### 使用标签组织 pod

创建带标签的 pod

- kubectl create -f kubia-manual-with-labels.yaml
    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: kubia-manual-v2
      labels:
        creation_method: manual
        env: prod
    spec:
      containers:
      - name: kubia
        image: luksa/kubia
        ports:
        - containerPort: 8080
          protocol: TCP
    ```

查看带标签列的 pod

- kubectl get pod --show-labels
- kubectl get pod -L label_key1,label_key2
- kubectl get pods -o wide --show-labels -A

修改现有 pod 的标签（新加 label 和 修改 label）

- kubectl label pod pod_name label_key=label_value
- kubectl label pod pod_name label_key=label_value --overwrite

使用标签选择器列出 pod 子集

- kubectl get pod -l label_key=label_value
- kubectl get pod -l key1=value1,key2=value2
- kubectl get pod -l label_key
- kubectl get pod -l '!label_key'
- key≠value、key in (v1, v2)、key notin (v1, v2)

### 使用标签限制 node

给一个节点添加 label

- kubectl label node node_name label_key=label_value

使用标签过滤器过滤 node

- kubectl get nodes -l label_key=label_value

创建一个 pod，并调度到指定 node

- kubectl create -f kubia-gpu.yaml
    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: kubia-gpu
    spec:
      nodeSelector:
        gpu: "true"
      containers:
      - image: luksa/kubia
        name: kubia
    ```

### 注解 pod

给 pod 添加注解

- kubectl annotate pod pod_name some_description

查看 pod 信息（Annotations 注解）

- kubectl describe pod pod_name

### Namespace

查看集群中的所有 namespace

- kubectl get namespace
- kubectl get ns

查看指定 namespace 中的 pod（默认 ns 为 default）

- kubectl get pod --namespace kube-system
- kubectl get pod -n kube-system

创建一个 namesapce（使用 kubeclt 命令或 yaml）

- kubectl create namespace custom-namespace
- kubectl create - f custom-namespace.yaml
    ```yaml
    apiVersion: v1
    kind: Namespace
    metadata:
      name: custom-namespace
    ```

创建资源时指定 namespace

- kubectl create -f your.yaml -n namespace_name

### 停止和移除 pod

删除 pod（实际上是终止 pod 内的所有容器）

- kubectl delete pod pod_name

使用标签选择器删除 pod

- kubectl delete pod -l label_key=label_value

通过删除整个命名空间来删除 pod

- kubectl delete ns namespace_name

删除命名空间中的所有 pod，但保留命名空间（这里 rc 创建的 pod，删除后会自动新建 pod，必须先删除 rc）

- kubectl delete pod --all

删除命名空间中的(几乎)所有资源（rc、pod、service 等）

- kubectl delete all --all
