# Go Modules

## Tool dependencies
1.24 tool dependencies
- https://go.dev/doc/modules/managing-dependencies#tools
- https://www.alexedwards.net/blog/how-to-manage-tool-dependencies-in-go-1.24-plus


## 实用命令
remove the entire module download cache
- `go clean --modcache`


## 配置 Go Module Go Proxy

默认 Go Module GO111MODULE=auto

修改 go env -w GO111MODULE=on

默认 Go Module GOPROXY="[https://proxy.golang.org](https://proxy.golang.org/),direct"

修改 go env -w GOPROXY="[https://goproxy.io](https://goproxy.io/),direct"

direct 表示当代理下载包出现404/410错误时，会回源到模块版本的源地址去下载

### GOSUMDB
是一个Go Checksum Database，用于在拉取模块版本时(无论是从源站拉取还是通 过Go Module Proxy拉取)保证拉取的模块版本数据未经过篡改，若发现不一致，则表示可能存 在篡改，将会立即中止。

GOSUMDB的默认值为sum.golang.org，在国内通常无法访问，但是GOSUMDB可以被Go 模块代理所代理。

因此可以通过设置GOPROXY来解决，而先前所设置的模块代理goproxy.cn就能支持代理 sum.golang.org，所以这一个问题在设置GOPROXY后。

另外若对GOSUMDB的值有自定义需求，则可自定义，其支持的格式如下:
- `<SUMDB_NAME>+<PUBLIC_KEY>`
- `<SUMDB_NAME>+<PUBLIC_KEY><SUMDB_URL`

也可以将其设置为off，也就是禁止Go在后续操作中校验模块版本。

### 访问私有仓库

步骤

1. 配置 go mod 环境变量（proxy 和 checksum database 相关）
    - `GOPRIVATE`
    - `GONOPROXY`
    - `GONOSUMDB`

一般建议直接设置GOPRIVATE，它的值将作为GONOPROXY和GONOSUMDB的默认值，所以建议直接使用GOPRIVATE。

设置多个私有仓库
- `go env-w GOPRIVATE="git.example.com,github.com/gzhh/etcd"`

通配符设置（所有模块路径为example.com的子域名(例如:git.example.com)都将不经过 Go Module Proxy和Go Checksum Database，但不包括example.com本身。）
- `go env-w GOPRIVATE="*.example.com"`


## go mod edit
- https://tpaschalis.me/go-mod-edit/

例子

```
$ go mod edit -go=1.21

$ go mod edit [-module=modname|-go=version|-toolchain=name]
```

### go mod edit -replace
使用场景：固定用某个第三方依赖库，则可以在go.mod文件中使用replace来强制指定对应的版本号
- https://eli.thegreenplace.net/2024/locally-patching-dependencies-in-go/

### go mod vendor
- https://go.dev/ref/mod#vendoring
- https://victoriametrics.com/blog/vendoring-go-mod-vendor/


## 参考：
- https://go.dev/ref/mod#private-modules
- https://go.dev/ref/mod#environment-variables
- https://go.dev/ref/mod#private-module-privacy
- go module 使用 gitlab 私有仓库 https://h1z3y3.me/posts/go-private-git-repository/
- Import Private Go Modules From GitLab Repositories https://duythhuynh.medium.com/import-private-go-modules-from-gitlab-repositories-8933fcd79c79
- gitlab 子模块 https://seankhliao.com/blog/12021-04-29-go-private-modules-in-gitlab/
- Gitlab Subgroups and Go Modules Names https://penkovski.com/post/gitlab-subgroups-go-modules/
- https://stackoverflow.com/questions/67924850/how-to-load-modules-from-gitlab-subgroup
