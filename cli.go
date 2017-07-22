package main

import (
	"fmt"
	"io"
	"os"
)

const VERSION = "0.2.4"

type CLI struct {
	out io.Writer
	err io.Writer
}

func NewCLI() *CLI {
	return &CLI{
		os.Stdout,
		os.Stderr,
	}
}

func (c *CLI) usage() {
	msg := "lacrosse version " + VERSION + "\n"
	msg += "Usage : lacrosse [domain] [type] [value] [TTL] [aws profile]"
	fmt.Fprintln(c.err, msg)
}

func (c *CLI) Run(args []string) int {
	if len(os.Args) != 6 {
		c.usage()
		return 1
	}

	r, err := NewRoute53(os.Args)
	if err != nil {
		fmt.Fprintln(c.err, err)
		return 1
	}

	err = r.RequestChange()
	if err != nil {
		fmt.Fprintln(c.err, err)
		return 1
	}

	return 0
}
