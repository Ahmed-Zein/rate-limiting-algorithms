package config

import "time"

type BucketBasedConfig struct {
	Capacity int
	Rate     float64
}

type WindowBasedConfig struct {
	WindowSize time.Duration
	Limit      int
}
