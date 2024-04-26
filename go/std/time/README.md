## Time
- https://pkg.go.dev/time

最佳实践：
- https://gobyexample.com/time
- https://gobyexample.com/time-formatting-parsing
- https://gobyexample.com/timeouts
- https://gobyexample.com/timers
- https://gobyexample.com/tickers

参考：
- https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/runtime/time.go
- https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/time/time.go
- https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/time/sleep.go
- https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/time/tick.go


### time.Parse and time.Format
时区相关
```
datetimeStr := "2023-08-26 09:00:00"

// Parse with utc
t, err := time.Parse(time.DateTime, datetimeStr)

// Parse with location info
// t, err := time.ParseInLocation(time.DateTime, datetimeStr, time.Local)

if err != nil {
    panic(err)
}

fmt.Println(t)
fmt.Println(t.Unix())

// Format with time instance
fmt.Println(t.Format(time.DateTime))

// Format with time instance, In set location information
fmt.Println(t.In(time.UTC).Format(time.DateTime))
fmt.Println(t.In(time.Local).Format(time.DateTime))
loc, err := time.LoadLocation("Asia/Shanghai")
if err != nil {
    panic(err)
}
fmt.Println(t.In(loc).Format(time.DateTime))
```

Ref
- time/time.go
- time/zoneinfo.go
