# Command

### 帮助命令

例子：查看 my.cnf 配置文件位置

mysql --verbose --help | grep my.cnf

### 系统变量

查看当前session

show variables like 'default_storeage_engine';

show variables like 'max_connections';

show variables like 'query_cache_size';

show variables like 'default%';

查看全局变量

show global like '';

**全局变量 GLOBAL**

影响服务器的整体操作

SET GLOBAL option=option_value

**会话变量 SESSION**

影响某个客户端连接的操作

SET option=option_value

### 状态变量

显示服务器程序的运行状态

show status like 'threads_connected%';

show status like 'innodb_rows_updated%';

show status like 'threads%';

### 字符集

**字符集 charset**

show character set like '';

show charset like 'utf8';

utf8 # mysql的utf8是阉割的utf8，字符编码长度为 1-3

utf8mb4 # mysql的utf8mb4才是正常utf8，字符编码长度为1-4

mysql8.0以前默认是utf8，以后是utf8mb4

**字符比较规则 collation**

show collation like 'utf8%';

utf8_general_ci # 通用比较规则

utf8mb4_general_ci

**字符集比较规则级别**

服务器级别

show variables like 'character_set_server';

show variables like 'collation_server';

数据库级别

show variables like 'character_set_database';

show variables like 'collation_database';

表级别

列级别
