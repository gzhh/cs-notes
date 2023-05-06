## Minikube

资源链接：

[Minikube github](https://github.com/kubernetes/minikube) implements a local Kubernetes cluster

[Minikube doc](https://minikube.sigs.k8s.io/docs/)

[Minikube - Kubernetes本地实验环境](https://yq.aliyun.com/articles/221687/)

前置条件：

HyperKit (默认装了Docker会自动安装，但是版本可能太老，需要卸载重新安装)

安装Minikube无法访问k8s.gcr.io的简单解决办法（使用阿里云镜像）

minikube start --driver=hyperkit

minikube start --driver=hyperkit --image-repository='registry.cn-hangzhou.aliyuncs.com/google_containers'

minikube config set driver hyperkit

minikube status

minikube stop

minikube delete

升级minikube

brew upgrade minikube

### 默认集群

查看当前集群

kubectl config get-contexts

切换集群为默认集群

kubectl config use-context docker-desktop

### 本地服务 Minikube

相当于单集群

ps -ef | grep minikube

minikube help

minikube start --vm-driver hyperkit

minikube dashboard

minikube status

minikube stop

minikube delete

minikube ssh

- docker images
- docker ps

minikube kubectl

minikube service xxx_service_name

minikube tunnel

minikube logs
