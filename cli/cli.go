package cli

import (
	"fmt"

  "github.com/mvisonneau/marefo/api"
  "github.com/mvisonneau/marefo/config"

	"github.com/urfave/cli"
)

func init() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Version)
	}
}

// Cli : Generates cli configuration for the application
func Run(version *string) (c *cli.App) {
	c = cli.NewApp()
	c.Name = "marefo"
	c.Version = *version
	c.Usage = "Manage the vulnerability of your running containers images across your Kubernetes clusters"
	c.EnableBashCompletion = true

	c.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "clair-endpoint,c",
			EnvVar:      "MAREFO_CLAIR_ENDPOINT",
			Usage:       "Clair endpoint to use",
			Value:       "localhost:6060",
			Destination: &config.Get().Clair.Endpoint,
		},
		cli.StringFlag{
			Name:        "log-level",
			EnvVar:      "MAREFO_LOG_LEVEL",
			Usage:       "log level (debug,info,warn,fatal,panic)",
			Value:       "info",
			Destination: &config.Get().Log.Level,
		},
		cli.StringFlag{
			Name:        "log-format",
			EnvVar:      "MAREFO_LOG_FORMAT",
			Usage:       "log format (json,text)",
			Value:       "text",
			Destination: &config.Get().Log.Format,
		},
	}

	c.Commands = []cli.Command{
		{
			Name:      "start",
			Usage:     "Start the API server",
			ArgsUsage: " ",
			Action:    api.Run,
		},
	}

	return
}
