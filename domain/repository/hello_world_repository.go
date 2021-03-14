package repository

import (
	"github.com/linzhengen/golang-cli-ddd-boilerplate/domain/entity"
	"github.com/linzhengen/golang-cli-ddd-boilerplate/infrastructure/persistence/hogestore"
)

// HelloWorldRepository is a interface.
type HelloWorldRepository interface {
	GetName() (*entity.HelloWorld, error)
}

type helloWorldRepository struct {
}

// NewHelloWorldRepository is HelloWorldRepository constructor.
func NewHelloWorldRepository() HelloWorldRepository {
	return &helloWorldRepository{}
}

func (h *helloWorldRepository) GetName() (*entity.HelloWorld, error) {
	return &entity.HelloWorld{Name: hogestore.NewHelloWorld().GetName()}, nil
}
