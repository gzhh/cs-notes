# Wireshark
- https://www.wireshark.org/

Learn
- https://www.wireshark.org/docs/
- Wireshark User’s Guide https://www.wireshark.org/docs/wsug_html_chunked/index.html
- Wireshark Wiki https://wiki.wireshark.org
- Questions - Ask Wireshark https://ask.wireshark.org/questions/
- Wireshark Mailing Lists https://www.wireshark.org/lists/
- Wireshark Wiki https://gitlab.com/wireshark/wireshark/-/wikis/home

Books
- Wireshark网络分析就这么简单 https://book.douban.com/subject/26268767/
- wireshark网络分析的艺术 https://book.douban.com/subject/26710788/
- Wireshark数据包分析实战(第3版) https://book.douban.com/subject/30387220/


## Usage
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

- wireshark怎么抓包、wireshark抓包详细图文教程 https://www.cnblogs.com/moonbaby/p/10528401.html

其他概念

- What is a TCP window update? https://stackoverflow.com/questions/1466307/what-is-a-tcp-window-update


