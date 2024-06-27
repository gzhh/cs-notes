# Cache
- https://en.wikipedia.org/wiki/Cache_(computing)


## 缓存最佳实践
- 如何优雅的设计和使用缓存？https://juejin.cn/post/6844903665845665805
- 美团万亿级 KV 存储架构与实践 https://tech.meituan.com/2020/07/01/kv-squirrel-cellar.html
- go-zero微服务实战系列（七、请求量这么高该如何优化）https://mp.weixin.qq.com/s/pPPSPZJispmITY9Wsi7hUg


## 缓存策略
### 缓存一致性
- 浅谈缓存最终一致性的解决方案 https://mp.weixin.qq.com/s/Y9S89MT0uAobzRKgYVrI9Q
- 干货 | 分布式缓存与DB秒级一致设计实践 https://mp.weixin.qq.com/s/_RlOAvjfTINJrhwHUogs7Q
- 干货 | 携程最终一致和强一致性缓存实践 https://mp.weixin.qq.com/s/E-chAZyHtaZOdA19mW59-Q
- 美团面试题：缓存一致性，我是这么回答的！https://mp.weixin.qq.com/s/esXWVZvgf74DPeDL7xbi1Q
- 并发环境下，先操作数据库还是先操作缓存？https://mp.weixin.qq.com/s/nCOaSbwgERAZDofWgRlyKg
- go-zero微服务实战系列（六、缓存一致性保证）https://mp.weixin.qq.com/s/422ZHs81y7nN9Sgb_ESsgg


## 缓存问题
介绍
- 缓存雪崩、击穿、穿透 https://xiaolincoding.com/redis/cluster/cache_problem.html
- 缓存雪崩、缓存穿透、缓存更新了解多少？https://segmentfault.com/a/1190000017882763
- go-zero微服务实战系列（五、缓存代码怎么写）https://mp.weixin.qq.com/s/QqrLOq7DcDVuIM_1YAaVTw

### 缓存雪崩

### 缓存击穿
解决方案
- [singleflight](https://pkg.go.dev/golang.org/x/sync/singleflight)
- [singleflight usage](https://www.liwenzhou.com/posts/Go/singleflight/)
- Go并发编程 —— I/O聚合优化（动画讲解）https://mp.weixin.qq.com/s/DBjO5jKnQHy0m2gIWVdgdw

### 缓存穿透
- [Bloom Filter](https://zh.wikipedia.org/zh-hans/布隆过滤器)
- [布隆过滤器使用](https://juejin.cn/post/7038779056996745224)
- [经典论文解读——布隆过滤器](https://mp.weixin.qq.com/s/IWq0GHbHspAwIuQJ9epCMA)
- https://pages.cs.wisc.edu/~cao/papers/summary-cache/node8.html


