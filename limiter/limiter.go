package limiter

type Limiter interface {
	Take() bool
	Empty() bool
}

var _ Limiter = (*TokenBucket)(nil)
