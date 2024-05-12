# Explain
查询性能 explain

https://dev.mysql.com/doc/refman/8.0/en/execution-plan-information.html

使用 explain 查看某个查询的执行计划

## 用法

**简单用法**

explain 查询语句

**高级用法**

explain format=json 查询语句 \G

EXPLAIN 输出格式
| Column | JSON Name | Meaning | version(默认5.6) |
|---|---|---|---|
| id | select_id | The SELECT identifier | |  |
| select_type | None | The SELECT type | |  |
| table | table_name | The table for the output row | |  |
| partitions | partitions | The matching partitions | 8.0 |  |
| type | access_type | The join type <br> 针对单表访问的方法 | |  |
| possible_keys | possible_keys | The possible indexes to choose | |  |
| key | key | The index actually chosen | |  |
| key_len | ken_length | The length of the chosen key <br> 表示本次查询中，所选择的索引长度有多少字节，通常我们可借此判断联合索引有多少列被选择了 | | 参考：https://imysql.com/2015/10/20/mysql-faq-key-len-in-explain.shtml |
| ref | ref | The columns compared to the index <br> 当使用索引列等值查询时，与索引列进行等值匹配的对象信息 | |  |
| rows | rows | Estimate of rows to be examined <br> 预估需要读取的记录条数 | |  |
| filtered | filtered | Percentage of rows filtered by table condition <br> 针对预估需要读取的记录，经过搜索条件过滤后剩余记录条数的百分比 | |  |
| Extra | None | Additional information |  |  |

### Explain 各字段解释
- select_type
    - value
    ```
    SIMPLE 简单子查询，不包含子查询和union
    PRIMARY 包含union或者子查询，最外层的部分标记为primary
    SUBQUERY 一般子查询中的子查询被标记为subquery，也就是位于select列表中的查询
    DERIVED 派生表——该临时表是从子查询派生出来的，位于form中的子查询
    UNION 位于union中第二个及其以后的子查询被标记为union，第一个就被标记为primary如果是union位于from中则标记为derived
    UNION RESULT 用来从匿名临时表里检索结果的select被标记为union result
    DEPENDENT UNION 顾名思义，首先需要满足UNION的条件，及UNION中第二个以及后面的SELECT语句，同时该语句依赖外部的查询
    DEPENDENT SUBQUERY和DEPENDENT UNION相对UNION一样
    ```
- type
    - value
    ```
    const,SYSTEM 单表中最多只有一条匹配行，查询起来非常迅速，所以这个匹配行中的其他列中的值可以被优化器在当前查询中当做常量来处理。

    eq_ref 相对于ref来说就是使用的是唯一索引，对于每个索引键值，只有唯一的一条匹配记录
    ref 使用非唯一性索引或者唯一索引的前缀扫描，返回匹配某个单独值的记录行
    ref_or_null

    RANGE 索引范围扫描 > < between in

    INDEX 索引全扫描

    ALL 全表扫描

    INDEX MERGE 使用多个索引扫描，将结果合并Extra 常显示(Using intersect(index1, index2))
    - intersect, union, sort union

    NULL MYSQL不用访问表或者索引就直接能到结果
    ```
    - value 解释
    ```
    system 当表中只有一条记录并且该表使用的存储引擎的统计数据是精确的，那么对该表的访问就是 system

    const 常量查询（主键和唯一键）

    eq_ref 连接查询，如果被驱动表是通过主键或者不允许存储NULL值的唯一二级索引列等值匹配的方式进行访问的，则对该被驱动表的访问方法就是 eq_ref

    ref 二级索引（等值比较，扫描的是单点区间）

    ref_or_null（等值比较，但多了NULL值比较；扫描区间是NULL区间和单点区间）

    .
    .
    .

    range 二级索引（扫描若干个单点区间或范围区间）

    index（扫描全部二级索引，不用回表操作就可以得到结果，覆盖索引）特例：select * from table_name order by id

    all 全表扫描
    ```
- key_len
    - 参考：https://imysql.com/2015/10/20/mysql-faq-key-len-in-explain.shtml
- ref
    - value
    ```
    当访问类型是const、eq_ref、ref、ref_or_null ... 其中一个时，ref列展示的就是与索引列进行等值匹配的信息。

    可能值有
    const、NULL、func、其他
    ```
- rows
    - value
    ```
    是 MySQL 认为它要检查的行数（仅做参考），而不是结果集里的行数；
    同时 SQL 里的 LIMIT 和这个也是没有直接关系的。

    这个rows就是mysql认为必须要逐行去检查和判断的记录的条数。
    举个例子来说，假如有一个语句 select * from t where column_a = 1 and column_b = 2;
    全表假设有100条记录，column_a字段有索引（非联合索引），column_b没有索引。
    column_a = 1 的记录有20条， column_a = 1 and column_b = 2 的记录有5条。
    那么最终查询结果应该显示5条记录。 explain结果中的rows应该是20. 因为这20条记录mysql引擎必须逐行检查是否满足where条件。
    ```
