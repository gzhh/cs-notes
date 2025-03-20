# Caddy
Fast and extensible multi-platform HTTP/1-2-3 web server with automatic HTTPS
- https://github.com/caddyserver/caddy
- https://caddyserver.com


## 配置本地 https
```
如果你想用 Caddy 将自定义端口和本地域名代理到一个后端的 HTTP 接口，可以通过 reverse_proxy 指令实现。以下是具体步骤：

假设场景
自定义域名：mydomain.local
自定义端口：8443
后端 HTTP 接口：http://127.0.0.1:8080（假设这是你的后端服务地址）

1. 编辑本地 hosts 文件
确保自定义域名映射到本地，如前所述：

编辑 /etc/hosts（Linux/macOS）或 C:\Windows\System32\drivers\etc\hosts（Windows），添加：

127.0.0.1 mydomain.local


2. 创建 Caddyfile
在工作目录下创建 Caddyfile，配置代理到后端 HTTP 接口：

mydomain.local:8443 {
  tls internal
  reverse_proxy http://127.0.0.1:8080
}

mydomain.local:8443: Caddy 监听 mydomain.local 的 8443 端口，提供 HTTPS 服务。
tls internal: 使用自签名证书（本地测试用）。
reverse_proxy http://127.0.0.1:8080: 将所有请求代理到后端的 HTTP 服务（运行在 8080 端口）。


3. 启动后端服务（如果尚未运行）
确保你的后端 HTTP 接口已经在 http://127.0.0.1:8080 上运行。例如，如果你用 Python 跑一个简单服务：
python -m http.server 8080
这会启动一个简单的 HTTP 文件服务器，供测试用。


4. 运行 Caddy
在 Caddyfile 所在目录下运行：
caddy run


5. 访问代理
在浏览器中输入：
https://mydomain.local:8443
Caddy 会将请求代理到 http://127.0.0.1:8080，你应该能看到后端服务的响应。由于是自签名证书，浏览器可能提示“证书不受信任”，选择信任即可。


6. 可选：全局端口配置
如果你想统一设置 HTTPS 端口，可以用全局配置：

{
  https_port 8443
}

mydomain.local:8443 {
  tls internal
  reverse_proxy http://127.0.0.1:8080
}


7. 高级配置（可选）
路径匹配: 如果只想代理特定路径：

mydomain.local:8443 {
  tls internal
  route /api/* {
    reverse_proxy http://127.0.0.1:8080
  }
  file_server
  root * /path/to/your/web/files
}
这里 /api/* 的请求会代理到后端，其他请求仍提供本地静态文件。
负载均衡: 如果有多个后端：

mydomain.local:8443 {
  tls internal
  reverse_proxy {
    to http://127.0.0.1:8080 http://127.0.0.1:8081
    lb_policy round_robin
  }
}
注意事项
后端服务: 确保 http://127.0.0.1:8080 已启动，否则 Caddy 会返回 502 错误。
端口冲突: 确认 8443 端口未被占用。
日志调试: 如果有问题，运行 caddy run --watch 查看详细日志。
这样，你就成功用 Caddy 将 https://mydomain.local:8443 代理到一个 HTTP 接口了！
```
