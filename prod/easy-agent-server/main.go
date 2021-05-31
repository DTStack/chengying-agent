package main

import (
	"fmt"
	"os"

	"easyagent/internal"
	"easyagent/internal/server/base"
	"easyagent/internal/server/log"
	"github.com/urfave/cli"
)

func main() {
	fmt.Println("EasyAgent Server " + internal.VERSION)
	fmt.Println("Copyright (c) 2017 DTStack Inc.")
	base.ConfigureProductVersion(internal.VERSION)

	app := cli.NewApp()
	app.Version = internal.VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config,c",
			Usage: "config path",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "debug info",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		log.SetDebug(ctx.Bool("debug"))
		if err := ParseConfig(ctx.String("config")); err != nil {
			return err
		}

		return base.Run()
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exit with failure: %v\n", err)
	}
}
