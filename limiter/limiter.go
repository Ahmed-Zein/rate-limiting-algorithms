package limiter

type Limiter interface {
	IsAllowed() bool
}

var _ Limiter = (*TokenBucket)(nil)
var _ Limiter = (*LeakyBucket)(nil)

var _ Limiter = (*FixedWindowCounter)(nil)
var _ Limiter = (*SlidingWindowLog)(nil)
var _ Limiter = (*SlidingWindowCounter)(nil)
