<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [The Go Programming Language](#the-go-programming-language)
  - [一、Go Version Management](#%E4%B8%80go-version-management)
  - [二、Go Package Management](#%E4%BA%8Cgo-package-management)
    - [Go Modules](#go-modules)
    - [Go multi-module workspaces](#go-multi-module-workspaces)
    - [Go Modules Services](#go-modules-services)
  - [三、Best Practice](#%E4%B8%89best-practice)
    - [gonew](#gonew)
    - [Standard](#standard)
  - [四、Go Community](#%E5%9B%9Bgo-community)
    - [Go Open-Source Projects](#go-open-source-projects)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# The Go Programming Language
- https://github.com/golang
- https://github.com/golang/go
- https://go.dev
- Go Packages https://pkg.go.dev

Third Party Organization
- Uber Go https://github.com/uber-go


## 一、Go Version Management
Versions Info
- All releases https://go.dev/dl/
- blog: Go 1.18 is released! https://go.dev/blog/go1.18
- blog: Go 1.22 is released!  https://go.dev/blog/go1.22
- doc: Go 1.22 Release Notes https://go.dev/doc/go1.22

gotip: The gotip command compiles and runs the go command from the development tree.
- https://pkg.go.dev/golang.org/dl/gotip
- tip doc: Go 1.22 Release Notes https://tip.golang.org/doc/go1.22

Go Version Management
- golang.org/dl https://github.com/golang/dl
- gvm https://github.com/moovweb/gvm
- goenv https://github.com/go-nv/goenv
- Goup https://github.com/owenthereal/goup


## 二、Go Package Management
Go Packages
- https://pkg.go.dev

### Go Modules
- Using Go Modules https://go.dev/blog/using-go-modules
- Tutorial: Create a Go module https://go.dev/doc/tutorial/create-module
- doc/modules https://go.dev/doc/modules/
  - https://go.dev/doc/modules/managing-dependencies
- Go Wiki: Go Modules https://go.dev/wiki/Modules
- Go Modules Reference https://go.dev/ref/mod

Go Modules 介绍
- 干货满满的 Go Modules https://juejin.cn/post/6844903954879348750
- go mod 使用 https://juejin.cn/post/6844903798658301960
- Go 包管理解决之道 —— Modules 初试_ https://www.haoyizebo.com/posts/deaff727/
- Go Modules 详解 https://www.sulinehk.com/post/go-modules-details/

Go Modules 使用
- 使用go module导入本地包 https://zhuanlan.zhihu.com/p/109828249
- How to get another branch instead of default branch with go get https://stackoverflow.com/questions/42761820/how-to-get-another-branch-instead-of-default-branch-with-go-get

vgo
- https://github.com/golang/vgo
- https://research.swtch.com/vgo
- https://research.swtch.com/vgo-import

V2 Version
- Evolving the Go Standard Library with math/rand/v2 https://go.dev/blog/randv2

### Go multi-module workspaces
- Get familiar with workspaces https://go.dev/blog/get-familiar-with-workspaces
- Tutorial: Getting started with multi-module workspaces https://go.dev/doc/tutorial/workspaces
- Go Modules Reference - Workspaces https://go.dev/ref/mod#workspaces

### Go Modules Services
> Go Module Mirror, Index, and Checksum Database
- https://proxy.golang.org
- https://sum.golang.org
- https://index.golang.org
- https://index.golang.org/index

> Go Modules Proxy
- https://goproxy.cn
- https://goproxy.io
- https://developer.aliyun.com/mirror/goproxy


## 三、Best Practice
### Roadmap
- Awesome Go https://github.com/avelino/awesome-go
- Go Developer Roadmap https://github.com/darius-khll/golang-developer-roadmap
- Learning Go in 2024 https://www.bytesizego.com/blog/learning-golang-2024
- GOLANG ROADMAP https://golangroadmap.com
  - https://m.golangroadmap.com/#/collection
  - https://m.golangroadmap.com/#/collectiondetail?id=6

### gonew
介绍
- Experimenting with project templates https://go.dev/blog/gonew
- 又有新功能！Go 将有生成新模板的 gonew 工具链 https://juejin.cn/post/7265319735746920506
- Go项目初始化不再困扰你：gonew全方位解析 https://tonybai.com/2023/08/11/introduction-to-the-gonew-tool/

模版
- https://github.com/GoogleCloudPlatform/go-templates
- https://github.com/ServiceWeaver/template

### Project Layout
Organizing a Go module
- https://go.dev/doc/modules/layout
- https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/

Third Party
- https://github.com/golang-standards/project-layout
  - https://github.com/golang-standards/project-layout/issues/117
  - https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/
- https://github.com/bxcodec/go-clean-arch
- https://github.com/evrone/go-clean-template
- Golang 简洁架构实战 https://mp.weixin.qq.com/s/Xzlt_WcdcfLWhofafY3c6g
- Golang 整洁架构实践 https://mp.weixin.qq.com/s/I2Fx2TIrwXV2kfLj_T5g5g

### Best Practice
- Go Style Best Practices https://github.com/google/styleguide/blob/gh-pages/go/best-practices.md
- Don't write clean code, write CRISP code https://bitfieldconsulting.com/posts/crisp-code
- Common Anti-Patterns in Go Web Applications https://threedots.tech/post/common-anti-patterns-in-go-web-applications/
- Teh Zen of Go https://dave.cheney.net/2020/02/23/the-zen-of-go
  - Go语言之禅 https://mp.weixin.qq.com/s/mKe0w3urHJEe9iYJ7svCuA
- Go 语言实战: 编写可维护 Go 语言代码建议 https://github.com/llitfkitfk/go-best-practice
- TOP 20 Go最佳实践 https://colobu.com/2023/11/17/golang-quick-reference-top-20-best-coding-practices/
- 腾讯
  - Golang 编程思维和工程实战 https://mp.weixin.qq.com/s/llmE9QpnrvA02AtvfHtqJQ
  - Go 语言中各式各样的优化手段 https://zhuanlan.zhihu.com/p/403417640
  - Go语言“正统”在中国？这6点教你写好Go代码！https://mp.weixin.qq.com/s/YTreb3YhEZMFFMlC5ru61Q
  - 如何真正写好Golang代码? https://mp.weixin.qq.com/s/OIHqmgK4V7Y26uYoFjsCyA
  - Golang与Java全方位对比总结 https://mp.weixin.qq.com/s/-N4eqdXb9a93uvOWfE4ScQ
- 100 Go Mistakes and How to Avoid Them
  - https://github.com/teivah/100-go-mistakes
  - https://book.douban.com/subject/36084407/
  - Go语言中常见100问题 https://mp.weixin.qq.com/s/tqAcoskLYd2UV6B_6AeIjw
  - [长文]从《100 Go Mistakes》我总结了什么？https://www.luozhiyun.com/archives/797

### Standard
Style
- Go Style | styleguide https://google.github.io/styleguide/go/index
  - https://github.com/google/styleguide/tree/gh-pages/go
- The Uber Go Style Guide https://github.com/uber-go/guide
  - https://github.com/uber-go/guide/blob/master/style.md
- The Ultimate Go Study Guide https://github.com/hoanhan101/ultimate-go

Lint
- golangci-lint https://github.com/golangci/golangci-lint
- Staticcheck - The advanced Go linter https://github.com/dominikh/go-tools
- https://github.com/golang/lint

Deadcode
- Finding unreachable functions with deadcode https://go.dev/blog/deadcode
- deadcode command https://pkg.go.dev/golang.org/x/tools/cmd/deadcode

安全
- 腾讯-Go安全指南 https://github.com/Tencent/secguide/blob/main/Go安全指南.md


## 四、Go Community
- golang-nuts https://groups.google.com/g/golang-nuts

Go 101
- https://github.com/go101
- https://go101.org
  - Go (Fundamentals) 101 https://go101.org/article/101.html
  - Go Generics 101 https://go101.org/generics/101.html
  - Go Optimizations 101 https://go101.org/optimizations/101.html
  - Go Details & Tips 101 https://go101.org/details-and-tips/101.html
  - Go Quizzes 101 https://go101.org/quizzes/101.html

Newsletter
- Golang Weekly https://golangweekly.com
- Microsoft for Go Developers https://devblogs.microsoft.com/go/

Forum
- LearnKu https://learnku.com/go

Go Training Class 
- https://github.com/ardanlabs/gotraining


### Go Open-Source Projects
Well-known
- Docker
- Kubernetes
- Prometheus
- Etcd
- CockroachDB
- Hugo
- Gin
- Gorilla Toolkit

Other
- https://github.com/gogs/gogs
- https://github.com/go-gitea/gitea
- https://github.com/apache/incubator-answer


