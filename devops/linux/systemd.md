## Systemd
- https://github.com/systemd/systemd
- https://www.freedesktop.org/software/systemd/man/latest/systemd
- https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html
- https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html
- https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html
- https://www.freedesktop.org/software/systemd/man/latest/systemd.path.html

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
