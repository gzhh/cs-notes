# 分库分表

## 分区表
```
ALTER  TABLE  `p_channel_data`  add  PARTITION(
    PARTITION p20210203	  VALUES  LESS THAN (738189) ENGINE = InnoDB,
    PARTITION p20210210	  VALUES  LESS THAN (738196) ENGINE = InnoDB,
    PARTITION p20210217	  VALUES  LESS THAN (738203) ENGINE = InnoDB
);
```

写脚本定时自动创建分区

参考 https://my.oschina.net/u/4327542/blog/3373310


## 分库分表

![multi-database-table.jpeg](images/multi-database-table.jpeg)
