package command

import (
	"github.com/linzhengen/ddd-gin-admin/interfaces/action"
	"github.com/urfave/cli/v2"
)

// NewCommand is Command constructor.
func NewCommand(a action.Action) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "hello-world",
			Usage: "Execute hello world",
			Action: func(context *cli.Context) error {
				return a.HelloWorld()
			},
		},
	}
}
