# Distributed systems
- Distributed computing https://en.wikipedia.org/wiki/Distributed_computing
- What is a distributed system? https://www.atlassian.com/microservices/microservices-architecture/distributed-architecture
- 分布式架构知识体系 https://mp.weixin.qq.com/s/pQqSzMuF_H_5OZsYKvRiuA


## Theory
### Consensus 共识
- https://en.wikipedia.org/wiki/Consensus_(computer_science)

Paxos
- https://en.wikipedia.org/wiki/Paxos_(computer_science)

Raft
- https://en.wikipedia.org/wiki/Raft_(algorithm)
- https://zh.wikipedia.org/wiki/Raft
- The Raft Consensus Algorithm https://raft.github.io/
- 容易理解的Distributed Consensus分布式共识算法 http://kailing.pub/raft/index.html
- 八分钟了解一致性算法 -- Raft算法 https://mp.weixin.qq.com/s/FFkGKaNpNwJtkpaO5yyYrg
- 分布式系统的一致性与共识（一）https://zhuanlan.zhihu.com/p/68743917
- Implementing Raft: Part 0 - Introduction https://eli.thegreenplace.net/2020/implementing-raft-part-0-introduction/
- Golang implementation of the Raft consensus protocol https://github.com/hashicorp/raft

### Clock synchronization
- https://en.wikipedia.org/wiki/Clock_synchronization
- Time, Clocks, and the Ordering of Events in a Distributed System https://lamport.azurewebsites.net/pubs/time-clocks.pdf
- 分布式领域最重要的一篇论文，到底讲了什么？https://mp.weixin.qq.com/s/FZnJLPeTh-bV0amLO5CnoQ


## 分布式问题
### 拜占庭将军问题
- https://zh.wikipedia.org/wiki/拜占庭将军问题

### 讨论
- 既然 mysql 可以靠 binlog 达成多副本共识，那为什么还需要 Paxos、Raft 等共识算法？https://v2ex.com/t/1014337
