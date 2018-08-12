package main

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/urfave/cli.v2"
)

// Version 自体はリリース前に手で書き換える
const Version = "1.0.0"

// GitCommit はビルドじに自動的に取得したい
var GitCommit = ""

// CompileTime も自動的に取得できたら嬉しい
var CompileTime = ""
var compileTime int64

func init() {
	compileTime, _ = strconv.ParseInt(CompileTime, 10, 64)

	cli.VersionPrinter = func(c *cli.Context) {
		rev := GitCommit
		// 空のままビルドした場合は dev という文字列を使う
		if rev == "" {
			rev = "dev"
		}
		compiled := ""
		if c.App.Compiled.Unix() == 0 {
			compiled = "development"
		} else {
			compiled = c.App.Compiled.Format(time.RFC3339)
		}
		fmt.Fprintf(c.App.Writer, "%v version v%v (%s) build time: %s\n",
			c.App.Name, c.App.Version, rev, compiled)
	}
}
