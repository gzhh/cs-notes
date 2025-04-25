# epoll


### epoll 基本步骤
- 创建epoll
  - `int epoll_create(int size);`
  - 在内核中实则创建了一棵红黑树(平衡二叉树)的根节点root，这个根节点的关系与epfd是相对应的。
- 控制epoll
  - `int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event);`
  - 创建一个用户态的事件，绑定到某个fd上，然后添加到内核中的epoll红黑树中。
- 等待epoll
  - `int epoll_wait(int epfd, struct epoll_event *event, int maxevents, int timeout);`
  - epoll_wait是一个阻塞的状态，如果内核检测到I/O的读写响应，则会抛给上层的 epoll_wait，返给用户态一个已经触发的事件队列，同时阻塞返回。开发者可以从队列中取出事件来处理，事件中有对应的fd具体是哪一个。


### epoll编程主流结构
```
int epfd = epoll_create(1000);

//将 listen_fd 添加进 epoll 中
epoll_ctl(epfd, EPOLL_CTL_ADD, listen_fd, &listen_event);

while (1) {
    //阻塞等待 epoll 中 的 fd 触发
    int active_cnt = epoll_wait(epfd, events, 1000, -1);

    for (i = 0; i < active_cnt; i++) {
        if (evnets[i].data.fd == listen_fd) {
            //accept，并且将新 accept 的 fd 加入 epoll 中
        }
        else if (events[i].events & EPOLLIN) {
            //对该 fd 进行读操作
        }
        else if (events[i].events & EPOLLOUT) {
            //对该 fd 进行写操作
        }
    }
}
```

