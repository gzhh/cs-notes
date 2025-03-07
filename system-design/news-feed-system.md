# News feed system
信息流系统

## 设计范围
- 支持范围
  - 移动 app 和 web app
- 主要功能
  - 用户能发布动态，并且在信息流页面能看到朋友发布的动态
- 动态排序规则
  - 简单的按发布时间倒序排序即可
- 一个用户可以有朋友的数量
  - 5000 左右
- 应用体量
  - 日活千万
- 信息流包含的内容
  - 文本、图片和视频


## 整体设计
### 信息流系统主要 APIs
- 动态发布 API
  - 接收 HTTP POST 请求来发布动态
- 信息流生成 API
  - 接收 HTTP GET 请求来生成一个信息流列表

### 动态发布
![Feed publishing](https://cdn.gzhh.tech/2022/10/news-feed-system-feed-publishing.png)

- 用户：通过移动端或 web 端应用调用 API 发布动态
- 负载均衡器：将流量分发到具体服务器
- web 服务器：重定向流量到不同的内部子服务
- 发布服务：持久化数据到 db 和 cache
- 扇出服务：推送新的动态到朋友的信息流列表，为了快速生成信息流列表将信息流数据存在 cache
- 提醒服务：当新的动态发布成功的时候，通知朋友

### 信息流生成
![Newsfeed building](https://cdn.gzhh.tech/2022/10/news-feed-system-newsfeed-building.png)

- 用户：调用 API 发送一个刷新信息流列表的请求
- 负载均衡器：将流量分发到具体的服务器
- web 服务器：路由请求到 newsfeed 子服务
- newsfeed 服务：从 cache 获取信息流列表
- newsfeed 缓存：存储 news feed IDs 相关数据


## 深入细节
### 动态发布细节设计
![Feed publishing detail](https://cdn.gzhh.tech/2022/10/news-feed-system-feed-publishing-detail.png)

- Web servers
  - authentication：只有认证用户才能发布动态
  - rate limiting：限制一个用户单位时间内可以发布动态的数量
- Fanout service
  - 作用是将动态同步给所有朋友
  - Fanout on write（推模式）
    - workflow：新动态发布时，立刻将动态同步到朋友的缓存
    - 优点：
      - 动态生能的同时就能立刻同步到朋友
      - 因为信息流数据是在动态发布的时候就写到缓存，所以获取信息流列表很快
    - 缺点：
      - 如果用户的朋友有很多，那么获取朋友列表和为他们生成信息流的操作就很慢，这也叫做 hotkey 问题
      - 对于非活跃用户来说会浪费计算资源
  - Fanout on read（拉模式）
    - workflow：当用户刷新信息流页面时，才主动生成信息流，也叫做按需模型
    - 优点：
      - 对于非活跃用户来说，不会在他们上面浪费计算资源
      - 数据不会推送给好友，因此不会有 hotkey 问题
    - 缺点：
      - 因为信息流数据没有被预处理，所以刷新信息流页面的时候响应就比较慢
  - 混合方式：
    - 结合推模式和拉模式的优点
    - 对于大多数用户，我们采用推模式，保证获取信息流列表的速度
    - 对于有大量朋友或 follower 的用户，采用拉模式，避免浪费过多的计算资源在非活跃用户上
- Fanout service workflow
  - 从图数据库获取朋友 IDs
  - 从用户缓存获取用户信息数据，并按配置做过滤
  - 发送朋友列表和动态 ID 到 MQ
  - Fanout workers 从 MQ 获取数据，并将信息流数据存在信息流缓存上（缓存数据需尽可能小）

### 信息流生成设计
![Newsfeed building detail](https://cdn.gzhh.tech/2022/10/news-feed-system-newsfeed-building-detail.png)

- 媒体素材数据存储在 CDN
- workflow
  - 用户发送一个刷新信息流列表的请求
  - 负载均衡器将流量分发到具体的 web 服务器
  - web 服务器调用 news feed 服务来获取信息流列表
  - 信息流服务从信息流缓存获取一个动态 IDs 列表
  - 用户的信息流数据就是动态 IDs 列表的扩展，需要从缓存获取完整的用户及动态数据，包括用户名、头像、动态内容、动态图片等附加信息
  - 完整的信息流数据以 json 的形式返回给客户端用来渲染页面


## 总结：有待优化的点
可扩展性
- 垂直和水平扩容
- SQL 和 NoSQL
- 主从复制
- 只读副本
- 一致性模型
- 数据库分片

其他方面
- 保持各个 web 服务器的无状态性
- 尽可能的缓存数据
- 支持多数据中心
- 使用 MQ 来解耦各个组件
- 监控关键指标，例如峰值 QPS 和用户刷新信息流页面的时延


## 参考
- [System Design Interview](https://book.douban.com/subject/35246417/)
- [Feed 流系统的架构设计方案](https://mp.weixin.qq.com/s/CyXO13C3zQLoa-p7N2qPMg)
- 腾讯频道Feed流系统架构设计 https://mp.weixin.qq.com/s/XVgdfRz3fkuEM0wFCvljqA
