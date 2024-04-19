# Docker
- https://www.docker.com/
- https://github.com/moby/moby

最佳实践
- Guides https://docs.docker.com/get-started/overview/
- Manuals https://docs.docker.com/manuals/
- Reference https://docs.docker.com/reference/
- Samples https://docs.docker.com/samples/

工具
- https://www.docker.com/products/docker-desktop/
- https://docs.docker.com/desktop/
- https://orbstack.dev/
- https://github.com/abiosoft/colima

## 一、概念
What's the difference between Docker Compose and Kubernetes?
- https://stackoverflow.com/questions/47536536/whats-the-difference-between-docker-compose-and-kubernetes

### 主要原理
1. Linux namespace
> 隔离机制：使每个进程只看得到它自己的系统视图（文件、进程、网络接口、主机名等）
- 让容器初始化进程误以为它自己就是一号进程
- 所以说，容器其实就是一种特殊的进程，设置了 namespace 的进程
- 容器是一个“单进程”模型
    - 只有 PID = 1 的进程才会被 Dockerd 控制，即 pid=1 的进程挂了 Dockerd 能感知到，但是其他进程却不受 Dockerd 的管理，所以可能会出现孤儿进程。
    - 容器与应用同生命周期

2. Linux cgroup
> 资源限制：限制一个进程组能够使用的资源上限（CPU、内存、磁盘、网络带宽等）
- /proc 文件系统对容器做了 cgroup 资源限制不清楚，所以会影响到容器内执行 top 命令
- https://cizixs.com/2017/08/04/docker-resources-limit/

3. 容器镜像，也叫 rootfs（根文件系统）
> Change Root 切换进程的根目录

- Linux 使用 Mount Namespace 来实现 chroot 功能，使用它用来为容器进程提供隔离后执行环境的文件系统，这个文件系统也叫容器镜像，或 rootfs（根文件系统）
- rootfs 解决了本地开发环境与生产环境的一致性问题（操作系统级别的运行环境）
- rootfs 里打包的不只是应用，而是整个操作系统的文件和目录，也就意味着应用以及它运行所需要的所有依赖，都被封装在了一起（不包括对内核的依赖）
- Docker 层（Layer）概念
    - 用户制作镜像的每一步操作，都会生成一个层，也就是基于 base rootfs 生成一个增量 rootfs ，避免每次都创建一个全新的 rootfs
    - Layer 概念使用到了 联合文件系统 Union File System 技术
    - Docker 的 Layer 分为（TODO 深入理解）
        - 可读写层（容器层）
            - 在没有写入文件前，这个目录是空的；而一旦在容器里做了写操作，修改产生的内容就会以增量的方式出现在这个层中
            - 只读 + whiteout 用来隐藏进行删除操作的文件
            - 增删改操作发生的位置
        - init 层
            - 只做为容器初始化时写入的一些值，docker commit 时不会随读写层被提交
        - 只读层（镜像层）
            - 都是只读的（只读 + whiteout  ）
        - 镜像文件的修改：
            - 所有的增删改都只会作用在容器层，相同的文件上层会覆盖掉下层
            - 首先会从上到下查找有没有这个文件，如果找到了，就复制到容器层中进行修改，修改的结果就会作用到下层的文件，这种方式也被称为 copy-on-write

### 其他概念

- docker build & docker commit 原理
- docker exec 命令实现原理
    - 一个进程可以选择加入到某个进程已有的 Namespace 当中，从而达到进入这个进程所在容器的目录
- docker volume 原理（TODO 深入理解）

### 小结

1. 一个正在运行的容器，可以被一分为二的看待
    - rootfs 容器镜像，容器的静态视图
    - namespace + cgroup，容器的动态视图（runtime）
2. 容器编排技术的发展
    - 从单个容器到容器云的跨越


## 二、常见操作
### docker context
查看 context 子命令
> docker context -h

列出所有可用 context
> docker context ls

查看当前 context
> docker context show

设置新的 context
> docker context use {CONTEXT_NAME}


### 磁盘及清除相关
查看磁盘使用
> docker system df

清除所有 hang <none> 镜像
> docker image prune

清除未使用的容器
> docker container prune

清除所有无用的数据卷
> docker volume prune

清除未使用的构建缓存
> docker builder prune

清除未使用数据
> docker system prune
