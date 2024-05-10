# redo log

### 简介

事务提交时，把 mysql 操作写入日志，刷到磁盘。即使之后系统崩溃，重启之后按照日志记录的操作内容重新更新一下数据页，那么该事务对数据库中所做的修改就可以被恢复出来，实现持久化的需求。

上述这样的日志称为 redo日志

**好处**

redo 日志占用空间非常小

redo 日志是顺序写入磁盘的，顺序I/O，写完一个 block 继续写下一个 block

redo 日志格式
| type | space ID | page number | data |
|--------|--------|--------|----|
| 这条redo日志的类型 | 表空间ID | 页号 | 这条redo日志的具体内容 |

由若干个512字节大小的block组成，每一个 block 分为3个部分

- log block header
- log block body
- log block trailer

### redo 日志文件

**redo 日志刷盘时机**

- log buffer 空间不足时
- 事务提交时
    
    只将执行过程中产生的redo日志刷新到磁盘，而不是将所有修改过的页面都刷新到磁盘
    
- 后台线程大约每秒一次的频率将log buffer中的redo日志刷新到磁盘
- 正常关闭服务器时
- 做 checkpoint 时

**在磁盘中的位置**

show variables like 'datadir';

具体文件

ib_buffer_pool

ib_logfile0

ib_logfile1

**checkpoint**

由于redo日志空间优先，需要覆盖可以被覆盖掉的redo日志
