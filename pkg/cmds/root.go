package cmds

import (
	"fmt"
	"runtime"

	"github.com/litekube/kinectl/pkg/myapp"
	"github.com/litekube/kinectl/pkg/version"
	"github.com/urfave/cli/v2"
)

var (
	Config ConfigFlags
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = myapp.AppName
	app.Usage = "kinectl, a commond-line tool to use kine database."
	app.Version = version.Version
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version %s\n", app.Name, app.Version)
		fmt.Printf("go version %s\n", runtime.Version())
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "cacert",
			EnvVars:     []string{"KINECTL_CACERT"},
			Usage:       "TLS CA certificate to connect to kine",
			Destination: &Config.CAFile,
		},
		&cli.StringFlag{
			Name:        "cert",
			EnvVars:     []string{"KINECTL_CERT"},
			Usage:       "TLS client certificate to connect to kine",
			Destination: &Config.CertFile,
		},
		&cli.StringFlag{
			Name:        "key",
			EnvVars:     []string{"KINECTL_KEY"},
			Usage:       "TLS client key to connect to kine",
			Destination: &Config.KeyFile,
		},
		&cli.BoolFlag{
			Name:        "leader-elect",
			EnvVars:     []string{"KINECTL_LEADELECT"},
			Usage:       "whether kine database enable leader-elect or not",
			Value:       false,
			Destination: &Config.LeaderElect,
		},
		&cli.StringFlag{
			Name:        "endpoints",
			EnvVars:     []string{"KINECTL_LEADELECT"},
			Usage:       "kine database endpoint, like \"www.example1.com:2379,www.example2.com:2379\"",
			Destination: &Config.Endpoints,
		},
	}

	return app
}
