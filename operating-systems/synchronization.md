## Synchronization 同步（原语）
- https://en.wikipedia.org/wiki/Synchronization_(computer_science)

### 同步经典问题
- Producer–consumer problem: https://en.wikipedia.org/wiki/Producer-consumer_problem
- Readers–writers problem: https://en.wikipedia.org/wiki/Readers-writers_problem
- Dining philosophers problem: https://en.wikipedia.org/wiki/Dining_philosophers_problem
- ABA problem: https://en.wikipedia.org/wiki/ABA_problem

### 同步要求
- Critical section 临界区: https://en.wikipedia.org/wiki/Critical_section
- Race condition 竞态条件: https://en.wikipedia.org/wiki/Race_condition#Software
- Concurrency 并发: https://en.wikipedia.org/wiki/Concurrency_(computer_science)
  - Concurrency control: https://en.wikipedia.org/wiki/Concurrency_control 
    - Lock: https://en.wikipedia.org/wiki/Lock_(computer_science)
- 同步还需要解决的问题：
  - Deadlock: https://en.wikipedia.org/wiki/Deadlock
  - Starvation: https://en.wikipedia.org/wiki/Starvation_(computer_science)
  - Priority inversion: https://en.wikipedia.org/wiki/Priority_inversion
  - Busy waiting: https://en.wikipedia.org/wiki/Busy_waiting

### 硬件同步
机器指令
- Test-and-set: https://en.wikipedia.org/wiki/Test-and-set
- Compare-and-swap: https://en.wikipedia.org/wiki/Compare-and-swap

### 编程语言支持
- Monitor: https://en.wikipedia.org/wiki/Monitor_(synchronization)
  - 例如 Java 中的 Monitor
- 在Go语言中同步和并发控制通常使用通道（Channel）和互斥锁（Mutex）等机制来实现，而不是像Java那样使用Monitor。

### 同步实现
- Spinlock: https://en.wikipedia.org/wiki/Spinlock
- Semaphore: https://en.wikipedia.org/wiki/Semaphore_(programming)
- Barrier: https://en.wikipedia.org/wiki/Barrier_(computer_science)
  - Memory barrier: https://en.wikipedia.org/wiki/Memory_barrier
- Mutual exclusion: https://en.wikipedia.org/wiki/Mutual_exclusion
- Readers–writer lock: https://en.wikipedia.org/wiki/Readers%E2%80%93writer_lock

Ref
- linux下信号量和互斥锁的区别
  - https://aiops.com/news/post/6191.html
  - ChatGPT

