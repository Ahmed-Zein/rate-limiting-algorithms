package limiter

type Limiter interface {
	Take() bool
}

var _ Limiter = (*TokenBucket)(nil)
