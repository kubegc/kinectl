package cmds

import (
	"context"
	"fmt"

	"github.com/litekube/kinectl/pkg/myapp"
	"github.com/urfave/cli/v2"
)

type DeleteArgs struct {
	Key      string
	Revision int64
}

var deleteArgs DeleteArgs

func NewDeleteCommand() *cli.Command {
	return &cli.Command{
		Name:      "delete",
		Usage:     "delete options",
		UsageText: myapp.AppName + " [global options] delete [Args]",
		Action:    delete,
		Flags:     DeleteFlags,
	}
}

var DeleteFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "key",
		Aliases:     []string{"k"},
		Usage:       "key to locate",
		Required:    true,
		Destination: &deleteArgs.Key,
	},
	&cli.Int64Flag{
		Name:        "revision",
		Aliases:     []string{"r"},
		Usage:       "revision of store data",
		Value:       0,
		Destination: &deleteArgs.Revision,
	},
}

func delete(*cli.Context) error {
	client, err_init := Config.Client()
	if err_init != nil {
		return fmt.Errorf("fail to init client: %s", err_init.Error())
	}
	defer client.Close()

	err := client.Delete(context.TODO(), deleteArgs.Key, deleteArgs.Revision)
	if err == nil {
		fmt.Println("success")
	}
	return err
}
