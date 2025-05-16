# Go DevOps

## Golang编译
### 编译
- 编译 Go 程序时，编译器只会关注那些直接被引用的库
- 跨平台交叉编译
    
    Golang 支持交叉编译，在一个平台上生成另一个平台的可执行程序
    
    Mac 下编译 Linux 64位可执行程序
    
    ```
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
    ```
    
    Linux 下编译 Mac 64位可执行程序
    
    ```
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
    ```
    
    GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
    
    GOARCH：目标平台的体系架构（386、amd64、arm）
    
    交叉编译不支持 CGO 所以要禁用它

### 编译优化
参考：https://blog.abnerzhao.com/posts/go-build/

参数
- `-ldflags "-w -s"` 选项目的是获取最小的二进制文件
- `-gcflags "-N -l"` 选项目的是在编译过程中禁止内联优化，加快编译速度减少开销

示例
- `CGO_ENABLED=0 go build -gcflags "-N -l" -ldflags="-s -w" -v -a -o target **.go`


## Graceful Restart
- Go实现进程优雅退出 https://zhuanlan.zhihu.com/p/571016778
- Go 程序如何实现优雅退出？来看看 K8s 是怎么做的 https://mp.weixin.qq.com/s/DiYPP0wPC8-yj94yK4yc4w
- Go 如何实现热重启 https://mp.weixin.qq.com/s/UVZKFmv8p4ghm8ICdz85wQ
- Go 平滑重启 https://mp.weixin.qq.com/s/LcYdLEkJc1qMCn-LESzfkw
- 请问你们是如何实现服务不中断的代码发布更新的 https://v2ex.com/t/922911
- 优雅的重启服务 https://eddycjy.com/posts/go/gin/2018-03-15-reload-http/
- endless 如何实现不停机重启 Go 程序？https://www.luozhiyun.com/archives/584
- Graceful restart or stop https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
- gRPC的平滑关闭和在Kubernetes上的服务摘流方案总结 https://mp.weixin.qq.com/s/lCTyZgSK3b-rJtV9l6PNYA
- Graceful Shutdown in Go: Practical Patterns https://victoriametrics.com/blog/go-graceful-shutdown/index.html


## Trace ID
- Go 语言如果没有 ctx 传递，如何让 Trace ID 连贯传递呢？ - V2EX https://v2ex.com/t/968060#reply17


## CICD
- Automatically https://github.com/uber-go/automaxprocs
