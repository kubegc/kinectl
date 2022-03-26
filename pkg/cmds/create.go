package cmds

import (
	"context"
	"fmt"

	"github.com/litekube/kinectl/pkg/myapp"
	"github.com/urfave/cli/v2"
)

type CreateArgs struct {
	Key   string
	Value string
}

var createArgs CreateArgs

func NewCreateCommand() *cli.Command {
	return &cli.Command{
		Name:      "create",
		Usage:     "create options",
		UsageText: myapp.AppName + " [global options] create [Args]",
		Action:    create,
		Flags:     CreateFlags,
	}
}

var CreateFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "key",
		Aliases:     []string{"k"},
		Usage:       "key to locate",
		Required:    true,
		Destination: &createArgs.Key,
	},
	&cli.StringFlag{
		Name:        "value",
		Aliases:     []string{"v"},
		Usage:       "value to store",
		Required:    true,
		Destination: &createArgs.Value,
	},
}

func create(*cli.Context) error {
	client, err_init := Config.Client()
	if err_init != nil {
		fmt.Println("fail to init client.")
		return err_init
	}
	defer client.Close()

	err := client.Create(context.Background(), createArgs.Key, []byte(createArgs.Value))
	if err == nil {
		fmt.Println("success")
	} else {
		fmt.Println(err.Error())
	}
	return err
}
