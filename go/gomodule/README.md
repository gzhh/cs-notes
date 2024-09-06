# Go Modules


## 实用命令
remove the entire module download cache
- `go clean --modcache`


## 配置 Go Module Go Proxy

默认 Go Module GO111MODULE=auto

修改 go env -w GO111MODULE=on

默认 Go Module GOPROXY="[https://proxy.golang.org](https://proxy.golang.org/),direct"

修改 go env -w GOPROXY="[https://goproxy.io](https://goproxy.io/),direct"

### 访问私有仓库

步骤

1. 配置 go mod 环境变量（proxy 和 checksum database 相关）
    - `GOPRIVATE`
    - `GONOPROXY`
    - `GONOSUMDB`


## go mod edit
- https://tpaschalis.me/go-mod-edit/

例子

```
$ go mod edit -go=1.21

$ go mod edit [-module=modname|-go=version|-toolchain=name]
```

### go mod edit -replace
- https://eli.thegreenplace.net/2024/locally-patching-dependencies-in-go/


## 参考：
- https://go.dev/ref/mod#private-modules
- https://go.dev/ref/mod#environment-variables
- https://go.dev/ref/mod#private-module-privacy
- go module 使用 gitlab 私有仓库 https://h1z3y3.me/posts/go-private-git-repository/
- Import Private Go Modules From GitLab Repositories https://duythhuynh.medium.com/import-private-go-modules-from-gitlab-repositories-8933fcd79c79
- gitlab 子模块 https://seankhliao.com/blog/12021-04-29-go-private-modules-in-gitlab/
- Gitlab Subgroups and Go Modules Names https://penkovski.com/post/gitlab-subgroups-go-modules/
- https://stackoverflow.com/questions/67924850/how-to-load-modules-from-gitlab-subgroup
