## 问题排查&日志分析

- Docker 问答录 https://blog.lab99.org/post/docker-2016-07-14-faq.html


### github 仓库 issue 搜索相关问题

[https://github.com/docker/for-mac/issues](https://github.com/docker/for-mac/issues)

### docker desktop 守护进程日志

/Users/gzhh/Library/Containers/com.docker.docker/Data/log/vm/*

~/Library/Containers/com.docker.docker/Data/log/vm/dockerd.log

~/Library/Containers/com.docker.docker/Data/log/vm/docker.log

系统预占空间太大
- 具体位置 /Library/Containers/com.docker.docker/Data/vms/0/data/Docker.raw
- 可在 docker desktop 中调整 Settings -> Resources -> Disk image size 


### docker destop 配置文件

~/Library/Group\ Containers/group.com.docker/settings.json

### 查看镜像相关信息

delv image_name
