## Redis commands

### zset 有序集合
zadd key score member
zincrby key increment member

zrange key start stop [WITHSCORES]
zrangebyscore key min max [WITHSCORES]
zrevrange key start stop [WITHSCORES]
zrevrangebyscore key max min [WITHSCORES]
zrank key member
zrevrank key member

根据score由小到大取top10数据
zrange key -inf inf byscore limit 0 10 [WITHSCORES]

zcard key
zcount key min max

zscore key member
zmscore key member [member ...]


zrem key member [member ...]
zremrangebyrank key start stop
zremrangebyscore key min max

zmpop numkeys key [key ...] MIN|MAX
zpopmax key
zpopmin key

zrangestore dst src min max
