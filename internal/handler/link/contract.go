package link

type repository interface {
	Set(key, value string) error
	Get(key string) (string, error)
}
