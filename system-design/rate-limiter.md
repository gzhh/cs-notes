# Rate Limiter 限流器

## 需要把限流器放在哪里
- 网关层
- 后端应用服务

## 限流算法
### Token Bucket 令牌桶
工作原理
- 有一个固定大小的桶，表示桶内最多可以放置固定大小的令牌。且会定时往桶内放置一定量的令牌，如果桶已经满了，则不处理，如果没满，则将令牌放入桶内。
- 每个请求会从桶中取一个令牌，如果拿到令牌则请求可以进行后续处理，如果没有拿到令牌，则请求直接被丢掉或放入消息队列等待后续处理。

令牌桶算法主要有两个参数
- 桶大小：桶内最多允许存放令牌的数量
- 令牌填充速率：单位时间（通常是秒）内放入桶内令牌的数量

### Leaky Bucket 漏桶
工作原理
- 通常使用一个 FIFO 的队列来代表桶，当有请求需要处理时，首先会检查队列是否已满；如果未满，那么请求直接添加到队列，如果已满，则请求被丢弃。请求会定时从队列里拉出来处理。

漏桶算法主要有两个参数
- 桶大小：即队列的大小，队列会持有将要被处理的请求
- 漏出速率：表示单位时间内（通常是秒）有多少请求能被处理

### 滑动窗口
工作原理
- 使用 Redis 的 ZSet，以当前时间到当前时间的前一个固定时间为一个窗口，窗口内给定一个固定正整数 limit，每次请求前检查这个数是否大于limit，如果大于则请求正常并zadd 1，否则放弃处理本次请求。
- 具体操作
  ```
  v = ZCOUNT key timeNow-conf.window timeNow
  if v > conf.limit
    return false
  else
    ZREMRANGEBYSCORE key timeNow-7200, timeNow-3600
    ZADD key timeNow, timeNow, token
    EXPIRE key 3600
    return true
  ```

## 限流整体架构
除了限流算法外，我们还需要计数器来追踪有多少个请求来自相同的用户、IP 等。

存储计数器一般都采用内存缓存技术来实现，比如 Redis，它提供了两个命令 INCR 和 EXPIRE。
- INCR：给计数器累加1
- EXPIRE：给计数器设置超时时间，如果过期，计数器会被自动删除

## 细节处理
### 限流规则
限流粒度
- 全局请求限流
- 指定 service 请求限流
- 指定 API 请求限流

限流频率控制
- 单位
- 单位时间内允许处理请求数

### 分布式系统中限流器需要考虑的问题
1. Race condition
  - 解决办法有：使用 lock，或使用 redis lua 脚本保证原子性
2. Synchronization issue
  - 解决办法有：多个限流器使用共享数据中心

### 性能
引入多数据中心减少用户访问数据的路径。

### 监控
验证限流器是否正常工作，和规则是否正常。


## 总结
避免请求被限流，客户端的最佳实践
- 使用缓存避免频繁请求 API
- 理解限流工作原理，不要短时间内频繁发送太多请求
- 客户端应该捕获异常和错误，支持从异常状态平滑恢复
- 添加有效的日志备份以便进行重试逻辑


## 参考
- [System Design Interview](https://book.douban.com/subject/35246417/)
- https://en.wikipedia.org/wiki/Rate_limiting
- https://en.wikipedia.org/wiki/Token_bucket
- https://en.wikipedia.org/wiki/Leaky_bucket
- 软件系统限流的底层原理解析 https://mp.weixin.qq.com/s/EJ68f40ebapdqKTgGOG8tw
- 网关层 Nginx 限流讨论
  - https://x.com/vikingmute/status/1853606107874464067
  - https://nginx.org/en/docs/http/ngx_http_limit_req_module.html
- Go工程化实践：基于Redis实现一个分布式限流器 https://zhuanlan.zhihu.com/p/619530958
