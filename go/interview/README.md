# Go interview
1. new 和 make 的区别

    它们均在堆空间分配内存

    make 一般用来初始化 slice、map、channel（内建类型）
    
    new 一般用来类型的内存分配、初始化对象、struct，返回指针（自建类型）
    
2. slice的底层原理
    
    有时候共用内存、有时候不是
    
    ```
    x := []int{1, 2, 3}
    y := x[:2]
    fmt.Printf("x = %v\n", x)
    fmt.Printf("y = %v\n", y)
    
    y = append(y, 50)
    fmt.Printf("x = %v\n", x)
    fmt.Printf("y = %v\n", y)
    
    y = append(y, 60)
    fmt.Printf("x = %v\n", x)
    fmt.Printf("y = %v\n", y)
    
    y[0] = 20
    fmt.Printf("x = %v\n", x)
    fmt.Printf("y = %v\n", y)
    ```
    
3. channel
    - 不能在close之后向channel在发送消息
    - 空channel和指定大小channel的区别
    - 不要用共享内存来操作变量，而是用channel来共享内存
    
4. goroutine 阻塞
    
    有某一个goroutine一直占用资源，gmp调度怎么处理
    
    goroutine遇到oom怎么处理
    
    - 用法及技巧
        - goroutine内的错误，子goroutine的错误
5. 内存管理
    
    gc
    
    TODO
    
    https://en.wikipedia.org/wiki/Out_of_memory
    
6. 函数传参传指针好还是传值好
    - 看场景，传指针一般用来改变参数的值，传值一般不需要改变参数值
    - 在go中传指针可能会发生逃逸，把原本放在栈上的内存逃逸到堆上
    - 在c中传指针可以节省内存拷贝
7. 进程、线程、协程
    
    进程
    
    线程
    
    - 线程模型：用户级线程和内核级线程的映射
        - 一对一模型
        - 一对多模型
        - 多对多模型
    
    协程
    
    区别
    
8. GMP 模型
    
    GMP概念
    
    - G
    - M
    - P
    
    gmp 中线程的几种状态？
    
9. 怎么定位线上问题
    
    pprof
    
10. 错误处理
    
    TODO
    
11. gRPC
    - grpc gateway
    - 服务发现
    - 负载均衡
12. gin 框架
    - 参数校验
    - 中间件
13. 结构体反射
    - 结构体tag解析
    - 应用
14. golang 锁机制 Mutex
    - 模型：正常模式、饥饿模式
    - 应用：并发map读写
        - 给map加锁
        - syncmap

其他：

MySQL

- 锁
- orm
    
    xorm、gorm
    
- 分库分表
    - 分库
        
        点击数据：按uid、imei等分
        
        回传数据：按android、ios分
        
    - 分表
        
        水平拆分：表分区、按uid hash拆分
        
        垂直拆分：按业务字段拆分表
        

Redis

- 分布式锁 setnx
- 集群模式、主从模式、哨兵模式？
- 几种数据结构
- 持久化：rdb、aof

负载均衡算法

- 随机
- 轮询
- hash
