# API performance
- API性能测试指标以及压测方式 https://zhuanlan.zhihu.com/p/609348456
- 命令 `wrk -t8 -c5000 -d30s http://your-api-server/health`


## 服务器调优
分析
- 1.查看已建立的 HTTP 连接 `netstat -anp | grep ':80' | grep ESTABLISHED | wc -l`
- 2.查看一个进程最多可打开的连接数 `ulimit -n`
- 3.查看 TCP 半连接队列的最大长度（backlog） `cat /proc/sys/net/core/somaxconn`

操作
- 1.提高 backlog 队列长度 `echo 1024 > /proc/sys/net/core/somaxconn`
- 2.查看并调整 ulimit -n（文件描述符限制）`ulimit -n 100000`
- 3.Nginx 层优化（若用 Nginx）
```
# worker_connections × worker_processes ≈ 最大连接数（建议值设高一些）
worker_rlimit_nofile 100000;

events {
    worker_connections 65535;
    use epoll;
}
```
- 4.其他内核参数（可选）
```
# 同时打开连接数限制更高
sysctl -w net.ipv4.ip_local_port_range="1024 65535"
sysctl -w net.ipv4.tcp_tw_reuse=1
```
