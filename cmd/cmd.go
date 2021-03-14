package main

import (
	"log"
	"os"

	"github.com/linzhengen/golang-cli-ddd-boilerplate/injector"

	"github.com/linzhengen/golang-cli-ddd-boilerplate/interfaces/command"
	"github.com/urfave/cli/v2"
)

// VERSION app version.
var VERSION = "dev"

func main() {
	app := cli.NewApp()
	app.Name = "hello world cli"
	app.Version = VERSION
	app.Usage = "golang cli DDD boilerplate"
	app.Commands = command.NewCommand(injector.BuildInjector())
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
