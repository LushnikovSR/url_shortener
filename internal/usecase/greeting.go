package usecase

type Greeting struct{}

func NewGreeting() *Greeting {
	return &Greeting{}
}

func (gu *Greeting) Hello() string {
	return "Hello"
}

func (gu *Greeting) HelloName(name string) string {
	return "Hello" + " " + name + "!"
}
