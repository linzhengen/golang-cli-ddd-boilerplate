package hogestore

import "github.com/bxcodec/faker/v3"

// HelloWorld is a interface.
type HelloWorld interface {
	GetName() string
}

type helloWorld struct {
}

// NewHelloWorld is HelloWorld constructor.
func NewHelloWorld() HelloWorld {
	return &helloWorld{}
}

func (h *helloWorld) GetName() string {
	return faker.Name()
}
