# API 设计

## 标准
OpenAPI
- https://www.openapis.org
- https://github.com/OAI
- https://github.com/OAI/OpenAPI-Specification

API design guide | Google Cloud
- https://cloud.google.com/apis/design/
- https://cloud.google.com/apis/design/errors

Microsoft REST API Guidelines
- https://github.com/microsoft/api-guidelines
- https://github.com/microsoft/api-guidelines/blob/master/Guidelines.md


## 最佳实践
### RESTful API
- https://restfulapi.net
- https://docs.github.com/en/rest
- https://docs.github.com/en/rest/guides
- RESTful API 设计指南 https://www.ruanyifeng.com/blog/2014/05/restful_api.html
- 实用 Web API 规范 https://blog.alswl.com/2023/04/web-api-guidelines/

### GraphQL
- https://graphql.org
- GraphQL Engine https://github.com/hasura/graphql-engine/
- GraphQL: From Excitement to Deception https://betterprogramming.pub/graphql-from-excitement-to-deception-f81f7c95b7cf
- GraphQL 为何没有火起来? https://www.zhihu.com/question/38596306
- GraphQL及元数据驱动架构在后端BFF中的实践 https://tech.meituan.com/2021/05/06/bff-graphql.html

### 介绍
- 前后端分离实践有感 https://juejin.cn/post/6844903545964249101
- API设计，从RPC、SOAP、REST到GraphQL
  - https://blog.enixjin.net/api_design_part1/
  - https://blog.enixjin.net/api_design_part2/
  - https://blog.enixjin.net/api_design_part3/
  - https://blog.enixjin.net/api_design_part4/
- REST vs. SOAP: Choosing the best web service https://www.techtarget.com/searchapparchitecture/tip/REST-vs-SOAP-Choosing-the-best-web-service
- “一把梭：REST API 全用 POST” https://coolshell.cn/articles/22173.html
- Architectural Styles and the Design of Network-based Software Architectures https://ics.uci.edu/~fielding/pubs/dissertation/top.htm

BFF
- The BFF Pattern (Backend for Frontend): An Introduction https://blog.bitsrc.io/bff-pattern-backend-for-frontend-an-introduction-e4fa965128bf

N+1 Problem in REST API
- https://restfulapi.net/rest-api-n-1-problem/
- https://learnku.com/laravel/t/15077/what-is-the-n1-problem-and-how-to-solve-the-n1-problem-in-laravel


## Swagger
- https://github.com/swagger-api
- https://github.com/swagger-api/swagger-editor
- https://github.com/swagger-api/swagger-ui
- https://github.com/swagger-api/swagger-codegen
- https://swagger.io
- https://editor.swagger.io
- https://swagger.io/tools/swagger-ui/
- https://swagger.io/tools/swagger-codegen/

介绍
- 如何编写基于OpenAPI规范的API文档 https://huangwenchao.gitbooks.io/swagger/content/


## 分页问题及其解决方案
问题1: 分页 page 越大，导致查询变慢
- 问题原因：
  - mysql 先查询 page*(page_size+1) 条数据，再过滤掉前面 page*page_size条数据，最后返回剩下 page_size 条数据
  - 虽然使用了二级索引，但是需要回表
- 解决方案：
  - 先用子查询查出符合条件的主键ID，再用主键ID做条件查出所有字段
  - 使用cursor 进行分页，不使用 page，page_size 模式
- 参考：https://www.zhihu.com/question/432910565

问题2: 滑动分页，新增数据导致丢失
- 解决方案：查询带 query_time，下拉翻页查询 query_time 不变，上拉刷新页面 query_time 重置。
- 参考：https://juejin.cn/post/7073519311213559822

问题3: 分页场景下，SQL 查询得到数据后，业务代码排除了部分内容，导致不足一页。
- 解决方案：
  - 1.客户端处理，加载非固定页大小的数据
  - 2.使用cursor 进行分页，不使用 page，page_size 模式
- 参考：
  - https://www.zhihu.com/question/619291617
  - https://www.v2ex.com/t/603295

问题4: 分页场景下，列表缓存策略。
- 解决方案
  - 1.直接缓存分页列表结果
  - 2.查询对象ID列表，再缓存每个对象条目
  - 3.缓存对象ID列表，同时缓存每个对象条目
- 参考：
  - https://www.cnblogs.com/makemylife/p/17425593.html
  - https://developer.aliyun.com/article/620529


## 工具
- Docgen - Transform your postman collection to HTML/Markdown documentation https://github.com/thedevsaddam/docgen
