## 常见操作示例（Linux）

### Kafka 配置
支持外网访问的配置(修改配置后需重启 kafka)

配置文件:/home/kafka/config/server.properties

配置内容:
- listeners=PLAINTEXT://0.0.0.0:9092
- advertised.listeners=PLAINTEXT://10.18.32.23:9092

Ref:
- http://openfile.shixinke.com/images/posts/2019/02/9efc02f4b076c8e8c521e2a0adbce19f.png

### Kafka 常见操作示例
**topic 相关**
查看 topics 列表
> /home/kafka/bin/kafka-topics.sh kafka-topics --list --bootstrap-server localhost:9092

查看 topic 详情
> /home/kafka/bin/kafka-topics.sh kafka-topics --describe --bootstrap-server localhost:9092 --topic topic-name

创建 topic
> /home/kafka/bin/kafka-topics.sh kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic topic-name

删除指定topic
> /home/kafka/bin/kafka-topics.sh kafka-topics --delete --bootstrap-server localhost:9092 --topic topic-name

**partition 相关**
修改分区数量(partition_num 为具体数值)
> /home/kafka/bin/kafka-topics.sh kafka-topics --alter --bootstrap-server localhost:9092 --topic test-topic --partitions partition_num

**生产者**
生产者 console
> /home/kafka/bin/kafka-console-producer.sh kafka-console-producer --bootstrap-server localhost:9092 --topic topic-name

**消费者**
消费者 console
> /home/kafka/bin/kafka-console-consumer.sh kafka-console-consumer --bootstrap-server localhost:9092 --topic topic-name

> /home/kafka/bin/kafka-console-consumer.sh kafka-console-consumer --bootstrap-server localhost:9092 --topic topic-name --from-beginning

消息消费时指定消费者组
> /home/kafka/bin/kafka-console-consumer.sh kafka-console-consumer --bootstrap-server localhost:9092 --group group-name --topic topic-name

**消费者组相关**
查看消费者组列表
> /home/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --list

查看消费者组订阅的所有 topic 的消费情况
> /home/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe --group {group_name}

**消息 offset 相关**
查看 topic 的 offset
> /home/kafka/bin/kafka-get-offsets.sh get-offset-shell --bootstrap-server localhost:9092 --topic topic-name

重置消费 offset 到当前最新位置(需要消费 group is inactive 状态，but the current state is Stable 也可以)
> /home/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group {group-name} --topic {topic-name} --reset-offsets --to-current --execute

重置消费 offset 到指定位置(offset_num 为具体数值)
> /home/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group {group-name} --topic {topic-name} --reset-offsets --to-offset {offset_num} --execute

**参考**
- Kafka 安装参考：https://linuxhostsupport.com/blog/how-to-install-apache-kafka-on-centos-7/
- Kafka 安装参考：https://segmentfault.com/a/1190000012730949
- Kafka 重置消费 offset 参考：https://cloud.tencent.com/developer/article/1436988
- Kafka Consumer 重置 Offset：https://juejin.cn/post/7069388633999933471

