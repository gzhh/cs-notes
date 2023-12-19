## Error handling

In Go, if an application panics, there is a bug — not an error.

Many developers new to Go confuse a panic with a runtime error in other languages. Go has errors for that. Panics should not be used to check if something works or not. Always use errors for that.

Catching panics can make sense to prevent unknown bugs from crashing a production server. Such a bug needs to be fixed though and using recover to catch it should only be a measure to prevent disaster in high-stake scenarios — never to be used as a way to do “alternative logic” as you put it.


blog
- [Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
- [Error handling and Go](https://go.dev/blog/error-handling-and-go)
- [Errors are values](https://go.dev/blog/errors-are-values)
- [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)


### Error
error type
```
type error interface {
    Error() string
}
```

package errors
- https://pkg.go.dev/errors


### Defer
defer statements rules
1. A deferred function’s arguments are evaluated when the defer statement is evaluated.
2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
3. Deferred functions may read and assign to the returning function’s named return values.

tour
- https://go.dev/tour/flowcontrol/12
- https://go.dev/tour/flowcontrol/13


### Panic
[builtin#panic](https://pkg.go.dev/builtin#panic)

Panic is a built-in function that stops the ordinary flow of control and begins panicking.


### Recover
[builtin#recover](https://pkg.go.dev/builtin#recover)

Recover is a built-in function that regains control of a panicking goroutine.


### Best Practice
Package errgroup
- https://pkg.go.dev/golang.org/x/sync/errgroup

注意点：
- 检查错误的类型
- 检查错误的值
- 不要处理错误两次
- 不要忘了在 defer 中处理错误


### Troubleshooting
- [Why does `defer recover()` not catch panics?](https://stackoverflow.com/questions/29518109/why-does-defer-recover-not-catch-panics)
- [Why recover() does not work in a nested deferred function?](https://stackoverflow.com/questions/49344478/why-recover-does-not-work-in-a-nested-deferred-function)


### Third party library
- https://github.com/pkg/errors
- https://github.com/go-errors/errors
- https://github.com/hashicorp/go-multierror
- https://github.com/cockroachdb/errors


### Ref
- Go Specification
  - https://go.dev/ref/spec#Defer_statements
  - https://go.dev/ref/spec#Handling_panics
  - https://go.dev/ref/spec#Errors
  - https://go.dev/ref/spec#Run_time_panics
- Go Tutorial
  - [Return and handle an error](https://go.dev/doc/tutorial/handle-errors)
- Go by Example
  - [Errors](https://gobyexample.com/errors)
  - [Panic](https://gobyexample.com/panic)
  - [Defer](https://gobyexample.com/defer)
  - [Recover](https://gobyexample.com/recover)
- Articles
  - [Panic recover in Go v.s. try catch in other languages](https://stackoverflow.com/questions/3413389/panic-recover-in-go-v-s-try-catch-in-other-languages)
  - [Panic and Recover](https://golangbot.com/panic-and-recover/)
  - [你对 Go 错误处理的 4 个误解！- 脑子进煎鱼了](https://mp.weixin.qq.com/s/Ey-yqIq__wpaLTlBAOHjxg)
  - [Go 业务开发中 Error & Context - 毛剑](https://github.com/gopherchina/conference/blob/master/2019/1.5%20Go%20%E4%B8%9A%E5%8A%A1%E5%BC%80%E5%8F%91%E4%B8%AD%20Error%20%26%20Context%20-%20%E6%AF%9B%E5%89%91.pdf)
  - [Go 错误处理总结与实践](https://www.pseudoyu.com/zh/2021/08/29/go_error_handling/)
  - [Use %w instead of %v or %s](https://stackoverflow.com/questions/61283248/format-errors-in-go-s-v-or-w)
  - [卷起来，老程序员也得了解errors包的新变化 - 鸟窝](https://colobu.com/2023/12/13/learn-more-about-errors/)
