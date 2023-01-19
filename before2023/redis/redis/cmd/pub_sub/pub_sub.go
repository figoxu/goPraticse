package main

import (
	"time"

	"figoxu.me/redis/pkg/ut"
	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

type Pub interface {
	Pub(channel, message string) error
}

type Sub interface {
	Sub(channel string, callBack SubCb)
}
type SubCb func(string)

type PubSub struct {
	Redis *redis.Redis
}

func NewPubSub(v *redis.Redis) *PubSub {
	return &PubSub{
		Redis: v,
	}
}

func (p *PubSub) Pub(channel, message string) error {
	_, err := p.Redis.Rpush(channel, message)
	return errors.WithStack(err)
}

func (p *PubSub) Sub(channel string, callBack SubCb) {
	exc := func() error {
		defer ut.Recovery()
		v, err := p.Redis.Lpop(channel)
		if err != nil {
			return err
		}
		if v == "" {
			time.Sleep(time.Millisecond * time.Duration(10))
		}
		callBack(v)
		return nil
	}
	loopExc := func() {
		for {
			err := exc()
			if err == redis.Nil {
				time.Sleep(time.Second)
			} else if err != nil {
				ut.Log().WithErr(err).Entry().Println(`panic for subscriber `, channel)
				time.Sleep(time.Second)
			}
		}
	}
	go loopExc()
}
