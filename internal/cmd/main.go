package cmd

import (
	"embed"
	"fmt"
	"io/fs"
	"os"

	"poetry/internal/app"

	"github.com/urfave/cli/v2"
)

const (
	PROCESS     = "poetry"
	DESCRIPTION = "chinese poetry"
	VERSION     = "0.1.0"
)

var (
	FS         embed.FS
	defaultApp = app.New()
)

func before(clx *cli.Context) error {
	return defaultApp.Init(clx.String("config"), clx.StringSlice("set-config")...)
}

func action(clx *cli.Context) error {
	if addr := clx.String("addr"); addr != "" {
		defaultApp.Config.Set("server.addr", addr)
	}
	if debug := clx.Bool("debug"); debug {
		defaultApp.Config.Set("server.mode", "dev")
	}

	webFS, err := fs.Sub(FS, "web/dist/spa")
	if err != nil {
		return err
	}
	return defaultApp.Run(webFS)
}

func Run() {
	app := &cli.App{
		Name:    PROCESS,
		Usage:   DESCRIPTION,
		Version: VERSION,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"D"},
				Usage:   "debug mode",
			},
			&cli.StringFlag{
				Name:    "addr",
				Aliases: []string{"a"},
				Usage:   "listen `ADDR`",
			},
			&cli.PathFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "load config from `FILE`",
				Value:   "config.yaml",
			},
			&cli.StringSliceFlag{
				Name:  "set-config",
				Usage: "set config from string",
			},
		},
		Before: before,
		Action: action,
		Commands: []*cli.Command{
			initCmd,
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
	}
}
