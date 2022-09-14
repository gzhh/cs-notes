# MySQL notes

## 锁相关

### 常见问题
1. 为什么 innodb 能够支持行锁，而 myisam 不行呢
- 答：innodb 是聚簇索引，索引和数据存一块；myisam 是非聚簇索引，索引和数据分开存储，实现行锁起来复杂。
