# Web crawler
Web 爬虫

爬虫使用场景
- 搜索引擎索引（最常见）
- web 存档
- web 挖掘
- web 监控

## 设计范围
### 实现爬虫的基本逻辑
- 提供一个种子 URL 集合，下载这些 URL 对应的所有 web 页面
- 解析这些 URL 页面并获取新的 URL
- 将新的 URL 放入待下载 web 页面的 URL 列表，重复这三个步骤

### 约束范围
- 用途：为搜索引擎提供索引
- 体量：每月下载大概十亿个 web 页面
- 内容类型：只包含 HTML web 页面
- 内容变动：新的或有变动的 web 页面需要被考虑
- 存储时间：下载的 web 页面需要存储至少 5 年
- 重复内容策略：重复的 web 页面只存一份

### 大型爬虫需尽可能满足
- 可伸缩性
- 健壮性
- 遵循爬虫礼仪
- 可扩展性

### 数据量粗略估算
- Assume 1 billion web pages are downloaded every month.
- QPS: 1,000,000,000 / 30 days / 24 hours / 3600 seconds = ~400 pages per second.
- PeakQPS = 2 x QPS = 800
- Assume the average web page size is 500k.
- 1-billion-page x 500k = 500 TB storage per month.
- Assuming data are stored for five years, 500 TB * 12 months * 5 years = 30 PB. A 30 PB storage is needed to store five-year content.


## 整体设计
![Web crawler workflow](https://cdn.gzhh.tech/2022/10/web-crawler-workflow.png)

- 种子 URL
- URL 队列
- HTML 下载器
- DNS 解析器
- 网页内容分析器
- 网页内容去重
- 网页内容存储
- URL 提取器
- URL 过滤器
- URL 去重：Bloom filter 和 hash table 是常见实现方案
- URL 存储

## 细节设计
- DFS vs BFS
  - BFS 更好，因为 DFS 可能会导致递归调用栈太深
- URL 队列
  - 爬虫礼仪
    - 需控制单位时间内并发请求的数量
    - 可以每个目标网站配置一个 URL 队列，多个队列并发处理即对多个目标网站并发处理
  - 优先级
    - 可以通过配置权重来调整不同类型页面处理的优先级，具体可以参考 PageRank 算法
  - 最新的页面内容
    - 针对页面可能的更新、添加、删除等操作，我们需要周期性的下载新的页面内容
  - URL 队列的存储
    - 大部分数据 URL 数据存在磁盘，少量数据存在内存来增加读取速度
- HTML 下载器
  - robots.txt 机器人互斥协议，指定允许下载的页面
  - 性能优化
    - 分布式爬虫
    - DNS 解析缓存
    - 针对不同网站的物理位置，部署更近位置的爬虫服务
    - 减少 timeout，避免长时间等
- 健壮性
  - 使用一致性哈希来支撑分布式下载器
  - 存储爬取状态和数据日志，方便重试
  - 引入异常处理来支撑崩溃时平滑恢复
  - 数据可视化机制能有效阻止系统错误
- 可扩展性
  - 不仅局限于 web HTML 内容，需要支持 PNG 等其他类型的内容 
- 检测并跳过问题 web 页面的下载
  - 冗余内容：可以通过 hash 或 checksum 来判断是否重复
  - 爬虫陷阱：避免错误的 URL 让爬虫陷入无限循环
  - 数据噪音：过滤无价值的 web 内容，例如广告、代码段等


## 总结
额外需要优化的地方
- 支持获取服务端渲染内容：例如 JavaScript 和 AJAX 动态渲染的内容
- 过滤掉非意愿页面减少存储压力
- 引入数据复制和分片，来提升可用性、可伸缩性、可依赖性
- 水平扩展
- 可用性、一致性、可靠性
- 数据统计及分析

## 参考
- [System Design Interview](https://book.douban.com/subject/35246417/)
