// +build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/linzhengen/golang-cli-ddd-boilerplate/application"
	"github.com/linzhengen/golang-cli-ddd-boilerplate/domain/repository"
	"github.com/linzhengen/golang-cli-ddd-boilerplate/interfaces/action"
)

// BuildInjector build injector use wire.
func BuildInjector() action.Action {
	wire.Build(
		// action
		action.NewAction,
		action.NewHelloWorldAction,
		// application
		application.NewHelloWorldApp,
		// repository
		repository.NewHelloWorldRepository,
	)
	return nil
}
