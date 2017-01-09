package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
)

var logpath = os.Getenv("HOME") + "/.cache/lacrosse.log"
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

func write_log(r *Route53) error {
	check_dir(path.Dir(logpath))
	logfile, err := os.OpenFile(logpath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	b, err := json.Marshal(r.Req)
	if err != nil {
		return err
	}

	logfile.Write(b)
	logfile.WriteString("\n")
	return nil
}

func main() {
	if len(os.Args) != 6 {
		usage()
	}

	has("host")
	has("aws")

	req, err := NewRequest()
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

	err = write_log(r)
	if err != nil {
		abort(err)
	}
}
