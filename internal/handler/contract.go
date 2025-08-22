package handler

type Greeting interface {
	Hello() string
	HelloName(name string) string
}

type Repository interface {
	Set(key, value string) error
	Get(key string) (string, error)
}
