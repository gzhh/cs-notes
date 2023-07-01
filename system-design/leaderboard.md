# Leaderboard system
排行榜系统

## 设计范围
排行榜类型
- 1.Absolute leaderboards
  - 根据某一指标排序所有集合内成员数据，然后展示集合内 TopK 成员数据。
- 2.Relative leaderboards
  - 根据数据的不同方面对成员进行排名，根据更狭窄或相对的标准将成员分组，这可能需要进行复杂的计算。例如，一个常见的游戏场景是显示一个特定竞争者的排名以及在他们前后的竞争者的视图。

## 整体设计
数据存储方案
- 1.传统RDMS
  - 写数据：数据表存所有需排序数据
  - 读数据：建立索引，然后根据索引列进行排序查询
- 2.NoSQL Redis
  - 直接写数据到 zset
  - 从 zset 直接读根据 scores 排序后的数据，当 score 一样时按 member 的字符序进行排序
- 3.混合存储
  - 写数据，数据持久化到 RDMS，需要排序的时候再写到 NoSQL，而且只写一小部分数据到 NoSQL
  - 直接从 NoSQL 中读数据

## 参考
- [Redis Gaming Leaderboards](https://redis.com/solutions/use-cases/leaderboards/)
- https://segmentfault.com/a/1190000039320528
