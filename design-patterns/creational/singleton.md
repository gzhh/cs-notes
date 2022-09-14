### 单例模式可以分为懒汉式和饿汉式。
懒汉式就是创建对象时比较懒，先不急着创建对象，在需要加载配置文件的时候再去创建。
饿汉式就是在系统初始化的时候我们已经把对象创建好了，需要用的时候直接拿过来用就好了。

### 懒汉式
Implementation
```
package singleton

type singleton map[string]string

var (
    once sync.Once

    instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}
```

Usage
```
s := singleton.New()

s["this"] = "that"

s2 := singleton.New()

fmt.Println("This is ", s2["this"])
// This is that
```


### 饿汉式
```
func init() {
    // TODO
}
```
