# MySQL
文档
- https://dev.mysql.com/doc/refman/8.0/en/
- https://dev.mysql.com/doc/refman/5.7/en/
- https://dev.mysql.com/doc/refman/8.0/en/faqs.html
- https://dev.mysql.com/doc/refman/5.7/en/faqs.html
- https://dev.mysql.com/doc/refman/8.0/en/indexes.html
- https://dev.mysql.com/doc/refman/8.0/en/glossary.html

MySQL 相关程序
- 服务端程序：https://dev.mysql.com/doc/refman/8.0/en/programs-server.html
- 客户端程序：https://dev.mysql.com/doc/refman/8.0/en/programs-client.html
- 管理效率程序: https://dev.mysql.com/doc/refman/8.0/en/programs-admin-utils.html

MySQL 服务器管理
- mysqld 服务器相关配置：https://dev.mysql.com/doc/refman/8.0/en/mysqld-server.html
- https://dev.mysql.com/doc/refman/8.0/en/data-directory.html
- https://dev.mysql.com/doc/refman/8.0/en/system-schema.html
- 服务器日志：https://dev.mysql.com/doc/refman/8.0/en/server-logs.html
- 调试 MySQL：https://dev.mysql.com/doc/refman/8.0/en/debugging-mysql.html

安全相关
- 访问控制：https://dev.mysql.com/doc/refman/8.0/en/access-control.html

备份和恢复
- 备份和恢复类型：https://dev.mysql.com/doc/refman/8.0/en/backup-types.html
- mysqldump 使用：https://dev.mysql.com/doc/refman/8.0/en/using-mysqldump.html
- 防灾设计：https://mp.weixin.qq.com/s/ZLHbfFAbR78JE3qBXXZkqw

优化相关
- https://dev.mysql.com/doc/refman/8.0/en/optimization.html
- 语句、索引、数据库结构、表、查询计划、查询优化器、缓冲和缓冲、锁操作、MySQL 服务器优化、性能

实现细节
- 语言结构：https://dev.mysql.com/doc/refman/8.0/en/language-structure.html
- 字符集：https://dev.mysql.com/doc/refman/8.0/en/charset.html
- 数据类型：https://dev.mysql.com/doc/refman/8.0/en/data-types.html
- 函数和操作符：https://dev.mysql.com/doc/refman/8.0/en/functions.html

SQL 语句
- https://dev.mysql.com/doc/refman/8.0/en/sql-statements.html
- https://dev.mysql.com/doc/refman/8.0/en/sql-data-definition-statements.html
- https://dev.mysql.com/doc/refman/8.0/en/sql-data-manipulation-statements.html
- ...

InnoDB 存储引擎
- https://dev.mysql.com/doc/refman/8.0/en/innodb-storage-engine.html
- https://dev.mysql.com/doc/refman/8.0/en/mysql-acid.html
- https://dev.mysql.com/doc/refman/8.0/en/innodb-multi-versioning.html
- InnoDB 行格式：https://dev.mysql.com/doc/refman/8.0/en/innodb-row-format.html

其他存储引擎
- https://dev.mysql.com/doc/refman/8.0/en/storage-engines.html

复制
- https://dev.mysql.com/doc/refman/8.0/en/replication.html
- https://dev.mysql.com/doc/refman/8.0/en/group-replication.html

Schema
- MySQL 数据目录：https://dev.mysql.com/doc/refman/8.0/en/data-directory.html
- MySQL 数据字典（数据库对象信息存储）：https://dev.mysql.com/doc/refman/8.0/en/data-dictionary.html
- https://dev.mysql.com/doc/refman/8.0/en/system-schema.html
- https://dev.mysql.com/doc/refman/8.0/en/information-schema.html
- https://dev.mysql.com/doc/refman/8.0/en/performance-schema.html
- https://dev.mysql.com/doc/refman/8.0/en/sys-schema.html


## 标准
阿里巴巴MySQL规范
- 建表 https://mp.weixin.qq.com/s/Kif4yvRJGO68jXPErxwNXw
- 建表 https://mp.weixin.qq.com/s/h58L7XxIq5MibmGN-GAHCQ
- https://mp.weixin.qq.com/s/xm-8caxn3sMwXKpeq1w0uA


## 配置相关
1. 查看 MySQL 配置文件路径

```
 ~ mysql --verbose --help | grep my.cnf
                      order of preference, my.cnf, $MYSQL_TCP_PORT,
/etc/my.cnf /etc/mysql/my.cnf /usr/local/mysql/etc/my.cnf ~/.my.cn
```

如果没有，则在上面路径中新建一个

2. 配置文件内容
参考
- https://dev.mysql.com/doc/refman/5.7/en/server-configuration-defaults.html

开启 binlog 日志
- https://www.cnblogs.com/sportsky/p/16357479.html

binlog 日志逻辑
- https://www.cnblogs.com/martinzhang/p/3454358.html


## 锁相关
### 常见问题
1. 为什么 innodb 能够支持行锁，而 myisam 不行呢
- 答：innodb 是聚簇索引，索引和数据存一块；myisam 是非聚簇索引，索引和数据分开存储，实现行锁起来复杂。
