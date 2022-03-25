package cmds

import (
	"strings"

	"github.com/k3s-io/kine/pkg/endpoint"
	"github.com/k3s-io/kine/pkg/tls"
	clientk "github.com/litekube/kinectl/pkg/client"
)

type ConfigFlags struct {
	Endpoints   string
	LeaderElect bool
	CAFile      string
	CertFile    string
	KeyFile     string
}

func (config ConfigFlags) ToETCDConfig() endpoint.ETCDConfig {
	endpoints := make([]string, 0)
	for _, endpoint := range strings.Split(config.Endpoints, ",") {
		endpoints = append(endpoints, strings.TrimSpace(endpoint))
	}

	return endpoint.ETCDConfig{
		Endpoints:   endpoints,
		LeaderElect: config.LeaderElect,
		TLSConfig: tls.Config{
			CAFile:   config.CAFile,
			CertFile: config.CertFile,
			KeyFile:  config.KeyFile,
		},
	}
}

func (config ConfigFlags) Client() (clientk.Client, error) {
	return clientk.New(config.ToETCDConfig())
}
