## Systemd
- https://github.com/systemd/systemd
- https://www.freedesktop.org/software/systemd/man/latest/systemd
- https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html
- https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html
- https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html
- https://www.freedesktop.org/software/systemd/man/latest/systemd.path.html

### Systemctl
systemctl log
```
systemctl status <service>：查看服务的状态和活动日志。

journalctl -u <service>：查看服务的所有日志，包括活动日志和错误日志。

journalctl -u <service> -f：实时查看服务的日志。

journalctl --since "YYYY-MM-DD HH:MM:SS" --until "YYYY-MM-DD HH:MM:SS"：查看指定时间范围内的所有日志。

journalctl -p <priority>：查看指定优先级的日志。优先级可以是debug、info、notice、warning、err、crit、alert或emerg之一。
```

### Example
- https://learnku.com/articles/34025

```
# cat /etc/systemd/system/test-gitlabci.path
[Unit]
Description = monitor test-gitlabci app

[Path]
PathChanged = /opt/test-gitlabci
PathModified = /opt/test-gitlabci/app
Unit = test-gitlabci.service

# cat /etc/systemd/system/test-gitlabci.service
[Unit]
Description=Test gitlab-ci

[Service]
Type=simple
WorkingDirectory=/opt/test-gitlabci
ExecStart=/opt/test-gitlabci/test-gitlabci

[Install]
WantedBy=multi-user.target
```

### Usage Ref
- https://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html
- https://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-part-two.html
- https://www.digitalocean.com/community/tutorials/understanding-systemd-units-and-unit-files
- https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/system_administrators_guide/chap-managing_services_with_systemd
- https://dunwu.github.io/linux-tutorial/linux/ops/systemd.html


### Environment variables Ref
-  https://www.jibing57.com/2022/10/09/environment-varilables-in-systemd-service/

### Path unit
- https://unix.stackexchange.com/questions/600642/systemd-path-how-to-tell-which-pathchanged
- https://unix.stackexchange.com/questions/644002/multiple-instances-of-unit-in-path-or-timer-unit
