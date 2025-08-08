# Redis 运维

## Best Practice
- https://help.aliyun.com/zh/redis/user-guide/

### 排查内存使用率高
- Ref https://help.aliyun.com/zh/redis/user-guide/troubleshoot-the-high-memory-usage-of-an-apsaradb-for-redis-instance


## 命令
- 参考 http://www.peter-zhou.com/fu-wu-bu-shu/redis/13redisyun-wei-chang-yong-ming-ling.html

### Info
- https://redis.io/docs/latest/commands/info/

> info

### Scanning for big keys
- https://redis.io/docs/latest/develop/connect/cli/#scanning-for-big-keys

> $ redis-cli --bigkeys


## 工具
### Another Redis Desktop Manager
- 目标 redis key path 目录点击右键，选中Memory Analysis，即可看到目标 key path 中的最大 key 列表
