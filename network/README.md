# 计算机网络
[https://zh.wikipedia.org/wiki/计算机网络](https://zh.wikipedia.org/wiki/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C)

## 网络组成

1. 网络连接
    
    物理传输介质（属于 OSI 模型的物理层及数据链路层）
    
    以太网是局域网的主流传输介质技术，以太网可以铜线或光纤电缆传输数据。无线局域网则一般会以无线电作传输介质。
    
    - 有线技术（例如双绞线、光纤等）
    - 无线技术（例如通信卫星、蜂窝网络、无线电与扩频技术 Wi-Fi 等）
    
2. 网络节点
    - 网络接口（即网卡，使得电脑能够访问传输介质上的资料，以太网中网卡都有唯一的 MAC 地址）
    - ~~中继器及集线器~~（用于增强信号的网络设备，由于交换器的功能相对较佳，价格亦相近，故两者皆已被交换器所取代）
    - ~~桥接器~~（桥接器连接两个独立的网段及过滤之间的流量，它在OSI模型的数据链路层中运作）
    - 交换器（一种依据MAC地址，来在端口之间转发和过滤数据链路层的帧的设备）
    - 路由器（路由器是一款互连网络设备，兼具了中继器、桥接器、集线器的功能。其依照数据包内的消息及路由表中的信息来选择数据包传递的路径。它必须拥有IP地址才可正常运作）
    - 调制解调器（即光猫，把节点的信号转换成其他非专用线路能够发送的信号。将数字信号调变到模拟信号上进行传输，并解调收到的模拟信号以得到数字信号的电子设备）
    - 防火墙（一种控制网络安全和访问规则的网络系统。它按特定规则来充许或阻止资料通过）


## 网络协议

### OSI七层模型

[https://zh.wikipedia.org/wiki/OSI模型](https://zh.wikipedia.org/wiki/OSI%E6%A8%A1%E5%9E%8B)

- 应用层
- 表示层
- 会话层
- 传输层
- 网络层
- 数据链路层
- 物理层

### TCP/IP四层模型

[https://zh.wikipedia.org/wiki/TCP/IP协议族](https://zh.wikipedia.org/wiki/TCP/IP%E5%8D%8F%E8%AE%AE%E6%97%8F)

- 应用层
- 传输层
- 网络层
- 链路层

### 常用的协议

**应用层**

运行在[TCP](https://zh.wikipedia.org/wiki/%E4%BC%A0%E8%BE%93%E6%8E%A7%E5%88%B6%E5%8D%8F%E8%AE%AE)协议上的协议：
- [HTTP](https://zh.wikipedia.org/wiki/%E8%B6%85%E6%96%87%E6%9C%AC%E4%BC%A0%E8%BE%93%E5%8D%8F%E8%AE%AE)（Hypertext Transfer Protocol，超文本传输协议），主要用于普通浏览。
- [HTTPS](https://zh.wikipedia.org/wiki/%E8%B6%85%E6%96%87%E6%9C%AC%E4%BC%A0%E8%BE%93%E5%AE%89%E5%85%A8%E5%8D%8F%E8%AE%AE)（Hypertext Transfer Protocol over Secure Socket Layer, or HTTP over SSL，安全超文本传输协议）,HTTP协议的安全版本。
- [FTP](https://zh.wikipedia.org/wiki/%E6%96%87%E4%BB%B6%E4%BC%A0%E8%BE%93%E5%8D%8F%E8%AE%AE)（File Transfer Protocol，文件传输协议），由名知义，用于文件传输。

- [TELNET](https://zh.wikipedia.org/wiki/Telnet)（Teletype over the Network，网络电传），通过一个终端（terminal）登陆到网络。
- [SSH](https://zh.wikipedia.org/wiki/Secure_Shell)（Secure Shell，用于替代安全性差的[TELNET](https://zh.wikipedia.org/wiki/TELNET)），用于加密安全登陆用。
- [POP3](https://zh.wikipedia.org/wiki/%E9%83%B5%E5%B1%80%E5%8D%94%E5%AE%9A)（Post Office Protocol, version 3，邮局协议），收邮件用。
- [SMTP](https://zh.wikipedia.org/wiki/%E7%AE%80%E5%8D%95%E9%82%AE%E4%BB%B6%E4%BC%A0%E8%BE%93%E5%8D%8F%E8%AE%AE)（Simple Mail Transfer Protocol，简单邮件传输协议），用来发送电子邮件。

运行在[UDP](https://zh.wikipedia.org/wiki/%E7%94%A8%E6%88%B7%E6%95%B0%E6%8D%AE%E6%8A%A5%E5%8D%8F%E8%AE%AE)协议上的协议：
- [DHCP](https://zh.wikipedia.org/wiki/%E5%8A%A8%E6%80%81%E4%B8%BB%E6%9C%BA%E8%AE%BE%E7%BD%AE%E5%8D%8F%E8%AE%AE)（Dynamic Host Configuration Protocol，动态主机配置协议），动态配置IP地址。

其他：
- [DNS](https://zh.wikipedia.org/wiki/%E5%9F%9F%E5%90%8D%E7%B3%BB%E7%BB%9F)（Domain Name Service，域名服务），用于完成地址查找，邮件转发等工作（运行在[TCP](https://zh.wikipedia.org/wiki/%E4%BC%A0%E8%BE%93%E6%8E%A7%E5%88%B6%E5%8D%8F%E8%AE%AE)和[UDP](https://zh.wikipedia.org/wiki/%E7%94%A8%E6%88%B7%E6%95%B0%E6%8D%AE%E6%8A%A5%E5%8D%8F%E8%AE%AE)协议上）。
- [ECHO](https://zh.wikipedia.org/w/index.php?title=ECHO&action=edit&redlink=1)（Echo Protocol，回绕协议），用于查错及测量应答时间（运行在[TCP](https://zh.wikipedia.org/wiki/%E4%BC%A0%E8%BE%93%E6%8E%A7%E5%88%B6%E5%8D%8F%E8%AE%AE)和[UDP](https://zh.wikipedia.org/wiki/%E7%94%A8%E6%88%B7%E6%95%B0%E6%8D%AE%E6%8A%A5%E5%8D%8F%E8%AE%AE)协议上）。
- [SNMP](https://zh.wikipedia.org/wiki/%E7%AE%80%E5%8D%95%E7%BD%91%E7%BB%9C%E7%AE%A1%E7%90%86%E5%8D%8F%E8%AE%AE)（Simple Network Management Protocol，简单网络管理协议），用于网络信息的收集和网络管理。

**表示层**
- SSL/TLS

**会话层**
- RPC [https://zh.wikipedia.org/wiki/遠程過程調用](https://zh.wikipedia.org/wiki/%E9%81%A0%E7%A8%8B%E9%81%8E%E7%A8%8B%E8%AA%BF%E7%94%A8)

**传输层**
- TCP
- UDP

**网络层**
- IPv4
- IPV6
- ICMP [https://zh.wikipedia.org/wiki/互联网控制消息协议](https://zh.wikipedia.org/wiki/%E4%BA%92%E8%81%94%E7%BD%91%E6%8E%A7%E5%88%B6%E6%B6%88%E6%81%AF%E5%8D%8F%E8%AE%AE)

对应应用层协议（ping、traceroute）

**链路层**
- [ARP](https://zh.wikipedia.org/wiki/%E5%9C%B0%E5%9D%80%E8%A7%A3%E6%9E%90%E5%8D%8F%E8%AE%AE)（Address Resolution Protocol，地址解析协议），用于动态解析以太网硬件的地址。

### TCP/UDP端口列表

[https://zh.wikipedia.org/wiki/TCP/UDP端口列表](https://zh.wikipedia.org/wiki/TCP/UDP%E7%AB%AF%E5%8F%A3%E5%88%97%E8%A1%A8)

# 常见端口

### 一、常见端口号

熟知：0 - 1023

20、21 FTP

22 SSH

23 TELNET

25 SMTP

53 DNS

80 HTTP

443 HTTPS

注册：1024 - 49151

3306 MySQL

6379 Redis

11211 Memcached

27017 MongoDB

动态：49152 - 65535

### ICMP

internet控制报文协议（位于TCP/IP 四层 模型中的网络层）

传输层 → TCP、UDP

> ICMP（诊断）IPV4、IPV6

网络层 → IP

icmp报文一般是携带在ip数据报中的

icmp报文类型

1. 差错类
2. 信息类

icmp对ip提供了功能

1. 回显请求（重定向路由）
2. 应答
3. 超时

**对应应用层命令**

ping

发送icmp会送请求消息给目的主机，目的主机返回icmp回送应答消息给源主机，如果源主机在一定时间内收到应答，则认为主机可答

用途：

计算间隔时间，有多少个包被送达，检测主机间是否能通讯

其他：

ping请求数据报在经过每一个路由器的时候，路由器都会把自己的IP放到该数据报中，然后目的主机会把这个ip列表复制到icmp数据包中发回主机（IP头大小有限）

