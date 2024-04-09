### Select
介绍
- The select statement lets a goroutine wait on multiple communication operations.
  - A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
- The default case in a select is run if no other case is ready.
  - Use a default case to try a send or receive without blocking:
- Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.

原理
- https://go.dev/src/runtime/select.go
- select 实现原理: https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-select/
- 使用反射 reflect.Select 操作 Channel: https://www.pursuitking.com/go_2/1-14.html

使用场景
```
// 1.多渠道的选择
func TestDemo1(t *testing.T) {
    select {
    case ret := <-retCh1:
        t.Logf("result %s", ret)
    case ret := <-retCh2:
        t.Logf("result %s", ret)
    default:
        t.Error("No one returned")
    }
}

// 2.超时控制
func TestDemo2(t *testing.T) {
    select {
    case ret := <-retCh:
        t.Logf("result %s", ret)
    case <-time.After(time.Second * 1):
        t.Error("time out")
    }
}
```

参考：
- https://go.dev/tour/concurrency/5
- https://go.dev/tour/concurrency/6
- https://gobyexample.com/select
