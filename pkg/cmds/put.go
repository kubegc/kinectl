package cmds

import (
	"context"
	"fmt"

	"github.com/litekube/kinectl/pkg/myapp"
	"github.com/urfave/cli/v2"
)

type PutArgs struct {
	Key   string
	Value string
	Force bool
}

var putArgs PutArgs

func NewPutCommand() *cli.Command {
	return &cli.Command{
		Name:      "put",
		Usage:     "put options",
		UsageText: myapp.AppName + " [global options] put [Args]",
		Action:    put,
		Flags:     PutFlags,
	}
}

var PutFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "key",
		Aliases:     []string{"k"},
		Usage:       "key to locate",
		Required:    true,
		Destination: &putArgs.Key,
	},
	&cli.StringFlag{
		Name:        "value",
		Aliases:     []string{"v"},
		Usage:       "value to store",
		Required:    true,
		Destination: &putArgs.Value,
	},
	&cli.BoolFlag{
		Name:        "force",
		Aliases:     []string{"f"},
		Usage:       "if value is not exist, create one",
		Value:       false,
		Destination: &putArgs.Force,
	},
}

func put(*cli.Context) error {
	client, err_init := Config.Client()
	if err_init != nil {
		return fmt.Errorf("fail to init client: %s", err_init.Error())
	}
	defer client.Close()

	err := client.Put(context.TODO(), putArgs.Key, []byte(putArgs.Value), putArgs.Force)
	if err == nil {
		fmt.Println("success")
	}
	return err
}
