## 常用命令

查看前十大的文件路径

-S

ls -alhS /opt/case/* | head -n 100

使用 nohup 命令

nohup command > /dev/null 2>&1 # doesn't create nohup.out

nohup command > /dev/null 2>&1 & # runs in background, still doesn't create nohup.out

nohup command > /your_ouput.log 2>&1 & # runs in background, doesn't create nohup.out, create log file

查看某一进程的所有线程

top -H -p <pid>

top -H -p 4361

### mount

**Linux mount挂载的作用，就是将一个设备（通常是存储设备）挂接到一个已存在的目录上。访问这个目录就是访问该存储设备。**

mount {device} {mountpoint}

umount {device} 或者 umount {mountpoint}

参考

- [https://www.zhihu.com/question/23550316](https://www.zhihu.com/question/23550316)
- [https://forum.ubuntu.org.cn/viewtopic.php?t=257333](https://forum.ubuntu.org.cn/viewtopic.php?t=257333)
- [https://juejin.cn/post/6864862445954039821](https://juejin.cn/post/6864862445954039821)


### Net client telnet and netcat
telnet 和 netcat 区别
- https://superuser.com/questions/1461609/what-is-the-difference-between-telnet-and-netcat

### lsof

常见的用法是查找应用程序打开的文件的名称和数目

lsof -p $pid

lsof -i:$port

查看进程监听的端口

lsof -nP -p $pid | grep LISTEN

### 查看Linux系统版本

查看内核版本

1. lsb_release (需额外安装)
    
    lsb_release -a
    
2. 只适用于 Redhat 系列
    
    cat /etc/redhat-release
    

查看发行版本

1. uname -a
2. cat /proc/version
