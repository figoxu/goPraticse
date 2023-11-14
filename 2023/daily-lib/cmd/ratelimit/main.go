package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/ory/dockertest/v3"
	"log"
	"net/url"
	"time"
)

//https://github.com/taptap/ratelimiter-spring-boot-starter/blob/ac08f28cfd3febbb95dda6748b9eff605c3624ce/src/main/java/com/taptap/ratelimiter/core/TokenBucketRateLimiter.java#L30
//https://github.com/taptap/ratelimiter-spring-boot-starter/blob/ac08f28cfd3febbb95dda6748b9eff605c3624ce/src/main/resources/META-INF/tokenBucket-rateLimit.lua

func tokenBucketRateLimit(client *redis.Client, tokensKey, timestampKey string, rate, capacity, requested int64) (bool, int64, error) {
	luaScript := `
local tokens_key = KEYS[1]
local timestamp_key = KEYS[2]
local rate = tonumber(ARGV[1])
local capacity = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local requested = tonumber(ARGV[4])

local fill_time = capacity/rate
local ttl = math.floor(fill_time*2)

local last_tokens = tonumber(redis.call("get", tokens_key))
if last_tokens == nil then
  last_tokens = capacity
end

local last_refreshed = tonumber(redis.call("get", timestamp_key))
if last_refreshed == nil then
  last_refreshed = 0
end

local delta = math.max(0, now-last_refreshed)
local filled_tokens = math.min(capacity, last_tokens+(delta*rate))
local allowed = filled_tokens >= requested
local new_tokens = filled_tokens
local allowed_num = 0
if allowed then
  new_tokens = filled_tokens - requested
  allowed_num = 1
end

redis.call("setex", tokens_key, ttl, new_tokens)
redis.call("setex", timestamp_key, ttl, now)

return { allowed_num, new_tokens }
`
	now := time.Now().Unix()
	cmd := client.Eval(context.Background(), luaScript, []string{tokensKey, timestampKey}, rate, capacity, now, requested)
	result, err := cmd.Result()
	if err != nil {
		return false, 0, err
	}

	values := result.([]interface{})
	allowed := values[0].(int64) == 1
	newTokens := values[1].(int64)

	return allowed, newTokens, nil
}

func main() {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.Run("redis", "latest", nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	redisURL := url.URL{
		Scheme: "redis",
		Host:   fmt.Sprintf("localhost:%s", resource.GetPort("6379/tcp")),
	}

	client := redis.NewClient(&redis.Options{
		Addr: redisURL.Host,
	})

	ctx := context.Background()

	if err := pool.Retry(func() error {
		_, err := client.Ping(ctx).Result()
		return err
	}); err != nil {
		log.Fatalf("Could not connect to Redis: %s", err)
	}

	allowed, newTokens, err := tokenBucketRateLimit(client, "tokens_key", "timestamp_key", 10, 100, 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Allowed:", allowed)
	fmt.Println("New Tokens:", newTokens)

	// When you're done, kill and remove the container
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	fmt.Println("Done")
}
