# Go Test

docs
- https://pkg.go.dev/testing
- https://go.dev/doc/tutorial/add-a-test
- https://gobyexample.com/testing

Learn Go with Tests (TDD)
- https://go.dev/blog/examples
- https://github.com/quii/learn-go-with-tests
- https://quii.gitbook.io/learn-go-with-tests

测试框架
- https://github.com/stretchr/testify
- https://github.com/onsi/ginkgo
- https://github.com/smartystreets/goconvey
- https://github.com/dvyukov/go-fuzz


## Go 单元测试
最佳实践
- 单元测试有必要吗？https://www.v2ex.com/t/821608
- 如何编写可测试的代码：两个核心三个思路 https://mp.weixin.qq.com/s/rkIqAElMQ68sLWm927ru6w
- Golang 编写易于单元测试的代码 https://blog.hackerpie.com/posts/testing/golang-write-testable-codes/
- GoLang快速上手单元测试（思想、框架、实践）https://learnku.com/articles/52896

李文周-Go单测从零到溜系列
- 单元测试基础 https://www.liwenzhou.com/posts/Go/unit-test-0/
- 网络测试 https://www.liwenzhou.com/posts/Go/unit-test-1/
- MySQL和Redis测试 https://www.liwenzhou.com/posts/Go/unit-test-2/
- mock接口测试 https://www.liwenzhou.com/posts/Go/unit-test-3/
- 使用monkey打桩 https://www.liwenzhou.com/posts/Go/unit-test-4/
- goconvey的使用 https://www.liwenzhou.com/posts/Go/unit-test-5/
- 编写可测试的代码 https://www.liwenzhou.com/posts/Go/unit-test-6/

### Go 单元测试 Mock
Mock 框架
- https://github.com/golang/mock
- https://github.com/uber-go/mock
- https://github.com/vektra/mockery
- https://github.com/DATA-DOG/go-sqlmock
- https://github.com/jarcoal/httpmock

- Go单元测试 Mock 方案总结: https://taoshu.in/go/mock.html
  - https://taoshu.in/go/monkey/monkey.html
  - Go语言实现猴子补丁
    - https://en.wikipedia.org/wiki/Monkey_patch
    - https://taoshu.in/go/monkey/monkey.html
    - https://taoshu.in/go/monkey/monkey-2.html
    - https://taoshu.in/go/monkey/monkey-3.html


## go test 命令
- https://pkg.go.dev/cmd/go#hdr-Testing_flags

执行当前目录下所有的 *_test.go 文件，不显示执行过程
- `go test`

执行当前目录下指定的 *_test.go 文件
- `go test utils_test.go utils.go`
- https://stackoverflow.com/questions/16935965/how-to-run-test-cases-in-a-specified-file

执行当前目录下所有的 *_test.go 文件
- `go test -v`

执行当前目录下及其所有子目录下的 *_test.go 文件
- `go test -v ./...`

执行指定目录下所有的 *_test.go 文件
- `go test -v ./{specific_dir}/`

执行指定目录下及其所有子目录下的 *_test.go 文件
- `go test -v ./{specific_dir}/...`

### go test -run 指定测试函数
- https://go.dev/blog/subtests

精准匹配
- `go test -v -run={func_name}`

正则匹配
- `go test -v -run "^{some_words}$"`

### go test -cover 测试覆盖率
- https://go.dev/blog/cover
- https://go.dev/blog/integration-test-coverage

当前目录及其子目录的测试覆盖率
- `go test -cover ./...`
