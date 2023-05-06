## HTTP version

### 介绍

HTTP 协议升级一般是，浏览器先向上支持更新的协议，然后服务器再实现

### HTTP服务

HTTP1.1

HTTP1.1 + SSL

HTTP2.0 + SSL

### HTTP2

标准化 RFC

### HTTP1 问题总结

1. 连接耗时
2. 拥塞控制耗时
3. HTTP 队头问题

### HTTP2 解决 HTTP1.1 不好的地方

单个连接 → 支持多个链接

默认6个并发连接 → HTTP2 连接多路复用，默认100个请求（解决http1队头堵塞问题）

[HTTP1 和 HTTP2 区别](images/http-version-diff.jpg)


### HTTP2 部署实践

1. 浏览器(一般各种协议都支持) + 代理服务器(HTTP2) + 后端服务器(HTTP1)
    
    浏览器和代理服务器之间用 HTTP2 通信
    
    代理服务器和后端服务器之间用HTTP1通信
    
    web服务器软件针对HTTP2的实现中，Apache可以用作代理和后端服务器，Nginx只能作反向代理服务器
    

### HTTP2 帧

HEADER 帧

DATA 帧


### HTTP2 问题

**TCP拥塞窗口**

慢启动，不断调整，不断自动调到最佳窗口（即使指数增长）

**丢包**影响

小概率的网络差情况下（实验证明超过2%丢包），丢包严重可能会导致 HTTP2 比 HTTP1 性能差

由于使用单个 TCP 连接，丢包导致 TCP 拥塞窗口减小，数据帧传输速度也会变慢

### QUIC 或 HTTP3

HTTP1 + TCP

|

HTTP2 + TCP (HTTP2 基于 SDPY - Google线上验证过)

|

QUIC + UDP (QUIC基于HTTP2 - Google线上验证)
