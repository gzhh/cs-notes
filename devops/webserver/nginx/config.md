## Nginx配置优化

```
# 开启gzip

gzip on;

# 启用gzip压缩的最小文件，小于设置值的文件将不会压缩

gzip_min_length 1k;

# gzip 压缩级别，1-10，数字越大压缩的越好，也越占用CPU时间，后面会有详细说明

gzip_comp_level 2;

# 进行压缩的文件类型。javascript有多种形式。其中的值可以在 mime.types 文件中找到。

gzip_types text/plain application/javascript application/x-javascript text/css application/xml text/javascript application/x-httpd-php image/jpeg image/gif image/png font/ttf font/otf image/svg+xml;

# 是否在http header中添加Vary: Accept-Encoding，建议开启

gzip_vary on;

# 禁用IE 6 gzip

gzip_disable "MSIE [1-6]\.";
```

### 配置参考

```
#Specifies the value for maximum file descriptors that can be opened by this process.

worker_rlimit_nofile 65535;

events {

use epoll;

worker_connections 65535;

}

http {

server_names_hash_bucket_size 128;

client_header_buffer_size 128k;

large_client_header_buffers 4 128k;

client_max_body_size 8m;

sendfile on;

tcp_nopush on;

keepalive_timeout 60;

tcp_nodelay on;

fastcgi_connect_timeout 300;

fastcgi_send_timeout 300;

fastcgi_read_timeout 300;

fastcgi_buffer_size 64k;

fastcgi_buffers 4 64k;

fastcgi_busy_buffers_size 128k;

fastcgi_temp_file_write_size 128k;

gzip on;

gzip_min_length 1k;

gzip_buffers 4 16k;

gzip_http_version 1.0;

gzip_comp_level 2;

gzip_types text/plain application/x-javascript text/css application/xml;

gzip_vary on;

#limit_zone crawler $binary_remote_addr 10m;

}
```
