package main

import (
	"fmt"
	"os"

	"github.com/OrlandoRomo/imgfltr/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "imgfltr"
	app.Usage = "CLI to apply different image filters"
	app.Description = `imgfltr is a CLI tool to apply a local picture o 
	image with the format (.png, .jpg, .jpeg) different filters such as grey scale, inverted colors, sepia etc`

	app.Commands = []*cli.Command{
		cmd.NewSetCommand(),
		cmd.NewListCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}
