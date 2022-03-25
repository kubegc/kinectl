package cmds

import (
	"context"
	"fmt"

	"github.com/litekube/kinectl/pkg/myapp"
	"github.com/urfave/cli/v2"
)

type GetArgs struct {
	Key string
}

var getArgs GetArgs

func NewGetCommand() *cli.Command {
	return &cli.Command{
		Name:      "get",
		Usage:     "get options",
		UsageText: myapp.AppName + " [global options] get [Args]",
		Action:    get,
		Flags:     GetFlags,
	}
}

var GetFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "key",
		Usage:       "key to locate",
		Required:    true,
		Destination: &getArgs.Key,
	},
}

func get(*cli.Context) error {
	client, err_init := Config.Client()
	if err_init != nil {
		fmt.Println("fail to init client.")
		return err_init
	}
	defer client.Close()

	value, err := client.Get(context.Background(), getArgs.Key)
	if err != nil {
		return err
	}

	fmt.Printf("key: %s\ndata: %s\nmodified times: %d\n", string(value.Key), string(value.Data), value.Modified)
	return nil
}
