package ds

import (
	"github.com/google/wire"

	"github.com/tal-tech/go-zero/core/stores/redis"
)

var Set = wire.NewSet(
	DefaultRedis,
)

func DefaultRedis() *redis.Redis {
	return redis.NewRedis("127.0.0.1:6379", redis.NodeType, "123123")
}
