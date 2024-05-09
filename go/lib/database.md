# Database Library
- https://pkg.go.dev/database/sql
- http://go-database-sql.org
- https://go.dev/doc/database/manage-connections
- https://go.dev/doc/tutorial/database-access


## 讨论
- 为什么要旗帜鲜明地反对 orm 和 sql builder https://mp.weixin.qq.com/s/5DIRKlWpr3pwx2h5YJhnAQ


## 非 ORM
- sqlx https://github.com/jmoiron/sqlx
- sqlc https://github.com/sqlc-dev/sqlc


## ORM & SQL Builder
### GORM
- https://github.com/go-gorm/gorm
- https://gorm.io
- https://gorm.io/docs/

深入理解
- gorm 框架原理&源码解析 https://mp.weixin.qq.com/s/VmAiRts3Mbbn0VsVDpP3Lg

Gen Tool
- https://gorm.io/gen/gen_tool.html
- 表结构转为 struct
  - `gentool -dsn "root:123456@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local" -tables "goods,test" --onlyModel=true -outPath "./modelDir"`

GORM Gen
- https://github.com/go-gorm/gen
- https://gorm.io/gen/
- GORM 强大的代码生成工具 —— gorm/gen https://juejin.cn/post/7133150674400837668

坑
- 哔了狗的 gorm. 很久没用 golang 了，生态这么操蛋么？https://v2ex.com/t/859178

### XORM
- https://github.com/go-xorm/xorm
- XORM 操作指南 https://books.studygolang.com/xorm/
- 表结构转为 struct
  - `xorm reverse mysql "user:password@(host:port)/dbname?charset=utf8" templates/goxorm your_model_path`
  - `xorm reverse mysql "user:password@(host:port)/dbname?charset=utf8" $GOPATH/src/github.com/go-xorm/cmd/xorm/templates/goxorm your_model_path`

常见问题
- SetMapper
    ```
    使用 conn.SetMapper(names.GonicMapper{})

    经常会出现 fatal error: sync: unlock of unlocked mutex

    解决办法生成结构体时使用，conn.SetMapper(names.GonicMapper{})，而不是每次操作时调用
    ```

### ent
- https://github.com/ent/ent

### SQLBoiler
- https://github.com/volatiletech/sqlboiler

### bun
- uptrace/bun https://github.com/uptrace/bun

### SQL builder for Go
- https://github.com/huandu/go-sqlbuilder


## SQL migration
- migrate https://github.com/golang-migrate/migrate
- sql-migrate https://github.com/rubenv/sql-migrate


## Third Party Tool
- SOAR: SQL Optimizer And Rewriter https://github.com/XiaoMi/soar
