# Socket
[https://zh.wikipedia.org/wiki/Berkeley套接字](https://zh.wikipedia.org/wiki/Berkeley%E5%A5%97%E6%8E%A5%E5%AD%97)

![socket-flow](images/socket-flow.png)

## IO模型

### 阻塞式IO

### 非阻塞式IO

### IO多路复用（I/O Multiplexing）

select和poll

epoll

kqueue

有点：单个process同时处理多个网络连接的IO

基本原理：select，poll，epoll，kqueue这些函数不断的轮询所负责的所有socket，当某个socket就绪（一般是读就绪或者写就绪），就通知用户进程

`当用户进程调用了select，那么整个进程会被block`，而同时，kernel会“监视”所有select负责的socket，当任何一个socket中的数据准备好了，select就会返回。这个时候用户进程再调用read操作，将数据从kernel拷贝到用户进程。

所以，`I/O 多路复用的特点是通过一种机制一个进程能同时等待多个文件描述符，而这些文件描述符（套接字描述符）其中的任意一个进入读就绪状态，select()函数就可以返回`

### 信号驱动

SIGIO

### 异步IO

POSIX的aio_系列函数

### Ref
- [https://zhuanlan.zhihu.com/p/115912936](https://zhuanlan.zhihu.com/p/115912936)
- [https://blog.csdn.net/ZWE7616175/article/details/80591587](https://blog.csdn.net/ZWE7616175/article/details/80591587)
- [https://blog.csdn.net/GJQI12/article/details/105118354](https://blog.csdn.net/GJQI12/article/details/105118354)
