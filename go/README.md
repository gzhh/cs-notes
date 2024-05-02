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

### Standard
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


