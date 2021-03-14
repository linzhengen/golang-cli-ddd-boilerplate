package action

import "github.com/linzhengen/golang-cli-ddd-boilerplate/application"

// HelloWorldAction is a interface.
type HelloWorldAction interface {
	HelloWorld() error
}

type helloWorldAction struct {
	a application.HelloWorldApp
}

// NewHelloWorldAction is HelloWorldAction constructor.
func NewHelloWorldAction(a application.HelloWorldApp) HelloWorldAction {
	return &helloWorldAction{a: a}
}

func (h *helloWorldAction) HelloWorld() error {
	return h.a.HelloWorld()
}
