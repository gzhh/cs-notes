## tmux

### **介绍**

登录远程服务器，有时候需要操作多个终端，如果又要用 ssh 去连接服务器很麻烦，这时我们可以使用终端复用工具 tmux，连接远程服务器后开启 tmux 服务，就能在一个终端中开启多个分屏会话了

### **列了些基本的使用命令**

开启 tmux : tmux

查看 tmux 终端：tmux list-session

切换终端编号：tmux a -t number

关闭所有终端：tmux kill-server

使用 tmux 命令 :

首先 - 输入 Ctrl + b 前缀

然后 - 输入具体命令

横向切屏：前缀 + "

竖向切屏：前缀 + %

光标切换：前缀 + 方向键

光标切换到上一个窗格：前缀 + ;

光标切换到下一个窗格：前缀 + o

### **参考链接**
- Tmux 使用教程 https://www.ruanyifeng.com/blog/2019/10/tmux.html

[https://www.cnblogs.com/kevingrace/p/6496899.html](https://www.cnblogs.com/kevingrace/p/6496899.html)

[https://gist.github.com/ryerh/14b7c24dfd623ef8edc7](https://gist.github.com/ryerh/14b7c24dfd623ef8edc7)

[https://tmuxcheatsheet.com/](https://tmuxcheatsheet.com/)
