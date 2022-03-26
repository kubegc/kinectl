package cmds

import (
	"fmt"

	"github.com/litekube/kinectl/pkg/myapp"
	"github.com/urfave/cli/v2"
)

func NewVersionCommand() *cli.Command {
	return &cli.Command{
		Name:      "version",
		Usage:     "view the version of Kine",
		UsageText: myapp.AppName + " [global options] version",
		Action: func(c *cli.Context) error {
			client, err_init := Config.Client()
			if err_init != nil {
				return fmt.Errorf("fail to init client: %s", err_init.Error())
			}
			defer client.Close()

			data, err := client.Version()
			if err != nil {
				return err
			}
			fmt.Println(data)
			return nil
		},
	}
}
