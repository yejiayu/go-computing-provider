package main

import (
	"github.com/swanchain/go-computing-provider/build"
	"github.com/urfave/cli/v2"
	"os"
)

const (
	FlagCpRepo = "repo"
)

var FlagRepo = &cli.StringFlag{
	Name:    FlagCpRepo,
	Usage:   "repo directory for computing-provider client",
	Value:   "~/.swan/computing",
	EnvVars: []string{"CP_PATH"},
}

func main() {
	app := &cli.App{
		Name:                 "computing-provider",
		Usage:                "Swanchain decentralized computing network client",
		EnableBashCompletion: true,
		Version:              build.UserVersion(),
		Flags: []cli.Flag{
			FlagRepo,
		},
		Commands: []*cli.Command{
			initCmd,
			runCmd,
			infoCmd,
			accountCmd,
			taskCmd,
			walletCmd,
			collateralCmd,
			ubiTaskCmd,
		},
	}
	app.Setup()

	if err := app.Run(os.Args); err != nil {
		os.Stderr.WriteString("Error: " + err.Error() + "\n")
	}
}
