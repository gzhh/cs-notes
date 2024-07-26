# Redis
- https://redis.io/documentation

架构设计
- Redis基础知识典藏版：架构设计 https://mp.weixin.qq.com/s/H6cGiahUreVFh34wbkFF8g


## 版本特性
- Redis4.0、5.0、6.0、7.0特性整理(持续更新) https://blog.csdn.net/zzhongcy/article/details/123903799
- https://help.aliyun.com/zh/redis/support/new-features-of-apsaradb-for-redis-7-0


## 使用场景
使用场景1
1. 缓存系统
2. 排行榜
3. 计数器
4. 社交网络（粉丝数、订阅数）
5. 消息队列系统
6. 实时系统

使用场景2
1. 会话
    
    替换高流量网站登陆时，cookie与关系型数据库的交互，使用cookie直接与redis交换来提高验证时修改数据的速度（实现会话cookie和记录用户活动）
    
2. 使用 redis 实现购物车
    
    hash结构能很方便的实现商品数据的增删改（每个用户的购物车）
    
3. 网页缓存
    
    针对页面内容长时间不发生变化的页面
    
    - 提高用户体验
    - 加速网站响应
    - 减少服务器数量
    
    缓存策略
    
    如果有缓存，就直接读取；如果没有，生成缓存，设置过期时间，再返回。
    
4. 数据行缓存（内容变化不大时）
    
    倒数促销活动时，只改变商品数量
    
5. 缓存部分高访问的页面

## 数据类型
- https://redis.io/docs/latest/develop/data-types/

基本数据结构使用场景
- Int
  - Count
  - Rate Limiter
  - Global ID
- String
  - Cache
  - Session
  - Distributed Lock
- List
  - MessageQueue
- Hash
  - Shoping Cart/Detail Info
- Set
  - User unique set
- ZSet
  - Leaderboard

附加数据结构
- Bitmaps
  - Bloom filter
  - Active user statistics
- Streams
  - MessageQueue

## 额外功能
- Best Practices https://www.jetbrains.com/guide/go/tutorials/getting_started_with_redis_go/best_practices/

发布订阅

事务

Lua脚本

Pipeline
