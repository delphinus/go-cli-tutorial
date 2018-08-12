package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v2"
)

func main() { NewApp().Run(os.Args) }

// NewApp creates the app
func NewApp() *cli.App {
	return &cli.App{
		Name:    "some app",
		Version: "0.0.1",
		Usage:   "some sample app",
		Action: func(c *cli.Context) error {
			fmt.Fprintln(c.App.Writer, "Hello, World!")
			return nil
		},
	}
}
