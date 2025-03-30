# URL Shortener
URL 短链服务

## 数据量大致预估
- 每天需要生成多少短链
- 每秒需要生成多少短链
- 短链多长时间有效
- 原始 URL 平均长度
- 最后预估需要的存储空间


## 整体设计
### API 入口
1. 生成 URL 短链
  - POST api/v1/data/shorten
  - request parameter: {longUrl: $longURLString}
  - return shortURL
2. URL 短链重定向到原始 URL
  - GET api/v1/${shortUrl}
  - return longURL for HTTP redirection

### 生成 URL 短链
必须满足条件
- 每个 longURL 必须被哈希为 hashValue
- 每个 hashValue 可以被映射回 longURL


## 生成 URL 短链和 URL 短链重定向的细节
### 数据模型
需要存储原始 URL 和短链 URL 的对应关系，因为估算的数据量可能很大，不能存储数据在内存 hash table 中，所以数据模型应该选择关系型数据库

### 使用 hash 函数来生成短链
缺点
- hash 函数生成的 hashValue 长度和我们最终需要的短链 hashValue 有较大差别，虽然可以使用通过取固定长度字符解决，但是这可能会带来哈希冲撞，需要额外的开销来解决哈希冲撞

### 使用 Base 62 conversion 方法生成短链
原理
- Base 62 算法可以将纯数字的 Unique ID 转为只包含大小写字母和数字的 62 个字符，可以理解为将 10 进制数字转位 62 进制，从而使 Unique ID 表示的长度变短

整体流程
- 根据客户端输入的 longURL 判断是否在数据库中存在
- 如果存在，意味着这个 longURL 之前已经被转化成 shortURL，此时直接返回 shortURL 即可
- 如果不存在，意味着这个 longURL 没有被处理过，需要使用 UID generator 生成一个 Unique ID
- 使用 Base 62 conversion 方法将 Unique ID 转化成 shortURL
- 在数据库中插入一条包含 Unique ID、shortURL、longURL 的新记录，

### URL 重定向
即将客户端输入的短链 URL 重定向为访问原始 URL
- 用户访问 shortURL
- 负载均衡器将请求发送到指定 web server
- 如果 shortURL 在缓存中存在，则直接返回对应 longURL
- 如果 shortURL 不在缓存中存在，则先去数据库中查询对应的 longURL，如果在数据库存在，则直接将 longURL 返回给客户端，如果在数据库不存在，那么则输入 shortURL 无效

## 总结
额外需要考虑的点
- 短链生成接口和短链重定向接口需要考虑限流
- web server 需要支持水平扩容
- database 也需要支持水平扩容和数据分片
- 服务的可用性、一致性、可靠性及可观测性也需要考虑

## 参考
- [System Design Interview](https://book.douban.com/subject/35246417/)
- Base62: https://en.wikipedia.org/wiki/Base62
- 最佳实践
  - https://github.com/dubinc/dub
  - https://github.com/ccbikai/Sink
  - https://github.com/yourselfhosted/slash
