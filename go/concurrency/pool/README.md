### Pool
原理
- https://pkg.go.dev/sync#Pool
- https://www.pursuitking.com/go_2/1-10.html

最佳实践
- 由浅入深聊聊Golang的sync.Pool https://juejin.cn/post/6844903903046320136
- Go sync.Pool 保姆级教程 https://juejin.cn/post/6989798306440282148
- Go sync.Pool and the Mechanics Behind It https://victoriametrics.com/blog/go-sync-pool/
- https://github.com/panjf2000/ants
  - 慢聊Golang协程池Ants实现原理 https://mp.weixin.qq.com/s/fZpPkG-C0wZ5Z45H2aUxAA
- https://github.com/fatih/pool

MySQL 连接池
- 标准库 database/sql 配置参数
  - SetMaxOpenConns
  - SetMaxIdleConns
  - SetConnMaxLifetime
  - SetConnMaxIdleTime
- Golang连接池的几种实现案例 https://juejin.cn/post/6844904077386596366
- Configuring sql.DB for Better Performance
  - https://www.alexedwards.net/blog/configuring-sqldb
  - https://colobu.com/2020/05/18/configuring-sql-DB-for-better-performance-2020/
- 详解连接池参数设置（边调边看）https://juejin.cn/post/7111500846575124488

Redis 连接池
- 
