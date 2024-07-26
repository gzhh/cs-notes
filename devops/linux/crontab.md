# crontab


## 环境变量
crontab 中执行的任务读取的环境变量可能和 shell 中读取的不一样，最好在 crontab 文件中显示设置环境变量

例如在文件顶部设置
```
PATH=/usr/local/bin:/usr/bin:/bin

RUN_MODE=prod
```


## 问题排查
- 分享一次Linux任务计划crontab不执行的问题排查过程 https://zhang.ge/5093.html

测试任务
```
* * * * * /bin/echo 'Hello' >> /tmp/test.log 2>&1
```
