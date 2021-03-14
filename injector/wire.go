// +build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/linzhengen/ddd-gin-admin/application"
	"github.com/linzhengen/ddd-gin-admin/domain/repository"
	"github.com/linzhengen/ddd-gin-admin/interfaces/action"
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
