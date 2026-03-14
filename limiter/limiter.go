package limiter

type Limiter interface {
	IsAllowed() bool
}
