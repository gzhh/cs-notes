# map
- https://go.dev/blog/maps
- https://go.dev/blog/swisstable


## 使用
- https://pkg.go.dev/maps
- https://pkg.go.dev/golang.org/x/exp/maps


## 原理
- https://en.wikipedia.org/wiki/Hash_table
  - 哈希表 https://www.hello-algo.com/chapter_hashing/hash_map
- https://en.wikipedia.org/wiki/Hash_collision
  - 哈希冲突 https://www.hello-algo.com/chapter_hashing/hash_collision/
- https://en.wikipedia.org/wiki/Hash_function
  - 哈希函数 https://www.hello-algo.com/chapter_hashing/hash_algorithm

Hash table
- https://go.dev/src/runtime/map.go
- 机智！生抠 map的哈希函数 https://colobu.com/2022/12/21/use-the-builtin-map-hasher/

Swiss Tables
- https://go.dev/src/runtime/map_swiss.go
- Go 1.24 - a new builtin map implementation based on Swiss Tables
  - more efficient memory allocation of small objects, and a new runtime-internal mutex implementation.
  - https://abseil.io/about/design/swisstables
- Swiss Tables https://news.ycombinator.com/item?id=43086437
- 替换标准库的map实现，SwissTable更快？https://colobu.com/2023/06/29/replace-std-map-faster/
- 简单了解下最近正火的SwissTable https://www.cnblogs.com/apocelipes/p/17562468.html
- SwissTable：高性能哈希表实现 https://juejin.cn/post/7443083267282239524

深入理解
- 解剖Go语言map底层实现 https://i6448038.github.io/2018/08/26/map-secret/
- 从底层到应用，想深入Map这篇文章千万不要错过！https://mp.weixin.qq.com/s/BKcRiDMc0pLbZ822LcRPuA
- Go Maps Explained: How Key-Value Pairs Are Actually Stored https://victoriametrics.com/blog/go-map/
