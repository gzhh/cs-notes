# ElasticSearch

![核心知识](images/elastic_search_notes.webp)

## 基本概念

### 索引

- 正向索引
    
    ID → 内容
    
- 倒排索引（反向索引）
    
    词汇 → 在文档中的位置
    
    解释：被用来存储在全文搜索下某个单词在一个文档或者一组文档中的存储位置的映射。它是文档检索系统中最常用的数据结构。

## 基本使用

### 安装配置

启动 Elasticsearch

`elasticsearch`

查看

[http://localhost:9200/](http://localhost:9200/)

查看插件

`elasticsearch-plugin list`

安装插件

`elasticsearch-plugin install analysis-icu`

### 集群

在开发机上运行多个 Elasticsearch 实例

`elasticsearch -E node.name=node0 -E cluster.name=geektime -E path.data=node0_data -d`

`elasticsearch -E node.name=node1 -E cluster.name=geektime -E path.data=node1_data -d`

`elasticsearch -E node.name=node2 -E cluster.name=geektime -E path.data=node2_data -d`

`elasticsearch -E node.name=node3 -E cluster.name=geektime -E path.data=node3_data -d`

查看集群节点

[http://localhost:9200/_cat/nodes](http://localhost:9200/_cat/nodes)

### Kibana

启动 Kibana

`kibana`

查看

[http://localhost:5601/](http://localhost:5601/)

es 管理工具 cerebro

[http://localhost:9000/](http://localhost:5601/)
