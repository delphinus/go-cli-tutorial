package main

import (
	"errors"
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
		// actionFunc() で自前の関数を包む
		Action: actionFunc(func(c *cli.Context) error {
			// HandleExitCoder() が処理できないはずのエラーを返す
			fmt.Fprintln(c.App.Writer, "Hello, World!")
			return errors.New("some critical error")
		}),
	}
}

func actionFunc(f cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) (err error) {
		if err = f(c); err == nil {
			return
		}
		if _, ok := err.(cli.ExitCoder); ok {
			return
		}
		if _, ok := err.(cli.MultiError); ok {
			return
		}
		return cli.Exit(err, 1)
	}
}
