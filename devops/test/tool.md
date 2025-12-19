# Test tool

## Traffic replay tool
GoReplay vs TCPCopy
- GoReplay https://github.com/buger/goreplay
  - 实时或离线复制线上 HTTP 流量，并发送到测试环境，用于无风险验证与压测。
  - 可以用它做：
    - 新版本验证（shadow test）
    - 性能测试（倍增）
    - 回归测试（离线）
    - 线上问题复现
    - 流量脱敏 + 模拟测试
  - 无需改线上代码、无侵入、易部署。
- TCPCopy https://github.com/session-replay-tools/tcpcopy
  - 抓取线上真实 TCP 请求 → 重放至测试服务 → 不影响线上
  - 适合：
    - 新版本验证
    - 压测
    - 线上故障排查
    - 功能回归测试
  - 是非常经典、稳定、成熟的流量回放方案。

| 项目             | GoReplay    | TCPCopy                    |
| -------------- | ----------- | -------------------------- |
| 层级             | **HTTP 层**  | **TCP 层**                  |
| 协议支持           | HTTP/HTTPS  | 任意 TCP（如 MySQL、Redis、HTTP） |
| 是否需要 intercept | ❌ 不需要       | ✔ 需要                       |
| 使用难度           | 简单          | 较复杂                        |
| 是否回放响应         | ❌ 不回传给用户    | ❌ 不回传给用户                   |
| 能否倍增压力         | ✔ 支持        | ✔ 支持                       |
| 支持流量修改         | ✔ 强         | 弱                          |
| 是否跨平台友好        | ✔ Go 编写，很方便 | 中等                         |
| 使用场景           | Web 服务验证与回放 | 需要模拟真实 TCP 状态的服务           |

一句话总结：
- GoReplay = HTTP 流量复制/回放的首选工具
- TCPCopy = 任意 TCP 协议的流量复制工具
