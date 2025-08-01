package usecases

type GreetingUsecase struct{}

func NewGreetingUsecase() *GreetingUsecase {
	return &GreetingUsecase{}
}

func (gu *GreetingUsecase) Hello() string {
	return "Hello"
}

func (gu *GreetingUsecase) HelloName(name string) string {
	return "Hello" + name + "!"
}
