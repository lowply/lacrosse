package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
)

var profile string

func abort(e error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", e.Error())
	os.Exit(1)
}

func usage() {
	abort(errors.New("Usage : lacrosse [domain] [type] [value] [TTL] [awscli profile]"))
}

func has(command string) {
	_, err := exec.LookPath(command)
	if err != nil {
		abort(err)
	}
}

func check_dir(dirname string) {
	_, err := os.Stat(dirname)

	if err != nil {
		err := os.Mkdir(dirname, 0777)
		if err != nil {
			abort(err)
		}
		fmt.Println(dirname + " has been created")
	}
}

func main() {
	check_dir(path.Dir(logpath))

	if len(os.Args) != 6 {
		usage()
	}

	has("host")
	has("aws")

	req, err := NewRequest(os.Args)
	if err != nil {
		abort(err)
	}

	r, err := NewRoute53(req)
	if err != nil {
		abort(err)
	}

	resp, err := r.RequestChange()
	if err != nil {
		abort(err)
	}

	r.CheckStatus(resp)

	r.Logger()
}
