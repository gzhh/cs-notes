# Version

### 1. SQL Mode

SELECT @@GLOBAL.sql_mode;
SELECT @@SESSION.sql_mode;

SET GLOBAL sql_mode = '*modes*';
SET SESSION sql_mode = '*modes*';

mysql5.7 & mysql8.0 默认（严格模式）

ONLY_FULL_GROUP_BY, STRICT_TRANS_TABLES, ..., NO_ENGINE_SUBSTITUTION

https://dev.mysql.com/doc/refman/8.0/en/sql-mode.html

mysql5.6 默认

NO_ENGINE_SUBSTITUTION

https://dev.mysql.com/doc/refman/5.7/en/sql-mode.html

可能导致问题

1). incorrect decimal value '' for column

严格模式下 decimal 类型的字段值不能从字符串转化，只能通过数值存

非严格模式下可以
