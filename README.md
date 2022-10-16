<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Notes](#notes)
  - [Golang](#golang)
    - [Golang Notes](#golang-notes)
    - [Trouble Shooting](#trouble-shooting)
  - [Operating System](#operating-system)
  - [Networking](#networking)
  - [Database](#database)
  - [DevOps](#devops)
    - [Automation](#automation)
    - [Docker](#docker)
    - [Kubernetes](#kubernetes)
    - [CI/CD](#cicd)
  - [Cache](#cache)
  - [Message Queue](#message-queue)
  - [Design Pattern](#design-pattern)
  - [System Design](#system-design)
    - [系统设计的大框架](#%E7%B3%BB%E7%BB%9F%E8%AE%BE%E8%AE%A1%E7%9A%84%E5%A4%A7%E6%A1%86%E6%9E%B6)
    - [常见设计方案](#%E5%B8%B8%E8%A7%81%E8%AE%BE%E8%AE%A1%E6%96%B9%E6%A1%88)
  - [Distributed System](#distributed-system)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Notes 

## Golang
### [Golang Notes](https://github.com/gzhh/golang-notes/tree/main/src)
### [Trouble Shooting](https://github.com/gzhh/golang-notes/tree/main/src/debug/README.md)

## Operating System

## Networking

## Database
[MySQL](https://github.com/gzhh/golang-notes/tree/main/database/README.md)

## DevOps
### [Automation]()
### [Docker]()
### [Kubernetes]()
### [CI/CD]()

## Cache

## Message Queue

## Design Pattern
设计模式
- [创建型](https://github.com/gzhh/golang-notes/tree/main/design-patterns/creational)
- [行为型](https://github.com/gzhh/golang-notes/tree/main/design-patterns/behavioral)
- [结构型](https://github.com/gzhh/golang-notes/tree/main/design-patterns/structural)

参考
- https://en.wikipedia.org/wiki/Software_design_pattern
- https://github.com/senghoo/golang-design-pattern
- https://github.com/tmrts/go-patterns

## System Design
### 系统设计的大框架
- 充分理解问题并确定设计范围
- 提出整体设计设计然后再继续
- 按优先级设计各个模块的细节
- 总结归纳及拓展

### 常见设计方案
- [Rate Limiter 限流器](https://github.com/gzhh/golang-notes/tree/main/system-design/rate-limiter.md)
- [Consistent Hashing 一致性哈希](https://github.com/gzhh/golang-notes/tree/main/system-design/consistent-hashing.md)
- [Unique ID Generator 唯一ID生成器](https://github.com/gzhh/golang-notes/tree/main/system-design/unique-id-generator.md)
- [URL SHORTENER URL短链](https://github.com/gzhh/golang-notes/tree/main/system-design/url-shortener.md)
- [Key-Value Store 键值存储](system-design/key-value-store.md)

偏场景
- [秒杀系统](system-design/spike-system.md)
- Web 爬虫
- 权限系统
- 任务调度系统
- 抢红包
- 排行榜
- 点赞功能
- [Notification System 通知系统](system-design/notification-system.md)
- News Feed System 新闻流系统
- Chat System 聊天系统
- Search Auto-Complete System 搜索自动补全系统

偏海量数据处理
- 查找两个大文件中的相同记录 [ref1](https://blog.csdn.net/Fly_as_tadpole/article/details/88375809) [ref2](https://www.zhihu.com/question/21827402)

TODO 偏技术
- 设计 Web Framework
- 设计 ORM
- 设计 RPC
- 设计 Web Crawler

参考
- [System Design Interview – An Insider's Guide](https://www.goodreads.com/book/show/54109255-system-design-interview-an-insider-s-guide)
- [System Design Interview – An Insider's Guide: Volume 2](https://www.goodreads.com/book/show/60631342-system-design-interview-an-insider-s-guide)
- [ByteByteGo System Design](https://blog.bytebytego.com)
- [Grokking the System Design Interview](https://www.educative.io/courses/grokking-the-system-design-interview)
- [极客兔兔：七天用Go从零实现系列](https://geektutu.com/post/gee.html)
- [李智慧 · 高并发架构实战课](https://time.geekbang.org/column/intro/100105701)

## Distributed System
CAP 理论
- [CAP theorem](https://en.wikipedia.org/wiki/CAP_theorem)
- [一文看懂｜分布式系统之CAP理论](https://cloud.tencent.com/developer/article/1860632)
- [一文吃透分布式理论【CAP,BASE深入解析】](https://juejin.cn/post/7021717177220726798)
