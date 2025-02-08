# Clickhouse
- https://clickhouse.com/docs/en/

概念
- 列式存储数据库 OLAP https://www.jianshu.com/p/a5bf490247ea


## Best Practice


## 常见操作

1. 按 partition 删除iiiii
    
    CREATE TABLE dmp.market_media_retention_data (
    
    ...
    ) ENGINE = AggregatingMergeTree() PARTITION BY dt
    
    ALTER TABLE dmp.market_media_retention_data DELETE WHERE dt = '20211230';
    
    ALTER TABLE dmp.market_media_retention_data DROP partition '20220121'; 
    
2. 更新
    
    *注意不能对索引列进行数据更新*
    
    ALTER TABLE [db.]table UPDATE column1 = expr1 [, ...] WHERE filter_expr;
