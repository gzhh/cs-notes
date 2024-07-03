## struct
A struct is a collection of fields.

### 最佳实践
- Go: Should I Use a Pointer instead of a Copy of my Struct? https://medium.com/a-journey-with-go/go-should-i-use-a-pointer-instead-of-a-copy-of-my-struct-44b43b104963
- 面试中结构体的那些事 https://mp.weixin.qq.com/s/5As91jd-PbC2SPgLuEnDbg
- 在 Go 中如何检查结构体是否为空 https://mp.weixin.qq.com/s/6KqG53y_9QwFCKAFe66D4w

### 结构体拷贝
浅拷贝

### struct 比较
引用类型，如果你要比较的话，是想去比较他的值还是说他的地址？这就相当于在比较的时候会有歧义产生。Go从语言层面上直接杜绝了引用类型的比较；引用类型不能比较，比如 slice、map、channel、func 等。

PS：引用类型可以和 nil 直接比较

Go中结构体能不能比较，分为两种情况：
- 一种是同类的结构体，如果是同类的结构体，并且结构体的字段没有不可比较的类型，如Slice、Map、Channel、Func，那么就可以比较，反之如果有不可比较的类型就不行；
- 另一种是不同类的结构体，不同类的结构体如果成员变量类型、变量名和顺序都相同，而且结构体没有不可比较字段时，那么进行显式类型转换后就可以比较，反之不能比较。


### 空结构体 struct{}
特点：空结构体的 sizeof 为 0

使用场景
- 实现方法接收者
- 实现集合类型
- 实现空通道

Ref
- https://juejin.cn/post/7018037514246029349

