## Docker network

### Lists all the networks
docker network ls

三种网络类型
- bridge （容器默认选择）
- host
- none/null

### Displays network detail
docker network inspect {NewtorkName}


### bridge 概念
docker0 为 bridge network 的 gateway，容器能访问 gateway，所有容器的网络通过 gateway 与宿主机网络互通，容器也能同 bridge network 里面的容器网络互通。

Bridge 网络下总结
- 容器与容器互通
- 容器与宿主互通
- 容器与互联网互通（互联网访问容器时容器需要对外暴露端口关联）

bridge 网络可以创建多个，默认的不能删除
- 默认的 bridge 网络一般用来做练习使用
- 自定义的 bridge 网络比较适合生产环境，可以自动 DNS，比较适合单宿主多容器


### host 网络
默认只能有一个，不能创建多个，也不能删除

host 总结
- 不需要像 bridge 网络那样，每个容器如果要与宿主通信必须要关联端口，bridge 不需要，访问容器就像访问宿主的一个应用那样。


### none 网络
特点
- 隔离
- 不能联网


### 参考
- https://www.bilibili.com/video/BV1Aj411r71b
