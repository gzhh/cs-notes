## 副本机制和其他控制器

### 1. 使用存活探针保持pod健康

介绍

- HTTP GET 探针
    
    如果 HTTP 响应状态码是 2xx 或 3xx，则认为探测成功，否则失败
    
- TCP 套接字探针
    
    如果 TCP 连接成功，则探测成功，否则失败
    
- Exec 探针
    
    退出状态码为 0 则探测成功，否则失败
    

配置

```yaml
...
		livenessProbe:
      httpGet: # 一个 HTTP GET 存活探针
        path: /
        port: 8080
			initialDelaySeconds: 15 # 在第一次探测前等待15秒
...
```

### 2. ReplicationController

旨在创建和管理一个pod的多个副本

ReplicationController的三部分

- label selector ( 标签选择器)， 用于确定ReplicationController作用域中有哪些
- replica count (副本个数)， 指定应运行的pod 数量
- podtemplate (pod模板)， 用于创建新的pod副本

使用 ReplicationController的好处

- 确保 一 个 pod (或多个 pod 副本)持续运行， 方法是在现有 pod 丢失时启动
一 个新 pod。
- 集群节点发生故障时， 它将为故障节 点 上运 行的所有 pod (即受
ReplicationController 控制的节点上的那些 pod) 创建替代副本。
- 它能轻松实现 pod的水平伸缩 手动和自动都可以(参见第 15 章中的
pod的水平自动伸缩)。

创建一个 ReplicationController

- kubectl create -f kubia-rc.yaml
    
    ```yaml
    apiVersion: v1
    kind: ReplicationController
    metadata:
      name: kubia
    spec:
      replicas: 3
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
    

查看 ReplicationController 信息

- kubectl get rc
- kubectl describe rc rc_name

**将 pod 移入或移除 ReplicationController 的作用域**

给 ReplicationController 管理的 pod 加标签

- kubectl label pod pod_name label_key=label_value

更改己托管的 pod 的标签

- kubectl label pod pod_name label_key=label_value --overwrite

**修改 pod 模版**

只会影响之后创建的 pod，不会影响现有的 pod

- kubectl edit rc rc_name

**水平缩放 pod**

用kubectl scale命令缩容

- kubectl scale re rc_name --replicas=10

还可以通过编辑 pod 模版来缩放 ReplicationController

- 编辑配置 replicas: 10

**删除ReplicationController**

删除 rc（rc 管理的 pod 也随之被删除）

- kubectl delete rc rc_name

删除 rc ，并不删除 pod，只是让 pod 脱离该 rc 的管理

- kubectl delete rc rc_name --cascade= false

### 3. 使用 ReplicaSet而不是 ReplicationController

ReplicaSet 使用方式机会和 ReplicationController，是新一代的 rc，rs 拥有更强大的标签选择器

创建一个 ReplicationSet

- kubectl create -f kubia-replicaset.yaml
    
    ```yaml
    apiVersion: v1
    kind: ReplicationSet
    metadata:
      name: kubia
    spec:
      replicas: 3
      selector:
    		matchLabels:
    	    app: kubia
      template:
        metadata:
          labels:
            app: kubia
        spec:
          containers:
          - name: kubia
            image: luksa/kubia
    ```
    

查看 ReplicationController 信息

- kubectl get rs
- kubectl describe rs rs_name

更丰富的标签选择器

```yaml
...
	selector:
    matchExpressions:
      - key: app
        operator: In
        values:
         - kubia
...
```

删除 rs

- kubectl delete rs rs_name

### 4. 使用 DaemonSet在每个节点上运行一个pod

守护进程特性

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: kubia
spec:
  replicas: 3
  selector:
    matchExpressions:
      - key: app
        operator: In
        values:
         - kubia
  template:
    metadata:
      labels:
        app: kubia
    spec:
      containers:
      - name: kubia
        image: luksa/kubia
```

- kubectl get ds

### 5. Job：运行执行单个任务的pod

一个可完成的任务中， 其进程终止后， 不再重新启动

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: multi-completion-batch-job
spec:
  completions: 5
  parallelism: 2
  template:
    metadata:
      labels:
        app: batch-job
    spec:
      restartPolicy: OnFailure
      containers:
      - name: main
        image: luksa/batch-job
```

- kubectl get jobs
- kubectl logs job_name

### 6. CronJob：安排Job定期运行或在将来运行一次

在计划的时间内，CronJob资源会创建 Job 资源，然后Job创建pod。

```yaml
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: batch-job-every-fifteen-minutes
spec:
  schedule: "0,15,30,45 * * * *"
  jobTemplate:
    spec:
      template:
        metadata:
          labels:
            app: periodic-batch-job
        spec:
          restartPolicy: OnFailure
          containers:
          - name: main
            image: luksa/batch-job
```

- kubectl get cronjobs
- kubectl get jobs
