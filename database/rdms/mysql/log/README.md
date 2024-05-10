## MySQL Log

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
