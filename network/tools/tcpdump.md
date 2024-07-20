# tcpdump
- https://en.wikipedia.org/wiki/Tcpdump
- https://github.com/the-tcpdump-group/tcpdump
- https://www.tcpdump.org
- https://www.tcpdump.org/manpages/tcpdump.1.html

最佳实践
- Let's learn tcpdump! https://wizardzines.com/zines/tcpdump/
- A tcpdump Tutorial with Examples https://danielmiessler.com/p/tcpdump/
- Tcpdump advanced filters https://blog.wains.be/2007/2007-10-01-tcpdump-advanced-filters/
- tcpdump for Dummies http://www.alexonlinux.com/tcpdump-for-dummies
- How to capture and analyze traffic with tcpdump https://www.techtarget.com/searchnetworking/tutorial/How-to-capture-and-analyze-traffic-with-tcpdump
- https://wangchujiang.com/linux-command/c/tcpdump.html
- 抓包神器 tcpdump 使用介绍 https://cizixs.com/2015/03/12/tcpdump-introduction/
- Linux tcpdump命令详解 https://www.cnblogs.com/ggjucheng/archive/2012/01/14/2322659.html
- TCPdump: A Powerful Network Packet Analyzer A Thread with examples https://x.com/devops_tech/status/1675418922282283009


## 常用参数
- Cheet sheet https://packetlife.net/media/library/12/tcpdump.pdf

`-D` 
- Display available interfaces.

`-i`
- Select an interface.

`-w`
- Write to a file (use the .pcap extension).

`-r`
- Read from a file with tcpdump.

`-n`
- Don't convert addresses.
- 主要用于提高捕获效率和避免不必要的 DNS 查询，非常适合在需要快速捕获和分析大量数据包的场景下使用。

`-S`
- Print absolute, rather than relative, TCP sequence numbers.
