package main

import (
	"time"

	"figoxu.me/redis/pkg/ut"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"

	"github.com/quexer/utee"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

type Barrier interface {
	Acquire() (bool, error)
	SlotLeft() (int, error)
}

func NewBarrier(Redis *redis.Redis, Bucket string, Duration, Slot int) Barrier {
	barrier := &TimeCountBarrier{
		Redis:    Redis,
		Bucket:   Bucket,
		Duration: Duration,
		Slot:     Slot,
	}
	go func() {
		defer ut.Recovery()
		time.Sleep(time.Minute)
		barrier.Clean()
	}()
	return barrier
}

type TimeCountBarrier struct {
	Redis    *redis.Redis
	Bucket   string // 桶名:沙箱隔离不同业务
	Duration int    // 计算周期,单位秒
	Slot     int    // 阈值
}

func (p *TimeCountBarrier) key() string {
	return "tcb_" + p.Bucket
}

var LockIsBusy = errors.New("当前有其他用户正在进行操作，请稍后重试")

func (p *TimeCountBarrier) tryLock(times int, gapDuration time.Duration) (*redis.RedisLock, error) {
	redisLock := redis.NewRedisLock(p.Redis, "lock_"+p.key())
	redisLock.SetExpire(10)
	for i := 0; i < times; i++ {
		ok, err := redisLock.Acquire()
		if ok && err == nil {
			return redisLock, nil
		}
		time.Sleep(gapDuration)
	}
	return nil, LockIsBusy
}

func (p *TimeCountBarrier) Acquire() (bool, error) {
	redisLock, err := p.tryLock(1000, time.Millisecond)
	if err != nil {
		return false, err
	}
	defer func() {
		recover()
		redisLock.Release()
	}()
	left, err := p.SlotLeft()
	if err != nil {
		return false, err
	} else if left < 1 {
		return false, nil
	}
	_, err = p.Redis.Zadd(p.key(), utee.Tick(), uuid.New())
	if err != nil {
		return false, errors.WithStack(err)
	}
	return true, nil
}

func (p *TimeCountBarrier) SlotLeft() (int, error) {
	endTick := utee.Tick()
	slotUsed, err := p.Redis.Zcount(p.key(), p.startTick(), endTick)
	if err != nil {
		return 0, err
	}
	return p.Slot - slotUsed, nil
}

func (p *TimeCountBarrier) startTick() int64 {
	return utee.Tick(time.Now().Add(time.Second * time.Duration(-1*p.Duration)))
}

func (p *TimeCountBarrier) Clean() error {
	_, err := p.Redis.Zremrangebyscore(p.key(), 0, p.startTick()-1)
	return errors.WithStack(err)
}
