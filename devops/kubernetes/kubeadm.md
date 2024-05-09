## kubeadm
- https://github.com/kubernetes/kubeadm
- https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/

### 安装 k8s 工具

需安装老版本的 k8s（例如1.23 版本），因为默认安装 k8s 1.24 版本，内置的 container runtime 是 conainer 而不是 docker，后续执行 kubeadm init 的时候可能会有问题

查看指定版本

`sudo apt-cache madison kubectl | grep 1.23`

安装

`sudo apt-get install -y kubelet=1.23.8-00 kubeadm=1.23.8-00 kubectl=1.23.8-00`

### 使用 kubeadm 安装 k8s

初始化并指定国内源（因为初始化时需要从 [gcp.io](http://gcp.io) 下载镜像，国内被墙不能之间下载）

`sudo kubeadm init --image-repository [registry.aliyuncs.com/google_containers](http://registry.aliyuncs.com/google_containers)`

安装 oci 网络插件（flannel 或者 calico）

```markdown
# flannel
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

# calico
kubectl apply -f [https://raw.githubusercontent.com/projectcalico/calico/master/manifests/calico.yaml](https://raw.githubusercontent.com/projectcalico/calico/master/manifests/calico.yaml)
```

### 注意

- 默认单节点的集群，pod 不能被调度到 master 节点，需要 删除master节点污点（taint）后才能调度

### 参考

安装参考

- Bootstrapping clusters with kubeadm
    
    [https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/)
    
- 用 kubeadm 在 Ubuntu 上快速构建 Kubernetes 测试集群
    
    [https://jimmysong.io/kubernetes-handbook/practice/install-kubernetes-on-ubuntu-server-16.04-with-kubeadm.html](https://jimmysong.io/kubernetes-handbook/practice/install-kubernetes-on-ubuntu-server-16.04-with-kubeadm.html)
    
- 如何在 Ubuntu 20.04 上安装 Kubernetes – Kubeadm 和 Minikube
    
    [https://os.51cto.com/article/704190.html](https://os.51cto.com/article/704190.html)
    

执行 kubeadm init … 失败参考

- k8s初始化 报错Error getting node“ err=“node
    
    [https://blog.csdn.net/weixin_66536807/article/details/124903478](https://blog.csdn.net/weixin_66536807/article/details/124903478)
    
- kubeadm init shows kubelet isn't running or healthy
    
    [https://stackoverflow.com/questions/52119985/kubeadm-init-shows-kubelet-isnt-running-or-healthy](https://stackoverflow.com/questions/52119985/kubeadm-init-shows-kubelet-isnt-running-or-healthy)
    

若安装失败则需重置 kubeadm

`sudo kubeadm reset`
