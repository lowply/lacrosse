package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"time"
)

var logpath = os.Getenv("HOME") + "/.cache/lacrosse.log"

type Request struct {
	Date    string
	Id      string
	Action  string
	Domain  string
	Type    string
	Value   string
	TTL     int64
	Profile string
}

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

func write_log(r *Request) error {
	check_dir(path.Dir(logpath))
	logfile, err := os.OpenFile(logpath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	b, err := json.Marshal(r)
	if err != nil {
		return err
	}

	logfile.Write(b)
	logfile.WriteString("\n")
	return nil
}

func create_request() (*Request, error) {
	r := &Request{
		Date:    time.Now().Format("2006/01/02 15:04:05 MST"),
		Action:  "UPSERT",
		Domain:  os.Args[1],
		Type:    os.Args[2],
		Value:   os.Args[3],
		Profile: os.Args[5],
	}

	svc = create_new_service(r.Profile)

	ttl, err := strconv.ParseInt(os.Args[4], 10, 64)
	if err != nil {
		return nil, err
	}
	r.TTL = ttl

	id, err := get_hosted_zone_id(r.Domain)
	if err != nil {
		return nil, err
	}
	r.Id = id
	return r, nil
}

func main() {
	if len(os.Args) != 6 {
		usage()
	}

	has("host")
	has("aws")

	r, err := create_request()
	if err != nil {
		abort(err)
	}

	params, err := create_new_params(r)
	if err != nil {
		abort(err)
	}

	fmt.Println("New param created: \n\n" + params.GoString() + "\n")

	resp, err := svc.ChangeResourceRecordSets(params)
	if err != nil {
		abort(err)
	}
	fmt.Println("Got response for the request: \n\n" + resp.GoString() + "\n")

	check_status(resp.ChangeInfo)

	err = write_log(r)
	if err != nil {
		abort(err)
	}
}
