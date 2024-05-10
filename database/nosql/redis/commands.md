# Redis commands

## 单线程为什么执行命令这么快

1. 一次只运行一条命令
    
    为什么这么快
    
    1）纯内存（主要原因）
    
    2）非阻塞 IO（Epoll）
    
    3）避免线程切换和竞态消耗
    
2. 拒绝长（慢）指令
    
    keys
    
    flushall
    
    flushdb
    
    slow lua script

## 通用命令

dbsize 计算key的总数

keys 基本不在生产环境用

type key

exists key

del key [key ...]

expire key seconds 设置key的过期时间

ttl key 查看key剩余的过期时间

persist key 去掉key的过期时间

