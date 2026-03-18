package limiter

type Limiter interface {
	Allow(id string) (bool, error)
	AllowN(id string, n int) (bool, error)
}
