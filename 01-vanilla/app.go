package main

import (
	"fmt"
	"os"
)

func main() {
	if err := process(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
func process(args []string) error {
	// 処理の中身はここ
	fmt.Println("Hello, World!")
	return nil
}
