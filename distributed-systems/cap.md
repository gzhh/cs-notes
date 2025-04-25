# CAP 理论
- [CAP theorem](https://en.wikipedia.org/wiki/CAP_theorem)
- [一文看懂｜分布式系统之CAP理论](https://cloud.tencent.com/developer/article/1860632)
- [一文吃透分布式理论【CAP,BASE深入解析】](https://juejin.cn/post/7021717177220726798)

CAP 三选二（由于是分布式系统，P 必选）
- 选C（要求强一致性，放弃可用性），例子：Redis, HBase, Zookeeper；场景：银行转账
- 选A（放弃强一致性，保留最终一致性，大部分系统），场景：淘宝订单退款、12306 买票

