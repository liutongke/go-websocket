package RdLine

import (
	"context"
	"errors"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

//https://learnku.com/articles/46788
//基于redis的分布式锁
type DisLockRedis struct {
	key       string             //锁名称
	ttl       int64              //锁超时时间
	isLocked  bool               //上锁成功标识
	cancelFun context.CancelFunc //用于取消自动续租携程

	redis *redis.Pool
	debug bool
}

func NewDisLockRedis(key string) *DisLockRedis {
	return &DisLockRedis{
		key:   key,
		ttl:   30,
		redis: GetRedisPool(),
	}
}

//上锁
func (this *DisLockRedis) TryLock() (err error) {
	if err = this.grant(); err != nil {
		return
	}

	ctx, cancelFun := context.WithCancel(context.TODO())

	this.cancelFun = cancelFun
	//自动续期
	this.renew(ctx)

	this.isLocked = true

	return nil
}

//释放锁
func (this *DisLockRedis) Unlock() (err error) {
	var res int
	if this.isLocked {
		if res, err = redis.Int(this.redisConn().Do("DEL", this.key)); err != nil {

			if this.debug {
				log.Println(err.Error())
			}
			return errors.New("释放锁失败")
		}

		if res == 1 {
			//释放成功，取消自动续租
			this.cancelFun()
			return
		}
	}

	return errors.New("释放锁失败")

}

//自动续期
func (this *DisLockRedis) renew(ctx context.Context) {

	go func() {

		for {
			select {
			case <-ctx.Done():
				return
			default:
				res, err := redis.Int(this.redisConn().Do("EXPIRE", this.key, this.ttl))
				if this.debug {
					if err != nil {
						log.Println("锁自动续期失败：", err)
					}

					if res != 1 {
						log.Println("锁自动续期失败")
					}
				}
			}

			time.Sleep(time.Duration(this.ttl/3) * time.Second)
		}
	}()

}

//创建租约
func (this *DisLockRedis) grant() (err error) {

	if res, err := redis.String(this.redisConn().Do("SET", this.key, "xxx", "NX", "EX", this.ttl)); err != nil {
		if this.debug {
			log.Println(err)
		}

	} else {
		if res == "OK" {
			return nil
		}
	}

	return errors.New("上锁失败")
}

func (this *DisLockRedis) redisConn() redis.Conn {
	return this.redis.Get()
}

func (this *DisLockRedis) Debug() *DisLockRedis {
	this.debug = true
	return this
}
