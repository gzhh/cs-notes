## MySQL Log

### 原理理解
WAL日志逻辑
- 刷盘，还是不刷盘 https://mp.weixin.qq.com/s/qQSkg3tp9LblnPUvfIdqhw
- 预写日志WAL的核心思路 https://mp.weixin.qq.com/s/Sduy43sJo_tR-tyoaa4CpA
- ARIES，数据恢复算法 https://mp.weixin.qq.com/s/cjv3YAzORYN1B0wVBvY2EQ


### MySQL Server Logs
- Binary Log: https://dev.mysql.com/doc/refman/8.0/en/binary-log.html
- Slow Query Log: https://dev.mysql.com/doc/refman/8.0/en/slow-query-log.html

mysq server binlog 归档（逻辑日志），追加写

通过开启慢查询日志，分析查询慢的sql语句，优化sql语句，提升程序性能。
- https://www.cnblogs.com/luyucheng/p/6265594.html

### InnoDB 日志
- Redo Log (physically): https://dev.mysql.com/doc/refman/8.0/en/innodb-redo-log.html
- Undo Log: https://dev.mysql.com/doc/refman/8.0/en/innodb-redo-log.html

Innodb redo log 实现 crash-safe（物理日志）

innodb undo log

