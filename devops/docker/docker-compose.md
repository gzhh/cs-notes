## Docker-Compose

查看命令

docker-compose -h

文件挂载
- 宿主机目录不存在则创建
- 宿主机文件不存在则错误，需先创建挂载文件
- nginx 构建要新建一个默认的 default.conf，否则会报错

网络
- 容器间通信一般用 bridge 模式

### 注意点
一、docker-compose version3 与 version 2 的区别
1. v3 的 network 代替之前的 links
2. container_name 很重要
