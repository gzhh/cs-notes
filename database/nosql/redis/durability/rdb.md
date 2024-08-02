# RDB
原理

1. fork 子进程
2. 写数据到临时 RDB 文件
3. 替换原来 rdb 文件

bgsave 问题
- Async-fork解读与Redis实战 https://mp.weixin.qq.com/s/uEdP6wQmENztkO1uX0NtqQ

### 触发机制

save 同步 ⇒ 通常会阻塞redis

缺点：阻塞客户端命令

bgsave 异步 ⇒ 不会阻塞redis，但是会fork一个子进程来创建rdb

优点：不阻塞客户端命令

自动（满足任一就会被执行）

配置 seconds change

save 900 1

save 300 10

save 60 10000

### RDB缺点

耗时、耗性能

fork：消耗内存，copy-on-write策略

Disk IO：IO性能

不可控、丢失数据
