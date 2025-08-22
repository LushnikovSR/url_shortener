package usecase

type Repository interface {
	Set(key, value string)
	Get(key string) (string, bool)
}
