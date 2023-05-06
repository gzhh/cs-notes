## HTTPS证书配置

### 初始配置流程参照

[https://ksmx.me/letsencrypt-ssl-https/](https://ksmx.me/letsencrypt-ssl-https/)

### 同一台服务器添加 HTTP 新站点

1. 添加证书配置文件
    
    例如：`/etc/letsencrypt/configs/uploads.gzhh.tech.conf`
    
2. 执行证书自动化生成命令
    
    需先关闭 nginx：`systemctl stop nginx`
    
    webroot模式（适用于80端口在用且服务已经在运行的情况）
    
    - 例如：`certbot certonly -c /etc/letsencrypt/configs/uploads.gzhh.tech.conf`
    
    standalone模式（适用于80端口未启用或已启用服务没有根目录的情况）
    
    - 再执行命令，例如： `certbot certonly --standalone -d uploads.gzhh.tech`
3. 查看是否生成证书
    
    `ll /etc/letsencrypt/live/uploads.gzhh.tech/`
    
4. 调整对应站点 nginx 配置使其支持 ssl
    
    例如：`/etc/nginx/conf.d/uploads.gzhh.tech.conf`
    
5. 重启 nginx
    
    `systemctl restart nginx`
    

### certbot 相关命令

1. 证书重新生成（crontab 配置自动生成）
    
    `certbot renew --pre-hook "systemctl stop nginx" --post-hook "systemctl start nginx" --renew-hook "systemctl reload nginx" --quiet`
    
2. 查看证书续期状态
    
    `certbot certificates`
    
3. 删除证书，可选操作
    
    `certbot delete`
    

### 参考

- [https://www.linuxmi.com/lets-encrypt-certbot-ssl.html](https://www.linuxmi.com/lets-encrypt-certbot-ssl.html)
