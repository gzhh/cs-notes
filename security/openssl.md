## Openssl

Wiki
- https://en.wikipedia.org/wiki/OpenSSL
- https://www.openssl.org

### 使用 openssl 生成密钥和证书
生成密钥
`openssl genrsa -out key.pem 2048`

生成自签的证书
`openssl req -new -x509 -key key.pem -out cert.pem -days 3650`

Ref
- https://luckymrwang.github.io/2021/04/06/cert-pem%E5%92%8Ckey-pem%E6%96%87%E4%BB%B6%E7%94%9F%E6%88%90/

