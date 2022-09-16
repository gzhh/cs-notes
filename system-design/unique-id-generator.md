# Unique ID Generator
单数据服务可能最先考虑使用传统数据库的自增 ID，但是在分布式系统中自增 ID 不适用，因为 ID 值可能会很大，并且需要跨多个数据服务实例


## Multi-master replication 多主复制集群
缺点
- 在多数据中心下很难扩容
- ID 不能跨多个服务节点的随时间增加
- 当有服务节点添加或删除时，不能扩容


## UUID
优点
- 生成 UUID 简单
- 容易扩容，因为每个节点都可以独立生成 UUID

缺点
- ID 是 128 位而不是常见的 64 位
- ID 不能随时间增长
- ID 不能是纯数字


## Ticket Server
优点
- 可以生成数字 ID
- 容易实现

缺点
- 单点故障问题，如果要引入多个节点避免单点故障，那么又会引入节点数据同步问题


## Twitter Snowflake approach 雪花算法
原理：使用分治思想，将 64 位的 ID 拆分成五部分
- Sign bit：1 位，保留空间
- Timestamp：41 位，存储毫秒时间戳，可以支撑将近 70 年的使用范围
- Datacenter ID：5 位，可以容纳的数据中心数量
- Machinne ID：5 位，每个数据中心可以容纳的机器数量
- Sequence number：12 位，对于每台机器/进程每次生成 ID 时，序列号会按自增 1 的逻辑进行生成. 该值每毫秒会被重置一次 0


## 额外需要考虑的点
- 需要考虑分布式系统中的时钟同步问题
- ID Generator 作为核心服务需要考虑高可用


## 参考
- [System Design Interview](https://book.douban.com/subject/35246417/)
- UUID: https://en.wikipedia.org/wiki/Universally_unique_identifier
- Twitter Snowflake: https://en.wikipedia.org/wiki/Snowflake_ID
- Generate or parse Twitter snowflake IDs with Golang: https://github.com/bwmarrin/snowflake
- Clock synchronization: https://en.wikipedia.org/wiki/Clock_synchronization
- Network Time Protocol: https://en.wikipedia.org/wiki/Network_Time_Protocol
