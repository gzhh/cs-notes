# DevOps

### mac下启动/停止/重启mysql服务

启动MySQL服务

sudo /usr/local/MySQL/support-files/mysql.server start

停止MySQL服务

sudo /usr/local/mysql/support-files/mysql.server stop

重启MySQL服务

sudo /usr/local/mysql/support-files/mysql.server restart


### 从一个表同步数据到另一个表
从一个表中复制所有的列插入到另一个已存在的表中：

INSERT INTO *table2*SELECT * FROM *table1;*

只复制希望的列插入到另一个已存在的表中：

INSERT INTO *table2(column_name(s))*SELECT *column_name(s)*FROM *table1;*


### 在线扩展表结构
- MySQL如何在线扩展表结构，内核原理 https://mp.weixin.qq.com/s/Jy5al99tPgxknt1ZArCveg


### MySQL Load file

load data local infile

配置

show variables like '%local%';

local_infile  OFF ⇒ ON
