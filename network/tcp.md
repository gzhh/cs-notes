# TCP
### TCP状态码

![tcp状态转换](images/tcp_status.png)

[https://zh.wikipedia.org/wiki/传输控制协议#状态编码](https://zh.wikipedia.org/wiki/%E4%BC%A0%E8%BE%93%E6%8E%A7%E5%88%B6%E5%8D%8F%E8%AE%AE#%E7%8A%B6%E6%80%81%E7%BC%96%E7%A0%81)

LISTEN

SYN-SEND

SYN-RECEIVED

ESTABLISHED

FIN-WAIT-1

FIN-WAIT-2

CLOSE-WAIT

CLOSING

LAST-ACK

TIME-WAIT

CLOSED

### TCP状态转化


### TCP报文

[https://zh.wikipedia.org/wiki/传输控制协议#封包結構](https://zh.wikipedia.org/wiki/%E4%BC%A0%E8%BE%93%E6%8E%A7%E5%88%B6%E5%8D%8F%E8%AE%AE#%E5%B0%81%E5%8C%85%E7%B5%90%E6%A7%8B)

### 滑动窗口

### 超时重传

### Ref
- [https://zh.wikipedia.org/wiki/传输控制协议](https://zh.wikipedia.org/wiki/%E4%BC%A0%E8%BE%93%E6%8E%A7%E5%88%B6%E5%8D%8F%E8%AE%AE)
- [https://zhuanlan.zhihu.com/p/53374516](https://zhuanlan.zhihu.com/p/53374516)
- [https://blog.csdn.net/imilli/article/details/50620104](https://blog.csdn.net/imilli/article/details/50620104)

## TCP 常见问题
### connection reset by peer
- https://stackoverflow.com/questions/1434451/what-does-connection-reset-by-peer-mean

### Ref
- https://plantegg.github.io/2021/02/14/TCP疑难问题案例汇总/


### TCP 状态码

一条tcp连接，主动关闭的一方不可能出现的连接状态是(A)

A.CLOSE_WAIT

B.FIN_WAIT2

C.TIME_WAIT

D.FIN_WAIT1

TCP链接中主动断开链接netstat观察可能出现的状态流转是（CD）

A.ESTABLISHED->CLOSE_WAIT->TIME_WAIT->CLOSED

B.ESTABLISHED->TIME_WAIT->CLOSE_WAIT->CLOSED

C.ESTABLISHED->FIN_WAIT_1->FIN_WAIT_2->TIME_WAIT->CLOSED

D.ESTABLISHED->FIN_WAIT_1->TIME_WAIT->CLOSED

### TCP设计问题

**【问题1】为什么连接的时候是三次握手，关闭的时候却是四次握手？**

答：因为当Server端收到Client端的SYN连接请求报文后，可以直接发送SYN+ACK报文。其中ACK报文是用来应答的，SYN报文是用来同步的。但是关闭连接时，当Server端收到FIN报文时，很可能并不会立即关闭SOCKET，所以只能先回复一个ACK报文，告诉Client端，"你发的FIN报文我收到了"。只有等到我Server端所有的报文都发送完了，我才能发送FIN报文，因此不能一起发送。故需要四步握手。

**【问题2】为什么TIME_WAIT状态需要经过2MSL(最大报文段生存时间)才能返回到CLOSE状态？**

答：虽然按道理，四个报文都发送完毕，我们可以直接进入CLOSE状态了，但是我们必须假象网络是不可靠的，有可能最后一个ACK丢失。所以TIME_WAIT状态就是用来重发可能丢失的ACK报文。

整个客户端所经历的状态：
