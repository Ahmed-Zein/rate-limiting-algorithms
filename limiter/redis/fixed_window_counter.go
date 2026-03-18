package redis

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/ahmed-zein/go_rate_limiting/config"
	"github.com/redis/go-redis/v9"
)

type FixedWindowCounter struct {
	domain     string
	limit      int
	windowSize time.Duration
	rdb        *redis.Client
	mu         sync.Mutex
}

func NewFixedWindowCounter(domain string, cfg *config.WindowBasedConfig, rdb *redis.Client) (*FixedWindowCounter, error) {
	if cfg.Limit <= 0 {
		return nil, errors.New("limit must be greater than 0")
	}

	return &FixedWindowCounter{
		limit:      cfg.Limit,
		windowSize: cfg.WindowSize,
		domain:     domain,
		rdb:        rdb,
	}, nil
}

func (fw *FixedWindowCounter) Allow(id string) (bool, error) {
	return fw.AllowN(id, 1)
}

func (fw *FixedWindowCounter) AllowN(id string, n int) (bool, error) {
	ctx := context.Background()
	key := generateKey(fw.domain, id)

	pipe := fw.rdb.Pipeline()
	incr := pipe.IncrBy(ctx, key, int64(n))
	pipe.Expire(ctx, key, fw.windowSize)

	if _, err := pipe.Exec(ctx); err != nil {
		return false, err
	}

	return incr.Val() <= int64(fw.limit), nil
}
