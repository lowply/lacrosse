package main

import (
	"fmt"
	"io"
	"os"
)

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

func (c *CLI) Run(args []string) int {
	if len(os.Args) != 6 {
		fmt.Fprintln(c.err, "Usage : lacrosse [domain] [type] [value] [TTL] [aws profile]")
		return 1
	}

	req, err := NewRequest(os.Args[1:])
	if err != nil {
		fmt.Fprintln(c.err, err)
		return 1
	}

	r, err := NewRoute53(req)
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
