# 事务 mvcc 

### 事务日志
1. 内存修改
2. 持久化到磁盘上的事务日志（追加方式）顺序IO
3. 日志中存在未写入磁盘，系统死机再次重启数据恢复

### 简介

Multi-Version Concurrency Control 多版本并发控制

每对记录进行一次改动，都会记录一条 undo 日志。每条 undo 日志也都有一个 roll_pointer 属性（ Insert 操作对应的 undo 日志没有该属性，因为 Insert 操作的记录并没有更早的版本），通过这个属性可以将这些 undo 日志串成一个链表。

在每次更新该记录后，都会将旧值放到一条 undo 日志中（就算是该记录的一个旧版本）。随着更新次数的增多，所有的版本都会被 roll_pointer 属性连接成一个链表，这个链表称为版本链。版本链的头节点就是当前记录的最新值。另外，每个版本中还包含生成该版本时对应的事务id。之后会用这个记录的版本链来控制并发事务访问相同记录时的行为，我们把这种机制称为“多版本并发控制”。

### ReadView

为了判断版本链中的哪个版本是当前事务可见的，提出了 ReadView 概念。

**ReadView内容**

- m_ids 在生成ReadView时，当前系统活跃的读写事务的事务id列表
- min_trx_id 在生成ReadView时，当前系统中活跃的读写事务中最小的事务i，也就是 m_ids 中的最小值
- max_trx_id 在生成ReadView时，系统应分配给下一个事务的事务id值
- creator_trx_id 生成该 ReadView 的事务的事务id

**判断记录的某个版本是否可见步骤**

- 如果被访问版本的trx_id属性值与ReadView中的creator_trx_id值相同，意味者当前事务在访问它自己修改过的记录，所以该版本可以被当前事务访问。
- 如果被访问版本的trx_id属性值小于ReadView中的min_trx_id值，表明生成该版本的事务在当前事务生成ReadView前已经提交，所以该版本可以被当前事务访问。
- 如果被访问版本的trx_id属性值大于或等于ReadView中max_trx_id值，表明生成该版本的事务在当前事务生成的ReadView后才开启，所以该版本不可以被当前事务访问。
- 如果被访问版本的trx_id属性值在ReadView的min_trx_id和max_trx_id之间，则需要判断trx_id属性值是否在m_ids列表中。如果在，锁门创建ReadView时生成该版本的事务还是活跃的，该版本不可以被访问；如果不在，说明创建ReadView时生成该版本的事务已经被提交，该版本可以被访问。

### 其他

InnoDB 利用了“所有数据都有多个版本”的这个特性，实现了“秒级创建快照”的能力。

InnoDB 的行数据有多个版本，每个数据版本有自己的 row trx_id，每个事务或者语句有自己的一致性视图。普通查询语句是一致性读，一致性读会根据 row trx_id 和一致性视图确定数据版本的可见性。

对于可重复读，查询只承认在事务启动前就已经提交完成的数据；

对于读提交，查询只承认在语句启动前就已经提交完成的数据；

而当前读，总是读取已经提交完成的最新版本。

为什么表结构不支持“可重复读”？这是因为表结构没有对应的行数据，也没有 row trx_id，因此只能遵循当前读的逻辑。

总结

1.innodb支持RC和RR隔离级别实现是用的一致性视图(consistent read view)

2.事务在启动时会拍一个快照,这个快照是基于整个库的.

基于整个库的意思就是说一个事务内,整个库的修改对于该事务都是不可见的(对于快照读的情况)

如果在事务内select t表,另外的事务执行了DDL t表,根据发生时间,要嘛锁住要嘛报错(参考第六章)

3.事务是如何实现的MVCC呢?

(1)每个事务都有一个事务ID,叫做transaction id(严格递增)

(2)事务在启动时,找到已提交的最大事务ID记为up_limit_id。

(3)事务在更新一条语句时,比如id=1改为了id=2.会把id=1和该行之前的row trx_id写到undo log里,

并且在数据页上把id的值改为2,并且把修改这条语句的transaction id记在该行行头

(4)再定一个规矩,一个事务要查看一条数据时,必须先用该事务的up_limit_id与该行的transaction id做比对,

如果up_limit_id>=transaction id,那么可以看.如果up_limit_id<transaction id,则只能去undo log里去取。去undo log查找数据的时候,也需要做比对,必须up_limit_id>transaction id,才返回数据

4.什么是当前读,由于当前读都是先读后写,只能读当前的值,所以为当前读.会更新事务内的up_limit_id为该事务的transaction id

5.为什么rr能实现可重复读而rc不能,分两种情况

(1)快照读的情况下,rr不能更新事务内的up_limit_id,

而rc每次会把up_limit_id更新为快照读之前最新已提交事务的transaction id,则rc不能可重复读

(2)当前读的情况下,rr是利用record lock+gap lock来实现的,而rc没有gap,所以rc不能可重复读

