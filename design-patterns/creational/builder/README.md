# 生成器模式

将复杂对象的建造过程抽象出来（抽象类别），使这个抽象过程的不同实现方法可以构造出不同表现（属性）的对象。

https://zh.wikipedia.org/wiki/生成器模式

使用场景
- 当创建复杂对象的算法应该独立于该对象的组成部分以及它们的装配方式时（当一个类的构造函数参数个数超过4个，而且这些参数有些是可选的参数，可以考虑使用构造者模式。）
- 当构造过程必须允许被构造的对象有不同的表示时。

扩展
- [Go 语言的函数选项模式](https://www.liwenzhou.com/posts/Go/functional-options-pattern/)

### 例子

```
// Config struct
type Config struct {
	Port int
}

// Config builder struct, containing an optional port
type ConfigBuilder struct {
	port *int
}

// Public method to set up the port
func (b *ConfigBuilder) Port(
	port int) *ConfigBuilder {
	b.port = &port
	return b
}

// Build method to create the config struct
func (b *ConfigBuilder) Build() (Config, error) {
	cfg := Config{}

	// Main logic related to port management
	if b.port == nil {
		cfg.Port = defaultHTTPPort
	} else {
		if *b.port == 0 {
			cfg.Port = randomPort()
		} else if *b.port < 0 {

			return Config{}, errors.New("port should be positive")
		} else {
			cfg.Port = *b.port
		}
	}

	return cfg, nil
}
```

使用
```
builder := ConfigBuilder{}  // Creates a builder config
builder.Port(8080)          // Sets the port
cfg, err := builder.Build() // Builds the config struct
if err != nil {
    return err
}

server, err := NewServer("localhost", cfg) // Passes the config struct
if err != nil {
    return err
}
```

### Ref
- https://zhuanlan.zhihu.com/p/58093669
