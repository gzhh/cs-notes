# MySQL的数据目录 
### 简介

查看 MySQL 数据的存放目录

show variables like "datadir";

在该目录下每个具体的数据库会会对应一个目录，子目录里面会存放具体的数据文件

InnoDB

数据根目录下有一个系统表空间文件（ibdata1，会不断扩容）

.frm 表结构

.ibd 表数据（独立表空间）

MyISAM

.frm 表结构

.MYD 表数据

.MYI 表索引
