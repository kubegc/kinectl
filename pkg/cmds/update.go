package cmds

import (
	"context"
	"fmt"

	"github.com/litekube/kinectl/pkg/myapp"
	"github.com/urfave/cli/v2"
)

type UpdateArgs struct {
	Key      string
	Revision int64
	Value    string
}

var updateArgs UpdateArgs

func NewUpdateCommand() *cli.Command {
	return &cli.Command{
		Name:      "update",
		Usage:     "update options",
		UsageText: myapp.AppName + " [global options] update [Args]",
		Action:    update,
		Flags:     UpdateFlags,
	}
}

var UpdateFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "key",
		Aliases:     []string{"k"},
		Usage:       "key to locate",
		Required:    true,
		Destination: &updateArgs.Key,
	},
	&cli.StringFlag{
		Name:        "value",
		Aliases:     []string{"v"},
		Usage:       "value to store",
		Required:    true,
		Destination: &updateArgs.Value,
	},
	&cli.Int64Flag{
		Name:        "revision",
		Aliases:     []string{"r"},
		Usage:       "revision of store data",
		Required:    true,
		Destination: &updateArgs.Revision,
	},
}

func update(*cli.Context) error {
	client, err_init := Config.Client()
	if err_init != nil {
		return fmt.Errorf("fail to init client: %s", err_init.Error())
	}
	defer client.Close()

	err := client.Update(context.TODO(), updateArgs.Key, updateArgs.Revision, []byte(updateArgs.Value))
	if err == nil {
		fmt.Println("success")
	}
	return err
}
