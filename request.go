package main

import (
	"os"
	"strconv"
	"time"
)

type Request struct {
	Date    string
	Action  string
	Domain  string
	Type    string
	Value   string
	TTL     int64
	Profile string
}

func NewRequest() (*Request, error) {
	ttli64, err := strconv.ParseInt(os.Args[4], 10, 64)
	if err != nil {
		return nil, err
	}

	req := &Request{
		Date:    time.Now().Format("2006/01/02 15:04:05 MST"),
		Action:  "UPSERT",
		Domain:  os.Args[1],
		Type:    os.Args[2],
		Value:   os.Args[3],
		TTL:     ttli64,
		Profile: os.Args[5],
	}

	return req, nil
}
