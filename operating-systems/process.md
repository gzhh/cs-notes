## 进程

- https://en.wikipedia.org/wiki/Process_(computing)
- https://en.wikipedia.org/wiki/Inter-process_communication

### 概念
关键几个函数
- fork
- exec系列
- _exit
- wait 和 waitpid

概念

1. 用fork创建进程时,在子进程中使用 getppid 得到的返回值为啥总是1
    
    由于父进程先退出了，造成子进程被init（ID=1）接管，所以用 getppid 出来的是1。


