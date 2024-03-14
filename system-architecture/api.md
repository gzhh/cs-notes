# api 设计

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
