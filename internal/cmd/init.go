package cmd

import (
	"poetry/internal/app/model"
	"poetry/internal/pkg/dataset"

	"github.com/urfave/cli/v2"
)

var (
	initCmd = &cli.Command{
		Name:  "init",
		Usage: "Create table and insert data with chinese-poetry",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "path",
				Usage: "chinese poetry path",
				Value: "chinese-poetry",
			},
			&cli.BoolFlag{
				Name:  "insert",
				Usage: "insert poetry from dataset",
			},
		},
		Action: func(clx *cli.Context) error {
			if path := clx.String("path"); path != "" {
				defaultApp.Config.Set("dataset.path", path)
			}
			if err := model.Init(defaultApp.Config, defaultApp.DB.DB); err != nil {
				return err
			}
			if clx.Bool("insert") {
				defaultApp.Config.Set("logger.level", "info")
				return dataset.Init(defaultApp)
			}
			return nil
		},
	}
)
