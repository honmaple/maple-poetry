package main

import (
	"embed"
	"poetry/internal/cmd"
)

var (
	//go:embed web/dist/spa/*
	webFS embed.FS
)

func main() {
	cmd.FS = webFS
	cmd.Run()
}
