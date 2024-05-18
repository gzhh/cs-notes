# HTTP
- https://en.wikipedia.org/wiki/HTTP
- https://developer.mozilla.org/en-US/docs/Web/HTTP


## Protocol
HTTP1.1
- https://datatracker.ietf.org/doc/html/rfc9112

HTTPS
- The Transport Layer Security (TLS) Protocol Version 1.3 https://datatracker.ietf.org/doc/html/rfc8446
- The Transport Layer Security (TLS) Protocol Version 1.2 https://datatracker.ietf.org/doc/html/rfc5246
- https://en.wikipedia.org/wiki/HTTPS
- https://en.wikipedia.org/wiki/Transport_Layer_Security
- https://tls13.xargs.org
- 有关 TLS/SSL 证书的一切 https://www.kawabangga.com/posts/5330

HTTP2
- https://datatracker.ietf.org/doc/html/rfc9113
- https://en.wikipedia.org/wiki/HTTP/2
- https://http2.github.io
- Does HTTP/2 make websockets obsolete? https://stackoverflow.com/questions/28582935/does-http-2-make-websockets-obsolete

HTTP3
- https://datatracker.ietf.org/doc/html/rfc9114
- https://datatracker.ietf.org/doc/html/rfc9000
- https://en.wikipedia.org/wiki/HTTP/3
- https://en.wikipedia.org/wiki/QUIC
- https://quic.xargs.org
- https://cangsdarm.github.io/illustrate/quic
- HTTP/3 From A To Z: Core Concepts https://www.smashingmagazine.com/2021/08/http3-core-concepts-part1/
- quic-go https://github.com/quic-go/quic-go


## 原理
- 了解 HTTP 看这一篇就够 https://mp.weixin.qq.com/s/JVQPy8hAVoq1pRq63HKBOw


## Content-Type
- https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type
- https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Basics_of_HTTP/MIME_types
- Do I need Content-Type: application/octet-stream for file download? https://stackoverflow.com/questions/20508788/do-i-need-content-type-application-octet-stream-for-file-download

### 请求报文

请求行 request line（请求方法、请求url、HTTP协议版本）

首部 headers

空行 blank line

请求体 request body

### 响应报文

状态行 status line（HTTP协议版本、状态码、状态码的文本描述）

响应头 header

空行 blank line

响应体 response body

### HTTP状态码
- https://en.wikipedia.org/wiki/List_of_HTTP_status_codes
- https://developer.mozilla.org/en-US/docs/Web/HTTP/Status

200 OK

301 Moved Permanently

302 Found（Moved Temporarily）

304 Not Modified

307 Temporary Redirect

400 Bad Request

401 Unauthorized

403 Forbidden

404 Not Found

405 Method Not Allowed

408 Request Timeout

500 Internal Server Error

501 Not Implemented

502 Bad Gateway

503 Service Unavailable

504 Gateway Timeout

### HTTP2对比HTTP1.1的优化

减少网络延迟，提高浏览器的页面加载速度

- 对[HTTP头字段](https://zh.wikipedia.org/wiki/HTTP%E5%A4%B4%E5%AD%97%E6%AE%B5)进行[数据压缩](https://zh.wikipedia.org/wiki/%E6%95%B0%E6%8D%AE%E5%8E%8B%E7%BC%A9)(即HPACK算法)；
- HTTP/2 服务端推送(Server Push)；
- 请求[流水线](https://zh.wikipedia.org/wiki/HTTP%E7%AE%A1%E7%B7%9A%E5%8C%96)；
- 修复HTTP/1.0版本以来未修复的 [队头阻塞](https://zh.wikipedia.org/wiki/%E9%98%9F%E5%A4%B4%E9%98%BB%E5%A1%9E) 问题；[https://zhuanlan.zhihu.com/p/330300133](https://zhuanlan.zhihu.com/p/330300133)
- 对数据传输采用[多路复用](https://zh.wikipedia.org/wiki/%E5%A4%9A%E8%B7%AF%E5%A4%8D%E7%94%A8)，让多个请求合并在同一 [TCP](https://zh.wikipedia.org/wiki/TCP) 连接内。

### Ref

[https://zh.wikipedia.org/wiki/超文本传输协议](https://zh.wikipedia.org/wiki/%E8%B6%85%E6%96%87%E6%9C%AC%E4%BC%A0%E8%BE%93%E5%8D%8F%E8%AE%AE)

[https://zh.wikipedia.org/wiki/HTTP状态码](https://zh.wikipedia.org/wiki/HTTP%E7%8A%B6%E6%80%81%E7%A0%81)

[https://zh.wikipedia.org/wiki/HTTP/2](https://zh.wikipedia.org/wiki/HTTP/2)

[https://zh.wikipedia.org/wiki/HTTP头字段](https://zh.wikipedia.org/wiki/HTTP%E5%A4%B4%E5%AD%97%E6%AE%B5)

[https://zh.wikipedia.org/wiki/表现层状态转换](https://zh.wikipedia.org/wiki/%E8%A1%A8%E7%8E%B0%E5%B1%82%E7%8A%B6%E6%80%81%E8%BD%AC%E6%8D%A2)

HTTP请求响应报文 [https://www.cnblogs.com/myseries/p/11239662.html](https://www.cnblogs.com/myseries/p/11239662.html)

HTTP1.0、HTTP1.1 和 HTTP2.0 的区别 [https://www.cnblogs.com/heluan/p/8620312.html](https://www.cnblogs.com/heluan/p/8620312.html)
