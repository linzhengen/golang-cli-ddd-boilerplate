package application

import (
	"log"

	"github.com/linzhengen/golang-cli-ddd-boilerplate/domain/repository"
)

// HelloWorldApp is a interface.
type HelloWorldApp interface {
	HelloWorld() error
}

type hellWorldApp struct {
	r repository.HelloWorldRepository
}

// NewHelloWorldApp is HelloWorldApp constructor.
func NewHelloWorldApp(r repository.HelloWorldRepository) HelloWorldApp {
	return &hellWorldApp{r}
}

func (h *hellWorldApp) HelloWorld() error {
	name, err := h.r.GetName()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("hello world: %s", name.Name)
	return nil
}
