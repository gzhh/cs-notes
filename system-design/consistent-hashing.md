# Consistennt Hashing 一致性哈希

## 最佳实践
- 9.4 什么是一致性哈希？https://xiaolincoding.com/os/8_network_system/hash.html
- 五分钟了解一致性哈希算法 https://mp.weixin.qq.com/s/-0A7ylm-M9ltURAhhvg89Q
- Consistent Hashing: An Overview and Implementation in Golang https://dev.to/permify/consistent-hashing-an-overview-and-implementation-in-golang-53je


## 介绍
### 为什么需要一致性哈希，解决了什么问题
如何均衡分配请求
- 由于单节点服务的机器配置限制，往往需要进行水平扩容，将请求分散到不同的节点进行处理
- 传统负载均衡算法，例如轮循加上加权逻辑，基本可以满足上述要求
- 但是如果是针对分布式 kv 服务，不同节点存储的是不同的数据，如果需要定位某个 key 位于哪个节点，上述的传统负载均衡无法处理

哈希算法的问题
- 接着上述如何定位 key 位于哪个节点，我们可以想到用哈希算法来定位，将 key 进行哈希处理，得到一个固定的值，然后按照 hash(key) % n 的公式确定 key 对应的 value 存在哪个节点（n 为节点数量）
- 但是如果节点数量 n 发生变化时，那么通过公式计算出 key 存储数据的节点位置就可能有变化，但是 key 对应的数据可能还在原来的节点上，这时候就需要做数据搬迁，将数据搬迁到 key 对应的新节点
- 普通的哈希算法面对节点调整时，需要搬迁的数据规模可能会很大，搬迁的成本太高


## 什么是一致性哈希，一致性哈希算法原理
### 原理
- 一致性哈希算法就很好地解决了分布式系统在节点扩容或者缩容时，发生过多的数据迁移的问题
- 一致哈希算法也用了取模运算，但与哈希算法不同的是，哈希算法是对节点的数量进行取模运算，而一致哈希算法是对 2^32 进行取模运算，是一个固定的值
- 一致性哈希是指将「存储节点」和「数据」都映射到一个首尾相连的哈希环上。映射的结果值往顺时针的方向的找到第一个节点，就是存储该数据的节点。
- 在一致哈希算法中，如果增加或者移除一个节点，仅影响该节点在哈希环上顺时针相邻的后继节点，其它数据也不会受到影响。

### 一致性哈希算法的问题
- 一致性哈希算法并不保证节点能够在哈希环上分布均匀，这样就会带来一个问题，会有大量的请求集中在一个节点上
- 不再将真实节点映射到哈希环上，而是将虚拟节点映射到哈希环上，并将虚拟节点映射到实际节点，所以这里有「两层」映射关系
- 带虚拟节点的一致性哈希方法不仅适合硬件配置不同的节点的场景，而且适合节点规模会发生变化的场景

### 一致性哈希的优点
- 优点：节点扩容和缩容的时候，可以保证减少数据搬迁
- 扩展：结合虚拟节点，


## 参考
- [System Design Interview](https://book.douban.com/subject/35246417/)
- https://en.wikipedia.org/wiki/Load_balancing_(computing)
- https://en.wikipedia.org/wiki/Hash_function
- https://en.wikipedia.org/wiki/Consistent_hashing
- https://xiaolincoding.com/os/8_network_system/hash.html
