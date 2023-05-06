## Kafka 使用（Mac）

### **安装和启动**

启动

brew services start kafka

或者

zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties &

kafka-server-start /usr/local/etc/kafka/server.properties

### **常用命令**

**1.topic**

查看所有topic

kafka-topics --list --bootstrap-server localhost:9092

创建topic

kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic topic-name

(--zookeeper将被弃用) kafka-topics --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic topic-name

查看topic

(--zookeeper将被弃用)kafka-topics --list --zookeeper localhost:2181

kafka-topics --describe --bootstrap-server localhost:9092 --topic topic-name

删除指定topic

kafka-topics --delete --bootstrap-server localhost:9092 --topic topic-name

**2.生产者消费者**

kafka-console-producer --bootstrap-server localhost:9092 --topic topic-name

kafka-console-consumer --bootstrap-server localhost:9092 --topic topic-name --from-beginning

**3.修改partition**

kafka-topics --alter --bootstrap-server localhost:9092 --topic test-topic --partitions partition_num

**4.消息相关**

查看 topic 消费 offset

/home/kafka/bin/kafka-get-offsets.sh get-offset-shell --bootstrap-server localhost:9092 --topic topic-name

重置消息

[https://cloud.tencent.com/developer/article/1436988](https://cloud.tencent.com/developer/article/1436988)

**命令介绍**

- zookeeper-server-start
- kafka-server-start
- kafka-topics
- kafka-console-producer
- kafka-console-consumer
- kafka-configs

kafka-topics

```
--bootstrap-server <String: server to connect to>
--config <String: name=value>
--create              Create a new topic
--delete              Delete a topic
--describe            List details for the given topics
--list                List all available topics
--partitions <Integer: # of partitions>
--replication-factor <Integer: replication factor>
--topic <String: topic>
```

### 其他常用命令
开启 zookeeper 服务（默认监听2181端口）

/usr/local/Cellar/kafka/2.5.0/bin/zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties

开启 kafka server 服务（默认监听9092端口）

/usr/local/Cellar/kafka/2.5.0/bin/kafka-server-start /usr/local/etc/kafka/server.properties


列出所有可用 topic

/usr/local/Cellar/kafka/2.5.0/bin/kafka-topics --list --bootstrap-server localhost:9092

创建 topic

/usr/local/Cellar/kafka/2.5.0/bin/kafka-topics --create --topic test --bootstrap-server localhost:9092

查看 topic 详细信息

/usr/local/Cellar/kafka/2.5.0/bin/kafka-topics --describe --topic test --bootstrap-server localhost:9092

删除 topic

/usr/local/Cellar/kafka/2.5.0/bin/kafka-topics --delete --topic test --bootstrap-server localhost:9092


消息生产者

/usr/local/Cellar/kafka/2.5.0/bin/kafka-console-producer --topic test --bootstrap-server localhost:9092

消息消费者

/usr/local/Cellar/kafka/2.5.0/bin/kafka-console-consumer --topic test --bootstrap-server localhost:9092

/usr/local/Cellar/kafka/2.5.0/bin/kafka-console-consumer --topic test --from-beginning --bootstrap-server localhost:9092

重置消费位点
```
kafka-consumer-groups \
    --bootstrap-server localhost:9092 \
    --group my-consumer-group \
    --topic myTopicName \
    --reset-offsets \
    --to-offfset 100 \
    --execute
```
