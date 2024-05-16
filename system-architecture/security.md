# 安全

## 加密
- https://zh.wikipedia.org/wiki/加密
- https://zh.wikipedia.org/wiki/散列函数
- https://zh.wikipedia.org/wiki/RSA加密算法
- https://zh.wikipedia.org/wiki/公开密钥加密
- https://zh.wikipedia.org/zh-cn/数字签名
- 公钥，私钥和数字签名这样最好理解 https://blog.csdn.net/21aspnet/article/details/7249401
- [Linux] 使用openssl实现RSA非对称加密 https://www.cnblogs.com/taoshihan/p/6340854.html
- What is a Digital Signature? http://youdzone.com/signature.html
- 数字签名是什么？https://www.ruanyifeng.com/blog/2011/08/what_is_a_digital_signature.html
- RSA算法原理（一）https://www.ruanyifeng.com/blog/2013/06/rsa_algorithm_part_one.html

OpenSSL
- https://www.openssl.org
- https://github.com/openssl/openssl

SSH
- https://en.wikipedia.org/wiki/Secure_Shell
- https://www.cloudflare.com/learning/access-management/what-is-ssh/


## 服务端应用安全

### 一、SQL注入（injection）

用预处理防护

### 二、文件上传漏洞

cig_pathinfo=1，会递归检查路径中的 scriptname

### 三、认证与会话管理

cookie

session

### 四、访问控制（OAuth）

OAuth2.0协议

### 五、加密算法与随机数

### 六、web框架安全

### 七、应用层拒绝服务攻击

 DDOS：资源耗尽，假大量 IP 请求

### 其他

1. 防盗链技术（别人网站盗用本站图片的链接）
    
    可以通过 HTTP 的请求头中的 REFERER 来判断访问图片的域名是否是本站
