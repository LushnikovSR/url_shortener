package greeting

type greeting interface {
	Hello() string
	HelloName(name string) string
}
