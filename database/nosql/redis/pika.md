# Pika
- https://github.com/OpenAtomFoundation/pikiwidb
- https://github.com/OpenAtomFoundation/pikiwidb/blob/unstable/README_CN.md

介绍
- Pika 是一个以 RocksDB 为存储引擎的的大容量、高性能、多租户、数据可持久化的弹性 KV 数据存储系统，完全兼容 Redis 协议，支持其常用的数据结构。

总结
- 数据存储在磁盘上（基于 RocksDB），RocksDB 原生持久化（写磁盘）
- 兼容 Redis 命令
- 基于 Binlog 的异步复制机制
- 能存储 TB 级数据
- 内存占用：只缓存热点数据，内存占用低
- 访问延迟：稍高（毫秒级），取决于磁盘性能
- 使用场景：大数据量 KV 持久化、冷数据缓存、队列存储
