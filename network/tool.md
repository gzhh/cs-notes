# 抓包工具使用

## 介绍

抓包概念介绍

- [https://www.uedbox.com/post/59475/](https://www.uedbox.com/post/59475/)

## tcpdump

## Wireshark

选择网卡

- 1.选择网卡
- 2.这时可以输入capture filter 来过滤，例如 host 127.0.0.1

PS：

wireshark抓取环回链路数据包

为什么wireshark抓包抓不到本机自己跟自己的通信包，因为本机发完本机的数据包不会经过网卡,而是经过环回链路返回本机，如果要监听环路链路，wireshark需要监听Loopback:lo0

进入主界面

- 1.输入过滤规则 display filter

捕捉过滤器（capture filter）和显示过滤器（Display filter）的区别

- 1.capture filter可以在抓包过程中将不符合过滤条件的包进行舍弃，只留复合过滤条件的包。而display filter是在已抓到的包中，将对应的包进行过滤，只显示满足条件的包。

参考

- [https://gitlab.com/wireshark/wireshark/-/wikis/home](https://gitlab.com/wireshark/wireshark/-/wikis/home)
- [https://www.cnblogs.com/moonbaby/p/10528401.html](https://www.cnblogs.com/moonbaby/p/10528401.html)

其他概念

- What is a TCP window update?
    
    [https://stackoverflow.com/questions/1466307/what-is-a-tcp-window-update](https://stackoverflow.com/questions/1466307/what-is-a-tcp-window-update)
    

## Fiddle

参考

- [https://docs.telerik.com/fiddler-everywhere/introduction](https://docs.telerik.com/fiddler-everywhere/introduction)

## Charles

参考

- [https://www.charlesproxy.com/documentation/](https://www.charlesproxy.com/documentation/)

## 对比

wireshark 更偏向tcp/ip层，可以抓http包内容，但是不支持解密抓到的https包加密的内容

fiddle更偏向应用层，支持http/https
