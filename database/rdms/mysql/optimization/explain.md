# Explain
查询性能 explain

https://dev.mysql.com/doc/refman/8.0/en/execution-plan-information.html

使用 explain 查看某个查询的执行计划

### 用法

**简单用法**

explain 查询语句

**高级用法**

explain format=json 查询语句 \G

EXPLAIN 输出格式
| Column | JSON Name | Meaning | version(默认5.6) | value | 解释 |
|---|---|---|---|---|---|
| id | select_id | The SELECT identifier | |  |  |
| select_type | None | The SELECT type | |  |  |
| table | table_name | The table for the output row | |  |  |
| partitions | partitions | The matching partitions | 8.0 |  |  |
| type | access_type | The join type <br> 针对单表访问的方法 | |  |  |
| possible_keys | possible_keys | The possible indexes to choose | |  |  |
| key | key | The index actually chosen | |  |  |
| key_len | ken_length | The length of the chosen key <br> 表示本次查询中，所选择的索引长度有多少字节，通常我们可借此判断联合索引有多少列被选择了 | | 参考：https://imysql.com/2015/10/20/mysql-faq-key-len-in-explain.shtml |  |
| ref | ref | The columns compared to the index <br> 当使用索引列等值查询时，与索引列进行等值匹配的对象信息 | |  |  |
| rows | rows | Estimate of rows to be examined <br> 预估需要读取的记录条数 | |  |  |
| filtered | filtered | Percentage of rows filtered by table condition <br> 针对预估需要读取的记录，经过搜索条件过滤后剩余记录条数的百分比 | |  |  |
| Extra | None | Additional information |  |  |  |

### 索引合并

1. insersection
    
    where c1 = 'c1v' and c2 = 'c2v'
    
    可以两个索引取交集后再回表
    
2. Union
    
    where c1 = 'c1v' or c2 = 'c2v'
    
    可以两个索引取并集后再回表
    
3. Sort-Union
    
    where c1 > 'c1v' or c2 > 'c2v'
    
    先排序，再合并两个索引取并集再回表
    

### 格式

explain selct ...

或

explain format=json select ...

查看成本信息 cost_info

### 例子 index: field1_field2

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

### 参考

https://blog.csdn.net/weixin_30126739/article/details/113122041
