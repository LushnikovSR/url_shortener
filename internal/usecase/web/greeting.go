package web

type Greeting struct{}

func New() *Greeting {
	return &Greeting{}
}

func (gu *Greeting) Hello() string {
	return "Hello"
}

func (gu *Greeting) HelloName(name string) string {
	return "Hello" + " " + name + "!"
}
