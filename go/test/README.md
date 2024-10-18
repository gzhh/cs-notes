# Go Test

docs
- https://pkg.go.dev/testing
- https://go.dev/doc/tutorial/add-a-test
- https://gobyexample.com/testing

Learn Go with Tests (TDD)
- https://go.dev/blog/examples
- Testing in Go - a crash article to get you going https://thedevelopercafe.com/articles/testing-in-go-929e2ad2

- https://github.com/quii/learn-go-with-tests
- https://quii.gitbook.io/learn-go-with-tests

测试框架
- https://github.com/stretchr/testify
- https://github.com/onsi/ginkgo
- https://github.com/smartystreets/goconvey
- https://github.com/dvyukov/go-fuzz

分类
1. 测试函数
    测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头，并且带有一个*testing.T类型的参数；
2. 基准测试函数
    以Benchmark为前缀名，并且带有一个*testing.B类型的参数；
3. 示例函数
    以Example为函数名开头


## Go Test Lib
- Package for comparing Go values in tests https://github.com/google/go-cmp


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
  - JetBrains Guide - Go
    - https://www.jetbrains.com/guide/go/tutorials/mock_testing_with_go/
    - https://github.com/JetBrains/go-code-samples/tree/main/mock-testing
    - https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/
    - https://github.com/JetBrains/go-code-samples/tree/main/testing-guide
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

执行当前目录下所有的 *_test.go 文件，参数-v可用于打印每个测试函数的名字和运行时间
- `go test -v`

执行当前目录下及其所有子目录下的 *_test.go 文件
- `go test -v ./...`

执行指定目录下所有的 *_test.go 文件
- `go test -v ./{specific_dir}/`

执行指定目录下及其所有子目录下的 *_test.go 文件
- `go test -v ./{specific_dir}/...`

### go test -run 指定测试函数
参数-run对应一个正则表达式，只有测试函数名被它正确匹配的测试函数才会被go test测试命令运行
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

### go test -bench

压力测试

go test -bench=.

内存相关分析

go test -bench=. -benchmem

### 子目录测试

测试当前目录及其所有的子目录

go test ./…

测试指定子目录

go test ./path_to_sub-directory/

### go list

我们可以用go list命令查看包对应目录中哪些Go源文件是产品代码，哪些是包内测试，还有哪些是外部测试包。我们以fmt包作为一个例子：GoFiles表示产品代码对应的Go源文件列表；也就是go build命令要编译的部分。

```
$ go list -f={{.GoFiles}} fmt
[doc.go format.go print.go scan.go]
```

TestGoFiles表示的是fmt包内部测试代码，以_test.go为后缀文件名，不过只在测试时被构建：

```
$ go list -f={{.TestGoFiles}} fmt
[export_test.go]
```

XTestGoFiles表示的是属于外部测试包的测试代码，也就是fmt_test包，因此它们必须先导入fmt包。同样，这些文件也只是在测试时被构建运行：

```
$ go list -f={{.XTestGoFiles}} fmt
[fmt_test.go scan_test.go stringer_test.go]
```

有时候外部测试包也需要访问被测试包内部的代码，例如在一个为了避免循环导入而被独立到外部测试包的白盒测试。在这种情况下，我们可以通过一些技巧解决：我们在包内的一个_test.go文件中导出一个内部的实现给外部测试包。因为这些代码只有在测试时才需要，因此一般会放在export_test.go文件中。]

为了确保fmt.isSpace和unicode.IsSpace函数的行为保持一致，fmt包谨慎地包含了一个测试。一个在外部测试包内的白盒测试，是无法直接访问到isSpace内部函数的，因此fmt通过一个后门导出了isSpace函数。export_test.go文件就是专门用于外部测试包的后门。

```
package fmt

var IsSpace = isSpace
```

这个测试文件并没有定义测试代码；它只是通过fmt.IsSpace简单导出了内部的isSpace函数，提供给外部测试包使用。这个技巧可以广泛用于位于外部测试包的白盒测试。
