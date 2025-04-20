# Distributed Locks

## 最佳实践
- 谈谈分布式锁 - 四火的唠叨 https://www.raychase.net/7671
- 再有人问你分布式锁，这篇文章扔给他 https://juejin.cn/post/6844903688088059912
- 聊聊分布式锁 https://mp.weixin.qq.com/s/-N4x6EkxwAYDGdJhwvmZLw
- 聊一聊分布式锁的设计模型 https://mp.weixin.qq.com/s/uA26VVmYMTfs-dWcLOY04w
- 分布式锁实现原理与最佳实践 https://mp.weixin.qq.com/s/JzCHpIOiFVmBoAko58ZuGw
- Redis、ZooKeeper、Etcd，谁有最好用的分布式锁？https://mp.weixin.qq.com/s/yZC6VJGxt1ANZkn0SljZBg
- 一文讲透Redis分布式锁安全问题 https://mp.weixin.qq.com/s/O8o31rRBVL1DwK-JfmurRw
- 深度剖析：Redis分布式锁到底安全吗？看完这篇文章彻底懂了！https://zhuanlan.zhihu.com/p/378797329
- 用 redis 做分布式锁这种骚操作是怎么流行起来的？https://www.v2ex.com/t/759583
- 基于Redis的分布式锁到底安全吗？
  - 上 http://zhangtielei.com/posts/blog-redlock-reasoning.html
  - 下 http://zhangtielei.com/posts/blog-redlock-reasoning-part2.html
- 面试官：你真的了解Redis分布式锁吗？https://www.cnblogs.com/yeya/p/14274948.html
- 如何1分钟实现一个分布式锁？https://mp.weixin.qq.com/s/BJ-8czCQLVgeSrIVr4JuIQ


## Third Lib
- Distributed Locks with Redis
  - https://redis.io/docs/latest/develop/use/patterns/distributed-locks/
  - https://github.com/go-redsync/redsync


## Demo
### 使用 redis setnx 实现分布式锁
```
package main

/*
基于redis的分布式锁，利用redis的setnx的特性，结合redis的expire超时删除key的机制。
1.redis分布式锁的核心思想是redis的setnx，利用setnx指令尝试向redis申请指定key，这个key就是我们的锁，当返回值为1时我们成功申请到分布式锁。
2.利用expire指令对申请到的锁设置超时时间，防止占有的进程一直不释放锁导致死锁。
3.需要提供两个加锁机制，一是尝试加锁后，加锁失败就直接返回；另一种是尝试加锁后，如果加锁失败就循环等待和重试，直到超时时间到了，再结束返回。
4.redis客户端作为分布式锁的高频操作者，注意不要反复创建连接，而应一直保持redis client实例。

TODO: Redis 分布式锁过期了，但业务还没有执行完，怎么办？
- Watchdog 进程定时给锁续命，给锁加上新的过期时间
- 参考：https://blog.csdn.net/qq_36594703/article/details/123500652
*/

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type Client struct {
	redis.UniversalClient
}

type Lock struct {
	timeout int
	tag     string
	client  *Client
}

func NewRedisClient() (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})

	err := client.Ping().Err()
	if err != nil {
		fmt.Println("client.Ping, ", err)
		return nil, errors.New("redis simple connect fail")
	}
	return &Client{client}, nil
}

var safeLock = redis.NewScript(`
	local key = KEYS[1]
	local r = redis.call("SETNX", key, 1)
	if (r == 0) then
		return 0
	end

	redis.call("EXPIRE", key, ARGV[1])
	return 1
`)

// 分布式锁key
func (lock *Lock) key() string {
	return "lock_" + lock.tag
}

func (lock *Lock) tryLock() (bool, error) {
	ok, err := safeLock.Run(lock.client, []string{lock.key()}, lock.timeout).Int64()
	if err != nil {
		return false, err
	}
	if ok == 0 {
		return false, nil
	}

	return true, nil
}

// 设置带超时的锁
func (lock *Lock) LockWithTimeout(client *Client, tag string, timeout int) (ok bool, err error) {
	if timeout < 0 {
		timeout = 0
	}
	lock.tag = tag
	lock.timeout = timeout
	lock.client = client
	ok, err = lock.tryLock()
	if !ok || err != nil {
		return
	}
	return
}

// 循环等待的锁，获取分布式锁失败，那就一直循环等待，一直循环等待maxWaitTime==-1
func (lock *Lock) LockWithWait(client *Client, tag string, timeout, maxWaitTime int) (ok bool, err error) {
	now := time.Now()
	for {
		ok, err = lock.LockWithTimeout(client, tag, timeout)
		if err != nil {
			fmt.Println("LockWithWait error, ", tag, timeout, maxWaitTime)
			return
		}
		if ok {
			return
		}
		fmt.Println("LockWithWait fail, ", tag, timeout, maxWaitTime)
		if maxWaitTime != -1 && int(time.Since(now).Seconds()) > maxWaitTime {
			return false, fmt.Errorf("maxWaitTime reach, %s %d %d", tag, timeout, maxWaitTime)
		}
		time.Sleep(time.Millisecond * 300)
	}
}

// lock
func (lock *Lock) Lock(client *Client, tag string) (ok bool, err error) {
	// lock without timeout
	// return lock.LockWithTimeout(client, tag, 1)

	// lock with wait forever
	// return lock.LockWithWait(client, tag, 1, -1)

	// lock with wait and waittime
	return lock.LockWithWait(client, tag, 1, 5)
}

// unlock
func (lock *Lock) Unlock() (err error) {
	err = lock.client.Del(lock.key()).Err()
	return
}

func doSomething(index int) {
	fmt.Printf("%d goroutine do!!!\n", index)
	time.Sleep(time.Second)
}

func main() {
	redisClient, err := NewRedisClient()
	if err != nil {
		fmt.Println("NewRedisClient fail, ", err)
		return
	}
	for i := 0; i < 10; i++ {
		go func(index int) {
			lock := Lock{}
			ok, err := lock.Lock(redisClient, "example_uid")
			if err != nil || !ok {
				fmt.Println("apply lock fail,", err)
				return
			}
			defer func() {
				lock.Unlock()
			}()
			doSomething(index)
		}(i)
	}

	select {}
}

```
