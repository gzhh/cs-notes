## Canal
Github 介绍
- https://github.com/alibaba/canal

主要用途
- 基于 MySQL 数据库增量日志解析，提供增量数据订阅和消费。

业务场景
- 数据库镜像
- 数据库实时备份
- 索引构建和实时维护(拆分异构索引、倒排索引等)
- 业务 cache 刷新
- 带业务逻辑的增量数据处理

工作原理
- MySQL主备复制原理
  - MySQL master 将数据变更写入二进制日志( binary log, 其中记录叫做二进制日志事件binary log events，可以通过 show binlog events 进行查看)
  - MySQL slave 将 master 的 binary log events 拷贝到它的中继日志(relay log)
  - MySQL slave 重放 relay log 中事件，将数据变更反映它自己的数据
- canal 工作原理
  - canal 模拟 MySQL slave 的交互协议，伪装自己为 MySQL slave ，向 MySQL master 发送dump 协议
  - MySQL master 收到 dump 请求，开始推送 binary log 给 slave (即 canal )
  - canal 解析 binary log 对象(原始为 byte 流)

client-server 模式
- 交互协议使用 protobuf 3.0 , client 端可采用不同语言实现不同的消费逻辑
- https://github.com/alibaba/canal?tab=readme-ov-file#多语言
- https://github.com/withlin/canal-go


