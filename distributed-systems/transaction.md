# Distributed transaction
- https://en.wikipedia.org/wiki/Distributed_transaction

事务特性 ACID，本地可以实现；但是因为有 CAP 理论约束，分布式事务不能实现ACID。


## 最佳实践
- 再有人问你分布式事务，把这篇扔给他 https://juejin.cn/post/6844903647197806605
- 拜托，面试请不要再问我TCC分布式事务的实现原理！https://juejin.cn/post/6844903716089233416
- 分布式事务 https://icyfenix.cn/architect-perspective/general-architecture/transaction/distributed.html
- 分布式事务概述与项目实战 https://mp.weixin.qq.com/s/0Io-X0S9AY-s0HeRb_jbag
- 分布式系统 - 分布式事务及实现方案 https://pdai.tech/md/arch/arch-z-transection.html
- 分布式事务最经典的七种解决方案 https://segmentfault.com/a/1190000040321750
- 七种分布式事务的解决方案，一次讲给你听！https://cloud.tencent.com/developer/article/1806989


## Third Lib
- Distributed Transactions Manager https://github.com/dtm-labs/dtm
  - go-zero微服务实战系列（十、分布式事务如何实现）https://mp.weixin.qq.com/s/OWlck8Vgcy6hvyEc3EYGSw
  - go-zero对接分布式事务dtm保姆式教程 https://github.com/Mikaelemmmm/gozerodtm
