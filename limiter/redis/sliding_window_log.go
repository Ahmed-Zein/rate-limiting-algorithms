package redis

import (
	"context"
	"strconv"
	"time"

	"github.com/ahmed-zein/go_rate_limiting/config"
	"github.com/redis/go-redis/v9"
)

type SlidingWindowLog struct {
	limit      int
	windowSize time.Duration
	domain     string
	rdb        *redis.Client
}

func NewSlidingWindowCounter(domain string, cfg *config.WindowBasedConfig, rdb *redis.Client) (*SlidingWindowLog, error) {
	return &SlidingWindowLog{
		limit:      cfg.Limit,
		windowSize: cfg.WindowSize,
		rdb:        rdb,
		domain:     domain,
	}, nil
}

func (sc *SlidingWindowLog) Allow(id string) (bool, error) {
	return sc.AllowN(id, 1)
}

func (sc *SlidingWindowLog) AllowN(id string, n int) (bool, error) {
	ctx := context.Background()
	key := generateKey(sc.domain, id)

	now := time.Now()
	windowStart := now.Add(-sc.windowSize)

	pipe := sc.rdb.Pipeline()
	pipe.ZRemRangeByScore(ctx, key, "0", strconv.FormatInt(windowStart.UnixNano(), 10))
	count := pipe.ZCard(ctx, key)
	pipe.Exec(ctx)
	if count.Val()+int64(n) > int64(sc.limit) {
		return false, nil
	}
	pipe = sc.rdb.Pipeline()

	for range n {
		nowUnixNano := time.Now().UnixNano()
		pipe.ZAdd(ctx, key, redis.Z{
			Score:  float64(nowUnixNano),
			Member: nowUnixNano,
		})
	}

	pipe.Expire(ctx, key, sc.windowSize)
	pipe.Exec(ctx)

	return true, nil
}
