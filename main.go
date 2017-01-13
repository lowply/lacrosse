package main

import "os"

func main() {
	cli := NewCLI()
	os.Exit(cli.Run(os.Args))
}
