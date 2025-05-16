# Cache
- https://en.wikipedia.org/wiki/Cache_(computing)


## 缓存最佳实践
- 如何优雅的设计和使用缓存？https://juejin.cn/post/6844903665845665805
- 美团万亿级 KV 存储架构与实践 https://tech.meituan.com/2020/07/01/kv-squirrel-cellar.html
- go-zero微服务实战系列（七、请求量这么高该如何优化）https://mp.weixin.qq.com/s/pPPSPZJispmITY9Wsi7hUg
- 后台开发必备：每个程序员都应掌握的缓存技术 https://mp.weixin.qq.com/s/pu0w46EGBD6CJmozkMXkrQ


## 缓存策略
### 缓存一致性
- 一文讲透数据库缓存一致性问题 https://mp.weixin.qq.com/s/HilTpiLm5VKp5pbhd-K9uQ
- 浅谈缓存最终一致性的解决方案 https://mp.weixin.qq.com/s/Y9S89MT0uAobzRKgYVrI9Q
- 干货 | 分布式缓存与DB秒级一致设计实践 https://mp.weixin.qq.com/s/_RlOAvjfTINJrhwHUogs7Q
- 干货 | 携程最终一致和强一致性缓存实践 https://mp.weixin.qq.com/s/E-chAZyHtaZOdA19mW59-Q
- 美团面试题：缓存一致性，我是这么回答的！https://mp.weixin.qq.com/s/esXWVZvgf74DPeDL7xbi1Q
- 并发环境下，先操作数据库还是先操作缓存？https://mp.weixin.qq.com/s/nCOaSbwgERAZDofWgRlyKg
- go-zero微服务实战系列（六、缓存一致性保证）https://mp.weixin.qq.com/s/422ZHs81y7nN9Sgb_ESsgg
- 奇怪的缓存一致性问题 https://mp.weixin.qq.com/s/hcmRh493yWJUJTJl0tlZWw
- 所以延迟双删有啥用 https://www.dtstack.com/bbs/article/9477
- 缓存一致性的探讨 https://zhuanlan.zhihu.com/p/606249255

缓存策略
- 缓存更新的四种策略及选取建议 https://www.cnblogs.com/reim/p/17414244.html
- Go开发者必知：五大缓存策略详解与选型指南 :bhttps://mp.weixin.qq.com/s/aPiWDFInEXQjqosvGB-IUg
- 缓存 从策略到实践概览 https://zhuanlan.zhihu.com/p/352325821
- 不知道这四种缓存模式，敢说懂缓存吗？https://www.51cto.com/article/713035.html
- 3种常用的缓存读写策略详解 https://javaguide.cn/database/redis/3-commonly-used-cache-read-and-write-strategies.html


## 缓存问题
介绍
- 缓存雪崩、击穿、穿透 https://xiaolincoding.com/redis/cluster/cache_problem.html
- 缓存雪崩、缓存穿透、缓存更新了解多少？https://segmentfault.com/a/1190000017882763
- go-zero微服务实战系列（五、缓存代码怎么写）https://mp.weixin.qq.com/s/QqrLOq7DcDVuIM_1YAaVTw

### 缓存雪崩

### 缓存击穿
解决方案
- [singleflight](https://pkg.go.dev/golang.org/x/sync/singleflight)
  - singleflight 使用 https://www.liwenzhou.com/posts/Go/singleflight/
  - Go 并发控制：singleflight 详解 https://mp.weixin.qq.com/s/3BJObWvlzsetMHb7cSR8jg
  - 利用singleflight应对缓存击穿 https://zhuanlan.zhihu.com/p/487059758
- Go并发编程 —— I/O聚合优化（动画讲解）https://mp.weixin.qq.com/s/DBjO5jKnQHy0m2gIWVdgdw

### 缓存穿透
- [Bloom Filter](https://zh.wikipedia.org/zh-hans/布隆过滤器)
- [布隆过滤器使用](https://juejin.cn/post/7038779056996745224)
- [经典论文解读——布隆过滤器](https://mp.weixin.qq.com/s/IWq0GHbHspAwIuQJ9epCMA)
- https://pages.cs.wisc.edu/~cao/papers/summary-cache/node8.html
- https://eli.thegreenplace.net/2025/bloom-filters/


