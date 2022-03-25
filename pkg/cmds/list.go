package cmds

import (
	"context"
	"fmt"

	"github.com/litekube/kinectl/pkg/myapp"
	"github.com/urfave/cli/v2"
)

type ListArgs struct {
	Key      string
	Revision int
}

var listArgs ListArgs

func NewListCommand() *cli.Command {
	return &cli.Command{
		Name:      "list",
		Usage:     "list options",
		UsageText: myapp.AppName + " [global options] list [Args]",
		Action:    list,
		Flags:     ListFlags,
	}
}

var ListFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "key",
		Aliases:     []string{"k"},
		Usage:       "key to locate",
		Required:    true,
		Destination: &listArgs.Key,
	},
	&cli.IntFlag{
		Name:        "revision",
		Aliases:     []string{"r"},
		Usage:       "revision of store data",
		Value:       0,
		Destination: &listArgs.Revision,
	},
}

func list(*cli.Context) error {
	client, err_init := Config.Client()
	if err_init != nil {
		fmt.Println("fail to init client.")
		return err_init
	}
	defer client.Close()

	values, err := client.List(context.Background(), listArgs.Key, listArgs.Revision)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	for _, value := range values {
		fmt.Printf("key: %s\ndata: %s\nmodified times: %d\n", string(value.Key), string(value.Data), value.Modified)
	}
	return nil
}
