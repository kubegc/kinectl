package main

import (
	"fmt"
	"os"

	"github.com/litekube/kinectl/pkg/cmds"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cmds.NewApp()
	app.Commands = []*cli.Command{
		cmds.NewCreateCommand(),
		cmds.NewDeleteCommand(),
		cmds.NewPutCommand(),
		cmds.NewUpdateCommand(),
		cmds.NewGetCommand(),
		cmds.NewListCommand(),
		cmds.NewVersionCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error options: %s\n", err.Error())
		os.Exit(-1)
	}
}