- Extra
    - value
    ```
    Distinct：查找distinct 值，当mysql找到了第一条匹配的结果时，将停止该值的查询，转为后面其他值查询。
    Full scan on NULL key：子查询中的一种优化方式，主要在遇到无法通过索引访问null值的使用。
    Range checked for each record (index map: N)：通过 MySQL 官方手册的描述，当 MySQL Query Optimizer 没有发现好的可以使用的索引时，如果发现前面表的列值已知，部分索引可以使用。对前面表的每个行组合，MySQL检查是否可以使用range或 index_merge访问方法来索取行。
    SELECT tables optimized away：当我们使用某些聚合函数来访问存在索引的某个字段时，MySQL Query Optimizer 会通过索引直接一次定位到所需的数据行完成整个查询。当然，前提是在 Query 中不能有 GROUP BY 操作。如使用MIN()或MAX()的时候。
    Using filesort：当Query 中包含 ORDER BY 操作，而且无法利用索引完成排序操作的时候，MySQL Query Optimizer 不得不选择相应的排序算法来实现。
    Using index：所需数据只需在 Index 即可全部获得，不须要再到表中取数据。
    Using index for group-by：数据访问和 Using index 一样，所需数据只须要读取索引，当Query 中使用GROUP BY或DISTINCT 子句时，如果分组字段也在索引中，Extra中的信息就会是 Using index for group-by。
    Using temporary：当 MySQL 在某些操作中必须使用临时表时，在 Extra 信息中就会出现Using temporary 。主要常见于 GROUP BY 和 ORDER BY 等操作中。
    Using where：如果不读取表的所有数据，或不是仅仅通过索引就可以获取所有需要的数据，则会出现 Using where 信息。
    Using where with pushed condition：这是一个仅仅在 NDBCluster存储引擎中才会出现的信息，而且还须要通过打开 Condition Pushdown 优化功能才可能被使用。控制参数为 engine_condition_pushdown 。
    Impossible WHERE noticed after reading const tables：MySQL Query Optimizer 通过收集到的统计信息判断出不可能存在结果。
    No tables：Query 语句中使用 FROM DUAL或不包含任何 FROM子句。
    Not exists：在某些左连接中，MySQL Query Optimizer通过改变原有 Query 的组成而使用的优化方法，可以部分减少数据访问次数。
    ```
    - value 解释
    ```
    No tables used 没有 from 子句

    Impossible WHERE 查询语句永远为 FALSE

    Using index 使用覆盖索引执行查询，不需要回表操作

    Using index condition 有些搜索条件中虽然出现了索引列，但却不能充当边界条件来形成扫描区间，也就是不能用来减少需要扫描的记录数量，将会提示该信息。比如 like '%xxx'

    Using where 当某个搜索条件需要在server层进行判断时

    Using join buffer

    Using intersect、union、sort_union

    Zeron limit 当limit的参数为0时

    Using filesort 对结果集中的记录进行排序时

    Using temporary 借助临时表来进行去重、排序等操作
    ```


## 索引合并

1. insersection
    
    where c1 = 'c1v' and c2 = 'c2v'
    
    可以两个索引取交集后再回表
    
2. Union
    
    where c1 = 'c1v' or c2 = 'c2v'
    
    可以两个索引取并集后再回表
    
3. Sort-Union
    
    where c1 > 'c1v' or c2 > 'c2v'
    
    先排序，再合并两个索引取并集再回表
    

## 格式

explain selct ...

或

explain format=json select ...

查看成本信息 cost_info

## 例子 index: field1_field2

where field1 in () and field2=

explain select *
from p_channel
where media_id in (1,126004,157002,217002,337003) and ad_id = 1693403861508

不会命中索引，会全表查询

```
{
	"data":
	[
		{
			"id": 1,
			"select_type": "SIMPLE",
			"table": "p_channel",
			"type": "ALL",
			"possible_keys": "uk_media_ad,media_account",
			"key": null,
			"key_len": null,
			"ref": null,
			"rows": 4373235,
			"Extra": "Using where"
		}
	]
}
```

where field1 = and field2=

explain select *
from p_channel
where media_id = 1 and ad_id = 1693403861508141 and channel_type = 1

会命中索引

```
{
	"data":
	[
		{
			"id": 1,
			"select_type": "SIMPLE",
			"table": "p_channel",
			"type": "ref",
			"possible_keys": "uk_media_ad,media_account",
			"key": "uk_media_ad",
			"key_len": "4",
			"ref": "const",
			"rows": 2186617,
			"Extra": "Using index condition; Using where"
		}
	]
}
```

## 参考

https://blog.csdn.net/weixin_30126739/article/details/113122041
