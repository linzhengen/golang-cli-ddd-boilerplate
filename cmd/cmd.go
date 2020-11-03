package main

import (
	"log"
	"os"

	"github.com/linzhengen/ddd-gin-admin/application"
	"github.com/linzhengen/ddd-gin-admin/domain/repository"

	"github.com/linzhengen/ddd-gin-admin/interfaces/action"

	"github.com/linzhengen/ddd-gin-admin/interfaces/command"
	"github.com/urfave/cli/v2"
)

// VERSION app version.
var VERSION = "dev"

func main() {
	app := cli.NewApp()
	app.Name = "hello world cli"
	app.Version = VERSION
	app.Usage = "golang cli DDD boilerplate"
	app.Commands = command.NewCommand(action.NewAction(action.NewHelloWorldAction(application.NewHelloWorldApp(repository.NewHelloWorldRepository()))))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
