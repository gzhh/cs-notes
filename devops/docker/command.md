## 常见命令
- https://docs.docker.com/reference/cli/docker/

### 查看 docker 信息

docker info

- 代理信息
    - HTTP Proxy
    - HTTPS Proxy
- 镜像仓库信息
    - Registry Mirrors

### 基础命令

基础镜像（scratch、alpine 都支持编译后 go 程序的允许）

镜像拉取

docker pull ubuntu:20.04

镜像查看

docker image ls

查看docker信息（包括镜像、容器等）

docker inspect alpine

容器查看

docker ps

运行容器（终端交互，退出容器后即删除容器）

docker run -it --rm ubuntu:20.04 /bin/bash

开启/关闭/重启容器

docker start/stop/restart

查看容器中正在运行的进程（top）

docker top container_id/container_name

查看镜像、容器、数据卷所占用的空间

docker system df 

查看和批量清除“虚悬镜像”-即仓库名为<none>的镜像

docker image ls -f dangling=true

docker image prune

docker container prune # PS: 慎重使用

删除镜像

docker image rm image_id 或 docker rmi image_id

删除容器

docker rm container_id

运行容器，并进行端口映射（host_port:container_port）

docker run --name webserver -d -p 80:80 nginx

进入容器

docker exec -it webserver bash

查看容器内部的具体改动

docker diff webserver

查看镜像内的历史记录

docker history

将容器保存为镜像（慎用 docker commit，是个黑盒操作；使用Dockerfile构建镜像）
```
docker commit \
    --author "Tao Wang <twang2218@gmail.com>" \
    --message "修改了默认网页" \
    webserver \
    nginx:v2
```

### 容器退出时自动重启
重启策略

`docker run --restart=always xxx`

Ref
- https://www.jianshu.com/p/14fe95533d12
- https://www.cnblogs.com/ExMan/p/12982898.html
- https://stackoverflow.com/questions/41555884/docker-what-does-docker-run-restart-always-actually-do
