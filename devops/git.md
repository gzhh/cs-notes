# Git
## 常见命令

### 分支处理

删除本地分支

git branch -a

git checkout master

git branch -d <BranchName>

删除远程分支

git push origin --delete <BranchName>

git branch -a

### 基本操作

remote 操作

git remote add upstream [git@github.com](mailto:git@github.com):gzhh/gzhh.github.io.git

git remote add origin [git@github.com](mailto:git@github.com):gzhh/gzhh.github.io.git

[Git Remote 操作](https://www.ruanyifeng.com/blog/2014/06/git_remote.html)

## 配置

### 用户名和邮箱

查看用户名和邮箱（当前仓库）：

git config user.name
git config user.email

设置用户名和邮箱（当前仓库）：

git config [user.name](http://user.name/) "gzhh"
git config [user.email](http://user.email) "zhonghuangao@gmail.com"

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


## Git Flow

git flow feature start target-board

git add .

git commit -m ""

git push

git flow feature finish target-board



