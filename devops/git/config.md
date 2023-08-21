## Git 配置

### 用户名和邮箱

查看用户名和邮箱（当前仓库）：

git config user.name

git config user.email

设置用户名和邮箱（当前仓库）：

git config user.name "gzhh"
git config user.email "zhonghuangao@gmail.com"

全局设置

git config --global 

配置大小写敏感

git config core.ignorecase false


### 代理

配置代理：

git config --global http.proxy proxy_addr

git config --global https.proxy proxy_addr

如果要取消代理设置，则执行

git config --global –unset http.proxy

git config --global –unset https.proxy

查看已设置的值：

git config http.proxy

git config https.proxy


