## launchctl 服务

可以管理服务的开机自动开启，类似于 Linux 下的 systemd

### 例子
```
1.显示当前的启动脚本
launchctl list

2.开机时自动启动 MySQL 服务器
sudo launchctl load -w /System/Library/LaunchDaemons/com.oracle.oss.mysql.mysqld.plist

3.设置开机启动并立即启动改服务
launchctl load -w **.plist

4. 设置开机启动但不立即启动服务
launchctl load **.plist

5. 停止正在运行的启动脚本
sudo launchctl unload **.plist

6. 再加上-w 选项即可去除开机启动
sudo launchctl unload -w **.plist

7. start 选项

8. stop 选项

9. enable 选项
```

