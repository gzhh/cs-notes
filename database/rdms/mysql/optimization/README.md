# Optimization
- https://dev.mysql.com/doc/refman/8.0/en/optimization.html

## 优化器

### 基于成本的优化

1. 根据搜索条件，找出所有可能使用的索引
2. 计算全表扫描的代价
3. 计算使用不同索引执行查询的代价
4. 对比各种执行方案的代价，找出成本最低的那个方案。

成本最低的方案即执行计划

### 基于规则的优化

1. 移除不必要的括号
2. 常量传递
3. 移除没用的条件
4. 表达式计算
5. ...


## 进一步优化optimizer trace
### 简介

show variables like 'optimizer_trace';

### 用法

set optimizer_trace="enabled=on"

select ...;

select * from infomation_schema.OPTIMIZER_TRACE\G;


## 成本分析
cost_info


## 查询优化总结

### 索引优化

### MySQL查询优化

分析查询慢的原因

优化数据访问

优化长难语句

优化特定类型语句

### 工具

pt_query_digest 工具分析（不使用查询日志-慢）

show profile（首先 set profiling=1）

show status

show processlist

explain [select 语句]

### 总结

1. select 指定需要列代替 select *
2. 优化长难语句
    
    切分查询
    
    分解关联查询
    
3. count(*) 和 count(id)
4. group by 和 order by
5. 子查询和关联查询
6. limit
7. union all 效率高于 union
