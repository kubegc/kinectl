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
		cmds.NewPutCommand(),
		cmds.NewDeleteCommand(),
		cmds.NewUpdateCommand(),
		cmds.NewGetCommand(),
		cmds.NewListCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("error: ", err.Error())
	}
}
