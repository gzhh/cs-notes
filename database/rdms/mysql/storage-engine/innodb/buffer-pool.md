# Buffer Pool
- https://dev.mysql.com/doc/refman/8.0/en/innodb-buffer-pool.html

### 简介

为了缓存磁盘中的页，在 MySQL 服务器启动时就向操作系统申请了一片连续的内存（Buffer Pool 缓冲池），默认有 128 MB。
show variables like "innodb_buffer_pool_size";

使用很多链表来管理 buffer pool

show engine innodb status\G;

### 缓存策略

LRU

LRU 优化 ⇒ 热数据、冷数据拆开