最佳实践参考
- [Canal介绍和使用指南](https://www.liwenzhou.com/posts/Go/canal/)
- [滴滴基于Binlog的采集架构与实践](https://mp.weixin.qq.com/s/2gSxNTqqsDMYHFWRWfM1Fg) 


### Canal 最佳实践 - Docker + canal-go
参考
- https://github.com/alibaba/canal/wiki/QuickStart
- https://github.com/alibaba/canal/wiki/Docker-QuickStart
- https://developer.aliyun.com/article/892293

步骤
- 修改 MySQL 配置文件 my.cnf，开启 Binlog 写入功能，配置 binlog-format 为 ROW 模式。
    ```
    [mysqld]
    log-bin=mysql-bin # 开启 binlog
    binlog-format=ROW # 选择 ROW 模式
    server_id=1 # 配置 MySQL replaction 需要定义，不要和 canal 的 slaveId 重复
    
    -- 验证配置是否生效
    show variables like 'binlog_format';
    show variables like 'log_bin';
    show master status;
    ```
- 授权 canal 链接 MySQL 账号具有作为 MySQL slave 的权限, 如果已有账户可直接 grant
    ```
    CREATE USER canal IDENTIFIED BY 'canal';  
    GRANT SELECT, REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'canal'@'%';
    -- GRANT ALL PRIVILEGES ON *.* TO 'canal'@'%' ;
    FLUSH PRIVILEGES;
    
    -- 验证用户授权是否正确
    show grants for 'canal'@'%';
    ```
- 运行 docker canal-server 服务
    ```
    # 下载脚本
    wget https://raw.githubusercontent.com/alibaba/canal/master/docker/run.sh 

    # 构建一个destination name为test的队列
    sh run.sh -e canal.auto.scan=false \
              -e canal.destinations=test \
              -e canal.instance.master.address=127.0.0.1:13306  \
              -e canal.instance.dbUsername=canal  \
              -e canal.instance.dbPassword=canal  \
              -e canal.instance.connectionCharset=UTF-8 \
              -e canal.instance.tsdb.enable=true \
              -e canal.instance.gtidon=false  \
    ```

- 或者使用 docker-compose 运行 docker canal-server 服务
    - mkdir canal-server & cd canal-server
    - touch docker-comopse.yml
        ```
        version: '3'
        services:
          canal-server:
            image: canal/canal-server:v1.1.4
            container_name: canal-server
            restart: unless-stopped
            network_mode: host
            ports: 
              - 11111:11111
            environment:
              - canal.auto.scan=false
              - canal.instance.master.address=127.0.0.1:13306
              - canal.instance.dbUsername=canal
              - canal.instance.dbPassword=canal
              - canal.instance.filter.regex=.*\\..*
              - canal.destinations=test
              - canal.instance.connectionCharset=UTF-8
              - canal.instance.tsdb.enable=true
            volumes:
              - ./log/:/home/admin/canal-server/logs/
        ```
    - 启动服务
        ```
        docker-compose up

         ✔ Container canal-server  Created                                                                         0.0s
            Attaching to canal-server
            canal-server  | DOCKER_DEPLOY_TYPE=VM
            canal-server  | ==> INIT /alidata/init/02init-sshd.sh
            canal-server  | ==> EXIT CODE: 0
            canal-server  | ==> INIT /alidata/init/fix-hosts.py
            canal-server  | ==> EXIT CODE: 0
            canal-server  | ==> INIT DEFAULT
            canal-server  | Starting sshd: [  OK  ]
            canal-server  | Starting crond: [  OK  ]
            canal-server  | ==> INIT DONE
            canal-server  | ==> RUN /home/admin/app.sh
            canal-server  | ==> START ...
            canal-server  | start canal ...
            canal-server  | start canal successful
            canal-server  | ==> START SUCCESSFUL ...
            ```
- Canal client
    - 使用 client canal-go 连接 canal-server 进行消费数据
    - https://github.com/withlin/canal-go/blob/master/docker/docker-compose.yml
    ```
    package main

    // mysql binlog -> canal -> kafka -> go consumer

    import (
        "fmt"
        "log"
        "time"

        "github.com/withlin/canal-go/client"
        pbe "github.com/withlin/canal-go/protocol/entry"
        "google.golang.org/protobuf/proto"
    )

    func main() {
        connector := client.NewSimpleCanalConnector("127.0.0.1", 11111, "", "", "test", 60000, 60*60*1000)
        err := connector.Connect()
        if err != nil {
            log.Fatal(err)
        }

        // mysql 数据解析关注的表-所有表：.*   or  .*\\..*
        err = connector.Subscribe(".*\\..*")
        if err != nil {
            log.Fatal(err)
        }

        for {
            message, err := connector.Get(100, nil, nil)
            if err != nil {
                log.Fatal(err)
            }

            batchId := message.Id
            if batchId == -1 || len(message.Entries) <= 0 {
                time.Sleep(500 * time.Millisecond)
                continue
            }

            printEntry(message.Entries)
        }
    }

    func printEntry(entrys []pbe.Entry) {
        for _, entry := range entrys {
            if entry.GetEntryType() == pbe.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == pbe.EntryType_TRANSACTIONEND {
                continue
            }

            rowChange := new(pbe.RowChange)
            err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
            if err != nil {
                log.Fatalf("Fatal error: %s", err.Error())
            }

            if rowChange != nil {
                eventType := rowChange.GetEventType()
                header := entry.GetHeader()
                fmt.Printf("================> binlog[%s:%d], name[%s,%s], eventType: %s\n", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType())

                for _, rowData := range rowChange.GetRowDatas() {
                    if eventType == pbe.EventType_DELETE {
                        printColumn(rowData.GetBeforeColumns())
                    } else if eventType == pbe.EventType_INSERT {
                        printColumn(rowData.GetAfterColumns())
                    } else {
                        fmt.Println("-------> before")
                        printColumn(rowData.GetBeforeColumns())
                        fmt.Println("-------> after")
                        printColumn(rowData.GetAfterColumns())
                    }
                }
            }
        }
    }

    func printColumn(columns []*pbe.Column) {
        for _, col := range columns {
            fmt.Printf("%s : %s  update=%t\n", col.GetName(), col.GetValue(), col.GetUpdated())
        }
    }

    ```
- 功能验证
  - MySQL 操作
      ```
        CREATE DATABASE test;

        CREATE TABLE `role` (
          `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
          `role_name` varchar(255) DEFAULT NULL,
          `remark` varchar(255) DEFAULT NULL,
          PRIMARY KEY (`id`)
        ) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

        # 然后依次执行下面三条 sql
        insert into role (id, role_name) values (10, 'admin');
        update role set role_name='hh' where id = 10;
        delete from role where id = 10;
      ```
  - 客户端程序输出
      ```
        ================> binlog[mysql-bin.000002:9394], name[test,role], eventType: INSERT
        id : 10  update=true
        role_name : admin  update=true
        remark :   update=true
        ================> binlog[mysql-bin.000002:9662], name[test,role], eventType: UPDATE
        -------> before
        id : 10  update=false
        role_name : admin  update=false
        remark :   update=false
        -------> after
        id : 10  update=false
        role_name : hh  update=true
        remark :   update=false
        ================> binlog[mysql-bin.000002:9940], name[test,role], eventType: UPDATE
        -------> before
        id : 10  update=false
        role_name : hh  update=false
        remark :   update=false
        -------> after
        id : 10  update=false
        role_name : hh  update=false
        remark : foo  update=true
      ```

### Canal 最佳实践 - Docker + MQ 
介绍
- canal 1.1.1版本之后, 默认支持将canal server接收到的binlog数据直接投递到MQ, 目前默认支持的MQ系统有:
  - kafka, RocketMQ, RabbitMQ, pulsarmq

参考
- https://github.com/alibaba/canal/wiki/Canal-Kafka-RocketMQ-QuickStart
- https://www.liwenzhou.com/posts/Go/canal/
- https://developer.aliyun.com/article/892844
- https://www.modb.pro/db/383497

步骤
- 提前安装好 Zookeeper 和 Kafka，并启动
    ```
    # 启动 Zookeeper
    zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties

    # 启动 Kafka
    kafka-server-start /usr/local/etc/kafka/server.properties
    ```
- 使用 docker-compose 运行 docker canal-server 服务
  - 由于 docker-compose canal-server 服务不支持配置 mq 相关的环境变量，所以需要做宿主机与 canal-server 配置文件的映射
  - docker-compose.yml
      ```
        version: '3'
        services:
          canal-server:
            image: canal/canal-server:v1.1.4
            container_name: canal-server
            restart: unless-stopped
            network_mode: host
            ports:
              - 11111:11111
            volumes:
              - ./canal.properties:/home/admin/canal-server/conf/canal.properties
              - ./instance.properties:/home/admin/canal-server/conf/test/instance.properties
              - ./log/:/home/admin/canal-server/logs/
      ```
  - instance.properties 改动以下配置
      ```
        # mysql config 
        canal.instance.master.address=127.0.0.1:13306
        canal.instance.dbUsername=canal
        canal.instance.dbPassword=canal

        # mq config
        canal.mq.topic=my-topic
        canal.mq.partition=0
        canal.mq.flatMessage=true
      ```
  - canal.properties 改动以下配置
      ```
        canal.serverMode = kafka

        # mq config
        canal.mq.servers = 127.0.0.1:9092
      ```
- 启动 canal-server
    ```
    docker-compose up

    ...
    ```
- 消费 kafka topic 中的数据
    - kafka 消费
        ```
        kafka-console-consumer --bootstrap-server localhost:9092 --topic my-topic

        # 执行下面三条 sql，分别打印
        insert into role (id, role_name) values (10, 'admin');
        update role set role_name='hh' where id = 10;
        delete from role where id = 10;

        # 输出
        {"data":[{"id":"10","role_name":"admin","remark":null}],"database":"test","es":1706716357000,"id":5,"isDdl":false,"mysqlType":{"id":"int(10) unsigned","role_name":"varchar(255)","remark":"varchar(255)"},"old":null,"pkNames":["id"],"sql":"","sqlType":{"id":4,"role_name":12,"remark":12},"table":"role","ts":1706716357853,"type":"INSERT"}
        {"data":[{"id":"10","role_name":"hh","remark":null}],"database":"test","es":1706716361000,"id":6,"isDdl":false,"mysqlType":{"id":"int(10) unsigned","role_name":"varchar(255)","remark":"varchar(255)"},"old":[{"role_name":"admin"}],"pkNames":["id"],"sql":"","sqlType":{"id":4,"role_name":12,"remark":12},"table":"role","ts":1706716361206,"type":"UPDATE"}
        {"data":[{"id":"10","role_name":"hh","remark":null}],"database":"test","es":1706716364000,"id":7,"isDdl":false,"mysqlType":{"id":"int(10) unsigned","role_name":"varchar(255)","remark":"varchar(255)"},"old":null,"pkNames":["id"],"sql":"","sqlType":{"id":4,"role_name":12,"remark":12},"table":"role","ts":1706716364384,"type":"DELETE"}

        ```


### Canal 最佳实践 - Canal Admin
背景
- canal-admin设计上是为canal提供整体配置管理、节点运维等面向运维的功能，提供相对友好的WebUI操作界面，方便更多用户快速和安全的操作



### Canal 最佳实践 - Docker + MQ + Redis/ES
背景：
- 近乎实时的缓存 db 数据到 redis
- 将业务与缓存逻辑解耦

需要考虑的问题
- 消息顺序性，如何确保缓存到 redis 的是不是最后修改的数据
  - 解决方法1：单 topic 单 partition 处理，但是并发消费性能差
  - 解决方法2：只有当日志中的 update_time 大于缓存中数据的 update_time 时，才覆盖写入

参考
- https://github.com/alibaba/canal/wiki/Canal-Admin-QuickStart
- https://github.com/alibaba/canal/wiki/Canal-Admin-Docker
- https://zhuanlan.zhihu.com/p/590639613
