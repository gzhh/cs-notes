## 镜像仓库&代理

### 配置国内镜像仓库
Docker 服务安装好后，默认 registry 指向 hub.docker.com

```
"registry-mirrors": [
    "registry.docker-cn.com",
    "https://hub-mirror.c.163.com",
    "https://mirror.baidubce.com"
]
```

### 在 mac 上配置命令行代理

1. 在 ~/.bash_profile 中设置，需要有自己的代理
```
export http_proxy="http://127.0.0.1:1080"
export https_proxy="http://127.0.0.1:1080"
```

2. source 使 .bash_profile 生效

3. 验证是否生效
```
curl www.google.com
curl --proxy http://127.0.0.1:41091 www.google.com
```

### 在 mac docker desktop 配置守护进程代理

- dockerd 代理
- Container 代理
- docker build 代理

配置 dockerd 代理即可
```
HTTP：http://127.0.0.1:1080
HTTPS： http://127.0.0.1:1080
```

PS：不同版本的 docker for mac 可能会有 bug

参考：[https://note.qidong.name/2020/05/docker-proxy/](https://note.qidong.name/2020/05/docker-proxy/)

### docker pull 不了镜像的解决办法
问题
- [gcr.io](http://gcr.io) 资源国内不能下载（k8s团队和第三方维护的镜像库）
- [quay.io](http://quay.io/) 资源国内下载很慢（red hat运营的镜像库）

解决办法
1. 使用国内镜像代理仓库（例如阿里云）
2. 使用代理下载
3. 使用 github 仓库
