# undo log

### 简介

事务执行到一半就结束，但是事务在执行过程中可能已经修改了很多东西。为了保证事务的原子性，我们需要改回原来的样子，这个过程就称为回滚。

### undo日志格式

end of record

undo type

undo no

table id

主键各列信息

start of record

### 关联概念

trx_id

roll_pointer
