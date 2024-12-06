# Goroutine scheduler
CSP
- Bell Labs and CSP Threads https://swtch.com/~rsc/thread/
- Share Memory By Communicating https://go.dev/blog/codelab-share

## 原理
实现原理
- Scalable Go Scheduler Design Doc https://golang.org/s/go11sched
- Scheduler structures https://go.dev/src/runtime/HACKING#Scheduler-structures
- 理解Go协程调度的本质 https://mp.weixin.qq.com/s/j9OpuIxXRWa9524oacGCzw

src/runtime 源码
- Goroutine scheduler https://go.dev/src/runtime/proc.go
- https://go.dev/src/runtime/runtime2.go

原理总结
- 详解 Go 程序的启动流程，你知道 g0，m0 是什么吗？https://mp.weixin.qq.com/s/YK-TD3bZGEgqC0j-8U6VkQ
- g0 特殊的goroutine https://blog.haohtml.com/archives/22353/
- Go 调度模型 GPM https://segmentfault.com/a/1190000022871460
- 深入golang runtime的调度 https://zboya.github.io/post/go_scheduler/
- 调度器 https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/
- 6.8 协作与抢占 | Go 语言原本 https://golang.design/under-the-hood/zh-cn/part2runtime/ch06sched/preemption/
- Go 是协作式调度，还是抢占式调度？ https://v2ex.com/t/927783
- 深度解密Go语言之基于信号的抢占式调度 https://mp.weixin.qq.com/s/ESfbVoCGUIdpwFOfGyvz1w
- 深入了解 Go 语言与并发编程 https://mp.weixin.qq.com/s/obFUsRnppgEsGkoo08nWeQ
- goroutine 调度器原理 https://mp.weixin.qq.com/s/eX32jJMQBXov-HS8vRxo0Q
- Go Scheduler 的 GMP 模型 https://mp.weixin.qq.com/s/1CY3E5daJ5U42orVwzCpaw
- goroutine调度器揭秘 https://colobu.com/2024/03/19/goroutine-scheduler-revealed/
- goroutine调度器揭秘 2 https://colobu.com/2024/03/24/goroutine-scheduler-2/
- Go语言的CSP模型 https://zhuanlan.zhihu.com/p/313763247
- GM到GMP，Golang经历了什么？ https://zhuanlan.zhihu.com/p/364590925
- go并发奥秘：GMP模型｜Go主题月 https://juejin.cn/post/6944925506340913188
- 动图图解！GMP模型里为什么要有P？背后的原因让人暖心 https://juejin.cn/post/6956008643456139301
- Golang调度器的GMP模型 http://www.liangyaopei.com/2020/10/03/scheduler-gmp-model/
- Go的CSP并发模型实现：M, P, G https://www.cnblogs.com/sunsky303/p/9115530.html#4807807
- 深入Golang调度器之GMP模型 https://www.cnblogs.com/sunsky303/p/9705727.html
- Head First of Golang Scheduler https://zhuanlan.zhihu.com/p/42057783
- GMP 原理与调度 https://www.topgoer.com/%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B/GMP%E5%8E%9F%E7%90%86%E4%B8%8E%E8%B0%83%E5%BA%A6.html
- 如何更直观地理解 Go 调度过程 https://mp.weixin.qq.com/s/xjRryI2htlI0VsoGypoQfQ
- 假如 Go 能说话，听听 GMP 的心声 https://juejin.cn/post/6964956868720623646
- Go的GMP模型真的“简单” https://mp.weixin.qq.com/s/9rBFZGlu8Kt5W8VgdoQkow
- 如何更直观地理解 Go 调度过程 https://mp.weixin.qq.com/s/HXDMcQGEHMhcwgvciqPs8A
- 深入理解 Go scheduler 调度器- GPM 源码分析 https://mp.weixin.qq.com/s/Z29hSMez4PKzp5JU-lV-hA
- Goroutine 数量控制在多少合适，会影响 GC 和调度？https://mp.weixin.qq.com/s/uWP2X6iFu7BtwjIv5H55vw
- https://embeddedgo.github.io/2020/06/21/playing_with_go_schedulers_on_a_dual-core_risc-v.html
