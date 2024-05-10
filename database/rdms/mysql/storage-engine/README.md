# Storage Engine

作用：使数据在内存与磁盘数据中进行交互

简单结构：

- 页
    - 组1
        - 行记录
    - 组2
    - 槽位

### InnoDB

优点

1. 默认，性能最好，应用最广泛
2. 共享空间表
3. 性能高

支持

1. 崩溃后的安全恢复
2. 行级锁
3. 事务处理
4. 外键

### MyISAM
- https://dev.mysql.com/doc/refman/8.0/en/myisam-storage-engine.html

特定

1. mysql5.1之前版本默认索引
2. 全文索引
3. 不支持事务、行及锁、安全恢复
4. 表存储在两个文件：MYD数据 和 MyI索引
5. 某些场景下性能好（绝大部分查询的场景下，以前是这样，现在innodb性能也不差）

### 其他引擎

Memory

Archive

Blackhole

CSV
